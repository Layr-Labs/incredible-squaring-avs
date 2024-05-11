// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractPriceFeedAdapter

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

// ContractPriceFeedAdapterMetaData contains all meta data concerning the ContractPriceFeedAdapter contract.
var ContractPriceFeedAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addFeed\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_feedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeds\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestPrice\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeFeed\",\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeedAdded\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"feedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeedRemoved\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6108a38061007e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100f057806390cab41614610101578063f2fde38b14610114578063f66a1b711461012757600080fd5b806349b3ce56146100825780636a811e3b14610097578063715018a6146100e8575b600080fd5b610095610090366004610660565b610148565b005b6100cb6100a5366004610660565b80516020818301810180516001825292820191909301209152546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b61009561022e565b6000546001600160a01b03166100cb565b61009561010f3660046106b9565b610242565b610095610122366004610707565b61036d565b61013a610135366004610660565b6103e6565b6040519081526020016100df565b610150610513565b60006001600160a01b031660018260405161016b9190610759565b908152604051908190036020019020546001600160a01b031614156101c95760405162461bcd60e51b815260206004820152600f60248201526e2332b2b2103737ba103337bab7321760891b60448201526064015b60405180910390fd5b6001816040516101d99190610759565b90815260405190819003602001812080546001600160a01b03191690557f06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce906102239083906107a1565b60405180910390a150565b610236610513565b610240600061056d565b565b61024a610513565b6001600160a01b0381166102a05760405162461bcd60e51b815260206004820152601c60248201527f4665656420616464726573732063616e6e6f74206265207a65726f2e0000000060448201526064016101c0565b60008251116102f15760405162461bcd60e51b815260206004820152601760248201527f53796d626f6c2063616e6e6f7420626520656d7074792e00000000000000000060448201526064016101c0565b806001836040516103029190610759565b90815260405190819003602001812080546001600160a01b03939093166001600160a01b0319909316929092179091557f917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f37049061036190849084906107b4565b60405180910390a15050565b610375610513565b6001600160a01b0381166103da5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101c0565b6103e38161056d565b50565b6000806001836040516103f99190610759565b908152604051908190036020019020546001600160a01b03169050806104535760405162461bcd60e51b815260206004820152600f60248201526e2332b2b2103737ba103337bab7321760891b60448201526064016101c0565b600080826001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610494573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b891906107f8565b50935050925050610e1081426104ce9190610848565b1061050b5760405162461bcd60e51b815260206004820152600d60248201526c44617461206973207374616c6560981b60448201526064016101c0565b509392505050565b6000546001600160a01b031633146102405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101c0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126105e457600080fd5b813567ffffffffffffffff808211156105ff576105ff6105bd565b604051601f8301601f19908116603f01168101908282118183101715610627576106276105bd565b8160405283815286602085880101111561064057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561067257600080fd5b813567ffffffffffffffff81111561068957600080fd5b610695848285016105d3565b949350505050565b80356001600160a01b03811681146106b457600080fd5b919050565b600080604083850312156106cc57600080fd5b823567ffffffffffffffff8111156106e357600080fd5b6106ef858286016105d3565b9250506106fe6020840161069d565b90509250929050565b60006020828403121561071957600080fd5b6107228261069d565b9392505050565b60005b8381101561074457818101518382015260200161072c565b83811115610753576000848401525b50505050565b6000825161076b818460208701610729565b9190910192915050565b6000815180845261078d816020860160208601610729565b601f01601f19169290920160200192915050565b6020815260006107226020830184610775565b6040815260006107c76040830185610775565b905060018060a01b03831660208301529392505050565b805169ffffffffffffffffffff811681146106b457600080fd5b600080600080600060a0868803121561081057600080fd5b610819866107de565b945060208601519350604086015192506060860151915061083c608087016107de565b90509295509295909350565b60008282101561086857634e487b7160e01b600052601160045260246000fd5b50039056fea26469706673582212209896d48cf3c3c5dd332b619e8e473734e0b651b5f391c5acb5de009be7782c3864736f6c634300080c0033",
}

// ContractPriceFeedAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractPriceFeedAdapterMetaData.ABI instead.
var ContractPriceFeedAdapterABI = ContractPriceFeedAdapterMetaData.ABI

// ContractPriceFeedAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractPriceFeedAdapterMetaData.Bin instead.
var ContractPriceFeedAdapterBin = ContractPriceFeedAdapterMetaData.Bin

