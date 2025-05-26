package operator

// TODO: add toml flags to the SDK operator config, removing most of this config attributes
type Config struct {
	OperatorAddress string `toml:"operator_address"`

	// Core deployment addresses
	AllocationManagerAddress    string `toml:"allocation_manager_address"`
	DelegationManagerAddress    string `toml:"delegation_manager_address"`
	RewardsCoordinatorAddress   string `toml:"rewards_coordinator_address"`
	PermissionControllerAddress string `toml:"permission_controller_address"`

	// Avs deployment addresses
	ServiceManagerAddress      string `toml:"service_manager_address"`
	RegistryCoordinatorAddress string `toml:"registry_coordinator_address"`
	TokenStrategyAddr          string `toml:"token_strategy_addr"`

	EcdsaPrivateKeyStorePath string `toml:"ecdsa_private_key_store_path"`
	BlsPrivateKeyStorePath   string `toml:"bls_private_key_store_path"`

	EthRpcUrl string `toml:"eth_http_url"`
	EthWsUrl  string `toml:"eth_ws_url"`

	AggregatorServerIpPortAddress string `toml:"aggregator_server_ip_port"`

	TaskManagerAddress string `toml:"task_manager_address"`
}
