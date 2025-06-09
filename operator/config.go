package operator

import "github.com/Layr-Labs/eigensdk-go/operator"

type Config struct {
	operator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}
