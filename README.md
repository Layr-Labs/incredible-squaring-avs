# Smart Contract Auditor AVS

## Project Overview

This AVS is designed to facilitate secure, efficient, and transparent smart contract auditing using Zero-Knowledge Proofs (ZKPs) for bid submissions and verifications. This README provides an overview of the AVS platform's design and architecture, outlining the roles of operators and auditors, the phases of development, and the use of ZKPs to ensure a secure and fair auditing process.

The platform aims to enhance the quality of smart contract audits by:

1. ensuring a decentralized process of verifying and selecting the best auditors.
2. enabling auditors to stake their own assets to earn greater rewards for their audits.

![overall-diagram](./diagrams/overall_architecture.png)

## Key Actors

- **Users:** Individuals or entities who request audits for their smart contracts.
- **Auditors:** Professionals who specialize in smart contract security and bid on audit requests. They provide comprehensive audit reports upon completion of the audit.
- **ZKP Reviewers:** Experts in Zero-Knowledge Proofs who verify the validity of the ZKPs submitted by auditors, ensuring that the bidding process is fair and transparent.

```mermaid
flowchart TD
    U[Users]
    AVS[AVS]
    A[Auditors]
    Z[ZKP Reviewers]

    U -->|Request Audits & Assigns Funds| AVS
    A -->|Submit Bids with ZKPs & Provides Reports| AVS
    Z -->|Verify ZKPs| AVS
    AVS -->|Selects Auditor, Provide Reports & Manages Funds| U

    subgraph ROLES
        U
        A
        Z
        AVS
    end

    U --> ROLES
    A --> ROLES
    Z --> ROLES
    AVS --> ROLES
```

## Main Functions of the AVS

- **Request Management:** Handles the creation, tracking, and management of audit requests, including bid deadlines, criteria, and ZKP reviewer selection.
- **ZKP Verification Coordination:** Facilitates the secure distribution of ZKPs to reviewers, collects their verification results, and determines consensus based on the predefined threshold.
- **Bid Selection:** Evaluates valid bids and selects the winning auditor based on specified criteria.
- **Escrow and Payment:** Manages the secure escrow of funds and their disbursement upon successful audit completion.
- **Report Verification:** Assesses the quality and adherence of audit reports to the initial criteria.
- **Reputation Management (Optional):** Maintains a reputation system for both auditors and ZKP reviewers, tracking their performance and reliability over time.
- **Dispute Resolution (Optional):** Provides a mechanism to resolve disputes that may arise between users, auditors, or ZKP reviewers regarding audit quality, ZKP validity, or payment issues.

```mermaid
flowchart TD
    D[AVS Functions]
    D --> D1[Request Creation]
    D --> D2[Auditor Bidding]
    D --> D3[ZKP Reviewer Selection]
    D --> D4[ZKP Verification]
    D --> D5[Consensus and Bid Selection]
    D --> D6[Funds Escrow]
    D --> D7[Audit Execution]
    D --> D8[Report Submission]
    D --> D9[Report Verification]
    D --> D10[Payment and Funds Release]
    D --> D11[Penalty Mechanism]
```

## Workflow

**Structural Flow of the AVS**

1. **Request Creation:** A user initiates an audit request by providing the smart contract repository link, budget in USDC, bid deadline, audit criteria, desired number of ZKP reviewers, and the consensus threshold for ZKP verification.
2. **Auditor Bidding:** Auditors assess the audit request and submit their bids in a Zero-Knowledge (ZK) manner. Each bid includes:
   - **Encrypted Bid Amount:** The price in USDC, encrypted to maintain confidentiality.
   - **Zero-Knowledge Proof (ZKP):** A cryptographic proof demonstrating that the auditor's qualifications meet the specified criteria without revealing the actual bid amount or specific details of their qualifications.
