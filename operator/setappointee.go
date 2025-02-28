package operator

import (
	"context"

	allocationManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	mockAvsServiceManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockAvsServiceManager"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (o *Operator) SetAppointee() error {
	// request:= elcontracts.SetPermissionRequest{}
	serviceManagerAddr := common.HexToAddress("0xb7278A61aa25c888815aFC32Ad3cC52fF24fE575")
	allocationManagerAddr := common.HexToAddress("0x8a791620dd6260079bf849dc5567adc3f2fdc318")
	registryCoordinatorAddr := common.HexToAddress("0xc351628eb244ec633d5f21fbd6621e1a683b1181")
	serviceManager, err := mockAvsServiceManager.NewContractMockAvsServiceManager(serviceManagerAddr, o.ethClient)
	noSendTxOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
	waitForReceipt := true
	if err != nil {
		return err
	}
	selector := [4]byte{211, 217, 111, 244}
	tx, err := serviceManager.SetAppointee(noSendTxOpts, o.operatorAddr, allocationManagerAddr, selector)
	if err != nil {
		return err
	}
	receipt, err := o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar appointee tx with err", err)
	}
	o.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	allocationManagerContract, err := allocationManager.NewContractAllocationManager(allocationManagerAddr, o.ethClient)

	tx, err = allocationManagerContract.SetAVSRegistrar(noSendTxOpts, serviceManagerAddr, registryCoordinatorAddr)
	receipt, err = o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar tx with err", err)
	}
	o.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	createOperatorSetsSelector := [4]byte{38, 31, 132, 224}
	tx, err = serviceManager.SetAppointee(noSendTxOpts, registryCoordinatorAddr, allocationManagerAddr, createOperatorSetsSelector)
	if err != nil {
		return err
	}
	receipt, err = o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send createOperatorSets appointee tx with err", err)
	}
	o.logger.Info("tx successfully included", "txHash", receipt.TxHash.String())

	instantSlasherAddr := common.HexToAddress("0x1fa02b2d6a771842690194cf62d91bdd92bfe28d")
	slashOperatorSelector := [4]byte{54, 53, 32, 87}
	tx, err = serviceManager.SetAppointee(noSendTxOpts, instantSlasherAddr, allocationManagerAddr, slashOperatorSelector)
	if err != nil {
		return err
	}
	receipt, err = o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send slashOperator appointee tx with err", err)
	}
	o.logger.Info("tx successfully included for slashOperator appointee tx ", "txHash", receipt.TxHash.String())

	return nil
}
