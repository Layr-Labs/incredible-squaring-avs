package chainio

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkcommon "github.com/Layr-Labs/incredible-squaring-avs/common"
	csservicemanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringServiceManager"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/MockERC20"
)

type AvsManagersBindings struct {
	ServiceManagerAddr common.Address
	TaskManager        *cstaskmanager.ContractIncredibleSquaringTaskManager
	ServiceManager     *csservicemanager.ContractIncredibleSquaringServiceManager
	ethClient          eth.HttpBackend
	logger             logging.Logger
}

func NewAvsManagersBindings(
	serviceManagerAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethclient sdkcommon.EthClientInterface,
	logger logging.Logger,
) (*AvsManagersBindings, error) {
	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		serviceManagerAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}
	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}
	contractTaskManager, err := cstaskmanager.NewContractIncredibleSquaringTaskManager(taskManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch IIncredibleSquaringTaskManager contract", "err", err)
		return nil, err
	}

	return &AvsManagersBindings{
		ServiceManagerAddr: serviceManagerAddr,
		ServiceManager:     contractServiceManager,
		TaskManager:        contractTaskManager,
		ethClient:          ethclient,
		logger:             logger,
	}, nil
}

func (b *AvsManagersBindings) GetErc20Mock(tokenAddr common.Address) (*erc20mock.ContractMockERC20, error) {
	contractErc20Mock, err := erc20mock.NewContractMockERC20(tokenAddr, b.ethClient)
	if err != nil {
		b.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return contractErc20Mock, nil
}
