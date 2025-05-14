package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/urfave/cli"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	csservicemanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringServiceManager"
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

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

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
		TaskManagerAbi:                taskManagerAbi,
	}

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		config.IncredibleSquaringServiceManager,
		&config.EthHttpClient,
	)

	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})

	taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		taskManagerAddr,
		taskManagerAbi,
		config.TxMgr,
		&config.EthHttpClient,
	)

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(config.Logger, taskResponder)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	agg, err := sdkaggregator.NewAggregator(
		cfg,
		taskProcessor,
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
