package aggregator

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
)

func TestProcessSignedTaskResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var NUMBER_TO_SQUARE = uint32(3)

	MOCK_OPERATOR_ECDSA_PRIVATE_KEY, err := ecdsa.NewPrivateKeyFromHex(MOCK_OPERATOR_ECDSA_PRIVATE_KEY_HEX_STRING)
	assert.Nil(t, err)

	aggregator, _, mockEcdsaAggServ, err := createMockAggregator(mockCtrl)
	assert.Nil(t, err)

	signedTaskResponse, err := createMockSignedTaskResponse(MockTask{
		TaskNum:        TASK_INDEX,
		BlockNumber:    BLOCK_NUMBER,
		NumberToSquare: NUMBER_TO_SQUARE,
	}, MOCK_OPERATOR_ECDSA_PRIVATE_KEY)
	assert.Nil(t, err)
	signedTaskResponseDigest, err := core.GetTaskResponseDigest(&signedTaskResponse.TaskResponse)
	assert.Nil(t, err)

	// TODO(samlaf): is this the right way to test writing to external service?
	// or is there some wisdom to "don't mock 3rd party code"?
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockEcdsaAggServ.EXPECT().ProcessNewSignature(context.Background(), TASK_INDEX, signedTaskResponseDigest,
		signedTaskResponse.EcdsaSignature, signedTaskResponse.OperatorId)
	err = aggregator.ProcessSignedTaskResponse(signedTaskResponse, nil)
	assert.Nil(t, err)
}

// mocks an operator signing on a task response
func createMockSignedTaskResponse(mockTask MockTask, privateKey *ecdsa.PrivateKey) (*SignedTaskResponse, error) {
	numberToSquareBigInt := big.NewInt(int64(mockTask.NumberToSquare))
	taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: mockTask.TaskNum,
		NumberSquared:      numberToSquareBigInt.Mul(numberToSquareBigInt, numberToSquareBigInt),
	}
	taskResponseHash, err := core.GetTaskResponseDigest(taskResponse)
	if err != nil {
		return nil, err
	}
	ecdsaSignature, err := ecdsa.SignMsg(taskResponseHash[:], privateKey)
	if err != nil {
		return nil, err
	}
	signedTaskResponse := &SignedTaskResponse{
		TaskResponse:   *taskResponse,
		EcdsaSignature: ecdsaSignature,
		OperatorId:     ecdsa.PrivateKeyToOperatorId(privateKey),
	}
	return signedTaskResponse, nil
}
