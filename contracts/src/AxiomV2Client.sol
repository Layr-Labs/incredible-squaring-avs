// SPDX-License-Identifier: MIT
// copied from https://github.com/axiom-crypto/autonomous-airdrop-example/blob/main/contracts/src/AxiomV2Client.sol
// because v2 contracts are still not released under a unified repo
pragma solidity ^0.8.9;

import "./interfaces/IAxiomV2Client.sol";

abstract contract AxiomV2Client is IAxiomV2Client {
    address public immutable axiomV2QueryAddress;

    constructor(address _axiomV2QueryAddress) {
        axiomV2QueryAddress = _axiomV2QueryAddress;
    }

    function axiomV2Callback(
        uint64 sourceChainId,
        address callerAddr,
        bytes32 querySchema,
        uint256 queryId,
        bytes32[] calldata axiomResults,
        bytes calldata callbackExtraData
    ) external {
        // zk proof is sent to AxiomV2Query contract, which calls this callback.
        // we make sure that the caller is indeed the AxiomV2Query contract.
        // see https://docs-v2.axiom.xyz/introduction/concepts/query#compute-query-axiomrepl
        require(msg.sender == axiomV2QueryAddress, "AxiomV2Client: caller must be axiomV2QueryAddress");
        emit AxiomV2Call(sourceChainId, callerAddr, querySchema, queryId);

        _validateAxiomV2Call(sourceChainId, callerAddr, querySchema);
        _axiomV2Callback(sourceChainId, callerAddr, querySchema, queryId, axiomResults, callbackExtraData);
    }

    function _validateAxiomV2Call(
        uint64 sourceChainId,
        address callerAddr,
        bytes32 querySchema
    ) internal virtual;

    function _axiomV2Callback(
        uint64 sourceChainId,
        address callerAddr,
        bytes32 querySchema,
        uint256 queryId,
        bytes32[] calldata axiomResults,
        bytes calldata callbackExtraData
    ) internal virtual;
}