package integration_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	sdktaskgenerator "github.com/Layr-Labs/eigensdk-go/task-generator"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	"github.com/Layr-Labs/incredible-squaring-avs/challenger"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	taskgenerator "github.com/Layr-Labs/incredible-squaring-avs/task-generator"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/crypto/sha3"
)

type IntegrationClients struct {
	Sdkclients clients.Clients
}

func TestIntegration(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	/* Start the anvil chain */
	anvilC := startAnvilTestContainer()
	// Not sure why but deferring anvilC.Terminate() causes a panic when the test finishes...
	// so letting it terminate silently for now
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	/* Prepare the config file for aggregator */
	var aggConfigRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	commonincredible.ReadYamlConfig(aggConfigFilePath, &aggConfigRaw)
	aggConfigRaw.EthRpcUrl = "http://" + anvilEndpoint
	aggConfigRaw.EthWsUrl = "ws://" + anvilEndpoint

	var credibleSquaringDeploymentRaw config.IncredibleSquaringDeploymentRaw
	credibleSquaringDeploymentFilePath := "../../contracts/script/deployments/incredible-squaring//31337.json"
	commonincredible.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(aggConfigRaw.Environment)

	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}
	ethRpcClient, err := ethclient.Dial(aggConfigRaw.EthRpcUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}
	ethWsClient, err := ethclient.Dial(aggConfigRaw.EthWsUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}

	aggregatorEcdsaPrivateKeyString := "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	if aggregatorEcdsaPrivateKeyString[:2] == "0x" {
		aggregatorEcdsaPrivateKeyString = aggregatorEcdsaPrivateKeyString[2:]
	}
	aggregatorEcdsaPrivateKey, err := crypto.HexToECDSA(aggregatorEcdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}
	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(aggregatorEcdsaPrivateKey)
	if err != nil {
		t.Fatalf("Cannot get operator address: %s", err.Error())
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("Cannot get chainId: %s", err.Error())
	}

	privateKeySigner, _, err := signerv2.SignerFromConfig(
		signerv2.Config{PrivateKey: aggregatorEcdsaPrivateKey},
		chainId,
	)
	if err != nil {
		t.Fatalf("Cannot create signer: %s", err.Error())
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, privateKeySigner, aggregatorAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, aggregatorAddr)

	config := &config.Config{
		EcdsaPrivateKey: aggregatorEcdsaPrivateKey,
		Logger:          logger,
		EthHttpRpcUrl:   aggConfigRaw.EthRpcUrl,
		EthHttpClient:   *ethRpcClient,
		EthWsRpcUrl:     aggConfigRaw.EthWsUrl,
		EthWsClient:     *ethWsClient,
		OperatorStateRetrieverAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.OperatorStateRetrieverAddr,
		),
		IncredibleSquaringRegistryCoordinatorAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.RegistryCoordinatorAddr,
		),
		AggregatorServerIpPortAddr: aggConfigRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:  aggConfigRaw.RegisterOperatorOnStartup,
		TxMgr:                      txMgr,
		AggregatorAddress:          aggregatorAddr,
		IncredibleSquaringServiceManager: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.IncredibleSquaringServiceManager,
		),
	}

	/* Prepare the config file for operator */
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err = commonincredible.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	thresholdNumerator := sdktypes.QuorumThresholdPercentage(100)
	quorumNumbers := sdktypes.QuorumNums{0}
	taskGenLogic, err := taskgenerator.NewTaskGenLogic(config, thresholdNumerator, quorumNumbers)

	taskGenerator, err := sdktaskgenerator.BuildTaskGenerator(config.Logger, taskGenLogic, 10)
	if err != nil {
		t.Fatalf("Failed to create task generator: %s", err.Error())
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	newTaskEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	taskRespondedEventHash := taskManagerAbi.Events["TaskResponded"].ID

	challenferCfg := sdkchallenger.ChallengerConfig{
		EthWsUrl: config.EthWsRpcUrl,
		Logger:   config.Logger,
	}

	challengerVerifier, err := challenger.NewChallengerVerifierImpl(config)
	if err != nil {
		config.Logger.Fatalf("Failed to create challenger logic from config: %v", err)
	}

	challenger, err := sdkchallenger.NewChallenger(
		challenferCfg,
		challengerVerifier,
		newTaskEventHash,
		taskRespondedEventHash,
		taskManagerAbi,
		&config.EthHttpClient,
	)
	if err != nil {
		config.Logger.Fatalf("Failed to create challenger from config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
	defer cancel()
	/* start operator */
	// the passwords are set to empty strings
	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")
	nodeConfig.BlsPrivateKeyStorePath = "../keys/test.bls.key.json"
	nodeConfig.EcdsaPrivateKeyStorePath = "../keys/test.ecdsa.key.json"
	nodeConfig.EthRpcUrl = "http://" + anvilEndpoint
	nodeConfig.EthWsUrl = "ws://" + anvilEndpoint

	blockHash := taskManagerAbi.Events["NewTaskCreated"].ID

	err = operator.RegisterOperatorOnStartup(nodeConfig, logger)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operatorConfig := sdkoperator.OperatorConfig{
		OperatorAddress:               nodeConfig.OperatorAddress,
		OperatorStateRetrieverAddress: nodeConfig.OperatorStateRetrieverAddress,
		ServiceManagerAddress:         nodeConfig.IncredibleSquaringServiceManager,
		AVSRegistryCoordinatorAddress: nodeConfig.AVSRegistryCoordinatorAddress,
		EthRpcUrl:                     nodeConfig.EthRpcUrl,
		EthWsUrl:                      nodeConfig.EthWsUrl,
		BlsPrivateKeyStorePath:        nodeConfig.BlsPrivateKeyStorePath,
		AggregatorServerIpPortAddress: nodeConfig.AggregatorServerIpPortAddress,
		RegisterOnStartup:             true,
	}
	operatorTaskProcessor := operator.NewOperatorTaskProcessor(operatorConfig, logger)
	operator, err := sdkoperator.NewOperatorFromConfig(
		operatorConfig,
		blockHash,
		operatorTaskProcessor,
		logger,
		taskManagerAbi,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	go operator.Start(ctx)
	log.Println("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	/* start aggregator */
	log.Println("starting aggregator for integration tests")
	aggConfig := sdkaggregator.AggregatorConfig{
		RegistryCoordinatorAddress:    config.IncredibleSquaringRegistryCoordinatorAddr,
		OperatorStateRetrieverAddress: config.OperatorStateRetrieverAddr,
		ServiceManagerAddress:         config.IncredibleSquaringServiceManager,
		EthHttpClient:                 &config.EthHttpClient,
		Logger:                        config.Logger,
		EthHttpUrl:                    config.EthHttpRpcUrl,
		EthWsUrl:                      config.EthWsRpcUrl,
		EcdsaPrivateKey:               config.EcdsaPrivateKey,
		AggregatorServerIpPortAddr:    config.AggregatorServerIpPortAddr,
	}

	taskProcessor, err := aggregator.NewTaskProcessor(config)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	go challenger.Start(ctx)

	// This is the same hash function used by the operator to hash the task response before signing it.
	hashFunction := func(taskResponse sdktypes.TaskResponse) (sdktypes.TaskResponseDigest, error) {
		// The order here has to match the field ordering of cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
		taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{
				Name: "referenceTaskIndex",
				Type: "uint32",
			},
			{
				Name: "numberSquared",
				Type: "uint256",
			},
		})
		if err != nil {
			return sdktypes.TaskResponseDigest{}, utils.WrapError("Error creating taskResponseType", err)
		}
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		taskResponseAgg, ok := taskResponse.(sdkchallenger.GenericOutputTaskResponse[*big.Int])
		if !ok {
			return sdktypes.TaskResponseDigest{}, errors.New(
				"task Response could not be converted to sdk aggregator's Task Response type",
			)
		}

		incredibleSquaringTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: taskResponseAgg.ReferenceTaskIndex,
			NumberSquared:      taskResponseAgg.OutputValue,
		}
		encodeTaskResponseByte, err := arguments.Pack(incredibleSquaringTaskResponse)
		if err != nil {
			return sdktypes.TaskResponseDigest{}, utils.WrapError("Error Packing taskResponse", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
	aggConfig.TaskResponseHashFn = hashFunction

	agg, err := sdkaggregator.NewAggregator[*big.Int, *big.Int](
		aggConfig,
		taskProcessor,
		blockHash,
		taskManagerAbi,
	)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	go agg.Start(ctx)

	go taskGenerator.Start(ctx)

	log.Println("Started aggregator and task generator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	// get avsRegistry client to interact with the chain
	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	// check if the task is recorded in the contract for task index 1
	taskHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskHashes(&bind.CallOpts{}, 1)
	if err != nil {
		t.Fatalf("Cannot get task hash: %s", err.Error())
	}
	if taskHash == [32]byte{} {
		t.Fatalf("Task hash is empty")
	}

	received := false
	for i := 0; i < 3; i++ {
		// check if the task response is recorded in the contract for task index 1
		taskResponseHash, err := avsReader.AvsServiceBindings.TaskManager.AllTaskResponses(&bind.CallOpts{}, 1)
		log.Printf("taskResponseHash: %v", taskResponseHash)
		if err != nil {
			t.Fatalf("Cannot get task response hash: %s", err.Error())
		}

		if taskResponseHash != [32]byte{} {
			received = true
			break
		}
	}

	if !received {
		t.Fatalf("Task response hash is empty")
	}
}

// TODO(samlaf): have to advance chain to a block where the task is answered
func startAnvilTestContainer() testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: filepath.Join(
						integrationDir,
						"../anvil/avs-and-eigenlayer-deployed-anvil-state/state.json",
					),
				},
				Target: "/state.json",
			},
		},
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--load-state", "/state.json"},
		ExposedPorts: []string{"8545/tcp"},

		WaitingFor: wait.ForLog("Listening on"),
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	// this is needed temporarily because anvil restarts at 0 block when we load a state...
	// see comment in start-anvil-chain-with-el-and-avs-deployed.sh
	advanceChain(anvilC)
	return anvilC
}

func advanceChain(anvilC testcontainers.Container) {
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		panic(err)
	}
	rpcUrl := "http://" + anvilEndpoint
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf(
			`cast rpc anvil_mine 100 --rpc-url %s`,
			rpcUrl),
	)
	cmd.Dir = "../../contracts"
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
