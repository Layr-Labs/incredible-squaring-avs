#!/bin/bash

# TODO: we could speed this up by making the block times configurable
#  then we can use lower block times in the devnet, and reduce both startup and sleep times

# This script receives the EL contract version to test against
# In case they aren't specified, we use a default one
# DEFAULT_REF=v0.3.3-mainnet-rewards
# EL_REF="${EL_REF:-$DEFAULT_REF}"
# EL_VERSION="${EL_VERSION:-$EL_REF}"

set -e -o nounset

# This is so the script can be called from anywhere
parent_path=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    pwd -P
)
cd $parent_path/..


# Stop the devnet in case it's already running
avs-devnet stop || true

cp devnet.yaml devnet.yaml.bak

# Update the devnet configuration with the EL contract versions to test against
# if [[ $EL_REF != $DEFAULT_REF ]] ; then
#     yq -i ".deployments.0.ref = \"$EL_REF\"" devnet.yaml
#     yq -i ".deployments.0.version = \"$EL_VERSION\"" devnet.yaml
# fi

# Stop the devnet after we finish
cleanup() {
    echo "Executing cleanup function..."
    set +e
    # Stop the devnet, and restore the old devnet configuration
    avs-devnet stop || true
    mv devnet.yaml.bak devnet.yaml || true
}
trap 'cleanup' EXIT

# Start the devnet and exit with a custom message if it fails
avs-devnet start || { echo "Failed to start the devnet" ; exit 42; }

# Fetch RPC URL and TaskManager address
RPC_URL=$(avs-devnet get-ports | yq .el-1-besu-lighthouse.rpc)
echo "Fetched RPC URL: $RPC_URL"

TASK_MANAGER_ADDR=$(avs-devnet get-address avs_addresses:incredibleSquaringTaskManager)
echo "Fetched TaskManager address: $TASK_MANAGER_ADDR"

# Fetch the latest task number
echo "Fetching latest task number..."
LATEST_TASK_NUM=$(cast call --rpc-url $RPC_URL $TASK_MANAGER_ADDR 'latestTaskNum()(uint32)')

echo -n "Got: $LATEST_TASK_NUM. "

# Check that task number is a number
case $LATEST_TASK_NUM in
    ''|*[!0-9]*) echo "Invalid task number" ; exit 1 ;;
    *) echo "OK" ;;
esac


# Sleep until the task is completed
# TODO: in paper, this should take only 2 blocks (24s), but it takes more sometimes. Is that OK?
SLEEP_TIME=60
echo "Sleeping for ${SLEEP_TIME}s so the AVS completes the task"
sleep $SLEEP_TIME


# Fetch the task response hash
echo "Fetching task response $LATEST_TASK_NUM..."
SOME_RESPONSE=$(cast call --rpc-url $RPC_URL $TASK_MANAGER_ADDR 'allTaskResponses(uint32)(bytes32)' $LATEST_TASK_NUM)

echo -n "Got: $SOME_RESPONSE. "

EMPTY_RESPONSE="0x0000000000000000000000000000000000000000000000000000000000000000"

# Check the task response hash is a 32 bytes hexstring
if [[ "$SOME_RESPONSE" =~ ^0x[0-9A-Fa-f]{64}$ ]]; then
    # Check that task number is non-empty
    if [[ "$SOME_RESPONSE" == "$EMPTY_RESPONSE" ]]; then
        echo "Empty response"
        exit 2
    fi
    echo "OK"
else
    echo "Invalid task response"
    exit 3
fi
