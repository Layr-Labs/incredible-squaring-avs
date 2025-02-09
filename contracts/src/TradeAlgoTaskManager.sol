// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "@eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {
    BLSSignatureChecker,
    IRegistryCoordinator
} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@eigenlayer-middleware/src/libraries/BN254.sol";
import "./ITradeAlgoTaskManager.sol";

contract TradeAlgoTaskManager is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    BLSSignatureChecker,
    OperatorStateRetriever,
    ITradeAlgoTaskManager
{
    using BN254 for BN254.G1Point;

    /* CONSTANTS */
    uint32 public immutable TASK_RESPONSE_WINDOW_BLOCK;
    uint32 public constant TASK_CHALLENGE_WINDOW_BLOCK = 100;
    uint256 internal constant THRESHOLD_DENOMINATOR = 100;

    /* STORAGE */
    uint32 public latestTaskNum;

    mapping(uint32 => bytes32) public allTaskHashes;
    mapping(uint32 => bytes32) public allTaskResponses;
    mapping(uint32 => bool) public taskSuccessfullyChallenged;

    address public aggregator;
    address public generator;

    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    modifier onlyTaskGenerator() {
        require(msg.sender == generator, "Task generator must be the caller");
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
        _setAggregator(_aggregator);
        _setGenerator(_generator);
    }

    function setGenerator(address newGenerator) external onlyOwner {
        _setGenerator(newGenerator);
    }

    function setAggregator(address newAggregator) external onlyOwner {
        _setAggregator(newAggregator);
    }

    /* FUNCTIONS */

    function createNewTask(
        uint256 strategyId,
        address investor,
        uint32 quorumThresholdPercentage,
        bytes calldata quorumNumbers
    ) external onlyTaskGenerator {
        Task memory newTask;
        newTask.strategyId = strategyId;
        newTask.investor = investor;
        newTask.taskCreatedBlock = uint32(block.number);
        newTask.quorumThresholdPercentage = quorumThresholdPercentage;
        newTask.quorumNumbers = quorumNumbers;

        allTaskHashes[latestTaskNum] = keccak256(abi.encode(newTask));
        emit NewTaskCreated(latestTaskNum, newTask);
        latestTaskNum = latestTaskNum + 1;
    }

    function submitTaskResponse(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        uint32 quorumThresholdPercentage,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        uint32 taskIndex = taskResponse.referenceTaskIndex;
        require(
            keccak256(abi.encode(taskResponse)) == allTaskHashes[taskIndex],
            "Invalid task response"
        );
        require(
            allTaskResponses[taskIndex] == bytes32(0),
            "Task already responded"
        );
        require(
            uint32(block.number) <= taskResponseMetadata.taskRespondedBlock + TASK_RESPONSE_WINDOW_BLOCK,
            "Response too late"
        );

        bytes32 message = keccak256(abi.encode(taskResponse));

        (QuorumStakeTotals memory quorumStakeTotals, bytes32 hashOfNonSigners) =
            checkSignatures(message, task.quorumNumbers, uint32(taskResponseMetadata.taskRespondedBlock), nonSignerStakesAndSignature);

        for (uint256 i = 0; i < taskResponseMetadata.taskRespondedBlock; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * THRESHOLD_DENOMINATOR
                    >= quorumStakeTotals.totalStakeForQuorum[i] * uint8(quorumThresholdPercentage),
                "Signatories below threshold"
            );
        }

        allTaskResponses[taskIndex] = keccak256(abi.encode(taskResponse, taskResponseMetadata));

        emit TaskResponded(taskResponse, taskResponseMetadata);
    }

    function taskNumber() external view returns (uint32) {
        return latestTaskNum;
    }

    function raiseAndResolveChallenge(
        Task calldata task,
        TaskResponse calldata taskResponse,
        TaskResponseMetadata calldata taskResponseMetadata,
        BN254.G1Point[] memory pubkeysOfNonSigningOperators
    ) external {
        uint32 referenceTaskIndex = taskResponse.referenceTaskIndex;
        require(
            allTaskResponses[referenceTaskIndex] != bytes32(0),
            "Task not responded"
        );
        require(
            allTaskResponses[referenceTaskIndex] ==
                keccak256(abi.encode(taskResponse, taskResponseMetadata)),
            "Task response mismatch"
        );
        require(
            !taskSuccessfullyChallenged[referenceTaskIndex],
            "Task already challenged"
        );
        require(
            uint32(block.number) <= taskResponseMetadata.taskRespondedBlock + TASK_CHALLENGE_WINDOW_BLOCK,
            "Challenge period expired"
        );

        uint256 expectedResult = task.strategyId; 
        bool isResponseValid = (expectedResult == taskResponse.calculatedResult);

        if (isResponseValid) {
            emit TaskChallengedUnsuccessfully(referenceTaskIndex, msg.sender);
            return;
        }

        bytes32[] memory hashesOfNonSigningOperators =
            new bytes32[](pubkeysOfNonSigningOperators.length);
        for (uint256 i = 0; i < pubkeysOfNonSigningOperators.length; i++) {
            hashesOfNonSigningOperators[i] = pubkeysOfNonSigningOperators[i].hashG1Point();
        }

        bytes32 signatoryRecordHash =
            keccak256(abi.encodePacked(task.taskCreatedBlock, hashesOfNonSigningOperators));
        require(
            signatoryRecordHash == taskResponseMetadata.hashOfNonSigners,
            "Invalid non-signers"
        );

        taskSuccessfullyChallenged[referenceTaskIndex] = true;
        emit TaskChallengedSuccessfully(referenceTaskIndex, msg.sender);
    }

    function getTaskResponseWindowBlock() external view returns (uint32) {
        return TASK_RESPONSE_WINDOW_BLOCK;
    }

    function _setGenerator(address newGenerator) internal {
        address oldGenerator = generator;
        generator = newGenerator;
        emit GeneratorUpdated(oldGenerator, newGenerator);
    }

    function _setAggregator(address newAggregator) internal {
        address oldAggregator = aggregator;
        aggregator = newAggregator;
        emit AggregatorUpdated(oldAggregator, newAggregator);
    }
}
