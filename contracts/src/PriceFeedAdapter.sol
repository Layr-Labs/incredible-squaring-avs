// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.9;

// import { AggregatorV3Interface } from "chainlink/v0.8/interfaces/AggregatorV3Interface.sol";

// contract PriceFeedAdapter {
//     AggregatorV3Interface internal dataFeed;

//     /**
//      * Network: Sepolia
//      * Aggregator: BTC/USD
//      * Address: 0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
//      */
//     constructor() {
//         dataFeed = AggregatorV3Interface(
//             0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43
//         );
//     }

//     /**
//      * Returns the latest answer.
//      */
//     function getChainlinkDataFeedLatestAnswer() public view returns (int) {
//         // prettier-ignore
//         (
//             /* uint80 roundID */,
//             int answer,
//             /*uint startedAt*/,
//             /*uint timeStamp*/,
//             /*uint80 answeredInRound*/
//         ) = dataFeed.latestRoundData();
//         return answer;
//     }
// }
