package types

import (
	"errors"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/CredibleSquaringTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ICredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.ICredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
