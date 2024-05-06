pragma solidity ^0.8.12;
import "forge-std/Script.sol";
import {PriceFeedAdapterV1} from "src/PriceFeedAdapter2.sol";

contract ChainlinkScript is Script {
    function setUp() public {}
     function run() public {
        vm.startBroadcast();
        PriceFeedAdapterV1 priceFeedAdapter = new PriceFeedAdapterV1();
        vm.stopBroadcast();
     }
}