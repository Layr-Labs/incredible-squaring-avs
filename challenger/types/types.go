package types

import (
	"errors"

	cstaskmanager "github.com/ehsueh/trade-algo-avs-avs/contracts/bindings/TradeAlgoTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.ITradeAlgoTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.ITradeAlgoTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
