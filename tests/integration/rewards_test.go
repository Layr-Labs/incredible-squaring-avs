package integration_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	rewardsCoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RewardsCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/Layr-Labs/incredible-squaring-avs/operator"
	"github.com/Layr-Labs/incredible-squaring-avs/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationRewards(t *testing.T) {
	log.Println("This test takes ~50 seconds to run...")

	/* Start the anvil chain */
	anvilC := startAnvilTestContainer()
	// Not sure why but deferring anvilC.Terminate() causes a panic when the test finishes...
	// so letting it terminate silently for now
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	// anvilEndpoint := "127.0.0.1:8545"

	/* Prepare the config file for aggregator */
	var aggConfigRaw config.ConfigRaw
	aggConfigFilePath := "../../config-files/aggregator.yaml"
	commonincredible.ReadYamlConfig(aggConfigFilePath, &aggConfigRaw)
	aggConfigRaw.EthRpcUrl = "http://" + anvilEndpoint
	aggConfigRaw.EthWsUrl = "ws://" + anvilEndpoint

	var credibleSquaringDeploymentRaw config.IncredibleSquaringDeploymentRaw
	credibleSquaringDeploymentFilePath := "../../contracts/script/deployments/incredible-squaring//31337.json"
	commonincredible.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(aggConfigRaw.Environment)

	if err != nil {
		t.Fatalf("Failed to create logger: %s", err.Error())
	}
	ethRpcClient, err := ethclient.Dial(aggConfigRaw.EthRpcUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}
	ethWsClient, err := ethclient.Dial(aggConfigRaw.EthWsUrl)
	if err != nil {
		t.Fatalf("Failed to create eth client: %s", err.Error())
	}

	aggregatorEcdsaPrivateKeyString := "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	if aggregatorEcdsaPrivateKeyString[:2] == "0x" {
		aggregatorEcdsaPrivateKeyString = aggregatorEcdsaPrivateKeyString[2:]
	}
	aggregatorEcdsaPrivateKey, err := crypto.HexToECDSA(aggregatorEcdsaPrivateKeyString)
	if err != nil {
		t.Fatalf("Cannot parse ecdsa private key: %s", err.Error())
	}
	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(aggregatorEcdsaPrivateKey)
	if err != nil {
		t.Fatalf("Cannot get operator address: %s", err.Error())
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		t.Fatalf("Cannot get chainId: %s", err.Error())
	}

	privateKeySigner, _, err := signerv2.SignerFromConfig(
		signerv2.Config{PrivateKey: aggregatorEcdsaPrivateKey},
		chainId,
	)
	if err != nil {
		t.Fatalf("Cannot create signer: %s", err.Error())
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, privateKeySigner, aggregatorAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, aggregatorAddr)

	config := &config.Config{
		EcdsaPrivateKey: aggregatorEcdsaPrivateKey,
		Logger:          logger,
		EthHttpRpcUrl:   aggConfigRaw.EthRpcUrl,
		EthHttpClient:   *ethRpcClient,
		EthWsRpcUrl:     aggConfigRaw.EthWsUrl,
		EthWsClient:     *ethWsClient,
		OperatorStateRetrieverAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.OperatorStateRetrieverAddr,
		),
		IncredibleSquaringRegistryCoordinatorAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.RegistryCoordinatorAddr,
		),
		AggregatorServerIpPortAddr: aggConfigRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:  aggConfigRaw.RegisterOperatorOnStartup,
		TxMgr:                      txMgr,
		AggregatorAddress:          aggregatorAddr,
		IncredibleSquaringServiceManager: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.IncredibleSquaringServiceManager,
		),
	}

	/* Prepare the config file for operator */
	nodeConfig := types.NodeConfig{}
	nodeConfigFilePath := "../../config-files/operator.anvil.yaml"
	err = commonincredible.ReadYamlConfig(nodeConfigFilePath, &nodeConfig)
	if err != nil {
		t.Fatalf("Failed to read yaml config: %s", err.Error())
	}

	/* Register operator*/
	// log.Println("registering operator for integration tests")
	// we need to do this dynamically and can't just hardcode a registered operator into the anvil
	// state because the anvil state dump doesn't also dump the receipts tree so we lose events,
	// and the aggregator thus can't get the operator's pubkey
	// operatorRegistrationCmd := exec.Command("bash", "./operator-registration.sh")
	// err = operatorRegistrationCmd.Run()
	// if err != nil {
	// 	t.Fatalf("Failed to register operator: %s", err.Error())
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
	defer cancel()
	/* start operator */
	// the passwords are set to empty strings
	log.Println("starting operator for integration tests")
	os.Setenv("OPERATOR_BLS_KEY_PASSWORD", "")
	os.Setenv("OPERATOR_ECDSA_KEY_PASSWORD", "")
	nodeConfig.BlsPrivateKeyStorePath = "../keys/test.bls.key.json"
	nodeConfig.EcdsaPrivateKeyStorePath = "../keys/test.ecdsa.key.json"
	nodeConfig.RegisterOperatorOnStartup = true
	nodeConfig.EthRpcUrl = "http://" + anvilEndpoint
	nodeConfig.EthWsUrl = "ws://" + anvilEndpoint
	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		t.Fatalf("Failed to create operator: %s", err.Error())
	}

	//
	//
	//
	//
	//
	//
	//
	//
	// Start operator
	go operator.Start(ctx)
	log.Println("Started operator. Sleeping 15 seconds to give it time to register...")
	time.Sleep(15 * time.Second)

	/* start aggregator */
	log.Println("starting aggregator for integration tests")
	agg, err := aggregator.NewAggregator(config)
	if err != nil {
		t.Fatalf("Failed to create aggregator: %s", err.Error())
	}
	go agg.Start(ctx)
	log.Println("Started aggregator. Sleeping 20 seconds to give operator time to answer task 1...")
	time.Sleep(20 * time.Second)

	// get avsRegistry client to interact with the chain
	avsReader, err := chainio.BuildAvsReaderFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Reader: %s", err.Error())
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(config)
	if err != nil {
		t.Fatalf("Cannot create AVS Writer: %s", err.Error())
	}

	mockToken, err := avsReader.GetErc20Mock(ctx, common.HexToAddress(nodeConfig.TokenAddr))
	require.NoError(t, err)

	// Check claimer balance in strategy is zero
	initialBalance, err := mockToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0000000000000000000000000000000000000001"))
	require.NoError(t, err)
	assert.Zero(t, initialBalance.Int64())

	// Now call the rewards scripts or related and assert the balance has changed.
	stratAndMul := []servicemanager.IRewardsCoordinatorTypesStrategyAndMultiplier{
		{
			Strategy:   common.HexToAddress(nodeConfig.TokenStrategyAddr),
			Multiplier: big.NewInt(1_000_000),
		},
	}
	rewardsSubmission := []servicemanager.IRewardsCoordinatorTypesRewardsSubmission{
		{
			StrategiesAndMultipliers: stratAndMul,
			Token:                    common.HexToAddress(nodeConfig.TokenAddr),
			Amount:                   big.NewInt(1000),
			StartTimestamp:           0,
			Duration:                 1010101010,
		},
	}

	receipt, err := avsWriter.CreateAVSRewardsSubmission(ctx, rewardsSubmission, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, 1)

	noSendTxOpts, err := txMgr.GetNoSendTxOpts()
	require.NoError(t, err)

	contractRewardsCoordinator, err := rewardsCoordinator.NewContractRewardsCoordinator(common.HexToAddress(nodeConfig.RewardsCoordinatorAddress), ethRpcClient)
	require.NoError(t, err)

	tx, err := contractRewardsCoordinator.SubmitRoot(noSendTxOpts, [32]byte{}, 100)
	require.NoError(t, err)

	receipt, err = txMgr.Send(ctx, tx, true)
	require.NoError(t, err)
	require.Equal(t, receipt.Status, 1)

	coreConfig := elcontracts.Config{
		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		AvsDirectoryAddress:         common.HexToAddress("0x610178da211fef7d417bc0e6fed39f05609ad788"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),
	}
	elWriter, err := elcontracts.NewWriterFromConfig(coreConfig, &config.EthHttpClient, logger, &metrics.EigenMetrics{}, txMgr)
	fmt.Printf("elWriter: %v\n", elWriter)

	//elWriter.ProcessClaim(ctx, )
}
