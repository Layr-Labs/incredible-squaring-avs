// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {ContractsRegistry} from "../src/ContractsRegistry.sol";


contract DeployContractsRegistry is Script{

    address private deployer;


    function setUp() public {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

    }

    function run() external {
        vm.startBroadcast(deployer);

        bytes memory bytecode = type(ContractsRegistry).creationCode;
        address deployedAddress;
        assembly {
            deployedAddress := create(0, add(bytecode, 0x20), mload(bytecode))
        }

        vm.stopBroadcast();
    }


}