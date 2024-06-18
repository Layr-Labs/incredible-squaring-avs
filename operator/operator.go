package operator

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	priceFeedAdapter "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/PriceFeedAdapter"
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
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	AVS_NAME                 = "incredible-squaring"
	SEM_VER                  = "0.0.1"
)

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.Client
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg       *prometheus.Registry
	metrics          metrics.Metrics
	nodeApi          *nodeapi.NodeApi
	avsWriter        *chainio.AvsWriter
	avsReader        chainio.AvsReaderer
	avsSubscriber    chainio.AvsSubscriberer
	eigenlayerReader sdkelcontracts.ELReader
	eigenlayerWriter sdkelcontracts.ELWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan            chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	newPriceUpdateTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerPriceUpdateRequested

	// needed when opting in to avs (allow this service manager contract to slash operator)
	credibleSquaringServiceManagerAddr common.Address

	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask
	tasksMu               sync.RWMutex
	taskResponses         map[uint32]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse
	taskDigestQuorum      map[uint32]TaskDigestQuorum
	taskResponsesMu       sync.RWMutex

	// needed to fetch the price of assets on different on-chain oracle networks
	priceFeedAdapter *priceFeedAdapter.ContractPriceFeedAdapter

	priceFSM *PriceFSM
}

type PriceUpdateRequest struct {
	FeedName  string
	TaskId    uint32
	LeaderUrl string
}

type PriceUpdateTaskResponse struct {
	Price    uint32
	Decimals uint32
	TaskId   uint32
	Source   string
}

type TaskDigestQuorum struct {
	finalTaskResponse           cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature
}

