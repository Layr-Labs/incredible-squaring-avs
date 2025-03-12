package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/exp/rand"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	sdkcommon "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/metrics"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const AVS_NAME = "incredible-squaring"
const SEM_VER = "0.0.1"

type Operator struct {
	Config    types.NodeConfig
	Logger    sdklogging.Logger
	EthClient sdkcommon.EthClientInterface
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	MetricsReg       *prometheus.Registry
	Metrics          metrics.Metrics
	NodeApi          *nodeapi.NodeApi
	AvsWriter        *chainio.AvsWriter
	AvsReader        chainio.AvsReaderer
	AvsSubscriber    chainio.AvsSubscriberer
	EigenlayerReader sdkelcontracts.ChainReader
	EigenlayerWriter sdkelcontracts.ChainWriter
	BlsKeypair       *bls.KeyPair
	OperatorId       sdktypes.OperatorId
	OperatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	NewTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	// ip address of aggregator
	AggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	AggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	CredibleSquaringServiceManagerAddr common.Address
	// If bigger than zero, submits wrong responses that many times every 100
	TimesFailing int
}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {

	var logLevel sdklogging.LogLevel
	if c.Production {
		logLevel = sdklogging.Production
	} else {
		logLevel = sdklogging.Development
	}
	logger, err := sdklogging.NewZapLogger(logLevel)
	if err != nil {
		return nil, err
	}
	reg := prometheus.NewRegistry()
	eigenMetrics := sdkmetrics.NewEigenMetrics(AVS_NAME, c.EigenMetricsIpPortAddress, reg, logger)
	avsAndEigenMetrics := metrics.NewAvsAndEigenMetrics(AVS_NAME, eigenMetrics, reg)

	// Setup Node Api
	nodeApi := nodeapi.NewNodeApi(AVS_NAME, SEM_VER, c.NodeApiIpPortAddress, logger)

	var ethRpcClient, ethWsClient sdkcommon.EthClientInterface
	if c.EnableMetrics {
		rpcCallsCollector := rpccalls.NewCollector(AVS_NAME, reg)
		ethRpcClient, err = eth.NewInstrumentedClient(c.EthRpcUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewInstrumentedClient(c.EthWsUrl, rpcCallsCollector)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	} else {
		ethRpcClient, err = ethclient.Dial(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = ethclient.Dial(c.EthWsUrl)
		if err != nil {
			logger.Errorf("Cannot create ws ethclient", "err", err)
			return nil, err
		}
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return nil, err
	}
	// TODO(samlaf): should we add the chainId to the config instead?
	// this way we can prevent creating a signer that signs on mainnet by mistake
	// if the config says chainId=5, then we can only create a goerli signer
	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: c.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	chainioConfig := clients.BuildAllConfig{
		EthHttpUrl:                  c.EthRpcUrl,
		EthWsUrl:                    c.EthWsUrl,
		RegistryCoordinatorAddr:     c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr:  c.OperatorStateRetrieverAddress,
		ServiceManagerAddress:       c.IncredibleSquaringServiceManager,
		RewardsCoordinatorAddress:   c.RewardsCoordinatorAddress,
		PermissionControllerAddress: c.PermissionControllerAddress,
		AvsName:                     AVS_NAME,
		PromMetricsIpPortAddress:    c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr,
		common.HexToAddress(c.IncredibleSquaringServiceManager),
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethWsClient,
		ethRpcClient,
		logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.IncredibleSquaringServiceManager),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		sdkClients.EthWsClient,
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriber(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.IncredibleSquaringServiceManager),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethWsClient,
		logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph
	// calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
	}

	operator := &Operator{
		Config:                     c,
		Logger:                     logger,
		MetricsReg:                 reg,
		Metrics:                    avsAndEigenMetrics,
		NodeApi:                    nodeApi,
		EthClient:                  ethRpcClient,
		AvsWriter:                  avsWriter,
		AvsReader:                  avsReader,
		AvsSubscriber:              avsSubscriber,
		EigenlayerReader:           *sdkClients.ElChainReader,
		EigenlayerWriter:           *sdkClients.ElChainWriter,
		BlsKeypair:                 blsKeyPair,
		OperatorAddr:               common.HexToAddress(c.OperatorAddress),
		AggregatorServerIpPortAddr: c.AggregatorServerIpPortAddress,
		AggregatorRpcClient:        aggregatorRpcClient,
		NewTaskCreatedChan: make(
			chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated,
		),
		CredibleSquaringServiceManagerAddr: common.HexToAddress(c.IncredibleSquaringServiceManager),
		OperatorId:                         [32]byte{0}, // this is set below
		TimesFailing:                       c.TimesFailing,
	}
	operatorSetsIds := []uint32{c.OperatorSetId}
	waitForReceipt := true
	operator.SetAppointee(
		common.HexToAddress(c.InstantSlasher),
		operator.CredibleSquaringServiceManagerAddr,
		common.HexToAddress(c.AllocationManagerAddress),
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
	)
	operator.CreateTotalDelegatedStakeQuorum(
		c.MaxOperatorCount,
		c.KickBIPsOfOperatorStake,
		c.KickBIPsOfTotalStake,
		c.MinimumStake,
		c.Multiplier,
	)

	if c.RegisterOperatorOnStartup {
		operator.registerOperatorOnStartup(
			operatorEcdsaPrivateKey,
			common.HexToAddress(c.TokenStrategyAddr),
			common.HexToAddress(c.AVSRegistryCoordinatorAddress),
			common.HexToAddress(c.IncredibleSquaringServiceManager),
			operatorSetsIds,
			waitForReceipt,
			*operator.BlsKeypair,
			c.Socket,
		)
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.OperatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.OperatorId = operatorId
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.BlsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.BlsKeypair.GetPubKeyG2(),
	)
	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.AvsReader.IsOperatorRegistered(&bind.CallOpts{}, o.OperatorAddr)
	if err != nil {
		o.Logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack
		// trace that hides the actual error message. This error msg is more explicit and doesn't require showing a
		// stack trace to the user.
		return fmt.Errorf(
			"operator is not registered. Registering operator using the operator-cli before starting operator",
		)
	}

	o.Logger.Infof("Starting operator.")

	if o.Config.EnableNodeApi {
		o.NodeApi.Start()
	}
	var metricsErrChan <-chan error
	if o.Config.EnableMetrics {
		metricsErrChan = o.Metrics.Start(ctx, o.MetricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.AvsSubscriber.SubscribeToNewTasks(o.NewTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.Logger.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.Logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.AvsSubscriber.SubscribeToNewTasks(o.NewTaskCreatedChan)
		case newTaskCreatedLog := <-o.NewTaskCreatedChan:
			o.Metrics.IncNumTasksReceived()
			taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.AggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(
	newTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated,
) *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse {
	o.Logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.Logger.Info("Received new task",
		"numberToBeSquared", newTaskCreatedLog.Task.NumberToBeSquared,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)
	numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)

	if o.TimesFailing > 0 {
		rand.Seed(uint64((time.Now().UnixNano())))
		num := rand.Intn(100)
		if num < o.TimesFailing {
			numberSquared = big.NewInt(908243203843)
		}
	}
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		NumberSquared:      numberSquared,
	}
	return taskResponse
	// numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)
}

func (o *Operator) SignTaskResponse(
	taskResponse *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
) (*aggregator.SignedTaskResponse, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.Logger.Error(
			"Error getting task response header hash. skipping task (this is not expected and should be investigated)",
			"err",
			err,
		)
		return nil, err
	}
	blsSignature := o.BlsKeypair.SignMessage(taskResponseHash)
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.OperatorId,
	}
	o.Logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
