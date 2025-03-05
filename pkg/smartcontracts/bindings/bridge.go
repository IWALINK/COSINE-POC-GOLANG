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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cosineToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_walletLinking\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operationId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"BridgeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"BridgeFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operationId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"}],\"name\":\"BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operationId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"OperationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operationId\",\"type\":\"uint256\"}],\"name\":\"checkOperationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"operationId\",\"type\":\"uint256\"}],\"name\":\"confirmOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosineToken\",\"outputs\":[{\"internalType\":\"contractCosineToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"initiateIncomingBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"initiateOutgoingBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBridgeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBridgeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOperationId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"l2Wallet\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLock\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"confirmationCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredConfirmations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalBridgedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"updateBridgeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"updateBridgeLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmations\",\"type\":\"uint256\"}],\"name\":\"updateRequiredConfirmations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletLinking\",\"outputs\":[{\"internalType\":\"contractWalletLinking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040526019600355678ac7230489e8000060045569d3c21bcecceda1000000600555600260065560016007553480156200003a57600080fd5b50604051620023d2380380620023d28339810160408190526200005d9162000301565b6001805460ff191690556001600160a01b038316620000c35760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420746f6b656e2061646472657373000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b0382166200011b5760405162461bcd60e51b815260206004820152601e60248201527f496e76616c69642077616c6c6574206c696e6b696e67206164647265737300006044820152606401620000ba565b60008111620001605760405162461bcd60e51b815260206004820152601060248201526f125b9d985b1a590818da185a5b88125160821b6044820152606401620000ba565b6200016d60003362000234565b620001997fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217753362000234565b620001c57f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989263362000234565b620001f17fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc43362000234565b60018054610100600160a81b0319166101006001600160a01b0395861602179055600280546001600160a01b031916929093169190911790915560805262000342565b62000240828262000244565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000240576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620002a03390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b80516001600160a01b0381168114620002fc57600080fd5b919050565b6000806000606084860312156200031757600080fd5b6200032284620002e4565b92506200033260208501620002e4565b9150604084015190509250925092565b6080516120436200038f600039600081816103d3015281816105e901528181610709015281816108ea0152818161097601528181611013015281816110f201526111d601526120436000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806382b12dd71161010f578063b2e9949d116100a2578063c49baebe11610071578063c49baebe14610502578063d547741f14610529578063db606d5d1461053c578063fc4ce3621461054f57600080fd5b8063b2e9949d14610406578063b37fd190146104ce578063bdd9ed31146104e1578063c3c22475146104f957600080fd5b8063926d7d7f116100de578063926d7d7f146103a75780639a8a0592146103ce578063a217fddf146103f5578063aff10e90146103fd57600080fd5b806382b12dd71461037a57806382e717f7146103835780638456cb591461038c57806391d148541461039457600080fd5b80633f4ba83a1161018757806360132c9c1161015657806360132c9c1461030257806375b238fc1461030b5780638034e3561461032057806381b2db501461036757600080fd5b80633f4ba83a146102c957806346d6731d146102d15780635a1c0366146102e45780635c975abb146102f757600080fd5b8063248a9ca3116101c3578063248a9ca3146102555780632e9923f2146102785780632f2ff15d146102a357806336568abe146102b657600080fd5b806301ffc9a7146101ea57806302892dd8146102125780632246ff0414610227575b600080fd5b6101fd6101f8366004611c85565b610562565b60405190151581526020015b60405180910390f35b610225610220366004611caf565b610599565b005b610247610235366004611cdb565b60096020526000908152604090205481565b604051908152602001610209565b610247610263366004611cdb565b60009081526020819052604090206001015490565b60025461028b906001600160a01b031681565b6040516001600160a01b039091168152602001610209565b6102256102b1366004611d10565b6109fb565b6102256102c4366004611d10565b610a25565b610225610aa3565b6102256102df366004611d3c565b610ac6565b6102256102f2366004611cdb565b610b2f565b60015460ff166101fd565b61024760055481565b610247600080516020611fee83398151915281565b61033361032e366004611cdb565b610bcc565b6040805195151586526001600160a01b0390941660208601529284019190915260608301521515608082015260a001610209565b610225610375366004611cdb565b610c45565b61024760035481565b61024760065481565b610225610cc0565b6101fd6103a2366004611d10565b610ce0565b6102477fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b6102477f000000000000000000000000000000000000000000000000000000000000000081565b610247600081565b61024760075481565b610476610414366004611cdb565b6008602052600090815260409020805460018201546002830154600384015460048501546005860154600687015460078801546009909801546001600160a01b039097169795969495939492939192909160ff8082169261010090920416908a565b604080516001600160a01b03909b168b5260208b0199909952978901969096526060880194909452608087019290925260a086015260c0850152151560e0840152151561010083015261012082015261014001610209565b6102256104dc366004611d5e565b610d09565b60015461028b9061010090046001600160a01b031681565b61024760045481565b6102477f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c9892681565b610225610537366004611d10565b610f2e565b61022561054a366004611d9a565b610f53565b61022561055d366004611cdb565b611286565b60006001600160e01b03198216637965db0b60e01b148061059357506301ffc9a760e01b6001600160e01b03198316145b92915050565b6105a1611466565b826105e75760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a5908130c881dd85b1b195d607a1b60448201526064015b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000082036106565760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742062726964676520746f2073616d6520636861696e000000000060448201526064016105de565b60045481101561069f5760405162461bcd60e51b8152602060048201526014602482015273416d6f756e742062656c6f77206d696e696d756d60601b60448201526064016105de565b6005548111156106e85760405162461bcd60e51b8152602060048201526014602482015273416d6f756e742061626f7665206d6178696d756d60601b60448201526064016105de565b6002546040516307e1ac5360e01b81523360048201819052602482018690527f00000000000000000000000000000000000000000000000000000000000000006044830152916000916001600160a01b03909116906307e1ac5390606401602060405180830381865afa158015610763573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107879190611dd3565b9050806107a65760405162461bcd60e51b81526004016105de90611df5565b60006107c96127106107c3600354876114ae90919063ffffffff16565b906114c1565b905060006107d785836114cd565b6001546040516323b872dd60e01b81526001600160a01b03878116600483015230602483015260448201899052929350610100909104909116906323b872dd906064016020604051808303816000875af1158015610839573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085d9190611dd3565b6108a15760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016105de565b60078054600091826108b283611e4d565b90915550600081815260086020908152604080832080546001600160a01b0319166001600160a01b038b1617815560018082018e90557f00000000000000000000000000000000000000000000000000000000000000006002830155600382018d9055600482018890556005820189905542600683015560078201805461ffff1916909117905560098082018590558c8552909252909120549192509061095990846114d9565b6000898152600960209081526040918290209290925580518b81527f00000000000000000000000000000000000000000000000000000000000000009281019290925281018990526060810184905260808101859052600160a08201526001600160a01b0387169083907f29a0a92de40eaa235167917a9583a15f5bdc9d868497696b06561a53000360899060c00160405180910390a3505050505050505050565b600082815260208190526040902060010154610a16816114e5565b610a2083836114ef565b505050565b6001600160a01b0381163314610a955760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084016105de565b610a9f8282611573565b5050565b600080516020611fee833981519152610abb816114e5565b610ac36115d8565b50565b600080516020611fee833981519152610ade816114e5565b81831115610b235760405162461bcd60e51b815260206004820152601260248201527109ad2dc40daeae6e840c4ca40787a40dac2f60731b60448201526064016105de565b50600491909155600555565b600080516020611fee833981519152610b47816114e5565b6101f4821115610b905760405162461bcd60e51b81526020600482015260146024820152734665652063616e6e6f742065786365656420352560601b60448201526064016105de565b60038290556040518281527f42dfb00d085d601e55327921154ae76c1b24270b026c5a0c51caee18eb4c401f9060200160405180910390a15050565b6000818152600860205260408120805482918291829182916001600160a01b0316610c095760008060008060009550955095509550955050610c3c565b805460048201546009830154600790930154600198506001600160a01b0390921696509450909250610100900460ff1690505b91939590929450565b600080516020611fee833981519152610c5d816114e5565b6001821015610cba5760405162461bcd60e51b8152602060048201526024808201527f4d7573742072657175697265206174206c65617374203120636f6e6669726d616044820152633a34b7b760e11b60648201526084016105de565b50600655565b600080516020611fee833981519152610cd8816114e5565b610ac361162a565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020611fee833981519152610d21816114e5565b6001600160a01b038216610d775760405162461bcd60e51b815260206004820152601b60248201527f43616e6e6f742073656e6420746f207a65726f2061646472657373000000000060448201526064016105de565b6001546001600160a01b03610100909104811690851603610eb4576001546040516370a0823160e01b815230600482015260009161010090046001600160a01b0316906370a0823190602401602060405180830381865afa158015610de0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e049190611e66565b9050600060015b6103e88111610e575760008181526009602052604090205415610e4557600081815260096020526040902054610e429083906114d9565b91505b80610e4f81611e4d565b915050610e0b565b5084610e6383836114cd565b1015610eb15760405162461bcd60e51b815260206004820152601c60248201527f43616e6e6f742072657363756520747261636b656420746f6b656e730000000060448201526064016105de565b50505b60405163a9059cbb60e01b81526001600160a01b0383811660048301526024820185905285169063a9059cbb906044016020604051808303816000875af1158015610f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f279190611dd3565b5050505050565b600082815260208190526040902060010154610f49816114e5565b610a208383611573565b7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4610f7d816114e5565b610f85611466565b6001600160a01b038516610fd05760405162461bcd60e51b8152602060048201526012602482015271496e76616c6964204c31206164647265737360701b60448201526064016105de565b836110115760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a5908130c881dd85b1b195d607a1b60448201526064016105de565b7f000000000000000000000000000000000000000000000000000000000000000083036110805760405162461bcd60e51b815260206004820152601d60248201527f43616e6e6f74206272696467652066726f6d2073616d6520636861696e00000060448201526064016105de565b6004548210156110c95760405162461bcd60e51b8152602060048201526014602482015273416d6f756e742062656c6f77206d696e696d756d60601b60448201526064016105de565b6002546040516307e1ac5360e01b81526001600160a01b038781166004830152602482018790527f0000000000000000000000000000000000000000000000000000000000000000604483015260009216906307e1ac5390606401602060405180830381865afa158015611141573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111659190611dd3565b9050806111845760405162461bcd60e51b81526004016105de90611df5565b600780546000918261119583611e4d565b90915550600081815260086020908152604080832080546001600160a01b038d166001600160a01b031990911681178255600182018c9055600282018b90557f000000000000000000000000000000000000000000000000000000000000000060038301819055600483018b90556005830186905542600684015560078301805461ffff191690556009830186905583518d81529485018c905284840152606084018a90526080840185905260a084019490945290519394509284917f29a0a92de40eaa235167917a9583a15f5bdc9d868497696b06561a5300036089919081900360c00190a35050505050505050565b7f21702c8af46127c7fa207f89d0b0a8441bb32959a0ac7df790e9ab1a25c989266112b0816114e5565b6112b8611466565b600082815260086020526040902080546001600160a01b031661131d5760405162461bcd60e51b815260206004820152601860248201527f4f7065726174696f6e20646f6573206e6f74206578697374000000000000000060448201526064016105de565b6007810154610100900460ff16156113775760405162461bcd60e51b815260206004820152601b60248201527f4f7065726174696f6e20616c726561647920636f6d706c65746564000000000060448201526064016105de565b33600090815260088201602052604090205460ff16156113e55760405162461bcd60e51b815260206004820152602360248201527f416c726561647920636f6e6669726d656420627920746869732076616c6964616044820152623a37b960e91b60648201526084016105de565b3360009081526008820160205260409020805460ff191660019081179091556009820154611412916114d9565b60098201819055604051908152339084907f46873a1bbb891f66c4c3ae91f074a92d9133db9b69be62530b9c225e1840edd69060200160405180910390a3600654816009015410610a2057610a2083611665565b60015460ff16156114ac5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016105de565b565b60006114ba8284611e7f565b9392505050565b60006114ba8284611e96565b60006114ba8284611eb8565b60006114ba8284611ecb565b610ac38133611a35565b6114f98282610ce0565b610a9f576000828152602081815260408083206001600160a01b03851684529091529020805460ff1916600117905561152f3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61157d8282610ce0565b15610a9f576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6115e0611a8e565b6001805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b611632611466565b6001805460ff1916811790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361160d565b60008181526008602052604090206007810154610100900460ff16156116cd5760405162461bcd60e51b815260206004820152601b60248201527f4f7065726174696f6e20616c726561647920636f6d706c65746564000000000060448201526064016105de565b600654816009015410156117235760405162461bcd60e51b815260206004820152601860248201527f4e6f7420656e6f75676820636f6e6669726d6174696f6e73000000000000000060448201526064016105de565b60078101805461ff001981166101001790915560ff1615611828576001805460038301546004808501546040516317727d3560e01b8152918201929092526024810191909152604481019290925261010090046001600160a01b0316906317727d3590606401600060405180830381600087803b1580156117a357600080fd5b505af11580156117b7573d6000803e3d6000fd5b50506001546005840154604051630852cd8d60e31b815260048101919091526101009091046001600160a01b031692506342966c689150602401600060405180830381600087803b15801561180b57600080fd5b505af115801561181f573d6000803e3d6000fd5b505050506119d9565b60015460028201546004808401546040516317727d3560e01b8152918201929092526024810191909152600060448201526101009091046001600160a01b0316906317727d3590606401600060405180830381600087803b15801561188c57600080fd5b505af11580156118a0573d6000803e3d6000fd5b5050505060048101546002820154600090815260096020526040902054106118fd57600481015460028201546000908152600960205260409020546118e4916114cd565b6002820154600090815260096020526040902055611912565b60028101546000908152600960205260408120555b600154815460048084015460405163a9059cbb60e01b81526001600160a01b03938416928101929092526024820152610100909204169063a9059cbb906044016020604051808303816000875af1158015611971573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119959190611dd3565b6119d95760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b60448201526064016105de565b8054600482015460078301546040805192835260ff909116151560208301526001600160a01b039092169184917fc47fe959e81eba1006063d7963508474d286e6641b2ce932875dba8b463abbb0910160405180910390a35050565b611a3f8282610ce0565b610a9f57611a4c81611ad7565b611a57836020611ae9565b604051602001611a68929190611f02565b60408051601f198184030181529082905262461bcd60e51b82526105de91600401611f77565b60015460ff166114ac5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016105de565b60606105936001600160a01b03831660145b60606000611af8836002611e7f565b611b03906002611ecb565b67ffffffffffffffff811115611b1b57611b1b611faa565b6040519080825280601f01601f191660200182016040528015611b45576020820181803683370190505b509050600360fc1b81600081518110611b6057611b60611fc0565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611b8f57611b8f611fc0565b60200101906001600160f81b031916908160001a9053506000611bb3846002611e7f565b611bbe906001611ecb565b90505b6001811115611c36576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611bf257611bf2611fc0565b1a60f81b828281518110611c0857611c08611fc0565b60200101906001600160f81b031916908160001a90535060049490941c93611c2f81611fd6565b9050611bc1565b5083156114ba5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016105de565b600060208284031215611c9757600080fd5b81356001600160e01b0319811681146114ba57600080fd5b600080600060608486031215611cc457600080fd5b505081359360208301359350604090920135919050565b600060208284031215611ced57600080fd5b5035919050565b80356001600160a01b0381168114611d0b57600080fd5b919050565b60008060408385031215611d2357600080fd5b82359150611d3360208401611cf4565b90509250929050565b60008060408385031215611d4f57600080fd5b50508035926020909101359150565b600080600060608486031215611d7357600080fd5b611d7c84611cf4565b925060208401359150611d9160408501611cf4565b90509250925092565b60008060008060808587031215611db057600080fd5b611db985611cf4565b966020860135965060408601359560600135945092505050565b600060208284031215611de557600080fd5b815180151581146114ba57600080fd5b60208082526022908201527f4c312061646472657373206e6f74206c696e6b656420746f204c322077616c6c604082015261195d60f21b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060018201611e5f57611e5f611e37565b5060010190565b600060208284031215611e7857600080fd5b5051919050565b808202811582820484141761059357610593611e37565b600082611eb357634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561059357610593611e37565b8082018082111561059357610593611e37565b60005b83811015611ef9578181015183820152602001611ee1565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611f3a816017850160208801611ede565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611f6b816028840160208801611ede565b01602801949350505050565b6020815260008251806020840152611f96816040850160208701611ede565b601f01601f19169190910160400192915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600081611fe557611fe5611e37565b50600019019056fea49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a26469706673582212202051b3a9879889919730a8619ba6d2d703a44e37c29c052ba2621a1e6bc5edab64736f6c63430008130033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _cosineToken common.Address, _walletLinking common.Address, _chainId *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, _cosineToken, _walletLinking, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) ADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.ADMINROLE(&_Bridge.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) ADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.ADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) VALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) VALIDATORROLE() ([32]byte, error) {
	return _Bridge.Contract.VALIDATORROLE(&_Bridge.CallOpts)
}

