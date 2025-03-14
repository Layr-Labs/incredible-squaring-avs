package chainio

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"

	sdkcommon "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/MockERC20"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsReaderer interface {
	//sdkavsregistry.ChainReader

	CheckSignatures(
		ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature,
	) (cstaskmanager.IBLSSignatureCheckerTypesQuorumStakeTotals, error)
	GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractMockERC20, error)
	GetOperatorId(
		opts *bind.CallOpts,
		operatorAddress common.Address,
	) ([32]byte, error)
	IsOperatorRegistered(
		opts *bind.CallOpts,
		operatorAddress common.Address,
	) (bool, error)
}

type AvsReader struct {
	sdkavsregistry.ChainReader
	AvsServiceBindings *AvsManagersBindings
	logger             logging.Logger
}

var _ AvsReaderer = (*AvsReader)(nil)

func BuildAvsReaderFromConfig(c *config.Config) (*AvsReader, error) {
	ethWsClient, err := ethclient.Dial(c.EthWsRpcUrl)
	if err != nil {
		return nil, utils.WrapError("Failed to create Eth WS client", err)
	}
	return BuildAvsReader(
		c.IncredibleSquaringRegistryCoordinatorAddr,
		c.IncredibleSquaringServiceManager,
		c.OperatorStateRetrieverAddr,
		ethWsClient,
		&c.EthHttpClient,
		c.Logger,
	)
}

func BuildAvsReader(
	registryCoordinatorAddr, serviceManagerAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	wsClient eth.WsBackend,
	ethHttpClient sdkcommon.EthClientInterface,
	logger logging.Logger,
) (*AvsReader, error) {
	avsManagersBindings, err := NewAvsManagersBindings(
		serviceManagerAddr,
		operatorStateRetrieverAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	config := sdkavsregistry.Config{
		RegistryCoordinatorAddress:    registryCoordinatorAddr,
		OperatorStateRetrieverAddress: operatorStateRetrieverAddr,
		DontUseAllocationManager:      false,
		ServiceManagerAddress:         serviceManagerAddr,
	}

	chainReader, _, _, err := sdkavsregistry.BuildReadClients(config, ethHttpClient, wsClient, logger)
	if err != nil {
		return nil, err
	}
	return NewAvsReader(*chainReader, avsManagersBindings, logger)
}

func NewAvsReader(
	avsRegistryReader sdkavsregistry.ChainReader,
	avsServiceBindings *AvsManagersBindings,
	logger logging.Logger,
) (*AvsReader, error) {
	return &AvsReader{
		ChainReader:        avsRegistryReader,
		AvsServiceBindings: avsServiceBindings,
		logger:             logger,
	}, nil
}

func (r *AvsReader) CheckSignatures(
	ctx context.Context,
	msgHash [32]byte,
	quorumNumbers []byte,
	referenceBlockNumber uint32,
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerTypesNonSignerStakesAndSignature,
) (cstaskmanager.IBLSSignatureCheckerTypesQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.AvsServiceBindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return cstaskmanager.IBLSSignatureCheckerTypesQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}

func (r *AvsReader) GetErc20Mock(
	ctx context.Context,
	tokenAddr gethcommon.Address,
) (*erc20mock.ContractMockERC20, error) {
	erc20Mock, err := r.AvsServiceBindings.GetErc20Mock(tokenAddr)
	if err != nil {
		r.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return erc20Mock, nil
}
