package operator

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

func (o *Operator) CreateTotalDelegatedStakeQuorum() error {
	operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{MaxOperatorCount: 3, KickBIPsOfOperatorStake: 100, KickBIPsOfTotalStake: 1000}
	minimumStake := big.NewInt(0)
	strategyParams := []regcoord.IStakeRegistryTypesStrategyParams{
		{
			Strategy:   common.HexToAddress(o.config.TokenStrategyAddr),
			Multiplier: big.NewInt(1),
		},
	}
	receipt, err := o.avsWriter.CreateTotalDelegatedStakeQuorum(context.Background(), operatorSetParams, minimumStake, strategyParams, true)
	if err != nil {
		o.logger.Error("Error creating total delegated stake quorum", "err", err)
		return err
	}
	o.logger.Info("CreateTotalDelegatedStakeQuorum successfully included", "txHash", receipt.TxHash.String())

	return nil
}
