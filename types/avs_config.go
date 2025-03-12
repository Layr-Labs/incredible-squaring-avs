package types

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                       bool   `yaml:"production"`
	OperatorAddress                  string `yaml:"operator_address"`
	OperatorStateRetrieverAddress    string `yaml:"operator_state_retriever_address"`
	IncredibleSquaringServiceManager string `yaml:"service_manager_address"`
	InstantSlasher                   string `yaml:"instant_slasher_address"`
	AVSRegistryCoordinatorAddress    string `yaml:"avs_registry_coordinator_address"`
	RewardsCoordinatorAddress        string `yaml:"rewards_coordinator_address"`
	PermissionControllerAddress      string `yaml:"permission_controller_address"`
	AllocationManagerAddress         string `yaml:"allocation_manager_address"`
	TokenStrategyAddr                string `yaml:"token_strategy_addr"`
	EthRpcUrl                        string `yaml:"eth_rpc_url"`
	EthWsUrl                         string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath           string `yaml:"bls_private_key_store_path"`
	EcdsaPrivateKeyStorePath         string `yaml:"ecdsa_private_key_store_path"`
	AggregatorServerIpPortAddress    string `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup        bool   `yaml:"register_operator_on_startup"`
	EigenMetricsIpPortAddress        string `yaml:"eigen_metrics_ip_port_address"`
	EnableMetrics                    bool   `yaml:"enable_metrics"`
	NodeApiIpPortAddress             string `yaml:"node_api_ip_port_address"`
	EnableNodeApi                    bool   `yaml:"enable_node_api"`
	OperatorSetId                    uint32 `yaml:"operator_set_id"`
	Socket                           string `yaml:"socket"`
	MaxOperatorCount                 uint32 `yaml:"max_operator_count"`
	KickBIPsOfOperatorStake          uint16 `yaml:"kick_bips_of_operator_stake"`
	KickBIPsOfTotalStake             uint16 `yaml:"kick_bips_of_total_stake"`
	MinimumStake                     int64  `yaml:"minimum_stake"`
	Multiplier                       int64  `yaml:"multiplier"`
	TimesFailing                     int    `yaml:"times_failing"`
}
