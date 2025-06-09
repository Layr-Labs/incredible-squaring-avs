package challenger

import "github.com/Layr-Labs/eigensdk-go/challenger"

type Config struct {
	TaskManagerAddress string `toml:"task_manager_address"`

	challenger.Config
}
