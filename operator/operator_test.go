package operator

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	aggtypes "github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	chainiomocks "github.com/Layr-Labs/incredible-squaring-avs/core/chainio/mocks"
	operatormocks "github.com/Layr-Labs/incredible-squaring-avs/operator/mocks"
)

func TestOperator(t *testing.T) {
	operator, err := createMockOperator()
	assert.Nil(t, err)
	const taskIndex = 1

	t.Run("ProcessNewTaskCreatedLog", func(t *testing.T) {
		var numberToBeSquared = big.NewInt(3)
		newTaskCreatedLog := &cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.IIncredibleSquaringTaskManagerTask{
				NumberToBeSquared:         numberToBeSquared,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS,
				QuorumThresholdPercentage: aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		got := operator.ProcessNewTaskCreatedLog(newTaskCreatedLog)
		numberSquared := big.NewInt(0).Mul(numberToBeSquared, numberToBeSquared)
		want := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: taskIndex,
			NumberSquared:      numberSquared,
		}
		assert.Equal(t, got, want)
	})

	t.Run("Start", func(t *testing.T) {
		var numberToBeSquared = big.NewInt(3)

		// new task event
		newTaskCreatedEvent := &cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated{
			TaskIndex: taskIndex,
			Task: cstaskmanager.IIncredibleSquaringTaskManagerTask{
				NumberToBeSquared:         numberToBeSquared,
				TaskCreatedBlock:          1000,
				QuorumNumbers:             aggtypes.QUORUM_NUMBERS,
				QuorumThresholdPercentage: aggtypes.QUORUM_THRESHOLD_NUMERATOR,
			},
			Raw: types.Log{},
		}
		fmt.Println("newTaskCreatedEvent", newTaskCreatedEvent)

		sig, err := ecdsa.SignMsg([]byte{1}, operator.avsEcdsaPrivateKey)
		assert.Nil(t, err)
		signedTaskResponse := &aggregator.SignedTaskResponse{
			TaskResponse: cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
				ReferenceTaskIndex: taskIndex,
				NumberSquared:      big.NewInt(0).Mul(numberToBeSquared, numberToBeSquared),
			},
			EcdsaSignature: sig,
			OperatorId: operator.operatorId,
		}
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mockAggregatorRpcClient := operatormocks.NewMockAggregatorRpcClienter(mockCtrl)
		mockAggregatorRpcClient.EXPECT().SendSignedTaskResponseToAggregator(signedTaskResponse)
		operator.aggregatorRpcClient = mockAggregatorRpcClient

		mockSubscriber := chainiomocks.NewMockAvsSubscriberer(mockCtrl)
		mockSubscriber.EXPECT().SubscribeToNewTasks(operator.newTaskCreatedChan).Return(event.NewSubscription(func(quit <-chan struct{}) error {
			// loop forever
			<-quit
			return nil
		}))
		operator.avsSubscriber = mockSubscriber

		mockReader := chainiomocks.NewMockAvsReaderer(mockCtrl)
		mockReader.EXPECT().IsOperatorRegistered(gomock.Any(), operator.operatorAddr).Return(true, nil)
		operator.avsReader = mockReader

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			err := operator.Start(ctx)
			assert.Nil(t, err)
		}()
		operator.newTaskCreatedChan <- newTaskCreatedEvent
		time.Sleep(1 * time.Second)

		cancel()
	})

}
