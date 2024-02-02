#!/bin/bash

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
# At this point we are in tests/integration
cd "$parent_path"

# start an empty anvil chain in the background and dump its state to a json file upon exit
anvil --dump-state eigenlayer-deployed-anvil-state.json &

cd ../../contracts/lib/eigenlayer-middleware/lib/eigenlayer-contracts
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
mv script/output/M2_from_scratch_deployment_data.json script/output/M2_from_scratch_deployment_data.json.bak
# M2_Deploy_From_Scratch.s.sol prepends "script/testing/" to the configFile passed as input (M2_deploy_from_scratch.anvil.config.json)
forge script script/testing/M2_Deploy_From_Scratch.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
mv script/output/M2_from_scratch_deployment_data.json ../../../../script/output/31337/eigenlayer_deployment_output.json
mv script/output/M2_from_scratch_deployment_data.json.bak script/output/M2_from_scratch_deployment_data.json

# # kill anvil to save its state
pkill anvil

# Anvil adds a block state, making the code to fail. We don't care about this, just the accounts and the deployed code
cd "$parent_path"

jq 'del(.block)' eigenlayer-deployed-anvil-state.json > eigenlayer-deployed-anvil-state-tmp.json

cp -f eigenlayer-deployed-anvil-state-tmp.json eigenlayer-deployed-anvil-state.json

rm eigenlayer-deployed-anvil-state-tmp.json
