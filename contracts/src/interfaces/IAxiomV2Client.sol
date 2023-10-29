// SPDX-License-Identifier: MIT
// copied from https://github.com/axiom-crypto/autonomous-airdrop-example/tree/main
// because v2 contracts are still not released under a repo
pragma solidity ^0.8.9;

interface IAxiomV2Client {
    event AxiomV2Call(
        uint64 indexed sourceChainId,
        address callerAddr,
        bytes32 indexed querySchema,
        uint256 indexed queryId
    );

    function axiomV2QueryAddress() external view returns (address);

    function axiomV2Callback(
        uint64 sourceChainId,
        address callerAddr,
        bytes32 querySchema,
        uint256 queryId,
        bytes32[] calldata axiomResults,
        bytes calldata callbackExtraData
    ) external;
}
