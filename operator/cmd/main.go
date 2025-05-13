package main

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"os"

	"github.com/urfave/cli"

	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
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
	nodeConfig := types.NodeConfig{}
	err := commonincredible.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Config:", string(configJson))

	logger, err := sdklogging.NewZapLogger(sdklogging.Production) // Change here if want to change logging level
	if err != nil {
		return err
	}

	if nodeConfig.RegisterOperatorOnStartup {
		log.Println("Registering operator on startup")
		err = operator.RegisterOperatorOnStartup(nodeConfig, logger)
		if err != nil {
			logger.Fatalf(err.Error())
		}
	}

	logger.Info("initializing operator")

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operatorConfig := sdkoperator.OperatorConfig{
		OperatorAddress:               nodeConfig.OperatorAddress,
		OperatorStateRetrieverAddress: nodeConfig.OperatorStateRetrieverAddress,
		ServiceManagerAddress:         nodeConfig.IncredibleSquaringServiceManager,
		AVSRegistryCoordinatorAddress: nodeConfig.AVSRegistryCoordinatorAddress,
		EthRpcUrl:                     nodeConfig.EthRpcUrl,
		EthWsUrl:                      nodeConfig.EthWsUrl,
		BlsPrivateKeyStorePath:        nodeConfig.BlsPrivateKeyStorePath,
		AggregatorServerIpPortAddress: nodeConfig.AggregatorServerIpPortAddress,
		RegisterOnStartup:             true,
		Logger:                        logger,
		TaskManagerAbi:                taskManagerAbi,
	}

	calculator := sdkoperator.NewFunctionResponseCalculator(square)

	failingFunction, err := sdkoperator.NewFailingResponseCalculator(calculator, 10, big.NewInt(0))
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if err != nil {
		logger.Fatalf(err.Error())
	}

	operator, err := sdkoperator.NewOperatorFromConfig(
		operatorConfig,
		failingFunction,
		nil,
	)
	if err != nil {
		return err
	}
	log.Println("initialized operator")

	log.Println("starting operator")
	err = operator.Start(context.Background())
	if err != nil {
		return err
	}
	log.Println("started operator")

	return nil

}

// This function computes the square of a number
func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
}
