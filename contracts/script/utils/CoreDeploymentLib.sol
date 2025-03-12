// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {console2} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {DelegationManager} from "@eigenlayer/contracts/core/DelegationManager.sol";
import {StrategyManager} from "@eigenlayer/contracts/core/StrategyManager.sol";
import {AVSDirectory} from "@eigenlayer/contracts/core/AVSDirectory.sol";
import {EigenPodManager} from "@eigenlayer/contracts/pods/EigenPodManager.sol";
import {RewardsCoordinator} from "@eigenlayer/contracts/core/RewardsCoordinator.sol";
import {StrategyBase} from "@eigenlayer/contracts/strategies/StrategyBase.sol";
import {EigenPod} from "@eigenlayer/contracts/pods/EigenPod.sol";
import {IETHPOSDeposit} from "@eigenlayer/contracts/interfaces/IETHPOSDeposit.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import {PauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {IStrategyManager} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {IEigenPodManager} from "@eigenlayer/contracts/interfaces/IEigenPodManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IPauserRegistry} from "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";

import {UpgradeableProxyLib} from "./UpgradeableProxyLib.sol";
import {CoreDeploymentLib} from "@eigenlayer-middleware/test/utils/CoreDeployLib.sol";

library CoreDeploymentParsingLib {
    using stdJson for *;
    using Strings for *;
    using UpgradeableProxyLib for address;

    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function readDeploymentConfigValues(
        string memory directoryPath,
        string memory fileName
    ) internal view returns (CoreDeploymentLib.DeploymentConfigData memory) {
        string memory pathToFile = string.concat(directoryPath, fileName);

        require(vm.exists(pathToFile), "CoreDeployment: Deployment config file does not exist");

        string memory json = vm.readFile(pathToFile);

        CoreDeploymentLib.DeploymentConfigData memory data;

        // StrategyManager start
        data.strategyManager.initPausedStatus = json.readUint(".strategyManager.initPausedStatus");
        data.strategyManager.initialOwner = json.readAddress(".strategyManager.initialOwner");
        data.strategyManager.initialStrategyWhitelister =
            json.readAddress(".strategyManager.initialStrategyWhitelister");
        // StrategyManager config end

        // DelegationManager config start
        data.delegationManager.initPausedStatus =
            json.readUint(".delegationManager.initPausedStatus");
        data.delegationManager.initialOwner = json.readAddress(".delegationManager.initialOwner");
        data.delegationManager.minWithdrawalDelayBlocks =
            uint32(json.readUint(".delegationManager.minWithdrawalDelayBlocks"));
        // DelegationManager config end

        // EigenPodManager config start
        data.eigenPodManager.initPausedStatus = json.readUint(".eigenPodManager.initPausedStatus");
        data.eigenPodManager.initialOwner = json.readAddress(".eigenPodManager.initialOwner");
        // EigenPodManager config end

        // AllocationManager config start
        data.allocationManager.initPausedStatus =
            json.readUint(".allocationManager.initPausedStatus");
        data.allocationManager.initialOwner = json.readAddress(".allocationManager.initialOwner");
        data.allocationManager.deallocationDelay =
            uint32(json.readUint(".allocationManager.deallocationDelay"));
        data.allocationManager.allocationConfigurationDelay =
            uint32(json.readUint(".allocationManager.allocationConfigurationDelay"));
        // AllocationManager config end

        // StrategyFactory config start
        data.strategyFactory.initPausedStatus = json.readUint(".strategyFactory.initPausedStatus");
        data.strategyFactory.initialOwner = json.readAddress(".strategyFactory.initialOwner");
        // StrategyFactory config end

        // AVSDirectory config start
        data.avsDirectory.initPausedStatus = json.readUint(".avsDirectory.initPausedStatus");
        data.avsDirectory.initialOwner = json.readAddress(".avsDirectory.initialOwner");
        // AVSDirectory config end

        // RewardsCoordinator config start
        data.rewardsCoordinator.initPausedStatus =
            json.readUint(".rewardsCoordinator.initPausedStatus");
        data.rewardsCoordinator.initialOwner = json.readAddress(".rewardsCoordinator.initialOwner");
        data.rewardsCoordinator.rewardsUpdater =
            json.readAddress(".rewardsCoordinator.rewardsUpdater");

        data.rewardsCoordinator.activationDelay =
            uint32(json.readUint(".rewardsCoordinator.activationDelay"));
        data.rewardsCoordinator.defaultSplitBips =
            uint16(json.readUint(".rewardsCoordinator.defaultSplitBips"));
        data.rewardsCoordinator.calculationIntervalSeconds =
            uint32(json.readUint(".rewardsCoordinator.calculationIntervalSeconds"));
        data.rewardsCoordinator.maxRewardsDuration =
            uint32(json.readUint(".rewardsCoordinator.maxRewardsDuration"));
        data.rewardsCoordinator.maxRetroactiveLength =
            uint32(json.readUint(".rewardsCoordinator.maxRetroactiveLength"));
        data.rewardsCoordinator.maxFutureLength =
            uint32(json.readUint(".rewardsCoordinator.maxFutureLength"));
        data.rewardsCoordinator.genesisRewardsTimestamp =
            uint32(json.readUint(".rewardsCoordinator.genesisRewardsTimestamp"));
        // RewardsCoordinator config end

        data.ethPOSDeposit.ethPOSDepositAddress = address(1);

        return data;
    }

    function readDeploymentConfigValues(
        string memory directoryPath,
        uint256 chainId
    ) internal view returns (CoreDeploymentLib.DeploymentConfigData memory) {
        return
            readDeploymentConfigValues(directoryPath, string.concat(vm.toString(chainId), ".json"));
    }

    function readDeploymentJson(
        string memory directoryPath,
        uint256 chainId
    ) internal view returns (CoreDeploymentLib.DeploymentData memory) {
        return readDeploymentJson(directoryPath, string.concat(vm.toString(chainId), ".json"));
    }

    function readDeploymentJson(
        string memory path,
        string memory fileName
    ) internal view returns (CoreDeploymentLib.DeploymentData memory) {
        string memory pathToFile = string.concat(path, fileName);

        require(vm.exists(pathToFile), "CoreDeployment: Deployment file does not exist");

        string memory json = vm.readFile(pathToFile);

        CoreDeploymentLib.DeploymentData memory data;
        data.delegationManager = json.readAddress(".addresses.delegationManager");
        data.avsDirectory = json.readAddress(".addresses.avsDirectory");
        data.strategyManager = json.readAddress(".addresses.strategyManager");
        data.eigenPodManager = json.readAddress(".addresses.eigenPodManager");
        data.allocationManager = json.readAddress(".addresses.allocationManager");
        data.eigenPodBeacon = json.readAddress(".addresses.eigenPodBeacon");
        data.pauserRegistry = json.readAddress(".addresses.pauserRegistry");
        data.strategyFactory = json.readAddress(".addresses.strategyFactory");
        data.strategyBeacon = json.readAddress(".addresses.strategyBeacon");
        data.rewardsCoordinator = json.readAddress(".addresses.rewardsCoordinator");
        data.permissionController = json.readAddress(".addresses.permissionController");

        return data;
    }

    /// TODO: Need to be able to read json from eigenlayer-contracts repo for holesky/mainnet and output the json here
    function writeDeploymentJson(
        CoreDeploymentLib.DeploymentData memory data
    ) internal {
        writeDeploymentJson("deployments/core/", block.chainid, data);
    }

    function writeDeploymentJson(
        string memory path,
        uint256 chainId,
        CoreDeploymentLib.DeploymentData memory data
    ) internal {
        address proxyAdmin = address(UpgradeableProxyLib.getProxyAdmin(data.strategyManager));

        string memory deploymentData = _generateDeploymentJson(data, proxyAdmin);

        string memory fileName = string.concat(path, vm.toString(chainId), ".json");
        if (!vm.exists(path)) {
            vm.createDir(path, true);
        }

        vm.writeFile(fileName, deploymentData);
        console2.log("Deployment artifacts written to:", fileName);
    }

    function _generateDeploymentJson(
        CoreDeploymentLib.DeploymentData memory data,
        address proxyAdmin
    ) private view returns (string memory) {
        return string.concat(
            '{"lastUpdate":{"timestamp":"',
            vm.toString(block.timestamp),
            '","block_number":"',
            vm.toString(block.number),
            '"},"addresses":',
            _generateContractsJson(data, proxyAdmin),
            "}"
        );
    }

    function _generateContractsJson(
        CoreDeploymentLib.DeploymentData memory data,
        address proxyAdmin
    ) private view returns (string memory) {
        /// TODO: namespace contracts -> {avs, core}
        return string.concat(
            '{"proxyAdmin":"',
            proxyAdmin.toHexString(),
            '","delegationManager":"',
            data.delegationManager.toHexString(),
            '","delegationManagerImpl":"',
            data.delegationManager.getImplementation().toHexString(),
            '","avsDirectory":"',
            data.avsDirectory.toHexString(),
            '","avsDirectoryImpl":"',
            data.avsDirectory.getImplementation().toHexString(),
            '","strategyManager":"',
            data.strategyManager.toHexString(),
            '","strategyManagerImpl":"',
            data.strategyManager.getImplementation().toHexString(),
            '","eigenPodManager":"',
            data.eigenPodManager.toHexString(),
            '","eigenPodManagerImpl":"',
            data.eigenPodManager.getImplementation().toHexString(),
            '","allocationManager":"',
            data.allocationManager.toHexString(),
            '","allocationManagerImpl":"',
            data.allocationManager.getImplementation().toHexString(),
            '","eigenPodBeacon":"',
            data.eigenPodBeacon.toHexString(),
            '","pauserRegistry":"',
            data.pauserRegistry.toHexString(),
            '","pauserRegistryImpl":"',
            data.pauserRegistry.getImplementation().toHexString(),
            '","strategyFactory":"',
            data.strategyFactory.toHexString(),
            '","strategyFactoryImpl":"',
            data.strategyFactory.getImplementation().toHexString(),
            '","strategyBeacon":"',
            data.strategyBeacon.toHexString(),
            '","rewardsCoordinator":"',
            data.rewardsCoordinator.toHexString(),
            '","rewardsCoordinatorImpl":"',
            data.rewardsCoordinator.getImplementation().toHexString(),
            '","permissionController":"',
            data.permissionController.toHexString(),
            '","permissionControllerImpl":"',
            data.permissionController.getImplementation().toHexString(),
            '"}'
        );
    }
}
