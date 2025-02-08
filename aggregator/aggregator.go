package aggregator

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	
	"github.com/ehsueh/trading-algo-avs/aggregator/types"
	"github.com/ehsueh/trading-algo-avs/core/chainio"
	"github.com/ehsueh/trading-algo-avs/core/config"
	"github.com/ehsueh/trading-algo-avs/core/utils"

	tradingTaskManager "github.com/kaohaohan/TradeAlgoAVS-frontend/blob/b7dd2c8fb10844d6c565ce7c06ddddea662a48f4/blockchain/contracts/TradingAlgoAVS"
)

const (
	taskChallengeWindowBlock = 100 // Blocks before a task expires // for future use
	blockTimeSeconds         = 12 * time.Second
	avsName                  = "trading-algo-avs"
	quorumThreshold          = 66 // Percentage of quorum required for validation
	smartContractAddress      = "0x1D3A33D45557bb72F05DBdeB6a50c6Ea9818183a"
)

// Aggregator is responsible for:
// - Listening to smart contract events for subscriptions
// - Generating daily tasks for active subscriptions
// - Aggregating results and validating against the quorum
// - Rewarding honest execution and penalizing incorrect execution

type Aggregator struct {
	logger          logging.Logger
	avsWriter       chainio.AvsWriterer
	subscriptions   map[uint256]map[common.Address]bool // Tracks active subscriptions
	results        map[uint256][]OperatorResult        // Tracks operator execution results

	subscriptionsMu sync.RWMutex
	resultsMu       sync.RWMutex
}

type OperatorResult struct {
	Roi           *big.Int
	Profitability *big.Int
	Risk          *big.Int
}

// Start the aggregator to listen for new strategy and subscription events
func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Info("Starting aggregator...")
	go agg.listenForSubscriptionEvents(ctx)
	// go agg.listenForStrategyEvents(ctx) // Future: Listen for strategy update events
	return nil
}

// Listen for subscription-related events from the blockchain
func (agg *Aggregator) listenForSubscriptionEvents(ctx context.Context) {
	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(smartContractAddress),
		},
	}

	sub, err := agg.rpcClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		agg.logger.Error("Failed to subscribe to contract logs", "error", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		case log := <-logs:
			agg.handleContractEvent(log)
		}
	}
}
// Listen for subscription-related events from the blockchain
func (agg *Aggregator) listenForSubscriptionEvents(ctx context.Context) {
	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(smartContractAddress),
		},
	}

	sub, err := agg.rpcClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		agg.logger.Error("Failed to subscribe to contract logs", "error", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		case log := <-logs:
			agg.handleContractEvent(log)
		}
	}
}

// Handle contract events and route them to appropriate handlers
func (agg *Aggregator) handleContractEvent(log types.Log) {
	switch log.Topics[0].Hex() {
	case "0xSubscription": // Subscription event
		agg.handleNewSubscription(log)
	case "0xSubscriptionCancelled": // Subscription cancelled event
		agg.handleSubscriptionDeletion(log)
	default:
		agg.logger.Warn("Unhandled event", "event", log.Topics[0].Hex())
	}
}

// Handle new subscription creation
func (agg *Aggregator) handleNewSubscription(log types.Log) {
	strategyId := new(big.Int).SetBytes(log.Topics[1].Bytes())
	subscriber := common.BytesToAddress(log.Topics[2].Bytes())
	
	agg.subscriptionsMu.Lock()
	if _, exists := agg.subscriptions[strategyId.Uint64()]; !exists {
		agg.subscriptions[strategyId.Uint64()] = make(map[common.Address]bool)
	}
	agg.subscriptions[strategyId.Uint64()][subscriber] = true
	agg.subscriptionsMu.Unlock()
	agg.logger.Info("New subscription added", "strategyId", strategyId, "subscriber", subscriber)
}

// Handle subscription deletion
func (agg *Aggregator) handleSubscriptionDeletion(log types.Log) {
	subscriptionId := new(big.Int).SetBytes(log.Topics[1].Bytes())
	
	agg.subscriptionsMu.Lock()
	for strategyId, subs := range agg.subscriptions {
		delete(subs, subscriptionId)
	}
	agg.subscriptionsMu.Unlock()
	agg.logger.Info("Subscription removed", "subscriptionId", subscriptionId)
}

