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

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	sdkchallengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	"github.com/Layr-Labs/eigensdk-go/utils"
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

	cfg := sdkchallenger.Config{
		EthWsUrl:       config.EthWsRpcUrl,
		Logger:         config.Logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      &config.EthHttpClient,
	}

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		config.IncredibleSquaringServiceManager,
		&config.EthHttpClient,
	)

	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})

	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[*big.Int, *big.Int](
		taskManagerAddr,
		taskManagerAbi,
		config.TxMgr,
		cfg.EthClient,
	)

	indexingChallengerProcessor, err := sdkchallengerprocessor.NewIndexingChallengerProcessor(
		config.Logger,
		squareValidation,
		challengerRaiser,
	)

	challenger, err := sdkchallenger.NewChallenger(
		cfg,
		indexingChallengerProcessor,
	)
	if err != nil {
		config.Logger.Fatalf("Failed to create challenger from config: %v", err)
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
