 pragma solidity 0.8.12;

 import "forge-std/Test.sol";
 import "forge-std/console2.sol";

 

 
 import {PriceFeedAdapterV2 } from "../src/priceAdapter3.sol";
 //import { diaAdapter } from "../src/diaAdapter.sol";

 interface IDIAOracleV2{
    function getValue(string memory) external returns (uint128, uint128);
}


 contract PriceAdapterTest is Test {
    PriceFeedAdapterV2 internal priceFeedAdapter;

    event TestValue(int value);
    event TestValue2(uint128 value);

    // Define a constant address for the Chainlink feed; make sure this is correct and matches the network
    address internal constant CHAINLINK_FEED_ADDRESS = address(0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43); // Sepolia BTC/USD

    function setUp() public {
        priceFeedAdapter = new PriceFeedAdapterV2();
        // Ensure that the symbol used in `addFeed` matches the one used in `getLatestPrice`
        priceFeedAdapter.addFeed("BTC/USD", CHAINLINK_FEED_ADDRESS);
    }

    function testFetchBtcUsdPrice() public {
        // Fetch the BTC/USD price using the PriceFeedAdapter
        int256 price = priceFeedAdapter.getLatestPrice("BTC/USD");
         emit TestValue(price);

         uint128 diaPrice = priceFeedAdapter.getPriceDia("BTC/USD", 1715446929); // https://www.unixtimestamp.com/
         emit TestValue2(diaPrice);
        

    }
}

