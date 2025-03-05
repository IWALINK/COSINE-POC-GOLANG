// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title IBridge
 * @dev Interface for the bridge contract
 */
interface IBridge {
    function initiateOutgoingBridge(
        bytes32 l2Wallet,
        uint256 destinationChainId,
        uint256 amount
    ) external;

    function initiateIncomingBridge(
        address l1Address,
        bytes32 l2Wallet,
        uint256 sourceChainId,
        uint256 amount
    ) external;

    function confirmOperation(uint256 operationId) external;

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
        );
}
