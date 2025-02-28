// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

// ContractsRegistry store the address of all the contracts deployed for eigenlayer and avss
// It is used for testing purpose only, so that we can retrieve the addresses without having to store them in a json file
// This way integration testing against an anvil chain (started with a saved db) is self-contained
// ASSUMPTION: right now we deploy this contract as the first deployment (nonce=0) using the first anvil address
// 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 to make sure it's always at the address 0x5FbDB2315678afecb367f032d93F642f64180aa3
// forge create src/ContractsRegistry.sol:ContractsRegistry --rpc-url $RPC_URL  --private-key $PRIVATE_KEY
contract ContractsRegistry {
    mapping(string => address) public contracts;
    mapping(uint256 => string) public contractNames;
    uint256 public contractCount;

    mapping(bytes32 => BlockAndTimestamp) public anvil_test;

    struct BlockAndTimestamp {
        uint256 timestamp;
        uint256 block_number;
        int256 index;
    }

    function registerContract(string memory name, address _contract) public {
        // we treat redeploys as a bug since this is only meant to be used for testing.
        // If new contracts need to be deployed just start from a fresh anvil state.
        require(contracts[name] == address(0), "contract already registered");
        contracts[name] = _contract;
        contractNames[contractCount] = name;
        contractCount++;
    }

    function store_test(
        string memory test_name,
        int256 index,
        uint256 timestamp,
        uint256 block_number
    ) public {
        require(anvil_test[keccak256(abi.encodePacked(test_name, index))].timestamp == 0);
        anvil_test[keccak256(abi.encodePacked(test_name))] =
            BlockAndTimestamp({timestamp: timestamp, block_number: block_number, index: index});
    }

    function get_test_values(
        string memory test_name,
        int256 index
    ) public view returns (uint256, uint256, int256) {
        BlockAndTimestamp memory test_details =
            anvil_test[keccak256(abi.encodePacked(test_name, index))];
        return (test_details.timestamp, test_details.block_number, test_details.index);
    }
}
