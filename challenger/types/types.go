package types

import (
	"errors"

	incsqtaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/ethereum/go-ethereum/common"
)

type TaskResponseData struct {
	TaskResponse         incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	SignersIds           []common.Address
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
