// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringServiceManager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractIncredibleSquaringServiceManagerMetaData contains all meta data concerning the ContractIncredibleSquaringServiceManager contract.
var ContractIncredibleSquaringServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractECDSAStakeRegistry\"},{\"name\":\"_incredibleSquaringTaskManager\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSquaringTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200178b3803806200178b83398101604081905262000035916200014f565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e051611538620002536000396000818161010c015261067e01526000818161039d015281816104f90152818161059001528181610b0601528181610c8a0152610d2901526000818161075f0152818161082801526108fc0152600081816101d001528181610260015281816102e0015281816107d4015281816108a001528181610a440152610be501526115386000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101465780639926ee7d14610157578063a364f4da1461016a578063c4d66de81461017d578063e481af9d14610190578063f2fde38b1461019857600080fd5b806333cfb7b7146100ae57806338c8ee64146100d7578063715018a6146100ec578063750521f5146100f457806377ef731d14610107575b600080fd5b6100c16100bc366004611033565b6101ab565b6040516100ce9190611057565b60405180910390f35b6100ea6100e5366004611033565b610673565b005b6100ea61072c565b6100ea610102366004611159565b610740565b61012e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ce565b6033546001600160a01b031661012e565b6100ea6101653660046111aa565b6107c9565b6100ea610178366004611033565b610895565b6100ea61018b366004611033565b61092b565b6100c1610a3e565b6100ea6101a6366004611033565b610e08565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610217573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023b9190611255565b6040516333e24c9b60e11b81526001600160a01b0380831660048301529192506000917f000000000000000000000000000000000000000000000000000000000000000016906367c4993690602401602060405180830381865afa1580156102a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cb9190611272565b6001600160c01b0316905080158061036557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561033c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610360919061129b565b60ff16155b1561038157505060408051600081526020810190915292915050565b600061038c82610e7e565b90506000805b8251811015610462577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103dc576103dc6112be565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610420573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044491906112d4565b61044e9083611303565b91508061045a8161131b565b915050610392565b5060008167ffffffffffffffff81111561047e5761047e6110a4565b6040519080825280602002602001820160405280156104a7578160200160208202803683370190505b5090506000805b84518110156106665760008582815181106104cb576104cb6112be565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610540573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056491906112d4565b905060005b81811015610650576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106029190611336565b60000151868681518110610618576106186112be565b6001600160a01b03909216602092830291909101909101528461063a8161131b565b95505080806106489061131b565b915050610569565b505050808061065e9061131b565b9150506104ae565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107295760405162461bcd60e51b815260206004820152604a60248201527f6f6e6c79496e6372656469626c655371756172696e675461736b4d616e61676560448201527f723a206e6f742066726f6d206372656469626c65207371756172696e6720746160648201526939b59036b0b730b3b2b960b11b608482015260a4015b60405180910390fd5b50565b610734610f41565b61073e6000610f9b565b565b610748610f41565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb355906107949084906004016113f3565b600060405180830381600087803b1580156107ae57600080fd5b505af11580156107c2573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108115760405162461bcd60e51b815260040161072090611406565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061085f908590859060040161147e565b600060405180830381600087803b15801561087957600080fd5b505af115801561088d573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108dd5760405162461bcd60e51b815260040161072090611406565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401610794565b600054610100900460ff161580801561094b5750600054600160ff909116105b806109655750303b158015610965575060005460ff166001145b6109c85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610720565b6000805460ff1916600117905580156109eb576000805461ff0019166101001790555b6109f482610f9b565b8015610a3a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac4919061129b565b60ff16905080610ae257505060408051600081526020810190915290565b6000805b82811015610b9757604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7991906112d4565b610b839083611303565b915080610b8f8161131b565b915050610ae6565b5060008167ffffffffffffffff811115610bb357610bb36110a4565b604051908082528060200260200182016040528015610bdc578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c65919061129b565b60ff16811015610dfe57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cfd91906112d4565b905060005b81811015610de9576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9b9190611336565b60000151858581518110610db157610db16112be565b6001600160a01b039092166020928302919091019091015283610dd38161131b565b9450508080610de19061131b565b915050610d02565b50508080610df69061131b565b915050610be3565b5090949350505050565b610e10610f41565b6001600160a01b038116610e755760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610720565b61072981610f9b565b6060600080610e8c84610fed565b61ffff1667ffffffffffffffff811115610ea857610ea86110a4565b6040519080825280601f01601f191660200182016040528015610ed2576020820181803683370190505b5090506000805b825182108015610eea575061010081105b15610dfe576001811b935085841615610f31578060f81b838381518110610f1357610f136112be565b60200101906001600160f81b031916908160001a9053508160010191505b610f3a8161131b565b9050610ed9565b6033546001600160a01b0316331461073e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610720565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b8215611018576110026001846114c9565b9092169180611010816114e0565b915050610ff1565b92915050565b6001600160a01b038116811461072957600080fd5b60006020828403121561104557600080fd5b81356110508161101e565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156110985783516001600160a01b031683529284019291840191600101611073565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156110dd576110dd6110a4565b60405290565b600067ffffffffffffffff808411156110fe576110fe6110a4565b604051601f8501601f19908116603f01168101908282118183101715611126576111266110a4565b8160405280935085815286868601111561113f57600080fd5b858560208301376000602087830101525050509392505050565b60006020828403121561116b57600080fd5b813567ffffffffffffffff81111561118257600080fd5b8201601f8101841361119357600080fd5b6111a2848235602084016110e3565b949350505050565b600080604083850312156111bd57600080fd5b82356111c88161101e565b9150602083013567ffffffffffffffff808211156111e557600080fd5b90840190606082870312156111f957600080fd5b6112016110ba565b82358281111561121057600080fd5b83019150601f8201871361122357600080fd5b611232878335602085016110e3565b815260208301356020820152604083013560408201528093505050509250929050565b60006020828403121561126757600080fd5b81516110508161101e565b60006020828403121561128457600080fd5b81516001600160c01b038116811461105057600080fd5b6000602082840312156112ad57600080fd5b815160ff8116811461105057600080fd5b634e487b7160e01b600052603260045260246000fd5b6000602082840312156112e657600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115611316576113166112ed565b500190565b600060001982141561132f5761132f6112ed565b5060010190565b60006040828403121561134857600080fd5b6040516040810181811067ffffffffffffffff8211171561136b5761136b6110a4565b60405282516113798161101e565b815260208301516bffffffffffffffffffffffff8116811461139a57600080fd5b60208201529392505050565b6000815180845260005b818110156113cc576020818501810151868301820152016113b0565b818111156113de576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061105060208301846113a6565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526114a860a08401826113a6565b90506020840151606084015260408401516080840152809150509392505050565b6000828210156114db576114db6112ed565b500390565b600061ffff808316818114156114f8576114f86112ed565b600101939250505056fea264697066735822122083a3a88a13317ff35eb452e18fb84adc1023d756b56313e17d4910ffb091685064736f6c634300080c0033",
}

// ContractIncredibleSquaringServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.ABI instead.
var ContractIncredibleSquaringServiceManagerABI = ContractIncredibleSquaringServiceManagerMetaData.ABI

// ContractIncredibleSquaringServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.Bin instead.
var ContractIncredibleSquaringServiceManagerBin = ContractIncredibleSquaringServiceManagerMetaData.Bin

// DeployContractIncredibleSquaringServiceManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringServiceManager to it.
func DeployContractIncredibleSquaringServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _incredibleSquaringTaskManager common.Address) (common.Address, *types.Transaction, *ContractIncredibleSquaringServiceManager, error) {
	parsed, err := ContractIncredibleSquaringServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringServiceManagerBin), backend, _delegationManager, _registryCoordinator, _stakeRegistry, _incredibleSquaringTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringServiceManager{ContractIncredibleSquaringServiceManagerCaller: ContractIncredibleSquaringServiceManagerCaller{contract: contract}, ContractIncredibleSquaringServiceManagerTransactor: ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, ContractIncredibleSquaringServiceManagerFilterer: ContractIncredibleSquaringServiceManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManager struct {
	ContractIncredibleSquaringServiceManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringServiceManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringServiceManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringServiceManagerSession struct {
	Contract     *ContractIncredibleSquaringServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringServiceManagerCallerSession struct {
	Contract *ContractIncredibleSquaringServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringServiceManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerRaw struct {
	Contract *ContractIncredibleSquaringServiceManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringServiceManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringServiceManager creates a new instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringServiceManager, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManager{ContractIncredibleSquaringServiceManagerCaller: ContractIncredibleSquaringServiceManagerCaller{contract: contract}, ContractIncredibleSquaringServiceManagerTransactor: ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, ContractIncredibleSquaringServiceManagerFilterer: ContractIncredibleSquaringServiceManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringServiceManagerCaller creates a new read-only instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringServiceManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringServiceManagerTransactor creates a new write-only instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringServiceManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringServiceManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringServiceManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringServiceManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.ContractIncredibleSquaringServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) IncredibleSquaringTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "incredibleSquaringTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) IncredibleSquaringTaskManager() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IncredibleSquaringTaskManager(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// IncredibleSquaringTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) IncredibleSquaringTaskManager() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IncredibleSquaringTaskManager(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Owner(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Owner(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FreezeOperator(&_ContractIncredibleSquaringServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FreezeOperator(&_ContractIncredibleSquaringServiceManager.TransactOpts, operatorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Initialize(&_ContractIncredibleSquaringServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.Initialize(&_ContractIncredibleSquaringServiceManager.TransactOpts, initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSquaringServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.SetMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.SetMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.TransferOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.TransferOwnership(&_ContractIncredibleSquaringServiceManager.TransactOpts, newOwner)
}

// ContractIncredibleSquaringServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringServiceManagerInitialized)
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
		it.Event = new(ContractIncredibleSquaringServiceManagerInitialized)
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
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringServiceManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerInitializedIterator{contract: _ContractIncredibleSquaringServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringServiceManagerInitialized)
				if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringServiceManagerInitialized, error) {
	event := new(ContractIncredibleSquaringServiceManagerInitialized)
	if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
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
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringServiceManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringServiceManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
