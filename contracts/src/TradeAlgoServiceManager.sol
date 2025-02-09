// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./ITradeAlgoTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entry point for procuring services from TradeAlgo.
 * @author
 * @notice Manages interaction between the TradeAlgo task manager and EigenLayer's AVS infrastructure.
 */
contract TradeAlgoServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    ITradeAlgoTaskManager public immutable tradeAlgoTaskManager;

    /// @notice Ensures that only the TradeAlgo task manager can call a function.
    modifier onlyTradeAlgoTaskManager() {
        require(
            msg.sender == address(tradeAlgoTaskManager),
            "onlyTradeAlgoTaskManager: not from TradeAlgo task manager"
        );
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRewardsCoordinator _rewardsCoordinator,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        ITradeAlgoTaskManager _tradeAlgoTaskManager
    )
        ServiceManagerBase(_avsDirectory, _rewardsCoordinator, _registryCoordinator, _stakeRegistry)
    {
        tradeAlgoTaskManager = _tradeAlgoTaskManager;
    }

    /// @notice Initializes the contract's owner and the rewards initiator.
    /// @param initialOwner The address to be set as the initial owner of the contract.
    /// @param rewardsInitiator The address to be set as the rewards initiator for the contract.
    function initialize(address initialOwner, address rewardsInitiator) external initializer {
        __ServiceManagerBase_init(initialOwner, rewardsInitiator);
    }

    /// @notice Called during challenge resolution to forward a call to the Slasher contract, freezing the `operator`.
    /// @dev The Slasher contract is still under active development, and its interface might change.
    function freezeOperator(address operatorAddr) external onlyTradeAlgoTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
