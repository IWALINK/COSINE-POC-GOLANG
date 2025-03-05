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

// WalletLinkingMetaData contains all meta data concerning the WalletLinking contract.
var WalletLinkingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challengeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryTime\",\"type\":\"uint256\"}],\"name\":\"ChallengeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"WalletLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"WalletLinkingRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BSC_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CHALLENGE_EXPIRATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHEREUM_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLANA_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"adminRevokeLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"}],\"name\":\"generateChallenge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"challengeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedL1Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"}],\"name\":\"getLinkedL2Wallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1ToL2Mapping\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"l2ToL1Mapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"linkWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"revokeLinking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"verifyLinking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isLinked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506001805460ff191690556200002960003362000087565b620000557fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217753362000087565b620000817f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b9293362000087565b62000137565b62000093828262000097565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000093576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620000f33390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61167880620001476000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80636f6ad327116100de57806391d1485411610097578063c84450f911610071578063c84450f9146103bb578063d547741f146103ce578063ead5a1b4146103e1578063f5b541a61461040a57600080fd5b806391d1485414610398578063a217fddf146103ab578063a80ad651146103b357600080fd5b80636f6ad327146102c957806375b238fc1461031457806378c236e8146103295780638456cb591461034957806389151650146103515780638bac8d5d1461038557600080fd5b8063248a9ca31161014b5780633f4ba83a116101255780633f4ba83a1461029a578063405e6e31146102a257806344cf713c146102b55780635c975abb146102be57600080fd5b8063248a9ca3146102515780632f2ff15d1461027457806336568abe1461028757600080fd5b806301ffc9a714610193578063067bd07a146101bb57806307e1ac53146101d257806308b43a19146101e55780631cb3cd1c146102345780631dac56d314610249575b600080fd5b6101a66101a13660046112c1565b610431565b60405190151581526020015b60405180910390f35b6101c46101f581565b6040519081526020016101b2565b6101a66101e0366004611307565b610468565b6102176101f336600461133a565b60046020526000908152604090208054600182015460029092015490919060ff1683565b6040805193845260208401929092521515908201526060016101b2565b610247610242366004611355565b6104bb565b005b6101c4600181565b6101c461025f366004611355565b60009081526020819052604090206001015490565b61024761028236600461136e565b610566565b61024761029536600461136e565b610590565b61024761060e565b6102476102b036600461139a565b610631565b6101c4610e1081565b60015460ff166101a6565b6102fc6102d736600461141a565b600090815260036020908152604080832093835292905220546001600160a01b031690565b6040516001600160a01b0390911681526020016101b2565b6101c460008051602061162383398151915281565b6101c461033736600461133a565b60026020526000908152604090205481565b6102476109ae565b6102fc61035f36600461141a565b60036020908152600092835260408084209091529082529020546001600160a01b031681565b61024761039336600461143c565b6109ce565b6101a66103a636600461136e565b610a94565b6101c4600081565b6101c4603881565b6101c46103c936600461133a565b610abd565b6102476103dc36600461136e565b610c06565b6101c46103ef36600461133a565b6001600160a01b031660009081526002602052604090205490565b6101c47f97667070c54ef182b0f5858b034beac1b6f3089aa2d3188bb1e8929f4fa9b92981565b60006001600160e01b03198216637965db0b60e01b148061046257506301ffc9a760e01b6001600160e01b03198316145b92915050565b6001600160a01b038316600090815260026020526040812054831480156104b3575060008281526003602090815260408083208684529091529020546001600160a01b038581169116145b949350505050565b6104c3610c2b565b33600090815260026020526040902054806104f95760405162461bcd60e51b81526004016104f090611466565b60405180910390fd5b3360008181526002602090815260408083208390558583526003825280832085845282529182902080546001600160a01b031916905590518481528392917fc7aa9cf8b2f84b63b5e5f9d6b8ce4343cb07718922076f2e250848dc50dbf5f1910160405180910390a35050565b60008281526020819052604090206001015461058181610c73565b61058b8383610c7d565b505050565b6001600160a01b03811633146106005760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016104f0565b61060a8282610d01565b5050565b60008051602061162383398151915261062681610c73565b61062e610d66565b50565b610639610c2b565b8361067a5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a5908130c881dd85b1b195d607a1b60448201526064016104f0565b33600090815260026020526040902054156106d75760405162461bcd60e51b815260206004820152601960248201527f4c31206164647265737320616c7265616479206c696e6b65640000000000000060448201526064016104f0565b60008381526003602090815260408083208784529091529020546001600160a01b0316156107565760405162461bcd60e51b815260206004820152602660248201527f4c322077616c6c657420616c7265616479206c696e6b6564206f6e20746869736044820152651031b430b4b760d11b60648201526084016104f0565b33600090815260046020526040902080546107a85760405162461bcd60e51b8152602060048201526012602482015271139bc818da185b1b195b99d948199bdd5b9960721b60448201526064016104f0565b600281015460ff16156107f65760405162461bcd60e51b815260206004820152601660248201527510da185b1b195b99d948185b1c9958591e481d5cd95960521b60448201526064016104f0565b610e1081600101544261080991906114c0565b111561084b5760405162461bcd60e51b815260206004820152601160248201527010da185b1b195b99d948195e1c1a5c9959607a1b60448201526064016104f0565b80546040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c810191909152600090605c0160405160208183030381529060405280519060200120905060006108e085858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508693925050610db89050565b90506001600160a01b038116331461092e5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016104f0565b6002838101805460ff19166001179055336000818152602092835260408082208b9055898252600384528082208b835284529081902080546001600160a01b031916831790555188815289927f1cc138f50384be7770a8e4a34eecaa0d0c718bfaa1641af49db32532e4a5b910910160405180910390a350505050505050565b6000805160206116238339815191526109c681610c73565b61062e610ddc565b6000805160206116238339815191526109e681610c73565b6001600160a01b03831660009081526002602052604090205480610a1c5760405162461bcd60e51b81526004016104f090611466565b6001600160a01b03841660008181526002602090815260408083208390558683526003825280832085845282529182902080546001600160a01b031916905590518581528392917fc7aa9cf8b2f84b63b5e5f9d6b8ce4343cb07718922076f2e250848dc50dbf5f1910160405180910390a350505050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b60006001600160a01b038216610b0a5760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964204c31206164647265737360701b60448201526064016104f0565b813042610b186001436114c0565b6040516bffffffffffffffffffffffff19606095861b811660208301529390941b9092166034840152604883015240606882015260880160408051808303601f1901815282825280516020918201206060840183528084524282850181815260008686018181526001600160a01b038a1680835260049096529590209551865551600186015592516002909401805460ff1916941515949094179093559192507f5d986aaa663939d4f1fa5f4b99325eb2afb9845edf56ae12b399cac46c5961ea908390610be990610e10906114d3565b6040805192835260208301919091520160405180910390a2919050565b600082815260208190526040902060010154610c2181610c73565b61058b8383610d01565b60015460ff1615610c715760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104f0565b565b61062e8133610e17565b610c878282610a94565b61060a576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055610cbd3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b610d0b8282610a94565b1561060a576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b610d6e610e70565b6001805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000806000610dc78585610eb9565b91509150610dd481610efe565b509392505050565b610de4610c2b565b6001805460ff1916811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610d9b565b610e218282610a94565b61060a57610e2e81611048565b610e3983602061105a565b604051602001610e4a92919061150a565b60408051601f198184030181529082905262461bcd60e51b82526104f09160040161157f565b60015460ff16610c715760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016104f0565b6000808251604103610eef5760208301516040840151606085015160001a610ee3878285856111fd565b94509450505050610ef7565b506000905060025b9250929050565b6000816004811115610f1257610f126115b2565b03610f1a5750565b6001816004811115610f2e57610f2e6115b2565b03610f7b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104f0565b6002816004811115610f8f57610f8f6115b2565b03610fdc5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104f0565b6003816004811115610ff057610ff06115b2565b0361062e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016104f0565b60606104626001600160a01b03831660145b606060006110698360026115c8565b6110749060026114d3565b67ffffffffffffffff81111561108c5761108c6115df565b6040519080825280601f01601f1916602001820160405280156110b6576020820181803683370190505b509050600360fc1b816000815181106110d1576110d16115f5565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611100576111006115f5565b60200101906001600160f81b031916908160001a90535060006111248460026115c8565b61112f9060016114d3565b90505b60018111156111a7576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611163576111636115f5565b1a60f81b828281518110611179576111796115f5565b60200101906001600160f81b031916908160001a90535060049490941c936111a08161160b565b9050611132565b5083156111f65760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016104f0565b9392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561123457506000905060036112b8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611288573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166112b1576000600192509250506112b8565b9150600090505b94509492505050565b6000602082840312156112d357600080fd5b81356001600160e01b0319811681146111f657600080fd5b80356001600160a01b038116811461130257600080fd5b919050565b60008060006060848603121561131c57600080fd5b611325846112eb565b95602085013595506040909401359392505050565b60006020828403121561134c57600080fd5b6111f6826112eb565b60006020828403121561136757600080fd5b5035919050565b6000806040838503121561138157600080fd5b82359150611391602084016112eb565b90509250929050565b600080600080606085870312156113b057600080fd5b8435935060208501359250604085013567ffffffffffffffff808211156113d657600080fd5b818701915087601f8301126113ea57600080fd5b8135818111156113f957600080fd5b88602082850101111561140b57600080fd5b95989497505060200194505050565b6000806040838503121561142d57600080fd5b50508035926020909101359150565b6000806040838503121561144f57600080fd5b611458836112eb565b946020939093013593505050565b60208082526024908201527f4e6f206c696e6b696e6720666f756e6420666f722074686973204c31206164646040820152637265737360e01b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610462576104626114aa565b80820180821115610462576104626114aa565b60005b838110156115015781810151838201526020016114e9565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516115428160178501602088016114e6565b7001034b99036b4b9b9b4b733903937b6329607d1b60179184019182015283516115738160288401602088016114e6565b01602801949350505050565b602081526000825180602084015261159e8160408501602087016114e6565b601f01601f19169190910160400192915050565b634e487b7160e01b600052602160045260246000fd5b8082028115828204841417610462576104626114aa565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60008161161a5761161a6114aa565b50600019019056fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a26469706673582212208b3a86841de3c2fae9c6daae82d63f1d1a7e9fc3c21d0fc8574d1a520a8d65cb64736f6c63430008130033",
}

// WalletLinkingABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletLinkingMetaData.ABI instead.
var WalletLinkingABI = WalletLinkingMetaData.ABI

// WalletLinkingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletLinkingMetaData.Bin instead.
var WalletLinkingBin = WalletLinkingMetaData.Bin

// DeployWalletLinking deploys a new Ethereum contract, binding an instance of WalletLinking to it.
func DeployWalletLinking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletLinking, error) {
	parsed, err := WalletLinkingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletLinkingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletLinking{WalletLinkingCaller: WalletLinkingCaller{contract: contract}, WalletLinkingTransactor: WalletLinkingTransactor{contract: contract}, WalletLinkingFilterer: WalletLinkingFilterer{contract: contract}}, nil
}

// WalletLinking is an auto generated Go binding around an Ethereum contract.
type WalletLinking struct {
	WalletLinkingCaller     // Read-only binding to the contract
	WalletLinkingTransactor // Write-only binding to the contract
	WalletLinkingFilterer   // Log filterer for contract events
}

// WalletLinkingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletLinkingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLinkingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletLinkingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLinkingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletLinkingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLinkingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletLinkingSession struct {
	Contract     *WalletLinking    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletLinkingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletLinkingCallerSession struct {
	Contract *WalletLinkingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletLinkingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletLinkingTransactorSession struct {
	Contract     *WalletLinkingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletLinkingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletLinkingRaw struct {
	Contract *WalletLinking // Generic contract binding to access the raw methods on
}

// WalletLinkingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletLinkingCallerRaw struct {
	Contract *WalletLinkingCaller // Generic read-only contract binding to access the raw methods on
}

// WalletLinkingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletLinkingTransactorRaw struct {
	Contract *WalletLinkingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletLinking creates a new instance of WalletLinking, bound to a specific deployed contract.
func NewWalletLinking(address common.Address, backend bind.ContractBackend) (*WalletLinking, error) {
	contract, err := bindWalletLinking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletLinking{WalletLinkingCaller: WalletLinkingCaller{contract: contract}, WalletLinkingTransactor: WalletLinkingTransactor{contract: contract}, WalletLinkingFilterer: WalletLinkingFilterer{contract: contract}}, nil
}

// NewWalletLinkingCaller creates a new read-only instance of WalletLinking, bound to a specific deployed contract.
func NewWalletLinkingCaller(address common.Address, caller bind.ContractCaller) (*WalletLinkingCaller, error) {
	contract, err := bindWalletLinking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingCaller{contract: contract}, nil
}

// NewWalletLinkingTransactor creates a new write-only instance of WalletLinking, bound to a specific deployed contract.
func NewWalletLinkingTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletLinkingTransactor, error) {
	contract, err := bindWalletLinking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingTransactor{contract: contract}, nil
}

