pragma solidity ^0.8.12;
import "forge-std/Script.sol";
import "forge-std/console.sol";
import {PriceFeedAdapter} from "src/PriceFeedAdapter.sol";

contract ChainlinkScript is Script {
    function setUp() public {}
     function run() public {
        vm.startBroadcast();
        PriceFeedAdapter priceFeedAdapter = PriceFeedAdapter(address(0x09635F643e140090A9A8Dcd712eD6285858ceBef));
        priceFeedAdapter.addFeed("btc/usd", address(0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43));
        console.logInt(priceFeedAdapter.getLatestPrice("btc/usd"));
        vm.stopBroadcast();
     }
}