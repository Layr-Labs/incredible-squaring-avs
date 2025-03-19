#!/bin/bash

# Supposing anvil is already running

# Build and deploy contracts
make build-contracts
make deploy-all

# Check that at first, claimer balance in token is zero
initialBalance=$(make claimer-account-token-balance)

# Create and claim distribution root
make create-avs-distributions-root
make claim-distributions

# Check that after claims, claimer balance in token is 100
balanceAfterClaim=$(make claimer-account-token-balance)
