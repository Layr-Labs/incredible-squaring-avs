// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {ISocketRegistry, SocketRegistry} from "@eigenlayer-middleware/src/SocketRegistry.sol";
import {
    ISlashingRegistryCoordinator,
    ISlashingRegistryCoordinatorTypes
} from "@eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";
import {SlashingRegistryCoordinator} from "@eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {IPermissionController} from "@eigenlayer/contracts/interfaces/IPermissionController.sol";
import {
    IncredibleSquaringServiceManager,
    IServiceManager,
    IIncredibleSquaringTaskManager
} from "../../src/IncredibleSquaringServiceManager.sol";
import {IncredibleSquaringTaskManager} from "../../src/IncredibleSquaringTaskManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
// import {Quorum} from "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistryEventsAndErrors.sol";
import {UpgradeableProxyLib} from "./UpgradeableProxyLib.sol";
import {CoreDeploymentLib} from "./CoreDeploymentLib.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {InstantSlasher} from "@eigenlayer-middleware/src/slashers/InstantSlasher.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {SocketRegistry} from "@eigenlayer-middleware/src/SocketRegistry.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IAllocationManager} from "@eigenlayer/contracts/interfaces/IAllocationManager.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {CoreDeploymentLib} from "./CoreDeploymentLib.sol";

import {
    RegistryCoordinator,
    IBLSApkRegistry,
    IIndexRegistry,
    IStakeRegistry
} from
// ISocketRegistry
"@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {IStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import {PauserRegistry, IPauserRegistry} from "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