3. **ZKP Reviewer Selection:** The AVS randomly selects (or selects based on reputation) a predetermined number of ZKP Reviewers from the registered pool.
4. **ZKP Verification:** The selected ZKP Reviewers independently verify the ZKPs submitted by the auditors. They assess whether the proofs are valid and satisfy the audit criteria without revealing the bid amount. Each reviewer submits their verification result (pass/fail) to the AVS.
5. **Consensus and Bid Selection:** The AVS aggregates the verification results from the ZKP Reviewers. If a bid's ZKP receives sufficient passing verifications (based on the predefined consensus threshold), it is considered valid. The AVS then selects the winning bid based on predefined criteria, such as the lowest valid bid amount or a combination of factors like price and reputation.
6. **Funds Escrow:** Upon bid selection, the user's USDC payment is securely held in escrow within the AVS smart contract.
7. **Audit Execution:** The winning auditor conducts a comprehensive audit of the smart contract, adhering to the specified criteria and industry best practices.
8. **Report Submission:** After completing the audit, the auditor submits a detailed report to the AVS, outlining their findings, identified vulnerabilities, and recommendations.
9. **Report Verification:** The AVS assesses the audit report, ensuring its quality, thoroughness, and adherence to the initial audit request criteria.
10. **Payment and Funds Release:** If the audit report is deemed satisfactory, the escrowed funds are released to the winning auditor. Any remaining funds are returned to the user.

```mermaid
flowchart TD
    A[Request Creation] --> B[Auditor Bidding]
    B --> C[ZKP Reviewer Selection]
    C --> D[ZKP Verification]
    D --> E[Consensus and Bid Selection]
    E --> F[Funds Escrow]
    F --> G[Audit Execution]
    G --> H[Report Submission]
    H --> I[Report Verification]
    I --> J[Payment and Funds Release]

    subgraph Request_Creation
        A1[User provides smart contract repo link, budget, bid deadline, audit criteria, ZKP reviewers number, consensus threshold]
    end
    A1 --> A

    subgraph Auditor_Bidding
        B1[Auditors submit bids]
        B2[Encrypted Bid Amount]
        B3[Zero-Knowledge Proof]
    end
    B1 --> B2
    B1 --> B3
    B2 --> B
    B3 --> B

    subgraph ZKP_Reviewer_Selection
        C1[Selects ZKP Reviewers from pool]
    end
    C1 --> C

    subgraph ZKP_Verification
        D1[Reviewers verify ZKPs]
        D2[Reviewers submit pass/fail result]
    end
    D1 --> D2
    D2 --> D

    subgraph Consensus_and_Bid_Selection
        E1[Aggregates verification results]
        E2[Selects winning bid based on criteria]
    end
    E1 --> E2
    E2 --> E

    subgraph Funds_Escrow
        F1[User's USDC payment held in escrow]
    end
    F1 --> F

    subgraph Audit_Execution
        G1[Winning auditor conducts audit]
    end
    G1 --> G

    subgraph Report_Submission
        H1[Auditor submits audit report]
    end
    H1 --> H

    subgraph Report_Verification
        I1[AVS assesses audit report]
    end
    I1 --> I

    subgraph Payment_and_Funds_Release
        J1[Escrowed funds released to auditor]
        J2[Remaining funds returned to user]
    end
    J1 --> J
    J2 --> J
```

## Details on the AVS Smart Contract

1. **AuditRequestManager:**
   - **Request Creation:** Users create audit requests by providing contract details (repository link, etc.), budget (in USDC), bid deadline, desired number of ZKP reviewers, and the required consensus threshold for ZKP verification.
   - **Bid Submission:** Operators submit bids in a ZK manner, including an encrypted bid amount and a ZKP proving their qualifications meet the audit criteria.
   - **ZKP Reviewer Selection:** Randomly selects (or selects based on reputation) a subset of operators registered as ZKP reviewers for each audit request.
   - **Verification Results Aggregation:** Collects and aggregates the verification results (the best bit) submitted by the ZKP reviewers.
   - **Consensus Determination:** Applies a consensus mechanism (e.g., majority vote) to determine the overall validity of each ZKP based on the aggregated results.
   - **Bid Selection:** Selects the winning bid based on predefined criteria (e.g., lowest valid bid, reputation, or a combination).
   - **Funds Escrow:** Manages the escrow of USDC funds from the user, releasing payment to the winning auditor upon successful audit completion.