// NewWalletLinkingFilterer creates a new log filterer instance of WalletLinking, bound to a specific deployed contract.
func NewWalletLinkingFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletLinkingFilterer, error) {
	contract, err := bindWalletLinking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingFilterer{contract: contract}, nil
}

// bindWalletLinking binds a generic wrapper to an already deployed contract.
func bindWalletLinking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletLinkingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletLinking *WalletLinkingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletLinking.Contract.WalletLinkingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletLinking *WalletLinkingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLinking.Contract.WalletLinkingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletLinking *WalletLinkingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletLinking.Contract.WalletLinkingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletLinking *WalletLinkingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletLinking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletLinking *WalletLinkingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLinking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletLinking *WalletLinkingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletLinking.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingSession) ADMINROLE() ([32]byte, error) {
	return _WalletLinking.Contract.ADMINROLE(&_WalletLinking.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCallerSession) ADMINROLE() ([32]byte, error) {
	return _WalletLinking.Contract.ADMINROLE(&_WalletLinking.CallOpts)
}

// BSCCHAINID is a free data retrieval call binding the contract method 0xa80ad651.
//
// Solidity: function BSC_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCaller) BSCCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "BSC_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BSCCHAINID is a free data retrieval call binding the contract method 0xa80ad651.
//
// Solidity: function BSC_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingSession) BSCCHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.BSCCHAINID(&_WalletLinking.CallOpts)
}

// BSCCHAINID is a free data retrieval call binding the contract method 0xa80ad651.
//
// Solidity: function BSC_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCallerSession) BSCCHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.BSCCHAINID(&_WalletLinking.CallOpts)
}

// CHALLENGEEXPIRATION is a free data retrieval call binding the contract method 0x44cf713c.
//
// Solidity: function CHALLENGE_EXPIRATION() view returns(uint256)
func (_WalletLinking *WalletLinkingCaller) CHALLENGEEXPIRATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "CHALLENGE_EXPIRATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHALLENGEEXPIRATION is a free data retrieval call binding the contract method 0x44cf713c.
//
// Solidity: function CHALLENGE_EXPIRATION() view returns(uint256)
func (_WalletLinking *WalletLinkingSession) CHALLENGEEXPIRATION() (*big.Int, error) {
	return _WalletLinking.Contract.CHALLENGEEXPIRATION(&_WalletLinking.CallOpts)
}

// CHALLENGEEXPIRATION is a free data retrieval call binding the contract method 0x44cf713c.
//
// Solidity: function CHALLENGE_EXPIRATION() view returns(uint256)
func (_WalletLinking *WalletLinkingCallerSession) CHALLENGEEXPIRATION() (*big.Int, error) {
	return _WalletLinking.Contract.CHALLENGEEXPIRATION(&_WalletLinking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WalletLinking.Contract.DEFAULTADMINROLE(&_WalletLinking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WalletLinking.Contract.DEFAULTADMINROLE(&_WalletLinking.CallOpts)
}

// ETHEREUMCHAINID is a free data retrieval call binding the contract method 0x1dac56d3.
//
// Solidity: function ETHEREUM_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCaller) ETHEREUMCHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "ETHEREUM_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHEREUMCHAINID is a free data retrieval call binding the contract method 0x1dac56d3.
//
// Solidity: function ETHEREUM_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingSession) ETHEREUMCHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.ETHEREUMCHAINID(&_WalletLinking.CallOpts)
}

// ETHEREUMCHAINID is a free data retrieval call binding the contract method 0x1dac56d3.
//
// Solidity: function ETHEREUM_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCallerSession) ETHEREUMCHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.ETHEREUMCHAINID(&_WalletLinking.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingSession) OPERATORROLE() ([32]byte, error) {
	return _WalletLinking.Contract.OPERATORROLE(&_WalletLinking.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WalletLinking *WalletLinkingCallerSession) OPERATORROLE() ([32]byte, error) {
	return _WalletLinking.Contract.OPERATORROLE(&_WalletLinking.CallOpts)
}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCaller) SOLANACHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "SOLANA_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingSession) SOLANACHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.SOLANACHAINID(&_WalletLinking.CallOpts)
}

