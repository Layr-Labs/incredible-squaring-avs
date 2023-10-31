#!/bin/bash

set -ue
CHAIN_ID=$(cast chain-id --rpc-url $RPC_URL)

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
# At this point we are in contracts/script
cd "$parent_path"

cd ../lib/eigenlayer-middleware/lib/eigenlayer-contracts
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
cp script/output/M2_from_scratch_deployment_data.json script/output/M2_from_scratch_deployment_data.json.bak
# M2_Deploy_From_Scratch.s.sol prepends "script/testing/" to the configFile passed as input (M2_deploy_from_scratch.anvil.config.json)
# we deploy using the anvil config file even if we deploy to goerli b/c this is just for testing
forge script script/testing/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
mv script/output/M2_from_scratch_deployment_data.json ../../../../script/output/$CHAIN_ID/eigenlayer_deployment_output.json
mv script/output/M2_from_scratch_deployment_data.json.bak script/output/M2_from_scratch_deployment_data.json