// DeployContractPriceFeedAdapter deploys a new Ethereum contract, binding an instance of ContractPriceFeedAdapter to it.
func DeployContractPriceFeedAdapter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractPriceFeedAdapter, error) {
	parsed, err := ContractPriceFeedAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractPriceFeedAdapterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractPriceFeedAdapter{ContractPriceFeedAdapterCaller: ContractPriceFeedAdapterCaller{contract: contract}, ContractPriceFeedAdapterTransactor: ContractPriceFeedAdapterTransactor{contract: contract}, ContractPriceFeedAdapterFilterer: ContractPriceFeedAdapterFilterer{contract: contract}}, nil
}

// ContractPriceFeedAdapter is an auto generated Go binding around an Ethereum contract.
type ContractPriceFeedAdapter struct {
	ContractPriceFeedAdapterCaller     // Read-only binding to the contract
	ContractPriceFeedAdapterTransactor // Write-only binding to the contract
	ContractPriceFeedAdapterFilterer   // Log filterer for contract events
}

// ContractPriceFeedAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractPriceFeedAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPriceFeedAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractPriceFeedAdapterSession struct {
	Contract     *ContractPriceFeedAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractPriceFeedAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractPriceFeedAdapterCallerSession struct {
	Contract *ContractPriceFeedAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractPriceFeedAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractPriceFeedAdapterTransactorSession struct {
	Contract     *ContractPriceFeedAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractPriceFeedAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractPriceFeedAdapterRaw struct {
	Contract *ContractPriceFeedAdapter // Generic contract binding to access the raw methods on
}

// ContractPriceFeedAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterCallerRaw struct {
	Contract *ContractPriceFeedAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// ContractPriceFeedAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractPriceFeedAdapterTransactorRaw struct {
	Contract *ContractPriceFeedAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractPriceFeedAdapter creates a new instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapter(address common.Address, backend bind.ContractBackend) (*ContractPriceFeedAdapter, error) {
	contract, err := bindContractPriceFeedAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapter{ContractPriceFeedAdapterCaller: ContractPriceFeedAdapterCaller{contract: contract}, ContractPriceFeedAdapterTransactor: ContractPriceFeedAdapterTransactor{contract: contract}, ContractPriceFeedAdapterFilterer: ContractPriceFeedAdapterFilterer{contract: contract}}, nil
}

// NewContractPriceFeedAdapterCaller creates a new read-only instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterCaller(address common.Address, caller bind.ContractCaller) (*ContractPriceFeedAdapterCaller, error) {
	contract, err := bindContractPriceFeedAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterCaller{contract: contract}, nil
}

// NewContractPriceFeedAdapterTransactor creates a new write-only instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractPriceFeedAdapterTransactor, error) {
	contract, err := bindContractPriceFeedAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterTransactor{contract: contract}, nil
}

// NewContractPriceFeedAdapterFilterer creates a new log filterer instance of ContractPriceFeedAdapter, bound to a specific deployed contract.
func NewContractPriceFeedAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractPriceFeedAdapterFilterer, error) {
	contract, err := bindContractPriceFeedAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFilterer{contract: contract}, nil
}

// bindContractPriceFeedAdapter binds a generic wrapper to an already deployed contract.
func bindContractPriceFeedAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractPriceFeedAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.ContractPriceFeedAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPriceFeedAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.contract.Transact(opts, method, params...)
}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) Feeds(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "feeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) Feeds(arg0 string) (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Feeds(&_ContractPriceFeedAdapter.CallOpts, arg0)
}

// Feeds is a free data retrieval call binding the contract method 0x6a811e3b.
//
// Solidity: function feeds(string ) view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) Feeds(arg0 string) (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Feeds(&_ContractPriceFeedAdapter.CallOpts, arg0)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) GetLatestPrice(opts *bind.CallOpts, _symbol string) (*big.Int, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "getLatestPrice", _symbol)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) GetLatestPrice(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetLatestPrice(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0xf66a1b71.
//
// Solidity: function getLatestPrice(string _symbol) view returns(int256)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) GetLatestPrice(_symbol string) (*big.Int, error) {
	return _ContractPriceFeedAdapter.Contract.GetLatestPrice(&_ContractPriceFeedAdapter.CallOpts, _symbol)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractPriceFeedAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) Owner() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Owner(&_ContractPriceFeedAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterCallerSession) Owner() (common.Address, error) {
	return _ContractPriceFeedAdapter.Contract.Owner(&_ContractPriceFeedAdapter.CallOpts)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) AddFeed(opts *bind.TransactOpts, _symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "addFeed", _symbol, _feedAddress)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) AddFeed(_symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.AddFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol, _feedAddress)
}

