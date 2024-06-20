// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

/**
 * @title contract consuming our priceFeedAvs prices
 * @dev This contract is used to consume the prices from the priceFeedAvs contract 
 */

import "./interfaces/taskManagerInterface.sol";

contract PriceConsumer {
    TaskManagerInterface public taskManager;

    constructor(address _taskManager) {
        taskManager = TaskManagerInterface(_taskManager);
        
    }

    function fetchPrice(string memory feedName) public view returns (uint256, uint256) {
        return taskManager.fetchLatestAggregatedPrice(feedName);
    }
}