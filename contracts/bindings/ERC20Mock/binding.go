// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractERC20Mock

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

// ContractERC20MockMetaData contains all meta data concerning the ContractERC20Mock contract.
var ContractERC20MockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b506040805160208082018084526000808452845192830190945292815281519192909161003f9160039161005b565b50805161005390600490602084019061005b565b50505061012f565b828054610067906100f4565b90600052602060002090601f01602090048101928261008957600085556100cf565b82601f106100a257805160ff19168380011785556100cf565b828001600101855582156100cf579182015b828111156100cf5782518255916020019190600101906100b4565b506100db9291506100df565b5090565b5b808211156100db57600081556001016100e0565b600181811c9082168061010857607f821691505b6020821081141561012957634e487b7160e01b600052602260045260246000fd5b50919050565b6109808061013e6000396000f3fe608060405234801561001057600080fd5b50600436106100a45760003560e01c806306fdde03146100a9578063095ea7b3146100c757806318160ddd146100ea57806323b872dd146100fc578063313ce5671461010f578063395093511461011e57806340c10f191461013157806370a082311461014657806395d89b411461016f578063a457c2d714610177578063a9059cbb1461018a578063dd62ed3e1461019d575b600080fd5b6100b16101b0565b6040516100be919061079d565b60405180910390f35b6100da6100d536600461080e565b610242565b60405190151581526020016100be565b6002545b6040519081526020016100be565b6100da61010a366004610838565b61025a565b604051601281526020016100be565b6100da61012c36600461080e565b61027e565b61014461013f36600461080e565b6102a0565b005b6100ee610154366004610874565b6001600160a01b031660009081526020819052604090205490565b6100b16102ae565b6100da61018536600461080e565b6102bd565b6100da61019836600461080e565b61033d565b6100ee6101ab366004610896565b61034b565b6060600380546101bf906108c9565b80601f01602080910402602001604051908101604052809291908181526020018280546101eb906108c9565b80156102385780601f1061020d57610100808354040283529160200191610238565b820191906000526020600020905b81548152906001019060200180831161021b57829003601f168201915b5050505050905090565b600033610250818585610376565b5060019392505050565b60003361026885828561049a565b610273858585610514565b506001949350505050565b600033610250818585610291838361034b565b61029b9190610904565b610376565b6102aa82826106d0565b5050565b6060600480546101bf906108c9565b600033816102cb828661034b565b9050838110156103305760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102738286868403610376565b600033610250818585610514565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103d85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610327565b6001600160a01b0382166104395760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610327565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006104a6848461034b565b9050600019811461050e57818110156105015760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610327565b61050e8484848403610376565b50505050565b6001600160a01b0383166105785760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610327565b6001600160a01b0382166105da5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610327565b6001600160a01b038316600090815260208190526040902054818110156106525760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610327565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610689908490610904565b92505081905550826001600160a01b0316846001600160a01b031660008051602061092b833981519152846040516106c391815260200190565b60405180910390a361050e565b6001600160a01b0382166107265760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610327565b80600260008282546107389190610904565b90915550506001600160a01b03821660009081526020819052604081208054839290610765908490610904565b90915550506040518181526001600160a01b0383169060009060008051602061092b8339815191529060200160405180910390a35050565b600060208083528351808285015260005b818110156107ca578581018301518582016040015282016107ae565b818111156107dc576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b038116811461080957600080fd5b919050565b6000806040838503121561082157600080fd5b61082a836107f2565b946020939093013593505050565b60008060006060848603121561084d57600080fd5b610856846107f2565b9250610864602085016107f2565b9150604084013590509250925092565b60006020828403121561088657600080fd5b61088f826107f2565b9392505050565b600080604083850312156108a957600080fd5b6108b2836107f2565b91506108c0602084016107f2565b90509250929050565b600181811c908216806108dd57607f821691505b602082108114156108fe57634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561092557634e487b7160e01b600052601160045260246000fd5b50019056feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa264697066735822122090a206a81686d242c9126f013e890370ba7c9c838c5adb856ea0cb549089547064736f6c634300080c0033",
}

// ContractERC20MockABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractERC20MockMetaData.ABI instead.
var ContractERC20MockABI = ContractERC20MockMetaData.ABI

