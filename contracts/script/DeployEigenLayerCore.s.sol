// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {UpgradeableProxyLib} from "./utils/UpgradeableProxyLib.sol";

contract DeployEigenLayerCore is Script {
    using CoreDeploymentLib for *;
    using UpgradeableProxyLib for address;

    address internal deployer;
    address internal proxyAdmin;
    CoreDeploymentLib.DeploymentData internal deploymentData;
    CoreDeploymentLib.DeploymentConfigData internal configData;

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
    }

    function run() external {
        configData = CoreDeploymentLib.readDeploymentConfigValues("config/core/", block.chainid);

        vm.startBroadcast(deployer);

        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
        deploymentData = CoreDeploymentLib.deployContracts(deployer, proxyAdmin, configData);
        vm.stopBroadcast();
        string memory deploymentPath = "script/deployments/core/";
        CoreDeploymentLib.writeDeploymentJson(deploymentPath, block.chainid, deploymentData);
    }
}
