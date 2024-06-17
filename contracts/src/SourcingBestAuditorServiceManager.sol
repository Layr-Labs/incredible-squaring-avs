// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./ISourcingBestAuditorTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from SourcingBestAuditorTaskManager.
 * @author Layr Labs, Inc.
 */
contract SourcingBestAuditorServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    ISourcingBestAuditorTaskManager
        public immutable sourcingBestAuditorTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlySourcingBestAuditorTaskManager() {
        require(
            msg.sender == address(sourcingBestAuditorTaskManager),
            "onlySourcingBestAuditorTaskManager: not from credible task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        ISourcingBestAuditorTaskManager _sourcingBestAuditorTaskManager
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // inc-sq doesn't need to deal with payments
            _registryCoordinator,
            _stakeRegistry
        )
    {
        sourcingBestAuditorTaskManager = _sourcingBestAuditorTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlySourcingBestAuditorTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
