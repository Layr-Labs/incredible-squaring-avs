package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
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
		TxMgr:                         config.TxMgr,
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

	blockHash := taskManagerAbi.Events["NewTaskCreated"].ID
	agg, err := sdkaggregator.NewAggregator(cfg, taskProcessor, blockHash)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background(), &aggregator.IncredibleSquaringTaskResponse{})
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	return nil

}
