package taskgenerator

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	csservicemanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringServiceManager"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Task generator logic code
type TaskGenLogic struct {
	avsWriter          *AvsWriter
	thresholdNumerator uint8
	quorumNumbers      []uint8
	logger             logging.Logger
}

func NewTaskGenLogic(c *AvsConfig, thresholdNumerator uint8, quorumNumbers []uint8) (*TaskGenLogic, error) {
	avsWriter, err := BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &TaskGenLogic{
		avsWriter, thresholdNumerator, quorumNumbers, c.Logger}, nil
}

func (tgl *TaskGenLogic) SendNewTask(taskNumber int64) error {
	err := tgl.avsWriter.SendNewTaskNumberToSquare(context.Background(), big.NewInt(taskNumber),
		tgl.thresholdNumerator, tgl.quorumNumbers)
	if err != nil {
		tgl.logger.Error("TaskGenerator failed to send number to square", "err", err)
		return err
	}

	return nil
}

func GetTxManager(logger logging.Logger, ethHttpClient *ethclient.Client, taskgeneratorPk string) (*txmgr.SimpleTxManager, error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(taskgeneratorPk)
	if err != nil {
		return nil, err
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethHttpClient.ChainID(rpcCtx)
	if err != nil {
		logger.Error("Cannot get chain id", "err", err)
		return nil, err
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		return nil, err
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethHttpClient, signerV2, senderAddr, logger)
	if err != nil {
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethHttpClient, logger, senderAddr)
	return txMgr, nil
}

// Avs Writer code
type AvsWriter struct {
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	taskManagerContract *cstaskmanager.ContractIncredibleSquaringTaskManager
}

type AvsConfig struct {
	Logger                        logging.Logger
	IncredibleSquaringServiceManager common.Address
	TxMgr                         txmgr.TxManager
	EthHttpClient                 *ethclient.Client
}

func BuildAvsWriterFromConfig(c *AvsConfig) (*AvsWriter, error) {
	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		c.IncredibleSquaringServiceManager,
		c.EthHttpClient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch IServiceManager contract", err)
	}

	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})
	if err != nil {
		c.Logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}
	contractTaskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(taskManagerAddr, c.EthHttpClient)
	if err != nil {
		c.Logger.Error("Failed to fetch IncredibleSquaringTaskManager contract", "err", err)
		return nil, err
	}

	return &AvsWriter{
		logger:              c.Logger,
		TxMgr:               c.TxMgr,
		taskManagerContract: contractTaskManager,
	}, nil
}

func (w *AvsWriter) SendNewTaskNumberToSquare(
	ctx context.Context,
	numToSquare *big.Int,
	quorumThresholdPercentage uint8,
	quorumNumbers []uint8,
) error {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	tx, err := w.taskManagerContract.CreateNewTask(
		txOpts,
		numToSquare,
		uint32(quorumThresholdPercentage),
		quorumNumbers,
	)
	if err != nil {
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}
	_, err = w.TxMgr.Send(ctx, tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	return nil
}
