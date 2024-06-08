// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./IIncredibleSquaringTaskManager.sol";

contract IncredibleSquaringTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    IIncredibleSquaringTaskManager
{
    using BN254 for BN254.G1Point;

    /* CONSTANT */
    // The number of blocks from the task initialization within which the aggregator has to respond to
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    string public ETH_USDC_FEED_NAME = "eth/usdc";
    string public BTC_USDC_FEED_NAME = "btc/usdc";

    /* STORAGE */
    // The latest task index
    uint32 public latestTaskNum;

    // mapping of task indices to all tasks hashes
    // when a task is created, task hash is stored here,
    // and responses need to pass the actual task,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allTaskHashes;

    // mapping of task indices to hash of abi.encode([] taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32[]) public allTaskResponses;

    mapping(uint32 => bool) public taskSuccesfullyChallenged;

    mapping(string => AggregatedPrice) private aggregatedPriceFeed;

    modifier onlyOperator() {
        IRegistryCoordinator.OperatorInfo memory operatorInfo = registryCoordinator.getOperator(msg.sender);
        require(operatorInfo.operatorId != bytes32(0));
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator,
        uint32 _taskResponseWindowBlock
    ) BLSSignatureChecker(_registryCoordinator) {
        TASK_RESPONSE_WINDOW_BLOCK = _taskResponseWindowBlock;
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator,
        address _generator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
    }

    /* FUNCTIONS */
    // NOTE: this function creates new task, assigns it a taskId
    // Add onlyOperator modifier
    function createNewTask(
        uint256 numberToBeSquared,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external {
        // create a new task struct
        Task memory newTask;
        newTask.numberToBeSquared = numberToBeSquared;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }

    // Add onlyOperator modifier
    function requestPriceFeedUpdate(
        string memory feedName,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers,
        uint8 minNumOfOracleNetworks
    ) external onlyOperator {
        PriceUpdateTask memory task;
        task.taskCreatedBlock = uint32(block.number);
        task.feedName = feedName;
        task.quorumNumbers = quorumNumbers;
        task.quorumThresholdPercentage = quorumThresholdPercentage;
        task.minNumOfOracleNetworks = minNumOfOracleNetworks;

        // store hash of task onchain, emit event, and increase taskNum
        allTaskHashes[latestTaskNum] = keccak256(abi.encode(task));
        emit PriceUpdateRequested(latestTaskNum, task);
        latestTaskNum = latestTaskNum + 1;
    }

    // NOTE: this function responds to existing tasks.
    // Add onlyOperator modifier
    function respondToTask(
        PriceUpdateTask calldata task,
        PriceUpdateTaskResponse[] calldata taskResponses, // Each price feed source has a different response
        NonSignerStakesAndSignature[] memory nonSignerStakesAndSignatures
    ) external onlyOperator {
        uint32 taskCreatedBlock = task.taskCreatedBlock;
        bytes calldata quorumNumbers = task.quorumNumbers;
        uint32 quorumThresholdPercentage = task.quorumThresholdPercentage;

        // some logical checks
            require(
                taskResponses.length > 0,
                "Aggregator must submit task response"
            );

            require(
                taskResponses.length == nonSignerStakesAndSignatures.length,
                "response length must match aggregate signature length"
            );

            require(
                allTaskResponses[taskResponses[0].taskId].length == 0,
                "Aggregator has already responded to the task"
            );

            require(
                uint32(block.number) <=
                    taskCreatedBlock + TASK_RESPONSE_WINDOW_BLOCK,
                "Aggregator has responded to the task too late"
            );

        // Iterate over each source submitted
        for (uint i = 0; i < taskResponses.length; i++) {
            // check that the task is valid, hasn't been responsed yet, and is being responsed in time
            // require(
            //     keccak256(abi.encode(task)) ==
            //         allTaskHashes[taskResponses[i].taskId],
            //     "supplied task does not match the one recorded in the contract"
            // );

            // bytes32 message = keccak256(abi.encode(taskResponses[i]));

            // // check the BLS signature matches msg hash (ie. {price, feedSource, taskId})
            // (
            //     QuorumStakeTotals memory quorumStakeTotals,
            //     bytes32 hashOfNonSigners
            // ) = checkSignatures(
            //         message,
            //         quorumNumbers,
            //         taskCreatedBlock,
            //         nonSignerStakesAndSignatures[i]
            //     );
            
            // // check that signatories own at least a threshold percentage of each quourm
            // for (uint j = 0; j < quorumNumbers.length; j++) {
            //     // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying
            //     // signed stake > total stake
            //     require(
            //         quorumStakeTotals.signedStakeForQuorum[j] *
            //             _THRESHOLD_DENOMINATOR >=
            //             quorumStakeTotals.totalStakeForQuorum[j] *
            //                 uint8(quorumThresholdPercentage),
            //         "Signatories do not own at least threshold percentage of a quorum"
            //     );
            // }

            TaskResponseMetadata memory taskResponseMetadata = TaskResponseMetadata(
                uint32(block.number),
                bytes32(0) // hashOfNonSigners
            );

            // updating the storage with task responses
            allTaskResponses[taskResponses[i].taskId].push(keccak256(
                abi.encode(taskResponses[i], taskResponseMetadata)
            ));

            // emitting event
            emit TaskResponded(taskResponses[i].taskId, taskResponses[i], taskResponseMetadata);
        }

        // Update aggregated price of feed
        aggregatedPriceFeed[task.feedName] = AggregatedPrice({price: taskResponses[0].price, decimals: uint8(taskResponses[0].decimals), lastBlockUpdated: uint32(block.number), lastUpdatedTaskId: taskResponses[0].price});
    }

    function fetchLatestAggregatedPrice(string calldata feedName) external view returns(AggregatedPrice memory) {
        require(aggregatedPriceFeed[feedName].lastBlockUpdated != 0, "Requested feed is not supported");

        AggregatedPrice memory data = aggregatedPriceFeed[feedName];
        require(uint32(block.number) - data.lastBlockUpdated < 100, "Requested feed has not been updated in 100 blocks");

        return data;
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    // NOTE: this function enables a challenger to raise and resolve a challenge.
    // TODO: require challenger to pay a bond for raising a challenge
    // TODO(samlaf): should we check that quorumNumbers is same as the one recorded in the task?
    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        // TODO
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }
}
