// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@eigenlayer/contracts/permissions/PauserRegistry.sol";

import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {
    IBLSApkRegistry,
    IIndexRegistry,
    IStakeRegistry,
    ISocketRegistry
} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import {SocketRegistry} from "eigenlayer-middleware/src/SocketRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import {
    IncredibleSquaringServiceManager,
    IServiceManager
} from "../src/IncredibleSquaringServiceManager.sol";
import {IncredibleSquaringTaskManager} from "../src/IncredibleSquaringTaskManager.sol";
import {IIncredibleSquaringTaskManager} from "../src/IIncredibleSquaringTaskManager.sol";
import "../src/MockERC20.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";

import {ContractsRegistry} from "../src/ContractsRegistry.sol";
import {IncredibleSquaringDeploymentLib} from "../script/utils/IncredibleSquaringDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";

import {FundOperator} from "./utils/FundOperator.sol";
// # To deploy and verify our contract
// forge script script/IncredibleSquaringDeployer.s.sol:IncredibleSquaringDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv

contract IncredibleSquaringDeployer is Script {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    address public AGGREGATOR_ADDR;
    address public TASK_GENERATOR_ADDR;
    address public CONTRACTS_REGISTRY_ADDR;
    address public OPERATOR_ADDR;
    address public OPERATOR_2_ADDR;
    ContractsRegistry contractsRegistry;

    StrategyBaseTVLLimits public erc20MockStrategy;

    address public rewardscoordinator;

    ProxyAdmin public incredibleSquaringProxyAdmin;
    PauserRegistry public incredibleSquaringPauserReg;

    regcoord.RegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    ISocketRegistry public socketRegistry;
    ISocketRegistry public socketRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    IncredibleSquaringServiceManager public incredibleSquaringServiceManager;
    IServiceManager public incredibleSquaringServiceManagerImplementation;

    IncredibleSquaringTaskManager public incredibleSquaringTaskManager;
    IIncredibleSquaringTaskManager public incredibleSquaringTaskManagerImplementation;
    CoreDeploymentLib.DeploymentData internal configData;
    IStrategy incredibleSquaringStrategy;
    address private deployer;
    MockERC20 public erc20Mock;
    IncredibleSquaringDeploymentLib.DeploymentData incrediblSquaringDeployment;

    using UpgradeableProxyLib for address;

    address proxyAdmin;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
    }

    function run() external {
        // Eigenlayer contracts
        vm.startBroadcast(deployer);
        IncredibleSquaringDeploymentLib.IncredibleSquaringSetupConfig memory isConfig =
        IncredibleSquaringDeploymentLib.readIncredibleSquaringConfigJson(
            "incredible_squaring_config"
        );
        configData = CoreDeploymentLib.readDeploymentJson("script/deployments/core/", block.chainid);

        erc20Mock = new MockERC20();
        console.log(address(erc20Mock));
        FundOperator.fund_operator(address(erc20Mock), isConfig.operator_addr, 15_000e18);
        console2.log("strategy_factory");
        console2.log(configData.strategyFactory);
        incredibleSquaringStrategy =
            IStrategy(StrategyFactory(configData.strategyFactory).deployNewStrategy(erc20Mock));
        rewardscoordinator = configData.rewardsCoordinator;

        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
        // IncredibleSquaringDeploymentLib.DeploymentParams memory params = IncredibleSquaringDeploymentLib.DeploymentParams(proxyAdmin,address(incredibleSquaringStrategy),isConfig,msg.sender);

        incrediblSquaringDeployment = IncredibleSquaringDeploymentLib.deployContracts(
            proxyAdmin, address(incredibleSquaringStrategy), isConfig, msg.sender
        );
        console.log("instantSlasher", incrediblSquaringDeployment.slasher);

        FundOperator.fund_operator(
            address(erc20Mock), incrediblSquaringDeployment.incredibleSquaringServiceManager, 1e18
        );
        incrediblSquaringDeployment.token = address(erc20Mock);

        IncredibleSquaringDeploymentLib.writeDeploymentJson(incrediblSquaringDeployment);

        vm.stopBroadcast();
    }
}
