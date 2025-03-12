// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Vm} from "forge-std/Vm.sol";
import {ContractsRegistry} from "../../src/ContractsRegistry.sol";
import {CoreDeploymentLib} from "./CoreDeploymentLib.sol";
import {IncredibleSquaringDeploymentLib} from "./IncredibleSquaringDeploymentLib.sol";

library WriteToContractsRegistryLib {
    function writeCoreContractsToRegistry(
        address contracts_registry_addr,
        CoreDeploymentLib.DeploymentData memory deploymentdata
    ) internal {
        ContractsRegistry contractsRegistry = ContractsRegistry(contracts_registry_addr);
        contractsRegistry.registerContract(
            "delegationManager", address(deploymentdata.delegationManager)
        );
        contractsRegistry.registerContract(
            "strategyManager", address(deploymentdata.strategyManager)
        );
        contractsRegistry.registerContract("avsDirectory", address(deploymentdata.avsDirectory));
        contractsRegistry.registerContract(
            "allocationManager", address(deploymentdata.allocationManager)
        );
        require(deploymentdata.permissionController != address(0));
        contractsRegistry.registerContract(
            "permissionController", address(deploymentdata.permissionController)
        );
    }

    function writeIncredibleSquaringContractsToRegistry(
        address contracts_registry_addr,
        IncredibleSquaringDeploymentLib.DeploymentData memory deploymentdata
    ) internal {
        ContractsRegistry contractsRegistry = ContractsRegistry(contracts_registry_addr);

        contractsRegistry.registerContract(
            "incredibleSquaringTaskManager",
            address(deploymentdata.incredibleSquaringTaskManager)
        );
        contractsRegistry.registerContract("erc20MockStrategy", address(deploymentdata.strategy));
        contractsRegistry.registerContract(
            "incredibleSquaringRegistryCoordinator", address(deploymentdata.registryCoordinator)
        );
        contractsRegistry.registerContract(
            "incredibleSquaringOperatorStateRetriever",
            address(deploymentdata.operatorStateRetriever)
        );
        contractsRegistry.registerContract(
            "incredibleSquaringServiceManager",
            address(deploymentdata.incredibleSquaringServiceManager)
        );
    }
}
