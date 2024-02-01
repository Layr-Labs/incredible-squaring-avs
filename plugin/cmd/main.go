package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSARegistryCoordinator"
	stakereg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAStakeRegistry"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
)

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
		EnvVar:   "CONFIG",
	}
	AvsEcdsaKeyPasswordFlag = cli.StringFlag{
		Name:     "tx-ecdsa-key-password",
		Required: false,
		Usage:    "Password to decrypt the avs ecdsa key (key associated with the incredible-squaring avs)",
		EnvVar:   "AVS_ECDSA_KEY_PASSWORD",
	}
	OperatorEcdsaKeyPasswordFlag = cli.StringFlag{
		Name:     "operator-ecdsa-key-password",
		Required: false,
		Usage:    "Password to decrypt the operator ecdsa key (key associated with the operator address)",
		EnvVar:   "OPERATOR_ECDSA_KEY_PASSWORD",
	}
	OperationFlag = cli.StringFlag{
		Name:     "operation-type",
		Required: true,
		Usage:    "Supported operations: opt-in, deposit",
		EnvVar:   "OPERATION_TYPE",
	}
	StrategyAddrFlag = cli.StringFlag{
		Name:     "strategy-addr",
		Required: false,
		Usage:    "Strategy address for deposit mock tokens, only used for deposit action",
		EnvVar:   "STRATEGY_ADDR",
	}
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		ConfigFileFlag,
		AvsEcdsaKeyPasswordFlag,
		OperatorEcdsaKeyPasswordFlag,
		OperationFlag,
		StrategyAddrFlag,
	}
	app.Name = "credible-squaring-plugin"
	app.Usage = "Credible Squaring Plugin"
	app.Description = "This is used to run one time operations like avs opt-in/opt-out"
	app.Action = plugin
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed.", "Message:", err)
	}
}

func plugin(ctx *cli.Context) {
	goCtx := context.Background()

	operationType := ctx.GlobalString(OperationFlag.Name)
	configPath := ctx.GlobalString(ConfigFileFlag.Name)

	avsConfig := types.NodeConfig{}
	err := utils.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avsConfig)

	logger, _ := logging.NewZapLogger(logging.Development)
	ethHttpClient, err := eth.NewClient(avsConfig.EthRpcUrl)
	if err != nil {
		fmt.Println("can't connect to eth client")
		fmt.Println(err)
		return
	}
	chainID, err := ethHttpClient.ChainID(goCtx)
	if err != nil {
		fmt.Println("can't get chain id")
		fmt.Println(err)
		return
	}
	ecdsaKeyPassword := ctx.GlobalString(OperatorEcdsaKeyPasswordFlag.Name)
	operatorAddress, err := sdkecdsa.GetAddressFromKeyStoreFile(avsConfig.OperatorEcdsaPrivateKeyStorePath)
	if err != nil {
		fmt.Println("can't get operator address")
		fmt.Println(err)
		return
	}
	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: avsConfig.OperatorEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainID)
	if err != nil {
		fmt.Println("can't create signer")
		fmt.Println(err)
		return
	}
	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		fmt.Println("can't create avs reader")
		fmt.Println(err)
		return
	}
	txMgr := txmgr.NewSimpleTxManager(ethHttpClient, logger, signerV2, operatorAddress)
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr,
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		fmt.Println("can't create avs reader")
		fmt.Println(err)
		return
	}

	registryCoordinator, err := regcoord.NewContractECDSARegistryCoordinator(common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress), ethHttpClient)
	if err != nil {
		logger.Fatal("Failed to create registry coordinator", "err", err)
	}
	stakeRegistryAddr, err := registryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch stake registry address", "err", err)
	}
	stakeRegistry, err := stakereg.NewContractECDSAStakeRegistry(stakeRegistryAddr, ethHttpClient)
	if err != nil {
		logger.Fatal("Failed to create stake registry", "err", err)
	}
	delegationManagerAddr, err := stakeRegistry.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Fatal("Failed to fetch delegation manager address", "err", err)
	}
	elChainReader, err := elcontracts.BuildELChainReader(delegationManagerAddr, ethHttpClient, logger)
	if err != nil {
		logger.Fatal("Failed to create ELChainReader", "err", err)
	}
	noopMetrics := metrics.NewNoopMetrics()
	elChainWriter, err := elcontracts.BuildELChainWriter(delegationManagerAddr, ethHttpClient, logger, noopMetrics, txMgr)
	if err != nil {
		logger.Fatal("Failed to create ELChainWriter", "err", err)
	}

	if operationType == "opt-in" {

		avsEcdsaKeyPassword := ctx.GlobalString(AvsEcdsaKeyPasswordFlag.Name)
		avsEcdsaKeypair, err := sdkecdsa.ReadKey(avsConfig.AvsEcdsaPrivateKeyStorePath, avsEcdsaKeyPassword)
		if err != nil {
			fmt.Println(err)
			return
		}

		operatorEcdsaKeyPassword := ctx.GlobalString(OperatorEcdsaKeyPasswordFlag.Name)
		operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
			avsConfig.OperatorEcdsaPrivateKeyStorePath,
			operatorEcdsaKeyPassword,
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Register with registry coordination
		quorumNumbers := []byte{0}
		socket := "Not Needed"
		sigValidForSeconds := int64(1_000_000)
		operatorToAvsRegistrationSigSalt := [32]byte{123}
		operatorToAvsRegistrationSigExpiry := big.NewInt(int64(time.Now().Unix()) + sigValidForSeconds)
		logger.Infof("Registering with registry coordination with quorum numbers %v and socket %s", quorumNumbers, socket)
		r, err := avsWriter.AvsEcdsaRegistryWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
			goCtx,
			operatorEcdsaPrivateKey, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
			avsEcdsaKeypair, quorumNumbers,
		)
		if err != nil {
			logger.Errorf("Error assembling CreateNewTask tx")
			fmt.Println(err)
			return
		}
		logger.Infof("Registered with registry coordination successfully with tx hash %s", r.TxHash.Hex())
	} else if operationType == "opt-out" {
		fmt.Println("Opting out of slashing - unimplemented")
	} else if operationType == "deposit" {
		starategyAddrString := ctx.GlobalString(StrategyAddrFlag.Name)
		if len(starategyAddrString) == 0 {
			fmt.Println("Strategy address is required for deposit operation")
			return
		}
		strategyAddr := common.HexToAddress(ctx.GlobalString(StrategyAddrFlag.Name))
		_, tokenAddr, err := elChainReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
		if err != nil {
			logger.Error("Failed to fetch strategy contract", "err", err)
			return
		}
		contractErc20Mock, err := avsReader.GetErc20Mock(context.Background(), tokenAddr)
		if err != nil {
			logger.Error("Failed to fetch ERC20Mock contract", "err", err)
			return
		}
		txOpts, err := avsWriter.TxMgr.GetNoSendTxOpts()
		if err != nil {
			logger.Errorf("Error getting tx opts")
			return
		}
		amount := big.NewInt(1000)
		tx, err := contractErc20Mock.Mint(txOpts, operatorAddress, amount)
		if err != nil {
			logger.Errorf("Error assembling Mint tx")
			return
		}
		_, err = avsWriter.TxMgr.Send(context.Background(), tx)
		if err != nil {
			logger.Errorf("Error submitting Mint tx")
			return
		}

		_, err = elChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
		if err != nil {
			logger.Errorf("Error depositing into strategy")
			return
		}
		return
	} else {
		fmt.Println("Invalid operation type")
	}
}