// TODO(samlaf): config is a mess right now, since the chainio client constructors
//
//	take the config in core (which is shared with aggregator and challenger)
func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	var logLevel logging.LogLevel
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

	var ethRpcClient, ethWsClient eth.Client
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
		ethRpcClient, err = eth.NewClient(c.EthRpcUrl)
		if err != nil {
			logger.Errorf("Cannot create http ethclient", "err", err)
			return nil, err
		}
		ethWsClient, err = eth.NewClient(c.EthWsUrl)
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
		EthHttpUrl:                 c.EthRpcUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress,
		AvsName:                    AVS_NAME,
		PromMetricsIpPortAddress:   c.EigenMetricsIpPortAddress,
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		c.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return nil, err
	}
	sdkClients, err := clients.BuildAll(chainioConfig, operatorEcdsaPrivateKey, logger)
	sdkClients.ElChainReader.IsOperatorRegistered(&bind.CallOpts{}, sdktypes.Operator{Address: c.OperatorAddress})
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, common.HexToAddress(c.OperatorAddress), logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, common.HexToAddress(c.OperatorAddress))

	avsWriter, err := chainio.BuildAvsWriter(
		txMgr, common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethRpcClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsWriter", "err", err)
		return nil, err
	}

	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress),
		ethRpcClient, logger)
	if err != nil {
		logger.Error("Cannot create AvsReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriber(common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		common.HexToAddress(c.OperatorStateRetrieverAddress), ethWsClient, logger,
	)
	if err != nil {
		logger.Error("Cannot create AvsSubscriber", "err", err)
		return nil, err
	}

	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	quorumNames := map[sdktypes.QuorumNum]string{
		0: "quorum0",
	}
	economicMetricsCollector := economic.NewCollector(
		sdkClients.ElChainReader, sdkClients.AvsRegistryChainReader,
		AVS_NAME, logger, common.HexToAddress(c.OperatorAddress), quorumNames)
	reg.MustRegister(economicMetricsCollector)

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(context.Background(), sdkClients.AvsRegistryChainSubscriber, sdkClients.AvsRegistryChainReader, logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, logger)

	priceFeedClient, err := priceFeedAdapter.NewContractPriceFeedAdapter(common.HexToAddress(c.PriceFeedAdapterAddress), ethRpcClient)
	if err != nil {
		logger.Error("Cannot create priceFeedClient. Is aggregator running?", "err", err)
		return nil, err
	}

	// setup raft consensus client
	consensusFSM := NewConcensusFSM(priceFeedClient, blsKeyPair, operatorEcdsaPrivateKey)

	taskResponses := make(map[uint32]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse)

	// start http server with additional raft endpoints
	h := NewService(c.HttpBindingURI, consensusFSM, blsAggregationService, &taskResponses, ethRpcClient)
	h.avsReader = avsReader
	h.logger = logger
	if err := h.Start(); err != nil {
		logger.Error("failed to start HTTP service: %s", err.Error())
	}

	operator := &Operator{
		config:                             c,
		logger:                             logger,
		metricsReg:                         reg,
		metrics:                            avsAndEigenMetrics,
		nodeApi:                            nodeApi,
		ethClient:                          ethRpcClient,
		avsWriter:                          avsWriter,
		avsReader:                          avsReader,
		avsSubscriber:                      avsSubscriber,
		eigenlayerReader:                   sdkClients.ElChainReader,
		eigenlayerWriter:                   sdkClients.ElChainWriter,
		blsKeypair:                         blsKeyPair,
		blsAggregationService:              blsAggregationService,
		operatorAddr:                       common.HexToAddress(c.OperatorAddress),
		newTaskCreatedChan:                 make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
		newPriceUpdateTaskCreatedChan:      make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerPriceUpdateRequested),
		credibleSquaringServiceManagerAddr: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                         [32]byte{0}, // this is set below
		priceFeedAdapter:                   priceFeedClient,
		priceFSM:                           consensusFSM,
		tasks:                              make(map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask),
		taskResponses:                      taskResponses,
		taskDigestQuorum:                   make(map[uint32]TaskDigestQuorum),
	}

	if c.RegisterOperatorOnStartup {
		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey, common.HexToAddress(c.TokenStrategyAddr))
		avsWriter.ResigterOperatorUrl(context.Background(), c.HttpBindingURI, c.RaftBindingURI)
		time.Sleep(5 * time.Second) // Ensure operator is registered
	}

	// Setup raft
	consensusFSM.RaftBind = c.RaftBindingURI
	consensusFSM.RaftDir = c.RaftDirectoryPath
	consensusFSM.RaftHttpBind = c.HttpBindingURI

	// Check for past operators via OperatorRegistered event on ServiceManager
	results, err := avsReader.GetRegistedOperatorUrls(context.Background())
	results.Next() // First log is always empty so skip it https://stackoverflow.com/questions/62831835/go-ethereum-ethclient-cannot-get-event-logs-data

	if err != nil {
		logger.Error("Failed to load urls", "err", err)
	}

	// If operator is joining an existing raft network make request to join
	// iterate over up to 10 urls
	hasJoinedCluster := false
	hasError := false

	for i := 0; i < 10; i++ {
		event := results.Event

		blockNumber, errr := ethRpcClient.BlockNumber(context.Background())
		if errr != nil {
			logger.Error("Cannot get blockNumber", "err", errr)
			return nil, errr
		}

		logger.Info("Latest block number", "block", blockNumber)

		// check if operatorId is different from this operator
		if event.OperatorId.String() == common.HexToAddress(c.OperatorAddress).String() {
			results.Next()
			continue
		} else {
			url := event.HttpUrl

			logger.Info("Attempting to join existing raft cluster", "joinUrl", url)

			// initialize raft consensus server
			consensusFSM.Initialize(false, c.OperatorAddress)

			if err := consensusFSM.JoinExistingNetwork(url, c.RaftBindingURI, c.OperatorAddress, blockNumber); err != nil {
				logger.Warn("failed to join node at", "url", url, "err", err)
				hasError = true
				results.Next()
			}
			// Joined network
			hasJoinedCluster = true
			break
		}

	}

	if !hasJoinedCluster && !hasError {
		consensusFSM.Initialize(true, c.OperatorAddress)
		logger.Info("Attempting to bootstrap raft cluster")
	} else if !hasJoinedCluster && hasError {
		logger.Error("Failed to join existing raft cluster")
		return nil, err
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
	operator.priceFSM.setOperatorId(operatorId)
	logger.Info("Operator info",
		"operatorId", operatorId,
		"operatorAddr", c.OperatorAddress,
		"operatorG1Pubkey", operator.blsKeypair.GetPubKeyG1(),
		"operatorG2Pubkey", operator.blsKeypair.GetPubKeyG2(),
	)

	return operator, nil

}

