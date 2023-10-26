#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

# start an anvil instance in the background that has eigenlayer contracts deployed
anvil --load-state eigenlayer-and-shared-avs-contracts-deployed-anvil-state.json --dump-state avs-and-eigenlayer-deployed-anvil-state.json &
cd ../../contracts
forge script script/IncredibleSquaringDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -v
 # we also do this here to make sure the operator has funds to register with the eigenlayer contracts
cast send 0x860B6912C2d0337ef05bbC89b0C2CB6CbAEAB4A5 --value 10ether --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
# kill anvil to save its state
pkill anvil
