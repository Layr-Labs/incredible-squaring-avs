package aggregator

import "github.com/Layr-Labs/eigensdk-go/aggregator"

// This config has the same attributes as the aggregator config and also includes the
// deployed TaskManager contract address
type Config struct {
	aggregator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}
