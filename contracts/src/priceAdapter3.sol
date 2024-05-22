// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

import "@openzeppelin/contracts/access/Ownable.sol";




interface IDIAOracleV2{
    function getValue(string memory) external returns (uint128, uint128);
}

contract PriceFeedAdapterV2 is Ownable {
        
    int public ChainlinkAdapterPrice;
    uint128 public DiaAdapterPrice;
   // int224  public Api3AdapterPrice;
   // uint128 public timestampOflatestPrice
    address immutable diaORACLE = 0xCD5F78206ca1FF96Ff4c043C61a2299B2Febf3cB; // sepolia



    mapping(string => AggregatorV3Interface) public feeds;

    event FeedAdded(string symbol, address feedAddress);
    event FeedRemoved(string symbol);
    event chainlinkAdapterPrice(string symbol, int256 _ChainlinkAdapterPrice);
    event diaAdapterPrice(string symbol, uint128 _DiaAdapterPrice);

    function addFeed(string memory _symbol, address _feedAddress) external onlyOwner  {
        require(_feedAddress != address(0), "Feed address cannot be zero.");
        require(bytes(_symbol).length > 0, "Symbol cannot be empty.");
        feeds[_symbol] = AggregatorV3Interface(_feedAddress);
        emit FeedAdded(_symbol, _feedAddress);
    }

    function removeFeed(string memory _symbol) external onlyOwner  {
        require(address(feeds[_symbol]) != address(0), "Feed not found.");
        delete feeds[_symbol];
        emit FeedRemoved(_symbol);
    }

    function getLatestPrice(string memory _symbol) external returns (int) {
        AggregatorV3Interface feed = feeds[_symbol];
        require(address(feed) != address(0), "Feed not found.");
        (, int price,, uint256 updatedAt,) = feed.latestRoundData();
        require(block.timestamp - updatedAt < 1 hours, "Data is stale"); // Stale data check
        ChainlinkAdapterPrice = price;
        emit chainlinkAdapterPrice(_symbol, ChainlinkAdapterPrice);
        return price;
        
    }

     function getPriceDia(string memory _symbol, uint128 _timestampOflatestPrice) external returns(uint128) {
        
        (DiaAdapterPrice, _timestampOflatestPrice) = IDIAOracleV2(diaORACLE).getValue(_symbol);
        emit diaAdapterPrice(_symbol, DiaAdapterPrice);
        return DiaAdapterPrice;
        
        
    }

/// Maybe.....

   // function getPriceApi3(address proxy) //0xc1EF355A02fAeC5fAAd99149aF4307A3B0a367C8 sepolia
   //     public
   //     view
   //     returns (int224 value, uint32 timestamp)
     //   {
       //   (value, timestamp) = IProxy(proxy).read();
          
          

       // }
   
    function checkPriceAgeDia(uint128 maxTimePassed, uint128 _timestampOflatestPrice) external view returns (bool inTime){
         if((block.timestamp - _timestampOflatestPrice) < maxTimePassed){
             inTime = true;
         } else {
             inTime = false;
         }
    }


}