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

// CreditVerificationVector2D is an auto generated low-level Go binding around an user-defined struct.
type CreditVerificationVector2D struct {
	X *big.Int
	Y *big.Int
}

// CreditVerificationMetaData contains all meta data concerning the CreditVerification contract.
var CreditVerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosineTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"walletLinkingAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCost\",\"type\":\"uint256\"}],\"name\":\"AccumulatedCostReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passedThreshold\",\"type\":\"bool\"}],\"name\":\"CreditVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMultiplier\",\"type\":\"uint256\"}],\"name\":\"FeeMultiplierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTrusted\",\"type\":\"bool\"}],\"name\":\"FeedbackProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"vectorX\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"vectorY\",\"type\":\"int256\"}],\"name\":\"WalletVectorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accumulatedCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"internalType\":\"structCreditVerification.Vector2D\",\"name\":\"v1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"internalType\":\"structCreditVerification.Vector2D\",\"name\":\"v2\",\"type\":\"tuple\"}],\"name\":\"calculateCosineSimilarity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"similarity\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"}],\"name\":\"calculateVerificationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"checkFeedbackRequirement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"requiresFeedback\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosineToken\",\"outputs\":[{\"internalType\":\"contractICosineToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasFeedbackPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastCheckedWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lastVerificationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVerificationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"postTransactionFeedback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isTrusted\",\"type\":\"bool\"}],\"name\":\"provideTransactionFeedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"vectorX\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"vectorY\",\"type\":\"int256\"}],\"name\":\"submitWalletVector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"additionalCost\",\"type\":\"uint256\"}],\"name\":\"updateAccumulatedCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMultiplier\",\"type\":\"uint256\"}],\"name\":\"updateFeeMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinFee\",\"type\":\"uint256\"}],\"name\":\"updateMinimumFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int256\",\"name\":\"x\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"internalType\":\"structCreditVerification.Vector2D\",\"name\":\"thresholdVector\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"thresholdSimilarity\",\"type\":\"int256\"}],\"name\":\"verifyCreditScore\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"similarity\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletLinking\",\"outputs\":[{\"internalType\":\"contractWalletLinking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526078600355662386f26fc100006004553480156200002157600080fd5b5060405162001a1538038062001a15833981016040819052620000449162000270565b6001805460ff191690556001600160a01b038216620000aa5760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420746f6b656e2061646472657373000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b038116620001025760405162461bcd60e51b815260206004820152601e60248201527f496e76616c69642077616c6c6574206c696e6b696e67206164647265737300006044820152606401620000a1565b6200010f600033620001a3565b6200013b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533620001a3565b620001677f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892633620001a3565b60018054610100600160a81b0319166101006001600160a01b0394851602179055600280546001600160a01b03191691909216179055620002a8565b620001af8282620001b3565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16620001af576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556200020f3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b80516001600160a01b03811681146200026b57600080fd5b919050565b600080604083850312156200028457600080fd5b6200028f8362000253565b91506200029f6020840162000253565b90509250929050565b61175d80620002b86000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806376e062831161010f578063c49baebe116100a2578063e8f5b97e11610071578063e8f5b97e14610470578063ec3492da14610490578063f3f9861e146104a3578063f5b541a6146104cf57600080fd5b8063c49baebe1461040d578063d1eb06a814610434578063d547741f14610454578063e5a70ef71461046757600080fd5b8063940777ad116100de578063940777ad146103c7578063a217fddf146103da578063bdd9ed31146103e2578063c17dd765146103fa57600080fd5b806376e06283146103615780638456cb59146103845780638da3608b1461038c57806391d14854146103b457600080fd5b806336568abe11610187578063704d0f3411610156578063704d0f34146102eb57806374f244ee1461030b578063755ea6071461031e57806375b238fc1461034c57600080fd5b806336568abe146102b25780633f4ba83a146102c55780634e421efa146102cd5780635c975abb146102e057600080fd5b80632164077b116101c35780632164077b1461023e578063248a9ca3146102515780632e9923f2146102745780632f2ff15d1461029f57600080fd5b806301ffc9a7146101ea578063173b411414610212578063194614df14610227575b600080fd5b6101fd6101f836600461126a565b6104f6565b60405190151581526020015b60405180910390f35b610225610220366004611294565b61052d565b005b61023060045481565b604051908152602001610209565b61023061024c3660046112b6565b6105e7565b61023061025f3660046112b6565b60009081526020819052604090206001015490565b600254610287906001600160a01b031681565b6040516001600160a01b039091168152602001610209565b6102256102ad3660046112eb565b61062e565b6102256102c03660046112eb565b610658565b6102256106d6565b6102256102db3660046112b6565b6106f9565b60015460ff166101fd565b6102306102f9366004611317565b60096020526000908152604090205481565b6102256103193660046112b6565b610717565b6101fd61032c366004611332565b600760209081526000928352604080842090915290825290205460ff1681565b61023060008051602061170883398151915281565b6101fd61036f366004611317565b60086020526000908152604090205460ff1681565b6102256107bc565b61039f61039a36600461135c565b6107dc565b60408051928352901515602083015201610209565b6101fd6103c23660046112eb565b6109e7565b6102256103d536600461139d565b610a10565b610230600081565b6001546102879061010090046001600160a01b031681565b6102256104083660046113d7565b610a7b565b6102307f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b6102306104423660046112b6565b60066020526000908152604090205481565b6102256104623660046112eb565b610be1565b61023060035481565b61023061047e3660046112b6565b60056020526000908152604090205481565b61023061049e36600461147a565b610c06565b6101fd6104b1366004611317565b6001600160a01b031660009081526008602052604090205460ff1690565b6102307f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b60006001600160e01b03198216637965db0b60e01b148061052757506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892661055781610d6f565b8261057d5760405162461bcd60e51b8152600401610574906114a6565b60405180910390fd5b6000838152600560205260409020546105969083610d79565b600084815260056020526040908190208290555184917fa498743dbfd7c09b930a25f95c0fb18d57d463927d531532e98af5eea161a297916105da91815260200190565b60405180910390a2505050565b600081815260056020526040812054808203610607575050600454919050565b610627606461062160035484610d8590919063ffffffff16565b90610d91565b9392505050565b60008281526020819052604090206001015461064981610d6f565b6106538383610d9d565b505050565b6001600160a01b03811633146106c85760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610574565b6106d28282610e21565b5050565b6000805160206117088339815191526106ee81610d6f565b6106f6610e86565b50565b60008051602061170883398151915261071181610d6f565b50600455565b60008051602061170883398151915261072f81610d6f565b60648210156107805760405162461bcd60e51b815260206004820152601f60248201527f4d756c7469706c696572206d757374206265206174206c6561737420312e30006044820152606401610574565b60038290556040518281527fc422d1349750089c716e39be62193f9bbd759f3b17b6dcd96538dcb7fc5bf0219060200160405180910390a15050565b6000805160206117088339815191526107d481610d6f565b6106f6610ed8565b6000806107e7610f13565b846108045760405162461bcd60e51b8152600401610574906114a6565b600061080f866105e7565b6001546040516323b872dd60e01b81523360048201523060248201526044810183905291925061010090046001600160a01b0316906323b872dd906064016020604051808303816000875af115801561086c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089091906114d1565b6108d25760405162461bcd60e51b8152602060048201526013602482015272119959481d1c985b9cd9995c8819985a5b1959606a1b6044820152606401610574565b600154604051630852cd8d60e31b8152600481018390526101009091046001600160a01b0316906342966c6890602401600060405180830381600087803b15801561091c57600080fd5b505af1158015610930573d6000803e3d6000fd5b5050336000908152600860209081526040808320805460ff19166001179055600990915281208990559150610966905087610f5b565b905061097b8161049e368990038901896114ee565b600088815260056020908152604080832083905560068252918290204290558151858152888412159181018290529296509450889133917f116024e715d3fa7124b8639232501ea19c4a9d50eecab039a79a1aaaeb240d6d910160405180910390a35050935093915050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c98926610a3a81610d6f565b604080518481526020810184905285917f26b9f6b5b621a14c92c48918075a8b027399cf129f9fd8d3866052bca90f1c62910160405180910390a250505050565b610a83610f13565b81610aa05760405162461bcd60e51b8152600401610574906114a6565b3360009081526008602052604090205460ff16610aff5760405162461bcd60e51b815260206004820152601f60248201527f4e6f2070656e64696e6720666565646261636b20726571756972656d656e74006044820152606401610574565b336000908152600960205260409020548214610b6e5760405162461bcd60e51b815260206004820152602860248201527f466565646261636b206d75737420626520666f72206c61737420636865636b6560448201526719081dd85b1b195d60c21b6064820152608401610574565b3360008181526007602090815260408083208684528252808320805486151560ff19918216811790925585855260088452938290208054909416909355519182528492917f6273ed15405e9610be646ed2f9d0dac0d880a7ec99ce2b08145506b87ee56780910160405180910390a35050565b600082815260208190526040902060010154610bfc81610d6f565b6106538383610e21565b600080670de0b6b3a764000083602001518560200151610c269190611520565b84518651610c349190611520565b610c3e9190611550565b610c48919061158e565b6020850151909150600090670de0b6b3a764000090610c679080611520565b8651610c739080611520565b610c7d9190611550565b610c87919061158e565b6020850151909150600090670de0b6b3a764000090610ca69080611520565b8651610cb29080611520565b610cbc9190611550565b610cc6919061158e565b90506000610cd383610fb1565b90506000610ce083610fb1565b9050811580610ced575080155b15610d0057600095505050505050610527565b610d0a8183611520565b610d1c86670de0b6b3a7640000611520565b610d26919061158e565b9550670de0b6b3a7640000861315610d4857670de0b6b3a76400009550610d64565b670de0b6b3a7640000861215610d6457670de0b6b3a764000095505b505050505092915050565b6106f6813361101a565b600061062782846115bc565b600061062782846115cf565b600061062782846115e6565b610da782826109e7565b6106d2576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055610ddd3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610e2b82826109e7565b156106d2576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b610e8e611073565b6001805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b610ee0610f13565b6001805460ff1916811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610ebb565b60015460ff1615610f595760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610574565b565b60408051808201909152600080825260208201526000610f7e620f4240846115fa565b610f8f9066038d7ea4c68000611520565b60408051808201909152908152670de0b6b3a764000060208201529392505050565b6000808213610fc257506000919050565b81806000610fd16002836115e6565b610fdc9060016115bc565b90505b8181101561101257905080600281610ff781866115e6565b61100191906115bc565b61100b91906115e6565b9050610fdf565b509392505050565b61102482826109e7565b6106d257611031816110bc565b61103c8360206110ce565b60405160200161104d929190611632565b60408051601f198184030181529082905262461bcd60e51b8252610574916004016116a7565b60015460ff16610f595760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610574565b60606105276001600160a01b03831660145b606060006110dd8360026115cf565b6110e89060026115bc565b67ffffffffffffffff81111561110057611100611407565b6040519080825280601f01601f19166020018201604052801561112a576020820181803683370190505b509050600360fc1b81600081518110611145576111456116da565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611174576111746116da565b60200101906001600160f81b031916908160001a90535060006111988460026115cf565b6111a39060016115bc565b90505b600181111561121b576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106111d7576111d76116da565b1a60f81b8282815181106111ed576111ed6116da565b60200101906001600160f81b031916908160001a90535060049490941c93611214816116f0565b90506111a6565b5083156106275760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610574565b60006020828403121561127c57600080fd5b81356001600160e01b03198116811461062757600080fd5b600080604083850312156112a757600080fd5b50508035926020909101359150565b6000602082840312156112c857600080fd5b5035919050565b80356001600160a01b03811681146112e657600080fd5b919050565b600080604083850312156112fe57600080fd5b8235915061130e602084016112cf565b90509250929050565b60006020828403121561132957600080fd5b610627826112cf565b6000806040838503121561134557600080fd5b61134e836112cf565b946020939093013593505050565b6000806000838503608081121561137257600080fd5b843593506040601f198201121561138857600080fd5b50602084019150606084013590509250925092565b6000806000606084860312156113b257600080fd5b505081359360208301359350604090920135919050565b80151581146106f657600080fd5b600080604083850312156113ea57600080fd5b8235915060208301356113fc816113c9565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60006040828403121561142f57600080fd5b6040516040810181811067ffffffffffffffff8211171561146057634e487b7160e01b600052604160045260246000fd5b604052823581526020928301359281019290925250919050565b6000806080838503121561148d57600080fd5b611497848461141d565b915061130e846040850161141d565b602080825260119082015270125b9d985b1a5908130c881dd85b1b195d607a1b604082015260600190565b6000602082840312156114e357600080fd5b8151610627816113c9565b60006040828403121561150057600080fd5b610627838361141d565b634e487b7160e01b600052601160045260246000fd5b80820260008212600160ff1b8414161561153c5761153c61150a565b81810583148215176105275761052761150a565b80820182811260008312801582168215821617156115705761157061150a565b505092915050565b634e487b7160e01b600052601260045260246000fd5b60008261159d5761159d611578565b600160ff1b8214600019841416156115b7576115b761150a565b500590565b808201808211156105275761052761150a565b80820281158282048414176105275761052761150a565b6000826115f5576115f5611578565b500490565b60008261160957611609611578565b500690565b60005b83811015611629578181015183820152602001611611565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e742000000000000000000081526000835161166a81601785016020880161160e565b7001034b99036b4b9b9b4b733903937b6329607d1b601791840191820152835161169b81602884016020880161160e565b01602801949350505050565b60208152600082518060208401526116c681604085016020870161160e565b601f01601f19169190910160400192915050565b634e487b7160e01b600052603260045260246000fd5b6000816116ff576116ff61150a565b50600019019056fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a26469706673582212206bf0d86fc4f26c991bad4e5f995e07b2bfca68c029f9cdc25495f232ec4ada2e64736f6c63430008130033",
}

// CreditVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditVerificationMetaData.ABI instead.
var CreditVerificationABI = CreditVerificationMetaData.ABI

// CreditVerificationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreditVerificationMetaData.Bin instead.
var CreditVerificationBin = CreditVerificationMetaData.Bin

// DeployCreditVerification deploys a new Ethereum contract, binding an instance of CreditVerification to it.
func DeployCreditVerification(auth *bind.TransactOpts, backend bind.ContractBackend, cosineTokenAddress common.Address, walletLinkingAddress common.Address) (common.Address, *types.Transaction, *CreditVerification, error) {
	parsed, err := CreditVerificationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreditVerificationBin), backend, cosineTokenAddress, walletLinkingAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreditVerification{CreditVerificationCaller: CreditVerificationCaller{contract: contract}, CreditVerificationTransactor: CreditVerificationTransactor{contract: contract}, CreditVerificationFilterer: CreditVerificationFilterer{contract: contract}}, nil
}

// CreditVerification is an auto generated Go binding around an Ethereum contract.
type CreditVerification struct {
	CreditVerificationCaller     // Read-only binding to the contract
	CreditVerificationTransactor // Write-only binding to the contract
	CreditVerificationFilterer   // Log filterer for contract events
}

// CreditVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditVerificationSession struct {
	Contract     *CreditVerification // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CreditVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditVerificationCallerSession struct {
	Contract *CreditVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CreditVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditVerificationTransactorSession struct {
	Contract     *CreditVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CreditVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditVerificationRaw struct {
	Contract *CreditVerification // Generic contract binding to access the raw methods on
}

// CreditVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditVerificationCallerRaw struct {
	Contract *CreditVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// CreditVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditVerificationTransactorRaw struct {
	Contract *CreditVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditVerification creates a new instance of CreditVerification, bound to a specific deployed contract.
func NewCreditVerification(address common.Address, backend bind.ContractBackend) (*CreditVerification, error) {
	contract, err := bindCreditVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditVerification{CreditVerificationCaller: CreditVerificationCaller{contract: contract}, CreditVerificationTransactor: CreditVerificationTransactor{contract: contract}, CreditVerificationFilterer: CreditVerificationFilterer{contract: contract}}, nil
}

// NewCreditVerificationCaller creates a new read-only instance of CreditVerification, bound to a specific deployed contract.
func NewCreditVerificationCaller(address common.Address, caller bind.ContractCaller) (*CreditVerificationCaller, error) {
	contract, err := bindCreditVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationCaller{contract: contract}, nil
}

// NewCreditVerificationTransactor creates a new write-only instance of CreditVerification, bound to a specific deployed contract.
func NewCreditVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditVerificationTransactor, error) {
	contract, err := bindCreditVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationTransactor{contract: contract}, nil
}

