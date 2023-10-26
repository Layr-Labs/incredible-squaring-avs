#!/bin/bash

RPC_URL=http://localhost:8545
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
# At this point we are in tests/integration
cd "$parent_path"

# start an empty anvil chain in the background and dump its state to a json file upon exit
anvil --load-state eigenlayer-deployed-anvil-state.json --dump-state eigenlayer-and-shared-avs-contracts-deployed-anvil-state.json &

cd ../../contracts/lib/eigenlayer-middleware
forge script script/DeploySharedContracts.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast
mv script/output/31337/shared_contracts_deployment_data.json ../../script/output/31337/shared_avs_contracts_deployment_output.json

# # kill anvil to save its state
pkill anvil
