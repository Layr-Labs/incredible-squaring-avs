package actions

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
)

func PrintOperatorStatus(ctx *cli.Context) error {

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

	avsConfig := avsregistry.Config{
		RegistryCoordinatorAddress: common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
	}

	logger, err := sdklogging.NewZapLogger(sdklogging.Production)
	if err != nil {
		return err
	}

	ethRpcClient, err := ethclient.Dial(nodeConfig.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	err = printOperatorStatus(
		avsConfig,
		ethRpcClient,
		logger,
		common.HexToAddress(nodeConfig.OperatorAddress),
		nodeConfig.BlsPrivateKeyStorePath,
	)
	if err != nil {
		return err
	}

	return nil
}

type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func printOperatorStatus(
	avsConfig avsregistry.Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	operatorAddr common.Address,
	blsPrivateKeyStorePath string,
) error {
	fmt.Println("Printing operator status")

	avsReader, err := avsregistry.NewReaderFromConfig(avsConfig, ethClient, logger)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	operatorId, err := avsReader.GetOperatorId(&bind.CallOpts{}, operatorAddr)
	if err != nil {
		return err
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(blsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return err
	}

	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      operatorAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          blsKeyPair.GetPubKeyG1().String(),
		G2Pubkey:          blsKeyPair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}
