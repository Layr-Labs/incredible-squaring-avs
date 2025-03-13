#!/bin/bash

# Enable the script to exit immediately if a command exits with a non-zero status
set -o errexit -o nounset -o pipefail

# Navigate to the script directory
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
cd "$parent_path"

root_dir=$(realpath "$parent_path/../..")

set -a
source $parent_path/utils.sh
# we overwrite some variables here because should always deploy to anvil (localhost)
RPC_URL=http://localhost:8545
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
set +a


# # start an empty anvil chain in the background and dump its state to a json file upon exit
start_anvil_docker "" $parent_path/avs-and-eigenlayer-deployed-anvil-state

# Deploy Contracts
cd "$root_dir/contracts"

forge script script/DeployEigenLayerCore.s.sol:DeployEigenLayerCore --rpc-url $RPC_URL --broadcast --private-key $PRIVATE_KEY 

forge script script/IncredibleSquaringDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast

forge script script/UAMPermissions.s.sol --rpc-url $RPC_URL --broadcast --slow --private-key $PRIVATE_KEY

forge script script/CreateQuorum.s.sol --rpc-url $RPC_URL --broadcast --slow --private-key $PRIVATE_KEY
