// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

/**
 * @title CosineToken
 * @dev Implementation of the COSINE token with vesting and distribution mechanisms
 * Total supply: 1,000,000,000 tokens
 * Distribution:
 * - Foundation & Core Team: 10%
 * - Advisors & Contributors: 5%
 * - Early Developer/Builder Grants: 10%
 * - Private Sale / Presale: 15%
 * - Public Sale: 20%
 * - Network Rewards & Ecosystem Funds: 40%
 */
contract CosineToken is ERC20Burnable, AccessControl, Pausable {
    using SafeMath for uint256;

    // Roles
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");
    bytes32 public constant DISTRIBUTOR_ROLE = keccak256("DISTRIBUTOR_ROLE");

    // Token distribution constants
    uint256 public constant TOTAL_SUPPLY = 1_000_000_000 * 10 ** 18; // 1 billion tokens with 18 decimals

    // Allocation percentages in basis points (100 = 1%)
    uint256 public constant FOUNDATION_ALLOCATION = 1000; // 10%
    uint256 public constant ADVISORS_ALLOCATION = 500; // 5%
    uint256 public constant DEVELOPERS_ALLOCATION = 1000; // 10%
    uint256 public constant PRIVATE_SALE_ALLOCATION = 1500; // 15%
    uint256 public constant PUBLIC_SALE_ALLOCATION = 2000; // 20%
    uint256 public constant NETWORK_REWARDS_ALLOCATION = 4000; // 40%

    // Vesting schedule tracking
    struct VestingSchedule {
        uint256 totalAmount;
        uint256 claimedAmount;
        uint256 startTime;
        uint256 cliffDuration;
        uint256 vestingDuration;
        bool isRevocable;
        bool isRevoked;
    }

    // Vesting schedules for different allocations
    mapping(address => VestingSchedule) public vestingSchedules;
    mapping(address => uint256) public vestingShares;

    // Total tokens allocated to all vesting schedules
    uint256 public totalVestingTokens;

    // Accounting for bridge-locked tokens
    mapping(uint256 => uint256) public bridgedTokens; // chainId => amount

    // Token distribution status
    bool public distributionsInitialized;

    // Events
    event TokensVested(address indexed beneficiary, uint256 amount);
    event VestingScheduleCreated(
        address indexed beneficiary,
        uint256 amount,
        uint256 startTime,
        uint256 cliffDuration,
        uint256 vestingDuration
    );
    event VestingScheduleRevoked(
        address indexed beneficiary,
        uint256 amountRecovered
    );
    event TokensBridged(uint256 indexed chainId, uint256 amount);

    /**
     * @dev Constructor that initializes the token and sets up roles
     */
    constructor() ERC20("COSINE", "COS") {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _setupRole(ADMIN_ROLE, msg.sender);
        _setupRole(DISTRIBUTOR_ROLE, msg.sender);

        // Mint all tokens to the contract initially
        _mint(address(this), TOTAL_SUPPLY);

        distributionsInitialized = false;
    }

    /**
     * @dev Initializes token distributions according to the allocations
     * @param foundation Foundation & Core Team address
     * @param advisors Advisors & Contributors address
     * @param developers Early Developer Grants address
     * @param privateSale Private Sale address
     * @param publicSale Public Sale address
     * @param networkRewards Network Rewards & Ecosystem address
     */
    function initializeDistributions(
        address foundation,
        address advisors,
        address developers,
        address privateSale,
        address publicSale,
        address networkRewards
    ) external onlyRole(DISTRIBUTOR_ROLE) {
        require(!distributionsInitialized, "Distributions already initialized");
        require(
            foundation != address(0) &&
                advisors != address(0) &&
                developers != address(0) &&
                privateSale != address(0) &&
                publicSale != address(0) &&
                networkRewards != address(0),
            "Zero addresses not allowed"
        );

        // Calculate token amounts
        uint256 foundationAmount = TOTAL_SUPPLY.mul(FOUNDATION_ALLOCATION).div(
            10000
        );
        uint256 advisorsAmount = TOTAL_SUPPLY.mul(ADVISORS_ALLOCATION).div(
            10000
        );
        uint256 developersAmount = TOTAL_SUPPLY.mul(DEVELOPERS_ALLOCATION).div(
            10000
        );
        uint256 privateSaleAmount = TOTAL_SUPPLY
            .mul(PRIVATE_SALE_ALLOCATION)
            .div(10000);
        uint256 publicSaleAmount = TOTAL_SUPPLY.mul(PUBLIC_SALE_ALLOCATION).div(
            10000
        );
        uint256 networkRewardsAmount = TOTAL_SUPPLY
            .mul(NETWORK_REWARDS_ALLOCATION)
            .div(10000);

        // Setup vesting schedules

        // Foundation: 4-year vesting with 1-year cliff
        _createVestingSchedule(
            foundation,
            foundationAmount,
            block.timestamp,
            365 days, // 1 year cliff
            1460 days, // 4 years total
            true // revocable
        );

        // Advisors: 2-year vesting with 3-month cliff
        _createVestingSchedule(
            advisors,
            advisorsAmount,
            block.timestamp,
            90 days, // 3 months cliff
            730 days, // 2 years total
            true // revocable
        );

        // Developers: Released based on milestones (no standard vesting)
        // We allocate to a controlled account for milestone-based release
        _transfer(address(this), developers, developersAmount);

        // Private Sale: 6-12 months lock, then linear unlock over 6-12 months
        // Average case: 9 months lock, 9 months linear unlock
        _createVestingSchedule(
            privateSale,
            privateSaleAmount,
            block.timestamp,
            270 days, // 9 months lock
            540 days, // 9 months additional vesting after cliff
            false // not revocable
        );

        // Public Sale: Unlocked immediately at Token Generation Event
        _transfer(address(this), publicSale, publicSaleAmount);

        // Network Rewards & Ecosystem: Distributed over 5-8 years
        // Using 6.5 years (2372 days) as the average
        _createVestingSchedule(
            networkRewards,
            networkRewardsAmount,
            block.timestamp,
            0, // No cliff
            2372 days, // 6.5 years total
            false // not revocable
        );

        distributionsInitialized = true;
    }

    /**
     * @dev Internal function to create a vesting schedule
     */
    function _createVestingSchedule(
        address beneficiary,
        uint256 amount,
        uint256 startTime,
        uint256 cliffDuration,
        uint256 vestingDuration,
        bool isRevocable
    ) internal {
        require(
            beneficiary != address(0),
            "Beneficiary cannot be zero address"
        );
        require(amount > 0, "Amount must be greater than zero");
        require(
            vestingSchedules[beneficiary].totalAmount == 0,
            "Vesting schedule already exists"
        );

        vestingSchedules[beneficiary] = VestingSchedule({
            totalAmount: amount,
            claimedAmount: 0,
            startTime: startTime,
            cliffDuration: cliffDuration,
            vestingDuration: vestingDuration,
            isRevocable: isRevocable,
            isRevoked: false
        });

        totalVestingTokens = totalVestingTokens.add(amount);

        emit VestingScheduleCreated(
            beneficiary,
            amount,
            startTime,
            cliffDuration,
            vestingDuration
        );
    }

    /**
     * @dev Allows a beneficiary to claim their vested tokens
     */
    function claimVestedTokens() external whenNotPaused {
        VestingSchedule storage schedule = vestingSchedules[msg.sender];

        require(schedule.totalAmount > 0, "No vesting schedule found");
        require(!schedule.isRevoked, "Vesting schedule revoked");

        uint256 vestedAmount = _calculateVestedAmount(msg.sender);
        uint256 claimableAmount = vestedAmount.sub(schedule.claimedAmount);

        require(claimableAmount > 0, "No tokens available to claim");

        schedule.claimedAmount = schedule.claimedAmount.add(claimableAmount);

        _transfer(address(this), msg.sender, claimableAmount);
        emit TokensVested(msg.sender, claimableAmount);
    }

    /**
     * @dev Calculates the amount of tokens that have vested for a given beneficiary
     * @param beneficiary The address to calculate vested tokens for
     * @return The amount of vested tokens
     */
    function _calculateVestedAmount(
        address beneficiary
    ) internal view returns (uint256) {
        VestingSchedule memory schedule = vestingSchedules[beneficiary];

        if (schedule.totalAmount == 0 || schedule.isRevoked) {
            return 0;
        }

        // If before cliff, nothing is vested
        if (block.timestamp < schedule.startTime.add(schedule.cliffDuration)) {
            return 0;
        }

        // If after vesting duration, everything is vested
        if (
            block.timestamp >= schedule.startTime.add(schedule.vestingDuration)
        ) {
            return schedule.totalAmount;
        }

        // Otherwise, calculate linearly vested amount
        uint256 timeFromStart = block.timestamp.sub(schedule.startTime);
        return
            schedule.totalAmount.mul(timeFromStart).div(
                schedule.vestingDuration
            );
    }

    /**
     * @dev Returns the vested amount for a beneficiary (external view)
     */
    function getVestedAmount(
        address beneficiary
    ) external view returns (uint256) {
        return _calculateVestedAmount(beneficiary);
    }

    /**
     * @dev Returns the claimable amount for a beneficiary
     */
    function getClaimableAmount(
        address beneficiary
    ) external view returns (uint256) {
        VestingSchedule memory schedule = vestingSchedules[beneficiary];

        if (schedule.totalAmount == 0 || schedule.isRevoked) {
            return 0;
        }

        uint256 vestedAmount = _calculateVestedAmount(beneficiary);
        return vestedAmount.sub(schedule.claimedAmount);
    }

    /**
     * @dev Revokes a vesting schedule (only for revocable schedules)
     * @param beneficiary The address to revoke the schedule for
     */
    function revokeVestingSchedule(
        address beneficiary
    ) external onlyRole(ADMIN_ROLE) {
        VestingSchedule storage schedule = vestingSchedules[beneficiary];

        require(schedule.totalAmount > 0, "No vesting schedule found");
        require(schedule.isRevocable, "Schedule is not revocable");
        require(!schedule.isRevoked, "Schedule already revoked");

        uint256 vestedAmount = _calculateVestedAmount(beneficiary);
        uint256 amountToRecover = schedule.totalAmount.sub(vestedAmount);

        schedule.isRevoked = true;
        totalVestingTokens = totalVestingTokens.sub(amountToRecover);

        emit VestingScheduleRevoked(beneficiary, amountToRecover);
    }

    /**
     * @dev Function for bridge operations (lock/mint or burn/redeem)
     * @param chainId The chain ID where tokens are being locked or released
     * @param amount The amount of tokens to bridge
     * @param isLock True if tokens are being locked, false if being released
     */
    function bridge(
        uint256 chainId,
        uint256 amount,
        bool isLock
    ) external onlyRole(BRIDGE_ROLE) {
        if (isLock) {
            // Lock tokens on this chain for minting on another chain
            _burn(msg.sender, amount);
            bridgedTokens[chainId] = bridgedTokens[chainId].add(amount);
        } else {
            // Release tokens that were locked on another chain (coming back)
            require(
                bridgedTokens[chainId] >= amount,
                "Not enough tokens bridged to this chain"
            );
            bridgedTokens[chainId] = bridgedTokens[chainId].sub(amount);
            _mint(msg.sender, amount);
        }

        emit TokensBridged(chainId, amount);
    }

    /**
     * @dev Pause token transfers
     */
    function pause() external onlyRole(ADMIN_ROLE) {
        _pause();
    }

    /**
     * @dev Unpause token transfers
     */
    function unpause() external onlyRole(ADMIN_ROLE) {
        _unpause();
    }

    /**
     * @dev Hook that is called before any transfer of tokens
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override whenNotPaused {
        super._beforeTokenTransfer(from, to, amount);
    }
}
