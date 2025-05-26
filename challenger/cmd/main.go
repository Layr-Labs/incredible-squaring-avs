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

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	sdkchallengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	"github.com/Layr-Labs/eigensdk-go/utils"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	EthHttpUrl string `toml:"eth_http_url"`
	EthWsUrl   string `toml:"eth_ws_url"`
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
	app.Name = "credible-squaring-challenger"
	app.Usage = "Credible Squaring Challenger"
	app.Description = "Service that challenges wrong response to the task."

	app.Action = challengerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func challengerMain(ctx *cli.Context) error {

	log.Println("Initializing Challenger...")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	challengerConfig := &Config{}
	err := commonincredible.ReadTomlConfig(configPath, challengerConfig)

	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	ethRpcClient, err := ethclient.Dial(challengerConfig.EthHttpUrl)

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

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

	cfg := sdkchallenger.Config{
		EthWsUrl:       challengerConfig.EthWsUrl,
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethRpcClient,
	}

	taskManagerAddr := challengerConfig.TaskManagerAddress
	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		common.HexToAddress(taskManagerAddr),
		taskManagerAbi,
		txMgr,
		cfg.EthClient,
	)

	indexingChallengerProcessor, err := sdkchallengerprocessor.NewIndexingChallengerProcessor(
		logger,
		squareValidation,
		challengerRaiser,
	)

	challenger, err := sdkchallenger.NewChallenger(
		cfg,
		indexingChallengerProcessor,
	)
	if err != nil {
		logger.Fatalf("Failed to create challenger from config: %v", err)
	}

	err = challenger.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}

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