// VALIDATORROLE is a free data retrieval call binding the contract method 0xc49baebe.
//
// Solidity: function VALIDATOR_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) VALIDATORROLE() ([32]byte, error) {
	return _Bridge.Contract.VALIDATORROLE(&_Bridge.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x82b12dd7.
//
// Solidity: function bridgeFee() view returns(uint256)
func (_Bridge *BridgeCaller) BridgeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x82b12dd7.
//
// Solidity: function bridgeFee() view returns(uint256)
func (_Bridge *BridgeSession) BridgeFee() (*big.Int, error) {
	return _Bridge.Contract.BridgeFee(&_Bridge.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x82b12dd7.
//
// Solidity: function bridgeFee() view returns(uint256)
func (_Bridge *BridgeCallerSession) BridgeFee() (*big.Int, error) {
	return _Bridge.Contract.BridgeFee(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCallerSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// CheckOperationStatus is a free data retrieval call binding the contract method 0x8034e356.
//
// Solidity: function checkOperationStatus(uint256 operationId) view returns(bool exists, address initiator, uint256 amount, uint256 confirmations, bool isCompleted)
func (_Bridge *BridgeCaller) CheckOperationStatus(opts *bind.CallOpts, operationId *big.Int) (struct {
	Exists        bool
	Initiator     common.Address
	Amount        *big.Int
	Confirmations *big.Int
	IsCompleted   bool
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "checkOperationStatus", operationId)

	outstruct := new(struct {
		Exists        bool
		Initiator     common.Address
		Amount        *big.Int
		Confirmations *big.Int
		IsCompleted   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Initiator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Confirmations = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsCompleted = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// CheckOperationStatus is a free data retrieval call binding the contract method 0x8034e356.
//
// Solidity: function checkOperationStatus(uint256 operationId) view returns(bool exists, address initiator, uint256 amount, uint256 confirmations, bool isCompleted)
func (_Bridge *BridgeSession) CheckOperationStatus(operationId *big.Int) (struct {
	Exists        bool
	Initiator     common.Address
	Amount        *big.Int
	Confirmations *big.Int
	IsCompleted   bool
}, error) {
	return _Bridge.Contract.CheckOperationStatus(&_Bridge.CallOpts, operationId)
}

// CheckOperationStatus is a free data retrieval call binding the contract method 0x8034e356.
//
// Solidity: function checkOperationStatus(uint256 operationId) view returns(bool exists, address initiator, uint256 amount, uint256 confirmations, bool isCompleted)
func (_Bridge *BridgeCallerSession) CheckOperationStatus(operationId *big.Int) (struct {
	Exists        bool
	Initiator     common.Address
	Amount        *big.Int
	Confirmations *big.Int
	IsCompleted   bool
}, error) {
	return _Bridge.Contract.CheckOperationStatus(&_Bridge.CallOpts, operationId)
}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_Bridge *BridgeCaller) CosineToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "cosineToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_Bridge *BridgeSession) CosineToken() (common.Address, error) {
	return _Bridge.Contract.CosineToken(&_Bridge.CallOpts)
}

// CosineToken is a free data retrieval call binding the contract method 0xbdd9ed31.
//
// Solidity: function cosineToken() view returns(address)
func (_Bridge *BridgeCallerSession) CosineToken() (common.Address, error) {
	return _Bridge.Contract.CosineToken(&_Bridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// MaxBridgeAmount is a free data retrieval call binding the contract method 0x60132c9c.
//
// Solidity: function maxBridgeAmount() view returns(uint256)
func (_Bridge *BridgeCaller) MaxBridgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "maxBridgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBridgeAmount is a free data retrieval call binding the contract method 0x60132c9c.
//
// Solidity: function maxBridgeAmount() view returns(uint256)
func (_Bridge *BridgeSession) MaxBridgeAmount() (*big.Int, error) {
	return _Bridge.Contract.MaxBridgeAmount(&_Bridge.CallOpts)
}

// MaxBridgeAmount is a free data retrieval call binding the contract method 0x60132c9c.
//
// Solidity: function maxBridgeAmount() view returns(uint256)
func (_Bridge *BridgeCallerSession) MaxBridgeAmount() (*big.Int, error) {
	return _Bridge.Contract.MaxBridgeAmount(&_Bridge.CallOpts)
}

// MinBridgeAmount is a free data retrieval call binding the contract method 0xc3c22475.
//
// Solidity: function minBridgeAmount() view returns(uint256)
func (_Bridge *BridgeCaller) MinBridgeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minBridgeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBridgeAmount is a free data retrieval call binding the contract method 0xc3c22475.
//
// Solidity: function minBridgeAmount() view returns(uint256)
func (_Bridge *BridgeSession) MinBridgeAmount() (*big.Int, error) {
	return _Bridge.Contract.MinBridgeAmount(&_Bridge.CallOpts)
}

// MinBridgeAmount is a free data retrieval call binding the contract method 0xc3c22475.
//
// Solidity: function minBridgeAmount() view returns(uint256)
func (_Bridge *BridgeCallerSession) MinBridgeAmount() (*big.Int, error) {
	return _Bridge.Contract.MinBridgeAmount(&_Bridge.CallOpts)
}

// NextOperationId is a free data retrieval call binding the contract method 0xaff10e90.
//
// Solidity: function nextOperationId() view returns(uint256)
func (_Bridge *BridgeCaller) NextOperationId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nextOperationId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOperationId is a free data retrieval call binding the contract method 0xaff10e90.
//
// Solidity: function nextOperationId() view returns(uint256)
func (_Bridge *BridgeSession) NextOperationId() (*big.Int, error) {
	return _Bridge.Contract.NextOperationId(&_Bridge.CallOpts)
}

// NextOperationId is a free data retrieval call binding the contract method 0xaff10e90.
//
// Solidity: function nextOperationId() view returns(uint256)
func (_Bridge *BridgeCallerSession) NextOperationId() (*big.Int, error) {
	return _Bridge.Contract.NextOperationId(&_Bridge.CallOpts)
}

// Operations is a free data retrieval call binding the contract method 0xb2e9949d.
//
// Solidity: function operations(uint256 ) view returns(address initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, uint256 timestamp, bool isLock, bool isCompleted, uint256 confirmationCount)
func (_Bridge *BridgeCaller) Operations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Initiator          common.Address
	L2Wallet           [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Amount             *big.Int
	Fee                *big.Int
	Timestamp          *big.Int
	IsLock             bool
	IsCompleted        bool
	ConfirmationCount  *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "operations", arg0)

	outstruct := new(struct {
		Initiator          common.Address
		L2Wallet           [32]byte
		SourceChainId      *big.Int
		DestinationChainId *big.Int
		Amount             *big.Int
		Fee                *big.Int
		Timestamp          *big.Int
		IsLock             bool
		IsCompleted        bool
		ConfirmationCount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Initiator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.L2Wallet = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SourceChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DestinationChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsLock = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsCompleted = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.ConfirmationCount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Operations is a free data retrieval call binding the contract method 0xb2e9949d.
//
// Solidity: function operations(uint256 ) view returns(address initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, uint256 timestamp, bool isLock, bool isCompleted, uint256 confirmationCount)
func (_Bridge *BridgeSession) Operations(arg0 *big.Int) (struct {
	Initiator          common.Address
	L2Wallet           [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Amount             *big.Int
	Fee                *big.Int
	Timestamp          *big.Int
	IsLock             bool
	IsCompleted        bool
	ConfirmationCount  *big.Int
}, error) {
	return _Bridge.Contract.Operations(&_Bridge.CallOpts, arg0)
}

// Operations is a free data retrieval call binding the contract method 0xb2e9949d.
//
// Solidity: function operations(uint256 ) view returns(address initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, uint256 timestamp, bool isLock, bool isCompleted, uint256 confirmationCount)
func (_Bridge *BridgeCallerSession) Operations(arg0 *big.Int) (struct {
	Initiator          common.Address
	L2Wallet           [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Amount             *big.Int
	Fee                *big.Int
	Timestamp          *big.Int
	IsLock             bool
	IsCompleted        bool
	ConfirmationCount  *big.Int
}, error) {
	return _Bridge.Contract.Operations(&_Bridge.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Bridge *BridgeCaller) RequiredConfirmations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "requiredConfirmations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Bridge *BridgeSession) RequiredConfirmations() (*big.Int, error) {
	return _Bridge.Contract.RequiredConfirmations(&_Bridge.CallOpts)
}

// RequiredConfirmations is a free data retrieval call binding the contract method 0x82e717f7.
//
// Solidity: function requiredConfirmations() view returns(uint256)
func (_Bridge *BridgeCallerSession) RequiredConfirmations() (*big.Int, error) {
	return _Bridge.Contract.RequiredConfirmations(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// TotalBridgedTokens is a free data retrieval call binding the contract method 0x2246ff04.
//
// Solidity: function totalBridgedTokens(uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) TotalBridgedTokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "totalBridgedTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgedTokens is a free data retrieval call binding the contract method 0x2246ff04.
//
// Solidity: function totalBridgedTokens(uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) TotalBridgedTokens(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TotalBridgedTokens(&_Bridge.CallOpts, arg0)
}

// TotalBridgedTokens is a free data retrieval call binding the contract method 0x2246ff04.
//
// Solidity: function totalBridgedTokens(uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalBridgedTokens(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.TotalBridgedTokens(&_Bridge.CallOpts, arg0)
}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_Bridge *BridgeCaller) WalletLinking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "walletLinking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_Bridge *BridgeSession) WalletLinking() (common.Address, error) {
	return _Bridge.Contract.WalletLinking(&_Bridge.CallOpts)
}

// WalletLinking is a free data retrieval call binding the contract method 0x2e9923f2.
//
// Solidity: function walletLinking() view returns(address)
func (_Bridge *BridgeCallerSession) WalletLinking() (common.Address, error) {
	return _Bridge.Contract.WalletLinking(&_Bridge.CallOpts)
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0xfc4ce362.
//
// Solidity: function confirmOperation(uint256 operationId) returns()
func (_Bridge *BridgeTransactor) ConfirmOperation(opts *bind.TransactOpts, operationId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "confirmOperation", operationId)
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0xfc4ce362.
//
// Solidity: function confirmOperation(uint256 operationId) returns()
func (_Bridge *BridgeSession) ConfirmOperation(operationId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ConfirmOperation(&_Bridge.TransactOpts, operationId)
}

// ConfirmOperation is a paid mutator transaction binding the contract method 0xfc4ce362.
//
// Solidity: function confirmOperation(uint256 operationId) returns()
func (_Bridge *BridgeTransactorSession) ConfirmOperation(operationId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ConfirmOperation(&_Bridge.TransactOpts, operationId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// InitiateIncomingBridge is a paid mutator transaction binding the contract method 0xdb606d5d.
//
// Solidity: function initiateIncomingBridge(address l1Address, bytes32 l2Wallet, uint256 sourceChainId, uint256 amount) returns()
func (_Bridge *BridgeTransactor) InitiateIncomingBridge(opts *bind.TransactOpts, l1Address common.Address, l2Wallet [32]byte, sourceChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initiateIncomingBridge", l1Address, l2Wallet, sourceChainId, amount)
}

// InitiateIncomingBridge is a paid mutator transaction binding the contract method 0xdb606d5d.
//
// Solidity: function initiateIncomingBridge(address l1Address, bytes32 l2Wallet, uint256 sourceChainId, uint256 amount) returns()
func (_Bridge *BridgeSession) InitiateIncomingBridge(l1Address common.Address, l2Wallet [32]byte, sourceChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.InitiateIncomingBridge(&_Bridge.TransactOpts, l1Address, l2Wallet, sourceChainId, amount)
}

// InitiateIncomingBridge is a paid mutator transaction binding the contract method 0xdb606d5d.
//
// Solidity: function initiateIncomingBridge(address l1Address, bytes32 l2Wallet, uint256 sourceChainId, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) InitiateIncomingBridge(l1Address common.Address, l2Wallet [32]byte, sourceChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.InitiateIncomingBridge(&_Bridge.TransactOpts, l1Address, l2Wallet, sourceChainId, amount)
}

// InitiateOutgoingBridge is a paid mutator transaction binding the contract method 0x02892dd8.
//
// Solidity: function initiateOutgoingBridge(bytes32 l2Wallet, uint256 destinationChainId, uint256 amount) returns()
func (_Bridge *BridgeTransactor) InitiateOutgoingBridge(opts *bind.TransactOpts, l2Wallet [32]byte, destinationChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initiateOutgoingBridge", l2Wallet, destinationChainId, amount)
}

// InitiateOutgoingBridge is a paid mutator transaction binding the contract method 0x02892dd8.
//
// Solidity: function initiateOutgoingBridge(bytes32 l2Wallet, uint256 destinationChainId, uint256 amount) returns()
func (_Bridge *BridgeSession) InitiateOutgoingBridge(l2Wallet [32]byte, destinationChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.InitiateOutgoingBridge(&_Bridge.TransactOpts, l2Wallet, destinationChainId, amount)
}

// InitiateOutgoingBridge is a paid mutator transaction binding the contract method 0x02892dd8.
//
// Solidity: function initiateOutgoingBridge(bytes32 l2Wallet, uint256 destinationChainId, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) InitiateOutgoingBridge(l2Wallet [32]byte, destinationChainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.InitiateOutgoingBridge(&_Bridge.TransactOpts, l2Wallet, destinationChainId, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xb37fd190.
//
// Solidity: function rescueTokens(address token, uint256 amount, address recipient) returns()
func (_Bridge *BridgeTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "rescueTokens", token, amount, recipient)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xb37fd190.
//
// Solidity: function rescueTokens(address token, uint256 amount, address recipient) returns()
func (_Bridge *BridgeSession) RescueTokens(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RescueTokens(&_Bridge.TransactOpts, token, amount, recipient)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xb37fd190.
//
// Solidity: function rescueTokens(address token, uint256 amount, address recipient) returns()
func (_Bridge *BridgeTransactorSession) RescueTokens(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RescueTokens(&_Bridge.TransactOpts, token, amount, recipient)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// UpdateBridgeFee is a paid mutator transaction binding the contract method 0x5a1c0366.
//
// Solidity: function updateBridgeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactor) UpdateBridgeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateBridgeFee", newFee)
}

// UpdateBridgeFee is a paid mutator transaction binding the contract method 0x5a1c0366.
//
// Solidity: function updateBridgeFee(uint256 newFee) returns()
func (_Bridge *BridgeSession) UpdateBridgeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateBridgeFee(&_Bridge.TransactOpts, newFee)
}

// UpdateBridgeFee is a paid mutator transaction binding the contract method 0x5a1c0366.
//
// Solidity: function updateBridgeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactorSession) UpdateBridgeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateBridgeFee(&_Bridge.TransactOpts, newFee)
}

// UpdateBridgeLimits is a paid mutator transaction binding the contract method 0x46d6731d.
//
// Solidity: function updateBridgeLimits(uint256 min, uint256 max) returns()
func (_Bridge *BridgeTransactor) UpdateBridgeLimits(opts *bind.TransactOpts, min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateBridgeLimits", min, max)
}

// UpdateBridgeLimits is a paid mutator transaction binding the contract method 0x46d6731d.
//
// Solidity: function updateBridgeLimits(uint256 min, uint256 max) returns()
func (_Bridge *BridgeSession) UpdateBridgeLimits(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateBridgeLimits(&_Bridge.TransactOpts, min, max)
}

// UpdateBridgeLimits is a paid mutator transaction binding the contract method 0x46d6731d.
//
// Solidity: function updateBridgeLimits(uint256 min, uint256 max) returns()
func (_Bridge *BridgeTransactorSession) UpdateBridgeLimits(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateBridgeLimits(&_Bridge.TransactOpts, min, max)
}

// UpdateRequiredConfirmations is a paid mutator transaction binding the contract method 0x81b2db50.
//
// Solidity: function updateRequiredConfirmations(uint256 confirmations) returns()
func (_Bridge *BridgeTransactor) UpdateRequiredConfirmations(opts *bind.TransactOpts, confirmations *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateRequiredConfirmations", confirmations)
}

// UpdateRequiredConfirmations is a paid mutator transaction binding the contract method 0x81b2db50.
//
// Solidity: function updateRequiredConfirmations(uint256 confirmations) returns()
func (_Bridge *BridgeSession) UpdateRequiredConfirmations(confirmations *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateRequiredConfirmations(&_Bridge.TransactOpts, confirmations)
}

// UpdateRequiredConfirmations is a paid mutator transaction binding the contract method 0x81b2db50.
//
// Solidity: function updateRequiredConfirmations(uint256 confirmations) returns()
func (_Bridge *BridgeTransactorSession) UpdateRequiredConfirmations(confirmations *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateRequiredConfirmations(&_Bridge.TransactOpts, confirmations)
}

// BridgeBridgeCompletedIterator is returned from FilterBridgeCompleted and is used to iterate over the raw logs and unpacked data for BridgeCompleted events raised by the Bridge contract.
type BridgeBridgeCompletedIterator struct {
	Event *BridgeBridgeCompleted // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeCompleted)
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
		it.Event = new(BridgeBridgeCompleted)
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
func (it *BridgeBridgeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeCompleted represents a BridgeCompleted event raised by the Bridge contract.
type BridgeBridgeCompleted struct {
	OperationId *big.Int
	Initiator   common.Address
	Amount      *big.Int
	IsLock      bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBridgeCompleted is a free log retrieval operation binding the contract event 0xc47fe959e81eba1006063d7963508474d286e6641b2ce932875dba8b463abbb0.
//
// Solidity: event BridgeCompleted(uint256 indexed operationId, address indexed initiator, uint256 amount, bool isLock)
func (_Bridge *BridgeFilterer) FilterBridgeCompleted(opts *bind.FilterOpts, operationId []*big.Int, initiator []common.Address) (*BridgeBridgeCompletedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeCompleted", operationIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeCompletedIterator{contract: _Bridge.contract, event: "BridgeCompleted", logs: logs, sub: sub}, nil
}

// WatchBridgeCompleted is a free log subscription operation binding the contract event 0xc47fe959e81eba1006063d7963508474d286e6641b2ce932875dba8b463abbb0.
//
// Solidity: event BridgeCompleted(uint256 indexed operationId, address indexed initiator, uint256 amount, bool isLock)
func (_Bridge *BridgeFilterer) WatchBridgeCompleted(opts *bind.WatchOpts, sink chan<- *BridgeBridgeCompleted, operationId []*big.Int, initiator []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeCompleted", operationIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeCompleted)
				if err := _Bridge.contract.UnpackLog(event, "BridgeCompleted", log); err != nil {
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

// ParseBridgeCompleted is a log parse operation binding the contract event 0xc47fe959e81eba1006063d7963508474d286e6641b2ce932875dba8b463abbb0.
//
// Solidity: event BridgeCompleted(uint256 indexed operationId, address indexed initiator, uint256 amount, bool isLock)
func (_Bridge *BridgeFilterer) ParseBridgeCompleted(log types.Log) (*BridgeBridgeCompleted, error) {
	event := new(BridgeBridgeCompleted)
	if err := _Bridge.contract.UnpackLog(event, "BridgeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBridgeFeeUpdatedIterator is returned from FilterBridgeFeeUpdated and is used to iterate over the raw logs and unpacked data for BridgeFeeUpdated events raised by the Bridge contract.
type BridgeBridgeFeeUpdatedIterator struct {
	Event *BridgeBridgeFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeFeeUpdated)
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
		it.Event = new(BridgeBridgeFeeUpdated)
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
func (it *BridgeBridgeFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeFeeUpdated represents a BridgeFeeUpdated event raised by the Bridge contract.
type BridgeBridgeFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeFeeUpdated is a free log retrieval operation binding the contract event 0x42dfb00d085d601e55327921154ae76c1b24270b026c5a0c51caee18eb4c401f.
//
// Solidity: event BridgeFeeUpdated(uint256 newFee)
func (_Bridge *BridgeFilterer) FilterBridgeFeeUpdated(opts *bind.FilterOpts) (*BridgeBridgeFeeUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeFeeUpdatedIterator{contract: _Bridge.contract, event: "BridgeFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeFeeUpdated is a free log subscription operation binding the contract event 0x42dfb00d085d601e55327921154ae76c1b24270b026c5a0c51caee18eb4c401f.
//
// Solidity: event BridgeFeeUpdated(uint256 newFee)
func (_Bridge *BridgeFilterer) WatchBridgeFeeUpdated(opts *bind.WatchOpts, sink chan<- *BridgeBridgeFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeFeeUpdated)
				if err := _Bridge.contract.UnpackLog(event, "BridgeFeeUpdated", log); err != nil {
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

// ParseBridgeFeeUpdated is a log parse operation binding the contract event 0x42dfb00d085d601e55327921154ae76c1b24270b026c5a0c51caee18eb4c401f.
//
// Solidity: event BridgeFeeUpdated(uint256 newFee)
func (_Bridge *BridgeFilterer) ParseBridgeFeeUpdated(log types.Log) (*BridgeBridgeFeeUpdated, error) {
	event := new(BridgeBridgeFeeUpdated)
	if err := _Bridge.contract.UnpackLog(event, "BridgeFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeBridgeInitiatedIterator is returned from FilterBridgeInitiated and is used to iterate over the raw logs and unpacked data for BridgeInitiated events raised by the Bridge contract.
type BridgeBridgeInitiatedIterator struct {
	Event *BridgeBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeInitiated)
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
		it.Event = new(BridgeBridgeInitiated)
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
func (it *BridgeBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeInitiated represents a BridgeInitiated event raised by the Bridge contract.
type BridgeBridgeInitiated struct {
	OperationId        *big.Int
	Initiator          common.Address
	L2Wallet           [32]byte
	SourceChainId      *big.Int
	DestinationChainId *big.Int
	Amount             *big.Int
	Fee                *big.Int
	IsLock             bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeInitiated is a free log retrieval operation binding the contract event 0x29a0a92de40eaa235167917a9583a15f5bdc9d868497696b06561a5300036089.
//
// Solidity: event BridgeInitiated(uint256 indexed operationId, address indexed initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, bool isLock)
func (_Bridge *BridgeFilterer) FilterBridgeInitiated(opts *bind.FilterOpts, operationId []*big.Int, initiator []common.Address) (*BridgeBridgeInitiatedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeInitiated", operationIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeInitiatedIterator{contract: _Bridge.contract, event: "BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgeInitiated is a free log subscription operation binding the contract event 0x29a0a92de40eaa235167917a9583a15f5bdc9d868497696b06561a5300036089.
//
// Solidity: event BridgeInitiated(uint256 indexed operationId, address indexed initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, bool isLock)
func (_Bridge *BridgeFilterer) WatchBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BridgeBridgeInitiated, operationId []*big.Int, initiator []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeInitiated", operationIdRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeInitiated)
				if err := _Bridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
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

// ParseBridgeInitiated is a log parse operation binding the contract event 0x29a0a92de40eaa235167917a9583a15f5bdc9d868497696b06561a5300036089.
//
// Solidity: event BridgeInitiated(uint256 indexed operationId, address indexed initiator, bytes32 l2Wallet, uint256 sourceChainId, uint256 destinationChainId, uint256 amount, uint256 fee, bool isLock)
func (_Bridge *BridgeFilterer) ParseBridgeInitiated(log types.Log) (*BridgeBridgeInitiated, error) {
	event := new(BridgeBridgeInitiated)
	if err := _Bridge.contract.UnpackLog(event, "BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOperationConfirmedIterator is returned from FilterOperationConfirmed and is used to iterate over the raw logs and unpacked data for OperationConfirmed events raised by the Bridge contract.
type BridgeOperationConfirmedIterator struct {
	Event *BridgeOperationConfirmed // Event containing the contract specifics and raw log

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
func (it *BridgeOperationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOperationConfirmed)
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
		it.Event = new(BridgeOperationConfirmed)
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
func (it *BridgeOperationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOperationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOperationConfirmed represents a OperationConfirmed event raised by the Bridge contract.
type BridgeOperationConfirmed struct {
	OperationId   *big.Int
	Validator     common.Address
	Confirmations *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperationConfirmed is a free log retrieval operation binding the contract event 0x46873a1bbb891f66c4c3ae91f074a92d9133db9b69be62530b9c225e1840edd6.
//
// Solidity: event OperationConfirmed(uint256 indexed operationId, address indexed validator, uint256 confirmations)
func (_Bridge *BridgeFilterer) FilterOperationConfirmed(opts *bind.FilterOpts, operationId []*big.Int, validator []common.Address) (*BridgeOperationConfirmedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OperationConfirmed", operationIdRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOperationConfirmedIterator{contract: _Bridge.contract, event: "OperationConfirmed", logs: logs, sub: sub}, nil
}

// WatchOperationConfirmed is a free log subscription operation binding the contract event 0x46873a1bbb891f66c4c3ae91f074a92d9133db9b69be62530b9c225e1840edd6.
//
// Solidity: event OperationConfirmed(uint256 indexed operationId, address indexed validator, uint256 confirmations)
func (_Bridge *BridgeFilterer) WatchOperationConfirmed(opts *bind.WatchOpts, sink chan<- *BridgeOperationConfirmed, operationId []*big.Int, validator []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OperationConfirmed", operationIdRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOperationConfirmed)
				if err := _Bridge.contract.UnpackLog(event, "OperationConfirmed", log); err != nil {
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

// ParseOperationConfirmed is a log parse operation binding the contract event 0x46873a1bbb891f66c4c3ae91f074a92d9133db9b69be62530b9c225e1840edd6.
//
// Solidity: event OperationConfirmed(uint256 indexed operationId, address indexed validator, uint256 confirmations)
func (_Bridge *BridgeFilterer) ParseOperationConfirmed(log types.Log) (*BridgeOperationConfirmed, error) {
	event := new(BridgeOperationConfirmed)
	if err := _Bridge.contract.UnpackLog(event, "OperationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
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
		it.Event = new(BridgeRoleAdminChanged)
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
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
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
		it.Event = new(BridgeRoleGranted)
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
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
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
		it.Event = new(BridgeRoleRevoked)
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
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
