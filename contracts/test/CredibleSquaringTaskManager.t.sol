// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/IncredibleSquaringServiceManager.sol" as incsqsm;
import "../src/IncredibleSquaringTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "forge-std/console.sol";

contract IncredibleSquaringTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.IncredibleSquaringServiceManager sm;
    incsqsm.IncredibleSquaringServiceManager smImplementation;
    IncredibleSquaringTaskManager tm;
    IncredibleSquaringTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    // axiom-related constants (hardcoded for goerli)
    address public constant AXIOM_V2_QUERY_GOERLI_ADDR =
        0x28CeE427fCD58e5EF1cE4C93F877b621E2Db66df;
    // TODO: update this query schema to the query schema given by https://repl.axiom.xyz/
    //       for the query we end up making.
    bytes32 public constant AXIOM_QUERY_SCHEMA = 0x0;
    uint64 public constant AXIOM_CHAIN_ID = 5;

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new IncredibleSquaringTaskManager(
            incsqsm.IBLSRegistryCoordinatorWithIndices(
                address(registryCoordinator)
            ),
            TASK_RESPONSE_WINDOW_BLOCK,
            AXIOM_V2_QUERY_GOERLI_ADDR
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = IncredibleSquaringTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        serviceManagerOwner,
                        aggregator,
                        generator,
                        AXIOM_CHAIN_ID,
                        AXIOM_QUERY_SCHEMA
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

    function test_axiomV2Callback() public {
        // first create a new task in the task manager
        uint256 numberToBeSquared = 2;
        uint32 quorumThresholdPercentage = 100;
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator);
        tm.createNewTask(
            numberToBeSquared,
            quorumThresholdPercentage,
            quorumNumbers
        );
        IIncredibleSquaringTaskManager.Task
            memory task = IIncredibleSquaringTaskManager.Task(
                numberToBeSquared,
                // block number should be the one from when tm.createNewTask was created
                // that call doesn't seem to create a new block so we're fine doing this.
                uint32(block.number),
                quorumNumbers,
                quorumThresholdPercentage
            );

        // then create the taskResponse that operators would sign off on
        IIncredibleSquaringTaskManager.TaskResponse
            memory taskResponse = IIncredibleSquaringTaskManager.TaskResponse(
                0,
                numberToBeSquared * numberToBeSquared
            );

        // then we create the axiomResults and extraData arguments that would constructed by the axiom prover
        // and be passed to the callback
        uint256 queryId = 0;
        // not adding nonsigners or checking the quorumThresholdPercentage here b/c
        // those are checked in the zk proof in the axiom contract
        bytes32[] memory axiomResults = new bytes32[](1);
        axiomResults[0] = keccak256(abi.encode(taskResponse));
        bytes memory extraData = abi.encode(task, taskResponse);
        cheats.prank(AXIOM_V2_QUERY_GOERLI_ADDR);
        tm.axiomV2Callback(
            AXIOM_CHAIN_ID,
            AXIOM_V2_QUERY_GOERLI_ADDR,
            AXIOM_QUERY_SCHEMA,
            queryId,
            axiomResults,
            extraData
        );
        // TODO: should we check that taskResponseDigests was updated?
    }
}
