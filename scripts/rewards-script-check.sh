#!/bin/bash

# Supposing anvil is already running

# Build and deploy contracts
# make build-contracts
make deploy-all

# Check that at first, claimer balance in token is zero
initialBalance=$(make claimer-account-token-balance | tail -n1 | tr -d '[:space:]')

echo "Initial balance: '$initialBalance'"

if [ "$initialBalance" -ne 0 ]; then
    echo "claimer balance in token should be zero"
    exit 2
fi

# Create and claim distribution root
echo "Creating distribution root:"
make create-avs-distributions-root

echo "Claiming distribution root:"
make claim-distributions

# Check that after claims, claimer balance in token is 100
balanceAfterClaim=$(make claimer-account-token-balance | tail -n1 | tr -d '[:space:]')

echo "Balance after claim: '$balanceAfterClaim'"

if [ "$balanceAfterClaim" -ne 100 ]; then
    echo "After first claim, claimer balance in token should be 100"
    exit 3
fi
