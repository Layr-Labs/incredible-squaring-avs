package challenger

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

type ChallengerLogicImpl struct {
	logger        logging.Logger
	ethClient     *ethclient.Client
	avsWriter     chainio.AvsWriterer
	tasks         map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask
	taskResponses map[uint32]TaskResponseData
}

var _ sdkchallenger.ChallengerLogic = (*ChallengerLogicImpl)(nil)

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
		logger:        c.Logger,
		ethClient:     &c.EthHttpClient,
		avsWriter:     avsWriter,
		tasks:         make(map[uint32]cstaskmanager.IIncredibleSquaringTaskManagerTask),
		taskResponses: make(map[uint32]TaskResponseData),
	}, nil
}

func (c *ChallengerLogicImpl) ProcessNewTaskCreatedLog(
	log types.Log,
) error {
	var newTaskCreatedLog cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())
	c.tasks[newTaskIndex] = newTaskCreatedLog.Task

	return nil
}

func (c *ChallengerLogicImpl) ProcessTaskResponseLog(
	log types.Log,
) error {
	var taskRespondedLog cstaskmanager.ContractIncredibleSquaringTaskManagerTaskResponded

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Fatalf("Error obtaining task manager ABI: %v", err)
	}

	err = taskManagerAbi.UnpackIntoInterface(&taskRespondedLog, "TaskResponded", log.Data)
	if err != nil {
		return fmt.Errorf("error unpacking the log: %w", err)
	}

	taskIndex := taskRespondedLog.TaskResponse.ReferenceTaskIndex

	// get the inputs necessary for raising a challenge
	nonSigningOperatorPubKeys := c.getNonSigningOperatorPubKeys(log.TxHash)
	taskResponseData := TaskResponseData{
		TaskResponse:              taskRespondedLog.TaskResponse,
		TaskResponseMetadata:      taskRespondedLog.TaskResponseMetadata,
		NonSigningOperatorPubKeys: nonSigningOperatorPubKeys,
	}

	c.taskResponses[taskIndex] = taskResponseData

	if _, found := c.tasks[taskIndex]; found {
		_ = c.verifyChallenge(taskIndex)
	}

	return nil
}

func (c *ChallengerLogicImpl) getNonSigningOperatorPubKeys(
	transactionHash common.Hash,
) []cstaskmanager.BN254G1Point {
	// get the nonSignerStakesAndSignature
	tx, _, err := c.ethClient.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		c.logger.Error("Error getting transaction by hash",
			"txHash", transactionHash,
			"err", err,
		)
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		c.logger.Error("Error getting Abi", "err", err)
	}

	calldata := tx.Data()
	methodSig := calldata[:4]
	method, err := taskManagerAbi.MethodById(methodSig)
	if err != nil {
		c.logger.Error("Error getting method", "err", err)
	}

	inputs, err := method.Inputs.Unpack(calldata[4:])
	if err != nil {
		c.logger.Error("Error unpacking calldata", "err", err)
	}

	nonSignerStakesAndSignatureInput := inputs[2].(struct {
		NonSignerQuorumBitmapIndices []uint32 "json:\"nonSignerQuorumBitmapIndices\""
		NonSignerPubkeys             []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"nonSignerPubkeys\""
		QuorumApks []struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"quorumApks\""
		ApkG2 struct {
			X [2]*big.Int "json:\"X\""
			Y [2]*big.Int "json:\"Y\""
		} "json:\"apkG2\""
		Sigma struct {
			X *big.Int "json:\"X\""
			Y *big.Int "json:\"Y\""
		} "json:\"sigma\""
		QuorumApkIndices      []uint32   "json:\"quorumApkIndices\""
		TotalStakeIndices     []uint32   "json:\"totalStakeIndices\""
		NonSignerStakeIndices [][]uint32 "json:\"nonSignerStakeIndices\""
	})

	// get pubkeys of non-signing operators and submit them to the contract
	nonSigningOperatorPubKeys := make(
		[]cstaskmanager.BN254G1Point,
		len(nonSignerStakesAndSignatureInput.NonSignerPubkeys),
	)
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
}

func (c *ChallengerLogicImpl) verifyChallenge(taskIndex uint32) error {
	numberToBeSquared := c.tasks[taskIndex].NumberToBeSquared
	answerInResponse := c.taskResponses[taskIndex].TaskResponse.NumberSquared
	trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer.Cmp(answerInResponse) != 0 {
		c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

		// raise challenge
		c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)

		_, err := c.avsWriter.RaiseChallenge(
			context.Background(),
			c.tasks[taskIndex],
			c.taskResponses[taskIndex].TaskResponse,
			c.taskResponses[taskIndex].TaskResponseMetadata,
			c.taskResponses[taskIndex].NonSigningOperatorPubKeys,
		)
		if err != nil {
			c.logger.Error("Challenger failed to raise challenge:", "err", err)
			return fmt.Errorf("challenger failed to raise challenge: %w", err)
		}

		return nil
	} else {
		c.logger.Info("The number squared is correct")
		return errors.New("100. Task response is valid")
	}
}
