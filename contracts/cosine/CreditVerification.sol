// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "../interfaces/ICosineToken.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./WalletLinking.sol";

/**
 * @title CreditVerification
 * @dev Contract for verifying COSINE credit scores and collecting verification fees
 * Implements cosine similarity checks as described in Sections 4.8 and 7.5
 */
contract CreditVerification is AccessControl, Pausable {
    using SafeMath for uint256;

    // Roles
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VALIDATOR_ROLE = keccak256("VALIDATOR_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    // Reference to the COSINE token contract
    ICosineToken public cosineToken;

    // Reference to the wallet linking contract
    WalletLinking public walletLinking;

    // Fee multiplier (kappa) as defined in Section 10.1
    uint256 public feeMultiplier = 120; // 1.2 as 120/100

    // Minimum verification fee
    uint256 public minVerificationFee = 1 * 10 ** 16; // 0.01 COSINE

    // Accumulated cost for each wallet
    mapping(bytes32 => uint256) public accumulatedCost;

    // Last verification timestamp for each wallet
    mapping(bytes32 => uint256) public lastVerificationTime;

    // Trust/No-Trust votes after transactions
    mapping(address => mapping(bytes32 => bool)) public postTransactionFeedback;

    // Whether an address has provided feedback after verification
    mapping(address => bool) public hasFeedbackPending;

    // Tracks the wallet that a verifier last checked
    mapping(address => bytes32) public lastCheckedWallet;

    // Events
    event CreditVerified(
        address indexed verifier,
        bytes32 indexed l2Wallet,
        uint256 fee,
        bool passedThreshold
    );
    event FeedbackProvided(
        address indexed provider,
        bytes32 indexed l2Wallet,
        bool isTrusted
    );
    event AccumulatedCostReset(bytes32 indexed l2Wallet, uint256 newCost);
    event FeeMultiplierUpdated(uint256 newMultiplier);

    // Struct to represent vector data for cosine similarity
    struct Vector2D {
        int256 x;
        int256 y;
    }

    /**
     * @dev Constructor sets up roles and references to other contracts
     * @param cosineTokenAddress Address of the COSINE token contract
     * @param walletLinkingAddress Address of the wallet linking contract
     */
    constructor(address cosineTokenAddress, address walletLinkingAddress) {
        require(cosineTokenAddress != address(0), "Invalid token address");
        require(
            walletLinkingAddress != address(0),
            "Invalid wallet linking address"
        );

        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ADMIN_ROLE, msg.sender);
        _setupRole(VALIDATOR_ROLE, msg.sender);

        cosineToken = ICosineToken(cosineTokenAddress);
        walletLinking = WalletLinking(walletLinkingAddress);
    }

    /**
     * @dev Verify a wallet's credit score using cosine similarity
     * @param l2Wallet The L2 wallet to verify
     * @param thresholdVector The threshold vector to compare against
     * @param thresholdSimilarity The minimum similarity to be considered "passing"
     * @return similarity The calculated cosine similarity
     * @return passed Whether the similarity exceeded the threshold
     */
    function verifyCreditScore(
        bytes32 l2Wallet,
        Vector2D calldata thresholdVector,
        int256 thresholdSimilarity
    ) external whenNotPaused returns (int256 similarity, bool passed) {
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");

        // Calculate fee based on accumulated cost
        uint256 fee = calculateVerificationFee(l2Wallet);

        // Collect the fee (burn it as per whitepaper section 10.2)
        require(
            cosineToken.transferFrom(msg.sender, address(this), fee),
            "Fee transfer failed"
        );
        cosineToken.burn(fee);

        // Record that this address needs to provide feedback before next verification
        hasFeedbackPending[msg.sender] = true;
        lastCheckedWallet[msg.sender] = l2Wallet;

        // Get credit score vector from validator nodes
        // In a real implementation, this would involve cross-communication with L2
        // For this contract, we'll assume validators submit this data separately
        Vector2D memory walletVector = getWalletVector(l2Wallet);

        // Calculate cosine similarity
        similarity = calculateCosineSimilarity(walletVector, thresholdVector);
        passed = similarity >= thresholdSimilarity;

        // Reset accumulated cost after verification
        accumulatedCost[l2Wallet] = 0;
        lastVerificationTime[l2Wallet] = block.timestamp;

        emit CreditVerified(msg.sender, l2Wallet, fee, passed);

        return (similarity, passed);
    }

    /**
     * @dev Provide post-transaction feedback about a wallet
     * @param l2Wallet The L2 wallet to provide feedback for
     * @param isTrusted Whether the transaction with this wallet was trustworthy
     */
    function provideTransactionFeedback(
        bytes32 l2Wallet,
        bool isTrusted
    ) external whenNotPaused {
        // Ensure the wallet is valid and was the last one checked by this address
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");
        require(
            hasFeedbackPending[msg.sender],
            "No pending feedback requirement"
        );
        require(
            lastCheckedWallet[msg.sender] == l2Wallet,
            "Feedback must be for last checked wallet"
        );

        // Store the feedback
        postTransactionFeedback[msg.sender][l2Wallet] = isTrusted;

        // Clear the pending feedback flag
        hasFeedbackPending[msg.sender] = false;

        emit FeedbackProvided(msg.sender, l2Wallet, isTrusted);
    }

    /**
     * @dev Calculate the verification fee based on accumulated cost
     * @param l2Wallet The L2 wallet to calculate fee for
     * @return fee The calculated fee
     */
    function calculateVerificationFee(
        bytes32 l2Wallet
    ) public view returns (uint256 fee) {
        uint256 cost = accumulatedCost[l2Wallet];
        if (cost == 0) {
            return minVerificationFee;
        }

        return cost.mul(feeMultiplier).div(100);
    }

    /**
     * @dev Calculate cosine similarity between two 2D vectors
     * Uses fixed-point arithmetic with 18 decimals precision
     * @param v1 The first vector
     * @param v2 The second vector
     * @return similarity The cosine similarity scaled by 10^18
     */
    function calculateCosineSimilarity(
        Vector2D memory v1,
        Vector2D memory v2
    ) public pure returns (int256 similarity) {
        // Calculate dot product: v1.x * v2.x + v1.y * v2.y
        int256 dotProduct = (v1.x * v2.x + v1.y * v2.y) / 10 ** 18;

        // Calculate magnitudes: sqrt(v1.x^2 + v1.y^2) * sqrt(v2.x^2 + v2.y^2)
        int256 v1MagnitudeSq = (v1.x * v1.x + v1.y * v1.y) / 10 ** 18;
        int256 v2MagnitudeSq = (v2.x * v2.x + v2.y * v2.y) / 10 ** 18;

        // Simplified approach for fixed-point sqrt based on v1MagnitudeSq and v2MagnitudeSq
        // A more accurate implementation would use a proper fixed-point sqrt function
        int256 v1Magnitude = sqrt(v1MagnitudeSq);
        int256 v2Magnitude = sqrt(v2MagnitudeSq);

        // Avoid division by zero
        if (v1Magnitude == 0 || v2Magnitude == 0) {
            return 0;
        }

        // Calculate similarity
        similarity = (dotProduct * 10 ** 18) / (v1Magnitude * v2Magnitude);

        // Ensure result is in [-10^18, 10^18] range
        if (similarity > 10 ** 18) {
            similarity = 10 ** 18;
        } else if (similarity < -10 ** 18) {
            similarity = -10 ** 18;
        }

        return similarity;
    }

    /**
     * @dev Simplified fixed-point square root function
     * @param x The input value (scaled by 10^18)
     * @return y The square root of x (scaled by 10^18)
     */
    function sqrt(int256 x) internal pure returns (int256 y) {
        if (x <= 0) return 0;

        // Handle x as if it were uint256 for simplicity
        uint256 z = uint256(x);

        // Initial estimate (Newton's method)
        uint256 result = z;
        uint256 temp = (z / 2) + 1;

        while (temp < result) {
            result = temp;
            temp = (z / temp + temp) / 2;
        }

        return int256(result);
    }

    /**
     * @dev Update accumulated cost for a wallet (called by validators)
     * @param l2Wallet The L2 wallet to update
     * @param additionalCost The additional cost to add
     */
    function updateAccumulatedCost(
        bytes32 l2Wallet,
        uint256 additionalCost
    ) external onlyRole(VALIDATOR_ROLE) {
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");

        accumulatedCost[l2Wallet] = accumulatedCost[l2Wallet].add(
            additionalCost
        );

        emit AccumulatedCostReset(l2Wallet, accumulatedCost[l2Wallet]);
    }

    /**
     * @dev Get wallet vector for cosine similarity calculation
     * This would involve communication with L2 validators in a real implementation
     * For this contract, we'll just provide a mock implementation
     * @param l2Wallet The L2 wallet to get vector for
     * @return walletVector The normalized wallet vector
     */
    function getWalletVector(
        bytes32 l2Wallet
    ) internal view returns (Vector2D memory walletVector) {
        // This is a mock implementation
        // In a real implementation, this would retrieve data from validators

        // For testing purposes, derive a deterministic vector from the wallet hash
        int256 xComponent = int256(uint256(l2Wallet) % 1000000) * 10 ** 15; // Scaled by 10^18
        int256 yComponent = 10 ** 18; // Constant y component (1.0)

        return Vector2D({x: xComponent, y: yComponent});
    }

    /**
     * @dev Submit wallet vector data (called by validators)
     * @param l2Wallet The L2 wallet the vector belongs to
     * @param vectorX The x component of the vector (normalized score)
     * @param vectorY The y component of the vector (typically 1.0)
     */
    function submitWalletVector(
        bytes32 l2Wallet,
        int256 vectorX,
        int256 vectorY
    ) external onlyRole(VALIDATOR_ROLE) {
        // In a real implementation, this would store the vector data on-chain
        // or emit an event for off-chain processing

        // For this implementation, we'll just emit an event
        emit WalletVectorUpdated(l2Wallet, vectorX, vectorY);
    }

    /**
     * @dev Event for wallet vector updates
     */
    event WalletVectorUpdated(
        bytes32 indexed l2Wallet,
        int256 vectorX,
        int256 vectorY
    );

    /**
     * @dev Update the fee multiplier
     * @param newMultiplier The new multiplier (scaled by 100, e.g., 120 = 1.2x)
     */
    function updateFeeMultiplier(
        uint256 newMultiplier
    ) external onlyRole(ADMIN_ROLE) {
        require(newMultiplier >= 100, "Multiplier must be at least 1.0");
        feeMultiplier = newMultiplier;
        emit FeeMultiplierUpdated(newMultiplier);
    }

    /**
     * @dev Update the minimum verification fee
     * @param newMinFee The new minimum fee
     */
    function updateMinimumFee(uint256 newMinFee) external onlyRole(ADMIN_ROLE) {
        minVerificationFee = newMinFee;
    }

    /**
     * @dev Check if an address has a feedback requirement pending
     * @param verifier The address to check
     * @return requiresFeedback Whether the address needs to provide feedback
     */
    function checkFeedbackRequirement(
        address verifier
    ) external view returns (bool requiresFeedback) {
        return hasFeedbackPending[verifier];
    }

    /**
     * @dev Pause the contract
     */
    function pause() external onlyRole(ADMIN_ROLE) {
        _pause();
    }

    /**
     * @dev Unpause the contract
     */
    function unpause() external onlyRole(ADMIN_ROLE) {
        _unpause();
    }
}