func (o *Operator) Start(ctx context.Context) error {
	operatorIsRegistered, err := o.avsReader.IsOperatorRegistered(&bind.CallOpts{}, o.operatorAddr)
	if err != nil {
		o.logger.Error("Error checking if operator is registered", "err", err)
		return err
	}
	if !operatorIsRegistered {
		// We bubble the error all the way up instead of using logger.Fatal because logger.Fatal prints a huge stack trace
		// that hides the actual error message. This error msg is more explicit and doesn't require showing a stack trace to the user.
		return fmt.Errorf("operator is not registered. Registering operator using the operator-cli before starting operator")
	}

	o.logger.Infof("Starting operator.")

	if o.config.EnableNodeApi {
		o.nodeApi.Start()
	}
	var metricsErrChan <-chan error
	if o.config.EnableMetrics {
		metricsErrChan = o.metrics.Start(ctx, o.metricsReg)
	} else {
		metricsErrChan = make(chan error, 1)
	}

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.avsSubscriber.SubscribeToNewPriceUpdateTask(o.newPriceUpdateTaskCreatedChan)
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-metricsErrChan:
			// TODO(samlaf); we should also register the service as unhealthy in the node api
			// https://eigen.nethermind.io/docs/spec/api/
			o.logger.Fatal("Error in metrics server", "err", err)
		case err := <-sub.Err():
			o.logger.Error("Error in websocket subscription", "err", err)
			// TODO(samlaf): write unit tests to check if this fixed the issues we were seeing
			sub.Unsubscribe()
			// TODO(samlaf): wrap this call with increase in avs-node-spec metric
			sub = o.avsSubscriber.SubscribeToNewPriceUpdateTask(o.newPriceUpdateTaskCreatedChan)
		case newPriceUpdateTaskCreatedLog := <-o.newPriceUpdateTaskCreatedChan:
			o.metrics.IncNumTasksReceived()
			o.ProcessNewPriceUpdateCreatedLog(newPriceUpdateTaskCreatedLog)
			if err != nil {
				continue
			}
		case blsAggServiceResp := <-o.blsAggregationService.GetResponseChannel():
			o.logger.Info("Received response from blsAggregationService")
			o.sendAggregatedTaskResponseToContract(blsAggServiceResp)
		case <-ticker.C:
			err := o.sendNewTask()
			if err != nil {
				continue
			}
		}
	}
}

func (o *Operator) ProcessNewPriceUpdateCreatedLog(newPriceUpdateTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerPriceUpdateRequested) error {
	// If not leader ignore request
	isLeader, _ := o.priceFSM.IsLeader()

	if !isLeader {
		o.logger.Info("Waiting for leader request for feed",
			"feedName", string(newPriceUpdateTaskCreatedLog.Task.FeedName[:]),
			"taskCreatedBlock", newPriceUpdateTaskCreatedLog.Task.TaskCreatedBlock,
		)

		return nil
	}

	o.logger.Info("Received new price updated task, sending request to followers",
		"feedName", string(newPriceUpdateTaskCreatedLog.Task.FeedName[:]),
		"taskCreatedBlock", newPriceUpdateTaskCreatedLog.Task.TaskCreatedBlock,
		"minNumberOfOracleNetworks", newPriceUpdateTaskCreatedLog.Task.MinNumOfOracleNetworks,
		"quorumThreshholdPercentage", newPriceUpdateTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	data := &PriceUpdateRequest{
		FeedName:  string(newPriceUpdateTaskCreatedLog.Task.FeedName[:]),
		TaskId:    newPriceUpdateTaskCreatedLog.TaskIndex,
		LeaderUrl: o.priceFSM.RaftHttpBind, // This is the url of the leader making the request
	}

	dataBytes, err := json.Marshal(data)

	if err != nil {
		o.logger.Error("Failed to request task", "err", err)
	}

	// Leader sends request to follower operators
	o.priceFSM.raft.Apply(dataBytes, raftTimeout)

	o.logger.Info("Task request sent to followers")
	return err
}

func (o *Operator) sendNewTask() error {
	// If not leader ignore request
	isLeader, _ := o.priceFSM.IsLeader()

	if !isLeader {
		return nil // Only leader create task
	}

	o.logger.Info("Sending new task")
	// Send number to square to the task manager contract
	newTask, taskIndex, err := o.avsWriter.SendNewPriceUpdate(context.Background())
	if err != nil {
		o.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}

	o.tasksMu.Lock()
	o.tasks[taskIndex] = newTask
	o.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}

	o.logger.Info("Initializing task", "taskId", taskIndex)
	o.blsAggregationService.InitializeNewTask(taskIndex, newTask.TaskCreatedBlock, quorumNums, quorumThresholdPercentages, taskTimeToExpiry)
	return nil
}

