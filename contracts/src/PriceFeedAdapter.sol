
// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract PriceFeedAdapter is Ownable {
    mapping(string => AggregatorV3Interface) public feeds;

    event FeedAdded(string symbol, address feedAddress);
    event FeedRemoved(string symbol);

    function addFeed(string memory _symbol, address _feedAddress) external onlyOwner {
        require(_feedAddress != address(0), "Feed address cannot be zero.");
        require(bytes(_symbol).length > 0, "Symbol cannot be empty.");
        feeds[_symbol] = AggregatorV3Interface(_feedAddress);
        emit FeedAdded(_symbol, _feedAddress);
    }

    function removeFeed(string memory _symbol) external onlyOwner {
        require(address(feeds[_symbol]) != address(0), "Feed not found.");
        delete feeds[_symbol];
        emit FeedRemoved(_symbol);
    }

    function getLatestPrice(string memory _symbol) external view returns (int) {
        AggregatorV3Interface feed = feeds[_symbol];
        require(address(feed) != address(0), "Feed not found.");
        (, int price,, uint256 updatedAt,) = feed.latestRoundData();
        require(block.timestamp - updatedAt < 1 hours, "Data is stale"); // Stale data check
        return price;
    }
}