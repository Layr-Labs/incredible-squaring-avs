# Incredible Squaring AVS

[![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/incredible-squaring-avs)](https://goreportcard.com/report/github.com/Layr-Labs/incredible-squaring-avs)

<b> Do not use it in Production, testnet only. </b>

Basic repo demoing a simple AVS middleware with full eigenlayer integration. See this [video walkthrough](https://www.loom.com/share/50314b3ec0f34e2ba386d45724602d76?sid=9d68d8cb-d2d5-4123-bd06-776de2076de0).

## Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation) and [zap-pretty](https://github.com/maoueh/zap-pretty) and docker to run the examples below.

``` bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty/cmd/zap-pretty@latest
```

You will also need to [install docker](https://docs.docker.com/get-docker/), and build the contracts:

``` bash
make build-contracts
```

You will also need to [install abigen](https://geth.ethereum.org/docs/tools/abigen) if you want to make changes to the smart contracts and then generate the go bindings:

``` bash
make bindings
```

## Running via make

This simple session illustrates the basic flow of the AVS. The makefile commands are hardcoded for a single operator, but it's however easy to create new operator config files, and start more operators manually (see the actual commands that the makefile calls).

Start anvil in a separate terminal:

```bash
anvil
```

Deploy contracts, set UAM permissions, and create a quorum in a single command:

```bash
make deploy-all
```

Start the aggregator:

```bash
make start-aggregator
```

Register the operator with eigenlayer and incredible-squaring, and then start the process:

```bash
make start-operator
```

By default, the `start-operator` command will also register the operator.
To disable this, set `register_operator_on_startup` to false in `config-files/operator.anvil.yaml`.
The operator can be manually registered by running `make cli-setup-operator`.

The operator will produce an invalid result 10 times out of 100, as it is set in the `times_failing` field of the config.
These failures result in slashing once they're challenged.
To see this in action, start the challenger with:

```bash
make start-challenger
```

## Running via docker compose

We wrote a [docker-compose.yml](./docker-compose.yml) file to run and test everything on a single machine. It will start an anvil instance, loading a [state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) where the eigenlayer and incredible-squaring contracts are deployed, start the aggregator, and finally one operator, along with prometheus and grafana servers. The grafana server will be available at <http://localhost:3000>, with user and password both set to `admin`. We have created a simple [grafana dashboard](./grafana/provisioning/dashboards/AVSs/incredible_squaring.json) which can be used as a starting example and expanded to include AVS specific metrics. The eigen metrics should not be added to this dashboard as they will be exposed on the main eigenlayer dashboard provided by the eigenlayer-cli.

## Creating and Claiming Distributions

The example exposes 3 scripts in the Makefile interface:

- Creating a distribution root, that implies creating an AVS rewards submission and submitting a payment root.
- Creating an operator directed distribution root, similar to previous one but with rewards to operators involved in the claim generation. Note: operators in this case are hardcoded in the script file.
- Claiming the created distribution, giving the rewards to an specific receiver account. Note: The receiver in this case is harcoded in the script file (address 0x01).

This leads to 2 possible workflows, distributing equally across all operators and using custom distribution for each operator.

### Distributing equally across all operators

First, start anvil in a separate terminal and deploy the contracts. To do that follow the instructions in [To run section](#to-run)

Then, run the command:

``` bash
make create-avs-distributions-root
```

This creates a claimable root, a root of the merkle tree that stores cumulative earnings per ERC20 reward token for each earner.

To claim against the root, use:

``` bash
make claim-distributions
```

If you want to check the balance of the claimer, you can run the following command:

``` bash
make claimer-account-token-balance
```

Note that the claimer address is not passed by parameter, because in the script that address is hardcoded.

### Using custom distribution for each operator

First, start anvil in a separate terminal and deploy the contracts. To do that follow the instructions in [To run section](#to-run)

Then, run the command:

``` bash
make create-operator-directed-distributions-root
```

This creates a claimable root, that differs from the previous one in the fact that also distributes the claim to the directed operators established in the script (currently hardcoded).

The payment leaves are available in `contracts/payments.json`. The payment leaves are the keccak256 hash of each earner leaf. An earner leaf is composed by the earner and the token root of the token leaves, and each token leaf is the result of hashing the token address with the token earnings.

To claim against the root, use:

``` bash
make claim-distributions
```

If you want to check the balance of the claimer, you can run the following command:

``` bash
make claimer-account-token-balance
```

Note that the claimer address is not passed by parameter, because in the script that address is hardcoded.

## Avs Task Description

The architecture of the AVS contains:

- [Eigenlayer core](https://github.com/Layr-Labs/eigenlayer-contracts/tree/master) contracts
- AVS contracts
  - [ServiceManager](contracts/src/IncredibleSquaringServiceManager.sol) which will eventually contain slashing logic but for M2 is just a placeholder.
  - [TaskManager](contracts/src/IncredibleSquaringTaskManager.sol) which contains [task creation](contracts/src/IncredibleSquaringTaskManager.sol#L83) and [task response](contracts/src/IncredibleSquaringTaskManager.sol#L102) logic.
  - The [challenge](contracts/src/IncredibleSquaringTaskManager.sol#L176) logic could be separated into its own contract, but we have decided to include it in the TaskManager for this simple task.
  - Set of [registry contracts](https://github.com/Layr-Labs/eigenlayer-middleware) to manage operators opted in to this avs
- Task Generator
  - in a real world scenario, this could be a separate entity, but for this simple demo, the aggregator also acts as the task generator
- Aggregator
  - aggregates BLS signatures from operators and posts the aggregated response to the task manager
  - For this simple demo, the aggregator is not an operator, and thus does not need to register with eigenlayer or the AVS contract. It's IP address is simply hardcoded into the operators' config.
- Operators
  - Square the number sent to the task manager by the task generator, sign it, and send it to the aggregator

![](./diagrams/architecture.png)

1. A task generator (in our case, same as the aggregator) publishes tasks once every regular interval (say 10 blocks, you are free to set your own interval) to the IncredibleSquaringTaskManager contract's [createNewTask](contracts/src/IncredibleSquaringTaskManager.sol#L83) function. Each task specifies an integer `numberToBeSquared` for which it wants the currently opted-in operators to determine its square `numberToBeSquared^2`. `createNewTask` also takes `quorumNumbers` and `quorumThresholdPercentage` which requests that each listed quorum (we only use quorumNumber 0 in incredible-squaring) needs to reach at least thresholdPercentage of operator signatures.

2. A [registry](https://github.com/Layr-Labs/eigenlayer-middleware/blob/master/src/BLSRegistryCoordinatorWithIndices.sol) contract is deployed that allows any eigenlayer operator with at least 1 delegated [mockerc20](contracts/src/ERC20Mock.sol) token to opt-in to this AVS and also de-register from this AVS.

3. [Operator] The operators who are currently opted-in with the AVS need to read the task number from the Task contract, compute its square, sign on that computed result (over the BN254 curve) and send their taskResponse and signature to the aggregator.

4. [Aggregator] The aggregator collects the signatures from the operators and aggregates them using BLS aggregation. If any response passes the [quorumThresholdPercentage](contracts/src/IIncredibleSquaringTaskManager.sol#L36) set by the task generator when posting the task, the aggregator posts the aggregated response to the Task contract.

5. If a response was sent within the [response window](contracts/src/IncredibleSquaringTaskManager.sol#L119), we enter the [Dispute resolution] period.
   - [Off-chain] A challenge window is launched during which anyone can [raise a dispute](contracts/src/IncredibleSquaringTaskManager.sol#L171) in a DisputeResolution contract (in our case, this is the same as the TaskManager contract)
   - [On-chain] The DisputeResolution contract resolves that a particular operator’s response is not the correct response (that is, not the square of the integer specified in the task) or the opted-in operator didn’t respond during the response window. If the dispute is resolved, the operator will be frozen in the Registration contract and the veto committee will decide whether to veto the freezing request or not.

Below is a more detailed uml diagram of the aggregator and operator processes:

![](./diagrams/uml.png)

## StakeUpdates Cronjob

AVS Registry contracts have a stale view of operator shares in the delegation manager contract. In order to update their stake table, they need to periodically call the [StakeRegistry.updateStakes()](https://github.com/Layr-Labs/eigenlayer-middleware/blob/f171a0812126bbb0bb6d44f53c622591a643e987/src/StakeRegistry.sol#L76) function. We are currently writing a cronjob binary to do this for you, will be open sourced soon!

## Integration Tests

See the integration tests [README](tests/anvil/README.md) for more details.

## Structure Documentation

This AVS has three main participants:

- Operator: The operator subscribes to NewTasks Events and, when a new task is created, completes it, calculates the response, signs it, and sends it to the BLS aggregation service.
- Aggregator: The one who creates new tasks for the operators (through the on-chain `TaskManager`) every certain time. It also collects aggregated responses from the BLS aggregation service and sends them to the on-chain `TaskManager`, which then emits a TaskRespondedEvent.
- Challenger: The Challenger subscribes to TaskRespondedEvents, and in case the response given by the aggregator differs from the Challenger calculated response, it raises a challenge, that calls on-chain `TaskManager`, which verifies if the aggregator response was right. If it was not right, then the operator that signed the task will be slashed.

Now we will focus on each to show how each one does each thing.

### Operator

The operator code can be found in [`/operator` folder](https://github.com/Layr-Labs/incredible-squaring-avs/tree/dev/operator).

The operator's main logic is focused on this segment from [operator.go](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/operator/operator.go#L340-L363):

```go
for {
    select {
    case <-ctx.Done():
        ...
    case err := <-metricsErrChan:
        ...
    case err := <-sub.Err():
        ...
    case newTaskCreatedLog := <-o.newTaskCreatedChan:
        ...
    }
}
```

The upper three cases are handling error cases, the fourth one is the one that pops from the channel subscribed to new task creation events, and handles the response logic:

```go
o.metrics.IncNumTasksReceived()
taskResponse := o.ProcessNewTaskCreatedLog(newTaskCreatedLog)
signedTaskResponse, err := o.SignTaskResponse(taskResponse)
if err != nil {
    continue
}
go o.aggregatorRpcClient.SendSignedTaskResponseToAggregator(signedTaskResponse)
```

The [`ProcessNewTaskCreatedLog` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/operator/operator.go#L368-L394) generates the response to the new task:

```go
func (o *Operator) ProcessNewTaskCreatedLog(
    newTaskCreatedLog *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated,
) *cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse {
    ...
    numberSquared := big.NewInt(0).Exp(newTaskCreatedLog.Task.NumberToBeSquared, big.NewInt(2), nil)

    ...
    taskResponse := &cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse{
        ReferenceTaskIndex: newTaskCreatedLog.TaskIndex,
        NumberSquared:      numberSquared,
    }
    return taskResponse
}
```

Here is the response calculation logic, and it would be the place to change if you wanted to compute, for example, the cubed number instead. Note that the `Response` struct includes the number square because it's part of the `TaskManager` contract bindings, which should be modified too.

After the ProcessNewTaskCreatedLog function, that response is signed (in [SignTaskResponse](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/operator/operator.go#L396-L416)), and sent to the BLS aggregation service in the goroutine executing the [`SendSignedTaskResponseToAggregator()` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/operator/rpc_client.go#L52). That function makes a call to the `ProcessSignedTaskResponse` method of aggregator (through RPC), that redirects the signed response to the BLS aggregation service.

### Aggregator

The aggregator code can be found in [`/aggregator` folder](https://github.com/Layr-Labs/incredible-squaring-avs/tree/dev/aggregator).

The main aggregator logic can be found on this loop from [aggregator.go](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/aggregator/aggregator.go#L194-L209):

```go
for {
    select {
    case <-ctx.Done():
         ...
    case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
        agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
        agg.sendAggregatedResponseToContract(blsAggServiceResp)
    case <-ticker.C:
        err := agg.sendNewTask(big.NewInt(taskNum))
        taskNum++
        if err != nil {
            continue
        }
    }
}
```

The first case covers the context-done error case. The second covers the case where a new aggregated response is received from the BLS aggregation service. Remember that this happens when the operator responses to the tasks reach a threshold or the time of the task expires. In this case, the [`sendAggregatedResponseToContract()` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/aggregator/aggregator.go#L212-L254) is called.

That method wraps the response into a more complex `TaskManager` type that encapsulates the response and sends it with the completed to the on-chain Task Manager’s [`respondToTask` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/contracts/src/IncredibleSquaringTaskManager.sol#L118-L122).

That method makes several checks on the task response, stores the responses metadata and emits a `TaskResponded` event, that will be catched by the challenger (see challenger section to continue).

The third case of the main loop is the one which spawns new tasks every 10 seconds for the operators to complete, calling aggregator [`sendNewTask()` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/aggregator/aggregator.go#L258-L297). There the aggregator calls the [`CreateNewTask()` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/contracts/src/IncredibleSquaringTaskManager.sol#L99-L103) of the on-chain `TaskManager` contract, that stores a hash of the new task and emits a `NewTaskCreated` event, that will be caught by the challenger (see challenger section to continue). After that call to the `TaskManager`, the aggregator will initialize a new task in the BLS aggregation service, where the operators will send their signed response to the created task.

### Challenger

The challenger code can be found on the [`/challenger` folder](https://github.com/Layr-Labs/incredible-squaring-avs/tree/dev/challenger).

The main behavior of the challenger is to subscribe to the `NewTaskCreated` and `TaskResponded` events emitted by the on-chain `TaskManager` contract and can be found on [challenger.go](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/challenger/challenger.go#L79-L117).

```go
for {
    select {
    case err := <-newTaskSub.Err():
        ...
    case err := <-taskResponseSub.Err():
        ...
    case newTaskCreatedLog := <-c.newTaskCreatedChan:
        ...
        taskIndex := c.processNewTaskCreatedLog(newTaskCreatedLog)
        if _, found := c.taskResponses[taskIndex]; found {
            _ = c.callChallengeModule(taskIndex)
        }
    case taskResponseLog := <-c.taskResponseChan:
        ...
        taskIndex := c.processTaskResponseLog(taskResponseLog)
        if _, found := c.tasks[taskIndex]; found {
            _ = c.callChallengeModule(taskIndex)
        }
    }
}
```

The first two cases handle errors in the subscribed event channels. The other two listen to events and process them. In the case of `NewTaskCreated`, it means saving the created task for future events. In the case of TaskResponse, it means generating and saving the taskResponseData, that could be sent to the `TaskManager` in case of a challenge.

After the processing, the newTaskCreated case checks if there is a task response with that index, and the TaskResponse case checks if there's an initialized task with that index, and in both cases there is a call to the [callChallengeModule method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/challenger/challenger.go#L152-L169).

```go
func (c *Challenger) callChallengeModule(taskIndex uint32) error {
    numberToBeSquared := c.tasks[taskIndex].NumberToBeSquared
    answerInResponse := c.taskResponses[taskIndex].TaskResponse.NumberSquared
    trueAnswer := numberToBeSquared.Exp(numberToBeSquared, big.NewInt(2), nil)

    // Checking if the answer in the response submitted by the aggregator is correct
    if trueAnswer.Cmp(answerInResponse) != 0 {
        c.logger.Info("The number squared is not correct", "expectedAnswer", trueAnswer, "gotAnswer", answerInResponse)

        // Raise challenge
        c.raiseChallenge(taskIndex)

        return nil
    } else {
        c.logger.Info("The number squared is correct")
        return types.NoErrorInTaskResponse
    }
}
```

In this method, the challenger calculates the response and compares it with the aggregators response. If the response is not equal, a challenge is raised, what means a call to on-chain `TaskManager` [`RaiseAndResolveChallenge()` method](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/contracts/src/IncredibleSquaringTaskManager.sol#L175-L180).

```solidity
function raiseAndResolveChallenge(
    Task calldata task,
    TaskResponse calldata taskResponse,
    TaskResponseMetadata calldata taskResponseMetadata,
    BN254.G1Point[] memory pubkeysOfNonSigningOperators
) external {
    ...        
    // // Logic for checking whether the challenge is valid or not
    uint256 actualSquaredOutput = numberToBeSquared * numberToBeSquared;
    bool isResponseCorrect = (actualSquaredOutput == taskResponse.numberSquared);
    // //If the response was correct, no slashing happens so we return
    if (isResponseCorrect == true) {
        emit TaskChallengedUnsuccessfully(referenceTaskIndex, msg.sender);
        return;
    }
    ...
}
```

In that method the `TaskManager` calculates the response and determines if the aggregated response is correct or not. In the first case, nothing happens, but in the second case, the signer operators will be slashed.

The slashing mechanism can be found in the [second part](https://github.com/Layr-Labs/incredible-squaring-avs/blob/f8c379b151d8db778a12a5de1ba0266436d85366/contracts/src/IncredibleSquaringTaskManager.sol#L261-L279) of the raiseAndResolveChallenge method, but in a simple way to explain, the Manager defines an amount of wads to slash from each operator, and calls the [`InstantSlasher.fulfillSlashingRequest()` method](https://github.com/Layr-Labs/eigenlayer-middleware/blob/4d63f27247587607beb67f96fdabec4b2c1321ef/src/slashers/InstantSlasher.sol#L22-L31), that ends up calling the [`allocationManager.slashOperator()` method](https://github.com/Layr-Labs/eigenlayer-contracts/blob/aa84b7a1d801510a9b893be2f2a91e8ef093faf6/src/contracts/core/AllocationManager.sol#L64-L67).

## Troubleshooting

### Received error from aggregator

When running on anvil, a typical log for the operator is

``` bash
[2024-04-09 18:25:08.647 PDT] INFO (logging/zap_logger.go:49) rpc client is nil. Dialing aggregator rpc client
[2024-04-09 18:25:08.650 PDT] INFO (logging/zap_logger.go:49) Sending signed task response header to aggregator {"signedTaskResponse":"\u0026aggregator.SignedTaskResponse{TaskResponse:contractIncredibleSquaringTaskManager.IIncredibleSquaringTaskManagerTaskResponse{ReferenceTaskIndex:0x2, NumberSquared:4}, BlsSignature:bls.Signature{G1Point:(*bls.G1Point)(0x14000282068)}, OperatorId:[32]uint8{0xc4, 0xc2, 0x10, 0x30, 0xe, 0x28, 0xab, 0x4b, 0xa7, 0xb, 0x7f, 0xbb, 0xe, 0xfa, 0x55, 0x7d, 0x2a, 0x2a, 0x5f, 0x1f, 0xbf, 0xa6, 0xf8, 0x56, 0xe4, 0xcf, 0x3e, 0x9d, 0x76, 0x6a, 0x21, 0xdc}}"}
[2024-04-09 18:25:08.651 PDT] INFO (logging/zap_logger.go:49) Received error from aggregator {"err":"task 2 not initialized or already completed"}
[2024-04-09 18:25:08.651 PDT] INFO (logging/zap_logger.go:69) Retrying in 2 seconds
[2024-04-09 18:25:10.679 PDT] INFO (logging/zap_logger.go:49) Signed task response header accepted by aggregator. {"reply":false}
```

The error `task 2 not initialized or already completed` is expected behavior. This is because the aggregator needs to setup its data structures before it can accept responses. But on a local anvil setup, the operator had time to receive the websocket event for the new task, square the number, sign the response, and send it to the aggregator process before the aggregator has finalized its setup. Hence, the operator retries sending the response 2 seconds later and it is accepted.
