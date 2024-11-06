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


deploy_eigenlayer_and_avs() {
    forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
    echo "deployment done"
    mv script/output/devnet/M2_from_scratch_deployment_data.json ../../../../script/output/31337/eigenlayer_deployment_output.json
    mv script/output/devnet/M2_from_scratch_deployment_data.json.bak script/output/devnet/M2_from_scratch_deployment_data.json
    echo "deployment output moved"
    echo "deploying avs"
    cd ../../../../../contracts
    forge script script/IncredibleSquaringDeployer.s.sol:IncredibleSquaringDeployer --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast -v
    echo "avs deployed"
}



cd ../../contracts/lib/eigenlayer-middleware/lib/eigenlayer-contracts
pwd
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
mv script/output/devnet/M2_from_scratch_deployment_data.json script/output/devnet/M2_from_scratch_deployment_data.json.bak

# private key used for operations multisig which is the strategy whitelister
export OPERATIONS_MULTISIG_PK="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

deploy_eigenlayer_and_avs

