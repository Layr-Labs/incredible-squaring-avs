package aggregator

import (
	"context"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "incredible-squaring"
)

// Aggregator sends tasks (numbers to square) onchain, then listens for operator signed TaskResponses.
// It aggregates responses signatures, and if any of the TaskResponses reaches the QuorumThresholdPercentage for each quorum
// (currently we only use a single quorum of the ERC20Mock token), it sends the aggregated TaskResponse and signature onchain.
//
// The signature is checked in the BLSSignatureChecker.sol contract, which expects a
//
//	struct NonSignerStakesAndSignature {
//		uint32[] nonSignerQuorumBitmapIndices;
//		BN254.G1Point[] nonSignerPubkeys;
//		BN254.G1Point[] quorumApks;
//		BN254.G2Point apkG2;
//		BN254.G1Point sigma;
//		uint32[] quorumApkIndices;
//		uint32[] totalStakeIndices;
//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
//	}
//
// A task can only be responded onchain by having enough operators sign on it such that their stake in each quorum reaches the QuorumThresholdPercentage.
// In order to verify this onchain, the Registry contracts store the history of stakes and aggregate pubkeys (apks) for each operators and each quorum. These are
// updated everytime an operator registers or deregisters with the BLSRegistryCoordinatorWithIndices.sol contract, or calls UpdateStakes() on the StakeRegistry.sol contract,
// after having received new delegated shares or having delegated shares removed by stakers queuing withdrawals. Each of these pushes to their respective datatype array a new entry.
//
// This is true for quorumBitmaps (represent the quorums each operator is opted into), quorumApks (apks per quorum), totalStakes (total stake per quorum), and nonSignerStakes (stake per quorum per operator).
// The 4 "indices" in NonSignerStakesAndSignature basically represent the index at which to fetch their respective data, given a blockNumber at which the task was created.
// Note that different data types might have different indices, since for eg QuorumBitmaps are updated for operators registering/deregistering, but not for UpdateStakes.
// Thankfully, we have deployed a helper contract BLSOperatorStateRetriever.sol whose function getCheckSignaturesIndices() can be used to fetch the indices given a block number.
//
// The 4 other fields nonSignerPubkeys, quorumApks, apkG2, and sigma, however, must be computed individually.
// apkG2 and sigma are just the aggregated signature and pubkeys of the operators who signed the task response (aggregated over all quorums, so individual signatures might be duplicated).
// quorumApks are the G1 aggregated pubkeys of the operators who signed the task response, but one per quorum, as opposed to apkG2 which is summed over all quorums.
// nonSignerPubkeys are the G1 pubkeys of the operators who did not sign the task response, but were opted into the quorum at the blocknumber at which the task was created.
// Upon sending a task onchain (or receiving a NewTaskCreated Event if the tasks were sent by an external task generator), the aggregator can get the list of all operators opted into each quorum at that
// block number by calling the getOperatorState() function of the BLSOperatorStateRetriever.sol contract.
type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        chainio.AvsWriterer
	// aggregation related fields
	// blsAggregationService blsagg.BlsAggregationService
	tasks           map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	tasksMu         sync.RWMutex
	taskResponses   map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	taskResponsesMu sync.RWMutex
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c *config.Config) (*Aggregator, error) {

	// avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	// if err != nil {
	// 	c.Logger.Error("Cannot create avsReader", "err", err)
	// 	return nil, err
	// }

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	// chainioConfig := sdkclients.BuildAllConfig{
	// 	EthHttpUrl:                 c.EthHttpRpcUrl,
	// 	EthWsUrl:                   c.EthWsRpcUrl,
	// 	RegistryCoordinatorAddr:    c.IncredibleSquaringRegistryCoordinatorAddr.String(),
	// 	OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
	// 	AvsName:                    avsName,
	// 	PromMetricsIpPortAddress:   ":9090",
	// }
	// clients, err := clients.BuildAll(chainioConfig, c.EcdsaPrivateKey, c.Logger)
	// if err != nil {
	// 	c.Logger.Errorf("Cannot create sdk clients", "err", err)
	// 	return nil, err
	// }

	// operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	// avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	// blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	return &Aggregator{
		logger:           c.Logger,
		serverIpPortAddr: c.AggregatorServerIpPortAddr,
		avsWriter:        avsWriter,
		// blsAggregationService: blsAggregationService,
		tasks:         make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses: make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse),
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting aggregator.")
	agg.logger.Infof("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	// TODO(soubhik): refactor task generation/sending into a separate function that we can run as goroutine
	ticker := time.NewTicker(10 * time.Second)
	agg.logger.Infof("Aggregator set to send new task every 10 seconds...")
	defer ticker.Stop()
	taskNum := int64(0)
	// ticker doesn't tick immediately, so we send the first task here
	// see https://github.com/golang/go/issues/17601
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		}
		// case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
		// 	agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
		// 	agg.sendAggregatedResponseToContract(blsAggServiceResp)
	}
}
