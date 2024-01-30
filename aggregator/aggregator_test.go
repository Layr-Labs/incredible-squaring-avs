package aggregator

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	ecdsaservmock "github.com/Layr-Labs/eigensdk-go/services/ecdsa/mocks"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/mocks"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	chainiomocks "github.com/Layr-Labs/incredible-squaring-avs/core/chainio/mocks"
)

var MOCK_OPERATOR_STAKE = big.NewInt(100)
var MOCK_OPERATOR_ECDSA_PRIVATE_KEY_HEX_STRING = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

type MockTask struct {
	TaskNum        uint32
	BlockNumber    uint32
	NumberToSquare uint32
}

func TestSendNewTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	aggregator, mockAvsWriterer, mockBlsAggService, err := createMockAggregator(mockCtrl)
	assert.Nil(t, err)

	var TASK_INDEX = uint32(0)
	var BLOCK_NUMBER = uint32(100)
	var NUMBER_TO_SQUARE = uint32(3)
	var NUMBER_TO_SQUARE_BIG_INT = big.NewInt(int64(NUMBER_TO_SQUARE))

	mockAvsWriterer.EXPECT().SendNewTaskNumberToSquare(
		context.Background(), NUMBER_TO_SQUARE_BIG_INT, types.QUORUM_THRESHOLD_NUMERATOR, types.QUORUM_NUMBERS,
	).Return(mocks.MockSendNewTaskNumberToSquareCall(BLOCK_NUMBER, TASK_INDEX, NUMBER_TO_SQUARE))

	// 100 blocks, each takes 12 seconds. We hardcode for now since aggregator also hardcodes this value
	taskTimeToExpiry := 100 * 12 * time.Second
	// make sure that initializeNewTask was called on the blsAggService
	// maybe there's a better way to do this? There's a saying "don't mock 3rd party code"
	// see https://hynek.me/articles/what-to-mock-in-5-mins/
	mockBlsAggService.EXPECT().InitializeNewTask(TASK_INDEX, BLOCK_NUMBER, types.QUORUM_NUMBERS, []uint32{types.QUORUM_THRESHOLD_NUMERATOR}, taskTimeToExpiry)

	err = aggregator.sendNewTask(NUMBER_TO_SQUARE_BIG_INT)
	assert.Nil(t, err)
}

func createMockAggregator(
	mockCtrl *gomock.Controller,
) (*Aggregator, *chainiomocks.MockAvsWriterer, *ecdsaservmock.MockEcdsaAggregationService, error) {
	logger := sdklogging.NewNoopLogger()
	mockAvsWriter := chainiomocks.NewMockAvsWriterer(mockCtrl)
	mockEcdsaAggregationService := ecdsaservmock.NewMockEcdsaAggregationService(mockCtrl)

	aggregator := &Aggregator{
		logger:                  logger,
		avsWriter:               mockAvsWriter,
		ecdsaAggregationService: mockEcdsaAggregationService,
		tasks:                   make(map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:           make(map[types.TaskIndex]map[sdktypes.TaskResponseDigest]cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse),
	}
	return aggregator, mockAvsWriter, mockEcdsaAggregationService, nil
}

// just a mock ethclient to pass to bindings
// so that we can access abi methods
func createMockEthClient() *backends.SimulatedBackend {
	genesisAlloc := map[common.Address]gethcore.GenesisAccount{}
	blockGasLimit := uint64(1000000)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	return client
}
