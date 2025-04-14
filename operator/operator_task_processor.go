package operator

import (
	"fmt"
	"math/big"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/ethereum/go-ethereum/core/types"
)

type OperatorTaskProcessor struct {
	logger logging.Logger
}

func NewOperatorTaskProcessor(c sdkoperator.OperatorConfig, logger logging.Logger) OperatorTaskProcessor {
	return OperatorTaskProcessor{
		logger: logger,
	}
}

// Takes a NewTaskCreatedLog struct as input and returns a TaskResponseHeader struct.
// The TaskResponseHeader struct is the struct that is signed and sent to the contract as a task response.
func (otp OperatorTaskProcessor) ProcessNewTaskCreatedLog(
	log types.Log,
) (sdkaggregator.TaskResponse, error) {
	var newTaskCreatedLog cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		otp.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return nil, fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	otp.logger.Debug("Received new task", "task", newTaskCreatedLog)
	otp.logger.Info("Received new task",
		"numberToBeSquared", newTaskCreatedLog.Task.NumberToBeSquared,
		"taskIndex", newTaskIndex,
		"taskCreatedBlock", newTaskCreatedLog.Task.TaskCreatedBlock,
		"quorumNumbers", newTaskCreatedLog.Task.QuorumNumbers,
		"QuorumThresholdPercentage", newTaskCreatedLog.Task.QuorumThresholdPercentage,
	)

	numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)

	// if otp.timesFailing > 0 {
	// 	rand.Seed(uint64((time.Now().UnixNano())))
	// 	num := rand.Intn(100)
	// 	if num < otp.timesFailing {
	// 		numberSquared = big.NewInt(908243203843)
	// 		otp.logger.Info("Operator computed wrong task result")
	// 	}
	// }

	taskResponse := aggregator.IncredibleSquaringTaskResponse{
		ReferenceTaskIndex: newTaskIndex,
		NumberSquared:      numberSquared,
	}
	return taskResponse, nil
}
