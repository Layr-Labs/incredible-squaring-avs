// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "./interface/IBidVerifier.sol"; // Assuming a separate ZKP verification contract
import "./interface/IReportVerifier.sol"; // Assuming a separate report verification contract
import "@openzeppelin/contracts/token/ERC20/IERC20.sol"; // Import the ERC20 interface

contract AuditRequestManager {
    // Events
    event AuditRequested(
        uint256 requestId,
        address user,
        uint256 budget,
        uint256 deadline
    );
    event BidSubmitted(
        uint256 requestId,
        address auditor,
        bytes encryptedBid,
        bytes zkp
    );
    event ZKPReviewersSelected(uint256 requestId, address[] reviewers);
    event ZKPVerificationResult(
        uint256 requestId,
        address reviewer,
        bool isValid
    );
    event BidStatusUpdated(uint256 requestId, address auditor, bool isValid);
    event AuditorSelected(uint256 requestId, address auditor);
    event AuditReportSubmitted(
        uint256 requestId,
        address auditor,
        string reportURI
    );
    event AuditCompleted(uint256 requestId, address auditor);
    event FundsReleased(uint256 requestId, address recipient, uint256 amount);

    // Data Structures
    struct AuditRequest {
        address user;
        uint256 budget;
        IERC20 paymentToken;
        uint256 deadline;
        string criteriaURI;
        uint8 requiredVerifications;
        address[] zkpReviewers;
        mapping(address => bool) zkpVerified;
        address selectedAuditor;
        string reportURI;
        bool paymentMade;
        bool completed;
    }

    // State Variables
    uint256 public requestIdCounter;
    mapping(uint256 => AuditRequest) public auditRequests;
    BidVerifier public bidVerifier;
    ReportVerifier public reportVerifier;
    // ... (potential operator registry or reputation system variables)

    // Constructor (Initialize with BidVerifier and ReportVerifier addresses)
    constructor(BidVerifier _bidVerifier, ReportVerifier _reportVerifier) {
        bidVerifier = _bidVerifier;
        reportVerifier = _reportVerifier;
    }

    // Modifiers (for access control)
    // ... (e.g., onlyUser, onlyAuditor)

    // Functions

    /**
     * @notice Creates a new audit request
     * @dev Requires a valid budget, deadline, criteria URI, and number of required ZKP verifications.
     * @param budget The budget (in the specified ERC20 token) allocated for the audit.
     * @param paymentToken The ERC20 token contract address in which the budget is denominated.
     * @param deadline The timestamp by which auditors must submit their bids (at least 3 days in the future).
     * @param criteriaURI A URI pointing to an IPFS document that defines the criteria for the audit.
     * @param requiredVerifications The minimum number of ZKP reviewers who must approve a bid for it to be considered valid.
     * @dev Payment will be transferred only after bid acceptance, using the `makePayment` function.
     */
    function createAuditRequest(
        uint256 budget,
        IERC20 paymentToken,
        uint256 deadline,
        string memory criteriaURI,
        uint8 requiredVerifications
    ) external {
        // Input Validation
        require(budget > 0, "Budget must be greater than zero");
        require(
            deadline > block.timestamp + 3 days,
            "Deadline must be at least 3 days in the future"
        );
        require(
            requiredVerifications > 0,
            "Must require at least one verification"
        );

        // Create Audit Request
        requestIdCounter++;
        auditRequests[requestIdCounter] = AuditRequest({
            user: msg.sender,
            budget: budget,
            paymentToken: paymentToken,
            deadline: deadline,
            criteriaURI: criteriaURI,
            requiredVerifications: requiredVerifications,
            zkpReviewers: new address[](0),
            selectedAuditor: address(0),
            reportURI: "",
            paymentMade: false,
            completed: false
        });

        emit AuditRequested(
            requestIdCounter,
            msg.sender,
            budget,
            deadline,
            criteriaURI
        );
    }

    /**
     * @notice Allows an auditor to submit a bid for a specific audit request.
     * @dev The bid includes an encrypted bid amount and a zero-knowledge proof (ZKP) demonstrating the auditor's qualifications.
     * @param requestId The ID of the audit request for which the bid is being submitted.
     * @param encryptedBid The bid amount (in USDC) encrypted using a secure method.
     * @param zkp The zero-knowledge proof demonstrating that the auditor meets the specified audit criteria.
     * @dev The function validates the bid submission against the request deadline and ZKP validity.
     * @dev If valid, the bid is stored and the BidSubmitted event is emitted.
     */
    function submitBid(
        uint256 requestId,
        bytes calldata encryptedBid,
        bytes calldata zkp
    ) external {
        AuditRequest storage request = auditRequests[requestId];

        // Input Validation
        require(
            block.timestamp <= request.deadline,
            "Bidding deadline has passed"
        );
        require(
            request.selectedAuditor == address(0),
            "Auditor already selected"
        );
        // ... (Additional validation of bid format, ensuring it's not a duplicate bid, etc.)

        // ZKP Verification
        require(
            bidVerifier.verifyBid(request.criteriaHash, zkp),
            "Invalid ZKP"
        ); //@TODO: this might change as we may want to not have a verifier contract verify the bid

        // Store Bid
        // ... (You'll likely store the encryptedBid and a reference to the auditor in a data structure associated with the request)

        emit BidSubmitted(requestId, msg.sender, encryptedBid, zkp);
    }

    /**
     * @notice Allows the user who initiated the audit request to select ZKP reviewers.
     * @dev Only callable by the user who created the audit request.
     * @param requestId The ID of the audit request for which reviewers are being selected.
     * @param reviewers An array of addresses representing the selected ZKP reviewers.
     * @dev The number of reviewers must match the 'requiredVerifications' specified in the audit request.
     * @dev Additional checks for reviewer eligibility (e.g., reputation, staking) could be added here.
     */
    function selectZKPReviewers(
        uint256 requestId,
        address[] calldata reviewers
    ) external onlyUser(requestId) {
        // @TODO: to work on the onlyUser modification
        AuditRequest storage request = auditRequests[requestId];

        // Input Validation
        require(
            reviewers.length == request.requiredVerifications,
            "Incorrect number of reviewers"
        );
        // ... (Potentially add checks for reviewer eligibility, reputation, etc.)

        // Store Reviewers
        request.zkpReviewers = reviewers;

        emit ZKPReviewersSelected(requestId, reviewers);
    }

    /**
     * @notice Allows a ZKP reviewer to submit their verification result for a specific audit request.
     * @dev Only callable by a ZKP reviewer who has been selected for the given audit request.
     * @param requestId The ID of the audit request for which the verification result is being submitted.
     * @param isValid A boolean value indicating whether the ZKP was verified as valid (`true`) or invalid (`false`).
     * @dev The function checks if the caller is a valid reviewer for the specified request.
     * @dev If valid, the reviewer's verification result is stored.
     * @dev The function then checks if enough verifications have been received to reach consensus and potentially triggers bid selection.
     */
    function submitZKPVerificationResult(
        uint256 requestId,
        bool isValid
    ) external {
        // ... (check if msg.sender is a valid reviewer for the request)
        auditRequests[requestId].zkpVerified[msg.sender] = isValid;
        emit ZKPVerificationResult(requestId, msg.sender, isValid);

        // ... (check if enough verifications have been received, update bid status, and potentially select the auditor)
    }

    /**
     * @notice Allows the selected auditor to submit the audit report for a specific audit request.
     * @dev Only callable by the auditor who was selected for the given audit request.
     * @param requestId The ID of the audit request for which the report is being submitted.
     * @param reportURI A URI (e.g., IPFS hash) pointing to the location of the detailed audit report.
     * @dev The function stores the report URI and marks the audit as completed.
     */
    function submitAuditReport(
        uint256 requestId,
        string calldata reportURI
    ) external onlyAuditor(requestId) {
        /* ... */
    }

    /**
     * @notice Verifies the submitted audit report and, if valid, marks the audit as completed.
     * @dev Only callable by the designated ReportVerifier contract.
     * @param requestId The ID of the audit request for which the report is being verified.
     * @dev This function calls the external ReportVerifier contract to assess the report's validity.
     * @dev If the report is valid, it marks the audit as completed.
     * @dev If the report is invalid, it may trigger a dispute resolution process or return the funds to the user (implementation-dependent).
     */
    function verifyReportAndReleaseFunds(uint256 requestId) external {
        require(
            msg.sender == address(reportVerifier),
            "Only ReportVerifier can call this function"
        );
        AuditRequest storage request = auditRequests[requestId];
        require(request.reportURI != "", "No report submitted");
        require(!request.completed, "Audit already completed");

        bool isReportValid = reportVerifier.verifyReport(
            requestId,
            request.reportURI
        );

        if (isReportValid) {
            request.completed = true;
            payable(request.selectedAuditor).transfer(request.budget);
            emit AuditCompleted(requestId, request.selectedAuditor);
            emit FundsReleased(
                requestId,
                request.selectedAuditor,
                request.budget
            );
        } else {
            // Handle invalid report (e.g., return funds to user, allow dispute resolution)
            // ...
        }
    }

    /**
     * @notice Allows the user who initiated the audit request to make payment to the selected auditor in ERC20 tokens.
     * @dev Can only be called by the user who created the request, after an auditor has been selected, and before the audit is completed.
     * @param requestId The ID of the audit request for which the payment is being made.
     */
    function makePayment(uint256 requestId) external {
        AuditRequest storage request = auditRequests[requestId];

        // .... (Input validation from the previous response)

        // Transfer ERC20 Tokens to Auditor (Directly)
        require(
            request.paymentToken.transferFrom(
                msg.sender,
                request.selectedAuditor,
                request.budget
            ),
            "Token transfer failed"
        );

        request.paymentMade = true; // Mark payment as made

        emit PaymentMade(requestId, msg.sender, request.budget);
    }
}
