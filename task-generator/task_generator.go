package taskgenerator

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/common"
)



type TaskGenerator struct {
    taskManagerAddr common.Address
    rpcUrl string
	avsWriter        chainio.AvsWriterer
	logger           logging.Logger
}

func BuildTaskGenerator(c *config.Config, taskManagerAddr common.Address, rpcUrl string) (TaskGenerator, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return TaskGenerator{}, err
	}

    return TaskGenerator{
        taskManagerAddr,
        rpcUrl,
		avsWriter,
		c.Logger,
    }, nil
}


func (taskGen *TaskGenerator) Start(ctx context.Context) error {
	time.Sleep(time.Duration(10*time.Second)) // wait for 10 seconds first
	
	taskGen.logger.Info("Starting Task Generator.")
	taskGen.logger.Info("Starting Task Generator rpc server.")

	ticker := time.NewTicker(10 * time.Second)
	taskGen.logger.Info("Task Generator set to send new task every 10 seconds...")
	defer ticker.Stop()
	taskNum := int64(0)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			_, _, err := taskGen.avsWriter.SendNewTaskNumberToSquare(
				context.Background(),
				big.NewInt(taskNum),
				types.QUORUM_THRESHOLD_NUMERATOR,
				types.QUORUM_NUMBERS,
			)
			if err != nil {
				taskGen.logger.Error("Aggregator failed to send number to square", "err", err)
				return err
			}
			taskNum++
		}
	}
}

