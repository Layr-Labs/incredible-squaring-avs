package actions

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/utils"
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
		log.Fatal(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Printf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}
	operatorEcdsaPrivKey, err := sdkecdsa.ReadKey(
		nodeConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	operatorSetsIds := []uint32{nodeConfig.OperatorSetId}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		return errors.New("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(nodeConfig.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		
		return utils.WrapError("Cannot parse bls private key", err)
	}

	err = operator.RegisterForOperatorSets(
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		operatorSetsIds,
		true,
		*blsKeyPair,
		"socket",
		operatorEcdsaPrivKey,
	)
	if err != nil {
		return err
	}

	return nil
}
