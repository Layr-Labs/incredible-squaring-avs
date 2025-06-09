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

	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/testutils"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	taskspammer "github.com/Layr-Labs/incredible-squaring-avs/task-spammer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	sdkintegration "github.com/Layr-Labs/eigensdk-go/integration-tests"
)

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

	tsConfig := &taskspammer.Config{}
	err = commonincredible.ReadTomlConfig("../../config-files/config.toml", tsConfig)
	if err != nil {
		t.Fatalf("Failed to read operator config: %s", err.Error())
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	responseCalculationBuilder := func() sdkoperator.ResponseCalculator[*big.Int, *big.Int] {
		return sdkoperator.NewFunctionResponseCalculator(square)
	}

	equalFn := func(a, b *big.Int) bool {
		return a.Cmp(b) == 0
	}

	avsConfig := sdkintegration.AvsConfig[*big.Int, *big.Int]{
		TaskManagerAddr: common.HexToAddress(aggCfg.TaskManagerAddress),
		TaskManagerAbi:  taskManagerAbi,

		EthHttpUrl: aggCfg.EthHttpUrl,
		EthWsUrl:   aggCfg.EthWsUrl,

		ResponseCalculatorBuilder: responseCalculationBuilder,
		EqualFn:                   equalFn,
		InputSequence:             NewNumberToSquareSequence(),

		AggregatorServerIpPortAddr: aggCfg.AggregatorServerIpPortAddr,

		RegistryCoordinatorAddress:    aggCfg.RegistryCoordinatorAddress,
		OperatorStateRetrieverAddress: aggCfg.OperatorStateRetrieverAddress,
		AvsAddress:                    opCfg.Registration.AvsAddress,

		OperatorPrivateKey:    "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		OperatorBlsPrivateKey: "0x2518600ef40ef39cb4ab8b828ce303b3e02ac01ec6ba6bd0d0cf0663e1252ff0",

		TaskSpammerPrivateKey: "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
		AggregatorPrivateKey:  "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6",
		ChallengerPrivateKey:  testutils.ANVIL_FIRST_PRIVATE_KEY,
		AmountToMint:          "1000000000000000000000",

		AllocationManagerAddr:       opCfg.Registration.AllocationManagerAddr,
		DelegationManagerAddress:    opCfg.Registration.DelegationManagerAddress,
		StrategyAddr:                opCfg.Registration.StrategyAddrs[0],
		RewardsCoordinatorAddress:   opCfg.Registration.RewardsCoordinatorAddress,
		PermissionControllerAddress: opCfg.Registration.PermissionControllerAddress,

		AllocatableMagnitude: opCfg.Registration.AllocatableMagnitudes[0],
		OperatorSetId:        opCfg.Registration.OperatorSetIds[0],
		OperatorAddr:         opCfg.OperatorAddress,

		TimeBetweenTasks:          tsConfig.TimeBetweenTasks,
		QuorumThresholdPercentage: tsConfig.QuorumThresholdPercentage,
		QuorumNumbers:             tsConfig.QuorumNumbers,
	}

	ctx, cancel := context.WithCancel(context.Background())

	avs := sdkintegration.StartAvs(t, ctx, avsConfig)

	timer := time.NewTimer(32 * time.Second)

	select {
	case err := <-avs.Aggregator:
		t.Fatal("Aggregator error:", err)
	case err := <-avs.Challenger:
		t.Fatal("Challenger error:", err)
	case err := <-avs.Operator:
		t.Fatal("Operator error:", err)
	case err := <-avs.TaskSpammer:
		// Here we handle the input generation termination (in that case returns nil but does not imply an error)
		if err == nil {
			cancel()
		} else {
			t.Fatal("Task Spammer error:", err)
		}
	case <-timer.C:
		// If reached this point, there must be a problem with the tasks generation.
		t.Fatal("Timer expired before the task spammer finished sending tasks")
	}

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

// Returns an iterator for the sequence 1, 2, 3, ...
func NewNumberToSquareSequence() iter.Seq[*big.Int] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	count := 0
	return func(yield func(*big.Int) bool) {
		for count < 3 {
			if !yield(acc) {
				break
			}
			acc.Add(acc, delta)
			count++
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
