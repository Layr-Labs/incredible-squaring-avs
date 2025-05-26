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

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/pelletier/go-toml/v2"
)

// This config has the same attributes as the aggregator config and also includes the
// deployed TaskManager contract address
type Config struct {
	aggregator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}

// This function reads the config from the .toml file at the path received as a parameter
// and returns a config with those values
func GetConfigFromPath(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(data, config)
	return config, err
}

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
	aggConfig, err := GetConfigFromPath(configFilePath)

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
