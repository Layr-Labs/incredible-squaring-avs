#!/bin/bash

set -ue
CHAIN_ID=$(cast chain-id --rpc-url $RPC_URL)

# cd to the directory of this script so that this can be run from anywhere
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
# At this point we are in contracts/script
cd "$parent_path"

cd ..
forge script script/IncredibleSquaringDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -v
