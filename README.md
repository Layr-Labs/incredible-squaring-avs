# Smart Contract Auditor AVS

## Project Overview

This AVS is designed to facilitate secure, efficient, and transparent smart contract auditing using Zero-Knowledge Proofs (ZKPs) for bid submissions and verifications. This README provides an overview of the AVS platform's design and architecture, outlining the roles of operators and auditors, the phases of development, and the use of ZKPs to ensure a secure and fair auditing process.

The platform aims to enhance the quality of smart contract audits by:

1. ensuring a decentralized process of verifying and selecting the best auditors.
2. enabling auditors to stake their own assets to earn greater rewards for their audits.

TODO: Add logic diagram

```mermaid
graph TD
    subgraph Operator Registration
        A[OperatorRegistry]
        A -->|Registers Auditors| B[Auditors]
        A -->|Registers ZKP Reviewers| C[ZKP Reviewers]
        B -->|Stake Tokens & Prove Qualifications| A
        C -->|Stake Tokens & Prove Qualifications| A
    end

    subgraph Bid Preparation & Submission
        B -->|Prepare Bids (Price + ZKP)| D[AuditRequestManager]
        D -->|Submit Bids (Price + ZKP)| E[Bid Submission]
    end

    subgraph ZKP Review
        D -->|Distribute ZKPs| F[BidVerifier]
        F -->|Verify ZKPs| C
        C -->|Submit Verification Results| F
    end

    subgraph Bid Selection
        F -->|Aggregate Verification Results| G[AuditRequestManager]
        G -->|Select Valid Bid| H[Selected Auditor]
    end

    subgraph Audit Execution & Report
        H -->|Conduct Audit| I[Smart Contract]
        H -->|Generate Audit Report| J[ReportVerifier]
        J -->|Submit Audit Report| K[Report Verification]
    end

    subgraph Report Verification & Payment
        J -->|Verify Report| L[Payment Release]
        L -->|Release Payment| H
    end

    style A fill:#f9f,stroke:#333,stroke-width:4px
    style D fill:#bbf,stroke:#333,stroke-width:4px
    style F fill:#ddf,stroke:#333,stroke-width:4px
    style J fill:#bff,stroke:#333,stroke-width:4px
```

## Workflows

TODO: Add interaction diagram

### Smart Contract Audit Requests

### Bidding for Audit Requests

Operators register with the AVS platform by opting into the OperatorRegistry contract. They must specify their desired roles (Auditor and/or ZKP Reviewer) and may need to stake a certain amount of tokens. Proof of qualifications or experience is required for eligibility

- Price Determination: Auditors determine their proposed fee for the audit in USDC.
- Zero-Knowledge Proof Generation: Auditors create a ZKP proving their qualifications without revealing the actual bid amount.
- Bid Submission: Auditors submit their bid (price and ZKP) to the AuditRequestManager contract.

TODO: Add interaction diagram between contracts

### Verification of Bids

Operators verify and determine if a bid's ZKP meets the required consensus.

TODO: Add diagram for Zkp verification

### Aggregation of Bids

### Winner Selection

The selected auditor conducts a comprehensive audit of the smart contract code according to specified criteria and industry best practices.

TODO: Add phase 1 system diagram

### Report Generation and Submission

Upon completing the audit, the auditor generates a detailed report summarizing their findings, potential vulnerabilities, and recommendations, then submits it to the ReportVerifier contract.

### Report Verification and Payment

The ReportVerifier contract assesses the audit report and releases the escrowed USDC payment to the auditor if it meets the required standards.

### Slashing on Hack Detection

## Phased Development

### Phase 1

Current

- add bids and request together
- only zkp for bid value verification
- random selection
- frequency-based aggregation

TODO: Add diagram

## Phase 2

- separation of bids and request
- allow node operators to be bidders; add stake
- off-chain selection

TODO: Add diagram

## Phase 3

- Audit Execution and Verification

### Report Generation and Submission

## Zero Knowledge Module

A ZKP circuit has implemented using `gnark`. This circuit verifies an auditor's bid without disclosing the bid amount.

TODO: Add the Circuit playground

## Further improvements
