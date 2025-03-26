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
import {
    IncredibleSquaringServiceManager,
    IServiceManager
} from "../src/IncredibleSquaringServiceManager.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {
    AllocationManager,
    IAllocationManager,
    IAllocationManagerTypes
} from "@eigenlayer/contracts/core/AllocationManager.sol";
import {OperatorSet} from "@eigenlayer/contracts/libraries/OperatorSetLib.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

contract ModifyAllocations is Script {
    using IncredibleSquaringDeploymentLib for *;

    address internal deployer;
    IncredibleSquaringDeploymentLib.DeploymentData deploymentData;
    CoreDeploymentLib.DeploymentData internal coreData;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
        deploymentData = IncredibleSquaringDeploymentLib.readDeploymentJson(block.chainid);
        coreData = CoreDeploymentLib.readDeploymentJson("script/deployments/core/", block.chainid);
    }

    function run() external {
        vm.startBroadcast(deployer);
        IServiceManager serviceManager =
            IServiceManager(deploymentData.incredibleSquaringServiceManager);
        IAllocationManager _allocationManager = IAllocationManager(coreData.allocationManager);

        OperatorSet memory operatorSet = OperatorSet({avs: address(serviceManager), id: 0});

        IStrategy[] memory istrategy = new IStrategy[](1);
        istrategy[0] = IStrategy(deploymentData.strategy);

        uint64[] memory newMagnitude = new uint64[](1);
        newMagnitude[0] = 100_000_000;

        // Construct the allocation parameters
        IAllocationManagerTypes.AllocateParams memory allocateParams = IAllocationManagerTypes
            .AllocateParams({
            operatorSet: operatorSet,
            strategies: istrategy,
            newMagnitudes: newMagnitude
        });

        IAllocationManagerTypes.AllocateParams[] memory allocateParamsArray =
            new IAllocationManagerTypes.AllocateParams[](1);
        allocateParamsArray[0] = allocateParams;

        _allocationManager.modifyAllocations(deployer, allocateParamsArray);

        vm.stopBroadcast();
    }
}
