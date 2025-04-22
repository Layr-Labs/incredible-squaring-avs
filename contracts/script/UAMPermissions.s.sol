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
    AllocationManager, IAllocationManager
} from "@eigenlayer/contracts/core/AllocationManager.sol";
import {IAVSRegistrar} from "@eigenlayer/contracts/interfaces/IAVSRegistrar.sol";

contract UAMPermissions is Script {
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
        serviceManager.setAppointee(
            deployer, coreData.allocationManager, AllocationManager.setAVSRegistrar.selector
        );

        IAllocationManager _allocationManager = IAllocationManager(coreData.allocationManager);
        _allocationManager.setAVSRegistrar(
            deploymentData.incredibleSquaringServiceManager,
            IAVSRegistrar(deploymentData.slashingRegistryCoordinator)
        );

        serviceManager.setAppointee(
            deploymentData.slashingRegistryCoordinator,
            coreData.allocationManager,
            AllocationManager.createOperatorSets.selector
        );

        serviceManager.setAppointee(
            deploymentData.slasher,
            coreData.allocationManager,
            AllocationManager.slashOperator.selector
        );

        // This should be in another contract
        serviceManager.setAppointee(
            deployer, coreData.allocationManager, AllocationManager.updateAVSMetadataURI.selector
        );

        _allocationManager.updateAVSMetadataURI(
            deploymentData.incredibleSquaringServiceManager, "metadataURI"
        );

        vm.stopBroadcast();
    }
}
