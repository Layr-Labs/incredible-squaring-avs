package operator

import (
	"context"

	allocationManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	mockAvsServiceManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockAvsServiceManager"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

func (o *Operator) SetAppointee(
	instantSlasherAddr common.Address,
	serviceManagerAddr common.Address,
	allocationManagerAddr common.Address,
	registryCoordinatorAddr common.Address,
) error {
	o.logger.Info(serviceManagerAddr.String())
	serviceManager, _ := mockAvsServiceManager.NewContractMockAvsServiceManager(serviceManagerAddr, o.ethClient)
	noSendTxOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
	waitForReceipt := true
	if err != nil {
		return err
	}
	selector := [4]byte{211, 217, 111, 244} // setAvsRegistrar call selector
	tx, err := serviceManager.SetAppointee(noSendTxOpts, o.operatorAddr, allocationManagerAddr, selector)
	if err != nil {
		return err
	}
	receipt, err := o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar appointee tx with err", err)
	}
	o.logger.Info(
		"tx successfully included for setAppointee for selector setAvsRegistrar ",
		"txHash",
		receipt.TxHash.String(),
	)

	allocationManagerContract, _ := allocationManager.NewContractAllocationManager(allocationManagerAddr, o.ethClient)

	tx, _ = allocationManagerContract.SetAVSRegistrar(noSendTxOpts, serviceManagerAddr, registryCoordinatorAddr)
	receipt, err = o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar tx with err", err)
	}
	o.logger.Info("tx successfully included for setAvsRegistrar ", "txHash", receipt.TxHash.String())

	createOperatorSetsSelector := [4]byte{38, 31, 132, 224} // createOperatorSets selector
	tx, err = serviceManager.SetAppointee(
		noSendTxOpts,
		registryCoordinatorAddr,
		allocationManagerAddr,
		createOperatorSetsSelector,
	)
	if err != nil {
		return err
	}
	receipt, err = o.avsWriter.TxMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send createOperatorSets appointee tx with err", err)
	}
	o.logger.Info(
		"tx successfully included for setAppointee for selector createOperatorSets",
		"txHash",
		receipt.TxHash.String(),
	)

	slashOperatorSelector := [4]byte{54, 53, 32, 87} // slashOperator selector
	tx, err = serviceManager.SetAppointee(
		noSendTxOpts,
		instantSlasherAddr,
		allocationManagerAddr,
		slashOperatorSelector,
	)
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
