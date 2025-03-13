// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {
    IRewardsCoordinator,
    IRewardsCoordinatorTypes
} from "@eigenlayer/contracts/interfaces/IRewardsCoordinator.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategyManager.sol";
import {ECDSAServiceManagerBase} from
    "@eigenlayer-middleware/src/unaudited/ECDSAServiceManagerBase.sol";
import {Vm} from "forge-std/Vm.sol";

library SetupDistributionsLib {
    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    struct PaymentLeaves {
        bytes32[] leaves;
        bytes32[] tokenLeaves;
    }

    function createAVSRewardsSubmissions(
        address incSquaringServiceManager,
        address strategy,
        uint256 numPayments,
        uint256 amountPerPayment,
        uint32 duration,
        uint32 startTimestamp
    ) internal {
        IRewardsCoordinatorTypes.RewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.RewardsSubmission[](numPayments);
        for (uint256 i = 0; i < numPayments; i++) {
            IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategiesAndMultipliers =
                new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
            strategiesAndMultipliers[0] = IRewardsCoordinatorTypes.StrategyAndMultiplier({
                strategy: IStrategy(strategy),
                multiplier: 10_000
            });

            IRewardsCoordinatorTypes.RewardsSubmission memory rewardsSubmission =
            IRewardsCoordinatorTypes.RewardsSubmission({
                strategiesAndMultipliers: strategiesAndMultipliers,
                token: IStrategy(strategy).underlyingToken(),
                amount: amountPerPayment,
                startTimestamp: startTimestamp,
                duration: duration
            });

            rewardsSubmissions[i] = rewardsSubmission;
        }
        ECDSAServiceManagerBase(incSquaringServiceManager).createAVSRewardsSubmission(
            rewardsSubmissions
        );
    }

    function createOperatorDirectedAVSRewardsSubmissions(
        address incSquaringServiceManager,
        address[] memory operators,
        address strategy,
        uint256 numPayments,
        uint256 amountPerPayment,
        uint32 duration,
        uint32 startTimestamp
    ) internal {
        uint256 operatorRewardAmount = amountPerPayment / operators.length;

        IRewardsCoordinatorTypes.OperatorReward[] memory operatorRewards =
            new IRewardsCoordinatorTypes.OperatorReward[](operators.length);
        for (uint256 i = 0; i < operators.length; i++) {
            operatorRewards[i] = IRewardsCoordinatorTypes.OperatorReward({
                operator: operators[i],
                amount: operatorRewardAmount
            });
        }

        IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[] memory rewardsSubmissions =
            new IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[](numPayments);
        for (uint256 i = 0; i < numPayments; i++) {
            IRewardsCoordinatorTypes.StrategyAndMultiplier[] memory strategiesAndMultipliers =
                new IRewardsCoordinatorTypes.StrategyAndMultiplier[](1);
            strategiesAndMultipliers[0] = IRewardsCoordinatorTypes.StrategyAndMultiplier({
                strategy: IStrategy(strategy),
                multiplier: 10_000
            });

            IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission memory rewardsSubmission =
            IRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission({
                strategiesAndMultipliers: strategiesAndMultipliers,
                token: IStrategy(strategy).underlyingToken(),
                operatorRewards: operatorRewards,
                startTimestamp: startTimestamp,
                duration: duration,
                description: "test"
            });

            rewardsSubmissions[i] = rewardsSubmission;
        }
        ECDSAServiceManagerBase(incSquaringServiceManager)
            .createOperatorDirectedAVSRewardsSubmission(rewardsSubmissions);
    }

    function processClaim(
        IRewardsCoordinator rewardsCoordinator,
        string memory filePath,
        uint256 indexToProve,
        address recipient,
        IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf memory earnerLeaf,
        uint256 NUM_TOKEN_EARNINGS,
        address strategy,
        uint32 amountPerPayment
    ) internal {
        PaymentLeaves memory paymentLeaves = parseLeavesFromJson(filePath);

        bytes memory proof = generateMerkleProof(paymentLeaves.leaves, indexToProve);
        //we only have one token leaf
        bytes memory tokenProof = generateMerkleProof(paymentLeaves.tokenLeaves, indexToProve);

        uint32[] memory tokenIndices = new uint32[](NUM_TOKEN_EARNINGS);
        bytes[] memory tokenProofs = new bytes[](NUM_TOKEN_EARNINGS);
        tokenProofs[0] = tokenProof;

        IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[] memory tokenLeaves =
            new IRewardsCoordinatorTypes.TokenTreeMerkleLeaf[](NUM_TOKEN_EARNINGS);
        tokenLeaves[0] = defaultTokenLeaf(amountPerPayment, strategy);

        // this workflow assumes a new root submitted for every payment claimed.  So we get the latest rood index to process a claim for
        uint256 rootIndex = rewardsCoordinator.getDistributionRootsLength() - 1;

        IRewardsCoordinatorTypes.RewardsMerkleClaim memory claim = IRewardsCoordinatorTypes
            .RewardsMerkleClaim({
            rootIndex: uint32(rootIndex),
            earnerIndex: uint32(indexToProve),
            earnerTreeProof: proof,
            earnerLeaf: earnerLeaf,
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenProofs,
            tokenLeaves: tokenLeaves
        });

        rewardsCoordinator.processClaim(claim, recipient);
    }

    function submitRoot(
        IRewardsCoordinator rewardsCoordinator,
        bytes32[] memory tokenLeaves,
        IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf[] memory earnerLeaves,
        address strategy,
        uint32 rewardsCalculationEndTimestamp,
        uint256 NUM_PAYMENTS,
        uint256 NUM_TOKEN_EARNINGS,
        string memory filePath
    ) internal {
        bytes32 paymentRoot = createPaymentRoot(
            rewardsCoordinator,
            tokenLeaves,
            earnerLeaves,
            NUM_PAYMENTS,
            NUM_TOKEN_EARNINGS,
            filePath
        );
        rewardsCoordinator.submitRoot(paymentRoot, rewardsCalculationEndTimestamp);
    }

    function createPaymentRoot(
        IRewardsCoordinator rewardsCoordinator,
        bytes32[] memory tokenLeaves,
        IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf[] memory earnerLeaves,
        uint256 NUM_PAYMENTS,
        uint256 NUM_TOKEN_EARNINGS,
        string memory filePath
    ) internal returns (bytes32) {
        require(
            earnerLeaves.length == NUM_PAYMENTS, "Number of earners must match number of payments"
        );
        bytes32[] memory leaves = new bytes32[](NUM_PAYMENTS);

        require(
            tokenLeaves.length == NUM_TOKEN_EARNINGS,
            "Number of token leaves must match number of token earnings"
        );
        for (uint256 i = 0; i < NUM_PAYMENTS; i++) {
            leaves[i] = rewardsCoordinator.calculateEarnerLeafHash(earnerLeaves[i]);
        }

        writeLeavesToJson(leaves, tokenLeaves, filePath);
        return (merkleizeKeccak(leaves));
    }

    function createEarnerLeaves(
        address[] calldata earners,
        bytes32[] memory tokenLeaves
    ) public pure returns (IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf[] memory) {
        IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf[] memory leaves =
            new IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf[](earners.length);
        for (uint256 i = 0; i < earners.length; i++) {
            leaves[i] = IRewardsCoordinatorTypes.EarnerTreeMerkleLeaf({
                earner: earners[i],
                earnerTokenRoot: createTokenRoot(tokenLeaves)
            });
        }
        return leaves;
    }

    function createTokenRoot(
        bytes32[] memory tokenLeaves
    ) public pure returns (bytes32) {
        return merkleizeKeccak(tokenLeaves);
    }

    function createTokenLeaves(
        IRewardsCoordinator rewardsCoordinator,
        uint256 NUM_TOKEN_EARNINGS,
        uint256 TOKEN_EARNINGS,
        address strategy
    ) internal view returns (bytes32[] memory) {
        bytes32[] memory leaves = new bytes32[](NUM_TOKEN_EARNINGS);
        for (uint256 i = 0; i < NUM_TOKEN_EARNINGS; i++) {
            IRewardsCoordinatorTypes.TokenTreeMerkleLeaf memory leaf =
                defaultTokenLeaf(TOKEN_EARNINGS, strategy);
            leaves[i] = rewardsCoordinator.calculateTokenLeafHash(leaf);
        }
        return leaves;
    }

    function defaultTokenLeaf(
        uint256 TOKEN_EARNINGS,
        address strategy
    ) internal view returns (IRewardsCoordinatorTypes.TokenTreeMerkleLeaf memory) {
        IRewardsCoordinatorTypes.TokenTreeMerkleLeaf memory leaf = IRewardsCoordinatorTypes
            .TokenTreeMerkleLeaf({
            token: IStrategy(strategy).underlyingToken(),
            cumulativeEarnings: TOKEN_EARNINGS
        });
        return leaf;
    }

    function writeLeavesToJson(
        bytes32[] memory leaves,
        bytes32[] memory tokenLeaves,
        string memory filePath
    ) internal {
        string memory parent_object = "parent_object";
        vm.serializeBytes32(parent_object, "leaves", leaves);
        string memory finalJson = vm.serializeBytes32(parent_object, "tokenLeaves", tokenLeaves);
        vm.writeJson(finalJson, filePath);
    }

    function parseLeavesFromJson(
        string memory filePath
    ) internal view returns (PaymentLeaves memory) {
        string memory json = vm.readFile(filePath);
        bytes memory data = vm.parseJson(json);
        return abi.decode(data, (PaymentLeaves));
    }

    function generateMerkleProof(
        bytes32[] memory leaves,
        uint256 index
    ) internal pure returns (bytes memory) {
        require(leaves.length > 0, "Leaves array cannot be empty");
        require(index < leaves.length, "Index out of bounds");

        leaves = padLeaves(leaves);

        uint256 n = leaves.length;
        uint256 depth = 0;
        while ((1 << depth) < n) {
            depth++;
        }

        bytes32[] memory proof = new bytes32[](depth);
        uint256 proofIndex = 0;

        for (uint256 i = 0; i < depth; i++) {
            uint256 levelSize = (n + 1) / 2;
            uint256 siblingIndex = (index % 2 == 0) ? index + 1 : index - 1;

            if (siblingIndex < n) {
                proof[proofIndex] = leaves[siblingIndex];
                proofIndex++;
            }

            for (uint256 j = 0; j < levelSize; j++) {
                if (2 * j + 1 < n) {
                    leaves[j] = keccak256(abi.encodePacked(leaves[2 * j], leaves[2 * j + 1]));
                } else {
                    leaves[j] = leaves[2 * j];
                }
            }

            n = levelSize;
            index /= 2;
        }

        return abi.encodePacked(proof);
    }

    /**
     * @notice this function returns the merkle root of a tree created from a set of leaves using keccak256 as its hash function
     *  @param leaves the leaves of the merkle tree
     *  @return The computed Merkle root of the tree.
     *  @dev This pads to the next power of 2. very inefficient! just for POC
     */
    function merkleizeKeccak(
        bytes32[] memory leaves
    ) internal pure returns (bytes32) {
        // uint256 paddedLength = 2;
        // while(paddedLength < leaves.length) {
        //     paddedLength <<= 1;
        // }

        // bytes32[] memory paddedLeaves = new bytes32[](paddedLength);
        // for (uint256 i = 0; i < leaves.length; i++) {
        //     paddedLeaves[i] = leaves[i];
        // }
        leaves = padLeaves(leaves);

        //there are half as many nodes in the layer above the leaves
        uint256 numNodesInLayer = leaves.length / 2;
        //create a layer to store the internal nodes
        bytes32[] memory layer = new bytes32[](numNodesInLayer);
        //fill the layer with the pairwise hashes of the leaves
        for (uint256 i = 0; i < numNodesInLayer; i++) {
            layer[i] = keccak256(abi.encodePacked(leaves[2 * i], leaves[2 * i + 1]));
        }
        //the next layer above has half as many nodes
        numNodesInLayer /= 2;
        //while we haven't computed the root
        while (numNodesInLayer != 0) {
            //overwrite the first numNodesInLayer nodes in layer with the pairwise hashes of their children
            for (uint256 i = 0; i < numNodesInLayer; i++) {
                layer[i] = keccak256(abi.encodePacked(layer[2 * i], layer[2 * i + 1]));
            }
            //the next layer above has half as many nodes
            numNodesInLayer /= 2;
        }
        //the first node in the layer is the root
        return layer[0];
    }

    function padLeaves(
        bytes32[] memory leaves
    ) internal pure returns (bytes32[] memory) {
        uint256 paddedLength = 2;
        while (paddedLength < leaves.length) {
            paddedLength <<= 1;
        }

        bytes32[] memory paddedLeaves = new bytes32[](paddedLength);
        for (uint256 i = 0; i < leaves.length; i++) {
            paddedLeaves[i] = leaves[i];
        }
        return paddedLeaves;
    }
}