// Triggered whenever a taskDigest (ie. hash of {price, taskId, source}) is at quorom
func (o *Operator) sendAggregatedTaskResponseToContract(blsAggServiceResp blsagg.BlsAggregationServiceResponse) {
	// If not leader ignore request
	isLeader, _ := o.priceFSM.IsLeader()

	if !isLeader {
		return // Only leader can submit aggregate task
	}

	if blsAggServiceResp.Err != nil {
		o.logger.Error("BlsAggregationServiceResponse contains an error", "err", blsAggServiceResp.Err)
		// panicing to help with debugging (fail fast), but we shouldn't panic if we run this in production
		panic(blsAggServiceResp.Err)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []cstaskmanager.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}

	o.tasksMu.RLock()
	task := o.tasks[blsAggServiceResp.TaskIndex]
	o.tasksMu.RUnlock()
	o.taskResponsesMu.RLock()
	taskResponse := o.taskResponses[blsAggServiceResp.TaskIndex][blsAggServiceResp.TaskResponseDigest]
	o.taskResponsesMu.RUnlock()

	o.logger.Info("Quorum reached on source. Checking is quorum achieved over minimum amount of sources",
		"taskIndex", blsAggServiceResp.TaskIndex,
		"source", taskResponse.Source,
	)

	o.tasksMu.Lock()
	// Cache threshold is achieved for this taskDigest (ie. source)
	o.taskDigestQuorum[blsAggServiceResp.TaskIndex] = TaskDigestQuorum{
		finalTaskResponse:           taskResponse,
		nonSignerStakesAndSignature: nonSignerStakesAndSignature,
	}

	// Iterate through all task digest quorums on this task and ensure enough price feed sources acheive quorom
	taskSubmissions := []TaskDigestQuorum{}
	sourcesAtQuorum := []string{}

	for _, submission := range o.taskDigestQuorum {
		doesContain := contains(sourcesAtQuorum, submission.finalTaskResponse.Source)
		// Each source can only have 1 quorum
		if !doesContain {
			sourcesAtQuorum = append(sourcesAtQuorum, submission.finalTaskResponse.Source)
			taskSubmissions = append(taskSubmissions, submission)
		}
	}

	// If quorum reached submit on chain
	if len(taskSubmissions) >= int(task.MinNumOfOracleNetworks) {

		var finalResponses []cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse
		var finalStakesAndSignatures []cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature

		for _, submission := range taskSubmissions {
			finalResponses = append(finalResponses, submission.finalTaskResponse)
			finalStakesAndSignatures = append(finalStakesAndSignatures, submission.nonSignerStakesAndSignature)
		}

		o.logger.Info("Quorum reached on enough sources. Submitting aggregate responses from each price feed source",
			"taskIndex", blsAggServiceResp.TaskIndex,
			"minNumOfOracleNetworks", task.MinNumOfOracleNetworks,
			"numOfSourcesWithQuorum", len(taskSubmissions),
		)

		_, err := o.avsWriter.SendAggregatedResponse(context.Background(), task, finalResponses, finalStakesAndSignatures)

		o.logger.Info("Send aggregate response", "finalRepsonses", finalResponses, "signatures", finalStakesAndSignatures)
		if err != nil {
			o.logger.Error("Aggregator failed to respond to task", "err", err)
		}

		// Elect new operator as leader
		o.priceFSM.TriggerElection()
	}
	o.tasksMu.Unlock()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
