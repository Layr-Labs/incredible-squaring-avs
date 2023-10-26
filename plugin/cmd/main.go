package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	sdkavsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdkelcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signer"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	err := utils.ReadYamlConfig(configPath, &avsConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avsConfig)

	ecdsaKeyPassword := ctx.GlobalString(EcdsaKeyPasswordFlag.Name)

	ethHttpClient, err := eth.NewClient(avsConfig.EthRpcUrl)
	if err != nil {
		fmt.Println("can't connect to eth client")
		fmt.Println(err)
		return
	}
	ethWsClient, err := eth.NewClient(avsConfig.EthWsUrl)
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

	signer, err := signer.NewPrivateKeyFromKeystoreSigner(avsConfig.EcdsaPrivateKeyStorePath, ecdsaKeyPassword, chainID)
	if err != nil {
		fmt.Println("can't create signer")
		fmt.Println(err)
		return
	}

	logger, _ := logging.NewZapLogger(logging.Development)
	avsWriter, err := chainio.NewAvsWriter(signer,
		common.HexToAddress(avsConfig.AVSServiceManagerAddress),
		common.HexToAddress(avsConfig.BLSOperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)

	if err != nil {
		fmt.Println("can't create avs writer")
		fmt.Println(err)
		return
	}

	avsServiceBindings, err := chainio.NewAvsServiceBindings(
		common.HexToAddress(avsConfig.AVSServiceManagerAddress),
		common.HexToAddress(avsConfig.BLSOperatorStateRetrieverAddress),
		ethHttpClient,
		logger,
	)
	if err != nil {
		fmt.Println("can't create avs service bindings")
		fmt.Println(err)
		return
	}

	blsRegistryCoordinatorAddr, err := avsServiceBindings.ServiceManager.RegistryCoordinator(&bind.CallOpts{})
	if err != nil {
		fmt.Println("can't get bls registry coordinator address")
		fmt.Println(err)
		return
	}

	stakeRegistryAddr, err := avsServiceBindings.ServiceManager.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		fmt.Println("can't get stake registry address")
		fmt.Println(err)
		return
	}
	blsPubkeyRegistryAddr, err := avsServiceBindings.ServiceManager.BlsPubkeyRegistry(&bind.CallOpts{})
	if err != nil {
		fmt.Println("can't get bls pubkey registry address")
		fmt.Println(err)
		return
	}

	avsRegistryContractClient, err := sdkclients.NewAvsRegistryContractsChainClient(
		blsRegistryCoordinatorAddr, common.HexToAddress(avsConfig.BLSOperatorStateRetrieverAddress), stakeRegistryAddr, blsPubkeyRegistryAddr, ethHttpClient, logger,
	)
	if err != nil {
		fmt.Println("can't create avs registry contract client")
		fmt.Println(err)
		return
	}

	avsRegistryReader, err := sdkavsregistry.NewAvsRegistryReader(avsRegistryContractClient, logger, ethHttpClient)
	if err != nil {
		fmt.Println("can't create avs registry reader")
		fmt.Println(err)
		return
	}

	avsReader, err := chainio.NewAvsReader(
		avsRegistryReader,
		avsServiceBindings,
		logger,
	)

	if err != nil {
		fmt.Println("can't create avs writer")
		fmt.Println(err)
		return
	}

	elContractsClient, err := sdkclients.NewELContractsChainClient(common.HexToAddress(avsConfig.ELSlasherAddress), common.HexToAddress(avsConfig.BlsPublicKeyCompendiumAddress), ethHttpClient, ethWsClient, logger)
	if err != nil {
		logger.Error("Cannot create ELContractsChainClient", "err", err)
		return
	}
	noopMetrics := metrics.NewNoopMetrics()
	eigenLayerWriter := sdkelcontracts.NewELChainWriter(
		elContractsClient,
		ethHttpClient,
		signer,
		logger,
		noopMetrics,
	)

	eigenlayerReader, err := sdkelcontracts.NewELChainReader(
		elContractsClient,
		logger,
		ethHttpClient,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println("can't create eigenlayer writer")
		fmt.Println(err)
		return
	}

	if operationType == "opt-in" {
		blsKeyPassword := ctx.GlobalString(BlsKeyPasswordFlag.Name)

		keypair, err := bls.ReadPrivateKeyFromFile(avsConfig.BlsPrivateKeyStorePath, blsKeyPassword)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Opt into slashing
		logger.Infof("Opting into slashing with AVS Service Manager address %s", avsConfig.AVSServiceManagerAddress)
		_, err = eigenLayerWriter.OptOperatorIntoSlashing(goCtx, common.HexToAddress(avsConfig.AVSServiceManagerAddress))
		if err != nil {
			fmt.Println(err)
			return
		}
		logger.Infof("Opted into slashing successfully")

		// Register with registry coordination
		quorumNumbers := []byte{0}
		socket := "Not Needed"

		pubkey := pubKeyG1ToBN254G1Point(keypair.GetPubKeyG1())
		g1Point := regcoord.BN254G1Point{
			X: pubkey.X,
			Y: pubkey.Y,
		}

		logger.Infof("Registering with registry coordination with quorum numbers %v and socket %s", quorumNumbers, socket)
		r, err := avsWriter.RegisterOperatorWithAVSRegistryCoordinator(goCtx, quorumNumbers, g1Point, socket)
		if err != nil {
			fmt.Println(err)
			logger.Errorf("Error assembling CreateNewTask tx")
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
		_, tokenAddr, err := eigenlayerReader.GetStrategyAndUnderlyingToken(context.Background(), strategyAddr)
		if err != nil {
			logger.Error("Failed to fetch strategy contract", "err", err)
			return
		}
		contractErc20Mock, err := avsReader.GetErc20Mock(context.Background(), tokenAddr)
		if err != nil {
			logger.Error("Failed to fetch ERC20Mock contract", "err", err)
			return
		}
		txOpts := avsWriter.Signer.GetTxOpts()
		amount := big.NewInt(1000)
		tx, err := contractErc20Mock.Mint(txOpts, common.HexToAddress(avsConfig.OperatorAddress), amount)
		if err != nil {
			logger.Errorf("Error assembling Mint tx")
			return
		}
		ethHttpClient.WaitForTransactionReceipt(context.Background(), tx.Hash())

		_, err = eigenLayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount)
		if err != nil {
			logger.Errorf("Error depositing into strategy")
			return
		}
		return
	} else {
		fmt.Println("Invalid operation type")
	}
}

func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
	return regcoord.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}
