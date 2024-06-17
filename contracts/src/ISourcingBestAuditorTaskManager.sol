// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer-middleware/src/libraries/BN254.sol";

interface ISourcingBestAuditorTaskManager {
    // EVENTS
    event NewTaskCreated(uint32 indexed taskIndex, Task task);

    event TaskResponded(
        TaskResponse taskResponse,
        TaskResponseMetadata taskResponseMetadata
    );

    event TaskCompleted(uint32 indexed taskIndex);

    event TaskCreated(
        uint256 taskId,
        uint256 budget,
        bytes auditJobSpecificationURI,
        bytes quorumNumbers,
        uint32 quorumThresholdPercentage
    );

    // STRUCTS
    struct Task {
        uint256 taskId;
        bytes auditJobSpecificationURI;
        uint256 budget;
        bytes quorumNumbers;
        uint32 quorumThresholdPercentage;
        Bid[] bids;
    }

    struct Bid {
        uint256 bidId;
        string zkp;
        string bidPitchDocURI;
    }

    struct TaskResponse {
        uint32 referenceTaskIndex; // Can be obtained by the operator from the event NewTaskCreated.
        string evaluatedBidURI; // URI pointing to the operator's evaluation of the bids.
        uint256 selectedBidId; // The ID of the selected best bid.
    }

    struct TaskResponseMetadata {
        uint32 taskResponsedBlock;
        bytes32 hashOfNonSigners;
    }

    // FUNCTIONS
    // NOTE: this function creates new task.
    function createNewTask(
        bytes calldata auditJobSpecificationURI,
        uint256 budgetInUSDC,
        bytes calldata quorumNumbers,
        uint32 quorumThresholdPercentage,
        Bid[] calldata bids // Submitting the array of all the bids
    ) external;

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view returns (uint32);

    /// @notice Returns the TASK_RESPONSE_WINDOW_BLOCK
    function getTaskResponseWindowBlock() external view returns (uint32);
}
