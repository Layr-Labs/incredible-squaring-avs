#!/bin/bash

set -e

RPC_URL=http://localhost:8545
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Navigate to the script directory
parent_path=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)
cd "$parent_path"

root_dir=$(realpath "$parent_path/../..")

# Deploy Contracts
cd "$root_dir/contracts"
forge script script/DeployEigenLayerCore.s.sol:DeployEigenLayerCore --rpc-url $RPC_URL --broadcast --private-key $PRIVATE_KEY 
