// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title IWalletLinking
 * @dev Interface for the wallet linking contract
 */
interface IWalletLinking {
    function generateChallenge(address l1Address) external returns (bytes32 challengeHash);
    function linkWallet(bytes32 l2Wallet, uint256 chainId, bytes calldata signature) external;
    function revokeLinking(uint256 chainId) external;
    function verifyLinking(address l1Address, bytes32 l2Wallet, uint256 chainId) external view returns (bool isLinked);
    function getLinkedL2Wallet(address l1Address) external view returns (bytes32 l2Wallet);
    function getLinkedL1Address(bytes32 l2Wallet, uint256 chainId) external view returns (address l1Address);
}