// ContractERC20MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractERC20MockMetaData.Bin instead.
var ContractERC20MockBin = ContractERC20MockMetaData.Bin

// DeployContractERC20Mock deploys a new Ethereum contract, binding an instance of ContractERC20Mock to it.
func DeployContractERC20Mock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractERC20Mock, error) {
	parsed, err := ContractERC20MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractERC20MockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractERC20Mock{ContractERC20MockCaller: ContractERC20MockCaller{contract: contract}, ContractERC20MockTransactor: ContractERC20MockTransactor{contract: contract}, ContractERC20MockFilterer: ContractERC20MockFilterer{contract: contract}}, nil
}

// ContractERC20Mock is an auto generated Go binding around an Ethereum contract.
type ContractERC20Mock struct {
	ContractERC20MockCaller     // Read-only binding to the contract
	ContractERC20MockTransactor // Write-only binding to the contract
	ContractERC20MockFilterer   // Log filterer for contract events
}

// ContractERC20MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractERC20MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractERC20MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractERC20MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractERC20MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractERC20MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractERC20MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractERC20MockSession struct {
	Contract     *ContractERC20Mock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractERC20MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractERC20MockCallerSession struct {
	Contract *ContractERC20MockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractERC20MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractERC20MockTransactorSession struct {
	Contract     *ContractERC20MockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractERC20MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractERC20MockRaw struct {
	Contract *ContractERC20Mock // Generic contract binding to access the raw methods on
}

// ContractERC20MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractERC20MockCallerRaw struct {
	Contract *ContractERC20MockCaller // Generic read-only contract binding to access the raw methods on
}

// ContractERC20MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractERC20MockTransactorRaw struct {
	Contract *ContractERC20MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractERC20Mock creates a new instance of ContractERC20Mock, bound to a specific deployed contract.
func NewContractERC20Mock(address common.Address, backend bind.ContractBackend) (*ContractERC20Mock, error) {
	contract, err := bindContractERC20Mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractERC20Mock{ContractERC20MockCaller: ContractERC20MockCaller{contract: contract}, ContractERC20MockTransactor: ContractERC20MockTransactor{contract: contract}, ContractERC20MockFilterer: ContractERC20MockFilterer{contract: contract}}, nil
}

// NewContractERC20MockCaller creates a new read-only instance of ContractERC20Mock, bound to a specific deployed contract.
func NewContractERC20MockCaller(address common.Address, caller bind.ContractCaller) (*ContractERC20MockCaller, error) {
	contract, err := bindContractERC20Mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractERC20MockCaller{contract: contract}, nil
}

// NewContractERC20MockTransactor creates a new write-only instance of ContractERC20Mock, bound to a specific deployed contract.
func NewContractERC20MockTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractERC20MockTransactor, error) {
	contract, err := bindContractERC20Mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractERC20MockTransactor{contract: contract}, nil
}

// NewContractERC20MockFilterer creates a new log filterer instance of ContractERC20Mock, bound to a specific deployed contract.
func NewContractERC20MockFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractERC20MockFilterer, error) {
	contract, err := bindContractERC20Mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractERC20MockFilterer{contract: contract}, nil
}

