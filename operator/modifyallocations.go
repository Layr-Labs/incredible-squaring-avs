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

func (o *Operator) modifyAllocations(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	allocationManagerAddr common.Address,
	serviceManagerAddr common.Address,
	strategies []common.Address,
	newMagnitudes []uint64,
	httpUrl string,
	txMgr txmgr.TxManager,
	id uint32,
) error {
	txOpts, _ := txMgr.GetNoSendTxOpts()

	operatorAddress := crypto.PubkeyToAddress(operatorEcdsaPrivateKey.PublicKey)
	ethRpcClient, _ := ethclient.Dial(httpUrl)
	waitForReceipt := true
	allocationManagerContract, _ := allocationmanager.NewContractAllocationManager(allocationManagerAddr, ethRpcClient)
	operatorSet := allocationmanager.OperatorSet{Avs: serviceManagerAddr, Id: id}
	var allocations []allocationmanager.IAllocationManagerTypesAllocateParams
	allocations_1 := allocationmanager.IAllocationManagerTypesAllocateParams{
		OperatorSet:   operatorSet,
		Strategies:    strategies,
		NewMagnitudes: newMagnitudes,
	}
	allocations = append(allocations, allocations_1)
	tx, err := allocationManagerContract.ModifyAllocations(txOpts, operatorAddress, allocations)
	if err != nil {
		return err
	}
	receipt, err := o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send modifyAllocations tx with err", err)
	}
	o.logger.Info(
		"tx successfully included for modifyAllocations  ",
		"txHash",
		receipt.TxHash.String(),
	)

	return nil
}
