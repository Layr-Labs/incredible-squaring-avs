package aggregator

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/trading-algo-avs/aggregator/types"
	"github.com/Layr-Labs/trading-algo-avs/core"
	"github.com/Layr-Labs/trading-algo-avs/core/chainio"
	"github.com/Layr-Labs/trading-algo-avs/core/config"

	tradingTaskManager "github.com/kaohaohan/TradeAlgoAVS-frontend/blob/b7dd2c8fb10844d6c565ce7c06ddddea662a48f4/blockchain/contracts/TradingAlgoAVS"
)

const (
	// Number of blocks after which a task is considered expired
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "trading-algo-avs"
)

// Aggregator listens to trading strategy events from the TradingAlgoAVS smart contract,
// collects operator outputs for (strategy, user, market) tasks, aggregates statistics,
// and sends the verified response on-chain if operator responses are consistent.
type Aggregator struct {
	logger                logging.Logger
	serverIpPortAddr      string
	avsWriter             chainio.AvsWriterer
	blsAggregationService blsagg.BlsAggregationService
	strategies            map[uint256]tradingTaskManager.Strategy
	subscriptions         map[uint256]map[common.Address]bool
	strategiesMu          sync.RWMutex
	subscriptionsMu       sync.RWMutex
}

// NewAggregator creates a new Aggregator for TradingAlgoAVS.
func NewAggregator(c *config.Config) (*Aggregator, error) {
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	clients, err := clients.BuildAll(clients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.TradingAlgoRegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
	}, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create SDK clients", "err", err)
		return nil, err
	}

	blsAggregationService := blsagg.NewBlsAggregatorService(nil, nil, c.Logger)

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		strategies:            make(map[uint256]tradingTaskManager.Strategy),
		subscriptions:         make(map[uint256]map[common.Address]bool),
	}, nil
}

// Start the aggregator to listen for new strategy and subscription events
func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Infof("Starting Trading Algo AVS Aggregator.")

	contractAddress := common.HexToAddress("0xYourContractAddress")
	tradingContract, err := tradingTaskManager.NewTradingAlgoAVS(contractAddress, nil) // Replace `nil` with actual blockchain client
	if err != nil {
		agg.logger.Error("Failed to load TradingAlgoAVS contract", "err", err)
		return err
	}

	// Subscribe to StrategyCreated event
	strategyCreatedChan := make(chan *tradingTaskManager.TradingAlgoAVSStrategyCreated)
	strategySub, err := tradingContract.WatchStrategyCreated(nil, strategyCreatedChan, nil, nil)
	if err != nil {
		agg.logger.Error("Failed to subscribe to StrategyCreated event", "err", err)
		return err
	}

	// Subscribe to SubscribedToStrategy event
	subscribedChan := make(chan *tradingTaskManager.TradingAlgoAVSSubscribedToStrategy)
	subscriptionSub, err := tradingContract.WatchSubscribedToStrategy(nil, subscribedChan, nil, nil)
	if err != nil {
		agg.logger.Error("Failed to subscribe to SubscribedToStrategy event", "err", err)
		return err
	}

	// Process events in a goroutine
	go func() {
		for {
			select {
			case strategyEvent := <-strategyCreatedChan:
				agg.processNewStrategy(strategyEvent)
			case subscriptionEvent := <-subscribedChan:
				agg.processNewSubscription(subscriptionEvent)
			case <-ctx.Done():
				strategySub.Unsubscribe()
				subscriptionSub.Unsubscribe()
				return
			}
		}
	}()

	return nil
}

// Process new strategies when detected
func (agg *Aggregator) processNewStrategy(event *tradingTaskManager.TradingAlgoAVSStrategyCreated) {
	agg.logger.Info("New strategy detected", "strategyId", event.StrategyId, "provider", event.Provider)

	agg.strategiesMu.Lock()
	agg.strategies[event.StrategyId] = tradingTaskManager.Strategy{
		Id:               event.StrategyId,
		Provider:         event.Provider,
		SubscriptionFee:  event.SubscriptionFee,
		SubscriptionPeriod: event.SubscriptionPeriod,
		StrategyUid:      event.StrategyUid,
		Roi:              event.Roi,
		Profitability:    event.Profitability,
		Risk:             event.Risk,
	}
	agg.strategiesMu.Unlock()
}

// Process new subscriptions when detected
func (agg *Aggregator) processNewSubscription(event *tradingTaskManager.TradingAlgoAVSSubscribedToStrategy) {
	agg.logger.Info("New subscription detected", "strategyId", event.StrategyId, "subscriber", event.Subscriber)

	agg.subscriptionsMu.Lock()
	if _, exists := agg.subscriptions[event.StrategyId]; !exists {
		agg.subscriptions[event.StrategyId] = make(map[common.Address]bool)
	}
	agg.subscriptions[event.StrategyId][event.Subscriber] = true
	agg.subscriptionsMu.Unlock()
}

// Aggregates operator statistics and sends them on-chain if consistent
func (agg *Aggregator) aggregateResults(strategyId uint256) {
	agg.strategiesMu.RLock()
	strategy, exists := agg.strategies[strategyId]
	agg.strategiesMu.RUnlock()

	if !exists {
		agg.logger.Error("Strategy not found", "strategyId", strategyId)
		return
	}

	// Fetch operator results (Mocked for now)
	operatorResults := []struct {
		Roi          *big.Int
		Profitability *big.Int
		Risk         *big.Int
	}{
		{big.NewInt(5), big.NewInt(80), big.NewInt(10)},
		{big.NewInt(5), big.NewInt(79), big.NewInt(11)},
		{big.NewInt(5), big.NewInt(81), big.NewInt(9)},
	}

	// Calculate mean and standard deviation
	// (Add standard deviation check logic here)

	// If within standard deviation, submit to blockchain
	agg.logger.Info("Sending aggregated trading statistics on-chain.", "strategyId", strategyId)
	_, err := agg.avsWriter.SendAggregatedTradingResults(context.Background(), strategy)
	if err != nil {
		agg.logger.Error("Failed to send aggregated results", "err", err)
	}
}
