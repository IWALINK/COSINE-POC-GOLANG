// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CosineTokenMetaData contains all meta data concerning the CosineToken contract.
var CosineTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensVested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cliffDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vestingDuration\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRecovered\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADVISORS_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEVELOPERS_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOUNDATION_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETWORK_REWARDS_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIVATE_SALE_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_SALE_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimVestedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionsInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"getVestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"foundation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"advisors\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"developers\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"privateSale\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"publicSale\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"networkRewards\",\"type\":\"address\"}],\"name\":\"initializeDistributions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"revokeVestingSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVestingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vestingSchedules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliffDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingDuration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRevocable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRevoked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vestingShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405180604001604052806006815260200165434f53494e4560d01b81525060405180604001604052806003815260200162434f5360e81b81525081600390816200005e9190620003b3565b5060046200006d8282620003b3565b50506006805460ff19169055506200008760003362000107565b620000b37fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217753362000107565b620000df7ffbd454f36a7e1a388bd6fc3ab10d434aa4578f811acbbcf33afb1c697486313c3362000107565b620000f7306b033b2e3c9fd0803ce800000062000117565b600b805460ff19169055620004a1565b620001138282620001ec565b5050565b6001600160a01b038216620001735760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b620001816000838362000276565b80600260008282546200019591906200047f565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b620001f8828262000298565b620001135760008281526005602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620002323390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b62000280620002c5565b620002938383836001600160e01b038416565b505050565b60008281526005602090815260408083206001600160a01b038516845290915290205460ff165b92915050565b60065460ff16156200030d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016200016a565b565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200033a57607f821691505b6020821081036200035b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200029357600081815260208120601f850160051c810160208610156200038a5750805b601f850160051c820191505b81811015620003ab5782815560010162000396565b505050505050565b81516001600160401b03811115620003cf57620003cf6200030f565b620003e781620003e0845462000325565b8462000361565b602080601f8311600181146200041f5760008415620004065750858301515b600019600386901b1c1916600185901b178555620003ab565b600085815260208120601f198616915b8281101562000450578886015182559484019460019091019084016200042f565b50858210156200046f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b80820180821115620002bf57634e487b7160e01b600052601160045260246000fd5b61225180620004b16000396000f3fe608060405234801561001057600080fd5b50600436106102745760003560e01c80638456cb5911610151578063c0fc859e116100c3578063dd62ed3e11610087578063dd62ed3e1461052b578063e12f3a611461053e578063e74f3fbb14610551578063f035a27214610559578063f0bd87cc1461056c578063fdb20ccb1461059357600080fd5b8063c0fc859e146104cf578063c26e2402146104dc578063c409b8f9146104e5578063d547741f14610505578063d5a73fdd1461051857600080fd5b806395d89b411161011557806395d89b4114610469578063a217fddf14610471578063a457c2d714610479578063a6c3e5311461048c578063a9059cbb14610495578063b5bfddea146104a857600080fd5b80638456cb59146104335780638b997b791461043b578063902d55a5146104445780639083275c1461034657806391d148541461045657600080fd5b8063313ce567116101ea57806342966c68116101ae57806342966c68146103b15780635a71c528146103c45780635c975abb146103d757806370a08231146103e257806375b238fc1461040b57806379cc67901461042057600080fd5b8063313ce5671461036b57806336568abe1461037a578063395093511461038d5780633dc762f0146103a05780633f4ba83a146103a957600080fd5b806321edb61b1161023c57806321edb61b146102f057806323b872dd14610310578063248a9ca314610323578063289e89bf146103465780632de2f0701461034f5780632f2ff15d1461035857600080fd5b806301ffc9a71461027957806306fdde03146102a1578063095ea7b3146102b657806317727d35146102c957806318160ddd146102de575b600080fd5b61028c610287366004611e55565b61061b565b60405190151581526020015b60405180910390f35b6102a9610652565b6040516102989190611ea3565b61028c6102c4366004611ef2565b6106e4565b6102dc6102d7366004611f1c565b6106fc565b005b6002545b604051908152602001610298565b6102e26102fe366004611f5a565b60086020526000908152604090205481565b61028c61031e366004611f75565b610848565b6102e2610331366004611fb1565b60009081526005602052604090206001015490565b6102e26103e881565b6102e2610fa081565b6102dc610366366004611fca565b61086c565b60405160128152602001610298565b6102dc610388366004611fca565b610896565b61028c61039b366004611ef2565b610914565b6102e26107d081565b6102dc610936565b6102dc6103bf366004611fb1565b610959565b6102dc6103d2366004611ff6565b610963565b60065460ff1661028c565b6102e26103f0366004611f5a565b6001600160a01b031660009081526020819052604090205490565b6102e26000805160206121fc83398151915281565b6102dc61042e366004611ef2565b610bfc565b6102dc610c11565b6102e260095481565b6102e2676765c793fa10079d601b1b81565b61028c610464366004611fca565b610c31565b6102a9610c5c565b6102e2600081565b61028c610487366004611ef2565b610c6b565b6102e26105dc81565b61028c6104a3366004611ef2565b610ce6565b6102e27f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f81565b600b5461028c9060ff1681565b6102e26101f481565b6102e26104f3366004611fb1565b600a6020526000908152604090205481565b6102dc610513366004611fca565b610cf4565b6102e2610526366004611f5a565b610d19565b6102e261053936600461206a565b610d24565b6102e261054c366004611f5a565b610d4f565b6102dc610e09565b6102dc610567366004611f5a565b610f95565b6102e27ffbd454f36a7e1a388bd6fc3ab10d434aa4578f811acbbcf33afb1c697486313c81565b6105e26105a1366004611f5a565b600760205260009081526040902080546001820154600283015460038401546004850154600590950154939492939192909160ff8082169161010090041687565b6040805197885260208801969096529486019390935260608501919091526080840152151560a0830152151560c082015260e001610298565b60006001600160e01b03198216637965db0b60e01b148061064c57506301ffc9a760e01b6001600160e01b03198316145b92915050565b60606003805461066190612094565b80601f016020809104026020016040519081016040528092919081815260200182805461068d90612094565b80156106da5780601f106106af576101008083540402835291602001916106da565b820191906000526020600020905b8154815290600101906020018083116106bd57829003601f168201915b5050505050905090565b6000336106f2818585611143565b5060019392505050565b7f52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f61072681611267565b8115610763576107363384611271565b6000848152600a602052604090205461074f90846113af565b6000858152600a6020526040902055610808565b6000848152600a60205260409020548311156107d65760405162461bcd60e51b815260206004820152602760248201527f4e6f7420656e6f75676820746f6b656e73206272696467656420746f20746869604482015266399031b430b4b760c91b60648201526084015b60405180910390fd5b6000848152600a60205260409020546107ef90846113c2565b6000858152600a602052604090205561080833846113ce565b837f9ae36a2ae9e50dd27b843521a327f5c2a422a56649281bc6c5754f15b4a01b7e8460405161083a91815260200190565b60405180910390a250505050565b600033610856858285611499565b610861858585611513565b506001949350505050565b60008281526005602052604090206001015461088781611267565b61089183836116c2565b505050565b6001600160a01b03811633146109065760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016107cd565b6109108282611748565b5050565b6000336106f28185856109278383610d24565b61093191906120e4565b611143565b6000805160206121fc83398151915261094e81611267565b6109566117af565b50565b6109563382611271565b7ffbd454f36a7e1a388bd6fc3ab10d434aa4578f811acbbcf33afb1c697486313c61098d81611267565b600b5460ff16156109ea5760405162461bcd60e51b815260206004820152602160248201527f446973747269627574696f6e7320616c726561647920696e697469616c697a656044820152601960fa1b60648201526084016107cd565b6001600160a01b03871615801590610a0a57506001600160a01b03861615155b8015610a1e57506001600160a01b03851615155b8015610a3257506001600160a01b03841615155b8015610a4657506001600160a01b03831615155b8015610a5a57506001600160a01b03821615155b610aa65760405162461bcd60e51b815260206004820152601a60248201527f5a65726f20616464726573736573206e6f7420616c6c6f77656400000000000060448201526064016107cd565b6000610acb612710610ac5676765c793fa10079d601b1b6103e8611801565b9061180d565b90506000610aec612710610ac5676765c793fa10079d601b1b6101f4611801565b90506000610b0d612710610ac5676765c793fa10079d601b1b6103e8611801565b90506000610b2e612710610ac5676765c793fa10079d601b1b6105dc611801565b90506000610b4f612710610ac5676765c793fa10079d601b1b6107d0611801565b90506000610b70612710610ac5676765c793fa10079d601b1b610fa0611801565b9050610b898d87426301e13380630784ce006001611819565b610b9f8c86426276a7006303c267006001611819565b610baa308c86611513565b610bc18a8442630163f5006302c7ea006000611819565b610bcc308a84611513565b610be08882426000630c3726006000611819565b5050600b805460ff191660011790555050505050505050505050565b610c07823383611499565b6109108282611271565b6000805160206121fc833981519152610c2981611267565b610956611a76565b60009182526005602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606004805461066190612094565b60003381610c798286610d24565b905083811015610cd95760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016107cd565b6108618286868403611143565b6000336106f2818585611513565b600082815260056020526040902060010154610d0f81611267565b6108918383611748565b600061064c82611ab3565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0381166000908152600760209081526040808320815160e081018352815480825260018301549482019490945260028201549281019290925260038101546060830152600481015460808301526005015460ff808216151560a084015261010090910416151560c0820152901580610dcf57508060c001515b15610ddd5750600092915050565b6000610de884611ab3565b9050610e018260200151826113c290919063ffffffff16565b949350505050565b610e11611bb5565b3360009081526007602052604090208054610e6a5760405162461bcd60e51b8152602060048201526019602482015278139bc81d995cdd1a5b99c81cd8da19591d5b1948199bdd5b99603a1b60448201526064016107cd565b6005810154610100900460ff1615610ec45760405162461bcd60e51b815260206004820152601860248201527f56657374696e67207363686564756c65207265766f6b6564000000000000000060448201526064016107cd565b6000610ecf33611ab3565b90506000610eea8360010154836113c290919063ffffffff16565b905060008111610f3c5760405162461bcd60e51b815260206004820152601c60248201527f4e6f20746f6b656e7320617661696c61626c6520746f20636c61696d0000000060448201526064016107cd565b6001830154610f4b90826113af565b6001840155610f5b303383611513565b60405181815233907fd4691cc79b8fc72aac1e8c0d15a2ca06d71a386c983f815266f0fce713dc5ea79060200160405180910390a2505050565b6000805160206121fc833981519152610fad81611267565b6001600160a01b0382166000908152600760205260409020805461100f5760405162461bcd60e51b8152602060048201526019602482015278139bc81d995cdd1a5b99c81cd8da19591d5b1948199bdd5b99603a1b60448201526064016107cd565b600581015460ff166110635760405162461bcd60e51b815260206004820152601960248201527f5363686564756c65206973206e6f74207265766f6361626c650000000000000060448201526064016107cd565b6005810154610100900460ff16156110bd5760405162461bcd60e51b815260206004820152601860248201527f5363686564756c6520616c7265616479207265766f6b6564000000000000000060448201526064016107cd565b60006110c884611ab3565b82549091506000906110da90836113c2565b60058401805461ff0019166101001790556009549091506110fb90826113c2565b6009556040518181526001600160a01b038616907fb71d84e80ebac87b433939f1a1191d542623a94ff5b940314af5589461f352759060200160405180910390a25050505050565b6001600160a01b0383166111a55760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016107cd565b6001600160a01b0382166112065760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016107cd565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6109568133611bfd565b6001600160a01b0382166112d15760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016107cd565b6112dd82600083611c56565b6001600160a01b038216600090815260208190526040902054818110156113515760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016107cd565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b60006113bb82846120e4565b9392505050565b60006113bb82846120f7565b6001600160a01b0382166114245760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016107cd565b61143060008383611c56565b806002600082825461144291906120e4565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b60006114a58484610d24565b9050600019811461150d57818110156115005760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016107cd565b61150d8484848403611143565b50505050565b6001600160a01b0383166115775760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016107cd565b6001600160a01b0382166115d95760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016107cd565b6115e4838383611c56565b6001600160a01b0383166000908152602081905260409020548181101561165c5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016107cd565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361150d565b6116cc8282610c31565b6109105760008281526005602090815260408083206001600160a01b03851684529091529020805460ff191660011790556117043390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6117528282610c31565b156109105760008281526005602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6117b7611c5e565b6006805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60006113bb828461210a565b60006113bb8284612121565b6001600160a01b03861661187a5760405162461bcd60e51b815260206004820152602260248201527f42656e65666963696172792063616e6e6f74206265207a65726f206164647265604482015261737360f01b60648201526084016107cd565b600085116118ca5760405162461bcd60e51b815260206004820181905260248201527f416d6f756e74206d7573742062652067726561746572207468616e207a65726f60448201526064016107cd565b6001600160a01b038616600090815260076020526040902054156119305760405162461bcd60e51b815260206004820152601f60248201527f56657374696e67207363686564756c6520616c7265616479206578697374730060448201526064016107cd565b6040518060e001604052808681526020016000815260200185815260200184815260200183815260200182151581526020016000151581525060076000886001600160a01b03166001600160a01b03168152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548160ff02191690831515021790555060c08201518160050160016101000a81548160ff021916908315150217905550905050611a18856009546113af90919063ffffffff16565b6009556040805186815260208101869052908101849052606081018390526001600160a01b038716907f47d46f58c7d60d60571092cff5715b8c8c5bc1347db2bb9890e639bf3680ba7b9060800160405180910390a2505050505050565b611a7e611bb5565b6006805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586117e43390565b6001600160a01b0381166000908152600760209081526040808320815160e081018352815480825260018301549482019490945260028201549281019290925260038101546060830152600481015460808301526005015460ff808216151560a084015261010090910416151560c0820152901580611b3357508060c001515b15611b415750600092915050565b60608101516040820151611b54916113af565b421015611b645750600092915050565b60808101516040820151611b77916113af565b4210611b84575192915050565b6000611b9d8260400151426113c290919063ffffffff16565b60808301518351919250610e0191610ac59084611801565b60065460ff1615611bfb5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016107cd565b565b611c078282610c31565b61091057611c1481611ca7565b611c1f836020611cb9565b604051602001611c30929190612143565b60408051601f198184030181529082905262461bcd60e51b82526107cd91600401611ea3565b610891611bb5565b60065460ff16611bfb5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016107cd565b606061064c6001600160a01b03831660145b60606000611cc883600261210a565b611cd39060026120e4565b67ffffffffffffffff811115611ceb57611ceb6121b8565b6040519080825280601f01601f191660200182016040528015611d15576020820181803683370190505b509050600360fc1b81600081518110611d3057611d306121ce565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611d5f57611d5f6121ce565b60200101906001600160f81b031916908160001a9053506000611d8384600261210a565b611d8e9060016120e4565b90505b6001811115611e06576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611dc257611dc26121ce565b1a60f81b828281518110611dd857611dd86121ce565b60200101906001600160f81b031916908160001a90535060049490941c93611dff816121e4565b9050611d91565b5083156113bb5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016107cd565b600060208284031215611e6757600080fd5b81356001600160e01b0319811681146113bb57600080fd5b60005b83811015611e9a578181015183820152602001611e82565b50506000910152565b6020815260008251806020840152611ec2816040850160208701611e7f565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114611eed57600080fd5b919050565b60008060408385031215611f0557600080fd5b611f0e83611ed6565b946020939093013593505050565b600080600060608486031215611f3157600080fd5b833592506020840135915060408401358015158114611f4f57600080fd5b809150509250925092565b600060208284031215611f6c57600080fd5b6113bb82611ed6565b600080600060608486031215611f8a57600080fd5b611f9384611ed6565b9250611fa160208501611ed6565b9150604084013590509250925092565b600060208284031215611fc357600080fd5b5035919050565b60008060408385031215611fdd57600080fd5b82359150611fed60208401611ed6565b90509250929050565b60008060008060008060c0878903121561200f57600080fd5b61201887611ed6565b955061202660208801611ed6565b945061203460408801611ed6565b935061204260608801611ed6565b925061205060808801611ed6565b915061205e60a08801611ed6565b90509295509295509295565b6000806040838503121561207d57600080fd5b61208683611ed6565b9150611fed60208401611ed6565b600181811c908216806120a857607f821691505b6020821081036120c857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561064c5761064c6120ce565b8181038181111561064c5761064c6120ce565b808202811582820484141761064c5761064c6120ce565b60008261213e57634e487b7160e01b600052601260045260246000fd5b500490565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081526000835161217b816017850160208801611e7f565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516121ac816028840160208801611e7f565b01602801949350505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000816121f3576121f36120ce565b50600019019056fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220a8fe04ac68e977e9fed838ad08b35ab61d9933000507e36c5dffd9a84e1e012564736f6c63430008130033",
}

// CosineTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CosineTokenMetaData.ABI instead.
var CosineTokenABI = CosineTokenMetaData.ABI

// CosineTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosineTokenMetaData.Bin instead.
var CosineTokenBin = CosineTokenMetaData.Bin

// DeployCosineToken deploys a new Ethereum contract, binding an instance of CosineToken to it.
func DeployCosineToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosineToken, error) {
	parsed, err := CosineTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosineTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosineToken{CosineTokenCaller: CosineTokenCaller{contract: contract}, CosineTokenTransactor: CosineTokenTransactor{contract: contract}, CosineTokenFilterer: CosineTokenFilterer{contract: contract}}, nil
}

// CosineToken is an auto generated Go binding around an Ethereum contract.
type CosineToken struct {
	CosineTokenCaller     // Read-only binding to the contract
	CosineTokenTransactor // Write-only binding to the contract
	CosineTokenFilterer   // Log filterer for contract events
}

// CosineTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosineTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosineTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosineTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosineTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosineTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosineTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosineTokenSession struct {
	Contract     *CosineToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosineTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosineTokenCallerSession struct {
	Contract *CosineTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CosineTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosineTokenTransactorSession struct {
	Contract     *CosineTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosineTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosineTokenRaw struct {
	Contract *CosineToken // Generic contract binding to access the raw methods on
}

// CosineTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosineTokenCallerRaw struct {
	Contract *CosineTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CosineTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosineTokenTransactorRaw struct {
	Contract *CosineTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosineToken creates a new instance of CosineToken, bound to a specific deployed contract.
func NewCosineToken(address common.Address, backend bind.ContractBackend) (*CosineToken, error) {
	contract, err := bindCosineToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosineToken{CosineTokenCaller: CosineTokenCaller{contract: contract}, CosineTokenTransactor: CosineTokenTransactor{contract: contract}, CosineTokenFilterer: CosineTokenFilterer{contract: contract}}, nil
}

// NewCosineTokenCaller creates a new read-only instance of CosineToken, bound to a specific deployed contract.
func NewCosineTokenCaller(address common.Address, caller bind.ContractCaller) (*CosineTokenCaller, error) {
	contract, err := bindCosineToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosineTokenCaller{contract: contract}, nil
}

// NewCosineTokenTransactor creates a new write-only instance of CosineToken, bound to a specific deployed contract.
func NewCosineTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CosineTokenTransactor, error) {
	contract, err := bindCosineToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosineTokenTransactor{contract: contract}, nil
}

// NewCosineTokenFilterer creates a new log filterer instance of CosineToken, bound to a specific deployed contract.
func NewCosineTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CosineTokenFilterer, error) {
	contract, err := bindCosineToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosineTokenFilterer{contract: contract}, nil
}

// bindCosineToken binds a generic wrapper to an already deployed contract.
func bindCosineToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CosineTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosineToken *CosineTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosineToken.Contract.CosineTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosineToken *CosineTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosineToken.Contract.CosineTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosineToken *CosineTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosineToken.Contract.CosineTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosineToken *CosineTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosineToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosineToken *CosineTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosineToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosineToken *CosineTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosineToken.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenSession) ADMINROLE() ([32]byte, error) {
	return _CosineToken.Contract.ADMINROLE(&_CosineToken.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCallerSession) ADMINROLE() ([32]byte, error) {
	return _CosineToken.Contract.ADMINROLE(&_CosineToken.CallOpts)
}

// ADVISORSALLOCATION is a free data retrieval call binding the contract method 0xc26e2402.
//
// Solidity: function ADVISORS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) ADVISORSALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "ADVISORS_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADVISORSALLOCATION is a free data retrieval call binding the contract method 0xc26e2402.
//
// Solidity: function ADVISORS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) ADVISORSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.ADVISORSALLOCATION(&_CosineToken.CallOpts)
}

