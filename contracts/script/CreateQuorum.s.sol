// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {IncredibleSquaringDeploymentLib} from "../script/utils/IncredibleSquaringDeploymentLib.sol";
import {SlashingRegistryCoordinator} from
    "@eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";
import {ISlashingRegistryCoordinatorTypes} from
    "@eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";
import {IStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";

contract CreateQuorum is Script {
    using IncredibleSquaringDeploymentLib for *;

    address internal deployer;
    IncredibleSquaringDeploymentLib.DeploymentData deploymentData;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
        deploymentData = IncredibleSquaringDeploymentLib.readDeploymentJson(block.chainid);
    }

    function run() external {
        vm.startBroadcast(deployer);
        ISlashingRegistryCoordinatorTypes.OperatorSetParam memory _operatorSetParam =
        ISlashingRegistryCoordinatorTypes.OperatorSetParam({
            maxOperatorCount: 3,
            kickBIPsOfOperatorStake: 100,
            kickBIPsOfTotalStake: 1000
        });
        uint96 minimumStake = 0;
        IStakeRegistryTypes.StrategyParams[] memory _strategyParams =
            new IStakeRegistryTypes.StrategyParams[](1);
        IStrategy istrategy = IStrategy(deploymentData.strategy);
        _strategyParams[0] =
            IStakeRegistryTypes.StrategyParams({strategy: istrategy, multiplier: 1});
        SlashingRegistryCoordinator regCoord =
            SlashingRegistryCoordinator(deploymentData.slashingRegistryCoordinator);
        regCoord.createTotalDelegatedStakeQuorum(_operatorSetParam, minimumStake, _strategyParams);

        vm.stopBroadcast();
    }
}
