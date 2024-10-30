#!/bin/bash

RPC_URL=http://localhost:8545

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd "$parent_path"

set -a
source ./utils.sh
set +a


deploy_eigenlayer() {
    forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
    echo "deployment done"
    mv script/output/devnet/M2_from_scratch_deployment_data.json ../../../../script/output/31337/eigenlayer_deployment_output.json
    mv script/output/devnet/M2_from_scratch_deployment_data.json.bak script/output/devnet/M2_from_scratch_deployment_data.json
    echo "deployment output moved"
}

deploy_avs() {
    echo "deploying avs"
    forge script ../../../../script/IncredibleSquaringDeployer.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -v
    echo "avs deployed"
}

start_anvil() {
    anvil
    sleep 2
}



cd ../../contracts/lib/eigenlayer-middleware/lib/eigenlayer-contracts
pwd
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
# mv script/output/devnet/M2_from_scratch_deployment_data.json script/output/devnet/M2_from_scratch_deployment_data.json.bak


start_anvil &
deploy_eigenlayer &
deploy_avs &

# cd ../../contracts
# # we need to restart the anvil chain at the correct block, otherwise the indexRegistry has a quorumUpdate at the block number
# # at which it was deployed (aka quorum was created/updated), but when we start anvil by loading state file it starts at block number 0
# # so calling getOperatorListAtBlockNumber reverts because it thinks there are no quorums registered at block 0
# # advancing chain manually like this is a current hack until https://github.com/foundry-rs/foundry/issues/6679 is merged
# cast rpc anvil_mine 100 --rpc-url $RPC_URL
# echo "advancing chain... current block-number:" $(cast block-number)

