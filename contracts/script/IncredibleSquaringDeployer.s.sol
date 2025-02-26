// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import "@eigenlayer/contracts/core/AllocationManager.sol";
import "@eigenlayer/contracts/permissions/PermissionController.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import "@eigenlayer/contracts/interfaces/IPauserRegistry.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";

import "@eigenlayer-middleware/test/utils/MiddlewareDeployLib.sol";

import {ISlashingRegistryCoordinator, ISlashingRegistryCoordinatorTypes} from "@eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";

import {IStakeRegistry, IStakeRegistryTypes} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";

import {
    IBLSApkRegistry,
    IIndexRegistry,
    IStakeRegistry
} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {
    IncredibleSquaringServiceManager,
    IServiceManager
} from "../src/IncredibleSquaringServiceManager.sol";
import {IncredibleSquaringTaskManager} from "../src/IncredibleSquaringTaskManager.sol";
import {IIncredibleSquaringTaskManager} from "../src/IIncredibleSquaringTaskManager.sol";
import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";

// # To deploy and verify our contract
// forge script script/IncredibleSquaringDeployer.s.sol:IncredibleSquaringDeployer --rpc-url $RPC_URL  --private-key $PRIVATE_KEY --broadcast -vvvv
contract IncredibleSquaringDeployer is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
    address public constant TASK_GENERATOR_ADDR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    // ERC20 and Strategy: we need to deploy this ERC20, create a strategy for it, and whitelist this strategy in the strategymanager

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

    // Incredible Squaring contracts
    ProxyAdmin public incredibleSquaringProxyAdmin;
    PauserRegistry public incredibleSquaringPauserReg;

    SlashingRegistryCoordinator public registryCoordinator;
    ISlashingRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    IncredibleSquaringServiceManager public incredibleSquaringServiceManager;
    IServiceManager public incredibleSquaringServiceManagerImplementation;

    IncredibleSquaringTaskManager public incredibleSquaringTaskManager;
    IIncredibleSquaringTaskManager public incredibleSquaringTaskManagerImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput("eigenlayer_deployment_output");
        IDelegationManager delegationManager = IDelegationManager(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.delegation")
        );
        IAVSDirectory avsDirectory = IAVSDirectory(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.avsDirectory")
        );
        IRewardsCoordinator rewardsCoordinator = IRewardsCoordinator(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.rewardsCoordinator")
        );

        address incredibleSquaringCommunityMultisig = msg.sender;
        address incredibleSquaringPauser = msg.sender;

        erc20MockStrategy = StrategyBaseTVLLimits(
            stdJson.readAddress(eigenlayerDeployedContracts, ".addresses.strategies.MockETH")
        );
        erc20Mock = ERC20Mock(address(erc20MockStrategy.underlyingToken()));

        vm.startBroadcast();
        _deployIncredibleSquaringContracts(
            delegationManager,
            avsDirectory,
            rewardsCoordinator,
            //erc20MockStrategy,
            incredibleSquaringCommunityMultisig,
            incredibleSquaringPauser
        );
        vm.stopBroadcast();
    }

    function _deployIncredibleSquaringContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IRewardsCoordinator rewardsCoordinator,
