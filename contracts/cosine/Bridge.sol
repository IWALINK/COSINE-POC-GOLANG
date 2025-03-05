// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./CosineToken.sol";
import "./WalletLinking.sol";

/**
 * @title Bridge
 * @dev Contract for bridging COSINE tokens across different chains
 * Implements lock-and-mint / burn-and-redeem bridging as described in Section 5.7
 */
contract Bridge is AccessControl, Pausable {
    using SafeMath for uint256;

    // Roles
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VALIDATOR_ROLE = keccak256("VALIDATOR_ROLE");
    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    // Reference to the COSINE token contract
    CosineToken public cosineToken;

    // Reference to the wallet linking contract
    WalletLinking public walletLinking;

    // Chain ID of this blockchain
    uint256 public immutable chainId;

    // Bridge fee (in basis points, 100 = 1%)
    uint256 public bridgeFee = 25; // 0.25%

    // Minimum and maximum bridge amounts
    uint256 public minBridgeAmount = 10 * 10 ** 18; // 10 COSINE
    uint256 public maxBridgeAmount = 1_000_000 * 10 ** 18; // 1M COSINE

    // Required confirmations for a bridge transaction
    uint256 public requiredConfirmations = 2;

    // Bridging operation ID counter
    uint256 public nextOperationId = 1;

    // Structure for bridge operations
    struct BridgeOperation {
        address initiator;
        bytes32 l2Wallet;
        uint256 sourceChainId;
        uint256 destinationChainId;
        uint256 amount;
        uint256 fee;
        uint256 timestamp;
        bool isLock; // true = lock (outgoing), false = redeem (incoming)
        bool isCompleted;
        mapping(address => bool) validatorConfirmations;
        uint256 confirmationCount;
    }

    // Mapping from operation ID to BridgeOperation
    mapping(uint256 => BridgeOperation) public operations;

    // Mapping to track total tokens bridged to each chain
    mapping(uint256 => uint256) public totalBridgedTokens;

    // Events
    event BridgeInitiated(
        uint256 indexed operationId,
        address indexed initiator,
        bytes32 l2Wallet,
        uint256 sourceChainId,
        uint256 destinationChainId,
        uint256 amount,
        uint256 fee,
        bool isLock
    );

    event OperationConfirmed(
        uint256 indexed operationId,
        address indexed validator,
        uint256 confirmations
    );

    event BridgeCompleted(
        uint256 indexed operationId,
        address indexed initiator,
        uint256 amount,
        bool isLock
    );

    event BridgeFeeUpdated(uint256 newFee);

    /**
     * @dev Constructor sets up roles and references to other contracts
     * @param _cosineToken Address of the COSINE token contract
     * @param _walletLinking Address of the wallet linking contract
     * @param _chainId The chain ID of this blockchain
     */
    constructor(
        address _cosineToken,
        address _walletLinking,
        uint256 _chainId
    ) {
        require(_cosineToken != address(0), "Invalid token address");
        require(_walletLinking != address(0), "Invalid wallet linking address");
        require(_chainId > 0, "Invalid chain ID");

        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ADMIN_ROLE, msg.sender);
        _setupRole(VALIDATOR_ROLE, msg.sender);
        _setupRole(RELAYER_ROLE, msg.sender);

        cosineToken = CosineToken(_cosineToken);
        walletLinking = WalletLinking(_walletLinking);
        chainId = _chainId;
    }

    /**
     * @dev Initiate a bridge operation to lock tokens on this chain
     * @param l2Wallet The COSINE L2 wallet
     * @param destinationChainId The destination chain ID
     * @param amount The amount of tokens to bridge
     */
    function initiateOutgoingBridge(
        bytes32 l2Wallet,
        uint256 destinationChainId,
        uint256 amount
    ) external whenNotPaused {
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");
        require(destinationChainId != chainId, "Cannot bridge to same chain");
        require(amount >= minBridgeAmount, "Amount below minimum");
        require(amount <= maxBridgeAmount, "Amount above maximum");

        // Verify wallet linking
        address l1Address = msg.sender;
        bool isLinked = walletLinking.verifyLinking(
            l1Address,
            l2Wallet,
            chainId
        );
        require(isLinked, "L1 address not linked to L2 wallet");

        // Calculate fee
        uint256 fee = amount.mul(bridgeFee).div(10000);
        uint256 amountAfterFee = amount.sub(fee);

        // Transfer tokens from user to bridge contract
        require(
            cosineToken.transferFrom(l1Address, address(this), amount),
            "Token transfer failed"
        );

        // Create bridge operation
        uint256 operationId = nextOperationId++;
        BridgeOperation storage operation = operations[operationId];
        operation.initiator = l1Address;
        operation.l2Wallet = l2Wallet;
        operation.sourceChainId = chainId;
        operation.destinationChainId = destinationChainId;
        operation.amount = amountAfterFee;
        operation.fee = fee;
        operation.timestamp = block.timestamp;
        operation.isLock = true;
        operation.isCompleted = false;
        operation.confirmationCount = 0;

        // Update total bridged tokens
        totalBridgedTokens[destinationChainId] = totalBridgedTokens[
            destinationChainId
        ].add(amountAfterFee);

        emit BridgeInitiated(
            operationId,
            l1Address,
            l2Wallet,
            chainId,
            destinationChainId,
            amountAfterFee,
            fee,
            true
        );
    }

    /**
     * @dev Initiate a bridge operation to redeem tokens on this chain (called by relayers)
     * @param l1Address The recipient's L1 address
     * @param l2Wallet The COSINE L2 wallet
     * @param sourceChainId The source chain ID
     * @param amount The amount of tokens to bridge
     */
    function initiateIncomingBridge(
        address l1Address,
        bytes32 l2Wallet,
        uint256 sourceChainId,
        uint256 amount
    ) external onlyRole(RELAYER_ROLE) whenNotPaused {
        require(l1Address != address(0), "Invalid L1 address");
        require(l2Wallet != bytes32(0), "Invalid L2 wallet");
        require(sourceChainId != chainId, "Cannot bridge from same chain");
        require(amount >= minBridgeAmount, "Amount below minimum");

        // Verify wallet linking
        bool isLinked = walletLinking.verifyLinking(
            l1Address,
            l2Wallet,
            chainId
        );
        require(isLinked, "L1 address not linked to L2 wallet");

        // Create bridge operation
        uint256 operationId = nextOperationId++;
        BridgeOperation storage operation = operations[operationId];
        operation.initiator = l1Address;
        operation.l2Wallet = l2Wallet;
        operation.sourceChainId = sourceChainId;
        operation.destinationChainId = chainId;
        operation.amount = amount;
        operation.fee = 0; // Fee was already taken on the source chain
        operation.timestamp = block.timestamp;
        operation.isLock = false;
        operation.isCompleted = false;
        operation.confirmationCount = 0;

        emit BridgeInitiated(
            operationId,
            l1Address,
            l2Wallet,
            sourceChainId,
            chainId,
            amount,
            0,
            false
        );
    }

    /**
     * @dev Confirm a bridge operation (called by validators)
     * @param operationId The operation ID to confirm
     */
    function confirmOperation(
        uint256 operationId
    ) external onlyRole(VALIDATOR_ROLE) whenNotPaused {
        BridgeOperation storage operation = operations[operationId];

        require(operation.initiator != address(0), "Operation does not exist");
        require(!operation.isCompleted, "Operation already completed");
        require(
            !operation.validatorConfirmations[msg.sender],
            "Already confirmed by this validator"
        );

        // Add confirmation
        operation.validatorConfirmations[msg.sender] = true;
        operation.confirmationCount = operation.confirmationCount.add(1);

        emit OperationConfirmed(
            operationId,
            msg.sender,
            operation.confirmationCount
        );

        // If enough confirmations, complete the operation
        if (operation.confirmationCount >= requiredConfirmations) {
            completeOperation(operationId);
        }
    }

    /**
     * @dev Complete a bridge operation after sufficient confirmations
     * @param operationId The operation ID to complete
     */
    function completeOperation(uint256 operationId) internal {
        BridgeOperation storage operation = operations[operationId];

        require(!operation.isCompleted, "Operation already completed");
        require(
            operation.confirmationCount >= requiredConfirmations,
            "Not enough confirmations"
        );

        operation.isCompleted = true;

        if (operation.isLock) {
            // For lock operations, tokens are already in this contract
            // Notify the token contract about the lock
            cosineToken.bridge(
                operation.destinationChainId,
                operation.amount,
                true
            );

            // Handle the fee (burn it or distribute to validators)
            cosineToken.burn(operation.fee);
        } else {
            // For redeem operations, mint tokens to the recipient
            cosineToken.bridge(
                operation.sourceChainId,
                operation.amount,
                false
            );

            // Update total bridged tokens
            if (
                totalBridgedTokens[operation.sourceChainId] >= operation.amount
            ) {
                totalBridgedTokens[
                    operation.sourceChainId
                ] = totalBridgedTokens[operation.sourceChainId].sub(
                    operation.amount
                );
            } else {
                totalBridgedTokens[operation.sourceChainId] = 0;
            }

            // Transfer the tokens to the recipient
            require(
                cosineToken.transfer(operation.initiator, operation.amount),
                "Token transfer failed"
            );
        }

        emit BridgeCompleted(
            operationId,
            operation.initiator,
            operation.amount,
            operation.isLock
        );
    }

    /**
     * @dev Check the status of a bridge operation
     * @param operationId The operation ID to check
     * @return exists Whether the operation exists
     * @return initiator The address that initiated the operation
     * @return amount The amount being bridged
     * @return confirmations The current number of confirmations
     * @return isCompleted Whether the operation is completed
     */
    function checkOperationStatus(
        uint256 operationId
    )
        external
        view
        returns (
            bool exists,
            address initiator,
            uint256 amount,
            uint256 confirmations,
            bool isCompleted
        )
    {
        BridgeOperation storage operation = operations[operationId];

        if (operation.initiator == address(0)) {
            return (false, address(0), 0, 0, false);
        }

        return (
            true,
            operation.initiator,
            operation.amount,
            operation.confirmationCount,
            operation.isCompleted
        );
    }

    /**
     * @dev Update the bridge fee
     * @param newFee The new fee in basis points (100 = 1%)
     */
    function updateBridgeFee(uint256 newFee) external onlyRole(ADMIN_ROLE) {
        require(newFee <= 500, "Fee cannot exceed 5%");
        bridgeFee = newFee;
        emit BridgeFeeUpdated(newFee);
    }

    /**
     * @dev Update the minimum and maximum bridge amounts
     * @param min The new minimum amount
     * @param max The new maximum amount
     */
    function updateBridgeLimits(
        uint256 min,
        uint256 max
    ) external onlyRole(ADMIN_ROLE) {
        require(min <= max, "Min must be <= max");
        minBridgeAmount = min;
        maxBridgeAmount = max;
    }

    /**
     * @dev Update the required confirmations for bridge operations
     * @param confirmations The new number of required confirmations
     */
    function updateRequiredConfirmations(
        uint256 confirmations
    ) external onlyRole(ADMIN_ROLE) {
        require(confirmations >= 1, "Must require at least 1 confirmation");
        requiredConfirmations = confirmations;
    }

    /**
     * @dev Rescue tokens accidentally sent to the contract
     * @param token The token address
     * @param amount The amount to rescue
     * @param recipient The address to send the tokens to
     */
    function rescueTokens(
        address token,
        uint256 amount,
        address recipient
    ) external onlyRole(ADMIN_ROLE) {
        require(recipient != address(0), "Cannot send to zero address");

        if (token == address(cosineToken)) {
            // For COSINE tokens, can only rescue excess tokens not accounted for in totalBridgedTokens
            uint256 contractBalance = cosineToken.balanceOf(address(this));
            uint256 trackedTokens = 0;

            // Sum up all bridged tokens
            for (uint256 i = 1; i <= 1000; i++) {
                // Limit to 1000 chains for gas reasons
                if (totalBridgedTokens[i] > 0) {
                    trackedTokens = trackedTokens.add(totalBridgedTokens[i]);
                }
            }

            require(
                contractBalance.sub(trackedTokens) >= amount,
                "Cannot rescue tracked tokens"
            );
        }

        IERC20(token).transfer(recipient, amount);
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
