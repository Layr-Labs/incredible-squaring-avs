// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {IncredibleSquaringServiceManager} from "../src/IncredibleSquaringServiceManager.sol";
import {MockAVSDeployer} from "@eigenlayer-middleware/test/utils/MockAVSDeployer.sol";
import {ECDSAStakeRegistry} from "@eigenlayer-middleware/src/unaudited/ECDSAStakeRegistry.sol";
import {Vm} from "forge-std/Vm.sol";
import {console2} from "forge-std/Test.sol";
import {IncredibleSquaringDeploymentLib} from "../script/utils/IncredibleSquaringDeploymentLib.sol";
import {CoreDeploymentLib} from "../script/utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "../script/utils/UpgradeableProxyLib.sol";
import {MockERC20} from "../src/MockERC20.sol";
import {IERC20, StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";

import "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistry.sol";
import {IStrategyManager} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {ISignatureUtils} from "@eigenlayer/contracts/interfaces/ISignatureUtils.sol";
import {AVSDirectory} from "@eigenlayer/contracts/core/AVSDirectory.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {Test, console2 as console} from "forge-std/Test.sol";
import {IncredibleSquaringServiceManager} from "../src/IncredibleSquaringServiceManager.sol";

contract IncredibleSquaringServiceManagerSetup is Test {
    IECDSAStakeRegistryTypes.Quorum internal quorum;

    struct Operator {
        Vm.Wallet key;
        Vm.Wallet signingKey;
    }

    struct TrafficGenerator {
        Vm.Wallet key;
    }

    Operator[] internal operators;
    TrafficGenerator internal generator;

    IncredibleSquaringDeploymentLib.DeploymentData internal incredibleSquaringDeployment;
    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;
    IncredibleSquaringDeploymentLib.IncredibleSquaringSetupConfig iSquaringConfig;

    MockERC20 public mockToken;

    mapping(address => IStrategy) public tokenToStrategy;
    address public deployer;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
        Vm.Wallet memory AGGREGATOR_ADDR = vm.createWallet("AGGREGATOR_AGGR");
        Vm.Wallet memory TASK_GENERATOR_ADDR = vm.createWallet("TASK_GENERATOR_ADDR");
        Vm.Wallet memory ADMIN = vm.createWallet("ADMIN");
        address proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();

        coreConfigData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
        coreDeployment = CoreDeploymentLib.deployContracts(deployer, proxyAdmin, coreConfigData);
        iSquaringConfig = IncredibleSquaringDeploymentLib.readIncredibleSquaringConfigJson(
            "incredible_squaring_config"
        );
        mockToken = new MockERC20();

        IStrategy strategy = addStrategy(address(mockToken));
        quorum.strategies.push(
            IECDSAStakeRegistryTypes.StrategyParams({strategy: strategy, multiplier: 10_000})
        );

        incredibleSquaringDeployment = IncredibleSquaringDeploymentLib.deployContracts(
            proxyAdmin, coreDeployment, address(strategy), iSquaringConfig, ADMIN.addr
        );
        labelContracts(coreDeployment, incredibleSquaringDeployment);
    }

    function addStrategy(
        address token
    ) public returns (IStrategy) {
        if (tokenToStrategy[token] != IStrategy(address(0))) {
            return tokenToStrategy[token];
        }

        StrategyFactory strategyFactory = StrategyFactory(coreDeployment.strategyFactory);
        IStrategy newStrategy = strategyFactory.deployNewStrategy(IERC20(token));
        tokenToStrategy[token] = newStrategy;
        return newStrategy;
    }

    function labelContracts(
        CoreDeploymentLib.DeploymentData memory coreDeployment,
        IncredibleSquaringDeploymentLib.DeploymentData memory incredibleSquaringDeployment
    ) internal {
        vm.label(coreDeployment.delegationManager, "DelegationManager");
        vm.label(coreDeployment.avsDirectory, "AVSDirectory");
        vm.label(coreDeployment.strategyManager, "StrategyManager");
        vm.label(coreDeployment.eigenPodManager, "EigenPodManager");
        vm.label(coreDeployment.rewardsCoordinator, "RewardsCoordinator");
        vm.label(coreDeployment.eigenPodBeacon, "EigenPodBeacon");
        vm.label(coreDeployment.pauserRegistry, "PauserRegistry");
        vm.label(coreDeployment.strategyFactory, "StrategyFactory");
        vm.label(coreDeployment.strategyBeacon, "StrategyBeacon");
        vm.label(
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            "IncredibleSquaringServiceManager"
        );
        vm.label(incredibleSquaringDeployment.stakeRegistry, "StakeRegistry");
    }
}