library IncredibleSquaringDeploymentLib {
    using stdJson for *;
    using Strings for *;
    using UpgradeableProxyLib for address;

    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct DeploymentData {
        address incredibleSquaringServiceManager;
        address incredibleSquaringTaskManager;
        address registryCoordinator;
        address operatorStateRetriever;
        address blsapkRegistry;
        address indexRegistry;
        address stakeRegistry;
        address socketRegistry;
        address strategy;
        address pauserRegistry;
        address token;
        address slasher;
    }

    struct IncredibleSquaringSetupConfig {
        uint256 numQuorums;
        uint256[] operatorParams;
        address operator_addr;
        address task_generator_addr;
        address aggregator_addr;
    }

    function deployContracts(
        address proxyAdmin,
        address strategy,
        IncredibleSquaringSetupConfig memory isConfig,
        address admin
    ) internal returns (DeploymentData memory) {
        /// read EL deployment address
        DeploymentData memory result;

        // First, deploy upgradeable proxy contracts that will point to the implementations.
        OperatorStateRetriever operatorStateRetriever = new OperatorStateRetriever();
        result.incredibleSquaringServiceManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.stakeRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.incredibleSquaringTaskManager = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.registryCoordinator = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.blsapkRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.indexRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.socketRegistry = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.slasher = UpgradeableProxyLib.setUpEmptyProxy(proxyAdmin);
        result.strategy = strategy;

        result.operatorStateRetriever = address(operatorStateRetriever);
        CoreDeploymentLib.DeploymentData memory coredata =
            readCoreDeploymentJson("script/deployments/core/", block.chainid);
        // Deploy the implementation contracts, using the proxy contracts as inputs

        address stakeRegistryImpl = address(
            new StakeRegistry(
                ISlashingRegistryCoordinator(result.registryCoordinator),
                IDelegationManager(coredata.delegationManager),
                IAVSDirectory(coredata.avsDirectory),
                IAllocationManager(coredata.allocationManager)
            )
        );

        address blsApkRegistryImpl = address(new BLSApkRegistry(IRegistryCoordinator(result.registryCoordinator)));
        address indexRegistryimpl = address(new IndexRegistry(IRegistryCoordinator(result.registryCoordinator)));
        address instantSlasherImpl = address(
            new InstantSlasher(
                IAllocationManager(coredata.allocationManager),
                ISlashingRegistryCoordinator(result.registryCoordinator),
                result.incredibleSquaringTaskManager
            )
        );

        address registryCoordinatorImpl = address(
            new SlashingRegistryCoordinator(
                IStakeRegistry(result.stakeRegistry),
                IBLSApkRegistry(result.blsapkRegistry),
                IIndexRegistry(result.indexRegistry),
                ISocketRegistry(result.socketRegistry),
                IAllocationManager(coredata.allocationManager),
                IPauserRegistry(coredata.pauserRegistry)
            )
        );

        address[] memory pausers = new address[](2);
        pausers[0] = admin;
        pausers[1] = admin;
        PauserRegistry pausercontract = new PauserRegistry(pausers, admin);

        IStrategy[] memory deployedStrategyArray = new IStrategy[](1);
        deployedStrategyArray[0] = IStrategy(strategy);
        uint256 numStrategies = deployedStrategyArray.length;

        uint256 numQuorums = isConfig.numQuorums;
        IRegistryCoordinator.OperatorSetParam[] memory quorumsOperatorSetParams =
            new IRegistryCoordinator.OperatorSetParam[](numQuorums);
        uint256[] memory operator_params = isConfig.operatorParams;

        for (uint256 i = 0; i < numQuorums; i++) {
            quorumsOperatorSetParams[i] = ISlashingRegistryCoordinatorTypes.OperatorSetParam({
                maxOperatorCount: uint32(operator_params[i]),
                kickBIPsOfOperatorStake: uint16(operator_params[i + 1]),
                kickBIPsOfTotalStake: uint16(operator_params[i + 2])
            });
        }
        // // set to 0 for every quorum
        uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
        IStakeRegistryTypes.StrategyParams[][] memory quorumsStrategyParams =
            new IStakeRegistryTypes.StrategyParams[][](numQuorums);
        for (uint256 i = 0; i < numQuorums; i++) {
            quorumsStrategyParams[i] = new IStakeRegistryTypes.StrategyParams[](numStrategies);
            for (uint256 j = 0; j < numStrategies; j++) {
                quorumsStrategyParams[i][j] = IStakeRegistryTypes.StrategyParams({
                    strategy: deployedStrategyArray[j],
                    // setting this to 1 ether since the divisor is also 1 ether
                    // therefore this allows an operator to register with even just 1 token
                    // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
                    //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                    multiplier: 1 ether
                });
            }
        }

        IStakeRegistryTypes.StakeType[] memory stake_type = new IStakeRegistryTypes.StakeType[](1);
        stake_type[0] = IStakeRegistryTypes.StakeType.TOTAL_SLASHABLE;
        uint32[] memory look_ahead_period = new uint32[](1);
        look_ahead_period[0] = 0;
        bytes memory upgradeCall = abi.encodeCall(
            SlashingRegistryCoordinator.initialize, (admin, admin, admin, 0, result.incredibleSquaringServiceManager)
        );

        UpgradeableProxyLib.upgrade(result.stakeRegistry, stakeRegistryImpl);
        UpgradeableProxyLib.upgrade(result.blsapkRegistry, blsApkRegistryImpl);
        UpgradeableProxyLib.upgrade(result.indexRegistry, indexRegistryimpl);
        UpgradeableProxyLib.upgradeAndCall(result.registryCoordinator, registryCoordinatorImpl, upgradeCall);
        console2.log("rewarr");
        console2.log(coredata.rewardsCoordinator);
        console2.log(coredata.permissionController);
        require(coredata.permissionController != address(0));
        IncredibleSquaringServiceManager incredibleSquaringServiceManagerImpl = new IncredibleSquaringServiceManager(
            (IAVSDirectory(coredata.avsDirectory)),
            IRegistryCoordinator(result.registryCoordinator),
            IStakeRegistry(result.stakeRegistry),
            coredata.rewardsCoordinator,
            IAllocationManager(coredata.allocationManager),
            IPermissionController(coredata.permissionController),
            IIncredibleSquaringTaskManager(result.incredibleSquaringTaskManager)
        );
        IncredibleSquaringTaskManager incredibleSquaringTaskManagerImpl = new IncredibleSquaringTaskManager(
            IRegistryCoordinator(result.registryCoordinator),
            IPauserRegistry(address(pausercontract)),
            30,
            result.incredibleSquaringServiceManager
        );
        bytes memory servicemanagerupgradecall =
            abi.encodeCall(IncredibleSquaringServiceManager.initialize, (admin, admin));
        UpgradeableProxyLib.upgradeAndCall(
            result.incredibleSquaringServiceManager,
            address(incredibleSquaringServiceManagerImpl),
            servicemanagerupgradecall
        );

        bytes memory taskmanagerupgradecall = abi.encodeCall(
            IncredibleSquaringTaskManager.initialize,
            (admin, isConfig.aggregator_addr, isConfig.task_generator_addr, coredata.allocationManager, result.slasher)
        );
        UpgradeableProxyLib.upgradeAndCall(
            result.incredibleSquaringTaskManager, address(incredibleSquaringTaskManagerImpl), (taskmanagerupgradecall)
        );

        UpgradeableProxyLib.upgrade(result.slasher, instantSlasherImpl);

        address socketRegistryImpl =
            address(new SocketRegistry(ISlashingRegistryCoordinator(result.registryCoordinator)));
        UpgradeableProxyLib.upgrade(result.socketRegistry, socketRegistryImpl);

        verify_deployment(result);

        return result;
    }

    function readDeploymentJson(uint256 chainId) internal returns (DeploymentData memory) {
        return readDeploymentJson("script/deployments/incredible-squaring/", chainId);
    }

    function readIncredibleSquaringConfigJson(string memory directoryPath)
        internal
        returns (IncredibleSquaringSetupConfig memory)
    {
        string memory fileName = string.concat(directoryPath, ".json");
        require(vm.exists(fileName), "Deployment file does not exist");
        string memory json = vm.readFile(fileName);

        IncredibleSquaringSetupConfig memory data;
        data.numQuorums = json.readUint(".num_quorums");
        data.operatorParams = json.readUintArray(".operator_params");
        data.aggregator_addr = json.readAddress(".aggregator_addr");
        data.operator_addr = json.readAddress(".operator_addr");
        data.task_generator_addr = json.readAddress(".task_generator_addr");
        return data;
    }

    function readDeploymentJson(string memory directoryPath, uint256 chainId)
        internal
        returns (DeploymentData memory)
    {
        string memory fileName = string.concat(directoryPath, vm.toString(chainId), ".json");

        require(vm.exists(fileName), "Deployment file does not exist");

        string memory json = vm.readFile(fileName);

        DeploymentData memory data;
        data.incredibleSquaringServiceManager = json.readAddress(".addresses.incredibleSquaringServiceManager");
        data.incredibleSquaringTaskManager = json.readAddress(".addresses.incredibleSquaringTaskManager");
        data.registryCoordinator = json.readAddress(".addresses.registryCoordinator");
        data.operatorStateRetriever = json.readAddress(".addresses.operatorStateRetriever");
        data.stakeRegistry = json.readAddress(".addresses.stakeRegistry");
        data.strategy = json.readAddress(".addresses.strategy");
        data.token = json.readAddress(".addresses.token");
        data.slasher = json.readAddress(".addresses.instantSlasher");

        return data;
    }

    /// write to default output path
    function writeDeploymentJson(DeploymentData memory data) internal {
        writeDeploymentJson("script/deployments/incredible-squaring/", block.chainid, data);
    }

    function writeDeploymentJson(string memory outputPath, uint256 chainId, DeploymentData memory data) internal {
        address proxyAdmin = address(UpgradeableProxyLib.getProxyAdmin(data.incredibleSquaringServiceManager));

        string memory deploymentData = _generateDeploymentJson(data, proxyAdmin);

        string memory fileName = string.concat(outputPath, vm.toString(chainId), ".json");
        if (!vm.exists(outputPath)) {
            vm.createDir(outputPath, true);
        }

        vm.writeFile(fileName, deploymentData);
        console2.log("Deployment artifacts written to:", fileName);
    }

    function _generateDeploymentJson(DeploymentData memory data, address proxyAdmin)
        private
        view
        returns (string memory)
    {
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

    function _generateContractsJson(DeploymentData memory data, address proxyAdmin)
        private
        view
        returns (string memory)
    {
        return string.concat(
            '{"proxyAdmin":"',
            proxyAdmin.toHexString(),
            '","incredibleSquaringServiceManager":"',
            data.incredibleSquaringServiceManager.toHexString(),
            '","incredibleSquaringServiceManagerImpl":"',
            data.incredibleSquaringServiceManager.getImplementation().toHexString(),
            '","incredibleSquaringTaskManager":"',
            data.incredibleSquaringTaskManager.toHexString(),
            '","registryCoordinator":"',
            data.registryCoordinator.toHexString(),
            '","blsapkRegistry":"',
            data.blsapkRegistry.toHexString(),
            '","indexRegistry":"',
            data.indexRegistry.toHexString(),
            '","stakeRegistry":"',
            data.stakeRegistry.toHexString(),
            '","operatorStateRetriever":"',
            data.operatorStateRetriever.toHexString(),
            '","strategy":"',
            data.strategy.toHexString(),
            '","pauserRegistry":"',
            data.pauserRegistry.toHexString(),
            '","token":"',
            data.token.toHexString(),
            '","instantSlasher":"',
            data.slasher.toHexString(),
            '"}'
        );
    }

    function readCoreDeploymentJson(string memory directoryPath, uint256 chainId)
        internal
        returns (CoreDeploymentLib.DeploymentData memory)
    {
        return readCoreDeploymentJson(directoryPath, string.concat(vm.toString(chainId), ".json"));
    }

    function readCoreDeploymentJson(string memory path, string memory fileName)
        internal
        returns (CoreDeploymentLib.DeploymentData memory)
    {
        string memory pathToFile = string.concat(path, fileName);

        require(vm.exists(pathToFile), "Deployment file does not exist");

        string memory json = vm.readFile(pathToFile);

        CoreDeploymentLib.DeploymentData memory data;
        data.strategyFactory = json.readAddress(".addresses.strategyFactory");
        data.strategyManager = json.readAddress(".addresses.strategyManager");
        data.eigenPodManager = json.readAddress(".addresses.eigenPodManager");
        data.delegationManager = json.readAddress(".addresses.delegation");
        data.avsDirectory = json.readAddress(".addresses.avsDirectory");
        data.pauserRegistry = json.readAddress(".addresses.pauserRegistry");
        data.rewardsCoordinator = json.readAddress(".addresses.rewardsCoordinator");
        data.permissionController = json.readAddress(".addresses.permissionController");
        data.allocationManager = json.readAddress(".addresses.allocationManager");

        return data;
    }

    function verify_deployment(DeploymentData memory result) internal view {
        IBLSApkRegistry blsapkregistry = IRegistryCoordinator(result.registryCoordinator).blsApkRegistry();
        require(address(blsapkregistry) != address(0));
        IStakeRegistry stakeregistry = IRegistryCoordinator(result.registryCoordinator).stakeRegistry();
        require(address(stakeregistry) != address(0));
        IDelegationManager delegationmanager = IStakeRegistry(address(stakeregistry)).delegation();
        require(address(delegationmanager) != address(0));
    }
}