//        IStrategy strat,      //Uncomment when use to set quorum params
        address incredibleSquaringCommunityMultisig,
        address incredibleSquaringPauser
    ) internal {
        // deploy proxy admin for ability to upgrade proxy contracts
        incredibleSquaringProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = incredibleSquaringPauser;
            pausers[1] = incredibleSquaringCommunityMultisig;
            incredibleSquaringPauserReg =
                new PauserRegistry(pausers, incredibleSquaringCommunityMultisig);
        }

        EmptyContract emptyContract = new EmptyContract();

        // This is a proxy contract that will point to the implementation, but SM contract is not deployed yet
        incredibleSquaringServiceManager = IncredibleSquaringServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract), address(incredibleSquaringProxyAdmin), ""
                )
            )
        );

        // Maybe we should receive this as parameter
        PermissionController permissionController = new PermissionController();

        AllocationManager allocationManager = new AllocationManager( 
            delegationManager,
            incredibleSquaringPauserReg,
            permissionController,
            uint32(0),
            uint32(0)
            );

        MiddlewareDeployLib.InstantSlasherConfig memory instantSlasherConfig = MiddlewareDeployLib.InstantSlasherConfig({ 
            initialOwner: incredibleSquaringCommunityMultisig,
            slasher: incredibleSquaringCommunityMultisig // The address authorized to send slashing requests.
            });

        MiddlewareDeployLib.SlashingRegistryCoordinatorConfig memory slashingRegConfig = MiddlewareDeployLib.SlashingRegistryCoordinatorConfig({ 
            // we set churnApprover and ejector to communityMultisig because we don't need them
            initialOwner: incredibleSquaringCommunityMultisig,
            churnApprover: incredibleSquaringCommunityMultisig,
            ejector: incredibleSquaringCommunityMultisig,
            initPausedStatus: 0, // 0 initialPausedStatus means everything unpaused
            serviceManager: address(incredibleSquaringServiceManager)
            });

        MiddlewareDeployLib.SocketRegistryConfig memory socketRegConfig = MiddlewareDeployLib.SocketRegistryConfig({ 
            initialOwner: incredibleSquaringCommunityMultisig
            });

        MiddlewareDeployLib.IndexRegistryConfig memory indexRegConfig = MiddlewareDeployLib.IndexRegistryConfig({ 
            initialOwner: incredibleSquaringCommunityMultisig
            });
        
        // Maybe we could add some strategy params
        IStakeRegistryTypes.StrategyParams[] memory strategyParams;
        MiddlewareDeployLib.StakeRegistryConfig memory stakeRegConfig = MiddlewareDeployLib.StakeRegistryConfig({ 
            initialOwner: incredibleSquaringCommunityMultisig,
            minimumStake: uint256(10),
            strategyParams: uint32(0),
            delegationManager: address(delegationManager),
            avsDirectory: address(avsDirectory),
            strategyParamsArray: strategyParams, 
            lookAheadPeriod: uint32(0),
            stakeType: IStakeRegistryTypes.StakeType.TOTAL_SLASHABLE
            });

        MiddlewareDeployLib.BLSApkRegistryConfig memory blsConfig = MiddlewareDeployLib.BLSApkRegistryConfig({ 
            initialOwner: incredibleSquaringCommunityMultisig
            });

        MiddlewareDeployLib.MiddlewareDeployConfig memory midDeployConfig = MiddlewareDeployLib.MiddlewareDeployConfig({
            instantSlasher: instantSlasherConfig,
            slashingRegistryCoordinator: slashingRegConfig,
            socketRegistry: socketRegConfig,
            indexRegistry: indexRegConfig,
            stakeRegistry: stakeRegConfig,
            blsApkRegistry: blsConfig
        });
        MiddlewareDeployLib.MiddlewareDeployData memory deployData = MiddlewareDeployLib.deployMiddleware(address(incredibleSquaringProxyAdmin), address(allocationManager), address(incredibleSquaringPauserReg), midDeployConfig);

        indexRegistry = IndexRegistry(deployData.indexRegistry);
        stakeRegistry = StakeRegistry(deployData.stakeRegistry);
        blsApkRegistry = BLSApkRegistry(deployData.blsApkRegistry);
        registryCoordinator = SlashingRegistryCoordinator(deployData.slashingRegistryCoordinator);
        // instantSlasher;
        // socketRegistry;

        /*
            This parameters should be used with some slashing registry coordinator method:

            // Adding this as a temporary fix to make the rest of the script work with a single strategy
            // since it was originally written to work with an array of strategies
            IStrategy[1] memory deployedStrategyArray = [strat];
            uint256 numStrategies = deployedStrategyArray.length;

            uint256 numQuorums = 1;
            // for each quorum to set up, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            ISlashingRegistryCoordinatorTypes.OperatorSetParam[] memory quorumsOperatorSetParams =
                new ISlashingRegistryCoordinatorTypes.OperatorSetParam[](numQuorums);
            for (uint256 i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = ISlashingRegistryCoordinatorTypes.OperatorSetParam({
                    maxOperatorCount: 10_000,
                    kickBIPsOfOperatorStake: 15_000,
                    kickBIPsOfTotalStake: 100
                });
            }
            // set to 0 for every quorum
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
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
         */

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        incredibleSquaringTaskManager = IncredibleSquaringTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract), address(incredibleSquaringProxyAdmin), ""
                )
            )
        );

        // See if it's being deployed in another place
        operatorStateRetriever = new OperatorStateRetriever();

        incredibleSquaringServiceManagerImplementation = new IncredibleSquaringServiceManager(
            avsDirectory,
            rewardsCoordinator,
            registryCoordinator,
            stakeRegistry,
            incredibleSquaringTaskManager,
            permissionController,
            allocationManager
        );
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        incredibleSquaringProxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(incredibleSquaringServiceManager))),
            address(incredibleSquaringServiceManagerImplementation)
        );

        incredibleSquaringTaskManagerImplementation =
            new IncredibleSquaringTaskManager(registryCoordinator, incredibleSquaringPauserReg, TASK_RESPONSE_WINDOW_BLOCK);

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        incredibleSquaringProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(incredibleSquaringTaskManager))),
            address(incredibleSquaringTaskManagerImplementation),
            abi.encodeWithSelector(
                incredibleSquaringTaskManager.initialize.selector,
                incredibleSquaringCommunityMultisig,
                AGGREGATOR_ADDR,
                TASK_GENERATOR_ADDR
            )
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(deployed_addresses, "erc20Mock", address(erc20Mock));
        vm.serializeAddress(deployed_addresses, "erc20MockStrategy", address(erc20MockStrategy));
        vm.serializeAddress(
            deployed_addresses,
            "credibleSquaringServiceManager",
            address(incredibleSquaringServiceManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "credibleSquaringServiceManagerImplementation",
            address(incredibleSquaringServiceManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "credibleSquaringTaskManager",
            address(incredibleSquaringTaskManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "credibleSquaringTaskManagerImplementation",
            address(incredibleSquaringTaskManagerImplementation)
        );
        vm.serializeAddress(deployed_addresses, "registryCoordinator", address(registryCoordinator));
        // vm.serializeAddress(
        //     deployed_addresses,
        //     "registryCoordinatorImplementation",
        //     address(registryCoordinatorImplementation)
        // );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses, "operatorStateRetriever", address(operatorStateRetriever)
        );

        // New release contracts
        vm.serializeAddress(
            deployed_addresses,
            "indexRegistry",
            address(indexRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "stakeRegistry",
            address(stakeRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "blsApkRegistry",
            address(blsApkRegistry)
        );
        vm.serializeAddress(
            deployed_addresses,
            "allocationManager",
            address(allocationManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "permissionController",
            address(permissionController)
        );

        // serialize all the data
        string memory finalJson =
            vm.serializeString(parent_object, deployed_addresses, deployed_addresses_output);

        writeOutput(finalJson, "credible_squaring_avs_deployment_output");
    }
}
