#!/bin/bash

SENDER_ADDR=0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

cd contracts

forge script script/SetupDistributions.s.sol:SetupDistributions --rpc-url http://localhost:8545 \
    --broadcast -v --sender $SENDER_ADDR

forge script script/SetupDistributions.s.sol:SetupDistributions --rpc-url http://localhost:8545 \
    --broadcast --sig \"executeProcessClaim()\" -v --sender $SENDER_ADDR
