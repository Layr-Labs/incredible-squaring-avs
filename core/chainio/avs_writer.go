package chainio

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signer"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewTaskNumberToSquare(
		ctx context.Context,
		numToSquare *big.Int,
		quorumThresholdPercentage uint32,
		quorumNumbers []byte,
	) (cstaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error)
	RaiseChallenge(
		ctx context.Context,
		task cstaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
		pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task cstaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsServiceBindings
	logger              logging.Logger
	Signer              signer.Signer
	client              eth.EthClient
}

var _ AvsWriterer = (*AvsWriter)(nil)

func NewAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return NewAvsWriter(c.Signer, c.IncredibleSquaringServiceManagerAddr, c.BlsOperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func NewAvsWriter(signer signer.Signer, serviceManagerAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsServiceBindings(serviceManagerAddr, blsOperatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	blsRegistryCoordinatorAddr, err := avsServiceBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := avsServiceBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsPubkeyRegistryAddr, err := avsServiceBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr, blsOperatorStateRetrieverAddr, stakeRegistryAddr, blsPubkeyRegistryAddr, ethHttpClient, logger,
	)
	if err != nil {
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.NewAvsRegistryWriter(avsRegistryContractClient, logger, signer, ethHttpClient)
	if err != nil {
		return nil, err
	}

	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		Signer:              signer,
		client:              ethHttpClient,
	}, nil
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskNumberToSquare(ctx context.Context, numToSquare *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (cstaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error) {
	txOpts := w.Signer.GetTxOpts()
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, numToSquare, quorumThresholdPercentage, quorumNumbers)
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSquaringTaskManagerFilterer.ParseNewTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedResponse(ctx context.Context, task cstaskmanager.IIncredibleSquaringTaskManagerTask, taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Receipt, error) {
	txOpts := w.Signer.GetTxOpts()
	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts := w.Signer.GetTxOpts()
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt := w.client.WaitForTransactionReceipt(ctx, tx.Hash())
	return receipt, nil
}
