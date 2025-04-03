package taskgenerator

import (
	"context"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/incredible-squaring-avs/aggregator/types"
	csservicemanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringServiceManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/chainio"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type TaskGenerator struct {
	taskManagerAddr common.Address
	avsWriter       chainio.AvsWriterer
	logger          logging.Logger
}

func BuildTaskGenerator(c *config.Config) (*TaskGenerator, error) {
	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	contractServiceManager, err := csservicemanager.NewContractIncredibleSquaringServiceManager(
		c.IncredibleSquaringServiceManager,
		&c.EthHttpClient,
	)
	if err != nil {
		c.Logger.Error("Failed to fetch IServiceManager contract", "err", err)
		return nil, err
	}
	taskManagerAddr, err := contractServiceManager.IncredibleSquaringTaskManager(&bind.CallOpts{})
	if err != nil {
		c.Logger.Error("Failed to fetch TaskManager address", "err", err)
		return nil, err
	}

	return &TaskGenerator{
		taskManagerAddr,
		avsWriter,
		c.Logger,
	}, nil
}

func (taskGen *TaskGenerator) Start(ctx context.Context) error {
	time.Sleep(time.Duration(2 * time.Second))

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
			taskGen.logger.Infof("Task Generator sending new task, number to square: %v", taskNum)
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
