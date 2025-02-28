// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Script} from "forge-std/Script.sol";
import {IncredibleSquaringDeploymentLib} from "./utils/IncredibleSquaringDeploymentLib.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {SetupPaymentsLib} from "./utils/SetupPaymentsLib.sol";
import {IRewardsCoordinatorTypes} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {console2} from "forge-std/console2.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {stdJson} from "forge-std/StdJson.sol";

contract OperatorDirectedPayments is Script {
    using stdJson for string;

    address private deployer;
    CoreDeploymentLib.DeploymentData coreDeployment;
    IncredibleSquaringDeploymentLib.DeploymentData incredibleSquaringDeployment;
    SetupPaymentsLib.OperatorConfig operatorRewardConfig;
    string internal constant filePath = "test/mockData/scratch/payment_info.json";

    function setUp() public {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

        coreDeployment =
            CoreDeploymentLib.readDeploymentJson("script/deployments/core/", block.chainid);
        incredibleSquaringDeployment = IncredibleSquaringDeploymentLib.readDeploymentJson(
            "script/deployments/incredible-squaring/", block.chainid
        );
        operatorRewardConfig = SetupPaymentsLib.readOperatorConfig("operator_reward_config");
    }

    function run() external {
        vm.startBroadcast(deployer);

        string memory json = vm.readFile(filePath);
        uint32 duration = uint32(json.readUint(".duration"));
        uint32 index_to_prove = uint32(json.readUint(".indexToProve"));
        uint256 num_payments = json.readUint(".numPayments");
        address recipient = json.readAddress(".recipient");
        address[] memory earners = json.readAddressArray(".earners");
        bytes32[] memory earner_token_roots = json.readBytes32Array(".earnerTokenRoots");
        uint32 start_time = uint32(previousDivisibleTimestamp(block.timestamp));

        IRewardsCoordinatorTypes.OperatorReward[] memory operator_reward =
            new IRewardsCoordinatorTypes.OperatorReward[](operatorRewardConfig.operator_addr.length);

        for (uint256 i = 0; i < operator_reward.length; i++) {
            operator_reward[i] = IRewardsCoordinatorTypes.OperatorReward({
                operator: operatorRewardConfig.operator_addr[i],
                amount: operatorRewardConfig.amount[i]
            });
        }

        SetupPaymentsLib.createOperatorDirectedAVSRewardsSubmissions(
            incredibleSquaringDeployment.strategy,
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            operator_reward,
            num_payments,
            duration,
            start_time
        );

        vm.stopBroadcast();
    }

    function previousDivisibleTimestamp(
        uint256 blockTimestamp
    ) public pure returns (uint256) {
        uint256 daySeconds = 86_400;

        // Calculate the remainder to check how far blockTimestamp is from the nearest multiple of daySeconds
        uint256 remainder = blockTimestamp % daySeconds;

        if (remainder == 0) {
            // If blockTimestamp is already divisible by daySeconds, go two steps back
            return blockTimestamp - (2 * daySeconds);
        } else {
            // Otherwise, move back to the previous divisible timestamp and then one more step
            return blockTimestamp - remainder - daySeconds;
        }
    }
}
