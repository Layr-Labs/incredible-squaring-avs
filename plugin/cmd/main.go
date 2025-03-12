package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli"
)

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
		EnvVar:   "CONFIG",
	}
	EcdsaKeyPasswordFlag = cli.StringFlag{
		Name:     "ecdsa-key-password",
		Required: false,
		Usage:    "Password to decrypt the ecdsa key",
		EnvVar:   "ECDSA_KEY_PASSWORD",
	}
	BlsKeyPasswordFlag = cli.StringFlag{
		Name:     "bls-key-password",
		Required: false,
		Usage:    "Password to decrypt the bls key",
		EnvVar:   "BLS_KEY_PASSWORD",
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

const AVS_NAME = "incredible-squaring"
const SEM_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		ConfigFileFlag,
		EcdsaKeyPasswordFlag,
		BlsKeyPasswordFlag,
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
	err := commonincredible.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avsConfig)

	ecdsaKeyPassword := ctx.GlobalString(EcdsaKeyPasswordFlag.Name)

	buildClientConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 avsConfig.EthRpcUrl,
		EthWsUrl:                   avsConfig.EthWsUrl,
		RegistryCoordinatorAddr:    avsConfig.AVSRegistryCoordinatorAddress,
		OperatorStateRetrieverAddr: avsConfig.OperatorStateRetrieverAddress,
		AvsName:                    "incredible-squaring",
		PromMetricsIpPortAddress:   avsConfig.EigenMetricsIpPortAddress,

		DontUseAllocationManager: true,
	}
	logger, _ := logging.NewZapLogger(logging.Development)
	ethHttpClient, err := ethclient.Dial(avsConfig.EthRpcUrl)
	if err != nil {
		fmt.Println("can't connect to eth client")
		fmt.Println(err)
		return
	}
	ethWsClient, err := ethclient.Dial(avsConfig.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return
	}
	chainID, err := ethHttpClient.ChainID(goCtx)
	if err != nil {
		fmt.Println("can't get chain id")
		fmt.Println(err)
		return
	}
	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{
		KeystorePath: avsConfig.EcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainID)
	if err != nil {
		fmt.Println("can't create signer")
		fmt.Println(err)
		return
	}
	operatorEcdsaPrivateKey, err := sdkecdsa.ReadKey(
		avsConfig.EcdsaPrivateKeyStorePath,
		ecdsaKeyPassword,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	clients, err := sdkclients.BuildAll(buildClientConfig, operatorEcdsaPrivateKey, logger)
	if err != nil {
		fmt.Println(err)
		return
	}
	avsReader, err := chainio.BuildAvsReader(
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.IncredibleSquaringServiceManager),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethWsClient,
		ethHttpClient,
		logger,
	)
	if err != nil {
		fmt.Println("can't create avs reader")
		fmt.Println(err)
		return
	}
	skWallet, err := wallet.NewPrivateKeyWallet(
		ethHttpClient,
		signerV2,
		common.HexToAddress(avsConfig.OperatorAddress),
		logger,
	)
	if err != nil {
		fmt.Println("can't create wallet")
		fmt.Println(err)
		return
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethHttpClient, logger, common.HexToAddress(avsConfig.OperatorAddress))
	avsWriter, err := chainio.BuildAvsWriter(
		txMgr,
		common.HexToAddress(avsConfig.IncredibleSquaringServiceManager),
		common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress),
		common.HexToAddress(avsConfig.OperatorStateRetrieverAddress),
		ethWsClient,
		ethHttpClient,
		logger,
	)
	if err != nil {
		fmt.Println("can't create avs reader")
		fmt.Println(err)
		return
	}

	if operationType == "opt-in" {
		blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

		blsKeypair, err := bls.ReadPrivateKeyFromFile(avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			logger.Error("Cannot create AggregatorRpcClient. Is aggregator running?", "err", err)
		}

		operatorSetIds := []uint32{0}
		waitForReceipt := true
		socket := "socket"

		// maxOperatorCount := 3
		// kickBpsOfOperatorStake := 100
		// kickBpsOfTotalStake := 1000
		// minimumStake := 0
		// multiplier := 1

		registrationRequest := elcontracts.RegistrationRequest{
			OperatorAddress: common.HexToAddress(avsConfig.OperatorAddress),
			AVSAddress:      common.HexToAddress(avsConfig.IncredibleSquaringServiceManager),
			OperatorSetIds:  operatorSetIds,
			WaitForReceipt:  waitForReceipt,
			BlsKeyPair:      blsKeypair,
			Socket:          socket,
		}
		r, err := clients.ElChainWriter.RegisterForOperatorSets(
			goCtx, common.HexToAddress(avsConfig.AVSRegistryCoordinatorAddress), registrationRequest)
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
		_, tokenAddr, err := clients.ElChainReader.GetStrategyAndUnderlyingToken(nil, strategyAddr)
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
		tx, err := contractErc20Mock.Mint(txOpts, common.HexToAddress(avsConfig.OperatorAddress), amount)
		if err != nil {
			logger.Errorf("Error assembling Mint tx")
			return
		}
		_, err = avsWriter.TxMgr.Send(context.Background(), tx, true)
		if err != nil {
			logger.Errorf("Error submitting Mint tx")
			return
		}

		_, err = clients.ElChainWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
		if err != nil {
			logger.Errorf("Error depositing into strategy")
			return
		}
		return
	} else {
		fmt.Println("Invalid operation type")
	}
}
