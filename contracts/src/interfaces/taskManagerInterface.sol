// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

interface TaskManagerInterface {
    function fetchLatestAggregatedPrice(string memory feedName) external view returns (uint256, uint256);
}