package aggregator

import (
	"context"
	"errors"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	chtypes "github.com/Layr-Labs/incredible-squaring-avs/challenger/types"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type IncredibleTaskProcessor struct {
	logger           logging.Logger
	avsWriter        chainio.AvsWriterer
	tasks                 map[sdktypes.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	tasksMu               sync.RWMutex
	taskResponses      map[uint32]chtypes.TaskResponseData
}



func NewTaskProcessor(c *config.Config) (*IncredibleTaskProcessor, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &IncredibleTaskProcessor{
		logger:                c.Logger,
		avsWriter:             avsWriter,
		tasks:                 make(map[sdktypes.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses:		   make(map[uint32]chtypes.TaskResponseData),
	}, nil
}

func (tp *IncredibleTaskProcessor) ProcessNewTask(ctx context.Context, event any) (blsagg.TaskMetadata, error) {
	newTaskCreatedLog, validConversion := event.(*cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated)
	if !validConversion {
		tp.logger.Error("Log cannot not be converted into a New task created log")
		return blsagg.TaskMetadata{}, errors.New("Log cannot not be converted into a New task created log")
	}
	tp.logger.Info("Aggregator received new task", "numberToSquare", newTaskCreatedLog.Task.NumberToBeSquared)

	newTask := newTaskCreatedLog.Task
	tp.tasksMu.Lock()
	tp.tasks[newTaskCreatedLog.TaskIndex] = newTask
	tp.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block
	// number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	metadata := blsagg.NewTaskMetadata(
		newTaskCreatedLog.TaskIndex,
		newTask.TaskCreatedBlock,
		quorumNums,
		quorumThresholdPercentages,
		taskTimeToExpiry,
	)

	return metadata, nil
}

// func (tp *IncredibleTaskProcessor) ProcessTaskResponse(ctx context.Context, event any) (B256, error) {

// }

func (tp *IncredibleTaskProcessor) ProcessAggregatedResponse(ctx context.Context, response blsagg.BlsAggregationServiceResponse) error {
	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for _, nonSignerPubkey := range response.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, core.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []cstaskmanager.BN254G1Point{}
	for _, quorumApk := range response.QuorumApksG1 {
		quorumApks = append(quorumApks, core.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(response.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(response.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: response.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             response.QuorumApkIndices,
		TotalStakeIndices:            response.TotalStakeIndices,
		NonSignerStakeIndices:        response.NonSignerStakeIndices,
	}

	tp.logger.Info("Threshold reached. Sending aggregated response onchain.","taskIndex", response.TaskIndex)

	tp.tasksMu.RLock()
	task := tp.tasks[response.TaskIndex]
	tp.tasksMu.RUnlock()
	taskResponse, _ := response.TaskResponse.(cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse)
	_, err := tp.avsWriter.SendAggregatedResponse(
		context.Background(),
		task,
		taskResponse,
		nonSignerStakesAndSignature,
	)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}
