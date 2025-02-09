// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractTradeAlgoServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractTradeAlgoServiceManagerMetaData contains all meta data concerning the ContractTradeAlgoServiceManager contract.
var ContractTradeAlgoServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_incredibleSquaringTaskManager\",\"type\":\"address\",\"internalType\":\"contractITradeAlgoTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSquaringTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITradeAlgoTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b5f38038062001b5f83398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118e86200027760003960008181610151015261093b01526000818161065a015281816107b60152818161084d01528181610c7601528181610dfa0152610e99015260008181610485015281816105140152818161059401528181610a0801528181610a9e01528181610bb40152610d5501526000818161031501526103f301526000818161010c01528181610a5c01528181610afa0152610b7901526118e86000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b146101735780639926ee7d14610184578063a364f4da14610197578063a98fb355146101aa578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a61461014457806377ef731d1461014c575b600080fd5b6100cc6100c736600461118e565b6101d8565b005b6100e16100dc366004611218565b610460565b6040516100ee919061123c565b60405180910390f35b6100cc610105366004611218565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109e9565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b031661012c565b6100cc61019236600461133e565b6109fd565b6100cc6101a5366004611218565b610a93565b6100cc6101b83660046113e9565b610b5a565b6100e1610bae565b6100cc6101d3366004611218565b610f78565b6101e0610fee565b60005b818110156103db578282828181106101fd576101fd61143a565b905060200281019061020f9190611450565b610220906040810190602001611218565b6001600160a01b03166323b872dd33308686868181106102425761024261143a565b90506020028101906102549190611450565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190611480565b508282828181106102e2576102e261143a565b90506020028101906102f49190611450565b610305906040810190602001611218565b6001600160a01b031663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008585858181106103465761034661143a565b90506020028101906103589190611450565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca9190611480565b506103d4816114b8565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a908590859060040161156c565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f0919061167a565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611693565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116bc565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b0316611048565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106106995761069961143a565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610701919061167a565b61070b90836116df565b915080610717816114b8565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b611289565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b84518110156109235760008582815181106107885761078861143a565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610821919061167a565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116f7565b600001518686815181106108d5576108d561143a565b6001600160a01b0390921660209283029190910190910152846108f7816114b8565b9550508080610905906114b8565b915050610826565b505050808061091b906114b8565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109e65760405162461bcd60e51b815260206004820152604a60248201527f6f6e6c79496e6372656469626c655371756172696e675461736b4d616e61676560448201527f723a206e6f742066726f6d206372656469626c65207371756172696e6720746160648201526939b59036b0b730b3b2b960b11b608482015260a4015b60405180910390fd5b50565b6109f1610fee565b6109fb600061110b565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a455760405162461bcd60e51b81526004016109dd90611756565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a908590859060040161181b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610adb5760405162461bcd60e51b81526004016109dd90611756565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b3f57600080fd5b505af1158015610b53573d6000803e3d6000fd5b5050505050565b610b62610fee565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b25908490600401611866565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3491906116bc565b60ff16905080610c5257505060408051600081526020810190915290565b6000805b82811015610d0757604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce9919061167a565b610cf390836116df565b915080610cff816114b8565b915050610c56565b5060008167ffffffffffffffff811115610d2357610d23611289565b604051908082528060200260200182016040528015610d4c578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610db1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd591906116bc565b60ff16811015610f6e57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d919061167a565b905060005b81811015610f59576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610ee7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0b91906116f7565b60000151858581518110610f2157610f2161143a565b6001600160a01b039092166020928302919091019091015283610f43816114b8565b9450508080610f51906114b8565b915050610e72565b50508080610f66906114b8565b915050610d53565b5090949350505050565b610f80610fee565b6001600160a01b038116610fe55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109dd565b6109e68161110b565b6033546001600160a01b031633146109fb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109dd565b60606000806110568461115d565b61ffff1667ffffffffffffffff81111561107257611072611289565b6040519080825280601f01601f19166020018201604052801561109c576020820181803683370190505b5090506000805b8251821080156110b4575061010081105b15610f6e576001811b9350858416156110fb578060f81b8383815181106110dd576110dd61143a565b60200101906001600160f81b031916908160001a9053508160010191505b611104816114b8565b90506110a3565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561118857611172600184611879565b909216918061118081611890565b915050611161565b92915050565b600080602083850312156111a157600080fd5b823567ffffffffffffffff808211156111b957600080fd5b818501915085601f8301126111cd57600080fd5b8135818111156111dc57600080fd5b8660208260051b85010111156111f157600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109e657600080fd5b60006020828403121561122a57600080fd5b813561123581611203565b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561127d5783516001600160a01b031683529284019291840191600101611258565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112c2576112c2611289565b60405290565b600067ffffffffffffffff808411156112e3576112e3611289565b604051601f8501601f19908116603f0116810190828211818310171561130b5761130b611289565b8160405280935085815286868601111561132457600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561135157600080fd5b823561135c81611203565b9150602083013567ffffffffffffffff8082111561137957600080fd5b908401906060828703121561138d57600080fd5b61139561129f565b8235828111156113a457600080fd5b83019150601f820187136113b757600080fd5b6113c6878335602085016112c8565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113fb57600080fd5b813567ffffffffffffffff81111561141257600080fd5b8201601f8101841361142357600080fd5b611432848235602084016112c8565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261146657600080fd5b9190910192915050565b803561147b81611203565b919050565b60006020828403121561149257600080fd5b8151801515811461123557600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114cc576114cc6114a2565b5060010190565b6bffffffffffffffffffffffff811681146109e657600080fd5b8183526000602080850194508260005b8581101561154d57813561151081611203565b6001600160a01b0316875281830135611528816114d3565b6bffffffffffffffffffffffff168784015260409687019691909101906001016114fd565b509495945050505050565b803563ffffffff8116811461147b57600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561166c57878303603f190184528135368b9003609e190181126115b157600080fd5b8a0160a0813536839003601e190181126115ca57600080fd5b8201803567ffffffffffffffff8111156115e357600080fd5b8060061b36038413156115f557600080fd5b828752611607838801828c85016114ed565b92505050611616888301611470565b6001600160a01b03168886015281870135878601526060611638818401611558565b63ffffffff1690860152608061164f838201611558565b63ffffffff1695019490945250928501929085019060010161158b565b509098975050505050505050565b60006020828403121561168c57600080fd5b5051919050565b6000602082840312156116a557600080fd5b81516001600160c01b038116811461123557600080fd5b6000602082840312156116ce57600080fd5b815160ff8116811461123557600080fd5b600082198211156116f2576116f26114a2565b500190565b60006040828403121561170957600080fd5b6040516040810181811067ffffffffffffffff8211171561172c5761172c611289565b604052825161173a81611203565b8152602083015161174a816114d3565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117f4576020818501810151868301820152016117d8565b81811115611806576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261184560a08401826117ce565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061123560208301846117ce565b60008282101561188b5761188b6114a2565b500390565b600061ffff808316818114156118a8576118a86114a2565b600101939250505056fea2646970667358221220939ffbc2731160a3acf904aa502ec683052f07166493f46f16b3d34f23164d3664736f6c634300080c0033",
}

// ContractTradeAlgoServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractTradeAlgoServiceManagerMetaData.ABI instead.
var ContractTradeAlgoServiceManagerABI = ContractTradeAlgoServiceManagerMetaData.ABI

// ContractTradeAlgoServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractTradeAlgoServiceManagerMetaData.Bin instead.
var ContractTradeAlgoServiceManagerBin = ContractTradeAlgoServiceManagerMetaData.Bin

// DeployContractTradeAlgoServiceManager deploys a new Ethereum contract, binding an instance of ContractTradeAlgoServiceManager to it.
func DeployContractTradeAlgoServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _incredibleSquaringTaskManager common.Address) (common.Address, *types.Transaction, *ContractTradeAlgoServiceManager, error) {
	parsed, err := ContractTradeAlgoServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractTradeAlgoServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _incredibleSquaringTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractTradeAlgoServiceManager{ContractTradeAlgoServiceManagerCaller: ContractTradeAlgoServiceManagerCaller{contract: contract}, ContractTradeAlgoServiceManagerTransactor: ContractTradeAlgoServiceManagerTransactor{contract: contract}, ContractTradeAlgoServiceManagerFilterer: ContractTradeAlgoServiceManagerFilterer{contract: contract}}, nil
}

// ContractTradeAlgoServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManager struct {
	ContractTradeAlgoServiceManagerCaller     // Read-only binding to the contract
	ContractTradeAlgoServiceManagerTransactor // Write-only binding to the contract
	ContractTradeAlgoServiceManagerFilterer   // Log filterer for contract events
}

// ContractTradeAlgoServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTradeAlgoServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTradeAlgoServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTradeAlgoServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTradeAlgoServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTradeAlgoServiceManagerSession struct {
	Contract     *ContractTradeAlgoServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractTradeAlgoServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTradeAlgoServiceManagerCallerSession struct {
	Contract *ContractTradeAlgoServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// ContractTradeAlgoServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTradeAlgoServiceManagerTransactorSession struct {
	Contract     *ContractTradeAlgoServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// ContractTradeAlgoServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManagerRaw struct {
	Contract *ContractTradeAlgoServiceManager // Generic contract binding to access the raw methods on
}

// ContractTradeAlgoServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManagerCallerRaw struct {
	Contract *ContractTradeAlgoServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTradeAlgoServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTradeAlgoServiceManagerTransactorRaw struct {
	Contract *ContractTradeAlgoServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractTradeAlgoServiceManager creates a new instance of ContractTradeAlgoServiceManager, bound to a specific deployed contract.
func NewContractTradeAlgoServiceManager(address common.Address, backend bind.ContractBackend) (*ContractTradeAlgoServiceManager, error) {
	contract, err := bindContractTradeAlgoServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManager{ContractTradeAlgoServiceManagerCaller: ContractTradeAlgoServiceManagerCaller{contract: contract}, ContractTradeAlgoServiceManagerTransactor: ContractTradeAlgoServiceManagerTransactor{contract: contract}, ContractTradeAlgoServiceManagerFilterer: ContractTradeAlgoServiceManagerFilterer{contract: contract}}, nil
}

// NewContractTradeAlgoServiceManagerCaller creates a new read-only instance of ContractTradeAlgoServiceManager, bound to a specific deployed contract.
func NewContractTradeAlgoServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractTradeAlgoServiceManagerCaller, error) {
	contract, err := bindContractTradeAlgoServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManagerCaller{contract: contract}, nil
}

// NewContractTradeAlgoServiceManagerTransactor creates a new write-only instance of ContractTradeAlgoServiceManager, bound to a specific deployed contract.
func NewContractTradeAlgoServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTradeAlgoServiceManagerTransactor, error) {
	contract, err := bindContractTradeAlgoServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManagerTransactor{contract: contract}, nil
}

// NewContractTradeAlgoServiceManagerFilterer creates a new log filterer instance of ContractTradeAlgoServiceManager, bound to a specific deployed contract.
func NewContractTradeAlgoServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTradeAlgoServiceManagerFilterer, error) {
	contract, err := bindContractTradeAlgoServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManagerFilterer{contract: contract}, nil
}

