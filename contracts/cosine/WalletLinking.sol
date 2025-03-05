// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @title WalletLinking
 * @dev Contract for linking L1 addresses to COSINE L2 wallets
 * Enforces one-to-one mapping and proof of ownership
 * As described in the COSINE whitepaper sections 6 and 7.7
 */
contract WalletLinking is AccessControl, Pausable {
    using ECDSA for bytes32;

    // Roles
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant OPERATOR_ROLE = keccak256("OPERATOR_ROLE");

    // Mapping from L1 address to L2 wallet
    mapping(address => bytes32) public l1ToL2Mapping;

    // Mapping from L2 wallet to L1 addresses on various chains
    // chainId => L2 wallet => L1 address
    mapping(uint256 => mapping(bytes32 => address)) public l2ToL1Mapping;

    // Chain identifiers for supported blockchains
    uint256 public constant ETHEREUM_CHAIN_ID = 1;
    uint256 public constant BSC_CHAIN_ID = 56;
    uint256 public constant SOLANA_CHAIN_ID = 501; // Example ID for Solana

    // Verification challenge expiration time (in seconds)
    uint256 public constant CHALLENGE_EXPIRATION = 1 hours;

    // Structure to store verification challenges
    struct Challenge {
        bytes32 challengeHash;
        uint256 timestamp;
        bool isUsed;
    }

    // Mapping to store challenges for each address
    mapping(address => Challenge) public challenges;

    // Events
    event ChallengeCreated(
        address indexed l1Address,
        bytes32 challengeHash,
        uint256 expiryTime
    );
    event WalletLinked(
        address indexed l1Address,
        bytes32 indexed l2Wallet,
        uint256 chainId
    );
    event WalletLinkingRevoked(
        address indexed l1Address,
        bytes32 indexed l2Wallet,
        uint256 chainId
    );

    /**
     * @dev Constructor sets up the default admin role
     */
    constructor() {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ADMIN_ROLE, msg.sender);
        _setupRole(OPERATOR_ROLE, msg.sender);
    }

    /**
     * @dev Generate a challenge for a wallet to prove ownership
     * @param l1Address The L1 address requesting the challenge
     * @return challengeHash The hash of the challenge to be signed
     */
    function generateChallenge(
        address l1Address
    ) external returns (bytes32 challengeHash) {
        require(l1Address != address(0), "Invalid L1 address");

        // Generate a unique challenge using the address, contract address, and current timestamp
        challengeHash = keccak256(
            abi.encodePacked(
                l1Address,
                address(this),
                block.timestamp,
                blockhash(block.number - 1)
            )
        );

        challenges[l1Address] = Challenge({
            challengeHash: challengeHash,
            timestamp: block.timestamp,
            isUsed: false
        });

        emit ChallengeCreated(
            l1Address,
            challengeHash,
            block.timestamp + CHALLENGE_EXPIRATION
        );

        return challengeHash;
    }

    /**
     * @dev Link an L1 address to an L2 wallet with proof of ownership
     * @param l2Wallet The L2 wallet identifier (bytes32 hash)
     * @param chainId The chain ID of the current blockchain
     * @param signature The signature proving ownership of the L1 address
     */
    function linkWallet(
        bytes32 l2Wallet,
        uint256 chainId,
        bytes calldata signature
    ) external whenNotPaused {
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");
        require(
            l1ToL2Mapping[msg.sender] == bytes32(0),
            "L1 address already linked"
        );
        require(
            l2ToL1Mapping[chainId][l2Wallet] == address(0),
            "L2 wallet already linked on this chain"
        );

        Challenge storage challenge = challenges[msg.sender];
        require(challenge.challengeHash != bytes32(0), "No challenge found");
        require(!challenge.isUsed, "Challenge already used");
        require(
            block.timestamp - challenge.timestamp <= CHALLENGE_EXPIRATION,
            "Challenge expired"
        );

        // Verify the signature
        bytes32 ethSignedMessageHash = keccak256(
            abi.encodePacked(
                "\x19Ethereum Signed Message:\n32",
                challenge.challengeHash
            )
        );

        address recoveredAddress = ethSignedMessageHash.recover(signature);
        require(recoveredAddress == msg.sender, "Invalid signature");

        // Mark challenge as used
        challenge.isUsed = true;

        // Link the addresses
        l1ToL2Mapping[msg.sender] = l2Wallet;
        l2ToL1Mapping[chainId][l2Wallet] = msg.sender;

        emit WalletLinked(msg.sender, l2Wallet, chainId);
    }

    /**
     * @dev Revoke the linking between an L1 address and L2 wallet
     * Only the owner of the L1 address can revoke the linking
     * @param chainId The chain ID where the linking should be revoked
     */
    function revokeLinking(uint256 chainId) external whenNotPaused {
        bytes32 l2Wallet = l1ToL2Mapping[msg.sender];
        require(l2Wallet != bytes32(0), "No linking found for this L1 address");

        // Clear the mappings
        l1ToL2Mapping[msg.sender] = bytes32(0);
        l2ToL1Mapping[chainId][l2Wallet] = address(0);

        emit WalletLinkingRevoked(msg.sender, l2Wallet, chainId);
    }

    /**
     * @dev Administrative function to force revoke a linking in case of disputes
     * @param l1Address The L1 address to unlink
     * @param chainId The chain ID where the linking should be revoked
     */
    function adminRevokeLink(
        address l1Address,
        uint256 chainId
    ) external onlyRole(ADMIN_ROLE) {
        bytes32 l2Wallet = l1ToL2Mapping[l1Address];
        require(l2Wallet != bytes32(0), "No linking found for this L1 address");

        // Clear the mappings
        l1ToL2Mapping[l1Address] = bytes32(0);
        l2ToL1Mapping[chainId][l2Wallet] = address(0);

        emit WalletLinkingRevoked(l1Address, l2Wallet, chainId);
    }

    /**
     * @dev Verify if an L1 address is linked to a specific L2 wallet
     * @param l1Address The L1 address to check
     * @param l2Wallet The L2 wallet to verify against
     * @param chainId The chain ID to check on
     * @return isLinked Whether the addresses are linked
     */
    function verifyLinking(
        address l1Address,
        bytes32 l2Wallet,
        uint256 chainId
    ) external view returns (bool isLinked) {
        return
            l1ToL2Mapping[l1Address] == l2Wallet &&
            l2ToL1Mapping[chainId][l2Wallet] == l1Address;
    }

    /**
     * @dev Get the L2 wallet linked to an L1 address
     * @param l1Address The L1 address to query
     * @return l2Wallet The linked L2 wallet
     */
    function getLinkedL2Wallet(
        address l1Address
    ) external view returns (bytes32 l2Wallet) {
        return l1ToL2Mapping[l1Address];
    }

    /**
     * @dev Get the L1 address linked to an L2 wallet on a specific chain
     * @param l2Wallet The L2 wallet to query
     * @param chainId The chain ID to query
     * @return l1Address The linked L1 address
     */
    function getLinkedL1Address(
        bytes32 l2Wallet,
        uint256 chainId
    ) external view returns (address l1Address) {
        return l2ToL1Mapping[chainId][l2Wallet];
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
