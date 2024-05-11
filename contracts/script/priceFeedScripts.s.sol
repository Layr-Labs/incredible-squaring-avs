pragma solidity ^0.8.12;
import "forge-std/Script.sol";
import {PriceFeedAdapter} from "src/PriceFeedAdapter.sol";

contract ChainlinkScript is Script {
    function setUp() public {}
     function run() public {
        vm.startBroadcast();
        PriceFeedAdapter priceFeedAdapter = new PriceFeedAdapter();
        vm.stopBroadcast();
     }
}