# Price Oracle AVS

[![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/incredible-squaring-avs)](https://goreportcard.com/report/github.com/Layr-Labs/incredible-squaring-avs)

<b> Do not use it in Production, testnet only. </b>



## Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation), [zap-pretty](https://github.com/maoueh/zap-pretty) and docker to run the examples below, and [geth](https://geth.ethereum.org/docs/getting-started/installing-geth)
```
curl -L https://foundry.paradigm.xyz | bash
foundryup
brew install ethereum
brew install maoueh/tap/zap-pretty
brew install jq
```
You will also need to [install docker](https://docs.docker.com/get-docker/), and build the contracts:
```
make build-contracts
```

## Running via make

This simple session illustrates the basic flow of the AVS. The makefile commands are hardcoded for a single operator, but it's however easy to create new operator config files, and start more operators manually (see the actual commands that the makefile calls).

Before running these scripts ensure you create a `.env` file that looks like the following.
```bash
SEPOLIA_RPC= #rpc endpoint must be an archival node
```

We fork of sepolia testnet to ensure our local anivl node has access to testnet data feeds.

Start anvil in a separate terminal:

```bash
make deploy-all-to-anvil-and-save-state && make bindings && make start-anvil-chain-with-el-and-avs-deployed
```

The above command starts a local anvil chain from a [saved state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) with eigenlayer and incredible-squaring contracts already deployed (but no operator registered).

In order to have a consensus the raft protocol requires at least 3 nodes. So to run three operators run each of these commands in a seperate terminal:

```bash
make start-operator-1
make start-operator-2
make start-operator-3
```

Start the task generator server to generate a price update task every 10 seconds:

```bash
make start-task-generator
```



> By default, the `start-operator` command will also setup the operator (see `register_operator_on_startup` flag in `config-files/operator.anvil.yaml`). To disable this, set `register_operator_on_startup` to false, and run `make cli-setup-operator` before running `start-operator`.

