package types

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                       bool   `yaml:"production"`
	OperatorAddress                  string `yaml:"operator_address"`
	ELSlasherAddress                 string `yaml:"el_slasher_address"`
	BLSOperatorStateRetrieverAddress string `yaml:"bls_operator_state_retriever_address"`
	BlsPublicKeyCompendiumAddress    string `yaml:"bls_public_key_compendium_address"`
	AVSServiceManagerAddress         string `yaml:"avs_service_manager_address"`
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
}
