package taskspammer

import taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"

type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	EthHttpUrl string `toml:"eth_http_url"`

	taskspammer.Config
}
