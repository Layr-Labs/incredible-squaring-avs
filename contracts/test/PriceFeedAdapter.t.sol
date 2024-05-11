pragma solidity 0.8.12;

import "forge-std/Test.sol";
import "forge-std/console2.sol";

import { PriceFeedAdapter } from "../src/PriceFeedAdapter.sol";

contract PriceAdapterTest {
    PriceFeedAdapter internal priceFeedAdapter;

    event TestValue(int value);

    function setUp() public {
        priceFeedAdapter = new PriceFeedAdapter();

        // BTC/USD on goreli
        priceFeedAdapter.addFeed("btc/usd", address(0xA39434A63A52E749F02807ae27335515BA4b07F7));
    }

    function test_fetch_btc_usd_price() public {
        emit TestValue(priceFeedAdapter.getLatestPrice("btc/usd"));
    }
}