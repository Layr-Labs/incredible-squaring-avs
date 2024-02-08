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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_incredibleSquaringTaskManager\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSquaringTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016e4380380620016e483398101604081905262000035916200014f565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e051611491620002536000396000818161010c01526106860152600081816103a5015281816105010152818161059801528181610b0e01528181610c920152610d310152600081816107670152818161083001526109040152600081816101d00152818161025f015281816102df015281816107dc015281816108a801528181610a4c0152610bed01526114916000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b146101465780639926ee7d14610157578063a364f4da1461016a578063c4d66de81461017d578063e481af9d14610190578063f2fde38b1461019857600080fd5b806333cfb7b7146100ae57806338c8ee64146100d7578063715018a6146100ec578063750521f5146100f457806377ef731d14610107575b600080fd5b6100c16100bc366004610fa4565b6101ab565b6040516100ce9190610fc8565b60405180910390f35b6100ea6100e5366004610fa4565b61067b565b005b6100ea610734565b6100ea6101023660046110ca565b610748565b61012e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ce565b6033546001600160a01b031661012e565b6100ea61016536600461111b565b6107d1565b6100ea610178366004610fa4565b61089d565b6100ea61018b366004610fa4565b610933565b6100c1610a46565b6100ea6101a6366004610fa4565b610e10565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610217573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023b91906111c6565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ca91906111df565b90506001600160c01b038116158061036457507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561033b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035f9190611208565b60ff16155b1561038057505060408051600081526020810190915292915050565b6000610394826001600160c01b0316610e86565b90506000805b825181101561046a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103e4576103e461122b565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610428573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044c91906111c6565b6104569083611257565b9150806104628161126f565b91505061039a565b5060008167ffffffffffffffff81111561048657610486611015565b6040519080825280602002602001820160405280156104af578160200160208202803683370190505b5090506000805b845181101561066e5760008582815181106104d3576104d361122b565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c91906111c6565b905060005b81811015610658576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060a919061128a565b600001518686815181106106205761062061122b565b6001600160a01b0390921660209283029190910190910152846106428161126f565b95505080806106509061126f565b915050610571565b50505080806106669061126f565b9150506104b6565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107315760405162461bcd60e51b815260206004820152604a60248201527f6f6e6c79496e6372656469626c655371756172696e675461736b4d616e61676560448201527f723a206e6f742066726f6d206372656469626c65207371756172696e6720746160648201526939b59036b0b730b3b2b960b11b608482015260a4015b60405180910390fd5b50565b61073c610ee3565b6107466000610f3d565b565b610750610ee3565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061079c908490600401611356565b600060405180830381600087803b1580156107b657600080fd5b505af11580156107ca573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108195760405162461bcd60e51b815260040161072890611369565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061086790859085906004016113e1565b600060405180830381600087803b15801561088157600080fd5b505af1158015610895573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108e55760405162461bcd60e51b815260040161072890611369565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da9060240161079c565b600054610100900460ff16158080156109535750600054600160ff909116105b8061096d5750303b15801561096d575060005460ff166001145b6109d05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610728565b6000805460ff1916600117905580156109f3576000805461ff0019166101001790555b6109fc82610f3d565b8015610a42576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610acc9190611208565b60ff16905080610aea57505060408051600081526020810190915290565b6000805b82811015610b9f57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8191906111c6565b610b8b9083611257565b915080610b978161126f565b915050610aee565b5060008167ffffffffffffffff811115610bbb57610bbb611015565b604051908082528060200260200182016040528015610be4578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6d9190611208565b60ff16811015610e0657604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0591906111c6565b905060005b81811015610df1576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da3919061128a565b60000151858581518110610db957610db961122b565b6001600160a01b039092166020928302919091019091015283610ddb8161126f565b9450508080610de99061126f565b915050610d0a565b50508080610dfe9061126f565b915050610beb565b5090949350505050565b610e18610ee3565b6001600160a01b038116610e7d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610728565b61073181610f3d565b60606000805b610100811015610edc576001811b915083821615610ecc57828160f81b604051602001610eba92919061142c565b60405160208183030381529060405292505b610ed58161126f565b9050610e8c565b5050919050565b6033546001600160a01b031633146107465760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610728565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038116811461073157600080fd5b600060208284031215610fb657600080fd5b8135610fc181610f8f565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156110095783516001600160a01b031683529284019291840191600101610fe4565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561104e5761104e611015565b60405290565b600067ffffffffffffffff8084111561106f5761106f611015565b604051601f8501601f19908116603f0116810190828211818310171561109757611097611015565b816040528093508581528686860111156110b057600080fd5b858560208301376000602087830101525050509392505050565b6000602082840312156110dc57600080fd5b813567ffffffffffffffff8111156110f357600080fd5b8201601f8101841361110457600080fd5b61111384823560208401611054565b949350505050565b6000806040838503121561112e57600080fd5b823561113981610f8f565b9150602083013567ffffffffffffffff8082111561115657600080fd5b908401906060828703121561116a57600080fd5b61117261102b565b82358281111561118157600080fd5b83019150601f8201871361119457600080fd5b6111a387833560208501611054565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156111d857600080fd5b5051919050565b6000602082840312156111f157600080fd5b81516001600160c01b0381168114610fc157600080fd5b60006020828403121561121a57600080fd5b815160ff81168114610fc157600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000821982111561126a5761126a611241565b500190565b600060001982141561128357611283611241565b5060010190565b60006040828403121561129c57600080fd5b6040516040810181811067ffffffffffffffff821117156112bf576112bf611015565b60405282516112cd81610f8f565b815260208301516bffffffffffffffffffffffff811681146112ee57600080fd5b60208201529392505050565b60005b838110156113155781810151838201526020016112fd565b83811115611324576000848401525b50505050565b600081518084526113428160208601602086016112fa565b601f01601f19169290920160200192915050565b602081526000610fc1602083018461132a565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261140b60a084018261132a565b90506020840151606084015260408401516080840152809150509392505050565b6000835161143e8184602088016112fa565b6001600160f81b031993909316919092019081526001019291505056fea2646970667358221220c0f6d6debbbb6ab7f774e839119d0e67e7b9075c74f04d6173d5b7da8c6c30cc64736f6c634300080c0033",
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
