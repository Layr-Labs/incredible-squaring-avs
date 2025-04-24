package operator

import (
	"math/big"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
)

type OperatorTaskProcessor struct {
	logger logging.Logger
}

var _ sdkoperator.OperatorTaskProcessor[*big.Int] = (*OperatorTaskProcessor)(nil)

func NewOperatorTaskProcessor(c sdkoperator.OperatorConfig, logger logging.Logger) OperatorTaskProcessor {
	return OperatorTaskProcessor{
		logger: logger,
	}
}

func (otp OperatorTaskProcessor) DigestResponse(response *sdkchallenger.GenericInputTaskResponse[*big.Int]) [32]byte {
	incredibleSquaringTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: response.ReferenceTaskIndex,
		NumberSquared:      response.InputValue,
	}
	taskResponseHash, err := core.GetTaskResponseDigest(&incredibleSquaringTaskResponse)
	if err != nil {
		return [32]byte{}
	}
	return taskResponseHash
}
