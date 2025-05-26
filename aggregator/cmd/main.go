package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
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

	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	configFilePath := ctx.String("config")
	if configFilePath == "" {
		logger.Fatal("Missing required flag: --config")
	}
	aggConfig := &aggregator.Config{}
	err = commonincredible.ReadTomlConfig(configFilePath, aggConfig)

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	cfg := aggConfig.Config

	ethRpcClient, err := ethclient.Dial(aggConfig.EthHttpUrl)

	privateKeyHex := ctx.String("ecdsa-private-key")
	if privateKeyHex == "" {
		logger.Fatal("Missing required flag: --ecdsa-private-key")
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		logger.Errorf("Cannot parse ECDSA private key", "err", err)
		return err
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethRpcClient, ecdsaPrivateKey)

	taskResponder, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		common.HexToAddress(aggConfig.TaskManagerAddress),
		taskManagerAbi,
		txMgr,
		ethRpcClient,
	)

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	agg, err := sdkaggregator.NewAggregator(
		cfg,
		logger,
		taskProcessor,
		taskManagerAbi,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = agg.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return nil
}
