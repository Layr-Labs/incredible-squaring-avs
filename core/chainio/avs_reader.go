package chainio

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	logging "github.com/Layr-Labs/eigensdk-go/logging"

	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/ERC20Mock"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsReaderer interface {
	sdkavsregistry.AvsRegistryReader

	CheckSignatures(
		ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (cstaskmanager.IBLSSignatureCheckerQuorumStakeTotals, error)
	GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error)
}

type AvsReader struct {
	sdkavsregistry.AvsRegistryReader
	AvsServiceBindings *AvsServiceBindings
	logger             logging.Logger
}

var _ AvsReaderer = (*AvsReader)(nil)

func NewAvsReaderFromConfig(c *config.Config) (*AvsReader, error) {

	avsContractBindings, err := NewAvsServiceBindings(c.IncredibleSquaringServiceManagerAddr, c.BlsOperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
	if err != nil {
		return nil, err
	}
	blsRegistryCoordinatorAddr, err := avsContractBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	stakeRegistryAddr, err := avsContractBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	blsPubkeyRegistryAddr, err := avsContractBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	ethClient, err := eth.NewClient(c.EthRpcUrl)
	if err != nil {
		return nil, err
	}
	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(blsRegistryCoordinatorAddr, c.BlsOperatorStateRetrieverAddr, stakeRegistryAddr, blsPubkeyRegistryAddr, ethClient, c.Logger)
	if err != nil {
		return nil, err
	}

	avsRegistryReader, err := sdkavsregistry.NewAvsRegistryReader(avsRegistryContractClient, c.Logger, ethClient)
	if err != nil {
		return nil, err
	}
	avsServiceBindings, err := NewAvsServiceBindings(c.IncredibleSquaringServiceManagerAddr, c.BlsOperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
	if err != nil {
		return nil, err
	}

	return NewAvsReader(avsRegistryReader, avsServiceBindings, c.Logger)
}

func NewAvsReader(avsRegistryReader sdkavsregistry.AvsRegistryReader, avsServiceBindings *AvsServiceBindings, logger logging.Logger) (*AvsReader, error) {
	return &AvsReader{
		AvsRegistryReader:  avsRegistryReader,
		AvsServiceBindings: avsServiceBindings,
		logger:             logger,
	}, nil
}

func (r *AvsReader) CheckSignatures(
	ctx context.Context, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (cstaskmanager.IBLSSignatureCheckerQuorumStakeTotals, error) {
	stakeTotalsPerQuorum, _, err := r.AvsServiceBindings.TaskManager.CheckSignatures(
		&bind.CallOpts{}, msgHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature,
	)
	if err != nil {
		return cstaskmanager.IBLSSignatureCheckerQuorumStakeTotals{}, err
	}
	return stakeTotalsPerQuorum, nil
}

func (r *AvsReader) GetErc20Mock(ctx context.Context, tokenAddr gethcommon.Address) (*erc20mock.ContractERC20Mock, error) {
	erc20Mock, err := r.AvsServiceBindings.GetErc20Mock(tokenAddr)
	if err != nil {
		r.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return nil, err
	}
	return erc20Mock, nil
}
