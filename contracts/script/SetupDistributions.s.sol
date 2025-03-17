// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {MockERC20} from "src/MockERC20.sol";
import {IncredibleSquaringDeploymentLib} from "./utils/IncredibleSquaringDeploymentLib.sol";
import {CoreDeploymentLib} from "./utils/CoreDeploymentLib.sol";
import {SetupDistributionsLib} from "./utils/SetupDistributionsLib.sol";
import {IRewardsCoordinator} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {RewardsCoordinator} from "@eigenlayer/contracts/core/RewardsCoordinator.sol";
import {IStrategy} from "@eigenlayer/contracts/interfaces/IStrategy.sol";

import "forge-std/Test.sol";

contract SetupDistributions is Script, Test {
    struct PaymentInfo {
        address recipient;
        uint32 numPayments;
        uint32 amountPerPayment;
        uint32 duration;
        uint32 startTimestamp;
        uint32 endTimestamp;
        uint256 indexToProve;
    }

    address private deployer;
    CoreDeploymentLib.DeploymentData coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfig;

    IncredibleSquaringDeploymentLib.DeploymentData incredibleSquaringDeployment;
    IncredibleSquaringDeploymentLib.IncredibleSquaringSetupConfig incredibleSquaringConfig;

    RewardsCoordinator rewardsCoordinator;
    string internal constant paymentInfofilePath = "test/mockData/scratch/payment_info.json";
    string internal constant filePath = "test/mockData/scratch/payments.json";

    uint32 constant CALCULATION_INTERVAL_SECONDS = 1 days;
    uint256 constant NUM_TOKEN_EARNINGS = 1;
    //duration MUST be a multiple of CALCULATION_INTERVAL_SECONDS .
    //https://github.com/Layr-Labs/eigenlayer-contracts/blob/865e723a6b5c634cf45cce1817dec0ea95f0e03b/src/contracts/core/RewardsCoordinator.sol#L439
    uint32 constant DURATION = 172_800;
    uint32 constant REWARDS_END_TIMESTAMP_GAP = 1 days;
    uint256 constant NUM_EARNERS = 8;

    uint32 numPayments = 8;
    uint32 indexToProve = 0;
    uint32 amountPerPayment = 100;

    address recipient = address(1);
    IRewardsCoordinator.EarnerTreeMerkleLeaf[] public earnerLeaves;
    address[] public earners;
    uint32 startTimestamp;
    uint32 endTimestamp;
    uint256 cumumlativePaymentMultiplier;
    address nonceSender = 0x998abeb3E57409262aE5b751f60747921B33613E;

    address operator1 = address(1);
    address operator2 = address(2);

    function setUp() public {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");

        coreDeployment =
            CoreDeploymentLib.readDeploymentJson("script/deployments/core/", block.chainid);
        incredibleSquaringDeployment = IncredibleSquaringDeploymentLib.readDeploymentJson(
            "script/deployments/incredible-squaring/", block.chainid
        );

        rewardsCoordinator = RewardsCoordinator(coreDeployment.rewardsCoordinator);

        // TODO: Get the filePath from config
    }

    function runAVSRewards() external {
        vm.startBroadcast(deployer);

        _checkStartTimestamp();

        // sets a multiplier based on block number such that cumulativeEarnings increase accordingly for multiple runs of this script in the same session
        uint256 nonce = rewardsCoordinator.getDistributionRootsLength();
        amountPerPayment = uint32(amountPerPayment * (nonce + 1));

        createAVSRewardsSubmissions(uint256(numPayments), uint256(amountPerPayment), startTimestamp);

        vm.stopBroadcast();
        vm.startBroadcast(deployer);
        earners = _getEarners(deployer);
        submitPaymentRoot(earners, endTimestamp, numPayments, amountPerPayment);
        vm.stopBroadcast();
    }

    function runOperatorDirected() external {
        vm.startBroadcast(deployer);

        _checkStartTimestamp();

        // sets a multiplier based on block number such that cumulativeEarnings increase accordingly for multiple runs of this script in the same session
        uint256 nonce = rewardsCoordinator.getDistributionRootsLength();
        amountPerPayment = uint32(amountPerPayment * (nonce + 1));

        createOperatorDirectedAVSRewardsSubmissions(
            numPayments, amountPerPayment, startTimestamp, DURATION
        );
        vm.stopBroadcast();
        vm.startBroadcast(deployer);
        earners = _getEarners(deployer);
        submitPaymentRoot(earners, endTimestamp, numPayments, amountPerPayment);
        vm.stopBroadcast();
    }

    function _checkStartTimestamp() internal {
        // Go back 4 days
        uint256 targetStartTimestamp = block.timestamp - 4 days;
        // Start Timestamp must be a multiple of CALCULATION_INTERVAL_SECONDS
        uint32 diff = (uint32(targetStartTimestamp) % CALCULATION_INTERVAL_SECONDS);
        startTimestamp = uint32(targetStartTimestamp) - diff;

        endTimestamp = uint32(block.timestamp) - REWARDS_END_TIMESTAMP_GAP;
        emit log_named_uint("startTimestamp", startTimestamp);
        emit log_named_uint("endTimestamp", endTimestamp);
        emit log_named_uint("block.timestamp", block.timestamp);
        emit log_named_uint("MAX_RETROACTIVE_LENGTH", rewardsCoordinator.MAX_RETROACTIVE_LENGTH());
        if (endTimestamp > block.timestamp) {
            revert("RewardsEndTimestampNotElapsed.  Please wait to generate new payments.");
        }
    }

    function executeProcessClaim() public {
        uint256 nonce = rewardsCoordinator.getDistributionRootsLength();
        amountPerPayment = uint32(amountPerPayment * nonce);

        vm.startBroadcast(deployer);
        earnerLeaves = _getEarnerLeaves(
            _getEarners(deployer), amountPerPayment, incredibleSquaringDeployment.strategy
        );
        processClaim(
            filePath, indexToProve, recipient, earnerLeaves[indexToProve], amountPerPayment
        );
        vm.stopBroadcast();
    }

    function createAVSRewardsSubmissions(
        uint256 numPayments,
        uint256 amountPerPayment,
        uint32 startTimestamp
    ) public {
        MockERC20(incredibleSquaringDeployment.token).mint(deployer, amountPerPayment * numPayments);
        MockERC20(incredibleSquaringDeployment.token).increaseAllowance(
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            amountPerPayment * numPayments
        );
        uint32 duration = rewardsCoordinator.MAX_REWARDS_DURATION();

        SetupDistributionsLib.createAVSRewardsSubmissions(
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            incredibleSquaringDeployment.strategy,
            numPayments,
            amountPerPayment,
            duration,
            startTimestamp
        );
    }

    function createOperatorDirectedAVSRewardsSubmissions(
        uint256 numPayments,
        uint256 amountPerPayment,
        uint32 startTimestamp,
        uint32 duration
    ) public {
        MockERC20(incredibleSquaringDeployment.token).mint(deployer, amountPerPayment * numPayments);
        MockERC20(incredibleSquaringDeployment.token).increaseAllowance(
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            amountPerPayment * numPayments
        );
        address[] memory operators = new address[](2);
        operators[0] = operator1;
        operators[1] = operator2;

        SetupDistributionsLib.createOperatorDirectedAVSRewardsSubmissions(
            incredibleSquaringDeployment.incredibleSquaringServiceManager,
            operators,
            incredibleSquaringDeployment.strategy,
            numPayments,
            amountPerPayment,
            duration,
            startTimestamp
        );
    }

    function processClaim(
        string memory filePath,
        uint256 indexToProve,
        address recipient,
        IRewardsCoordinator.EarnerTreeMerkleLeaf memory earnerLeaf,
        uint32 amountPerPayment
    ) public {
        SetupDistributionsLib.processClaim(
            IRewardsCoordinator(coreDeployment.rewardsCoordinator),
            filePath,
            indexToProve,
            recipient,
            earnerLeaf,
            NUM_TOKEN_EARNINGS,
            incredibleSquaringDeployment.strategy,
            amountPerPayment
        );
    }

    function submitPaymentRoot(
        address[] memory earners,
        uint32 endTimestamp,
        uint32 numPayments,
        uint32 amountPerPayment
    ) public {
        emit log_named_uint("cumumlativePaymentMultiplier", cumumlativePaymentMultiplier);
        bytes32[] memory tokenLeaves = SetupDistributionsLib.createTokenLeaves(
            IRewardsCoordinator(coreDeployment.rewardsCoordinator),
            NUM_TOKEN_EARNINGS,
            amountPerPayment,
            incredibleSquaringDeployment.strategy
        );
        IRewardsCoordinator.EarnerTreeMerkleLeaf[] memory earnerLeaves =
            SetupDistributionsLib.createEarnerLeaves(earners, tokenLeaves);
        emit log_named_uint("Earner Leaves Length", earnerLeaves.length);
        emit log_named_uint("numPayments", numPayments);

        SetupDistributionsLib.submitRoot(
            IRewardsCoordinator(coreDeployment.rewardsCoordinator),
            tokenLeaves,
            earnerLeaves,
            incredibleSquaringDeployment.strategy,
            endTimestamp,
            numPayments,
            NUM_TOKEN_EARNINGS,
            filePath
        );
    }

    function _getEarnerLeaves(
        address[] memory earners,
        uint32 amountPerPayment,
        address strategy
    ) internal view returns (IRewardsCoordinator.EarnerTreeMerkleLeaf[] memory) {
        bytes32[] memory tokenLeaves = SetupDistributionsLib.createTokenLeaves(
            IRewardsCoordinator(coreDeployment.rewardsCoordinator),
            NUM_TOKEN_EARNINGS,
            amountPerPayment,
            strategy
        );

        IRewardsCoordinator.EarnerTreeMerkleLeaf[] memory earnerLeaves =
            SetupDistributionsLib.createEarnerLeaves(earners, tokenLeaves);

        return earnerLeaves;
    }

    function _getEarners(
        address deployer
    ) internal pure returns (address[] memory) {
        address[] memory earners = new address[](NUM_EARNERS);
        for (uint256 i = 0; i < earners.length; i++) {
            earners[i] = deployer;
        }
        return earners;
    }
}
