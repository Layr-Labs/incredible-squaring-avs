// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

/**
 * @title foundry testing for PriceConsumer contract
 *
 * @dev  This contract is used to test the PriceConsumer contract
 */
import "forge-std/Test.sol";
import "forge-std/console2.sol";

import { PriceConsumer } from "../src/priceConsumer.sol";

contract PriceConsumerTest {
    PriceConsumer internal priceConsumer;

    event TestValue(uint256 value, uint256 value2);
    
    function setUp() public {
        priceConsumer = new PriceConsumer(address(0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5));
    }

    function test_fetch_btc_usd_price01() public  {
        (uint256 priceFeed1, uint256 priceFeed2) = priceConsumer.fetchPrice("BTC/USD");
        emit TestValue(priceFeed1, priceFeed2);
    }
}
