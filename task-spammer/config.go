package taskspammer

type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	EthHttpUrl string `toml:"eth_http_url"`
}
