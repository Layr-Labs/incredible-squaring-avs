package chainio

import (
	"context"
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	incsqtaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsecdsaregistry.AvsEcdsaRegistryWriter

	SendNewTaskNumberToSquare(
		ctx context.Context,
		numToSquare *big.Int,
		quorumThresholdPercentage uint32,
		quorumNumbers []byte,
	) (incsqtaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error)
	SendAggregatedResponse(ctx context.Context,
		task incsqtaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		signerStakeIndicesAndSignatures incsqtaskmanager.ECDSASignatureCheckerSignerStakeIndicesAndSignatures,
	) (*types.Receipt, error)
	RaiseChallenge(
		ctx context.Context,
		task incsqtaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		taskResponseMetadata incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
		signerIds []gethcommon.Address,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsecdsaregistry.AvsEcdsaRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	ethHttpClient       eth.EthClient
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.IncredibleSquaringRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.EthClient, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsecdsaregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, ethHttpClient, logger, txMgr), nil
}
func NewAvsWriter(avsEcdsaRegistryWriter avsecdsaregistry.AvsEcdsaRegistryWriter, avsServiceBindings *AvsManagersBindings, ethHttpClient eth.EthClient, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsEcdsaRegistryWriter: avsEcdsaRegistryWriter,
		AvsContractBindings:    avsServiceBindings,
		logger:                 logger,
		TxMgr:                  txMgr,
		ethHttpClient:          ethHttpClient,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskNumberToSquare(ctx context.Context, numToSquare *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (incsqtaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return incsqtaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, numToSquare, quorumThresholdPercentage, quorumNumbers)
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return incsqtaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return incsqtaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSquaringTaskManagerFilterer.ParseNewTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return incsqtaskmanager.IIncredibleSquaringTaskManagerTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task incsqtaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	signerStakeIndicesAndSignatures incsqtaskmanager.ECDSASignatureCheckerSignerStakeIndicesAndSignatures,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponse, signerStakeIndicesAndSignatures)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task incsqtaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	taskResponseMetadata incsqtaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
	signerIds []gethcommon.Address,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, signerIds)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return nil, err
	}
	return receipt, nil
}
