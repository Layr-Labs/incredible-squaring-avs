package aggregator

import (
	"context"
	"math/big"

	sdkaggregator "github.com/Layr-Labs/eigensdk-go/aggregator"
	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/utils"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type IncredibleTaskProcessor struct {
	logger    logging.Logger
	avsWriter chainio.AvsWriterer
}

var _ sdkaggregator.TaskProcessor[*big.Int] = (*IncredibleTaskProcessor)(nil)

func NewTaskProcessor(c *config.Config) (*IncredibleTaskProcessor, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &IncredibleTaskProcessor{
		logger:    c.Logger,
		avsWriter: avsWriter,
	}, nil
}

func (tp *IncredibleTaskProcessor) ProcessTaskResponse(
	ctx context.Context,
	response sdkaggregator.TaskResponse,
) ([32]byte, error) {
	return response.Digest(), nil
}

// This method sends the aggregated response to the on-chain Task Manager contract
func (tp *IncredibleTaskProcessor) ProcessAggregatedResponse(
	ctx context.Context,
	response blsagg.BlsAggregationServiceResponse,
	task sdkchallenger.GenericInputTask[*big.Int],
) error {
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

	tp.logger.Info("Threshold reached. Sending aggregated response onchain.", "taskIndex", response.TaskIndex)

	taskResponseAgg, ok := response.TaskResponse.(sdkchallenger.GenericOutputTaskResponse[*big.Int])
	if !ok {
		tp.logger.Error("task Response could not be converted to sdk aggregator's Task Response type")
	}

	taskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
		ReferenceTaskIndex: taskResponseAgg.ReferenceTaskIndex,
		NumberSquared:      taskResponseAgg.OutputValue,
	}

	incredibleSquaringTask := cstaskmanager.IIncredibleSquaringTaskManagerTask{
		NumberToBeSquared:         task.InputValue,
		TaskCreatedBlock:          task.TaskCreatedBlock,
		QuorumNumbers:             task.QuorumNumbers,
		QuorumThresholdPercentage: task.QuorumThresholdPercentage,
	}

	_, err := tp.avsWriter.SendAggregatedResponse(
		context.Background(),
		incredibleSquaringTask,
		taskResponse,
		nonSignerStakesAndSignature,
	)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}