// NewCreditVerificationFilterer creates a new log filterer instance of CreditVerification, bound to a specific deployed contract.
func NewCreditVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditVerificationFilterer, error) {
	contract, err := bindCreditVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationFilterer{contract: contract}, nil
}

// bindCreditVerification binds a generic wrapper to an already deployed contract.
func bindCreditVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditVerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditVerification *CreditVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditVerification.Contract.CreditVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditVerification *CreditVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditVerification.Contract.CreditVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditVerification *CreditVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditVerification.Contract.CreditVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditVerification *CreditVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CreditVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditVerification *CreditVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditVerification *CreditVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CreditVerification.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) ADMINROLE() ([32]byte, error) {
	return _CreditVerification.Contract.ADMINROLE(&_CreditVerification.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) ADMINROLE() ([32]byte, error) {
	return _CreditVerification.Contract.ADMINROLE(&_CreditVerification.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CreditVerification.Contract.DEFAULTADMINROLE(&_CreditVerification.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CreditVerification.Contract.DEFAULTADMINROLE(&_CreditVerification.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) OPERATORROLE() ([32]byte, error) {
	return _CreditVerification.Contract.OPERATORROLE(&_CreditVerification.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) OPERATORROLE() ([32]byte, error) {
	return _CreditVerification.Contract.OPERATORROLE(&_CreditVerification.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) VALIDATORROLE() ([32]byte, error) {
	return _CreditVerification.Contract.VALIDATORROLE(&_CreditVerification.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _CreditVerification.Contract.VALIDATORROLE(&_CreditVerification.CallOpts)
}

// AccumulatedCost is a free data retrieval call binding the contract method 0xe8f5b97e.
//
// Solidity: function accumulatedCost(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationCaller) AccumulatedCost(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "accumulatedCost", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedCost is a free data retrieval call binding the contract method 0xe8f5b97e.
//
// Solidity: function accumulatedCost(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationSession) AccumulatedCost(arg0 [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.AccumulatedCost(&_CreditVerification.CallOpts, arg0)
}

// AccumulatedCost is a free data retrieval call binding the contract method 0xe8f5b97e.
//
// Solidity: function accumulatedCost(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationCallerSession) AccumulatedCost(arg0 [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.AccumulatedCost(&_CreditVerification.CallOpts, arg0)
}

// CalculateCosineSimilarity is a free data retrieval call binding the contract method 0xec3492da.
//
// Solidity: function calculateCosineSimilarity((int256,int256) v1, (int256,int256) v2) pure returns(int256 similarity)
func (_CreditVerification *CreditVerificationCaller) CalculateCosineSimilarity(opts *bind.CallOpts, v1 CreditVerificationVector2D, v2 CreditVerificationVector2D) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "calculateCosineSimilarity", v1, v2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCosineSimilarity is a free data retrieval call binding the contract method 0xec3492da.
//
// Solidity: function calculateCosineSimilarity((int256,int256) v1, (int256,int256) v2) pure returns(int256 similarity)
func (_CreditVerification *CreditVerificationSession) CalculateCosineSimilarity(v1 CreditVerificationVector2D, v2 CreditVerificationVector2D) (*big.Int, error) {
	return _CreditVerification.Contract.CalculateCosineSimilarity(&_CreditVerification.CallOpts, v1, v2)
}

// CalculateCosineSimilarity is a free data retrieval call binding the contract method 0xec3492da.
//
// Solidity: function calculateCosineSimilarity((int256,int256) v1, (int256,int256) v2) pure returns(int256 similarity)
func (_CreditVerification *CreditVerificationCallerSession) CalculateCosineSimilarity(v1 CreditVerificationVector2D, v2 CreditVerificationVector2D) (*big.Int, error) {
	return _CreditVerification.Contract.CalculateCosineSimilarity(&_CreditVerification.CallOpts, v1, v2)
}

// CalculateVerificationFee is a free data retrieval call binding the contract method 0x2164077b.
//
// Solidity: function calculateVerificationFee(bytes32 l2Wallet) view returns(uint256 fee)
func (_CreditVerification *CreditVerificationCaller) CalculateVerificationFee(opts *bind.CallOpts, l2Wallet [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "calculateVerificationFee", l2Wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateVerificationFee is a free data retrieval call binding the contract method 0x2164077b.
//
// Solidity: function calculateVerificationFee(bytes32 l2Wallet) view returns(uint256 fee)
func (_CreditVerification *CreditVerificationSession) CalculateVerificationFee(l2Wallet [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.CalculateVerificationFee(&_CreditVerification.CallOpts, l2Wallet)
}

// CalculateVerificationFee is a free data retrieval call binding the contract method 0x2164077b.
//
// Solidity: function calculateVerificationFee(bytes32 l2Wallet) view returns(uint256 fee)
func (_CreditVerification *CreditVerificationCallerSession) CalculateVerificationFee(l2Wallet [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.CalculateVerificationFee(&_CreditVerification.CallOpts, l2Wallet)
}

// CheckFeedbackRequirement is a free data retrieval call binding the contract method 0xf3f9861e.
//
// Solidity: function checkFeedbackRequirement(address verifier) view returns(bool requiresFeedback)
func (_CreditVerification *CreditVerificationCaller) CheckFeedbackRequirement(opts *bind.CallOpts, verifier common.Address) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "checkFeedbackRequirement", verifier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFeedbackRequirement is a free data retrieval call binding the contract method 0xf3f9861e.
//
// Solidity: function checkFeedbackRequirement(address verifier) view returns(bool requiresFeedback)
func (_CreditVerification *CreditVerificationSession) CheckFeedbackRequirement(verifier common.Address) (bool, error) {
	return _CreditVerification.Contract.CheckFeedbackRequirement(&_CreditVerification.CallOpts, verifier)
}

// CheckFeedbackRequirement is a free data retrieval call binding the contract method 0xf3f9861e.
//
// Solidity: function checkFeedbackRequirement(address verifier) view returns(bool requiresFeedback)
func (_CreditVerification *CreditVerificationCallerSession) CheckFeedbackRequirement(verifier common.Address) (bool, error) {
	return _CreditVerification.Contract.CheckFeedbackRequirement(&_CreditVerification.CallOpts, verifier)
}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_CreditVerification *CreditVerificationCaller) CosineToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "cosineToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_CreditVerification *CreditVerificationSession) CosineToken() (common.Address, error) {
	return _CreditVerification.Contract.CosineToken(&_CreditVerification.CallOpts)
}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_CreditVerification *CreditVerificationCallerSession) CosineToken() (common.Address, error) {
	return _CreditVerification.Contract.CosineToken(&_CreditVerification.CallOpts)
}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() view returns(uint256)
func (_CreditVerification *CreditVerificationCaller) FeeMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "feeMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() view returns(uint256)
func (_CreditVerification *CreditVerificationSession) FeeMultiplier() (*big.Int, error) {
	return _CreditVerification.Contract.FeeMultiplier(&_CreditVerification.CallOpts)
}

// FeeMultiplier is a free data retrieval call binding the contract method 0xe5a70ef7.
//
// Solidity: function feeMultiplier() view returns(uint256)
func (_CreditVerification *CreditVerificationCallerSession) FeeMultiplier() (*big.Int, error) {
	return _CreditVerification.Contract.FeeMultiplier(&_CreditVerification.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CreditVerification.Contract.GetRoleAdmin(&_CreditVerification.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CreditVerification.Contract.GetRoleAdmin(&_CreditVerification.CallOpts, role)
}

// HasFeedbackPending is a free data retrieval call binding the contract method 0x76e06283.
//
// Solidity: function hasFeedbackPending(address ) view returns(bool)
func (_CreditVerification *CreditVerificationCaller) HasFeedbackPending(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "hasFeedbackPending", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasFeedbackPending is a free data retrieval call binding the contract method 0x76e06283.
//
// Solidity: function hasFeedbackPending(address ) view returns(bool)
func (_CreditVerification *CreditVerificationSession) HasFeedbackPending(arg0 common.Address) (bool, error) {
	return _CreditVerification.Contract.HasFeedbackPending(&_CreditVerification.CallOpts, arg0)
}

// HasFeedbackPending is a free data retrieval call binding the contract method 0x76e06283.
//
// Solidity: function hasFeedbackPending(address ) view returns(bool)
func (_CreditVerification *CreditVerificationCallerSession) HasFeedbackPending(arg0 common.Address) (bool, error) {
	return _CreditVerification.Contract.HasFeedbackPending(&_CreditVerification.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CreditVerification *CreditVerificationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CreditVerification *CreditVerificationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CreditVerification.Contract.HasRole(&_CreditVerification.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CreditVerification *CreditVerificationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CreditVerification.Contract.HasRole(&_CreditVerification.CallOpts, role, account)
}

// LastCheckedWallet is a free data retrieval call binding the contract method 0x704d0f34.
//
// Solidity: function lastCheckedWallet(address ) view returns(bytes32)
func (_CreditVerification *CreditVerificationCaller) LastCheckedWallet(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "lastCheckedWallet", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastCheckedWallet is a free data retrieval call binding the contract method 0x704d0f34.
//
// Solidity: function lastCheckedWallet(address ) view returns(bytes32)
func (_CreditVerification *CreditVerificationSession) LastCheckedWallet(arg0 common.Address) ([32]byte, error) {
	return _CreditVerification.Contract.LastCheckedWallet(&_CreditVerification.CallOpts, arg0)
}

// LastCheckedWallet is a free data retrieval call binding the contract method 0x704d0f34.
//
// Solidity: function lastCheckedWallet(address ) view returns(bytes32)
func (_CreditVerification *CreditVerificationCallerSession) LastCheckedWallet(arg0 common.Address) ([32]byte, error) {
	return _CreditVerification.Contract.LastCheckedWallet(&_CreditVerification.CallOpts, arg0)
}

// LastVerificationTime is a free data retrieval call binding the contract method 0xd1eb06a8.
//
// Solidity: function lastVerificationTime(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationCaller) LastVerificationTime(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "lastVerificationTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVerificationTime is a free data retrieval call binding the contract method 0xd1eb06a8.
//
// Solidity: function lastVerificationTime(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationSession) LastVerificationTime(arg0 [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.LastVerificationTime(&_CreditVerification.CallOpts, arg0)
}

// LastVerificationTime is a free data retrieval call binding the contract method 0xd1eb06a8.
//
// Solidity: function lastVerificationTime(bytes32 ) view returns(uint256)
func (_CreditVerification *CreditVerificationCallerSession) LastVerificationTime(arg0 [32]byte) (*big.Int, error) {
	return _CreditVerification.Contract.LastVerificationTime(&_CreditVerification.CallOpts, arg0)
}

// MinVerificationFee is a free data retrieval call binding the contract method 0x194614df.
//
// Solidity: function minVerificationFee() view returns(uint256)
func (_CreditVerification *CreditVerificationCaller) MinVerificationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "minVerificationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVerificationFee is a free data retrieval call binding the contract method 0x194614df.
//
// Solidity: function minVerificationFee() view returns(uint256)
func (_CreditVerification *CreditVerificationSession) MinVerificationFee() (*big.Int, error) {
	return _CreditVerification.Contract.MinVerificationFee(&_CreditVerification.CallOpts)
}

// MinVerificationFee is a free data retrieval call binding the contract method 0x194614df.
//
// Solidity: function minVerificationFee() view returns(uint256)
func (_CreditVerification *CreditVerificationCallerSession) MinVerificationFee() (*big.Int, error) {
	return _CreditVerification.Contract.MinVerificationFee(&_CreditVerification.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditVerification *CreditVerificationCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditVerification *CreditVerificationSession) Paused() (bool, error) {
	return _CreditVerification.Contract.Paused(&_CreditVerification.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CreditVerification *CreditVerificationCallerSession) Paused() (bool, error) {
	return _CreditVerification.Contract.Paused(&_CreditVerification.CallOpts)
}

// PostTransactionFeedback is a free data retrieval call binding the contract method 0x755ea607.
//
// Solidity: function postTransactionFeedback(address , bytes32 ) view returns(bool)
func (_CreditVerification *CreditVerificationCaller) PostTransactionFeedback(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "postTransactionFeedback", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PostTransactionFeedback is a free data retrieval call binding the contract method 0x755ea607.
//
// Solidity: function postTransactionFeedback(address , bytes32 ) view returns(bool)
func (_CreditVerification *CreditVerificationSession) PostTransactionFeedback(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _CreditVerification.Contract.PostTransactionFeedback(&_CreditVerification.CallOpts, arg0, arg1)
}

// PostTransactionFeedback is a free data retrieval call binding the contract method 0x755ea607.
//
// Solidity: function postTransactionFeedback(address , bytes32 ) view returns(bool)
func (_CreditVerification *CreditVerificationCallerSession) PostTransactionFeedback(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _CreditVerification.Contract.PostTransactionFeedback(&_CreditVerification.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreditVerification *CreditVerificationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreditVerification *CreditVerificationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreditVerification.Contract.SupportsInterface(&_CreditVerification.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CreditVerification *CreditVerificationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CreditVerification.Contract.SupportsInterface(&_CreditVerification.CallOpts, interfaceId)
}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_CreditVerification *CreditVerificationCaller) WalletLinking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CreditVerification.contract.Call(opts, &out, "walletLinking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_CreditVerification *CreditVerificationSession) WalletLinking() (common.Address, error) {
	return _CreditVerification.Contract.WalletLinking(&_CreditVerification.CallOpts)
}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_CreditVerification *CreditVerificationCallerSession) WalletLinking() (common.Address, error) {
	return _CreditVerification.Contract.WalletLinking(&_CreditVerification.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.GrantRole(&_CreditVerification.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.GrantRole(&_CreditVerification.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditVerification *CreditVerificationTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditVerification *CreditVerificationSession) Pause() (*types.Transaction, error) {
	return _CreditVerification.Contract.Pause(&_CreditVerification.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CreditVerification *CreditVerificationTransactorSession) Pause() (*types.Transaction, error) {
	return _CreditVerification.Contract.Pause(&_CreditVerification.TransactOpts)
}

// ProvideTransactionFeedback is a paid mutator transaction binding the contract method 0xc17dd765.
//
// Solidity: function provideTransactionFeedback(bytes32 l2Wallet, bool isTrusted) returns()
func (_CreditVerification *CreditVerificationTransactor) ProvideTransactionFeedback(opts *bind.TransactOpts, l2Wallet [32]byte, isTrusted bool) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "provideTransactionFeedback", l2Wallet, isTrusted)
}

// ProvideTransactionFeedback is a paid mutator transaction binding the contract method 0xc17dd765.
//
// Solidity: function provideTransactionFeedback(bytes32 l2Wallet, bool isTrusted) returns()
func (_CreditVerification *CreditVerificationSession) ProvideTransactionFeedback(l2Wallet [32]byte, isTrusted bool) (*types.Transaction, error) {
	return _CreditVerification.Contract.ProvideTransactionFeedback(&_CreditVerification.TransactOpts, l2Wallet, isTrusted)
}

// ProvideTransactionFeedback is a paid mutator transaction binding the contract method 0xc17dd765.
//
// Solidity: function provideTransactionFeedback(bytes32 l2Wallet, bool isTrusted) returns()
func (_CreditVerification *CreditVerificationTransactorSession) ProvideTransactionFeedback(l2Wallet [32]byte, isTrusted bool) (*types.Transaction, error) {
	return _CreditVerification.Contract.ProvideTransactionFeedback(&_CreditVerification.TransactOpts, l2Wallet, isTrusted)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.RenounceRole(&_CreditVerification.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.RenounceRole(&_CreditVerification.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.RevokeRole(&_CreditVerification.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CreditVerification *CreditVerificationTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CreditVerification.Contract.RevokeRole(&_CreditVerification.TransactOpts, role, account)
}

// SubmitWalletVector is a paid mutator transaction binding the contract method 0x940777ad.
//
// Solidity: function submitWalletVector(bytes32 l2Wallet, int256 vectorX, int256 vectorY) returns()
func (_CreditVerification *CreditVerificationTransactor) SubmitWalletVector(opts *bind.TransactOpts, l2Wallet [32]byte, vectorX *big.Int, vectorY *big.Int) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "submitWalletVector", l2Wallet, vectorX, vectorY)
}

// SubmitWalletVector is a paid mutator transaction binding the contract method 0x940777ad.
//
// Solidity: function submitWalletVector(bytes32 l2Wallet, int256 vectorX, int256 vectorY) returns()
func (_CreditVerification *CreditVerificationSession) SubmitWalletVector(l2Wallet [32]byte, vectorX *big.Int, vectorY *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.SubmitWalletVector(&_CreditVerification.TransactOpts, l2Wallet, vectorX, vectorY)
}

// SubmitWalletVector is a paid mutator transaction binding the contract method 0x940777ad.
//
// Solidity: function submitWalletVector(bytes32 l2Wallet, int256 vectorX, int256 vectorY) returns()
func (_CreditVerification *CreditVerificationTransactorSession) SubmitWalletVector(l2Wallet [32]byte, vectorX *big.Int, vectorY *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.SubmitWalletVector(&_CreditVerification.TransactOpts, l2Wallet, vectorX, vectorY)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditVerification *CreditVerificationTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditVerification *CreditVerificationSession) Unpause() (*types.Transaction, error) {
	return _CreditVerification.Contract.Unpause(&_CreditVerification.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CreditVerification *CreditVerificationTransactorSession) Unpause() (*types.Transaction, error) {
	return _CreditVerification.Contract.Unpause(&_CreditVerification.TransactOpts)
}

// UpdateAccumulatedCost is a paid mutator transaction binding the contract method 0x173b4114.
//
// Solidity: function updateAccumulatedCost(bytes32 l2Wallet, uint256 additionalCost) returns()
func (_CreditVerification *CreditVerificationTransactor) UpdateAccumulatedCost(opts *bind.TransactOpts, l2Wallet [32]byte, additionalCost *big.Int) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "updateAccumulatedCost", l2Wallet, additionalCost)
}

// UpdateAccumulatedCost is a paid mutator transaction binding the contract method 0x173b4114.
//
// Solidity: function updateAccumulatedCost(bytes32 l2Wallet, uint256 additionalCost) returns()
func (_CreditVerification *CreditVerificationSession) UpdateAccumulatedCost(l2Wallet [32]byte, additionalCost *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateAccumulatedCost(&_CreditVerification.TransactOpts, l2Wallet, additionalCost)
}

// UpdateAccumulatedCost is a paid mutator transaction binding the contract method 0x173b4114.
//
// Solidity: function updateAccumulatedCost(bytes32 l2Wallet, uint256 additionalCost) returns()
func (_CreditVerification *CreditVerificationTransactorSession) UpdateAccumulatedCost(l2Wallet [32]byte, additionalCost *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateAccumulatedCost(&_CreditVerification.TransactOpts, l2Wallet, additionalCost)
}

// UpdateFeeMultiplier is a paid mutator transaction binding the contract method 0x74f244ee.
//
// Solidity: function updateFeeMultiplier(uint256 newMultiplier) returns()
func (_CreditVerification *CreditVerificationTransactor) UpdateFeeMultiplier(opts *bind.TransactOpts, newMultiplier *big.Int) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "updateFeeMultiplier", newMultiplier)
}

// UpdateFeeMultiplier is a paid mutator transaction binding the contract method 0x74f244ee.
//
// Solidity: function updateFeeMultiplier(uint256 newMultiplier) returns()
func (_CreditVerification *CreditVerificationSession) UpdateFeeMultiplier(newMultiplier *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateFeeMultiplier(&_CreditVerification.TransactOpts, newMultiplier)
}

// UpdateFeeMultiplier is a paid mutator transaction binding the contract method 0x74f244ee.
//
// Solidity: function updateFeeMultiplier(uint256 newMultiplier) returns()
func (_CreditVerification *CreditVerificationTransactorSession) UpdateFeeMultiplier(newMultiplier *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateFeeMultiplier(&_CreditVerification.TransactOpts, newMultiplier)
}

// UpdateMinimumFee is a paid mutator transaction binding the contract method 0x4e421efa.
//
// Solidity: function updateMinimumFee(uint256 newMinFee) returns()
func (_CreditVerification *CreditVerificationTransactor) UpdateMinimumFee(opts *bind.TransactOpts, newMinFee *big.Int) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "updateMinimumFee", newMinFee)
}

// UpdateMinimumFee is a paid mutator transaction binding the contract method 0x4e421efa.
//
// Solidity: function updateMinimumFee(uint256 newMinFee) returns()
func (_CreditVerification *CreditVerificationSession) UpdateMinimumFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateMinimumFee(&_CreditVerification.TransactOpts, newMinFee)
}

// UpdateMinimumFee is a paid mutator transaction binding the contract method 0x4e421efa.
//
// Solidity: function updateMinimumFee(uint256 newMinFee) returns()
func (_CreditVerification *CreditVerificationTransactorSession) UpdateMinimumFee(newMinFee *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.UpdateMinimumFee(&_CreditVerification.TransactOpts, newMinFee)
}

// VerifyCreditScore is a paid mutator transaction binding the contract method 0x8da3608b.
//
// Solidity: function verifyCreditScore(bytes32 l2Wallet, (int256,int256) thresholdVector, int256 thresholdSimilarity) returns(int256 similarity, bool passed)
func (_CreditVerification *CreditVerificationTransactor) VerifyCreditScore(opts *bind.TransactOpts, l2Wallet [32]byte, thresholdVector CreditVerificationVector2D, thresholdSimilarity *big.Int) (*types.Transaction, error) {
	return _CreditVerification.contract.Transact(opts, "verifyCreditScore", l2Wallet, thresholdVector, thresholdSimilarity)
}

// VerifyCreditScore is a paid mutator transaction binding the contract method 0x8da3608b.
//
// Solidity: function verifyCreditScore(bytes32 l2Wallet, (int256,int256) thresholdVector, int256 thresholdSimilarity) returns(int256 similarity, bool passed)
func (_CreditVerification *CreditVerificationSession) VerifyCreditScore(l2Wallet [32]byte, thresholdVector CreditVerificationVector2D, thresholdSimilarity *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.VerifyCreditScore(&_CreditVerification.TransactOpts, l2Wallet, thresholdVector, thresholdSimilarity)
}

// VerifyCreditScore is a paid mutator transaction binding the contract method 0x8da3608b.
//
// Solidity: function verifyCreditScore(bytes32 l2Wallet, (int256,int256) thresholdVector, int256 thresholdSimilarity) returns(int256 similarity, bool passed)
func (_CreditVerification *CreditVerificationTransactorSession) VerifyCreditScore(l2Wallet [32]byte, thresholdVector CreditVerificationVector2D, thresholdSimilarity *big.Int) (*types.Transaction, error) {
	return _CreditVerification.Contract.VerifyCreditScore(&_CreditVerification.TransactOpts, l2Wallet, thresholdVector, thresholdSimilarity)
}

// CreditVerificationAccumulatedCostResetIterator is returned from FilterAccumulatedCostReset and is used to iterate over the raw logs and unpacked data for AccumulatedCostReset events raised by the CreditVerification contract.
type CreditVerificationAccumulatedCostResetIterator struct {
	Event *CreditVerificationAccumulatedCostReset // Event containing the contract specifics and raw log

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
func (it *CreditVerificationAccumulatedCostResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationAccumulatedCostReset)
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
		it.Event = new(CreditVerificationAccumulatedCostReset)
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
func (it *CreditVerificationAccumulatedCostResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationAccumulatedCostResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationAccumulatedCostReset represents a AccumulatedCostReset event raised by the CreditVerification contract.
type CreditVerificationAccumulatedCostReset struct {
	L2Wallet [32]byte
	NewCost  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccumulatedCostReset is a free log retrieval operation binding the contract event 0xa498743dbfd7c09b930a25f95c0fb18d57d463927d531532e98af5eea161a297.
//
// Solidity: event AccumulatedCostReset(bytes32 indexed l2Wallet, uint256 newCost)
func (_CreditVerification *CreditVerificationFilterer) FilterAccumulatedCostReset(opts *bind.FilterOpts, l2Wallet [][32]byte) (*CreditVerificationAccumulatedCostResetIterator, error) {

	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "AccumulatedCostReset", l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationAccumulatedCostResetIterator{contract: _CreditVerification.contract, event: "AccumulatedCostReset", logs: logs, sub: sub}, nil
}

// WatchAccumulatedCostReset is a free log subscription operation binding the contract event 0xa498743dbfd7c09b930a25f95c0fb18d57d463927d531532e98af5eea161a297.
//
// Solidity: event AccumulatedCostReset(bytes32 indexed l2Wallet, uint256 newCost)
func (_CreditVerification *CreditVerificationFilterer) WatchAccumulatedCostReset(opts *bind.WatchOpts, sink chan<- *CreditVerificationAccumulatedCostReset, l2Wallet [][32]byte) (event.Subscription, error) {

	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "AccumulatedCostReset", l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationAccumulatedCostReset)
				if err := _CreditVerification.contract.UnpackLog(event, "AccumulatedCostReset", log); err != nil {
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

// ParseAccumulatedCostReset is a log parse operation binding the contract event 0xa498743dbfd7c09b930a25f95c0fb18d57d463927d531532e98af5eea161a297.
//
// Solidity: event AccumulatedCostReset(bytes32 indexed l2Wallet, uint256 newCost)
func (_CreditVerification *CreditVerificationFilterer) ParseAccumulatedCostReset(log types.Log) (*CreditVerificationAccumulatedCostReset, error) {
	event := new(CreditVerificationAccumulatedCostReset)
	if err := _CreditVerification.contract.UnpackLog(event, "AccumulatedCostReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationCreditVerifiedIterator is returned from FilterCreditVerified and is used to iterate over the raw logs and unpacked data for CreditVerified events raised by the CreditVerification contract.
type CreditVerificationCreditVerifiedIterator struct {
	Event *CreditVerificationCreditVerified // Event containing the contract specifics and raw log

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
func (it *CreditVerificationCreditVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationCreditVerified)
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
		it.Event = new(CreditVerificationCreditVerified)
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
func (it *CreditVerificationCreditVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationCreditVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationCreditVerified represents a CreditVerified event raised by the CreditVerification contract.
type CreditVerificationCreditVerified struct {
	Verifier        common.Address
	L2Wallet        [32]byte
	Fee             *big.Int
	PassedThreshold bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreditVerified is a free log retrieval operation binding the contract event 0x116024e715d3fa7124b8639232501ea19c4a9d50eecab039a79a1aaaeb240d6d.
//
// Solidity: event CreditVerified(address indexed verifier, bytes32 indexed l2Wallet, uint256 fee, bool passedThreshold)
func (_CreditVerification *CreditVerificationFilterer) FilterCreditVerified(opts *bind.FilterOpts, verifier []common.Address, l2Wallet [][32]byte) (*CreditVerificationCreditVerifiedIterator, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "CreditVerified", verifierRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationCreditVerifiedIterator{contract: _CreditVerification.contract, event: "CreditVerified", logs: logs, sub: sub}, nil
}

// WatchCreditVerified is a free log subscription operation binding the contract event 0x116024e715d3fa7124b8639232501ea19c4a9d50eecab039a79a1aaaeb240d6d.
//
// Solidity: event CreditVerified(address indexed verifier, bytes32 indexed l2Wallet, uint256 fee, bool passedThreshold)
func (_CreditVerification *CreditVerificationFilterer) WatchCreditVerified(opts *bind.WatchOpts, sink chan<- *CreditVerificationCreditVerified, verifier []common.Address, l2Wallet [][32]byte) (event.Subscription, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "CreditVerified", verifierRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationCreditVerified)
				if err := _CreditVerification.contract.UnpackLog(event, "CreditVerified", log); err != nil {
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

// ParseCreditVerified is a log parse operation binding the contract event 0x116024e715d3fa7124b8639232501ea19c4a9d50eecab039a79a1aaaeb240d6d.
//
// Solidity: event CreditVerified(address indexed verifier, bytes32 indexed l2Wallet, uint256 fee, bool passedThreshold)
func (_CreditVerification *CreditVerificationFilterer) ParseCreditVerified(log types.Log) (*CreditVerificationCreditVerified, error) {
	event := new(CreditVerificationCreditVerified)
	if err := _CreditVerification.contract.UnpackLog(event, "CreditVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationFeeMultiplierUpdatedIterator is returned from FilterFeeMultiplierUpdated and is used to iterate over the raw logs and unpacked data for FeeMultiplierUpdated events raised by the CreditVerification contract.
type CreditVerificationFeeMultiplierUpdatedIterator struct {
	Event *CreditVerificationFeeMultiplierUpdated // Event containing the contract specifics and raw log

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
func (it *CreditVerificationFeeMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationFeeMultiplierUpdated)
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
		it.Event = new(CreditVerificationFeeMultiplierUpdated)
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
func (it *CreditVerificationFeeMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationFeeMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationFeeMultiplierUpdated represents a FeeMultiplierUpdated event raised by the CreditVerification contract.
type CreditVerificationFeeMultiplierUpdated struct {
	NewMultiplier *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeMultiplierUpdated is a free log retrieval operation binding the contract event 0xc422d1349750089c716e39be62193f9bbd759f3b17b6dcd96538dcb7fc5bf021.
//
// Solidity: event FeeMultiplierUpdated(uint256 newMultiplier)
func (_CreditVerification *CreditVerificationFilterer) FilterFeeMultiplierUpdated(opts *bind.FilterOpts) (*CreditVerificationFeeMultiplierUpdatedIterator, error) {

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "FeeMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return &CreditVerificationFeeMultiplierUpdatedIterator{contract: _CreditVerification.contract, event: "FeeMultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeMultiplierUpdated is a free log subscription operation binding the contract event 0xc422d1349750089c716e39be62193f9bbd759f3b17b6dcd96538dcb7fc5bf021.
//
// Solidity: event FeeMultiplierUpdated(uint256 newMultiplier)
func (_CreditVerification *CreditVerificationFilterer) WatchFeeMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *CreditVerificationFeeMultiplierUpdated) (event.Subscription, error) {

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "FeeMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationFeeMultiplierUpdated)
				if err := _CreditVerification.contract.UnpackLog(event, "FeeMultiplierUpdated", log); err != nil {
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

// ParseFeeMultiplierUpdated is a log parse operation binding the contract event 0xc422d1349750089c716e39be62193f9bbd759f3b17b6dcd96538dcb7fc5bf021.
//
// Solidity: event FeeMultiplierUpdated(uint256 newMultiplier)
func (_CreditVerification *CreditVerificationFilterer) ParseFeeMultiplierUpdated(log types.Log) (*CreditVerificationFeeMultiplierUpdated, error) {
	event := new(CreditVerificationFeeMultiplierUpdated)
	if err := _CreditVerification.contract.UnpackLog(event, "FeeMultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationFeedbackProvidedIterator is returned from FilterFeedbackProvided and is used to iterate over the raw logs and unpacked data for FeedbackProvided events raised by the CreditVerification contract.
type CreditVerificationFeedbackProvidedIterator struct {
	Event *CreditVerificationFeedbackProvided // Event containing the contract specifics and raw log

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
func (it *CreditVerificationFeedbackProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationFeedbackProvided)
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
		it.Event = new(CreditVerificationFeedbackProvided)
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
func (it *CreditVerificationFeedbackProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationFeedbackProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationFeedbackProvided represents a FeedbackProvided event raised by the CreditVerification contract.
type CreditVerificationFeedbackProvided struct {
	Provider  common.Address
	L2Wallet  [32]byte
	IsTrusted bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeedbackProvided is a free log retrieval operation binding the contract event 0x6273ed15405e9610be646ed2f9d0dac0d880a7ec99ce2b08145506b87ee56780.
//
// Solidity: event FeedbackProvided(address indexed provider, bytes32 indexed l2Wallet, bool isTrusted)
func (_CreditVerification *CreditVerificationFilterer) FilterFeedbackProvided(opts *bind.FilterOpts, provider []common.Address, l2Wallet [][32]byte) (*CreditVerificationFeedbackProvidedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "FeedbackProvided", providerRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationFeedbackProvidedIterator{contract: _CreditVerification.contract, event: "FeedbackProvided", logs: logs, sub: sub}, nil
}

// WatchFeedbackProvided is a free log subscription operation binding the contract event 0x6273ed15405e9610be646ed2f9d0dac0d880a7ec99ce2b08145506b87ee56780.
//
// Solidity: event FeedbackProvided(address indexed provider, bytes32 indexed l2Wallet, bool isTrusted)
func (_CreditVerification *CreditVerificationFilterer) WatchFeedbackProvided(opts *bind.WatchOpts, sink chan<- *CreditVerificationFeedbackProvided, provider []common.Address, l2Wallet [][32]byte) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "FeedbackProvided", providerRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationFeedbackProvided)
				if err := _CreditVerification.contract.UnpackLog(event, "FeedbackProvided", log); err != nil {
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

// ParseFeedbackProvided is a log parse operation binding the contract event 0x6273ed15405e9610be646ed2f9d0dac0d880a7ec99ce2b08145506b87ee56780.
//
// Solidity: event FeedbackProvided(address indexed provider, bytes32 indexed l2Wallet, bool isTrusted)
func (_CreditVerification *CreditVerificationFilterer) ParseFeedbackProvided(log types.Log) (*CreditVerificationFeedbackProvided, error) {
	event := new(CreditVerificationFeedbackProvided)
	if err := _CreditVerification.contract.UnpackLog(event, "FeedbackProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CreditVerification contract.
type CreditVerificationPausedIterator struct {
	Event *CreditVerificationPaused // Event containing the contract specifics and raw log

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
func (it *CreditVerificationPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationPaused)
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
		it.Event = new(CreditVerificationPaused)
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
func (it *CreditVerificationPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationPaused represents a Paused event raised by the CreditVerification contract.
type CreditVerificationPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditVerification *CreditVerificationFilterer) FilterPaused(opts *bind.FilterOpts) (*CreditVerificationPausedIterator, error) {

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CreditVerificationPausedIterator{contract: _CreditVerification.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CreditVerification *CreditVerificationFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CreditVerificationPaused) (event.Subscription, error) {

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationPaused)
				if err := _CreditVerification.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CreditVerification *CreditVerificationFilterer) ParsePaused(log types.Log) (*CreditVerificationPaused, error) {
	event := new(CreditVerificationPaused)
	if err := _CreditVerification.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CreditVerification contract.
type CreditVerificationRoleAdminChangedIterator struct {
	Event *CreditVerificationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CreditVerificationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationRoleAdminChanged)
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
		it.Event = new(CreditVerificationRoleAdminChanged)
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
func (it *CreditVerificationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationRoleAdminChanged represents a RoleAdminChanged event raised by the CreditVerification contract.
type CreditVerificationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CreditVerification *CreditVerificationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CreditVerificationRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationRoleAdminChangedIterator{contract: _CreditVerification.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CreditVerification *CreditVerificationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CreditVerificationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationRoleAdminChanged)
				if err := _CreditVerification.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CreditVerification *CreditVerificationFilterer) ParseRoleAdminChanged(log types.Log) (*CreditVerificationRoleAdminChanged, error) {
	event := new(CreditVerificationRoleAdminChanged)
	if err := _CreditVerification.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CreditVerification contract.
type CreditVerificationRoleGrantedIterator struct {
	Event *CreditVerificationRoleGranted // Event containing the contract specifics and raw log

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
func (it *CreditVerificationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationRoleGranted)
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
		it.Event = new(CreditVerificationRoleGranted)
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
func (it *CreditVerificationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationRoleGranted represents a RoleGranted event raised by the CreditVerification contract.
type CreditVerificationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CreditVerification *CreditVerificationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CreditVerificationRoleGrantedIterator, error) {

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

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationRoleGrantedIterator{contract: _CreditVerification.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CreditVerification *CreditVerificationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CreditVerificationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationRoleGranted)
				if err := _CreditVerification.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CreditVerification *CreditVerificationFilterer) ParseRoleGranted(log types.Log) (*CreditVerificationRoleGranted, error) {
	event := new(CreditVerificationRoleGranted)
	if err := _CreditVerification.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CreditVerification contract.
type CreditVerificationRoleRevokedIterator struct {
	Event *CreditVerificationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CreditVerificationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationRoleRevoked)
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
		it.Event = new(CreditVerificationRoleRevoked)
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
func (it *CreditVerificationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationRoleRevoked represents a RoleRevoked event raised by the CreditVerification contract.
type CreditVerificationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CreditVerification *CreditVerificationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CreditVerificationRoleRevokedIterator, error) {

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

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationRoleRevokedIterator{contract: _CreditVerification.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CreditVerification *CreditVerificationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CreditVerificationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationRoleRevoked)
				if err := _CreditVerification.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CreditVerification *CreditVerificationFilterer) ParseRoleRevoked(log types.Log) (*CreditVerificationRoleRevoked, error) {
	event := new(CreditVerificationRoleRevoked)
	if err := _CreditVerification.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CreditVerification contract.
type CreditVerificationUnpausedIterator struct {
	Event *CreditVerificationUnpaused // Event containing the contract specifics and raw log

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
func (it *CreditVerificationUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationUnpaused)
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
		it.Event = new(CreditVerificationUnpaused)
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
func (it *CreditVerificationUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationUnpaused represents a Unpaused event raised by the CreditVerification contract.
type CreditVerificationUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditVerification *CreditVerificationFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CreditVerificationUnpausedIterator, error) {

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CreditVerificationUnpausedIterator{contract: _CreditVerification.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CreditVerification *CreditVerificationFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CreditVerificationUnpaused) (event.Subscription, error) {

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationUnpaused)
				if err := _CreditVerification.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CreditVerification *CreditVerificationFilterer) ParseUnpaused(log types.Log) (*CreditVerificationUnpaused, error) {
	event := new(CreditVerificationUnpaused)
	if err := _CreditVerification.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditVerificationWalletVectorUpdatedIterator is returned from FilterWalletVectorUpdated and is used to iterate over the raw logs and unpacked data for WalletVectorUpdated events raised by the CreditVerification contract.
type CreditVerificationWalletVectorUpdatedIterator struct {
	Event *CreditVerificationWalletVectorUpdated // Event containing the contract specifics and raw log

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
func (it *CreditVerificationWalletVectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditVerificationWalletVectorUpdated)
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
		it.Event = new(CreditVerificationWalletVectorUpdated)
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
func (it *CreditVerificationWalletVectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditVerificationWalletVectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditVerificationWalletVectorUpdated represents a WalletVectorUpdated event raised by the CreditVerification contract.
type CreditVerificationWalletVectorUpdated struct {
	L2Wallet [32]byte
	VectorX  *big.Int
	VectorY  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletVectorUpdated is a free log retrieval operation binding the contract event 0x26b9f6b5b621a14c92c48918075a8b027399cf129f9fd8d3866052bca90f1c62.
//
// Solidity: event WalletVectorUpdated(bytes32 indexed l2Wallet, int256 vectorX, int256 vectorY)
func (_CreditVerification *CreditVerificationFilterer) FilterWalletVectorUpdated(opts *bind.FilterOpts, l2Wallet [][32]byte) (*CreditVerificationWalletVectorUpdatedIterator, error) {

	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.FilterLogs(opts, "WalletVectorUpdated", l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &CreditVerificationWalletVectorUpdatedIterator{contract: _CreditVerification.contract, event: "WalletVectorUpdated", logs: logs, sub: sub}, nil
}

// WatchWalletVectorUpdated is a free log subscription operation binding the contract event 0x26b9f6b5b621a14c92c48918075a8b027399cf129f9fd8d3866052bca90f1c62.
//
// Solidity: event WalletVectorUpdated(bytes32 indexed l2Wallet, int256 vectorX, int256 vectorY)
func (_CreditVerification *CreditVerificationFilterer) WatchWalletVectorUpdated(opts *bind.WatchOpts, sink chan<- *CreditVerificationWalletVectorUpdated, l2Wallet [][32]byte) (event.Subscription, error) {

	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _CreditVerification.contract.WatchLogs(opts, "WalletVectorUpdated", l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditVerificationWalletVectorUpdated)
				if err := _CreditVerification.contract.UnpackLog(event, "WalletVectorUpdated", log); err != nil {
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

// ParseWalletVectorUpdated is a log parse operation binding the contract event 0x26b9f6b5b621a14c92c48918075a8b027399cf129f9fd8d3866052bca90f1c62.
//
// Solidity: event WalletVectorUpdated(bytes32 indexed l2Wallet, int256 vectorX, int256 vectorY)
func (_CreditVerification *CreditVerificationFilterer) ParseWalletVectorUpdated(log types.Log) (*CreditVerificationWalletVectorUpdated, error) {
	event := new(CreditVerificationWalletVectorUpdated)
	if err := _CreditVerification.contract.UnpackLog(event, "WalletVectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
