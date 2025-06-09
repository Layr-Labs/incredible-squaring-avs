package main

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.ConfigFileFlag}
	app.Name = "credible-squaring-operator"
	app.Usage = "Credible Squaring Operator"
	app.Description = "Service that reads numbers onchain, squares, signs, and sends them to the aggregator."

	app.Action = operatorMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func operatorMain(ctx *cli.Context) error {

	log.Println("Initializing Operator")
	configPath := ctx.GlobalString(config.ConfigFileFlag.Name)

	opConfig := &operator.Config{}
	err := commonincredible.ReadTomlConfig(configPath, opConfig)

	logger, err := sdklogging.NewZapLogger(sdklogging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	logger.Infof("Config is %#v", opConfig)

	logger.Info("initializing operator")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	ecdsaConfig := sdkoperator.EcdsaSignerConfig{
		KeystorePath: opConfig.EcdsaPrivateKeyStorePath,
	}

	blsConfig := sdkoperator.BlsSignerConfig{
		KeystorePath: opConfig.BlsPrivateKeyStorePath,
	}

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationCfg := sdkoperator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress(opConfig.AllocationManagerAddress),
		AvsAddress:            common.HexToAddress(opConfig.ServiceManagerAddress),
		StrategyAddrs:         []common.Address{common.HexToAddress(opConfig.TokenStrategyAddr)},

		DelegationManagerAddress:    common.HexToAddress(opConfig.DelegationManagerAddress),
		RewardsCoordinatorAddress:   common.HexToAddress(opConfig.RewardsCoordinatorAddress),
		PermissionControllerAddress: common.HexToAddress(opConfig.PermissionControllerAddress),

		EcdsaSignerCfg: ecdsaConfig,

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},
	}

	operatorConfig := sdkoperator.Config{
		OperatorAddress:               opConfig.OperatorAddress,
		RegistryCoordinatorAddress:    common.HexToAddress(opConfig.RegistryCoordinatorAddress),
		EthRpcUrl:                     opConfig.EthRpcUrl,
		EthWsUrl:                      opConfig.EthWsUrl,
		BlsSignerCfg:                  blsConfig,
		AggregatorServerIpPortAddress: opConfig.AggregatorServerIpPortAddress,

		Registration: registrationCfg,
	}

	calculator := sdkoperator.NewFunctionResponseCalculator(square)

	failingFunction, err := sdkoperator.NewFailingResponseCalculator(calculator, 10, big.NewInt(0))
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if err != nil {
		logger.Fatalf(err.Error())
	}

	operator, err := sdkoperator.NewOperator(
		logger,
		operatorConfig,
		taskManagerAbi,
		failingFunction,
		nil,
	)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	log.Println("initialized operator")

	log.Println("starting operator")
	err = <-operator.Start(context.Background())
	if err != nil {
		logger.Fatalf(err.Error())
	}
	log.Println("started operator")

	return nil

}

// This function computes the square of a number
func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
}