2. **BidVerifier:**
   - **ZKP Distribution:** Facilitates the secure distribution of ZKP bids to the selected ZKP reviewers.
   - **Result Validation:** Validates the format and signatures of the verification results submitted by the reviewers.
3. **ReportVerifier:**
   - **Report Assessment:** Checks the audit report submitted by the winning auditor against the initial audit request criteria.
   - **Additional Checks:** May perform further analysis, such as automated code checks or manual review, depending on the complexity and requirements of the audit.
   - **Success Determination:** Evaluates the audit report and additional checks to determine if the audit was successful, triggering payment release from escrow if so.
4. **OperatorRegistry:**
   - **Operator Registration:** Maintains a registry of operators who have opted into the AVS system, differentiating between two types:
     - **Auditors:** Perform the actual smart contract audits.
     - **ZKP Reviewers:** Specialize in verifying the ZKPs submitted by auditors.
   - **Reputation Tracking:** Tracks the reputation or performance metrics (e.g., successful audits, timely submissions, accurate ZKP verifications) of both types of operators, potentially using a separate ReputationManager contract.
   - **Stake Management:** May require operators to stake a certain amount of tokens upon registration as a form of commitment and to incentivize honest behavior.

**Optional Contracts:**

- **ReputationManager:** (If applicable) Tracks and manages the reputation of both auditors and ZKP reviewers, incorporating various performance metrics.
- **DisputeResolution:** (If applicable) Provides a mechanism for resolving disputes between users, auditors, and ZKP reviewers regarding audit results, ZKP validity, or payment issues.

**Additional Notes:**

- **ServiceManager:** This contract remains a placeholder for potential future implementations, such as slashing mechanisms for operator misconduct.
- **EigenLayer Integration:** If desired, additional contracts can be implemented to integrate with EigenLayer, enabling restaking and rewards for operators.
- **Payment Mechanism:** The exact details of the payment release process (e.g., full payment upon completion, milestones, installment-based) should be defined in the AuditRequestManager contract.

## Phased Development

### Phase 1

- **Add Bids and Request Together**: In this phase, the bid submission process is integrated with the initial audit request.
- **Only ZKP for Bid Value Verification**: Zero-Knowledge Proofs are used solely for verifying the bid values without revealing them.
- **Random Selection**: Auditors and ZKP Reviewers are selected randomly from the pool of available participants.
  Frequency-Based Aggregation: Aggregation of results is based on the frequency of verified values.

```mermaid
flowchart TD
    subgraph Phase1
        U[Users]
        A[Auditors]
        Z[ZKP Reviewers]

        U -->|Submit Audit Requests| A
        A -->|Submit Bids with ZKPs| Z
        Z -->|Verify ZKPs| A
        A -->|Submit Verified Bids| A
        A -->|Most Frequent Bid Aggregation| U
        U -->|Audit Requests Submitted| AVS_Platform
        AVS_Platform -->|Random Selection| A
        AVS_Platform -->|Random Selection| Z
    end
```

## Phase 2

- **Separation of Bids and Request**: The bid submission process is decoupled from the initial audit request, allowing for more flexibility and better management.
- **Allow node operators to be bidders and add stake**: Node operators can participate as bidders and stake their assets, increasing their potential rewards for successful audits.
- **Off-Chain Selection**: The selection process for auditors and ZKP Reviewers is moved off-chain to improve efficiency and scalability.

## Phase 3

- **Audit Execution and Verification**: This phase focuses on the comprehensive execution of the audit and verification of the audit reports, ensuring thoroughness and adherence to the specified criteria.