// ADVISORSALLOCATION is a free data retrieval call binding the contract method 0xc26e2402.
//
// Solidity: function ADVISORS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) ADVISORSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.ADVISORSALLOCATION(&_CosineToken.CallOpts)
}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCaller) BRIDGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "BRIDGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenSession) BRIDGEROLE() ([32]byte, error) {
	return _CosineToken.Contract.BRIDGEROLE(&_CosineToken.CallOpts)
}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCallerSession) BRIDGEROLE() ([32]byte, error) {
	return _CosineToken.Contract.BRIDGEROLE(&_CosineToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CosineToken.Contract.DEFAULTADMINROLE(&_CosineToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CosineToken.Contract.DEFAULTADMINROLE(&_CosineToken.CallOpts)
}

// DEVELOPERSALLOCATION is a free data retrieval call binding the contract method 0x9083275c.
//
// Solidity: function DEVELOPERS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) DEVELOPERSALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "DEVELOPERS_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEVELOPERSALLOCATION is a free data retrieval call binding the contract method 0x9083275c.
//
// Solidity: function DEVELOPERS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) DEVELOPERSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.DEVELOPERSALLOCATION(&_CosineToken.CallOpts)
}

// DEVELOPERSALLOCATION is a free data retrieval call binding the contract method 0x9083275c.
//
// Solidity: function DEVELOPERS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) DEVELOPERSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.DEVELOPERSALLOCATION(&_CosineToken.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCaller) DISTRIBUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "DISTRIBUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _CosineToken.Contract.DISTRIBUTORROLE(&_CosineToken.CallOpts)
}

