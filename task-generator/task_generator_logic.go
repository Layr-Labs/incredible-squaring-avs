package taskgenerator

import (
	"context"
	"math/big"
	"time"

	sdktaskgenerator "github.com/Layr-Labs/eigensdk-go/task-generator"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TaskGenLogic struct {
	avsWriter          chainio.AvsWriterer
	thresholdNumerator types.QuorumThresholdPercentage
	quorumNumbers      types.QuorumNums
	logger             logging.Logger
}

var _ sdktaskgenerator.TaskGeneratorLogic = (*TaskGenLogic)(nil)

func NewTaskGenLogic(c *config.Config, thresholdNumerator types.QuorumThresholdPercentage, quorumNumbers types.QuorumNums) (*TaskGenLogic, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	return &TaskGenLogic{
		avsWriter, thresholdNumerator, quorumNumbers, c.Logger}, nil
}

func (tgl *TaskGenLogic) SendNewTask(taskNumber int64) error {
	_, _, err := tgl.avsWriter.SendNewTaskNumberToSquare(
		context.Background(),
		big.NewInt(taskNumber),
		tgl.thresholdNumerator,
		tgl.quorumNumbers,
	)
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
