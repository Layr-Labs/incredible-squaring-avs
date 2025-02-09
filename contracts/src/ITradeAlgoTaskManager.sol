// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";
import {IBLSSignatureChecker} from "@eigenlayer-middleware/src/interfaces/IBLSSignatureChecker.sol";

interface ITradeAlgoTaskManager is IBLSSignatureChecker {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(TaskResponse taskResponse, TaskResponseMetadata taskResponseMetadata);

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger);

    event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger);

    event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator);

    event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator);

    // STRUCTS
    struct Task {
        uint256 strategyId; // Identifier for the trading strategy
        address investor; // Investor associated with the task
        uint32 taskCreatedBlock;
        bytes quorumNumbers; // Required quorums for validation
        uint32 quorumThresholdPercentage; // Quorum threshold for task validation
    }

    struct TaskResponse {
        uint32 referenceTaskIndex; // Task being responded to
        uint256 calculatedResult; // The computed output based on strategy execution
        bytes operatorSignature; // BLS signature from the operator
    }

    struct TaskResponseMetadata {
        uint32 taskRespondedBlock;
        bytes32 hashOfNonSigners;
    }
    
    // FUNCTIONS

    /// @notice Creates a new task for an investor following a subscribed strategy
    function createNewTask(
        uint256 strategyId,
        address investor,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external;

    /// @notice Retrieves the latest task index
    function taskNumber() external view returns (uint32);

    /// @notice Allows operators to respond with computed strategy results
    function submitTaskResponse(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        uint32 quorumThresholdPercentage,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature    
    ) external;

    /// @notice Challenges a task response for potential inaccuracies or fraud
    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external;

    /// @notice Retrieves the challenge response window duration
    function getTaskResponseWindowBlock() external view returns (uint32);
}
