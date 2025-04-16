package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/incredible-squaring-avs/challenger"
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

	newTaskEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	taskRespondedEventHash := taskManagerAbi.Events["TaskResponded"].ID

	cfg := sdkchallenger.ChallengerConfig{
		EthWsUrl: config.EthWsRpcUrl,
		Logger:   config.Logger,
	}

	challengerLogicImpl, err := challenger.NewChallengerLogicImpl(config)
	if err != nil {
		config.Logger.Errorf("Failed to create challenger logic from config: %v", err)
		return err
	}

	challenger, err := sdkchallenger.NewChallenger[
		NewTaskCreatedEvent, 
		TaskRespondedEvent,
		cstaskmanager.IIncredibleSquaringTaskManagerTask,
		](cfg, challengerLogicImpl, newTaskEventHash, taskRespondedEventHash, taskManagerAbi)
	if err != nil {
		config.Logger.Errorf("Failed to create challenger from config: %v", err)
		return err
	}

	err = challenger.Start(context.Background())
	if err != nil {
		return err
	}

	return nil

}

type NewTaskCreatedEvent struct {
	TaskIndex uint32
	Task      cstaskmanager.IIncredibleSquaringTaskManagerTask
	Raw       types.Log
}

func (newTaskEvent NewTaskCreatedEvent) InnerTask() (sdkchallenger.GenericTask){
	return newTaskEvent.Task
}

type TaskRespondedEvent struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []sdkchallenger.BN254G1Point
}


func (taskRespEvent TaskRespondedEvent)TaskIndex()(uint32){
	return taskRespEvent.TaskResponse.ReferenceTaskIndex
}

func (taskRespEvent TaskRespondedEvent)GetTaskResponse()(sdkchallenger.GenericTaskResponse){
	return taskRespEvent.TaskResponse
}

func (taskRespEvent TaskRespondedEvent)GetTaskResponseMetadata()(sdkchallenger.GenericTaskResponseMetadata){
	return taskRespEvent.TaskResponseMetadata
}
