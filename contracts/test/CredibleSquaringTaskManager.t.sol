// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/CredibleSquaringServiceManager.sol" as incsqsm;
import {CredibleSquaringTaskManager} from "../src/CredibleSquaringTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract CredibleSquaringTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.CredibleSquaringServiceManager sm;
    incsqsm.CredibleSquaringServiceManager smImplementation;
    CredibleSquaringTaskManager tm;
    CredibleSquaringTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new CredibleSquaringTaskManager(
            incsqsm.IBLSRegistryCoordinatorWithIndices(
                address(registryCoordinator)
            ),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = CredibleSquaringTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        serviceManagerOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        tm.createNewTask(2, 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
