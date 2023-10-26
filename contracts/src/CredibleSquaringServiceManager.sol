// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

// import "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./ICredibleSquaringTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

// import "./ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from CredibleSquaringServiceManager.
 * @author Layr Labs, Inc.
 * @notice This contract is used for:
 * - initializing the data store by the disperser
 * - confirming the data store by the disperser with inferred aggregated signatures of the quorum
 * - freezing operators as the result of various "challenges"
 */
contract CredibleSquaringServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    /// @notice The current task number
    uint32 public taskNum;

    /// @notice Unit of measure (in blocks) for which Ethereum censorship window will last.

    ICredibleSquaringTaskManager public immutable credibleSquaringTaskManager;
    uint32 public immutable ETHEREUM_CENSORSHIP_WINDOW;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyCredibleSquaringTaskManager() {
        require(
            msg.sender == address(credibleSquaringTaskManager),
            "onlyCredibleSquaringTaskManager: not from credible squaring task manager"
        );
        _;
    }

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
        ISlasher _slasher,
        ICredibleSquaringTaskManager _credibleSquaringTaskManager,
        uint32 _ethereumCensorshipWindow
    ) ServiceManagerBase(_registryCoordinator, _slasher) {
        credibleSquaringTaskManager = _credibleSquaringTaskManager;
        ETHEREUM_CENSORSHIP_WINDOW = _ethereumCensorshipWindow;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    function freezeOperator(
        address operatorAddr
    ) external override onlyCredibleSquaringTaskManager {
        // require(
        //     msg.sender == address(???),
        //     "CredibleSquaringServiceManager.freezeOperator: Only ??? can slash operators"
        // );
        slasher.freezeOperator(operatorAddr);
    }

    /// @notice Returns the current 'taskNumber' for the middleware
    function taskNumber() external view override returns (uint32) {
        return credibleSquaringTaskManager.taskNumber();
    }

    /// @notice Returns the block until which operators must serve.
    function latestServeUntilBlock() public view override returns (uint32) {
        return
            uint32(block.number) +
            credibleSquaringTaskManager.getTaskResponseWindowBlock() +
            ETHEREUM_CENSORSHIP_WINDOW;
    }
}
