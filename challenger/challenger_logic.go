package challenger

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

type ChallengerLogicImpl struct {
	logger    logging.Logger
	ethClient *ethclient.Client
	avsWriter chainio.AvsWriterer
}

//var _ sdkchallenger.ChallengerLogic = (*ChallengerLogicImpl)(nil)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

func NewChallengerLogicImpl(c *config.Config) (*ChallengerLogicImpl, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &ChallengerLogicImpl{
		logger:    c.Logger,
		ethClient: &c.EthHttpClient,
		avsWriter: avsWriter,
	}, nil
}

func (c *ChallengerLogicImpl) VerifyChallenge(
	taskIndex uint32,
	task sdkchallenger.GenericInputTask[*big.Int],
	responseData sdkchallenger.TaskResponseData[*big.Int],
) error {
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for i, pubkey := range responseData.NonSigningOperatorPubKeys {
		nonSignerPubkeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	numberToBeSquared := task.InputValue
	answerInResponse := responseData.TaskResponse.InputValue
	trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer.Cmp(answerInResponse) != 0 {
		c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

		// raise challenge
		c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)

		incredibleSquaringTask := cstaskmanager.IIncredibleSquaringTaskManagerTask{
			NumberToBeSquared:         task.InputValue,
			TaskCreatedBlock:          task.TaskCreatedBlock,
			QuorumNumbers:             task.QuorumNumbers,
			QuorumThresholdPercentage: task.QuorumThresholdPercentage,
		}

		incredibleSquaringTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: responseData.TaskResponse.ReferenceTaskIndex,
			NumberSquared:      responseData.TaskResponse.InputValue,
		}

		incredibleSquaringTaskResponseMetadata := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata{
			TaskRespondedBlock: responseData.TaskResponseMetadata.TaskRespondedBlock,
			HashOfNonSigners:   responseData.TaskResponseMetadata.HashOfNonSigners,
		}

		_, err := c.avsWriter.RaiseChallenge(
			context.Background(),
			incredibleSquaringTask,
			incredibleSquaringTaskResponse,
			incredibleSquaringTaskResponseMetadata,
			nonSignerPubkeys,
		)
		if err != nil {
			c.logger.Error("Challenger failed to raise challenge:", "err", err)
			return fmt.Errorf("challenger failed to raise challenge: %w", err)
		}

		return nil
	} else {
		c.logger.Info("The number squared is correct")
		return nil
	}
}
