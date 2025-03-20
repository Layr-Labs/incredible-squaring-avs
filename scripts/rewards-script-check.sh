#!/bin/bash

# This flags makes the script exit if any command has a non-zero exit status, or 
# if tries to use a non defined variable
set -e -o nounset

# Execute anvil in background
anvil -q &

# Deploy contracts
make deploy-all

# Check that at first, claimer balance in token is zero
initialBalance=$(make claimer-account-token-balance | tail -n1 | tr -d '[:space:]')

echo "Initial balance: '$initialBalance'"

if [ "$initialBalance" -ne 0 ]; then
    echo "claimer balance in token should be zero"
    exit 2
fi

# Create and claim normal distribution root
echo "Creating distribution root:"
make create-avs-distributions-root

echo "Claiming distribution root:"
make claim-distributions

# Check that after claim, claimer balance in token is 100
balanceAfterClaim=$(make claimer-account-token-balance | tail -n1 | tr -d '[:space:]')

echo "Balance after first claim: '$balanceAfterClaim'"

if [ "$balanceAfterClaim" -ne 100 ]; then
    echo "After first claim, claimer balance in token should be 100"
    exit 3
fi

# Create and claim operator directed distribution root
echo "Creating operator directed distribution root:"
make create-operator-directed-distributions-root

echo "Claiming distribution root:"
make claim-distributions

# Check that after another claim, claimer balance in token is 200
balanceAfterClaim=$(make claimer-account-token-balance | tail -n1 | tr -d '[:space:]')

echo "Balance after second claim: '$balanceAfterClaim'"

if [ "$balanceAfterClaim" -ne 200 ]; then
    echo "After second claim, claimer balance in token should be 200"
    exit 3
fi

# Kill anvil executing in background
kill $(pgrep anvil)
