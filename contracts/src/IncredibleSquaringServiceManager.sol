// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IIncredibleSquaringTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";

/**
 * @title Primary entrypoint for procuring services from IncredibleSquaring.
 * @author Layr Labs, Inc.
 */
contract IncredibleSquaringServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IIncredibleSquaringTaskManager
        public immutable incredibleSquaringTaskManager;

    mapping(address => string) public operatorHttpUrls;
    mapping(address => string) public operatorRpcUrls;

    event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl);

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyIncredibleSquaringTaskManager() {
        require(
            msg.sender == address(incredibleSquaringTaskManager),
            "onlyIncredibleSquaringTaskManager: not from credible squaring task manager"
        );
        _;
    }

    modifier onlyOperator() {
        IRegistryCoordinator.OperatorInfo memory operatorInfo = _registryCoordinator.getOperator(msg.sender);
        require(operatorInfo.operatorId != bytes32(0));
        _;
    }

    constructor(
        IAVSDirectory _avsDirectory,
        IRegistryCoordinator _registryCoordinator,
        IStakeRegistry _stakeRegistry,
        IIncredibleSquaringTaskManager _incredibleSquaringTaskManager
    )
        ServiceManagerBase(
            _avsDirectory,
            IPaymentCoordinator(address(0)), // inc-sq doesn't need to deal with payments
            _registryCoordinator,
            _stakeRegistry
        )
    {
        incredibleSquaringTaskManager = _incredibleSquaringTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external onlyIncredibleSquaringTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }

    function registerOperatorConsensusUrl(string memory httpUrl, string memory rpcUrl) onlyOperator external {
        operatorHttpUrls[msg.sender] = httpUrl;
        operatorRpcUrls[msg.sender] = rpcUrl;
        emit OperatorUrlRegistered(msg.sender, httpUrl, rpcUrl);
    }

    function fetchOperatorUrls(address operatorAddress) external view returns(string memory httpUrl, string memory rpcUrl) {
        bytes memory httpUrl = bytes(operatorHttpUrls[operatorAddress]);
        bytes memory rpcUrl = bytes(operatorRpcUrls[operatorAddress]);
        require (httpUrl.length > 0 , "No url registered for requested address");
        require (rpcUrl.length > 0 , "No url registered for requested address");
        return (operatorHttpUrls[operatorAddress], operatorRpcUrls[operatorAddress]);
    }

    function isValidOperator(address operatorAddress) external view returns(bool) {
        IRegistryCoordinator.OperatorInfo memory operatorInfo = _registryCoordinator.getOperator(operatorAddress);
        return operatorInfo.operatorId != bytes32(0);
    }
}