// bindContractTradeAlgoServiceManager binds a generic wrapper to an already deployed contract.
func bindContractTradeAlgoServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractTradeAlgoServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTradeAlgoServiceManager.Contract.ContractTradeAlgoServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.ContractTradeAlgoServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.ContractTradeAlgoServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractTradeAlgoServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTradeAlgoServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.AvsDirectory(&_ContractTradeAlgoServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.AvsDirectory(&_ContractTradeAlgoServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractTradeAlgoServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractTradeAlgoServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractTradeAlgoServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractTradeAlgoServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.GetRestakeableStrategies(&_ContractTradeAlgoServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.GetRestakeableStrategies(&_ContractTradeAlgoServiceManager.CallOpts)
}

// TradeAlgoTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCaller) TradeAlgoTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTradeAlgoServiceManager.contract.Call(opts, &out, "incredibleSquaringTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeAlgoTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) TradeAlgoTaskManager() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.TradeAlgoTaskManager(&_ContractTradeAlgoServiceManager.CallOpts)
}

// TradeAlgoTaskManager is a free data retrieval call binding the contract method 0x77ef731d.
//
// Solidity: function incredibleSquaringTaskManager() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerSession) TradeAlgoTaskManager() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.TradeAlgoTaskManager(&_ContractTradeAlgoServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractTradeAlgoServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) Owner() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.Owner(&_ContractTradeAlgoServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractTradeAlgoServiceManager.Contract.Owner(&_ContractTradeAlgoServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractTradeAlgoServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractTradeAlgoServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.FreezeOperator(&_ContractTradeAlgoServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.FreezeOperator(&_ContractTradeAlgoServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.PayForRange(&_ContractTradeAlgoServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.PayForRange(&_ContractTradeAlgoServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.RegisterOperatorToAVS(&_ContractTradeAlgoServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.RegisterOperatorToAVS(&_ContractTradeAlgoServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.RenounceOwnership(&_ContractTradeAlgoServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.RenounceOwnership(&_ContractTradeAlgoServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.TransferOwnership(&_ContractTradeAlgoServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.TransferOwnership(&_ContractTradeAlgoServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.UpdateAVSMetadataURI(&_ContractTradeAlgoServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractTradeAlgoServiceManager.Contract.UpdateAVSMetadataURI(&_ContractTradeAlgoServiceManager.TransactOpts, _metadataURI)
}

// ContractTradeAlgoServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractTradeAlgoServiceManager contract.
type ContractTradeAlgoServiceManagerInitializedIterator struct {
	Event *ContractTradeAlgoServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractTradeAlgoServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTradeAlgoServiceManagerInitialized)
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
		it.Event = new(ContractTradeAlgoServiceManagerInitialized)
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
func (it *ContractTradeAlgoServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTradeAlgoServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTradeAlgoServiceManagerInitialized represents a Initialized event raised by the ContractTradeAlgoServiceManager contract.
type ContractTradeAlgoServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractTradeAlgoServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractTradeAlgoServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManagerInitializedIterator{contract: _ContractTradeAlgoServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractTradeAlgoServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractTradeAlgoServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTradeAlgoServiceManagerInitialized)
				if err := _ContractTradeAlgoServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractTradeAlgoServiceManagerInitialized, error) {
	event := new(ContractTradeAlgoServiceManagerInitialized)
	if err := _ContractTradeAlgoServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTradeAlgoServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractTradeAlgoServiceManager contract.
type ContractTradeAlgoServiceManagerOwnershipTransferredIterator struct {
	Event *ContractTradeAlgoServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractTradeAlgoServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTradeAlgoServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractTradeAlgoServiceManagerOwnershipTransferred)
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
func (it *ContractTradeAlgoServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTradeAlgoServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTradeAlgoServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractTradeAlgoServiceManager contract.
type ContractTradeAlgoServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractTradeAlgoServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTradeAlgoServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractTradeAlgoServiceManagerOwnershipTransferredIterator{contract: _ContractTradeAlgoServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractTradeAlgoServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractTradeAlgoServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTradeAlgoServiceManagerOwnershipTransferred)
				if err := _ContractTradeAlgoServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractTradeAlgoServiceManager *ContractTradeAlgoServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractTradeAlgoServiceManagerOwnershipTransferred, error) {
	event := new(ContractTradeAlgoServiceManagerOwnershipTransferred)
	if err := _ContractTradeAlgoServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
