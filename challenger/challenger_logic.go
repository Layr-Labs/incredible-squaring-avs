package challenger

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

type ChallengerLogicImpl struct {
	logger        logging.Logger
	ethClient     *ethclient.Client
	avsWriter     chainio.AvsWriterer
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
		logger:        c.Logger,
		ethClient:     &c.EthHttpClient,
		avsWriter:     avsWriter,
	}, nil
}

func (c *ChallengerLogicImpl) GetNonSigningOperatorPubKeys(
	transactionHash common.Hash,
) []sdkchallenger.BN254G1Point {
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
		[]sdkchallenger.BN254G1Point,
		len(nonSignerStakesAndSignatureInput.NonSignerPubkeys),
	)
	for i, pubkey := range nonSignerStakesAndSignatureInput.NonSignerPubkeys {
		nonSigningOperatorPubKeys[i] = sdkchallenger.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	return nonSigningOperatorPubKeys
}

func (c *ChallengerLogicImpl) VerifyChallenge(taskIndex uint32, task sdkchallenger.GenericTask, responseData sdkchallenger.TaskResponseData) error {
	incredibleSquaringTask, ok := task.(cstaskmanager.IIncredibleSquaringTaskManagerTask)
	if !ok {
		return fmt.Errorf("task was not a incredible squaring task. Task %v", task)
	}
	incredibleSquaringTaskResponse, ok := responseData.TaskResponse.(cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse)
	if !ok {
		return fmt.Errorf("task was not a incredible squaring task response. Task response: %v", responseData.TaskResponse)
	}
	incredibleSquaringTaskResponseMetadata, ok := responseData.TaskResponseMetadata.(cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata)
	if !ok {
		return fmt.Errorf("task was not a incredible squaring task response metadata. Task response metadata: %v", responseData.TaskResponseMetadata)
	}
	nonSignerPubkeys := []cstaskmanager.BN254G1Point{}
	for i, pubkey := range responseData.NonSigningOperatorPubKeys {
		nonSignerPubkeys[i] = cstaskmanager.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}
	}

	numberToBeSquared := incredibleSquaringTask.NumberToBeSquared
	answerInResponse := incredibleSquaringTaskResponse.NumberSquared
	trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer.Cmp(answerInResponse) != 0 {
		c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

		// raise challenge
		c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)

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
