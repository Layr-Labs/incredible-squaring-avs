package chainio

import (
	"context"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewPriceUpdate(ctx context.Context) (cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask, uint32, error)
	ResigterOperatorUrl(ctx context.Context, raftHttpUrl string, raftRpcUrl string) error
	RaiseChallenge(
		ctx context.Context,
		task cstaskmanager.IIncredibleSquaringTaskManagerTask,
		taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
		taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
		pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
	) (*types.Receipt, error)
	SendAggregatedResponse(ctx context.Context,
		task cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask,
		taskResponses []cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse,
		nonSignerStakesAndSignatures []cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.IncredibleSquaringRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

func (w *AvsWriter) ResigterOperatorUrl(ctx context.Context, raftHttpUrl string, raftRpcUrl string) error {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()

	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return err
	}

	tx, err := w.AvsContractBindings.ServiceManager.RegisterOperatorConsensusUrl(txOpts, raftHttpUrl, raftRpcUrl)

	if err != nil {
		w.logger.Errorf("Error assembling RegisterOperatorConsensusUrl tx")
		return err
	}

	receipt, err := w.TxMgr.Send(ctx, tx)

	if err != nil {
		w.logger.Errorf("Error submitting RegisterOperatorConsensusUrl tx")
		return err
	}

	url, err := w.AvsContractBindings.ServiceManager.ContractIncredibleSquaringServiceManagerFilterer.ParseOperatorUrlRegistered(*receipt.Logs[0])

	if err != nil {
		w.logger.Error("Operator failed to register consensus url", "err", err)
		return err
	}

	w.logger.Info("Registered consensus url at", "url", url)

	return nil
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewPriceUpdate(ctx context.Context) (cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask{}, 0, err
	}
	feedName := "btc/usdc"
	tx, err := w.AvsContractBindings.TaskManager.RequestPriceFeedUpdate(txOpts, feedName, uint32(sdktypes.QuorumThresholdPercentage(66)), sdktypes.QuorumNums{0}.UnderlyingType(), uint8(1))
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSquaringTaskManagerFilterer.ParsePriceUpdateRequested(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) SendAggregatedResponse(
	ctx context.Context, task cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTask,
	taskResponses []cstaskmanager.IIncredibleSquaringTaskManagerPriceUpdateTaskResponse,
	nonSignerStakesAndSignatures []cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponses, nonSignerStakesAndSignatures)
	if err != nil {
		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting respondToTask tx")
		return nil, err
	}
	return receipt, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting RaiseChallenge tx")
		return nil, err
	}
	return receipt, nil
}
