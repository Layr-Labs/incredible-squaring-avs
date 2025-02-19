#!/bin/bash

# Default anvil URL
RPC_URL=http://localhost:8545
# Default anvil private key 0
PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# cd to the directory of this script so that this can be run from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
# At this point we are in tests/anvil
cd "$parent_path"

set -a
source ./utils.sh
set +a

cleanup() {
    echo "Executing cleanup function..."
    set +e
    docker rm -f anvil
    exit_status=$?
    if [ $exit_status -ne 0 ]; then
        echo "Script exited due to set -e on line $1 with command '$2'. Exit status: $exit_status"
    fi
}
trap 'cleanup $LINENO "$BASH_COMMAND"' EXIT

# start an empty anvil chain in the background and dump its state to a json file upon exit
start_anvil_docker "" $parent_path/eigenlayer-deployed-anvil-state.json

cd $parent_path/../../contracts

forge build

# Deploy the token and save the address
token_address=$(forge create --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
    --broadcast src/ERC20Mock.sol:ERC20Mock \
    | awk '/Deployed to: .*/{{ print $3 }}' | tr -d '\"\n')

cd lib/eigenlayer-middleware/lib/eigenlayer-contracts
# deployment overwrites this file, so we save it as backup, because we want that output in our local files, and not in the eigenlayer-contracts submodule files
mv script/output/devnet/M2_from_scratch_deployment_data.json script/output/devnet/M2_from_scratch_deployment_data.json.bak

# Add a strategy using the deployed token to the config file
strategy_info="{\"max_deposits\":115792089237316195423570985008687907853269984665640564039457584007913129639935,\"max_per_deposit\":115792089237316195423570985008687907853269984665640564039457584007913129639935,\"token_address\":\"$token_address\",\"token_symbol\":\"MockETH\"}"
cp script/configs/devnet/M2_deploy_from_scratch.anvil.config.json script/configs/devnet/M2_deploy_from_scratch.anvil.config.json.bak
sed -i '' "s#\"strategies\": \[\],#\"strategies\": \[$strategy_info\],#g" script/configs/devnet/M2_deploy_from_scratch.anvil.config.json

# M2_Deploy_From_Scratch.s.sol prepends "script/testing/" to the configFile passed as input (M2_deploy_from_scratch.anvil.config.json)
forge script script/deploy/devnet/M2_Deploy_From_Scratch.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY \
    --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json

mv script/output/devnet/M2_from_scratch_deployment_data.json ../../../../script/output/31337/eigenlayer_deployment_output.json
mv script/output/devnet/M2_from_scratch_deployment_data.json.bak script/output/devnet/M2_from_scratch_deployment_data.json
mv script/configs/devnet/M2_deploy_from_scratch.anvil.config.json.bak script/configs/devnet/M2_deploy_from_scratch.anvil.config.json