// GenerateDailyTasks creates daily tasks for active subscriptions
// Each active subscription corresponds to a task that will be executed by an operator
func (agg *Aggregator) GenerateDailyTasks() {
	agg.logger.Info("Generating daily tasks")
	agg.subscriptionsMu.RLock()
	for strategyId, subscribers := range agg.subscriptions {
		for subscriber := range subscribers {
			agg.logger.Info("Creating task for subscriber", "strategyId", strategyId, "subscriber", subscriber)
			
		}
	}
	agg.subscriptionsMu.RUnlock()
}

// ApplyRewardsAndPenalties rewards accurate operators and penalizes deviators based on deviation threshold
// Any operator whose result deviates more than 1 standard deviation from the mean is penalized
func (agg *Aggregator) ApplyRewardsAndPenalties() {
	agg.logger.Info("Applying rewards and penalties")
	agg.resultsMu.RLock()
	// TODO: Implement reward and penalty application logic
	agg.resultsMu.RUnlock()
}

// sendNewTask generates and submits a new task for an active subscription.
func (agg *Aggregator) sendNewTask(strategyId uint256, investor common.Address) error {
	agg.logger.Info("Generating new task", "strategyId", strategyId, "investor", investor)

	// Define task parameters (strategy, investor)
	task := types.Task{
		StrategyId: strategyId,
		Investor:   investor,
		Timestamp:  time.Now().Unix(),
	}

	// Lock subscriptions before accessing
	agg.subscriptionsMu.RLock()
	_, exists := agg.subscriptions[strategyId][investor]
	agg.subscriptionsMu.RUnlock()

	if !exists {
		agg.logger.Warn("Task generation failed - subscription not found", "strategyId", strategyId, "investor", investor)
		return errors.New("subscription not found")
	}

	// Assign a single operator (in a real-world scenario, an operator selection mechanism would be used)
	operator := selectOperatorForTask(strategyId)

	agg.logger.Info("Assigned operator for task", "operator", operator, "task", task)

	// Store the task details
	agg.resultsMu.Lock()
	if _, exists := agg.results[task.StrategyId]; !exists {
		agg.results[task.StrategyId] = []OperatorResult{}
	}
	agg.resultsMu.Unlock()

	// Send the task for execution (in practice, this could involve an RPC call or message queue)
	err := chainio.SubmitTaskToOperator(operator, task)
	if err != nil {
		agg.logger.Error("Failed to send task to operator", "error", err)
		return err
	}

	agg.logger.Info("Task successfully sent", "task", task)
	return nil
}

// selectOperatorForTask selects an operator for a given strategy task.
func selectOperatorForTask(strategyId uint256) common.Address {
	// Placeholder logic: Select an operator (this can be enhanced with a selection mechanism)
	// Here, a simple round-robin or load-balancing mechanism can be implemented.
	return common.HexToAddress("0xOperatorAddressPlaceholder")
}

// sendAggregatedResponseToContract sends the validated and aggregated task results on-chain with BLS signatures
func (agg *Aggregator) sendAggregatedResponseToContract(taskId uint256) error {
	agg.resultsMu.RLock()
	results, exists := agg.results[taskId]
	agg.resultsMu.RUnlock()

	if !exists || len(results) == 0 {
		agg.logger.Warn("No results available to aggregate", "taskId", taskId)
		return errors.New("no results available")
	}

	// Validate and aggregate results using standard deviation check
	validResults, meanRoi, meanProfitability, meanRisk, err := agg.ValidateResults(taskId, results)
	if err != nil {
		agg.logger.Warn("Failed to validate results", "taskId", taskId, "error", err)
		return err
	}

	// Retrieve operator states to identify signers and non-signers
	operatorState, err := chainio.GetOperatorState(taskId)
	if err != nil {
		agg.logger.Error("Failed to fetch operator state", "taskId", taskId, "error", err)
		return err
	}

	// Aggregate BLS signatures for quorum validation
	quorumApks := []BN254G1Point{}
	nonSignerPubkeys := []BN254G1Point{}

	for _, operator := range operatorState.Operators {
		if operator.Signed {
			quorumApks = append(quorumApks, operator.Pubkey)
		} else {
			nonSignerPubkeys = append(nonSignerPubkeys, operator.Pubkey)
		}
	}

	// Check if quorum threshold is met
	if len(quorumApks)*100/len(operatorState.Operators) < quorumThreshold {
		agg.logger.Warn("Quorum not met, not submitting response", "taskId", taskId)
		return errors.New("quorum not met")
	}

	// Use the same BLSSignatureChecker.sol contract as in the EigenLayer example, which expects a
	//
	//	struct NonSignerStakesAndSignature {
	//		uint32[] nonSignerQuorumBitmapIndices;
	//		BN254.G1Point[] nonSignerPubkeys;
	//		BN254.G1Point[] quorumApks;
	//		BN254.G2Point apkG2;
	//		BN254.G1Point sigma;
	//		uint32[] quorumApkIndices;
	//		uint32[] totalStakeIndices;
	//		uint32[][] nonSignerStakeIndices; // nonSignerStakeIndices[quorumNumberIndex][nonSignerIndex]
	//	}
	// Construct BLS signature structure for contract verification
	blsSignature := cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        core.ConvertToBN254G2Point(operatorState.SignersApkG2),
		Sigma:                        core.ConvertToBN254G1Point(operatorState.SignersAggSigG1),
		NonSignerQuorumBitmapIndices: operatorState.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             operatorState.QuorumApkIndices,
		TotalStakeIndices:            operatorState.TotalStakeIndices,
		NonSignerStakeIndices:        operatorState.NonSignerStakeIndices,
	}

	// Send aggregated result with signature verification
	agg.logger.Info("Quorum met, submitting result on-chain", "taskId", taskId,
		"meanRoi", meanRoi, "meanProfitability", meanProfitability, "meanRisk", meanRisk)

	err = chainio.SubmitAggregatedResultWithSignature(taskId, meanRoi, meanProfitability, meanRisk, blsSignature)
	if err != nil {
		agg.logger.Error("Failed to submit aggregated result with signature", "taskId", taskId, "error", err)
		return err
	}

	agg.logger.Info("Aggregated response with signature successfully submitted", "taskId", taskId)
	return nil
}


