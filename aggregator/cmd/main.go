package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/urfave/cli"
	"golang.org/x/crypto/sha3"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

var (
	// Version is the version of the binary.
	Version   string
	GitCommit string
	GitDate   string
)

func main() {

	app := cli.NewApp()
	app.Flags = config.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "credible-squaring-aggregator"
	app.Usage = "Credible Squaring Aggregator"
	app.Description = "Service that sends number to be credibly squared by operator nodes."

	app.Action = aggregatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func aggregatorMain(ctx *cli.Context) error {

	log.Println("Initializing Aggregator")
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	cfg := sdkaggregator.AggregatorConfig{
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

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

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

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			return sdktypes.TaskResponseDigest{}, utils.WrapError("Error Packing taskResponse", err)
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}
	cfg.TaskResponseHashFn = hashFunction

	newTaskCreatedEventHash := taskManagerAbi.Events["NewTaskCreated"].ID

	agg, err := sdkaggregator.NewAggregator[aggregator.IncredibleSquaringTaskResponse, *big.Int, *big.Int](
		cfg,
		taskProcessor,
		newTaskCreatedEventHash,
		taskManagerAbi,
	)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background())
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	return nil
}
