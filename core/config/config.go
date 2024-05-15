package config

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

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
	EthHttpClient                             eth.Client
	EthWsClient                               eth.Client
	OperatorStateRetrieverAddr                common.Address
	IncredibleSquaringRegistryCoordinatorAddr common.Address
	AggregatorServerIpPortAddr                string
	RegisterOperatorOnStartup                 bool
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement marshalJson
	SignerFn          signerv2.SignerFn `json:"-"`
	TxMgr             txmgr.TxManager
	AggregatorAddress common.Address
}

// These are read from CredibleSquaringDeploymentFileFlag
type IncredibleSquaringDeploymentRaw struct {
	Addresses IncredibleSquaringContractsRaw `json:"addresses"`
}
type IncredibleSquaringContractsRaw struct {
	RegistryCoordinatorAddr    string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr string `json:"operatorStateRetriever"`
}

// NewConfig parses config file to read from flags or environment variables
// Note: This config is shared by challenger and aggregator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context) (*Config, error) {
	Environment := sdklogging.LogLevel(ctx.GlobalString(EnvironmentFlag.Name))
	logger, err := sdklogging.NewZapLogger(Environment)
	if err != nil {
		return nil, err
	}

	EthRpcUrl := ctx.GlobalString(EthRpcUrlFlag.Name)
	if EthRpcUrl == "" {
		logger.Errorf("EthRpcUrl is required")
	}

	EthWsUrl := ctx.GlobalString(EthWsUrlFlag.Name)
	if EthWsUrl == "" {
		logger.Errorf("EthWsUrl is required")
	}

	AggregatorServerIpPortAddr := ctx.GlobalString(AggregatorServerIpPortAddressFlag.Name)
	if AggregatorServerIpPortAddr == "" {
		logger.Errorf("AggregatorServerIpPortAddress is required")
	}

	RegisterOperatorOnStartup := true // default value

	var credibleSquaringDeploymentRaw IncredibleSquaringDeploymentRaw
	credibleSquaringDeploymentFilePath := ctx.GlobalString(CredibleSquaringDeploymentFileFlag.Name)
	if _, err := os.Stat(credibleSquaringDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + credibleSquaringDeploymentFilePath + " does not exist")
	}
	sdkutils.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)

	ethRpcClient, err := eth.NewClient(EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(EthWsUrl)
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
		EcdsaPrivateKey:            ecdsaPrivateKey,
		Logger:                     logger,
		EthWsRpcUrl:                EthWsUrl,
		EthHttpRpcUrl:              EthRpcUrl,
		EthHttpClient:              ethRpcClient,
		EthWsClient:                ethWsClient,
		OperatorStateRetrieverAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		IncredibleSquaringRegistryCoordinatorAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr:                AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:                 RegisterOperatorOnStartup,
		SignerFn:                                  signerV2,
		TxMgr:                                     txMgr,
		AggregatorAddress:                         aggregatorAddr,
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
	/* Optional Flags */
	EnvironmentFlag = cli.StringFlag{
		Name:     "environment",
		Required: false,
		Usage:    "Set the environment (production or development)",
		Value:    "development", // default value
	}
	EthRpcUrlFlag = cli.StringFlag{
		Name:     "eth-rpc-url",
		Required: false,
		Usage:    "Ethereum RPC URL",
		EnvVar:   "ETH_RPC_URL",
	}
	EthWsUrlFlag = cli.StringFlag{
		Name:     "eth-ws-url",
		Required: false,
		Usage:    "Ethereum WS URL",
		EnvVar:   "ETH_WS_URL",
	}
	AggregatorServerIpPortAddressFlag = cli.StringFlag{
		Name:     "aggregator-server-ip-port-address",
		Required: false,
		Usage:    "Aggregator server IP PORT address",
		EnvVar:   "AGGREGATOR_SERVER_IP_PORT_ADDRESS",
	}
)

var requiredFlags = []cli.Flag{
	CredibleSquaringDeploymentFileFlag,
	EcdsaPrivateKeyFlag,
}

var optionalFlags = []cli.Flag{
	EnvironmentFlag,
	EthRpcUrlFlag,
	EthWsUrlFlag,
	AggregatorServerIpPortAddressFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