// AddFeed is a paid mutator transaction binding the contract method 0x90cab416.
//
// Solidity: function addFeed(string _symbol, address _feedAddress) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) AddFeed(_symbol string, _feedAddress common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.AddFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol, _feedAddress)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) RemoveFeed(opts *bind.TransactOpts, _symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "removeFeed", _symbol)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) RemoveFeed(_symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RemoveFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol)
}

// RemoveFeed is a paid mutator transaction binding the contract method 0x49b3ce56.
//
// Solidity: function removeFeed(string _symbol) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) RemoveFeed(_symbol string) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RemoveFeed(&_ContractPriceFeedAdapter.TransactOpts, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RenounceOwnership(&_ContractPriceFeedAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.RenounceOwnership(&_ContractPriceFeedAdapter.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.TransferOwnership(&_ContractPriceFeedAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractPriceFeedAdapter.Contract.TransferOwnership(&_ContractPriceFeedAdapter.TransactOpts, newOwner)
}

// ContractPriceFeedAdapterFeedAddedIterator is returned from FilterFeedAdded and is used to iterate over the raw logs and unpacked data for FeedAdded events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedAddedIterator struct {
	Event *ContractPriceFeedAdapterFeedAdded // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterFeedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterFeedAdded)
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
		it.Event = new(ContractPriceFeedAdapterFeedAdded)
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
func (it *ContractPriceFeedAdapterFeedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterFeedAdded represents a FeedAdded event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedAdded struct {
	Symbol      string
	FeedAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeedAdded is a free log retrieval operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterFeedAdded(opts *bind.FilterOpts) (*ContractPriceFeedAdapterFeedAddedIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFeedAddedIterator{contract: _ContractPriceFeedAdapter.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

// WatchFeedAdded is a free log subscription operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterFeedAdded) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "FeedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterFeedAdded)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

// ParseFeedAdded is a log parse operation binding the contract event 0x917a6ab4f9d1b23331c7307351d940aa91024649e9745f66c8d7d9ccdb5f3704.
//
// Solidity: event FeedAdded(string symbol, address feedAddress)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseFeedAdded(log types.Log) (*ContractPriceFeedAdapterFeedAdded, error) {
	event := new(ContractPriceFeedAdapterFeedAdded)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterFeedRemovedIterator is returned from FilterFeedRemoved and is used to iterate over the raw logs and unpacked data for FeedRemoved events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedRemovedIterator struct {
	Event *ContractPriceFeedAdapterFeedRemoved // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterFeedRemoved)
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
		it.Event = new(ContractPriceFeedAdapterFeedRemoved)
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
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterFeedRemoved represents a FeedRemoved event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterFeedRemoved struct {
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedRemoved is a free log retrieval operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterFeedRemoved(opts *bind.FilterOpts) (*ContractPriceFeedAdapterFeedRemovedIterator, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterFeedRemovedIterator{contract: _ContractPriceFeedAdapter.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

// WatchFeedRemoved is a free log subscription operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterFeedRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "FeedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterFeedRemoved)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

// ParseFeedRemoved is a log parse operation binding the contract event 0x06aa2b76883f4d9b405ce5d8cb618072ed06ad9a9f4a3a638d6a76f995f950ce.
//
// Solidity: event FeedRemoved(string symbol)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseFeedRemoved(log types.Log) (*ContractPriceFeedAdapterFeedRemoved, error) {
	event := new(ContractPriceFeedAdapterFeedRemoved)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPriceFeedAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterOwnershipTransferredIterator struct {
	Event *ContractPriceFeedAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPriceFeedAdapterOwnershipTransferred)
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
		it.Event = new(ContractPriceFeedAdapterOwnershipTransferred)
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
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPriceFeedAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPriceFeedAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the ContractPriceFeedAdapter contract.
type ContractPriceFeedAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractPriceFeedAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceFeedAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPriceFeedAdapterOwnershipTransferredIterator{contract: _ContractPriceFeedAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractPriceFeedAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractPriceFeedAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPriceFeedAdapterOwnershipTransferred)
				if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractPriceFeedAdapter *ContractPriceFeedAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*ContractPriceFeedAdapterOwnershipTransferred, error) {
	event := new(ContractPriceFeedAdapterOwnershipTransferred)
	if err := _ContractPriceFeedAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
