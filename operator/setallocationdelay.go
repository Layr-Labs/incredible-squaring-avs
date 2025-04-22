package operator

import (
	"context"
	"crypto/ecdsa"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	allocationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (o *Operator) setAllocationDelay(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	allocationManagerAddr common.Address,
	httpUrl string,
	txMgr txmgr.TxManager,
	delay uint32,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	operatorAddress := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	ethRpcClient, _ := ethclient.Dial(httpUrl)
	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethRpcClient)

	tx, err := allocationManagerContract.SetAllocationDelay(txOpts, operatorAddress, delay)
	if err != nil {
		return err
	}
	receipt, err := o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAllocationDelay tx with err", err)
	}
	o.logger.Info(
		"tx successfully included for SetAllocationDelay  ",
		"txHash",
		receipt.TxHash.String(),
	)

	return nil
}
