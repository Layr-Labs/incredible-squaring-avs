package operator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"

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
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkmetrics "github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/metrics/collectors/economic"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/nodeapi"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const (
	AVS_NAME = "incredible-squaring"
	SEM_VER  = "0.0.1"
)

type Operator struct {
	config    types.NodeConfig
	logger    logging.Logger
	ethClient sdkcommon.EthClientInterface
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
	eigenlayerReader sdkelcontracts.ChainReader
	eigenlayerWriter sdkelcontracts.ChainWriter
	blsKeypair       *bls.KeyPair
	operatorId       sdktypes.OperatorId
	operatorAddr     common.Address
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
		eigenlayerReader:                   *sdkClients.ElChainReader,
		eigenlayerWriter:                   *sdkClients.ElChainWriter,
		blsKeypair:                         blsKeyPair,
		operatorAddr:                       common.HexToAddress(c.OperatorAddress),
		aggregatorServerIpPortAddr:         c.AggregatorServerIpPortAddress,
		aggregatorRpcClient:                aggregatorRpcClient,
		newTaskCreatedChan:                 make(chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated),
		credibleSquaringServiceManagerAddr: common.HexToAddress(c.AVSRegistryCoordinatorAddress),
		operatorId:                         [32]byte{0}, // this is set below

	}

	if c.RegisterOperatorOnStartup {
		operator.registerOperatorOnStartup(operatorEcdsaPrivateKey, common.HexToAddress(c.TokenStrategyAddr))
	}

	// OperatorId is set in contract during registration so we get it after registering operator.
	operatorId, err := sdkClients.AvsRegistryChainReader.GetOperatorId(&bind.CallOpts{}, operator.operatorAddr)
	if err != nil {
		logger.Error("Cannot get operator id", "err", err)
		return nil, err
	}
	operator.operatorId = operatorId
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

// QueryLocationsResponse represents the JSON response from the Rust API.
type QueryLocationsResponse struct {
	Local   bool             `json:"local"`
	Nodes   [][2]interface{} `json:"nodes"` // each element is a tuple: [nodeId (hex string), score (float64)]
	Message string           `json:"message"`
}

// fetchAndTransformScores calls the /api/v0/query/<fileHash> endpoint,
// decodes the response, and transforms it into the desired task response type.
// In case of any network error or processing failure, it returns an empty taskResponse.
func fetchAndTransformScores(fileHash [32]byte, taskIndex uint32) *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse {
	// Prepare an empty task response.
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: taskIndex,
		Providers:          make([]common.Address, 0),
		Scores:             make([]*big.Int, 0),
	}

	// Convert fileHash ([32]byte) to its hex string representation.
	hexFileHash := hex.EncodeToString(fileHash[:])

	log.Printf("hexFileHash: %s", hexFileHash)

	// Construct the API URL (adjust the base URL as needed)
	url := fmt.Sprintf("http://localhost:8080/api/v0/query/%s", hexFileHash)

	log.Printf("url: %s", url)

	// Make the HTTP GET request.
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to fetch API: %v", err)
		return taskResponse
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %s", resp.Status)
		return taskResponse
	}

	// Read and decode the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response: %v", err)
		return taskResponse
	}

	log.Printf("Response body: %s", string(body))

	var queryResp QueryLocationsResponse
	if err := json.Unmarshal(body, &queryResp); err != nil {
		log.Printf("failed to decode JSON: %v", err)
		return taskResponse
	}

	// Iterate over each node tuple and convert the data.
	for _, node := range queryResp.Nodes {
		// Extract the nodeId, expected to be a 32-byte hex string.
		nodeId, ok := node[0].(string)
		if !ok {
			log.Printf("unexpected nodeId type: %v", node[0])
			continue
		}

		// Extract the score, expected to be a float64.
		scoreFloat, ok := node[1].(float64)
		if !ok {
			log.Printf("unexpected score type: %v", node[1])
			continue
		}

		// Convert the nodeId to a common.Address.
		provider := common.HexToAddress(nodeId)

		// Convert the float64 score to a *big.Int using basis points.
		scoreInt := big.NewInt(int64(scoreFloat * 10000))

		// Append the converted values.
		taskResponse.Providers = append(taskResponse.Providers, provider)
		taskResponse.Scores = append(taskResponse.Scores, scoreInt)
	}

	return taskResponse
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (o *Operator) ProcessNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated) *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse {
	o.logger.Debug("Received new task", "task", newTaskCreatedLog)
	o.logger.Info("Received new task",
		"fileHash", newTaskCreatedLog.Task.FileHash,
		"taskIndex", newTaskCreatedLog.TaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	taskResponse := fetchAndTransformScores(newTaskCreatedLog.Task.FileHash, newTaskCreatedLog.TaskIndex)

	// Use taskResp as needed.
	fmt.Printf("Transformed Task Response: %+v\n", taskResponse)

	return taskResponse
}

func (o *Operator) SignTaskResponse(taskResponse *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse) (*aggregator.SignedTaskResponse, error) {
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		o.logger.Error("Error getting task response header hash. skipping task (this is not expected and should be investigated)", "err", err)
		return nil, err
	}
	blsSignature := o.blsKeypair.SignMessage(taskResponseHash)
	signedTaskResponse := &aggregator.SignedTaskResponse{
		TaskResponse: *taskResponse,
		BlsSignature: *blsSignature,
		OperatorId:   o.operatorId,
	}
	o.logger.Debug("Signed task response", "signedTaskResponse", signedTaskResponse)
	return signedTaskResponse, nil
}
