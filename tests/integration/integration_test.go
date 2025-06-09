package integration_test

import (
	"context"
	"fmt"
	"iter"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktaskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	"github.com/Layr-Labs/incredible-squaring-avs/challenger"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	taskspammer "github.com/Layr-Labs/incredible-squaring-avs/task-spammer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationClients struct {
	Sdkclients clients.Clients
}

func TestIntegration(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	// Start the anvil chain and get the anvil endpoint
	anvilC := startAnvilTestContainer()
	// Not sure why but deferring anvilC.Terminate() causes a panic when the test finishes...
	// so letting it terminate silently for now
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	// Read the configs from the toml config file
	logger, err := sdklogging.NewZapLogger(sdklogging.Production)
	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}

	aggCfg := &aggregator.Config{}
	err = commonincredible.ReadTomlConfig("../../config-files/config.toml", aggCfg)
	if err != nil {
		t.Fatalf("Failed to read aggregator config: %s", err.Error())
	}

	aggCfg.EthHttpUrl = "http://" + anvilEndpoint
	aggCfg.EthWsUrl = "ws://" + anvilEndpoint

	opCfg := &operator.Config{}
	err = commonincredible.ReadTomlConfig("../../config-files/config.toml", opCfg)
	if err != nil {
		t.Fatalf("Failed to read operator config: %s", err.Error())
	}

	opCfg.EthRpcUrl = "http://" + anvilEndpoint
	opCfg.EthWsUrl = "ws://" + anvilEndpoint

	opCfg.BlsSignerCfg.KeystorePath = "../keys/test.bls.key.json"
	opCfg.Registration.EcdsaSignerCfg.KeystorePath = "../keys/test.ecdsa.key.json"

	challengerCfg := &challenger.Config{}
	err = commonincredible.ReadTomlConfig("../../config-files/config.toml", challengerCfg)
	if err != nil {
		t.Fatalf("Failed to read operator config: %s", err.Error())
	}

	challengerCfg.EthHttpUrl = "http://" + anvilEndpoint
	challengerCfg.EthWsUrl = "ws://" + anvilEndpoint

	tsConfig := &taskspammer.Config{}
	err = commonincredible.ReadTomlConfig("../../config-files/config.toml", tsConfig)
	if err != nil {
		t.Fatalf("Failed to read operator config: %s", err.Error())
	}

	ethRpcClient, err := ethclient.Dial(aggCfg.EthHttpUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}

	aggregatorEcdsaPrivateKeyString := "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	aggregatorEcdsaPrivateKey, err := crypto.HexToECDSA(aggregatorEcdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethRpcClient, aggregatorEcdsaPrivateKey)

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// We use the same contract for the aggregator, challenger and task spammer because the task manager methods are
	// permissionless excepting RespondToTask() (that is why we use the aggregator private key)
	taskManagerContract, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		common.HexToAddress(aggCfg.TaskManagerAddress),
		taskManagerAbi,
		txMgr,
		ethRpcClient,
	)
	if err != nil {
		t.Fatalf("Failed to create task manager contract: %s", err.Error())
	}

	taskSpammer, err := sdktaskspammer.NewTaskSpammer(logger, tsConfig.Config, taskManagerContract, NewNumberToSquareSequence())
	if err != nil {
		t.Fatalf("Failed to create task spammer: %s", err.Error())
	}

	indexingChallengerProcessor, err := sdkchallenger.NewIndexingProcessor(
		logger,
		squareValidation,
		taskManagerContract,
	)

	challenger, err := sdkchallenger.NewChallenger(
		logger,
		challengerCfg.Config,
		taskManagerAbi,
		indexingChallengerProcessor,
	)
	if err != nil {
		logger.Fatalf("Failed to create challenger from config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
	defer cancel()
	/* start operator */
	// the passwords are set to empty strings
	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")

	calculator := sdkoperator.NewFunctionResponseCalculator(square)

	failingFunction, err := sdkoperator.NewFailingResponseCalculator(calculator, 10, big.NewInt(0))
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operator, err := sdkoperator.NewOperator(
		logger,
		opCfg.Config,
		taskManagerAbi,
		failingFunction,
		nil,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	go operator.Start(ctx)
	log.Println("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	/* start aggregator */
	log.Println("starting aggregator for integration tests")
	aggConfig := aggCfg.Config

	taskProcessor, err := sdkaggregator.NewIndexingProcessor(logger, taskManagerContract)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	go challenger.Start(ctx)
	agg, err := sdkaggregator.NewAggregator(
		logger,
		aggConfig,
		taskManagerAbi,
		taskProcessor,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	go agg.Start(ctx)

	go taskSpammer.Start(ctx)

	log.Println("Started aggregator and task spammer. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	// get avsRegistry client to interact with the chain
	avsReader, err := chainio.BuildAvsReaderFromConfig(opCfg, aggCfg.OperatorStateRetrieverAddress, logger)
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

// This function computes the square of a number
func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
}

func squareValidation(taskIndex uint32, numberToSquare *big.Int, numberSquared *big.Int) (bool, error) {
	result, err := square(taskIndex, numberToSquare)
	if err != nil {
		return false, utils.WrapError("failed to calculate square", err)
	}

	return result.Cmp(numberSquared) == 0, nil
}

// Returns an iterator for the sequence 1, 2, 3, ...
func NewNumberToSquareSequence() iter.Seq[*big.Int] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	return func(yield func(*big.Int) bool) {
		for {
			if !yield(acc) {
				break
			}
			acc.Add(acc, delta)
		}
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
