// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title ICreditVerification
 * @dev Interface for the credit verification contract
 */
interface ICreditVerification {
    struct Vector2D {
        int256 x;
        int256 y;
    }
    
    function verifyCreditScore(
        bytes32 l2Wallet,
        Vector2D calldata thresholdVector,
        int256 thresholdSimilarity
    ) external returns (int256 similarity, bool passed);
    
    function provideTransactionFeedback(bytes32 l2Wallet, bool isTrusted) external;
    function calculateVerificationFee(bytes32 l2Wallet) external view returns (uint256 fee);
    function calculateCosineSimilarity(Vector2D memory v1, Vector2D memory v2) external pure returns (int256 similarity);
    function checkFeedbackRequirement(address verifier) external view returns (bool requiresFeedback);
}