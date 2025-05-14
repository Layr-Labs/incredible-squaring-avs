package main

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/urfave/cli"

	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktaskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
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
	config, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}
	fmt.Println("Config:", string(configJson))

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(config.IncredibleSquaringServiceManager, &config.EthHttpClient)

	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	taskCreator, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](taskManagerAddr, taskManagerAbi, config.TxMgr, &config.EthHttpClient)
	if err != nil {
		config.Logger.Fatalf(err.Error())
	}

	taskSpammerCfg := sdktaskspammer.Config{
		Logger:                    config.Logger,
		TimeBetweenTasks:          10 * time.Second,
		QuorumThresholdPercentage: uint32(100),
		QuorumNumbers:             []uint8{0},
	}

	taskSpammer, err := sdktaskspammer.NewTaskSpammer(taskCreator, taskSpammerCfg)
	if err != nil {
		config.Logger.Fatalf("Failed to create task spammer: %s", err.Error())
	}

	seq := NewNumberToSquareSequence()

	err = taskSpammer.Start(context.Background(), seq)
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
