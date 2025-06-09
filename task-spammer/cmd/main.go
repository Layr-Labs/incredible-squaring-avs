package main

import (
	"context"
	"fmt"
	"iter"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktaskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	taskspammer "github.com/Layr-Labs/incredible-squaring-avs/task-spammer"

	"github.com/Layr-Labs/incredible-squaring-avs/core/config"

	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
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
	app.Name = "credible-squaring-task-spammer"
	app.Usage = "Credible Squaring Task Spammer"
	app.Description = "Service that generates tasks and sends them to Task Manager."

	app.Action = taskSpammerMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func taskSpammerMain(ctx *cli.Context) error {

	log.Println("Initializing Task Spammer...")

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	tsConfig := &taskspammer.Config{}
	err := commonincredible.ReadTomlConfig(configPath, tsConfig)

	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	ethRpcClient, err := ethclient.Dial(tsConfig.EthHttpUrl)

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

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskManagerAddr := common.HexToAddress(tsConfig.TaskManagerAddress)
	taskCreator, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](taskManagerAddr, taskManagerAbi, txMgr, ethRpcClient)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	taskSpammerCfg := sdktaskspammer.Config{
		TimeBetweenTasks:          10 * time.Second,
		QuorumThresholdPercentage: uint32(100),
		QuorumNumbers:             []uint8{0},
	}

	seq := NewNumberToSquareSequence()

	taskSpammer, err := sdktaskspammer.NewTaskSpammer(logger, taskSpammerCfg, taskCreator, seq)
	if err != nil {
		logger.Fatalf("Failed to create task spammer: %s", err.Error())
	}

	err = <-taskSpammer.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

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
