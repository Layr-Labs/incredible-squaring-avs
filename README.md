# Trading Algo AVS

This is TradingAlgoAVS's AVS component. For the full project repo, visit [TradingAlgoAVS](https://github.com/ahsueh1996/TradeAlgoAVS)

Note that this repo is in early development. Do not use it in production! Testnet only!

AVS is built on top of Layr Labs's Eigen Layer AVS, bootstrapped from their opensource [incredible-squaring-avs example](https://github.com/Layr-Labs/incredible-squaring-avs/tree/master)
 [![Go Report Card](https://goreportcard.com/badge/github.com/Layr-Labs/incredible-squaring-avs)](https://goreportcard.com/report/github.com/Layr-Labs/incredible-squaring-avs)

---

## 1. AVS Design

### 1.1 Trust Model
TradeAlgoAVS ensures that Operators:
- **Execute trades correctly** according to the strategy logic.
- **Do not have custody of user funds**.
- **Keep strategy logic confidential**.
- **Distribute fees based on the agreed subscription model**.

**Malicious Behaviors and Slashing**
- Operators are **slashed** if:
  - They deviate significantly from the expected strategy performance.
  - They execute trades incorrectly or inefficiently.
  - They engage in front-running or exploitative behavior.

- **Evidence of misconduct** is submitted on-chain via **fraud-proof mechanisms**.

---

## 2. Task Design
Each **trading strategy execution** is modeled as a **Task** in the AVS network.

### Task Workflow:
1. **Strategy provider submits a strategy hash** and execution rules.
2. **Subscribers choose strategies and sign a fee agreement**.
3. **AI Agent Operators on Autonome fetch execution signals and trade on behalf of subscribers**.
4. **Operators submit signed execution proofs on-chain**.
5. **Validators verify execution and distribute fees accordingly**.

---

## 3. Validating Operator Execution

### M-of-N Aggregation Based on Strategy Execution Statistics
TradeAlgoAVS employs an **M-of-N aggregation model** where multiple Operators execute the same strategy, and their **performance statistics are compared**. 

The assumption is that:
1. **The same trading strategy, under identical market conditions, should yield statistically similar results within an expected standard deviation**.
2. **Investors (strategy users) do not care about the individual order-level execution but rather the final performance metrics (profit, ROI, drawdown, Sharpe ratio, etc.)**.

### Validation Workflow:
1. **Multiple Operators execute the strategy independently**.
2. **Their trade performance metrics are compared**:
   - ROI
   - Profit percentage
   - Execution slippage
   - Sharpe ratio
3. **On-chain smart contracts aggregate these results**.
4. **If an Operator's results deviate significantly** from the expected statistical range, they are:
   - **Flagged for fraud**.
   - **Slashed if proven malicious**.

### Why This Approach Works:
✅ **No need to verify every single trade step-by-step**.  
✅ **Ensures Operators remain accountable to overall performance**.  
✅ **Allows AI Agents to autonomously optimize execution while staying within acceptable performance bounds**.  
✅ **Efficient and scalable** compared to full trade-by-trade validation.

---

## 4. Self-Curating Strategy Selection and Operator Reputation System

### 4.1 How Operator's Good Behaviour is Incentivized

- **Slashing Mechanisms**
  - Operators **lose staked funds** for malicious or incorrect execution (that deviates from the quorum norm)
  - Fraudulent executions can result in **disqualification from the network**.

- **Reputation System**
  - High-reputation Operators gain **higher trade execution priority**.
  - Bad actors get **blacklisted** from the system.


### 4.2 How TradeAlgoAVS Self-Curates Strategies

#### A. Operators Are Penalized for Running Poor-Performing Strategies

Each Operator’s reputation is tied to their execution performance. If an Operator executes a strategy that consistently underperforms (or performs worse than advertised), their reputation score decreases. Since Operators compete for execution fees, bad reputation reduces their ability to attract users.

#### B. Operators Have an Incentive to Select Profitable Strategies

Since their reputation is at stake, Operators will avoid executing bad strategies. Two ways Operators can self-curate strategies:

1. Reject Strategies with Poor Historical Performance

2. Charge Higher Fees for Riskier Strategies

#### C. How This Creates a Self-Sustaining Ecosystem

Bad strategies naturally get filtered out because Operators won’t want to execute them. Only high-quality, profitable strategies remain, ensuring better returns for users. Strategy providers are incentivized to improve their strategies to ensure Operators will run them.

### 4.2 Smart Contract Implementation for Reputation & Strategy Selection

Operators’ reputation scores are stored on-chain and updated based on execution performance. Strategies have an on-chain performance rating, influencing Operator selection. Operators adjust execution fees dynamically based on strategy risk.

---
## 5. Why EigenLayer for TradeAlgoAVS?

### Security
- Leverages **EigenLayer's restaking model**.
- Operators **stake ETH**, ensuring security.

### Efficiency
- **M-of-N statistical aggregation** reduces gas costs.
- **Batch processing** improves execution speed.

### Decentralization
- No single Operator controls execution.
- **Multiple AI Agents verify execution**.

### Confidentiality
- Users never see the **strategy code**.
- Strategy providers retain **full IP control**.

---
## 6. Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation) and [zap-pretty](https://github.com/maoueh/zap-pretty) and docker to run the examples below.
```
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty@latest
```
You will also need to [install docker](https://docs.docker.com/get-docker/), and build the contracts:
```
make build-contracts
```
You will also need to [install abigen](https://geth.ethereum.org/docs/tools/abigen) if you want to make changes to the smart contracts and then generate the go bindings 
```
make bindings
```

---
## 7. Running The Repo

### 7.1 Make

This simple session illustrates the basic flow of the AVS. The makefile commands are hardcoded for a single operator, but it's however easy to create new operator config files, and start more operators manually (see the actual commands that the makefile calls).

Start anvil in a separate terminal:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

The above command starts a local anvil chain from a [saved state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) with eigenlayer and incredible-squaring contracts already deployed (but no operator registered).

Start the aggregator:

```bash
make start-aggregator
```

Register the operator with eigenlayer and incredible-squaring, and then start the process:

```bash
make start-operator
```

> By default, the `start-operator` command will also setup the operator (see `register_operator_on_startup` flag in `config-files/operator.anvil.yaml`). To disable this, set `register_operator_on_startup` to false, and run `make cli-setup-operator` before running `start-operator`.

### 7.2 Running via docker compose

See [docker-compose.yml](./docker-compose.yml) file to run and test everything on a single machine. It will start an anvil instance, loading a [state](./tests/anvil/avs-and-eigenlayer-deployed-anvil-state.json) where the eigenlayer and incredible-squaring contracts are deployed, start the aggregator, and finally one operator, along with prometheus and grafana servers. The grafana server will be available at http://localhost:3000, with user and password both set to `admin`. 

See a simple [grafana dashboard](./grafana/provisioning/dashboards/AVSs/incredible_squaring.json) which can be used as a starting example and expanded to include AVS specific metrics. The eigen metrics should not be added to this dashboard as they will be exposed on the main eigenlayer dashboard provided by the eigenlayer-cli.