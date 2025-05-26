package challenger

type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	EthHttpUrl string `toml:"eth_http_url"`
	EthWsUrl   string `toml:"eth_ws_url"`
}
