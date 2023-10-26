# Integration Tests

We store an anvil state files in this directory, so that we can start an anvil chain with the correct state for integration tests.
```
anvil --load-state STATE_FILE.json
```

## Eigenlayer deployment state file
`eigenlayer-deployed-anvil-state.json` contains the eigenlayer deployment.

It was created by running this [deploy script](https://github.com/Layr-Labs/eigenlayer-contracts/blob/2cb9ed107c6c918b9dfbac94cd71b4ab7c94e8c2/script/testing/M2_Deploy_From_Scratch.s.sol). If you ever need to redeploy a new version of eigenlayer contracts, first start an anvil chain that dumps its state after exiting
```
anvil --dump-state eigenlayer-deployed-anvil-state.json
```
Then run the deploy script
```
forge script script/testing/M2_Deploy_From_Scratch.s.sol --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast --sig "run(string memory configFile)" -- M2_deploy_from_scratch.anvil.config.json
```
and finally kill the anvil chain with `Ctrl-C`. Make sure to copy the deployment [output file](https://github.com/Layr-Labs/eigenlayer-contracts/blob/master/script/output/M2_from_scratch_deployment_data.json) to [eigenlayer_deployment_output.json](../../contracts/script/output/31337/eigenlayer_deployment_output.json) so that the tests can find the deployed contracts.

See the main [README](../../README.md#dependencies) to understand why we deploy from the `experimental-reduce-strategy-manager-bytecode-size` branch of eigenlayer-contracts.