// ValidateResults checks consistency across operator results using standard deviation
func (agg *Aggregator) ValidateResults(taskId uint256, results []OperatorResult) ([]OperatorResult, *big.Int, *big.Int, *big.Int, error) {
	agg.logger.Info("Validating operator results", "taskId", taskId)

	var totalRoi, totalProfitability, totalRisk big.Int
	count := big.NewInt(int64(len(results)))

	// Compute sum for mean calculation
	for _, result := range results {
		totalRoi.Add(&totalRoi, result.Roi)
		totalProfitability.Add(&totalProfitability, result.Profitability)
		totalRisk.Add(&totalRisk, result.Risk)
	}

	meanRoi := new(big.Int).Div(&totalRoi, count)
	meanProfitability := new(big.Int).Div(&totalProfitability, count)
	meanRisk := new(big.Int).Div(&totalRisk, count)

	// Compute variance for standard deviation
	var varianceRoi, varianceProfitability, varianceRisk big.Int
	for _, result := range results {
		diffRoi := new(big.Int).Sub(result.Roi, meanRoi)
		diffProfitability := new(big.Int).Sub(result.Profitability, meanProfitability)
		diffRisk := new(big.Int).Sub(result.Risk, meanRisk)

		varianceRoi.Add(&varianceRoi, new(big.Int).Mul(diffRoi, diffRoi))
		varianceProfitability.Add(&varianceProfitability, new(big.Int).Mul(diffProfitability, diffProfitability))
		varianceRisk.Add(&varianceRisk, new(big.Int).Mul(diffRisk, diffRisk))
	}

	stdDevRoi := new(big.Int).Sqrt(new(big.Int).Div(&varianceRoi, count))
	stdDevProfitability := new(big.Int).Sqrt(new(big.Int).Div(&varianceProfitability, count))
	stdDevRisk := new(big.Int).Sqrt(new(big.Int).Div(&varianceRisk, count))

	// Filter valid results within one standard deviation
	validResults := []OperatorResult{}
	for _, result := range results {
		if isWithinOneStdDev(result.Roi, meanRoi, stdDevRoi) &&
			isWithinOneStdDev(result.Profitability, meanProfitability, stdDevProfitability) &&
			isWithinOneStdDev(result.Risk, meanRisk, stdDevRisk) {
			validResults = append(validResults, result)
		}
	}

	if len(validResults) == 0 {
		return nil, nil, nil, nil, errors.New("no valid results within standard deviation")
	}

	return validResults, meanRoi, meanProfitability, meanRisk, nil
}

// isWithinOneStdDev checks if a value is within one standard deviation of the mean
func isWithinOneStdDev(value, mean, stdDev *big.Int) bool {
	lowerBound := new(big.Int).Sub(mean, stdDev)
	upperBound := new(big.Int).Add(mean, stdDev)
	return value.Cmp(lowerBound) >= 0 && value.Cmp(upperBound) <= 0
}