```mermaid
flowchart TD
    subgraph Phase3
        U[Users]
        A[Auditors]
        Z[ZKP Reviewers]
        AVS[AVS Platform]

        U -->|Submit Audit Requests| AVS
        AVS -->|Select Auditors and ZKP Reviewers| A
        AVS -->|Select Auditors and ZKP Reviewers| Z
        A -->|Conduct Audits| AVS
        A -->|Submit Audit Reports| AVS
        AVS -->|Verify Reports| Z
        Z -->|Verified Reports| AVS
        AVS -->|Release Funds| A
        AVS -->|Provide Reports| U
    end
```

## ZKP Implementation for Operator Bidding

The goal is to allow auditors to submit bids privately, proving they meet the audit criteria without revealing their actual bid amount. This is achieved through the use of Zero-Knowledge Proofs (ZKPs) and the involvement of ZKP Reviewers.

**1. Circuit Design:**

- A ZKP circuit needs to be designed using a framework like ZoKrates, Circom, or SnarkJS. This circuit encapsulates the logic for verifying an auditor's qualifications and bid without disclosing the bid amount itself.
- **Private Inputs:**
  - **Auditor's Bid:** The actual USDC amount the auditor is bidding.
  - **Auditor's Credentials:** Evidence demonstrating the auditor's experience, qualifications, past audits, etc. (The exact format depends on the chosen criteria).
- **Public Inputs:**
  - **Minimum Criteria Hash:** A hash of the minimum qualifications required for the specific audit request.

**2. Proof Generation:**

- The auditor generates a ZKP using the designed circuit and their private inputs. This proof cryptographically attests that:
  - The auditor's bid meets or exceeds the minimum criteria.
  - The auditor's credentials are valid (without revealing the specifics).

**3. Bid Submission:**

- The auditor submits the following to the AuditRequestManager contract:
  - **Encrypted Bid:** The bid amount, encrypted using a secure method like RSA or Elliptic Curve Cryptography, is sent along with the ZKP to the AuditRequestManager contract..
  - **Zero-Knowledge Proof:** The proof generated in step 2.

**4. ZKP Distribution and Verification:**

- The AuditRequestManager contract selects a predetermined number of ZKP Reviewers.
- The BidVerifier contract securely distributes the auditor's ZKP to the selected ZKP Reviewers.
- Each ZKP Reviewer independently verifies the ZKP against the public inputs (minimum criteria hash) to confirm its validity without revealing the bid amount.
- Reviewers submit their verification results (pass/fail) to the AuditRequestManager.

**5. Consensus and Bid Selection:**

- The AuditRequestManager contract aggregates the verification results from the ZKP Reviewers.
- If the ZKP receives sufficient passing verifications based on the predefined consensus threshold, the bid is considered valid.
- Valid bids proceed to the bid selection stage, where the winning auditor is chosen based on predefined criteria (lowest valid bid, reputation, or a combination).

**Benefits of Using ZKPs with Reviewers:**

- **Enhanced Privacy:** Maintains the confidentiality of bid amounts, preventing potential collusion or undercutting among auditors.
- **Decentralized Verification:** Distributes the responsibility of ZKP verification across multiple parties, reducing the reliance on a single entity and enhancing trust in the system.
- **Specialized Expertise:** Allows for the involvement of ZKP experts who can rigorously assess the validity of the proofs.
- **Fairness and Transparency:** Ensures bids are evaluated solely on merit (qualifications) and promotes transparency in the selection process.

**Challenges and Considerations:**

- **Complexity:** Designing and implementing ZKP circuits and coordinating the verification process across multiple reviewers can be complex.
- **Performance:** ZKP generation and verification can be computationally intensive, requiring optimizations for scalability.
- **Incentive Alignment:** Ensuring that ZKP Reviewers are incentivized to participate honestly and accurately is crucial. This could involve rewards (tokens, reputation) or penalties for incorrect verifications.

TODO: Add the Circuit playground

## Further improvements

- Point 1

```

```

```

```
