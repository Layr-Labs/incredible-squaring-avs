#!/bin/bash

AVS_CONTRACTS_NAME="$1"

AVS_LOWER_PREFIX=$(tr '[:upper:]' '[:lower:]' <<< ${AVS_CONTRACTS_NAME:0:1})${AVS_CONTRACTS_NAME:1}

TASK_MANAGER_CONTRACT_NAME=${AVS_CONTRACTS_NAME}TaskManager
SERVICE_MANAGER_CONTRACT_NAME=${AVS_CONTRACTS_NAME}ServiceManager

AVS_WITH_CAMEL=$(echo "$AVS_CONTRACTS_NAME" | sed 's/\([^A-Z]\)\([A-Z]\)/\1-\2/g')
MODULE_NAME=$(echo "$AVS_WITH_CAMEL" | tr '[:upper:]' '[:lower:]')-avs

git init

rm -rf contracts/lib/eigenlayer-middleware

git submodule add git@github.com:Layr-Labs/eigenlayer-middleware.git contracts/lib/eigenlayer-middleware

rm -rf contracts/lib/forge-std

git submodule add https://github.com/foundry-rs/forge-std.git contracts/lib/forge-std

cd contracts/lib/eigenlayer-middleware

git checkout v1.4.0-testnet-holesky

cd ../../..

git add *

git submodule update --init --recursive

mv contracts/src/IIncredibleSquaringTaskManager.sol contracts/src/I${TASK_MANAGER_CONTRACT_NAME}.sol
mv contracts/src/IncredibleSquaringServiceManager.sol contracts/src/${SERVICE_MANAGER_CONTRACT_NAME}.sol
mv contracts/src/IncredibleSquaringTaskManager.sol contracts/src/${TASK_MANAGER_CONTRACT_NAME}.sol

find . -type f -name "*.sol" -exec sed -i '' "s/IncredibleSquaringServiceManager/${SERVICE_MANAGER_CONTRACT_NAME}/g" {} +
find . -type f -name "*.sol" -exec sed -i '' "s/incredibleSquaringServiceManager/${AVS_LOWER_PREFIX}ServiceManager/g" {} +
find . -type f -name "*.sol" -exec sed -i '' "s/IncredibleSquaringTaskManager/${TASK_MANAGER_CONTRACT_NAME}/g" {} +
find . -type f -name "*.sol" -exec sed -i '' "s/incredibleSquaringTaskManager/${AVS_LOWER_PREFIX}TaskManager/g" {} +

sed -i '' "s|IncredibleSquaringTaskManager|${TASK_MANAGER_CONTRACT_NAME}|g" contracts/generate-go-bindings.sh
sed -i '' "s|IncredibleSquaringServiceManager|${SERVICE_MANAGER_CONTRACT_NAME}|g" contracts/generate-go-bindings.sh

make bindings

# Replace module name in go.mod file
sed -i '' "s|github.com/Layr-Labs/incredible-squaring-avs|github.com/Layr-Labs/${MODULE_NAME}|g" go.mod

# Use the new module name in the module local imports
find . -type f -name "*.go" -exec sed -i '' "s|\"github.com/Layr-Labs/incredible-squaring-avs|\"github.com/Layr-Labs/${MODULE_NAME}|g" {} +

# Change the binding import in the module files
find . -type f -name "*.go" -exec sed -i '' "s|cstaskmanager \"github.com/Layr-Labs/${MODULE_NAME}/contracts/bindings/IncredibleSquaringTaskManager\"|cstaskmanager \"github.com/Layr-Labs/${MODULE_NAME}/contracts/bindings/${TASK_MANAGER_CONTRACT_NAME}\"|g" {} +

# Replace the name of the contract when getting the Task Manager ABI
find . -type f -name "*.go" -exec sed -i '' "s|taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()|taskManagerAbi, err := cstaskmanager.Contract${TASK_MANAGER_CONTRACT_NAME}MetaData.GetAbi()|g" {} +

# Extra replaces (not necessary to have the system working)

find . -type f -name "*.go" -exec sed -i '' "s|csservicemanager \"github.com/Layr-Labs/${MODULE_NAME}/contracts/bindings/IncredibleSquaringServiceManager\"|csservicemanager \"github.com/Layr-Labs/${MODULE_NAME}/contracts/bindings/${SERVICE_MANAGER_CONTRACT_NAME}\"|g" {} +

find . -type f -name "*.go" -exec sed -i '' "s|ContractIncredibleSquaring|Contract${AVS_CONTRACTS_NAME}|g" {} +

find . -type f -name "*.go" -exec sed -i '' "s|contractServiceManager.IncredibleSquaringTaskManager|contractServiceManager.${TASK_MANAGER_CONTRACT_NAME}|g" {} +
