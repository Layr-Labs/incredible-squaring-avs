// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.27;

import "forge-std/Test.sol";
import "../script/utils/SetupPaymentsLib.sol";
import "../script/utils/CoreDeploymentLib.sol";
import "../script/utils/IncredibleSquaringDeploymentLib.sol";
import "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import "@eigenlayer/contracts/interfaces/IStrategy.sol";
import "@eigenlayer/contracts/libraries/Merkle.sol";
import "../script/DeployEigenLayerCore.s.sol";
import "../script/IncredibleSquaringDeployer.s.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";
import {console2} from "forge-std/console2.sol";
import {StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";
import "@eigenlayer-middleware/src/interfaces/IECDSAStakeRegistry.sol";
import {IERC20, StrategyFactory} from "@eigenlayer/contracts/strategies/StrategyFactory.sol";

contract TestConstants {
    uint256 constant NUM_PAYMENTS = 8;
    uint256 constant NUM_TOKEN_EARNINGS = 1;
    uint256 constant TOKEN_EARNINGS = 100;

    address RECIPIENT = address(1);
    address EARNER = address(2);
    uint256 INDEX_TO_PROVE = 0;
    uint256 NUM_EARNERS = 4;
}

contract SetupPaymentsLibTest is Test, TestConstants {
    using SetupPaymentsLib for *;

    Vm cheats = Vm(VM_ADDRESS);
    IncredibleSquaringDeploymentLib.DeploymentData internal incredibleSquaringDeployment;
    CoreDeploymentLib.DeploymentData internal coreDeployment;
    CoreDeploymentLib.DeploymentConfigData coreConfigData;
    IRewardsCoordinator public rewardsCoordinator;
    IStrategy public strategy;
    address proxyAdmin;
    address public deployer;
    MockERC20 public mockToken;
    IncredibleSquaringDeploymentLib.IncredibleSquaringSetupConfig iSquaringConfig;
    mapping(address => IStrategy) public tokenToStrategy;

    string internal constant filePath = "test/mockData/scratch/payments.json";

    function setUp() public virtual {
        deployer = vm.rememberKey(vm.envUint("PRIVATE_KEY"));
        vm.label(deployer, "Deployer");
        proxyAdmin = UpgradeableProxyLib.deployProxyAdmin();
        mockToken = new MockERC20();

        coreConfigData =
            CoreDeploymentLib.readDeploymentConfigValues("test/mockData/config/core/", 1337);
        coreDeployment = CoreDeploymentLib.deployContracts(deployer, proxyAdmin, coreConfigData);
        rewardsCoordinator = IRewardsCoordinator(coreDeployment.rewardsCoordinator);

        strategy = addStrategy(address(mockToken));

        mockToken.mint(address(this), 100_000);
        mockToken.mint(address(rewardsCoordinator), 100_000);
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

    function testSubmitRoot() public {
        address[] memory earners = new address[](NUM_EARNERS);
        for (uint256 i = 0; i < earners.length; i++) {
            earners[i] = address(1);
        }
        uint32 endTimestamp = rewardsCoordinator.currRewardsCalculationEndTimestamp() + 1 weeks;
        cheats.warp(endTimestamp + 1);

        bytes32[] memory tokenLeaves = SetupPaymentsLib.createTokenLeaves(
            rewardsCoordinator, NUM_TOKEN_EARNINGS, TOKEN_EARNINGS, address(strategy)
        );
        IRewardsCoordinator.EarnerTreeMerkleLeaf[] memory earnerLeaves =
            SetupPaymentsLib.createEarnerLeaves(earners, tokenLeaves);

        cheats.startPrank(deployer);
        SetupPaymentsLib.submitRoot(
            rewardsCoordinator,
            tokenLeaves,
            earnerLeaves,
            address(strategy),
            endTimestamp,
            NUM_EARNERS,
            1,
            filePath
        );
        cheats.stopPrank();
    }

    function testWriteLeavesToJson() public {
        bytes32[] memory leaves = new bytes32[](2);
        leaves[0] = bytes32(uint256(1));
        leaves[1] = bytes32(uint256(2));

        bytes32[] memory tokenLeaves = new bytes32[](2);
        tokenLeaves[0] = bytes32(uint256(3));
        tokenLeaves[1] = bytes32(uint256(4));

        string memory filePath = ("payments.json");

        SetupPaymentsLib.writeLeavesToJson(leaves, tokenLeaves, filePath);

        assertTrue(vm.exists("payments.json"), "JSON file should be created");
    }

    function testParseLeavesFromJson() public {
        string memory filePath = "test_parse_payments.json";
        string memory jsonContent = '{"leaves":["0x1234"], "tokenLeaves":["0x5678"]}';
        vm.writeFile(filePath, jsonContent);

        SetupPaymentsLib.PaymentLeaves memory paymentLeaves =
            SetupPaymentsLib.parseLeavesFromJson(filePath);

        assertEq(paymentLeaves.leaves.length, 1, "Incorrect number of leaves");
        assertEq(paymentLeaves.tokenLeaves.length, 1, "Incorrect number of token leaves");
    }

    function testGenerateMerkleProof() public {
        SetupPaymentsLib.PaymentLeaves memory paymentLeaves =
            SetupPaymentsLib.parseLeavesFromJson("test/mockData/scratch/payments.json");

        bytes32[] memory leaves = paymentLeaves.leaves;
        uint256 indexToProve = 0;

        bytes32[] memory proof = new bytes32[](2);
        proof[0] = leaves[1];
        proof[1] = keccak256(abi.encodePacked(leaves[2], leaves[3]));

        bytes memory proofBytesConstructed = abi.encodePacked(proof);
        bytes memory proofBytesCalculated =
            SetupPaymentsLib.generateMerkleProof(leaves, indexToProve);

        require(
            keccak256(proofBytesConstructed) == keccak256(proofBytesCalculated),
            "Proofs do not match"
        );

        bytes32 root = SetupPaymentsLib.merkleizeKeccak(leaves);

        require(
            Merkle.verifyInclusionKeccak(
                proofBytesCalculated, root, leaves[indexToProve], indexToProve
            )
        );
    }

    function testProcessClaim() public {
        emit log_named_address("token address", address(mockToken));
        string memory filePath = "test/mockData/scratch/payments.json";

        address[] memory earners = new address[](NUM_EARNERS);
        for (uint256 i = 0; i < earners.length; i++) {
            earners[i] = address(1);
        }
        uint32 endTimestamp = rewardsCoordinator.currRewardsCalculationEndTimestamp() + 1 weeks;
        cheats.warp(endTimestamp + 1);
        bytes32[] memory tokenLeaves = SetupPaymentsLib.createTokenLeaves(
            rewardsCoordinator, NUM_TOKEN_EARNINGS, TOKEN_EARNINGS, address(strategy)
        );
        IRewardsCoordinator.EarnerTreeMerkleLeaf[] memory earnerLeaves =
            SetupPaymentsLib.createEarnerLeaves(earners, tokenLeaves);

        console2.log(rewardsCoordinator.rewardsUpdater());
        cheats.startPrank(deployer);
        SetupPaymentsLib.submitRoot(
            rewardsCoordinator,
            tokenLeaves,
            earnerLeaves,
            address(strategy),
            endTimestamp,
            NUM_EARNERS,
            1,
            filePath
        );
        cheats.stopPrank();

        cheats.warp(block.timestamp + 2 weeks);

        cheats.startPrank(earnerLeaves[INDEX_TO_PROVE].earner, earnerLeaves[INDEX_TO_PROVE].earner);
        SetupPaymentsLib.processClaim(
            rewardsCoordinator,
            filePath,
            INDEX_TO_PROVE,
            RECIPIENT,
            earnerLeaves[INDEX_TO_PROVE],
            NUM_TOKEN_EARNINGS,
            address(strategy)
        );
        cheats.stopPrank();
    }

    function testCreateAVSRewardsSubmissions() public {
        uint256 numPayments = 5;
        uint256 amountPerPayment = 100;
        uint32 duration = rewardsCoordinator.MAX_REWARDS_DURATION();
        uint32 startTimestamp = 10 days;
        cheats.warp(startTimestamp + 1);
        mockToken.approve(address(rewardsCoordinator), amountPerPayment * numPayments);

        SetupPaymentsLib.createAVSRewardsSubmissions(
            rewardsCoordinator,
            address(strategy),
            numPayments,
            amountPerPayment,
            duration,
            startTimestamp
        );
    }
}
