package aggregator

import (
	"context"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/mocks"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	chainiomocks "github.com/Layr-Labs/incredible-squaring-avs/core/chainio/mocks"
)

var MOCK_OPERATOR_ID = sdktypes.OperatorId{207, 73, 226, 221, 104, 100, 123, 41, 192, 3, 9, 119, 90, 83, 233, 159, 231, 151, 245, 96, 150, 48, 144, 27, 102, 253, 39, 101, 1, 26, 135, 173}
var MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING = "50"

// MockBlsAggregationService is a mock of BlsAggregationService interface.
type MockBlsAggregationService struct {
	ctrl     *gomock.Controller
	recorder *MockBlsAggregationServiceMockRecorder
}

type MockBlsAggregationServiceMockRecorder struct {
	mock *MockBlsAggregationService
}

func NewMockBlsAggregationService(ctrl *gomock.Controller) *MockBlsAggregationService {
	mock := &MockBlsAggregationService{ctrl: ctrl}
	mock.recorder = &MockBlsAggregationServiceMockRecorder{mock}
	return mock
}

func (m *MockBlsAggregationService) EXPECT() *MockBlsAggregationServiceMockRecorder {
	return m.recorder
}

func (m *MockBlsAggregationService) GetResponseChannel() <-chan blsagg.BlsAggregationServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseChannel")
	ret0, _ := ret[0].(<-chan blsagg.BlsAggregationServiceResponse)
	return ret0
}

func (mr *MockBlsAggregationServiceMockRecorder) GetResponseChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseChannel", reflect.TypeOf((*MockBlsAggregationService)(nil).GetResponseChannel))
}

func (m *MockBlsAggregationService) InitializeNewTask(metadata blsagg.TaskMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeNewTask", metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockBlsAggregationServiceMockRecorder) InitializeNewTask(metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeNewTask", reflect.TypeOf((*MockBlsAggregationService)(nil).InitializeNewTask), metadata)
}

func (m *MockBlsAggregationService) ProcessNewSignature(ctx context.Context, task blsagg.TaskSignature) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNewSignature", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockBlsAggregationServiceMockRecorder) ProcessNewSignature(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNewSignature", reflect.TypeOf((*MockBlsAggregationService)(nil).ProcessNewSignature), ctx, task)
}

func createMockAggregator(
	mockCtrl *gomock.Controller, operatorPubkeyDict map[sdktypes.OperatorId]types.OperatorInfo,
) (*Aggregator, *chainiomocks.MockAvsWriterer, *MockBlsAggregationService, error) {
	logger := testutils.GetTestLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockBlsAggregationService := NewMockBlsAggregationService(mockCtrl)

	aggregator := &Aggregator{
		logger:                logger,
		avsWriter:             mockAvsWriter,
		blsAggregationService: mockBlsAggregationService,
		tasks:                 make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
	}
	return aggregator, mockAvsWriter, mockBlsAggregationService, nil
}

func TestSendNewTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	MOCK_OPERATOR_BLS_PRIVATE_KEY, err := bls.NewPrivateKey(MOCK_OPERATOR_BLS_PRIVATE_KEY_STRING)
	assert.Nil(t, err)
	MOCK_OPERATOR_KEYPAIR := bls.NewKeyPair(MOCK_OPERATOR_BLS_PRIVATE_KEY)
	MOCK_OPERATOR_G1PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG1()
	MOCK_OPERATOR_G2PUBKEY := MOCK_OPERATOR_KEYPAIR.GetPubKeyG2()

	operatorPubkeyDict := map[sdktypes.OperatorId]types.OperatorInfo{
		MOCK_OPERATOR_ID: {
			OperatorPubkeys: sdktypes.OperatorPubkeys{
				G1Pubkey: MOCK_OPERATOR_G1PUBKEY,
				G2Pubkey: MOCK_OPERATOR_G2PUBKEY,
			},
			OperatorAddr: common.Address{},
		},
	}

	aggregator, mockAvsWriter, mockBlsAggService, err := createMockAggregator(mockCtrl, operatorPubkeyDict)
	assert.Nil(t, err)

	var TASK_INDEX = sdktypes.TaskIndex(0)
	var BLOCK_NUMBER = uint32(100)
	var NUMBER_TO_SQUARE = uint32(3)
	var NUMBER_TO_SQUARE_BIG_INT = big.NewInt(int64(NUMBER_TO_SQUARE))

	mockAvsWriter.EXPECT().SendNewTaskNumberToSquare(
		context.Background(), NUMBER_TO_SQUARE_BIG_INT, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS,
	).Return(mocks.MockSendNewTaskNumberToSquareCall(BLOCK_NUMBER, TASK_INDEX, NUMBER_TO_SQUARE))

	taskMetadata := blsagg.NewTaskMetadata(TASK_INDEX, BLOCK_NUMBER, types.QUORUM_NUMBERS, sdktypes.QuorumThresholdPercentages{types.QUORUM_THRESHOLD_NUMERATOR}, 100*12*time.Second)
	mockBlsAggService.EXPECT().InitializeNewTask(taskMetadata)

	err = aggregator.sendNewTask(NUMBER_TO_SQUARE_BIG_INT)
	assert.Nil(t, err)
}
