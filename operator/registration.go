package operator

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	erc20mock "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/MockERC20"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RegisterOperatorOnStartup(nodeConfig types.NodeConfig, logger logging.Logger) error {
	ethRpcClient, err := ethclient.Dial(nodeConfig.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return err
	}

	elcontractsConfig := elcontracts.Config{
		DelegationManagerAddress:    common.HexToAddress(nodeConfig.DelegationManagerAddress),
		RewardsCoordinatorAddress:   common.HexToAddress(nodeConfig.RewardsCoordinatorAddress),
		PermissionControllerAddress: common.HexToAddress(nodeConfig.PermissionControllerAddress),
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
		logger.Fatalf("Failed to register operator with EigenLayer on startup: %v", err.Error())
	}

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	err = DepositIntoStrategyForOperator(
		logger,
		elcontractsConfig,
		ethRpcClient,
		common.HexToAddress(nodeConfig.TokenStrategyAddr),
		txMgr,
		common.HexToAddress(nodeConfig.OperatorAddress),
		amount,
	)
	if err != nil {
		logger.Fatalf("Failed to deposit into strategy for operator on startup: %v", err.Error())
	}

	blsKeyPassword, ok := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	if !ok {
		logger.Warnf("OPERATOR_BLS_KEY_PASSWORD env var not set. using empty string")
	}
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(nodeConfig.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil {
		logger.Errorf("Cannot parse bls private key", "err", err)
		return err
	}

	err = sdkoperator.RegisterForOperatorSets(
		common.HexToAddress(nodeConfig.OperatorAddress),
		logger,
		elcontractsConfig,
		ethRpcClient,
		txMgr,
		common.HexToAddress(nodeConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(nodeConfig.IncredibleSquaringServiceManager),
		[]uint32{0},
		*blsKeyPair,
		"",
	)
	if err != nil {
		logger.Fatalf("Failed to register operator for operator sets on startup: %v", err.Error())
	}

	err = sdkoperator.SetAllocationDelay(
		logger,
		common.HexToAddress(nodeConfig.OperatorAddress),
		ethRpcClient,
		common.HexToAddress(nodeConfig.AllocationManagerAddress),
		txMgr,
		0,
	)
	if err != nil {
		logger.Fatalf("Failed to set allocation delay: %v", err.Error())
	}

	return nil
}

func DepositIntoStrategyForOperator(
	logger logging.Logger,
	elcontractsConfig elcontracts.Config,
	ethClient *ethclient.Client,
	strategyAddr common.Address,
	txMgr txmgr.TxManager,
	operatorAddr common.Address,
	amount *big.Int,
) error {
	elReader, err := elcontracts.NewReaderFromConfig(elcontractsConfig, ethClient, logger)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	elWriter, err := elcontracts.NewWriterFromConfig(
		elcontractsConfig,
		ethClient,
		logger,
		&metrics.EigenMetrics{},
		txMgr,
	)
	if err != nil {
		logger.Error("Error creating eigenlayer chain writer", "err", err)
		return err
	}

	_, tokenAddr, err := elReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
	if err != nil {
		logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	logger.Info(tokenAddr.String())

	contractErc20Mock, err := erc20mock.NewContractMockERC20(tokenAddr, ethClient)
	if err != nil {
		logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := txMgr.GetNoSendTxOpts()
	if err != nil {
		logger.Errorf("Error in GetNoSendTxOpts")
		return err
	}

	tx, err := contractErc20Mock.Mint(txOpts, operatorAddr, amount)
	if err != nil {
		logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = txMgr.Send(context.Background(), tx, true)
	if err != nil {
		logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = elWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
	if err != nil {
		logger.Errorf("Error depositing into strategy", "err", err)
		return err
	}

	return nil
}
