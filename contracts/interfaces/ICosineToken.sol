// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title ICosineToken
 * @dev Interface for the COSINE token contract
 */
interface ICosineToken is IERC20 {
    // Add the burn function
    function burn(uint256 amount) external;

    // Other functions...
    function bridge(uint256 chainId, uint256 amount, bool isLock) external;
    function initializeDistributions(
        address foundation,
        address advisors,
        address developers,
        address privateSale,
        address publicSale,
        address networkRewards
    ) external;
    function claimVestedTokens() external;
    function getVestedAmount(
        address beneficiary
    ) external view returns (uint256);
    function getClaimableAmount(
        address beneficiary
    ) external view returns (uint256);

    // Role functions
    function BRIDGE_ROLE() external view returns (bytes32);
    function DISTRIBUTOR_ROLE() external view returns (bytes32);
    function DEFAULT_ADMIN_ROLE() external view returns (bytes32);
    function grantRole(bytes32 role, address account) external;
    function hasRole(
        bytes32 role,
        address account
    ) external view returns (bool);

    // Distribution state
    function distributionsInitialized() external view returns (bool);

    // Allocation constants
    function FOUNDATION_ALLOCATION() external view returns (uint256);
    function ADVISORS_ALLOCATION() external view returns (uint256);
    function DEVELOPERS_ALLOCATION() external view returns (uint256);
    function PRIVATE_SALE_ALLOCATION() external view returns (uint256);
    function PUBLIC_SALE_ALLOCATION() external view returns (uint256);
    function NETWORK_REWARDS_ALLOCATION() external view returns (uint256);
}
