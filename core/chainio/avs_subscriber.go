package chainio

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"

	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/CredibleSquaringTaskManager"
	"github.com/Layr-Labs/incredible-squaring-avs/core/config"
)

type AvsSubscriberer interface {
	SubscribeToNewTasks(newTaskCreatedChan chan *cstaskmanager.ContractCredibleSquaringTaskManagerNewTaskCreated) event.Subscription
	SubscribeToTaskResponses(taskResponseLogs chan *cstaskmanager.ContractCredibleSquaringTaskManagerTaskResponded) event.Subscription
	ParseTaskResponded(rawLog types.Log) (*cstaskmanager.ContractCredibleSquaringTaskManagerTaskResponded, error)
}

// Subscribers use a ws connection instead of http connection like Readers
// kind of stupid that the geth client doesn't have a unified interface for both...
// it takes a single url, so the bindings, even though they have watcher functions, those can't be used
// with the http connection... seems very very stupid. Am I missing something?
type AvsSubscriber struct {
	AvsContractBindings *AvsServiceBindings
	logger              sdklogging.Logger
}

func NewAvsSubscriberFromConfig(config *config.Config) (*AvsSubscriber, error) {
	return NewAvsSubscriber(
		config.CredibleSquaringServiceManagerAddr,
		config.BlsOperatorStateRetrieverAddr,
		config.EthWsClient,
		config.Logger,
	)
}

func NewAvsSubscriber(serviceManagerAddr, blsOperatorStateRetrieverAddr gethcommon.Address, ethclient eth.EthClient, logger sdklogging.Logger) (*AvsSubscriber, error) {
	avsContractBindings, err := NewAvsServiceBindings(serviceManagerAddr, blsOperatorStateRetrieverAddr, ethclient, logger)
	if err != nil {
		logger.Errorf("Failed to create contract bindings", "err", err)
		return nil, err
	}
	return &AvsSubscriber{
		AvsContractBindings: avsContractBindings,
		logger:              logger,
	}, nil
}

func (s *AvsSubscriber) SubscribeToNewTasks(newTaskCreatedChan chan *cstaskmanager.ContractCredibleSquaringTaskManagerNewTaskCreated) event.Subscription {
	sub, err := s.AvsContractBindings.TaskManager.WatchNewTaskCreated(
		&bind.WatchOpts{}, newTaskCreatedChan, nil,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to new TaskManager tasks", "err", err)
	}
	s.logger.Infof("Subscribed to new TaskManager tasks")
	return sub
}

func (s *AvsSubscriber) SubscribeToTaskResponses(taskResponseChan chan *cstaskmanager.ContractCredibleSquaringTaskManagerTaskResponded) event.Subscription {
	sub, err := s.AvsContractBindings.TaskManager.WatchTaskResponded(
		&bind.WatchOpts{}, taskResponseChan,
	)
	if err != nil {
		s.logger.Error("Failed to subscribe to TaskResponded events", "err", err)
	}
	s.logger.Infof("Subscribed to TaskResponded events")
	return sub
}

func (s *AvsSubscriber) ParseTaskResponded(rawLog types.Log) (*cstaskmanager.ContractCredibleSquaringTaskManagerTaskResponded, error) {
	return s.AvsContractBindings.TaskManager.ContractCredibleSquaringTaskManagerFilterer.ParseTaskResponded(rawLog)
}
