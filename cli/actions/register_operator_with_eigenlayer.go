package actions

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"

	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
)

func RegisterOperatorWithEigenlayer(ctx *cli.Context) error {

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

	logger, err := sdklogging.NewZapLogger(sdklogging.Production)
	if err != nil {
		return err
	}

	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress: common.HexToAddress(nodeConfig.DelegationManagerAddress),
	}

	ethRpcClient, err := ethclient.Dial(nodeConfig.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	chainid, err := ethRpcClient.ChainID(rpcCtx)
	if err != nil {
		logger.Error("Cannot get chain id", "err", err)
		return err
	}

	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	operatorEcdsaPrivateKey, err := ecdsa.ReadKey(
		nodeConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		return err
	}

	signerV2, senderAddr, err := signerv2.SignerFromConfig(signerv2.Config{
		PrivateKey: operatorEcdsaPrivateKey,
	}, chainid)
	if err != nil {
		panic(err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, senderAddr, logger)
	if err != nil {
		return err
	}

	txMgr := txmgr.NewSimpleTxManager(pkWallet, ethRpcClient, logger, senderAddr)

	err = sdkoperator.RegisterOperatorWithEigenlayer(
		common.HexToAddress(nodeConfig.OperatorAddress),
		elcontractsConfig,
		ethRpcClient,
		logger,
		txMgr,
	)
	if err != nil {
		return err
	}

	return nil
}
