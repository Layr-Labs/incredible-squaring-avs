// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMockERC20

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

// ContractMockERC20MetaData contains all meta data concerning the ContractMockERC20 contract.
var ContractMockERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040523461031357604080519081016001600160401b03811182821017610226576040908152600a82526926b7b1b5902a37b5b2b760b11b602083015280519081016001600160401b038111828210176102265760405260038152624d434b60e81b602082015281516001600160401b03811161022657600354600181811c91168015610309575b602082101461020857601f81116102a6575b50602092601f821160011461024557928192935f9261023a575b50508160011b915f199060031b1c1916176003555b80516001600160401b03811161022657600454600181811c9116801561021c575b602082101461020857601f81116101a5575b50602091601f8211600114610145579181925f9261013a575b50508160011b915f199060031b1c1916176004555b6040516108b490816103188239f35b015190505f80610116565b601f1982169260045f52805f20915f5b85811061018d57508360019510610175575b505050811b0160045561012b565b01515f1960f88460031b161c191690555f8080610167565b91926020600181928685015181550194019201610155565b60045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f830160051c810191602084106101fe575b601f0160051c01905b8181106101f357506100fd565b5f81556001016101e6565b90915081906101dd565b634e487b7160e01b5f52602260045260245ffd5b90607f16906100eb565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b5565b601f1982169360035f52805f20915f5b86811061028e5750836001959610610276575b505050811b016003556100ca565b01515f1960f88460031b161c191690555f8080610268565b91926020600181928685015181550194019201610255565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c810191602084106102ff575b601f0160051c01905b8181106102f4575061009b565b5f81556001016102e7565b90915081906102de565b90607f1690610089565b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816306fdde03146104b757508063095ea7b31461049157806318160ddd1461047457806323b872dd14610447578063313ce5671461042c57806339509351146103de57806340c10f191461031a57806370a08231146102e357806395d89b41146101c8578063a457c2d714610125578063a9059cbb146100f45763dd62ed3e146100a0575f80fd5b346100f05760403660031901126100f0576100b96105b0565b6100c16105c6565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b5f80fd5b346100f05760403660031901126100f05761011a6101106105b0565b6024359033610701565b602060405160018152f35b346100f05760403660031901126100f05761013e6105b0565b60243590335f52600160205260405f2060018060a01b0382165f5260205260405f2054918083106101755761011a920390336105fd565b60405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608490fd5b346100f0575f3660031901126100f0576040515f6004548060011c906001811680156102d9575b6020831081146102c5578285529081156102a95750600114610254575b50819003601f01601f191681019067ffffffffffffffff8211818310176102405761023c82918260405282610586565b0390f35b634e487b7160e01b5f52604160045260245ffd5b905060045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5f905b8282106102935750602091508201018261020c565b600181602092548385880101520191019061027e565b90506020925060ff191682840152151560051b8201018261020c565b634e487b7160e01b5f52602260045260245ffd5b91607f16916101ef565b346100f05760203660031901126100f0576001600160a01b036103046105b0565b165f525f602052602060405f2054604051908152f35b346100f05760403660031901126100f0576103336105b0565b6001600160a01b03166024358115610399577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020826103765f946002546105dc565b6002558484528382526040842061038e8282546105dc565b9055604051908152a3005b60405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606490fd5b346100f05760403660031901126100f05761011a6103fa6105b0565b335f52600160205260405f2060018060a01b0382165f5260205261042560405f2060243590546105dc565b90336105fd565b346100f0575f3660031901126100f057602060405160128152f35b346100f05760603660031901126100f05761011a6104636105b0565b61046b6105c6565b60443591610701565b346100f0575f3660031901126100f0576020600254604051908152f35b346100f05760403660031901126100f05761011a6104ad6105b0565b60243590336105fd565b346100f0575f3660031901126100f0575f6003548060011c9060018116801561057c575b6020831081146102c5578285529081156102a957506001146105275750819003601f01601f191681019067ffffffffffffffff8211818310176102405761023c82918260405282610586565b905060035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5f905b8282106105665750602091508201018261020c565b6001816020925483858801015201910190610551565b91607f16916104db565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b03821682036100f057565b602435906001600160a01b03821682036100f057565b919082018092116105e957565b634e487b7160e01b5f52601160045260245ffd5b6001600160a01b03169081156106b0576001600160a01b03169182156106605760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b031690811561082b576001600160a01b03169182156107da57815f525f60205260405f205481811061078657817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92602092855f525f84520360405f2055845f525f825260405f2061077b8282546105dc565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea2646970667358221220d20ddd7e0acc21c8810cd29f18143294f1478b4900957b29508413fa31f0bb8464736f6c634300081b0033",
}

// ContractMockERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMockERC20MetaData.ABI instead.
var ContractMockERC20ABI = ContractMockERC20MetaData.ABI

// ContractMockERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMockERC20MetaData.Bin instead.
var ContractMockERC20Bin = ContractMockERC20MetaData.Bin

// DeployContractMockERC20 deploys a new Ethereum contract, binding an instance of ContractMockERC20 to it.
func DeployContractMockERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractMockERC20, error) {
	parsed, err := ContractMockERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMockERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMockERC20{ContractMockERC20Caller: ContractMockERC20Caller{contract: contract}, ContractMockERC20Transactor: ContractMockERC20Transactor{contract: contract}, ContractMockERC20Filterer: ContractMockERC20Filterer{contract: contract}}, nil
}

// ContractMockERC20 is an auto generated Go binding around an Ethereum contract.
type ContractMockERC20 struct {
	ContractMockERC20Caller     // Read-only binding to the contract
	ContractMockERC20Transactor // Write-only binding to the contract
	ContractMockERC20Filterer   // Log filterer for contract events
}

// ContractMockERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMockERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMockERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMockERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMockERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMockERC20Session struct {
	Contract     *ContractMockERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractMockERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMockERC20CallerSession struct {
	Contract *ContractMockERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractMockERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMockERC20TransactorSession struct {
	Contract     *ContractMockERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractMockERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMockERC20Raw struct {
	Contract *ContractMockERC20 // Generic contract binding to access the raw methods on
}

// ContractMockERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMockERC20CallerRaw struct {
	Contract *ContractMockERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ContractMockERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMockERC20TransactorRaw struct {
	Contract *ContractMockERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMockERC20 creates a new instance of ContractMockERC20, bound to a specific deployed contract.
func NewContractMockERC20(address common.Address, backend bind.ContractBackend) (*ContractMockERC20, error) {
	contract, err := bindContractMockERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20{ContractMockERC20Caller: ContractMockERC20Caller{contract: contract}, ContractMockERC20Transactor: ContractMockERC20Transactor{contract: contract}, ContractMockERC20Filterer: ContractMockERC20Filterer{contract: contract}}, nil
}

// NewContractMockERC20Caller creates a new read-only instance of ContractMockERC20, bound to a specific deployed contract.
func NewContractMockERC20Caller(address common.Address, caller bind.ContractCaller) (*ContractMockERC20Caller, error) {
	contract, err := bindContractMockERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20Caller{contract: contract}, nil
}

// NewContractMockERC20Transactor creates a new write-only instance of ContractMockERC20, bound to a specific deployed contract.
func NewContractMockERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ContractMockERC20Transactor, error) {
	contract, err := bindContractMockERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20Transactor{contract: contract}, nil
}

// NewContractMockERC20Filterer creates a new log filterer instance of ContractMockERC20, bound to a specific deployed contract.
func NewContractMockERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ContractMockERC20Filterer, error) {
	contract, err := bindContractMockERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20Filterer{contract: contract}, nil
}

// bindContractMockERC20 binds a generic wrapper to an already deployed contract.
func bindContractMockERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMockERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMockERC20 *ContractMockERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMockERC20.Contract.ContractMockERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMockERC20 *ContractMockERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.ContractMockERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMockERC20 *ContractMockERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.ContractMockERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMockERC20 *ContractMockERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMockERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMockERC20 *ContractMockERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMockERC20 *ContractMockERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractMockERC20.Contract.Allowance(&_ContractMockERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractMockERC20.Contract.Allowance(&_ContractMockERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractMockERC20.Contract.BalanceOf(&_ContractMockERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractMockERC20.Contract.BalanceOf(&_ContractMockERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractMockERC20 *ContractMockERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractMockERC20 *ContractMockERC20Session) Decimals() (uint8, error) {
	return _ContractMockERC20.Contract.Decimals(&_ContractMockERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractMockERC20 *ContractMockERC20CallerSession) Decimals() (uint8, error) {
	return _ContractMockERC20.Contract.Decimals(&_ContractMockERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractMockERC20 *ContractMockERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractMockERC20 *ContractMockERC20Session) Name() (string, error) {
	return _ContractMockERC20.Contract.Name(&_ContractMockERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractMockERC20 *ContractMockERC20CallerSession) Name() (string, error) {
	return _ContractMockERC20.Contract.Name(&_ContractMockERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractMockERC20 *ContractMockERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractMockERC20 *ContractMockERC20Session) Symbol() (string, error) {
	return _ContractMockERC20.Contract.Symbol(&_ContractMockERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractMockERC20 *ContractMockERC20CallerSession) Symbol() (string, error) {
	return _ContractMockERC20.Contract.Symbol(&_ContractMockERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMockERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20Session) TotalSupply() (*big.Int, error) {
	return _ContractMockERC20.Contract.TotalSupply(&_ContractMockERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractMockERC20 *ContractMockERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ContractMockERC20.Contract.TotalSupply(&_ContractMockERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Approve(&_ContractMockERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Approve(&_ContractMockERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.DecreaseAllowance(&_ContractMockERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.DecreaseAllowance(&_ContractMockERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.IncreaseAllowance(&_ContractMockERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractMockERC20 *ContractMockERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.IncreaseAllowance(&_ContractMockERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractMockERC20 *ContractMockERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractMockERC20 *ContractMockERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Mint(&_ContractMockERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractMockERC20 *ContractMockERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Mint(&_ContractMockERC20.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Transfer(&_ContractMockERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.Transfer(&_ContractMockERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.TransferFrom(&_ContractMockERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractMockERC20 *ContractMockERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractMockERC20.Contract.TransferFrom(&_ContractMockERC20.TransactOpts, from, to, amount)
}

// ContractMockERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractMockERC20 contract.
type ContractMockERC20ApprovalIterator struct {
	Event *ContractMockERC20Approval // Event containing the contract specifics and raw log

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
func (it *ContractMockERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMockERC20Approval)
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
		it.Event = new(ContractMockERC20Approval)
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
func (it *ContractMockERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMockERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMockERC20Approval represents a Approval event raised by the ContractMockERC20 contract.
type ContractMockERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractMockERC20 *ContractMockERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractMockERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractMockERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20ApprovalIterator{contract: _ContractMockERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractMockERC20 *ContractMockERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractMockERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractMockERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMockERC20Approval)
				if err := _ContractMockERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ContractMockERC20 *ContractMockERC20Filterer) ParseApproval(log types.Log) (*ContractMockERC20Approval, error) {
	event := new(ContractMockERC20Approval)
	if err := _ContractMockERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMockERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractMockERC20 contract.
type ContractMockERC20TransferIterator struct {
	Event *ContractMockERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ContractMockERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMockERC20Transfer)
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
		it.Event = new(ContractMockERC20Transfer)
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
func (it *ContractMockERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMockERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMockERC20Transfer represents a Transfer event raised by the ContractMockERC20 contract.
type ContractMockERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractMockERC20 *ContractMockERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractMockERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractMockERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractMockERC20TransferIterator{contract: _ContractMockERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractMockERC20 *ContractMockERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractMockERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractMockERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMockERC20Transfer)
				if err := _ContractMockERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ContractMockERC20 *ContractMockERC20Filterer) ParseTransfer(log types.Log) (*ContractMockERC20Transfer, error) {
	event := new(ContractMockERC20Transfer)
	if err := _ContractMockERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
