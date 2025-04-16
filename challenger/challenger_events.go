package challenger

import (
	"math/big"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"

	"github.com/ethereum/go-ethereum/core/types"
)

type NewTaskCreatedEvent struct {
	TaskIndex uint32
	Task      sdkchallenger.GenericInputTask[*big.Int]
	Raw       types.Log
}

func (newTaskEvent NewTaskCreatedEvent) InnerTask() sdkchallenger.GenericInputTask[*big.Int] {
	return newTaskEvent.Task
}

type TaskRespondedEvent struct {
	TaskResponse              sdkchallenger.GenericInputTaskResponse[*big.Int]
	TaskResponseMetadata      sdkchallenger.GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []sdkchallenger.BN254G1Point
}

func (taskRespEvent TaskRespondedEvent) TaskIndex() uint32 {
	return taskRespEvent.TaskResponse.ReferenceTaskIndex
}

func (taskRespEvent TaskRespondedEvent) GetTaskResponse() sdkchallenger.GenericInputTaskResponse[*big.Int] {
	return taskRespEvent.TaskResponse
}

func (taskRespEvent TaskRespondedEvent) GetTaskResponseMetadata() sdkchallenger.GenericTaskResponseMetadata {
	return taskRespEvent.TaskResponseMetadata
}
