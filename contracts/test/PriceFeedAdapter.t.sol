 pragma solidity 0.8.12;

 import "forge-std/Test.sol";
 import "forge-std/console2.sol";

 

 import { PriceFeedAdapterV1 } from "../src/PriceFeedAdapter2.sol";
 import {PriceFeedAdapterV2 } from "../src/priceAdapter3.sol";
 //import { diaAdapter } from "../src/diaAdapter.sol";

 interface IDIAOracleV2{
    function getValue(string memory) external returns (uint128, uint128);
}


 contract PriceAdapterTest is Test{
    PriceFeedAdapterV2 internal priceFeedAdapter;
   // diaAdapter internal diaFeed;
    address immutable ORACLE = 0xa93546947f3015c986695750b8bbEa8e26D65856;

     function setUp() public {
         priceFeedAdapter = new PriceFeedAdapterV2();
     //    diaFeed = new diaAdapter();
     }

     function testPriceInfo() public {
        
        priceFeedAdapter.getPriceDia("ETH/USD", 1715446929);
    

        
    }

       /**
    * @dev Test that the storage varibles in the sample contract are
    * equal to values returned from DIAOracle 
    * */
    

     

    
  }