// SOLANACHAINID is a free data retrieval call binding the contract method 0x067bd07a.
//
// Solidity: function SOLANA_CHAIN_ID() view returns(uint256)
func (_WalletLinking *WalletLinkingCallerSession) SOLANACHAINID() (*big.Int, error) {
	return _WalletLinking.Contract.SOLANACHAINID(&_WalletLinking.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bytes32 challengeHash, uint256 timestamp, bool isUsed)
func (_WalletLinking *WalletLinkingCaller) Challenges(opts *bind.CallOpts, arg0 common.Address) (struct {
	ChallengeHash [32]byte
	Timestamp     *big.Int
	IsUsed        bool
}, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		ChallengeHash [32]byte
		Timestamp     *big.Int
		IsUsed        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChallengeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsUsed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bytes32 challengeHash, uint256 timestamp, bool isUsed)
func (_WalletLinking *WalletLinkingSession) Challenges(arg0 common.Address) (struct {
	ChallengeHash [32]byte
	Timestamp     *big.Int
	IsUsed        bool
}, error) {
	return _WalletLinking.Contract.Challenges(&_WalletLinking.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bytes32 challengeHash, uint256 timestamp, bool isUsed)
func (_WalletLinking *WalletLinkingCallerSession) Challenges(arg0 common.Address) (struct {
	ChallengeHash [32]byte
	Timestamp     *big.Int
	IsUsed        bool
}, error) {
	return _WalletLinking.Contract.Challenges(&_WalletLinking.CallOpts, arg0)
}

// GetLinkedL1Address is a free data retrieval call binding the contract method 0x6f6ad327.
//
// Solidity: function getLinkedL1Address(bytes32 l2Wallet, uint256 chainId) view returns(address l1Address)
func (_WalletLinking *WalletLinkingCaller) GetLinkedL1Address(opts *bind.CallOpts, l2Wallet [32]byte, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "getLinkedL1Address", l2Wallet, chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedL1Address is a free data retrieval call binding the contract method 0x6f6ad327.
//
// Solidity: function getLinkedL1Address(bytes32 l2Wallet, uint256 chainId) view returns(address l1Address)
func (_WalletLinking *WalletLinkingSession) GetLinkedL1Address(l2Wallet [32]byte, chainId *big.Int) (common.Address, error) {
	return _WalletLinking.Contract.GetLinkedL1Address(&_WalletLinking.CallOpts, l2Wallet, chainId)
}

// GetLinkedL1Address is a free data retrieval call binding the contract method 0x6f6ad327.
//
// Solidity: function getLinkedL1Address(bytes32 l2Wallet, uint256 chainId) view returns(address l1Address)
func (_WalletLinking *WalletLinkingCallerSession) GetLinkedL1Address(l2Wallet [32]byte, chainId *big.Int) (common.Address, error) {
	return _WalletLinking.Contract.GetLinkedL1Address(&_WalletLinking.CallOpts, l2Wallet, chainId)
}

// GetLinkedL2Wallet is a free data retrieval call binding the contract method 0xead5a1b4.
//
// Solidity: function getLinkedL2Wallet(address l1Address) view returns(bytes32 l2Wallet)
func (_WalletLinking *WalletLinkingCaller) GetLinkedL2Wallet(opts *bind.CallOpts, l1Address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "getLinkedL2Wallet", l1Address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedL2Wallet is a free data retrieval call binding the contract method 0xead5a1b4.
//
// Solidity: function getLinkedL2Wallet(address l1Address) view returns(bytes32 l2Wallet)
func (_WalletLinking *WalletLinkingSession) GetLinkedL2Wallet(l1Address common.Address) ([32]byte, error) {
	return _WalletLinking.Contract.GetLinkedL2Wallet(&_WalletLinking.CallOpts, l1Address)
}

// GetLinkedL2Wallet is a free data retrieval call binding the contract method 0xead5a1b4.
//
// Solidity: function getLinkedL2Wallet(address l1Address) view returns(bytes32 l2Wallet)
func (_WalletLinking *WalletLinkingCallerSession) GetLinkedL2Wallet(l1Address common.Address) ([32]byte, error) {
	return _WalletLinking.Contract.GetLinkedL2Wallet(&_WalletLinking.CallOpts, l1Address)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletLinking *WalletLinkingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletLinking *WalletLinkingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WalletLinking.Contract.GetRoleAdmin(&_WalletLinking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WalletLinking *WalletLinkingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WalletLinking.Contract.GetRoleAdmin(&_WalletLinking.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletLinking *WalletLinkingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletLinking *WalletLinkingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WalletLinking.Contract.HasRole(&_WalletLinking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WalletLinking *WalletLinkingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WalletLinking.Contract.HasRole(&_WalletLinking.CallOpts, role, account)
}

// L1ToL2Mapping is a free data retrieval call binding the contract method 0x78c236e8.
//
// Solidity: function l1ToL2Mapping(address ) view returns(bytes32)
func (_WalletLinking *WalletLinkingCaller) L1ToL2Mapping(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "l1ToL2Mapping", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1ToL2Mapping is a free data retrieval call binding the contract method 0x78c236e8.
//
// Solidity: function l1ToL2Mapping(address ) view returns(bytes32)
func (_WalletLinking *WalletLinkingSession) L1ToL2Mapping(arg0 common.Address) ([32]byte, error) {
	return _WalletLinking.Contract.L1ToL2Mapping(&_WalletLinking.CallOpts, arg0)
}

// L1ToL2Mapping is a free data retrieval call binding the contract method 0x78c236e8.
//
// Solidity: function l1ToL2Mapping(address ) view returns(bytes32)
func (_WalletLinking *WalletLinkingCallerSession) L1ToL2Mapping(arg0 common.Address) ([32]byte, error) {
	return _WalletLinking.Contract.L1ToL2Mapping(&_WalletLinking.CallOpts, arg0)
}

// L2ToL1Mapping is a free data retrieval call binding the contract method 0x89151650.
//
// Solidity: function l2ToL1Mapping(uint256 , bytes32 ) view returns(address)
func (_WalletLinking *WalletLinkingCaller) L2ToL1Mapping(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "l2ToL1Mapping", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Mapping is a free data retrieval call binding the contract method 0x89151650.
//
// Solidity: function l2ToL1Mapping(uint256 , bytes32 ) view returns(address)
func (_WalletLinking *WalletLinkingSession) L2ToL1Mapping(arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	return _WalletLinking.Contract.L2ToL1Mapping(&_WalletLinking.CallOpts, arg0, arg1)
}

// L2ToL1Mapping is a free data retrieval call binding the contract method 0x89151650.
//
// Solidity: function l2ToL1Mapping(uint256 , bytes32 ) view returns(address)
func (_WalletLinking *WalletLinkingCallerSession) L2ToL1Mapping(arg0 *big.Int, arg1 [32]byte) (common.Address, error) {
	return _WalletLinking.Contract.L2ToL1Mapping(&_WalletLinking.CallOpts, arg0, arg1)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WalletLinking *WalletLinkingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WalletLinking *WalletLinkingSession) Paused() (bool, error) {
	return _WalletLinking.Contract.Paused(&_WalletLinking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WalletLinking *WalletLinkingCallerSession) Paused() (bool, error) {
	return _WalletLinking.Contract.Paused(&_WalletLinking.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletLinking *WalletLinkingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletLinking *WalletLinkingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletLinking.Contract.SupportsInterface(&_WalletLinking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletLinking *WalletLinkingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletLinking.Contract.SupportsInterface(&_WalletLinking.CallOpts, interfaceId)
}

// VerifyLinking is a free data retrieval call binding the contract method 0x07e1ac53.
//
// Solidity: function verifyLinking(address l1Address, bytes32 l2Wallet, uint256 chainId) view returns(bool isLinked)
func (_WalletLinking *WalletLinkingCaller) VerifyLinking(opts *bind.CallOpts, l1Address common.Address, l2Wallet [32]byte, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _WalletLinking.contract.Call(opts, &out, "verifyLinking", l1Address, l2Wallet, chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyLinking is a free data retrieval call binding the contract method 0x07e1ac53.
//
// Solidity: function verifyLinking(address l1Address, bytes32 l2Wallet, uint256 chainId) view returns(bool isLinked)
func (_WalletLinking *WalletLinkingSession) VerifyLinking(l1Address common.Address, l2Wallet [32]byte, chainId *big.Int) (bool, error) {
	return _WalletLinking.Contract.VerifyLinking(&_WalletLinking.CallOpts, l1Address, l2Wallet, chainId)
}

// VerifyLinking is a free data retrieval call binding the contract method 0x07e1ac53.
//
// Solidity: function verifyLinking(address l1Address, bytes32 l2Wallet, uint256 chainId) view returns(bool isLinked)
func (_WalletLinking *WalletLinkingCallerSession) VerifyLinking(l1Address common.Address, l2Wallet [32]byte, chainId *big.Int) (bool, error) {
	return _WalletLinking.Contract.VerifyLinking(&_WalletLinking.CallOpts, l1Address, l2Wallet, chainId)
}

// AdminRevokeLink is a paid mutator transaction binding the contract method 0x8bac8d5d.
//
// Solidity: function adminRevokeLink(address l1Address, uint256 chainId) returns()
func (_WalletLinking *WalletLinkingTransactor) AdminRevokeLink(opts *bind.TransactOpts, l1Address common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "adminRevokeLink", l1Address, chainId)
}

// AdminRevokeLink is a paid mutator transaction binding the contract method 0x8bac8d5d.
//
// Solidity: function adminRevokeLink(address l1Address, uint256 chainId) returns()
func (_WalletLinking *WalletLinkingSession) AdminRevokeLink(l1Address common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.Contract.AdminRevokeLink(&_WalletLinking.TransactOpts, l1Address, chainId)
}

// AdminRevokeLink is a paid mutator transaction binding the contract method 0x8bac8d5d.
//
// Solidity: function adminRevokeLink(address l1Address, uint256 chainId) returns()
func (_WalletLinking *WalletLinkingTransactorSession) AdminRevokeLink(l1Address common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.Contract.AdminRevokeLink(&_WalletLinking.TransactOpts, l1Address, chainId)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xc84450f9.
//
// Solidity: function generateChallenge(address l1Address) returns(bytes32 challengeHash)
func (_WalletLinking *WalletLinkingTransactor) GenerateChallenge(opts *bind.TransactOpts, l1Address common.Address) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "generateChallenge", l1Address)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xc84450f9.
//
// Solidity: function generateChallenge(address l1Address) returns(bytes32 challengeHash)
func (_WalletLinking *WalletLinkingSession) GenerateChallenge(l1Address common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.GenerateChallenge(&_WalletLinking.TransactOpts, l1Address)
}

// GenerateChallenge is a paid mutator transaction binding the contract method 0xc84450f9.
//
// Solidity: function generateChallenge(address l1Address) returns(bytes32 challengeHash)
func (_WalletLinking *WalletLinkingTransactorSession) GenerateChallenge(l1Address common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.GenerateChallenge(&_WalletLinking.TransactOpts, l1Address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.GrantRole(&_WalletLinking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.GrantRole(&_WalletLinking.TransactOpts, role, account)
}

// LinkWallet is a paid mutator transaction binding the contract method 0x405e6e31.
//
// Solidity: function linkWallet(bytes32 l2Wallet, uint256 chainId, bytes signature) returns()
func (_WalletLinking *WalletLinkingTransactor) LinkWallet(opts *bind.TransactOpts, l2Wallet [32]byte, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "linkWallet", l2Wallet, chainId, signature)
}

// LinkWallet is a paid mutator transaction binding the contract method 0x405e6e31.
//
// Solidity: function linkWallet(bytes32 l2Wallet, uint256 chainId, bytes signature) returns()
func (_WalletLinking *WalletLinkingSession) LinkWallet(l2Wallet [32]byte, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _WalletLinking.Contract.LinkWallet(&_WalletLinking.TransactOpts, l2Wallet, chainId, signature)
}

// LinkWallet is a paid mutator transaction binding the contract method 0x405e6e31.
//
// Solidity: function linkWallet(bytes32 l2Wallet, uint256 chainId, bytes signature) returns()
func (_WalletLinking *WalletLinkingTransactorSession) LinkWallet(l2Wallet [32]byte, chainId *big.Int, signature []byte) (*types.Transaction, error) {
	return _WalletLinking.Contract.LinkWallet(&_WalletLinking.TransactOpts, l2Wallet, chainId, signature)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WalletLinking *WalletLinkingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WalletLinking *WalletLinkingSession) Pause() (*types.Transaction, error) {
	return _WalletLinking.Contract.Pause(&_WalletLinking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WalletLinking *WalletLinkingTransactorSession) Pause() (*types.Transaction, error) {
	return _WalletLinking.Contract.Pause(&_WalletLinking.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.RenounceRole(&_WalletLinking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.RenounceRole(&_WalletLinking.TransactOpts, role, account)
}

// RevokeLinking is a paid mutator transaction binding the contract method 0x1cb3cd1c.
//
// Solidity: function revokeLinking(uint256 chainId) returns()
func (_WalletLinking *WalletLinkingTransactor) RevokeLinking(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "revokeLinking", chainId)
}

// RevokeLinking is a paid mutator transaction binding the contract method 0x1cb3cd1c.
//
// Solidity: function revokeLinking(uint256 chainId) returns()
func (_WalletLinking *WalletLinkingSession) RevokeLinking(chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.Contract.RevokeLinking(&_WalletLinking.TransactOpts, chainId)
}

// RevokeLinking is a paid mutator transaction binding the contract method 0x1cb3cd1c.
//
// Solidity: function revokeLinking(uint256 chainId) returns()
func (_WalletLinking *WalletLinkingTransactorSession) RevokeLinking(chainId *big.Int) (*types.Transaction, error) {
	return _WalletLinking.Contract.RevokeLinking(&_WalletLinking.TransactOpts, chainId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.RevokeRole(&_WalletLinking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WalletLinking *WalletLinkingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WalletLinking.Contract.RevokeRole(&_WalletLinking.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WalletLinking *WalletLinkingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLinking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WalletLinking *WalletLinkingSession) Unpause() (*types.Transaction, error) {
	return _WalletLinking.Contract.Unpause(&_WalletLinking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WalletLinking *WalletLinkingTransactorSession) Unpause() (*types.Transaction, error) {
	return _WalletLinking.Contract.Unpause(&_WalletLinking.TransactOpts)
}

// WalletLinkingChallengeCreatedIterator is returned from FilterChallengeCreated and is used to iterate over the raw logs and unpacked data for ChallengeCreated events raised by the WalletLinking contract.
type WalletLinkingChallengeCreatedIterator struct {
	Event *WalletLinkingChallengeCreated // Event containing the contract specifics and raw log

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
func (it *WalletLinkingChallengeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingChallengeCreated)
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
		it.Event = new(WalletLinkingChallengeCreated)
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
func (it *WalletLinkingChallengeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingChallengeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingChallengeCreated represents a ChallengeCreated event raised by the WalletLinking contract.
type WalletLinkingChallengeCreated struct {
	L1Address     common.Address
	ChallengeHash [32]byte
	ExpiryTime    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChallengeCreated is a free log retrieval operation binding the contract event 0x5d986aaa663939d4f1fa5f4b99325eb2afb9845edf56ae12b399cac46c5961ea.
//
// Solidity: event ChallengeCreated(address indexed l1Address, bytes32 challengeHash, uint256 expiryTime)
func (_WalletLinking *WalletLinkingFilterer) FilterChallengeCreated(opts *bind.FilterOpts, l1Address []common.Address) (*WalletLinkingChallengeCreatedIterator, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "ChallengeCreated", l1AddressRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingChallengeCreatedIterator{contract: _WalletLinking.contract, event: "ChallengeCreated", logs: logs, sub: sub}, nil
}

// WatchChallengeCreated is a free log subscription operation binding the contract event 0x5d986aaa663939d4f1fa5f4b99325eb2afb9845edf56ae12b399cac46c5961ea.
//
// Solidity: event ChallengeCreated(address indexed l1Address, bytes32 challengeHash, uint256 expiryTime)
func (_WalletLinking *WalletLinkingFilterer) WatchChallengeCreated(opts *bind.WatchOpts, sink chan<- *WalletLinkingChallengeCreated, l1Address []common.Address) (event.Subscription, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "ChallengeCreated", l1AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingChallengeCreated)
				if err := _WalletLinking.contract.UnpackLog(event, "ChallengeCreated", log); err != nil {
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

// ParseChallengeCreated is a log parse operation binding the contract event 0x5d986aaa663939d4f1fa5f4b99325eb2afb9845edf56ae12b399cac46c5961ea.
//
// Solidity: event ChallengeCreated(address indexed l1Address, bytes32 challengeHash, uint256 expiryTime)
func (_WalletLinking *WalletLinkingFilterer) ParseChallengeCreated(log types.Log) (*WalletLinkingChallengeCreated, error) {
	event := new(WalletLinkingChallengeCreated)
	if err := _WalletLinking.contract.UnpackLog(event, "ChallengeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WalletLinking contract.
type WalletLinkingPausedIterator struct {
	Event *WalletLinkingPaused // Event containing the contract specifics and raw log

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
func (it *WalletLinkingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingPaused)
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
		it.Event = new(WalletLinkingPaused)
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
func (it *WalletLinkingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingPaused represents a Paused event raised by the WalletLinking contract.
type WalletLinkingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WalletLinking *WalletLinkingFilterer) FilterPaused(opts *bind.FilterOpts) (*WalletLinkingPausedIterator, error) {

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WalletLinkingPausedIterator{contract: _WalletLinking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WalletLinking *WalletLinkingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WalletLinkingPaused) (event.Subscription, error) {

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingPaused)
				if err := _WalletLinking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WalletLinking *WalletLinkingFilterer) ParsePaused(log types.Log) (*WalletLinkingPaused, error) {
	event := new(WalletLinkingPaused)
	if err := _WalletLinking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the WalletLinking contract.
type WalletLinkingRoleAdminChangedIterator struct {
	Event *WalletLinkingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *WalletLinkingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingRoleAdminChanged)
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
		it.Event = new(WalletLinkingRoleAdminChanged)
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
func (it *WalletLinkingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingRoleAdminChanged represents a RoleAdminChanged event raised by the WalletLinking contract.
type WalletLinkingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WalletLinking *WalletLinkingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*WalletLinkingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingRoleAdminChangedIterator{contract: _WalletLinking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WalletLinking *WalletLinkingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *WalletLinkingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingRoleAdminChanged)
				if err := _WalletLinking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_WalletLinking *WalletLinkingFilterer) ParseRoleAdminChanged(log types.Log) (*WalletLinkingRoleAdminChanged, error) {
	event := new(WalletLinkingRoleAdminChanged)
	if err := _WalletLinking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the WalletLinking contract.
type WalletLinkingRoleGrantedIterator struct {
	Event *WalletLinkingRoleGranted // Event containing the contract specifics and raw log

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
func (it *WalletLinkingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingRoleGranted)
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
		it.Event = new(WalletLinkingRoleGranted)
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
func (it *WalletLinkingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingRoleGranted represents a RoleGranted event raised by the WalletLinking contract.
type WalletLinkingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletLinking *WalletLinkingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WalletLinkingRoleGrantedIterator, error) {

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

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingRoleGrantedIterator{contract: _WalletLinking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletLinking *WalletLinkingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *WalletLinkingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingRoleGranted)
				if err := _WalletLinking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_WalletLinking *WalletLinkingFilterer) ParseRoleGranted(log types.Log) (*WalletLinkingRoleGranted, error) {
	event := new(WalletLinkingRoleGranted)
	if err := _WalletLinking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the WalletLinking contract.
type WalletLinkingRoleRevokedIterator struct {
	Event *WalletLinkingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *WalletLinkingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingRoleRevoked)
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
		it.Event = new(WalletLinkingRoleRevoked)
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
func (it *WalletLinkingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingRoleRevoked represents a RoleRevoked event raised by the WalletLinking contract.
type WalletLinkingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletLinking *WalletLinkingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WalletLinkingRoleRevokedIterator, error) {

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

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingRoleRevokedIterator{contract: _WalletLinking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WalletLinking *WalletLinkingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *WalletLinkingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingRoleRevoked)
				if err := _WalletLinking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_WalletLinking *WalletLinkingFilterer) ParseRoleRevoked(log types.Log) (*WalletLinkingRoleRevoked, error) {
	event := new(WalletLinkingRoleRevoked)
	if err := _WalletLinking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WalletLinking contract.
type WalletLinkingUnpausedIterator struct {
	Event *WalletLinkingUnpaused // Event containing the contract specifics and raw log

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
func (it *WalletLinkingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingUnpaused)
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
		it.Event = new(WalletLinkingUnpaused)
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
func (it *WalletLinkingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingUnpaused represents a Unpaused event raised by the WalletLinking contract.
type WalletLinkingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WalletLinking *WalletLinkingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WalletLinkingUnpausedIterator, error) {

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WalletLinkingUnpausedIterator{contract: _WalletLinking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WalletLinking *WalletLinkingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WalletLinkingUnpaused) (event.Subscription, error) {

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingUnpaused)
				if err := _WalletLinking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WalletLinking *WalletLinkingFilterer) ParseUnpaused(log types.Log) (*WalletLinkingUnpaused, error) {
	event := new(WalletLinkingUnpaused)
	if err := _WalletLinking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingWalletLinkedIterator is returned from FilterWalletLinked and is used to iterate over the raw logs and unpacked data for WalletLinked events raised by the WalletLinking contract.
type WalletLinkingWalletLinkedIterator struct {
	Event *WalletLinkingWalletLinked // Event containing the contract specifics and raw log

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
func (it *WalletLinkingWalletLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingWalletLinked)
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
		it.Event = new(WalletLinkingWalletLinked)
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
func (it *WalletLinkingWalletLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingWalletLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingWalletLinked represents a WalletLinked event raised by the WalletLinking contract.
type WalletLinkingWalletLinked struct {
	L1Address common.Address
	L2Wallet  [32]byte
	ChainId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletLinked is a free log retrieval operation binding the contract event 0x1cc138f50384be7770a8e4a34eecaa0d0c718bfaa1641af49db32532e4a5b910.
//
// Solidity: event WalletLinked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) FilterWalletLinked(opts *bind.FilterOpts, l1Address []common.Address, l2Wallet [][32]byte) (*WalletLinkingWalletLinkedIterator, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "WalletLinked", l1AddressRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingWalletLinkedIterator{contract: _WalletLinking.contract, event: "WalletLinked", logs: logs, sub: sub}, nil
}

// WatchWalletLinked is a free log subscription operation binding the contract event 0x1cc138f50384be7770a8e4a34eecaa0d0c718bfaa1641af49db32532e4a5b910.
//
// Solidity: event WalletLinked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) WatchWalletLinked(opts *bind.WatchOpts, sink chan<- *WalletLinkingWalletLinked, l1Address []common.Address, l2Wallet [][32]byte) (event.Subscription, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "WalletLinked", l1AddressRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingWalletLinked)
				if err := _WalletLinking.contract.UnpackLog(event, "WalletLinked", log); err != nil {
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

// ParseWalletLinked is a log parse operation binding the contract event 0x1cc138f50384be7770a8e4a34eecaa0d0c718bfaa1641af49db32532e4a5b910.
//
// Solidity: event WalletLinked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) ParseWalletLinked(log types.Log) (*WalletLinkingWalletLinked, error) {
	event := new(WalletLinkingWalletLinked)
	if err := _WalletLinking.contract.UnpackLog(event, "WalletLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletLinkingWalletLinkingRevokedIterator is returned from FilterWalletLinkingRevoked and is used to iterate over the raw logs and unpacked data for WalletLinkingRevoked events raised by the WalletLinking contract.
type WalletLinkingWalletLinkingRevokedIterator struct {
	Event *WalletLinkingWalletLinkingRevoked // Event containing the contract specifics and raw log

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
func (it *WalletLinkingWalletLinkingRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLinkingWalletLinkingRevoked)
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
		it.Event = new(WalletLinkingWalletLinkingRevoked)
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
func (it *WalletLinkingWalletLinkingRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLinkingWalletLinkingRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLinkingWalletLinkingRevoked represents a WalletLinkingRevoked event raised by the WalletLinking contract.
type WalletLinkingWalletLinkingRevoked struct {
	L1Address common.Address
	L2Wallet  [32]byte
	ChainId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletLinkingRevoked is a free log retrieval operation binding the contract event 0xc7aa9cf8b2f84b63b5e5f9d6b8ce4343cb07718922076f2e250848dc50dbf5f1.
//
// Solidity: event WalletLinkingRevoked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) FilterWalletLinkingRevoked(opts *bind.FilterOpts, l1Address []common.Address, l2Wallet [][32]byte) (*WalletLinkingWalletLinkingRevokedIterator, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _WalletLinking.contract.FilterLogs(opts, "WalletLinkingRevoked", l1AddressRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return &WalletLinkingWalletLinkingRevokedIterator{contract: _WalletLinking.contract, event: "WalletLinkingRevoked", logs: logs, sub: sub}, nil
}

// WatchWalletLinkingRevoked is a free log subscription operation binding the contract event 0xc7aa9cf8b2f84b63b5e5f9d6b8ce4343cb07718922076f2e250848dc50dbf5f1.
//
// Solidity: event WalletLinkingRevoked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) WatchWalletLinkingRevoked(opts *bind.WatchOpts, sink chan<- *WalletLinkingWalletLinkingRevoked, l1Address []common.Address, l2Wallet [][32]byte) (event.Subscription, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2WalletRule []interface{}
	for _, l2WalletItem := range l2Wallet {
		l2WalletRule = append(l2WalletRule, l2WalletItem)
	}

	logs, sub, err := _WalletLinking.contract.WatchLogs(opts, "WalletLinkingRevoked", l1AddressRule, l2WalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLinkingWalletLinkingRevoked)
				if err := _WalletLinking.contract.UnpackLog(event, "WalletLinkingRevoked", log); err != nil {
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

// ParseWalletLinkingRevoked is a log parse operation binding the contract event 0xc7aa9cf8b2f84b63b5e5f9d6b8ce4343cb07718922076f2e250848dc50dbf5f1.
//
// Solidity: event WalletLinkingRevoked(address indexed l1Address, bytes32 indexed l2Wallet, uint256 chainId)
func (_WalletLinking *WalletLinkingFilterer) ParseWalletLinkingRevoked(log types.Log) (*WalletLinkingWalletLinkingRevoked, error) {
	event := new(WalletLinkingWalletLinkingRevoked)
	if err := _WalletLinking.contract.UnpackLog(event, "WalletLinkingRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