// DISTRIBUTORROLE is a free data retrieval call binding the contract method 0xf0bd87cc.
//
// Solidity: function DISTRIBUTOR_ROLE() view returns(bytes32)
func (_CosineToken *CosineTokenCallerSession) DISTRIBUTORROLE() ([32]byte, error) {
	return _CosineToken.Contract.DISTRIBUTORROLE(&_CosineToken.CallOpts)
}

// FOUNDATIONALLOCATION is a free data retrieval call binding the contract method 0x289e89bf.
//
// Solidity: function FOUNDATION_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) FOUNDATIONALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "FOUNDATION_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FOUNDATIONALLOCATION is a free data retrieval call binding the contract method 0x289e89bf.
//
// Solidity: function FOUNDATION_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) FOUNDATIONALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.FOUNDATIONALLOCATION(&_CosineToken.CallOpts)
}

// FOUNDATIONALLOCATION is a free data retrieval call binding the contract method 0x289e89bf.
//
// Solidity: function FOUNDATION_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) FOUNDATIONALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.FOUNDATIONALLOCATION(&_CosineToken.CallOpts)
}

// NETWORKREWARDSALLOCATION is a free data retrieval call binding the contract method 0x2de2f070.
//
// Solidity: function NETWORK_REWARDS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) NETWORKREWARDSALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "NETWORK_REWARDS_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NETWORKREWARDSALLOCATION is a free data retrieval call binding the contract method 0x2de2f070.
//
// Solidity: function NETWORK_REWARDS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) NETWORKREWARDSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.NETWORKREWARDSALLOCATION(&_CosineToken.CallOpts)
}

// NETWORKREWARDSALLOCATION is a free data retrieval call binding the contract method 0x2de2f070.
//
// Solidity: function NETWORK_REWARDS_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) NETWORKREWARDSALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.NETWORKREWARDSALLOCATION(&_CosineToken.CallOpts)
}

// PRIVATESALEALLOCATION is a free data retrieval call binding the contract method 0xa6c3e531.
//
// Solidity: function PRIVATE_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) PRIVATESALEALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "PRIVATE_SALE_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRIVATESALEALLOCATION is a free data retrieval call binding the contract method 0xa6c3e531.
//
// Solidity: function PRIVATE_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) PRIVATESALEALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.PRIVATESALEALLOCATION(&_CosineToken.CallOpts)
}

// PRIVATESALEALLOCATION is a free data retrieval call binding the contract method 0xa6c3e531.
//
// Solidity: function PRIVATE_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) PRIVATESALEALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.PRIVATESALEALLOCATION(&_CosineToken.CallOpts)
}

// PUBLICSALEALLOCATION is a free data retrieval call binding the contract method 0x3dc762f0.
//
// Solidity: function PUBLIC_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCaller) PUBLICSALEALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "PUBLIC_SALE_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICSALEALLOCATION is a free data retrieval call binding the contract method 0x3dc762f0.
//
// Solidity: function PUBLIC_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenSession) PUBLICSALEALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.PUBLICSALEALLOCATION(&_CosineToken.CallOpts)
}

