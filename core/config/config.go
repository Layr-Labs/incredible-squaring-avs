package config

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	commonincredible "github.com/Layr-Labs/incredible-squaring-avs/common"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
)

// Config contains all of the configuration information for a credible squaring aggregators and challengers.
// Operators use a separate config. (see config-files/operator.anvil.yaml)
type Config struct {
	EcdsaPrivateKey           *ecdsa.PrivateKey
	BlsPrivateKey             *bls.PrivateKey
	Logger                    sdklogging.Logger
	EigenMetricsIpPortAddress string
	// we need the url for the eigensdk currently... eventually standardize api so as to
	// only take an ethclient or an rpcUrl (and build the ethclient at each constructor site)
	EthHttpRpcUrl                             string
	EthWsRpcUrl                               string
	EthHttpClient                             ethclient.Client
	EthWsClient                               ethclient.Client
	OperatorStateRetrieverAddr                common.Address
	IncredibleSquaringRegistryCoordinatorAddr common.Address
	IncredibleSquaringServiceManager          common.Address
	AggregatorServerIpPortAddr                string
	RegisterOperatorOnStartup                 bool
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement
	// marshalJson
	SignerFn              signerv2.SignerFn `json:"-"`
	TxMgr                 txmgr.TxManager
	AggregatorAddress     common.Address
	DelegationManagerAddr common.Address
	TokenStrategyAddr     common.Address
}

// These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                  string              `yaml:"eth_rpc_url"`
	EthWsUrl                   string              `yaml:"eth_ws_url"`
	AggregatorServerIpPortAddr string              `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup  bool                `yaml:"register_operator_on_startup"`
	delegationManagerAddr      string              `yaml:"delegation_manager_address"`
}

// These are read from CredibleSquaringDeploymentFileFlag
type IncredibleSquaringDeploymentRaw struct {
	Addresses IncredibleSquaringContractsRaw `json:"addresses"`
}
type IncredibleSquaringContractsRaw struct {
	RegistryCoordinatorAddr          string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr       string `json:"operatorStateRetriever"`
	IncredibleSquaringServiceManager string `json:"IncredibleSquaringServiceManager"`
	TokenStrategyAddr                string `json:"strategy"`
}

type EigenLayerDeploymentRaw struct {
	Addresses EigenLayerContractsRaw `json:"addresses"`
}

type EigenLayerContractsRaw struct {
	DelegationManagerAddr string `json:"delegation"`
}

// NewConfig parses config file to read from from flags or environment variables
// Note: This config is shared by challenger and aggregator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context) (*Config, error) {

	var configRaw ConfigRaw
	configFilePath := ctx.GlobalString(ConfigFileFlag.Name)
	if configFilePath != "" {
		commonincredible.ReadYamlConfig(configFilePath, &configRaw)
	}

	var credibleSquaringDeploymentRaw IncredibleSquaringDeploymentRaw
	credibleSquaringDeploymentFilePath := ctx.GlobalString(CredibleSquaringDeploymentFileFlag.Name)
	var coreDeploymentRaw EigenLayerDeploymentRaw
	coreDeploymentFilePath := ctx.GlobalString(CoreDeploymentFileFlag.Name)
	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	logger.Info(credibleSquaringDeploymentFilePath)
	if _, err := os.Stat(credibleSquaringDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + credibleSquaringDeploymentFilePath + " does not exist")
	}
	if _, err := os.Stat(coreDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + coreDeploymentFilePath + " does not exist")
	}
	commonincredible.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)
	commonincredible.ReadJsonConfig(coreDeploymentFilePath, &coreDeploymentRaw)

	if err != nil {
		return nil, err
	}

	ethRpcClient, err := ethclient.Dial(configRaw.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}

	ethWsClient, err := ethclient.Dial(configRaw.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return nil, err
	}

	ecdsaPrivateKeyString := ctx.GlobalString(EcdsaPrivateKeyFlag.Name)
	if ecdsaPrivateKeyString[:2] == "0x" {
		ecdsaPrivateKeyString = ecdsaPrivateKeyString[2:]
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivateKeyString)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}

	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		logger.Error("Cannot get operator address", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, aggregatorAddr, logger)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, aggregatorAddr)

	config := &Config{
		EcdsaPrivateKey: ecdsaPrivateKey,
		Logger:          logger,
		EthWsRpcUrl:     configRaw.EthWsUrl,
		EthHttpRpcUrl:   configRaw.EthRpcUrl,
		EthHttpClient:   *ethRpcClient,
		EthWsClient:     *ethWsClient,
		OperatorStateRetrieverAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.OperatorStateRetrieverAddr,
		),
		IncredibleSquaringRegistryCoordinatorAddr: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.RegistryCoordinatorAddr,
		),
		AggregatorServerIpPortAddr: configRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:  configRaw.RegisterOperatorOnStartup,
		IncredibleSquaringServiceManager: common.HexToAddress(
			credibleSquaringDeploymentRaw.Addresses.IncredibleSquaringServiceManager,
		),
		SignerFn:          signerV2,
		TxMgr:             txMgr,
		AggregatorAddress: aggregatorAddr,
		DelegationManagerAddr: common.HexToAddress(
			coreDeploymentRaw.Addresses.DelegationManagerAddr,
		),
		TokenStrategyAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.TokenStrategyAddr),
	}
	config.validate()
	return config, nil
}

func (c *Config) validate() {
	// TODO: make sure every pointer is non-nil
	if c.OperatorStateRetrieverAddr == common.HexToAddress("") {
		panic("Config: BLSOperatorStateRetrieverAddr is required")
	}
	if c.IncredibleSquaringRegistryCoordinatorAddr == common.HexToAddress("") {
		panic("Config: IncredibleSquaringRegistryCoordinatorAddr is required")
	}
}

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	CredibleSquaringDeploymentFileFlag = cli.StringFlag{
		Name:     "credible-squaring-deployment",
		Required: true,
		Usage:    "Load credible squaring contract addresses from `FILE`",
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "Ethereum private key",
		Required: true,
		EnvVar:   "ECDSA_PRIVATE_KEY",
	}
	CoreDeploymentFileFlag = cli.StringFlag{
		Name:     "core-deployment",
		Required: true,
		Usage:    "Load core contract addresses from `FILE`",
	}
	/* Optional Flags */
)

var requiredFlags = []cli.Flag{
	ConfigFileFlag,
	CredibleSquaringDeploymentFileFlag,
	EcdsaPrivateKeyFlag,
	CoreDeploymentFileFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
