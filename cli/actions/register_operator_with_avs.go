package actions

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	allocationManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AllocationManager"
	mockAvsServiceManager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockAvsServiceManager"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

	operatorSetIds := []uint32{0}
	socket := "socket"
	ecdsaKeyPassword, ok := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	if !ok {
		log.Printf("OPERATOR_ECDSA_KEY_PASSWORD env var not set. using empty string")
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(nodeConfig.BlsPrivateKeyStorePath, blsKeyPassword)

	logger, err := sdklogging.NewZapLogger(sdklogging.Production) // Change here if want to change logging level
	if err != nil {
		return err
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

	err = SetAppointee(
		common.HexToAddress(nodeConfig.InstantSlasher),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		common.HexToAddress(nodeConfig.AllocationManagerAddress),
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(nodeConfig.OperatorAddress),
		logger,
		ethRpcClient,
		txMgr,
	)
	if err != nil {
		return err
	}

	maxOperatorCount := 3
	kickBpsOfOperatorStake := 100
	kickBpsOfTotalStake := 1000
	minimumStake := 0
	multiplier := 1
	err = CreateTotalDelegatedStakeQuorum(
		uint32(maxOperatorCount),
		uint16(kickBpsOfOperatorStake),
		uint16(kickBpsOfTotalStake),
		int64(minimumStake),
		int64(multiplier),
		common.HexToAddress(nodeConfig.TokenStrategyAddr),
		ethRpcClient,
		txMgr,
		logger,
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
	)
	if err != nil {
		return err
	}

	err = sdkoperator.RegisterForOperatorSets(
		common.HexToAddress(nodeConfig.OperatorAddress),
		logger,
		elcontracts.Config{},
		ethRpcClient,
		txMgr,
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		operatorSetIds,
		*blsKeyPair,
		socket,
	)
	if err != nil {
		return err
	}

	return nil
}

func SetAppointee(
	instantSlasherAddr common.Address,
	serviceManagerAddr common.Address,
	allocationManagerAddr common.Address,
	registryCoordinatorAddr common.Address,
	operatorAddr common.Address,
	logger logging.Logger,
	ethClient *ethclient.Client,
	txMgr txmgr.TxManager,
) error {
	logger.Info(serviceManagerAddr.String())
	serviceManager, _ := mockAvsServiceManager.NewContractMockAvsServiceManager(serviceManagerAddr, ethClient)
	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	waitForReceipt := true
	if err != nil {
		return err
	}
	selector := [4]byte{211, 217, 111, 244} // setAvsRegistrar call selector
	tx, err := serviceManager.SetAppointee(noSendTxOpts, operatorAddr, allocationManagerAddr, selector)
	if err != nil {
		return err
	}
	receipt, err := txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar appointee tx with err", err)
	}
	logger.Info(
		"tx successfully included for setAppointee for selector setAvsRegistrar ",
		"txHash",
		receipt.TxHash.String(),
	)

	allocationManagerContract, _ := allocationManager.NewContractAllocationManager(allocationManagerAddr, ethClient)

	tx, _ = allocationManagerContract.SetAVSRegistrar(noSendTxOpts, serviceManagerAddr, registryCoordinatorAddr)
	receipt, err = txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send setAvsRegistrar tx with err", err)
	}
	logger.Info("tx successfully included for setAvsRegistrar ", "txHash", receipt.TxHash.String())

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
	receipt, err = txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send createOperatorSets appointee tx with err", err)
	}
	logger.Info(
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
	receipt, err = txMgr.Send(context.Background(), tx, waitForReceipt)
	if err != nil {
		return utils.WrapError("failed to send slashOperator appointee tx with err", err)
	}
	logger.Info("tx successfully included for slashOperator appointee tx ", "txHash", receipt.TxHash.String())

	return nil
}

func CreateTotalDelegatedStakeQuorum(
	maxOperatorCount uint32,
	kickBIPsOfOperatorStake uint16,
	kickBIPsOfTotalStake uint16,
	minimumStake int64,
	multiplier int64,
	tokenStrategyAddr common.Address,
	ethClient *ethclient.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	registryCoordinatorAddr common.Address,
) error {
	avsConfig := avsregistry.Config{
		RegistryCoordinatorAddress: registryCoordinatorAddr,
	}

	avsWriter, err := avsregistry.NewWriterFromConfig(avsConfig, ethClient, txMgr, logger)
	if err != nil {
		logger.Error("Error creating avsregistry chain writer: ", "err", err)
		return err
	}

	operatorSetParams := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
		MaxOperatorCount:        maxOperatorCount,
		KickBIPsOfOperatorStake: kickBIPsOfOperatorStake,
		KickBIPsOfTotalStake:    kickBIPsOfTotalStake,
	}
	strategyParams := []regcoord.IStakeRegistryTypesStrategyParams{
		{
			Strategy:   tokenStrategyAddr,
			Multiplier: big.NewInt(multiplier),
		},
	}
	receipt, err := avsWriter.CreateTotalDelegatedStakeQuorum(
		context.Background(),
		operatorSetParams,
		big.NewInt(minimumStake),
		strategyParams,
		true,
	)
	if err != nil {
		logger.Error("Error creating total delegated stake quorum", "err", err)
		return err
	}
	logger.Info("CreateTotalDelegatedStakeQuorum successfully included", "txHash", receipt.TxHash.String())

	return nil
}