// PUBLICSALEALLOCATION is a free data retrieval call binding the contract method 0x3dc762f0.
//
// Solidity: function PUBLIC_SALE_ALLOCATION() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) PUBLICSALEALLOCATION() (*big.Int, error) {
	return _CosineToken.Contract.PUBLICSALEALLOCATION(&_CosineToken.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_CosineToken *CosineTokenCaller) TOTALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "TOTAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_CosineToken *CosineTokenSession) TOTALSUPPLY() (*big.Int, error) {
	return _CosineToken.Contract.TOTALSUPPLY(&_CosineToken.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) TOTALSUPPLY() (*big.Int, error) {
	return _CosineToken.Contract.TOTALSUPPLY(&_CosineToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosineToken *CosineTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosineToken *CosineTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosineToken.Contract.Allowance(&_CosineToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosineToken.Contract.Allowance(&_CosineToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosineToken *CosineTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosineToken *CosineTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosineToken.Contract.BalanceOf(&_CosineToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosineToken.Contract.BalanceOf(&_CosineToken.CallOpts, account)
}

// BridgedTokens is a free data retrieval call binding the contract method 0xc409b8f9.
//
// Solidity: function bridgedTokens(uint256 ) view returns(uint256)
func (_CosineToken *CosineTokenCaller) BridgedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "bridgedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgedTokens is a free data retrieval call binding the contract method 0xc409b8f9.
//
// Solidity: function bridgedTokens(uint256 ) view returns(uint256)
func (_CosineToken *CosineTokenSession) BridgedTokens(arg0 *big.Int) (*big.Int, error) {
	return _CosineToken.Contract.BridgedTokens(&_CosineToken.CallOpts, arg0)
}

// BridgedTokens is a free data retrieval call binding the contract method 0xc409b8f9.
//
// Solidity: function bridgedTokens(uint256 ) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) BridgedTokens(arg0 *big.Int) (*big.Int, error) {
	return _CosineToken.Contract.BridgedTokens(&_CosineToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosineToken *CosineTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosineToken *CosineTokenSession) Decimals() (uint8, error) {
	return _CosineToken.Contract.Decimals(&_CosineToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosineToken *CosineTokenCallerSession) Decimals() (uint8, error) {
	return _CosineToken.Contract.Decimals(&_CosineToken.CallOpts)
}

// DistributionsInitialized is a free data retrieval call binding the contract method 0xc0fc859e.
//
// Solidity: function distributionsInitialized() view returns(bool)
func (_CosineToken *CosineTokenCaller) DistributionsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "distributionsInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DistributionsInitialized is a free data retrieval call binding the contract method 0xc0fc859e.
//
// Solidity: function distributionsInitialized() view returns(bool)
func (_CosineToken *CosineTokenSession) DistributionsInitialized() (bool, error) {
	return _CosineToken.Contract.DistributionsInitialized(&_CosineToken.CallOpts)
}

// DistributionsInitialized is a free data retrieval call binding the contract method 0xc0fc859e.
//
// Solidity: function distributionsInitialized() view returns(bool)
func (_CosineToken *CosineTokenCallerSession) DistributionsInitialized() (bool, error) {
	return _CosineToken.Contract.DistributionsInitialized(&_CosineToken.CallOpts)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenCaller) GetClaimableAmount(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "getClaimableAmount", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenSession) GetClaimableAmount(beneficiary common.Address) (*big.Int, error) {
	return _CosineToken.Contract.GetClaimableAmount(&_CosineToken.CallOpts, beneficiary)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) GetClaimableAmount(beneficiary common.Address) (*big.Int, error) {
	return _CosineToken.Contract.GetClaimableAmount(&_CosineToken.CallOpts, beneficiary)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CosineToken *CosineTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CosineToken *CosineTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CosineToken.Contract.GetRoleAdmin(&_CosineToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CosineToken *CosineTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CosineToken.Contract.GetRoleAdmin(&_CosineToken.CallOpts, role)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenCaller) GetVestedAmount(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "getVestedAmount", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenSession) GetVestedAmount(beneficiary common.Address) (*big.Int, error) {
	return _CosineToken.Contract.GetVestedAmount(&_CosineToken.CallOpts, beneficiary)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address beneficiary) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) GetVestedAmount(beneficiary common.Address) (*big.Int, error) {
	return _CosineToken.Contract.GetVestedAmount(&_CosineToken.CallOpts, beneficiary)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CosineToken *CosineTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CosineToken *CosineTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CosineToken.Contract.HasRole(&_CosineToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CosineToken *CosineTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CosineToken.Contract.HasRole(&_CosineToken.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosineToken *CosineTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosineToken *CosineTokenSession) Name() (string, error) {
	return _CosineToken.Contract.Name(&_CosineToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosineToken *CosineTokenCallerSession) Name() (string, error) {
	return _CosineToken.Contract.Name(&_CosineToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CosineToken *CosineTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CosineToken *CosineTokenSession) Paused() (bool, error) {
	return _CosineToken.Contract.Paused(&_CosineToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CosineToken *CosineTokenCallerSession) Paused() (bool, error) {
	return _CosineToken.Contract.Paused(&_CosineToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosineToken *CosineTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosineToken *CosineTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CosineToken.Contract.SupportsInterface(&_CosineToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CosineToken *CosineTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CosineToken.Contract.SupportsInterface(&_CosineToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosineToken *CosineTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosineToken *CosineTokenSession) Symbol() (string, error) {
	return _CosineToken.Contract.Symbol(&_CosineToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosineToken *CosineTokenCallerSession) Symbol() (string, error) {
	return _CosineToken.Contract.Symbol(&_CosineToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosineToken *CosineTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosineToken *CosineTokenSession) TotalSupply() (*big.Int, error) {
	return _CosineToken.Contract.TotalSupply(&_CosineToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CosineToken.Contract.TotalSupply(&_CosineToken.CallOpts)
}

// TotalVestingTokens is a free data retrieval call binding the contract method 0x8b997b79.
//
// Solidity: function totalVestingTokens() view returns(uint256)
func (_CosineToken *CosineTokenCaller) TotalVestingTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "totalVestingTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVestingTokens is a free data retrieval call binding the contract method 0x8b997b79.
//
// Solidity: function totalVestingTokens() view returns(uint256)
func (_CosineToken *CosineTokenSession) TotalVestingTokens() (*big.Int, error) {
	return _CosineToken.Contract.TotalVestingTokens(&_CosineToken.CallOpts)
}

// TotalVestingTokens is a free data retrieval call binding the contract method 0x8b997b79.
//
// Solidity: function totalVestingTokens() view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) TotalVestingTokens() (*big.Int, error) {
	return _CosineToken.Contract.TotalVestingTokens(&_CosineToken.CallOpts)
}

// VestingSchedules is a free data retrieval call binding the contract method 0xfdb20ccb.
//
// Solidity: function vestingSchedules(address ) view returns(uint256 totalAmount, uint256 claimedAmount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration, bool isRevocable, bool isRevoked)
func (_CosineToken *CosineTokenCaller) VestingSchedules(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	ClaimedAmount   *big.Int
	StartTime       *big.Int
	CliffDuration   *big.Int
	VestingDuration *big.Int
	IsRevocable     bool
	IsRevoked       bool
}, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "vestingSchedules", arg0)

	outstruct := new(struct {
		TotalAmount     *big.Int
		ClaimedAmount   *big.Int
		StartTime       *big.Int
		CliffDuration   *big.Int
		VestingDuration *big.Int
		IsRevocable     bool
		IsRevoked       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClaimedAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CliffDuration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VestingDuration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsRevocable = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsRevoked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// VestingSchedules is a free data retrieval call binding the contract method 0xfdb20ccb.
//
// Solidity: function vestingSchedules(address ) view returns(uint256 totalAmount, uint256 claimedAmount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration, bool isRevocable, bool isRevoked)
func (_CosineToken *CosineTokenSession) VestingSchedules(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	ClaimedAmount   *big.Int
	StartTime       *big.Int
	CliffDuration   *big.Int
	VestingDuration *big.Int
	IsRevocable     bool
	IsRevoked       bool
}, error) {
	return _CosineToken.Contract.VestingSchedules(&_CosineToken.CallOpts, arg0)
}

// VestingSchedules is a free data retrieval call binding the contract method 0xfdb20ccb.
//
// Solidity: function vestingSchedules(address ) view returns(uint256 totalAmount, uint256 claimedAmount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration, bool isRevocable, bool isRevoked)
func (_CosineToken *CosineTokenCallerSession) VestingSchedules(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	ClaimedAmount   *big.Int
	StartTime       *big.Int
	CliffDuration   *big.Int
	VestingDuration *big.Int
	IsRevocable     bool
	IsRevoked       bool
}, error) {
	return _CosineToken.Contract.VestingSchedules(&_CosineToken.CallOpts, arg0)
}

// VestingShares is a free data retrieval call binding the contract method 0x21edb61b.
//
// Solidity: function vestingShares(address ) view returns(uint256)
func (_CosineToken *CosineTokenCaller) VestingShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosineToken.contract.Call(opts, &out, "vestingShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingShares is a free data retrieval call binding the contract method 0x21edb61b.
//
// Solidity: function vestingShares(address ) view returns(uint256)
func (_CosineToken *CosineTokenSession) VestingShares(arg0 common.Address) (*big.Int, error) {
	return _CosineToken.Contract.VestingShares(&_CosineToken.CallOpts, arg0)
}

// VestingShares is a free data retrieval call binding the contract method 0x21edb61b.
//
// Solidity: function vestingShares(address ) view returns(uint256)
func (_CosineToken *CosineTokenCallerSession) VestingShares(arg0 common.Address) (*big.Int, error) {
	return _CosineToken.Contract.VestingShares(&_CosineToken.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Approve(&_CosineToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Approve(&_CosineToken.TransactOpts, spender, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0x17727d35.
//
// Solidity: function bridge(uint256 chainId, uint256 amount, bool isLock) returns()
func (_CosineToken *CosineTokenTransactor) Bridge(opts *bind.TransactOpts, chainId *big.Int, amount *big.Int, isLock bool) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "bridge", chainId, amount, isLock)
}

// Bridge is a paid mutator transaction binding the contract method 0x17727d35.
//
// Solidity: function bridge(uint256 chainId, uint256 amount, bool isLock) returns()
func (_CosineToken *CosineTokenSession) Bridge(chainId *big.Int, amount *big.Int, isLock bool) (*types.Transaction, error) {
	return _CosineToken.Contract.Bridge(&_CosineToken.TransactOpts, chainId, amount, isLock)
}

// Bridge is a paid mutator transaction binding the contract method 0x17727d35.
//
// Solidity: function bridge(uint256 chainId, uint256 amount, bool isLock) returns()
func (_CosineToken *CosineTokenTransactorSession) Bridge(chainId *big.Int, amount *big.Int, isLock bool) (*types.Transaction, error) {
	return _CosineToken.Contract.Bridge(&_CosineToken.TransactOpts, chainId, amount, isLock)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CosineToken *CosineTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CosineToken *CosineTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Burn(&_CosineToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CosineToken *CosineTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Burn(&_CosineToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_CosineToken *CosineTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_CosineToken *CosineTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.BurnFrom(&_CosineToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_CosineToken *CosineTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.BurnFrom(&_CosineToken.TransactOpts, account, amount)
}

// ClaimVestedTokens is a paid mutator transaction binding the contract method 0xe74f3fbb.
//
// Solidity: function claimVestedTokens() returns()
func (_CosineToken *CosineTokenTransactor) ClaimVestedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "claimVestedTokens")
}

// ClaimVestedTokens is a paid mutator transaction binding the contract method 0xe74f3fbb.
//
// Solidity: function claimVestedTokens() returns()
func (_CosineToken *CosineTokenSession) ClaimVestedTokens() (*types.Transaction, error) {
	return _CosineToken.Contract.ClaimVestedTokens(&_CosineToken.TransactOpts)
}

// ClaimVestedTokens is a paid mutator transaction binding the contract method 0xe74f3fbb.
//
// Solidity: function claimVestedTokens() returns()
func (_CosineToken *CosineTokenTransactorSession) ClaimVestedTokens() (*types.Transaction, error) {
	return _CosineToken.Contract.ClaimVestedTokens(&_CosineToken.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosineToken *CosineTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosineToken *CosineTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.DecreaseAllowance(&_CosineToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosineToken *CosineTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.DecreaseAllowance(&_CosineToken.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.GrantRole(&_CosineToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.GrantRole(&_CosineToken.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosineToken *CosineTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosineToken *CosineTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.IncreaseAllowance(&_CosineToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosineToken *CosineTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.IncreaseAllowance(&_CosineToken.TransactOpts, spender, addedValue)
}

// InitializeDistributions is a paid mutator transaction binding the contract method 0x5a71c528.
//
// Solidity: function initializeDistributions(address foundation, address advisors, address developers, address privateSale, address publicSale, address networkRewards) returns()
func (_CosineToken *CosineTokenTransactor) InitializeDistributions(opts *bind.TransactOpts, foundation common.Address, advisors common.Address, developers common.Address, privateSale common.Address, publicSale common.Address, networkRewards common.Address) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "initializeDistributions", foundation, advisors, developers, privateSale, publicSale, networkRewards)
}

// InitializeDistributions is a paid mutator transaction binding the contract method 0x5a71c528.
//
// Solidity: function initializeDistributions(address foundation, address advisors, address developers, address privateSale, address publicSale, address networkRewards) returns()
func (_CosineToken *CosineTokenSession) InitializeDistributions(foundation common.Address, advisors common.Address, developers common.Address, privateSale common.Address, publicSale common.Address, networkRewards common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.InitializeDistributions(&_CosineToken.TransactOpts, foundation, advisors, developers, privateSale, publicSale, networkRewards)
}

// InitializeDistributions is a paid mutator transaction binding the contract method 0x5a71c528.
//
// Solidity: function initializeDistributions(address foundation, address advisors, address developers, address privateSale, address publicSale, address networkRewards) returns()
func (_CosineToken *CosineTokenTransactorSession) InitializeDistributions(foundation common.Address, advisors common.Address, developers common.Address, privateSale common.Address, publicSale common.Address, networkRewards common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.InitializeDistributions(&_CosineToken.TransactOpts, foundation, advisors, developers, privateSale, publicSale, networkRewards)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CosineToken *CosineTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CosineToken *CosineTokenSession) Pause() (*types.Transaction, error) {
	return _CosineToken.Contract.Pause(&_CosineToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CosineToken *CosineTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _CosineToken.Contract.Pause(&_CosineToken.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RenounceRole(&_CosineToken.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RenounceRole(&_CosineToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RevokeRole(&_CosineToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CosineToken *CosineTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RevokeRole(&_CosineToken.TransactOpts, role, account)
}

// RevokeVestingSchedule is a paid mutator transaction binding the contract method 0xf035a272.
//
// Solidity: function revokeVestingSchedule(address beneficiary) returns()
func (_CosineToken *CosineTokenTransactor) RevokeVestingSchedule(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "revokeVestingSchedule", beneficiary)
}

// RevokeVestingSchedule is a paid mutator transaction binding the contract method 0xf035a272.
//
// Solidity: function revokeVestingSchedule(address beneficiary) returns()
func (_CosineToken *CosineTokenSession) RevokeVestingSchedule(beneficiary common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RevokeVestingSchedule(&_CosineToken.TransactOpts, beneficiary)
}

// RevokeVestingSchedule is a paid mutator transaction binding the contract method 0xf035a272.
//
// Solidity: function revokeVestingSchedule(address beneficiary) returns()
func (_CosineToken *CosineTokenTransactorSession) RevokeVestingSchedule(beneficiary common.Address) (*types.Transaction, error) {
	return _CosineToken.Contract.RevokeVestingSchedule(&_CosineToken.TransactOpts, beneficiary)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Transfer(&_CosineToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.Transfer(&_CosineToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.TransferFrom(&_CosineToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CosineToken *CosineTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosineToken.Contract.TransferFrom(&_CosineToken.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CosineToken *CosineTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosineToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CosineToken *CosineTokenSession) Unpause() (*types.Transaction, error) {
	return _CosineToken.Contract.Unpause(&_CosineToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CosineToken *CosineTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _CosineToken.Contract.Unpause(&_CosineToken.TransactOpts)
}

// CosineTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CosineToken contract.
type CosineTokenApprovalIterator struct {
	Event *CosineTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenApproval represents a Approval event raised by the CosineToken contract.
type CosineTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosineToken *CosineTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CosineTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenApprovalIterator{contract: _CosineToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosineToken *CosineTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CosineTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenApproval)
				if err := _CosineToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosineToken *CosineTokenFilterer) ParseApproval(log types.Log) (*CosineTokenApproval, error) {
	event := new(CosineTokenApproval)
	if err := _CosineToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CosineToken contract.
type CosineTokenPausedIterator struct {
	Event *CosineTokenPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenPaused represents a Paused event raised by the CosineToken contract.
type CosineTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CosineToken *CosineTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*CosineTokenPausedIterator, error) {

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CosineTokenPausedIterator{contract: _CosineToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CosineToken *CosineTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CosineTokenPaused) (event.Subscription, error) {

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenPaused)
				if err := _CosineToken.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CosineToken *CosineTokenFilterer) ParsePaused(log types.Log) (*CosineTokenPaused, error) {
	event := new(CosineTokenPaused)
	if err := _CosineToken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CosineToken contract.
type CosineTokenRoleAdminChangedIterator struct {
	Event *CosineTokenRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenRoleAdminChanged represents a RoleAdminChanged event raised by the CosineToken contract.
type CosineTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CosineToken *CosineTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CosineTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenRoleAdminChangedIterator{contract: _CosineToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CosineToken *CosineTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CosineTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenRoleAdminChanged)
				if err := _CosineToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CosineToken *CosineTokenFilterer) ParseRoleAdminChanged(log types.Log) (*CosineTokenRoleAdminChanged, error) {
	event := new(CosineTokenRoleAdminChanged)
	if err := _CosineToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CosineToken contract.
type CosineTokenRoleGrantedIterator struct {
	Event *CosineTokenRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenRoleGranted represents a RoleGranted event raised by the CosineToken contract.
type CosineTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CosineTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenRoleGrantedIterator{contract: _CosineToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CosineTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenRoleGranted)
				if err := _CosineToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) ParseRoleGranted(log types.Log) (*CosineTokenRoleGranted, error) {
	event := new(CosineTokenRoleGranted)
	if err := _CosineToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CosineToken contract.
type CosineTokenRoleRevokedIterator struct {
	Event *CosineTokenRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenRoleRevoked represents a RoleRevoked event raised by the CosineToken contract.
type CosineTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CosineTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenRoleRevokedIterator{contract: _CosineToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CosineTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenRoleRevoked)
				if err := _CosineToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CosineToken *CosineTokenFilterer) ParseRoleRevoked(log types.Log) (*CosineTokenRoleRevoked, error) {
	event := new(CosineTokenRoleRevoked)
	if err := _CosineToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenTokensBridgedIterator is returned from FilterTokensBridged and is used to iterate over the raw logs and unpacked data for TokensBridged events raised by the CosineToken contract.
type CosineTokenTokensBridgedIterator struct {
	Event *CosineTokenTokensBridged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenTokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenTokensBridged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenTokensBridged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenTokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenTokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenTokensBridged represents a TokensBridged event raised by the CosineToken contract.
type CosineTokenTokensBridged struct {
	ChainId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokensBridged is a free log retrieval operation binding the contract event 0x9ae36a2ae9e50dd27b843521a327f5c2a422a56649281bc6c5754f15b4a01b7e.
//
// Solidity: event TokensBridged(uint256 indexed chainId, uint256 amount)
func (_CosineToken *CosineTokenFilterer) FilterTokensBridged(opts *bind.FilterOpts, chainId []*big.Int) (*CosineTokenTokensBridgedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "TokensBridged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenTokensBridgedIterator{contract: _CosineToken.contract, event: "TokensBridged", logs: logs, sub: sub}, nil
}

// WatchTokensBridged is a free log subscription operation binding the contract event 0x9ae36a2ae9e50dd27b843521a327f5c2a422a56649281bc6c5754f15b4a01b7e.
//
// Solidity: event TokensBridged(uint256 indexed chainId, uint256 amount)
func (_CosineToken *CosineTokenFilterer) WatchTokensBridged(opts *bind.WatchOpts, sink chan<- *CosineTokenTokensBridged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "TokensBridged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenTokensBridged)
				if err := _CosineToken.contract.UnpackLog(event, "TokensBridged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensBridged is a log parse operation binding the contract event 0x9ae36a2ae9e50dd27b843521a327f5c2a422a56649281bc6c5754f15b4a01b7e.
//
// Solidity: event TokensBridged(uint256 indexed chainId, uint256 amount)
func (_CosineToken *CosineTokenFilterer) ParseTokensBridged(log types.Log) (*CosineTokenTokensBridged, error) {
	event := new(CosineTokenTokensBridged)
	if err := _CosineToken.contract.UnpackLog(event, "TokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenTokensVestedIterator is returned from FilterTokensVested and is used to iterate over the raw logs and unpacked data for TokensVested events raised by the CosineToken contract.
type CosineTokenTokensVestedIterator struct {
	Event *CosineTokenTokensVested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenTokensVestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenTokensVested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenTokensVested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenTokensVestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenTokensVestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenTokensVested represents a TokensVested event raised by the CosineToken contract.
type CosineTokenTokensVested struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensVested is a free log retrieval operation binding the contract event 0xd4691cc79b8fc72aac1e8c0d15a2ca06d71a386c983f815266f0fce713dc5ea7.
//
// Solidity: event TokensVested(address indexed beneficiary, uint256 amount)
func (_CosineToken *CosineTokenFilterer) FilterTokensVested(opts *bind.FilterOpts, beneficiary []common.Address) (*CosineTokenTokensVestedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "TokensVested", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenTokensVestedIterator{contract: _CosineToken.contract, event: "TokensVested", logs: logs, sub: sub}, nil
}

// WatchTokensVested is a free log subscription operation binding the contract event 0xd4691cc79b8fc72aac1e8c0d15a2ca06d71a386c983f815266f0fce713dc5ea7.
//
// Solidity: event TokensVested(address indexed beneficiary, uint256 amount)
func (_CosineToken *CosineTokenFilterer) WatchTokensVested(opts *bind.WatchOpts, sink chan<- *CosineTokenTokensVested, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "TokensVested", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenTokensVested)
				if err := _CosineToken.contract.UnpackLog(event, "TokensVested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensVested is a log parse operation binding the contract event 0xd4691cc79b8fc72aac1e8c0d15a2ca06d71a386c983f815266f0fce713dc5ea7.
//
// Solidity: event TokensVested(address indexed beneficiary, uint256 amount)
func (_CosineToken *CosineTokenFilterer) ParseTokensVested(log types.Log) (*CosineTokenTokensVested, error) {
	event := new(CosineTokenTokensVested)
	if err := _CosineToken.contract.UnpackLog(event, "TokensVested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CosineToken contract.
type CosineTokenTransferIterator struct {
	Event *CosineTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenTransfer represents a Transfer event raised by the CosineToken contract.
type CosineTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosineToken *CosineTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CosineTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenTransferIterator{contract: _CosineToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosineToken *CosineTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CosineTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenTransfer)
				if err := _CosineToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosineToken *CosineTokenFilterer) ParseTransfer(log types.Log) (*CosineTokenTransfer, error) {
	event := new(CosineTokenTransfer)
	if err := _CosineToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CosineToken contract.
type CosineTokenUnpausedIterator struct {
	Event *CosineTokenUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenUnpaused represents a Unpaused event raised by the CosineToken contract.
type CosineTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CosineToken *CosineTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CosineTokenUnpausedIterator, error) {

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CosineTokenUnpausedIterator{contract: _CosineToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CosineToken *CosineTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CosineTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenUnpaused)
				if err := _CosineToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CosineToken *CosineTokenFilterer) ParseUnpaused(log types.Log) (*CosineTokenUnpaused, error) {
	event := new(CosineTokenUnpaused)
	if err := _CosineToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenVestingScheduleCreatedIterator is returned from FilterVestingScheduleCreated and is used to iterate over the raw logs and unpacked data for VestingScheduleCreated events raised by the CosineToken contract.
type CosineTokenVestingScheduleCreatedIterator struct {
	Event *CosineTokenVestingScheduleCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenVestingScheduleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenVestingScheduleCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenVestingScheduleCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenVestingScheduleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenVestingScheduleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenVestingScheduleCreated represents a VestingScheduleCreated event raised by the CosineToken contract.
type CosineTokenVestingScheduleCreated struct {
	Beneficiary     common.Address
	Amount          *big.Int
	StartTime       *big.Int
	CliffDuration   *big.Int
	VestingDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleCreated is a free log retrieval operation binding the contract event 0x47d46f58c7d60d60571092cff5715b8c8c5bc1347db2bb9890e639bf3680ba7b.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 amount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration)
func (_CosineToken *CosineTokenFilterer) FilterVestingScheduleCreated(opts *bind.FilterOpts, beneficiary []common.Address) (*CosineTokenVestingScheduleCreatedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "VestingScheduleCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenVestingScheduleCreatedIterator{contract: _CosineToken.contract, event: "VestingScheduleCreated", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleCreated is a free log subscription operation binding the contract event 0x47d46f58c7d60d60571092cff5715b8c8c5bc1347db2bb9890e639bf3680ba7b.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 amount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration)
func (_CosineToken *CosineTokenFilterer) WatchVestingScheduleCreated(opts *bind.WatchOpts, sink chan<- *CosineTokenVestingScheduleCreated, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "VestingScheduleCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenVestingScheduleCreated)
				if err := _CosineToken.contract.UnpackLog(event, "VestingScheduleCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVestingScheduleCreated is a log parse operation binding the contract event 0x47d46f58c7d60d60571092cff5715b8c8c5bc1347db2bb9890e639bf3680ba7b.
//
// Solidity: event VestingScheduleCreated(address indexed beneficiary, uint256 amount, uint256 startTime, uint256 cliffDuration, uint256 vestingDuration)
func (_CosineToken *CosineTokenFilterer) ParseVestingScheduleCreated(log types.Log) (*CosineTokenVestingScheduleCreated, error) {
	event := new(CosineTokenVestingScheduleCreated)
	if err := _CosineToken.contract.UnpackLog(event, "VestingScheduleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosineTokenVestingScheduleRevokedIterator is returned from FilterVestingScheduleRevoked and is used to iterate over the raw logs and unpacked data for VestingScheduleRevoked events raised by the CosineToken contract.
type CosineTokenVestingScheduleRevokedIterator struct {
	Event *CosineTokenVestingScheduleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosineTokenVestingScheduleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosineTokenVestingScheduleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosineTokenVestingScheduleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosineTokenVestingScheduleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosineTokenVestingScheduleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosineTokenVestingScheduleRevoked represents a VestingScheduleRevoked event raised by the CosineToken contract.
type CosineTokenVestingScheduleRevoked struct {
	Beneficiary     common.Address
	AmountRecovered *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRevoked is a free log retrieval operation binding the contract event 0xb71d84e80ebac87b433939f1a1191d542623a94ff5b940314af5589461f35275.
//
// Solidity: event VestingScheduleRevoked(address indexed beneficiary, uint256 amountRecovered)
func (_CosineToken *CosineTokenFilterer) FilterVestingScheduleRevoked(opts *bind.FilterOpts, beneficiary []common.Address) (*CosineTokenVestingScheduleRevokedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.FilterLogs(opts, "VestingScheduleRevoked", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CosineTokenVestingScheduleRevokedIterator{contract: _CosineToken.contract, event: "VestingScheduleRevoked", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRevoked is a free log subscription operation binding the contract event 0xb71d84e80ebac87b433939f1a1191d542623a94ff5b940314af5589461f35275.
//
// Solidity: event VestingScheduleRevoked(address indexed beneficiary, uint256 amountRecovered)
func (_CosineToken *CosineTokenFilterer) WatchVestingScheduleRevoked(opts *bind.WatchOpts, sink chan<- *CosineTokenVestingScheduleRevoked, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CosineToken.contract.WatchLogs(opts, "VestingScheduleRevoked", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosineTokenVestingScheduleRevoked)
				if err := _CosineToken.contract.UnpackLog(event, "VestingScheduleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVestingScheduleRevoked is a log parse operation binding the contract event 0xb71d84e80ebac87b433939f1a1191d542623a94ff5b940314af5589461f35275.
//
// Solidity: event VestingScheduleRevoked(address indexed beneficiary, uint256 amountRecovered)
func (_CosineToken *CosineTokenFilterer) ParseVestingScheduleRevoked(log types.Log) (*CosineTokenVestingScheduleRevoked, error) {
	event := new(CosineTokenVestingScheduleRevoked)
	if err := _CosineToken.contract.UnpackLog(event, "VestingScheduleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
