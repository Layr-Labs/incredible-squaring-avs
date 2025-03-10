package operator

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

func (o *Operator) CreateTotalDelegatedStakeQuorum(maxOperatorCount uint32,kickBIPsOfOperatorStake uint16,kickBIPsOfTotalStake uint16, minimumStake int64,multiplier int64) error {
	operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{MaxOperatorCount: maxOperatorCount, KickBIPsOfOperatorStake: kickBIPsOfOperatorStake, KickBIPsOfTotalStake: kickBIPsOfTotalStake}
	strategyParams := []regcoord.IStakeRegistryTypesStrategyParams{
		{
			Strategy:   common.HexToAddress(o.Config.TokenStrategyAddr),
			Multiplier: big.NewInt(multiplier),
		},
	}
	receipt, err := o.AvsWriter.CreateTotalDelegatedStakeQuorum(context.Background(), operatorSetParams,  big.NewInt(minimumStake), strategyParams, true)
	if err != nil {
		o.Logger.Error("Error creating total delegated stake quorum", "err", err)
		return err
	}
	o.Logger.Info("CreateTotalDelegatedStakeQuorum successfully included", "txHash", receipt.TxHash.String())

	return nil
}