// bindContractERC20Mock binds a generic wrapper to an already deployed contract.
func bindContractERC20Mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractERC20MockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractERC20Mock *ContractERC20MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractERC20Mock.Contract.ContractERC20MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractERC20Mock *ContractERC20MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.ContractERC20MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractERC20Mock *ContractERC20MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.ContractERC20MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractERC20Mock *ContractERC20MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractERC20Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractERC20Mock *ContractERC20MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractERC20Mock *ContractERC20MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractERC20Mock.Contract.Allowance(&_ContractERC20Mock.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractERC20Mock.Contract.Allowance(&_ContractERC20Mock.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractERC20Mock.Contract.BalanceOf(&_ContractERC20Mock.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractERC20Mock.Contract.BalanceOf(&_ContractERC20Mock.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractERC20Mock *ContractERC20MockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractERC20Mock *ContractERC20MockSession) Decimals() (uint8, error) {
	return _ContractERC20Mock.Contract.Decimals(&_ContractERC20Mock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ContractERC20Mock *ContractERC20MockCallerSession) Decimals() (uint8, error) {
	return _ContractERC20Mock.Contract.Decimals(&_ContractERC20Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractERC20Mock *ContractERC20MockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractERC20Mock *ContractERC20MockSession) Name() (string, error) {
	return _ContractERC20Mock.Contract.Name(&_ContractERC20Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ContractERC20Mock *ContractERC20MockCallerSession) Name() (string, error) {
	return _ContractERC20Mock.Contract.Name(&_ContractERC20Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractERC20Mock *ContractERC20MockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractERC20Mock *ContractERC20MockSession) Symbol() (string, error) {
	return _ContractERC20Mock.Contract.Symbol(&_ContractERC20Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ContractERC20Mock *ContractERC20MockCallerSession) Symbol() (string, error) {
	return _ContractERC20Mock.Contract.Symbol(&_ContractERC20Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractERC20Mock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockSession) TotalSupply() (*big.Int, error) {
	return _ContractERC20Mock.Contract.TotalSupply(&_ContractERC20Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractERC20Mock *ContractERC20MockCallerSession) TotalSupply() (*big.Int, error) {
	return _ContractERC20Mock.Contract.TotalSupply(&_ContractERC20Mock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Approve(&_ContractERC20Mock.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Approve(&_ContractERC20Mock.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.DecreaseAllowance(&_ContractERC20Mock.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.DecreaseAllowance(&_ContractERC20Mock.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.IncreaseAllowance(&_ContractERC20Mock.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.IncreaseAllowance(&_ContractERC20Mock.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractERC20Mock *ContractERC20MockTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractERC20Mock *ContractERC20MockSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Mint(&_ContractERC20Mock.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractERC20Mock *ContractERC20MockTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Mint(&_ContractERC20Mock.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Transfer(&_ContractERC20Mock.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.Transfer(&_ContractERC20Mock.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.TransferFrom(&_ContractERC20Mock.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractERC20Mock *ContractERC20MockTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractERC20Mock.Contract.TransferFrom(&_ContractERC20Mock.TransactOpts, from, to, amount)
}

// ContractERC20MockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractERC20Mock contract.
type ContractERC20MockApprovalIterator struct {
	Event *ContractERC20MockApproval // Event containing the contract specifics and raw log

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
func (it *ContractERC20MockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractERC20MockApproval)
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
		it.Event = new(ContractERC20MockApproval)
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
func (it *ContractERC20MockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractERC20MockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractERC20MockApproval represents a Approval event raised by the ContractERC20Mock contract.
type ContractERC20MockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractERC20Mock *ContractERC20MockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractERC20MockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractERC20Mock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractERC20MockApprovalIterator{contract: _ContractERC20Mock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractERC20Mock *ContractERC20MockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractERC20MockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractERC20Mock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractERC20MockApproval)
				if err := _ContractERC20Mock.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ContractERC20Mock *ContractERC20MockFilterer) ParseApproval(log types.Log) (*ContractERC20MockApproval, error) {
	event := new(ContractERC20MockApproval)
	if err := _ContractERC20Mock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractERC20MockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractERC20Mock contract.
type ContractERC20MockTransferIterator struct {
	Event *ContractERC20MockTransfer // Event containing the contract specifics and raw log

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
func (it *ContractERC20MockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractERC20MockTransfer)
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
		it.Event = new(ContractERC20MockTransfer)
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
func (it *ContractERC20MockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractERC20MockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractERC20MockTransfer represents a Transfer event raised by the ContractERC20Mock contract.
type ContractERC20MockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractERC20Mock *ContractERC20MockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractERC20MockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractERC20Mock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractERC20MockTransferIterator{contract: _ContractERC20Mock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractERC20Mock *ContractERC20MockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractERC20MockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractERC20Mock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractERC20MockTransfer)
				if err := _ContractERC20Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ContractERC20Mock *ContractERC20MockFilterer) ParseTransfer(log types.Log) (*ContractERC20MockTransfer, error) {
	event := new(ContractERC20MockTransfer)
	if err := _ContractERC20Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
