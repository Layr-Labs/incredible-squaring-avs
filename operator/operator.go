package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/metrics"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSARegistryCoordinator"
	stakereg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAStakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const AVS_NAME = "incredible-squaring"
const SEM_VER = "0.0.1"

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient eth.EthClient
	// TODO(samlaf): remove both avsWriter and eigenlayerWrite from operator
	// they are only used for registration, so we should make a special registration package
	// this way, auditing this operator code makes it obvious that operators don't need to
	// write to the chain during the course of their normal operations
	// writing to the chain should be done via the cli only
	metricsReg         *prometheus.Registry
	metrics            metrics.Metrics
	nodeApi            *nodeapi.NodeApi
	avsWriter          *chainio.AvsWriter
	avsReader          chainio.AvsReaderer
	avsSubscriber      chainio.AvsSubscriberer
	eigenlayerReader   sdkelcontracts.ELReader
	eigenlayerWriter   sdkelcontracts.ELWriter
	avsEcdsaPrivateKey *ecdsa.PrivateKey
	operatorId         sdktypes.EcdsaOperatorId
	operatorAddr       common.Address
	// receive new tasks in this chan (typically from listening to onchain event)
	newTaskCreatedChan chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	// ip address of aggregator
	aggregatorServerIpPortAddr string
	// rpc client to send signed task responses to aggregator
	aggregatorRpcClient AggregatorRpcClienter
	// needed when opting in to avs (allow this service manager contract to slash operator)
	credibleSquaringServiceManagerAddr common.Address
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

	var ethRpcClient, ethWsClient eth.EthClient
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

	avsEcdsaKeyPassword, ok := os.LookupEnv("AVS_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("AVS_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}
	avsEcdsaPrivateKey, err := ecdsa.ReadKey(c.AvsEcdsaPrivateKeyStorePath, avsEcdsaKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse avs ecdsa private key", "err", err)
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
		KeystorePath: c.OperatorEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil {
		panic(err)
	}

	operatorAddress, err := sdkecdsa.GetAddressFromKeyStoreFile(c.OperatorEcdsaPrivateKeyStorePath)
	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, operatorAddress)

	registryCoordinator, err := regcoord.NewContractECDSARegistryCoordinator(common.HexToAddress(c.AVSRegistryCoordinatorAddress), ethRpcClient)
	if err != nil {
		logger.Fatal("Failed to create registry coordinator", "err", err)
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch stake registry address", "err", err)
	}
	stakeRegistry, err := stakereg.NewContractECDSAStakeRegistry(stakeRegistryAddr, ethRpcClient)
	if err != nil {
		logger.Fatal("Failed to create stake registry", "err", err)
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch delegation manager address", "err", err)
	}
	elChainReader, err := elcontracts.BuildELChainReader(delegationManagerAddr, ethRpcClient, logger)
	if err != nil {
		logger.Fatal("Failed to create ELChainReader", "err", err)
	}
	noopMetrics := metrics.NewNoopMetrics()
	elChainWriter, err := elcontracts.BuildELChainWriter(delegationManagerAddr, ethRpcClient, logger, noopMetrics, txMgr)
	if err != nil {
		logger.Fatal("Failed to create ELChainWriter", "err", err)
	}

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

	// TODO: economic collector doesn't work with ecdsa registry atm
	// We must register the economic metrics separately because they are exported metrics (from jsonrpc or subgraph calls)
	// and not instrumented metrics: see https://prometheus.io/docs/instrumenting/writing_clientlibs/#overall-structure
	// quorumNames := map[sdktypes.QuorumNum]string{
	// 	0: "quorum0",
	// }
	// economicMetricsCollector := economic.NewCollector(
	// 	elChainReader, avsReader.AvsEcdsaRegistryReader,
	// 	AVS_NAME, logger, operatorAddress, quorumNames)
	// reg.MustRegister(economicMetricsCollector)

	aggregatorRpcClient, err := NewAggregatorRpcClient(c.AggregatorServerIpPortAddress, logger, avsAndEigenMetrics)
	if err != nil {
		logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		return nil, err
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
		eigenlayerReader:                   elChainReader,
		eigenlayerWriter:                   elChainWriter,
		avsEcdsaPrivateKey:                 avsEcdsaPrivateKey,
		operatorAddr:                       operatorAddress,
		aggregatorServerIpPortAddr:         c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:                aggregatorRpcClient,
		newTaskCreatedChan:                 make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
		credibleSquaringServiceManagerAddr: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
	}

	if c.RegisterOperatorOnStartup {
		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			c.OperatorEcdsaPrivateKeyStorePath,
			ecdsaKeyPassword,
		)
		if err != nil {
			return nil, err
		}
		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey, common.HexToAddress(c.TokenStrategyAddr))
	}

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
		return fmt.Errorf("operator is not registered. Register operator using the operator-cli before starting operator")
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

	// TODO(samlaf): wrap this call with increase in avs-node-spec metric
	sub := o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
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
			sub = o.avsSubscriber.SubscribeToNewTasks(o.newTaskCreatedChan)
		case newTaskCreatedLog := <-o.newTaskCreatedChan:
			o.metrics.IncNumTasksReceived()
			taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
			signedTaskResponse, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				continue
			}
			go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
		}
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated) *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse {
	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"numberToBeSquared", newTaskCreatedLog.Task.NumberToBeSquared,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)
	numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
		NumberSquared:      numberSquared,
	}
	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse) (*aggregator.SignedTaskResponse, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	ecdsaSignature, err := sdkecdsa.SignMsg(taskResponseHash[:], o.avsEcdsaPrivateKey)
	if err != nil {
		o.logger.Error("Error signing task response", "err", err)
		return nil, err
	}
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse:   *taskResponse,
		EcdsaSignature: ecdsaSignature,
		OperatorId:     o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
