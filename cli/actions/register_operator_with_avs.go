package actions

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

func RegisterOperatorWithAvs(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := commonincredible.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	operatorSetIds := []uint32{0}
	waitForReceipt := true
	socket := "socket"
	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Printf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}
	operatorEcdsaPrivKey, err := sdkecdsa.ReadKey(
		nodeConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(nodeConfig.BlsPrivateKeyStorePath, blsKeyPassword)

	operator.SetAppointee(
		common.HexToAddress(nodeConfig.InstantSlasher),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		common.HexToAddress(nodeConfig.AllocationManagerAddress),
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
	)
	maxOperatorCount := 3
	kickBpsOfOperatorStake := 100
	kickBpsOfTotalStake := 1000
	minimumStake := 0
	multiplier := 1
	operator.CreateTotalDelegatedStakeQuorum(
		uint32(maxOperatorCount),
		uint16(kickBpsOfOperatorStake),
		uint16(kickBpsOfTotalStake),
		int64(minimumStake),
		int64(multiplier),
	)
	err = operator.RegisterForOperatorSets(
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		operatorSetIds,
		waitForReceipt,
		*blsKeyPair,
		socket,
		operatorEcdsaPrivKey,
	)
	if err != nil {
		return err
	}

	return nil
}
