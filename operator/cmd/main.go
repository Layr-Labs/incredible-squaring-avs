package main

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/urfave/cli"

	sdkchallenger "github.com/Layr-Labs/eigensdk-go/challenger"
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

	blockHash := taskManagerAbi.Events["NewTaskCreated"].ID

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
	}

	responseCalculationFn := func(task sdkchallenger.GenericInputTask[*big.Int], taskIndex uint32) (sdkchallenger.GenericOutputTaskResponse[*big.Int], error) {
		numberSquared := big.NewInt(0).Exp(task.InputValue, big.NewInt(2), nil)

		taskResponse := sdkchallenger.GenericOutputTaskResponse[*big.Int]{
			ReferenceTaskIndex: taskIndex,
			OutputValue:        numberSquared,
		}

		return taskResponse, nil
	}

	abiEncondingFn := func(taskResponse sdkchallenger.GenericOutputTaskResponse[*big.Int]) ([]byte, error) {
		// The order here has to match the field ordering of cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
		taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{
				Name: "referenceTaskIndex",
				Type: "uint32",
			},
			{
				Name: "numberSquared",
				Type: "uint256",
			},
		})
		if err != nil {
			return nil, err
		}
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		incredibleTaskResponse := cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
			ReferenceTaskIndex: taskResponse.ReferenceTaskIndex,
			NumberSquared:      taskResponse.OutputValue,
		}

		bytes, err := arguments.Pack(incredibleTaskResponse)
		if err != nil {
			return nil, err
		}

		return bytes, nil
	}

	operator, err := sdkoperator.NewOperatorFromConfig(
		operatorConfig,
		blockHash,
		logger,
		taskManagerAbi,
		responseCalculationFn,
		abiEncondingFn,
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
