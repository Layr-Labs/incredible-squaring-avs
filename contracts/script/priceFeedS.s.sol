pragma solidity ^0.8.12;
import "forge-std/Script.sol";


import {PriceFeedAdapterV2} from "src/priceAdapter3.sol";

contract ChainlinkScript is Script {
    function setUp() public {}
     function run() public {
        vm.startBroadcast();
        PriceFeedAdapterV2 priceFeedAdapter = new PriceFeedAdapterV2();
        vm.stopBroadcast();
     }
}