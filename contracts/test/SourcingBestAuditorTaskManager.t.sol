// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/SourcingBestAuditorServiceManager.sol" as incsqsm;
import {SourcingBestAuditorTaskManager} from "../src/SourcingBestAuditorTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract SourcingBestAuditorTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.SourcingBestAuditorServiceManager sm;
    incsqsm.SourcingBestAuditorServiceManager smImplementation;
    SourcingBestAuditorTaskManager tm;
    SourcingBestAuditorTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new SourcingBestAuditorTaskManager(
            incsqsm.IRegistryCoordinator(address(registryCoordinator)),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = SourcingBestAuditorTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    // function createNewTask(
    //     bytes calldata auditJobSpecificationURI,
    //     uint256 budgetInUSDC,
    //     bytes calldata quorumNumbers,
    //     uint32 quorumThresholdPercentage,
    //     Bid[] calldata bids // Submitting the array of all the bids
    // ) external onlyTaskGenerator {

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        SourcingBestAuditorTaskManager.Bid[]
            memory bids = new SourcingBestAuditorTaskManager.Bid[](1);
        bids[0] = SourcingBestAuditorTaskManager.Bid({
            bidId: 1,
            zkp: "zkp_example",  //@TODO: @akshatamohanty
            bidPitchDocURI: "uri_example"
        });
        cheats.prank(generator, generator);
        tm.createNewTask("specificationURI", 100, quorumNumbers, 50, bids);
        assertEq(tm.latestTaskNum(), 1);
    }
}
