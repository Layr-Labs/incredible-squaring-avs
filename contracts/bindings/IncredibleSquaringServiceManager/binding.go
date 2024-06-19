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

// ContractIncredibleSquaringServiceManagerMetaData contains all meta data concerning the ContractIncredibleSquaringServiceManager contract.
var ContractIncredibleSquaringServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_incredibleSquaringTaskManager\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fetchOperatorUrls\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"httpUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSquaringTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSquaringTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidOperator\",\"inputs\":[{\"name\":\"operatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorHttpUrls\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorRpcUrls\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorConsensusUrl\",\"inputs\":[{\"name\":\"httpUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorUrlRegistered\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"httpUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620023bc380380620023bc83398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e0516101005161213762000285600039600081816101ee0152610a1c01526000818161073b015281816108970152818161092e015281816112970152818161141b01526114ba015260008181610566015281816105f50152818161067501528181610ae201528181610cb601528181610d4c0152818161115b015281816111d501526113760152600081816103f601526104d401526000818161019601528181610d0a01528181610da8015261110201526121376000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806377ef731d116100a2578063a803be5711610071578063a803be5714610247578063a98fb35514610268578063e2d53d7c1461027b578063e481af9d1461029e578063f2fde38b146102a657600080fd5b806377ef731d146101e95780638da5cb5b146102105780639926ee7d14610221578063a364f4da1461023457600080fd5b806360707ef6116100de57806360707ef6146101745780636b3aa72e146101945780636fcb7a75146101ce578063715018a6146101e157600080fd5b80631b4455161461011057806333cfb7b71461012557806338c8ee641461014e5780634fa7ab3614610161575b600080fd5b61012361011e366004611848565b6102b9565b005b6101386101333660046118d2565b610541565b60405161014591906118f6565b60405180910390f35b61012361015c3660046118d2565b610a11565b61012361016f366004611a3b565b610aca565b6101876101823660046118d2565b610be4565b6040516101459190611aec565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b039091168152602001610145565b6101876101dc3660046118d2565b610c7e565b610123610c97565b6101b67f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166101b6565b61012361022f366004611aff565b610cab565b6101236102423660046118d2565b610d41565b61025a6102553660046118d2565b610e08565b604051610145929190611baa565b610123610276366004611bd8565b6110e3565b61028e6102893660046118d2565b611137565b6040519015158152602001610145565b6101386111cf565b6101236102b43660046118d2565b611599565b6102c161160f565b60005b818110156104bc578282828181106102de576102de611c15565b90506020028101906102f09190611c2b565b6103019060408101906020016118d2565b6001600160a01b03166323b872dd333086868681811061032357610323611c15565b90506020028101906103359190611c2b565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af115801561038c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b09190611c5b565b508282828181106103c3576103c3611c15565b90506020028101906103d59190611c2b565b6103e69060408101906020016118d2565b6001600160a01b031663095ea7b37f000000000000000000000000000000000000000000000000000000000000000085858581811061042757610427611c15565b90506020028101906104399190611c2b565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af1158015610487573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ab9190611c5b565b506104b581611c93565b90506102c4565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061050b9085908590600401611d47565b600060405180830381600087803b15801561052557600080fd5b505af1158015610539573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156105ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d19190611e55565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561063c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106609190611e6e565b90506001600160c01b03811615806106fa57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f59190611e97565b60ff16155b1561071657505060408051600081526020810190915292915050565b600061072a826001600160c01b0316611669565b90506000805b8251811015610800577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f584838151811061077a5761077a611c15565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156107be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e29190611e55565b6107ec9083611eba565b9150806107f881611c93565b915050610730565b5060008167ffffffffffffffff81111561081c5761081c611943565b604051908082528060200260200182016040528015610845578160200160208202803683370190505b5090506000805b8451811015610a0457600085828151811061086957610869611c15565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156108de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109029190611e55565b905060005b818110156109ee576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561097c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a09190611ed2565b600001518686815181106109b6576109b6611c15565b6001600160a01b0390921660209283029190910190910152846109d881611c93565b95505080806109e690611c93565b915050610907565b50505080806109fc90611c93565b91505061084c565b5090979650505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ac75760405162461bcd60e51b815260206004820152604a60248201527f6f6e6c79496e6372656469626c655371756172696e675461736b4d616e61676560448201527f723a206e6f742066726f6d206372656469626c65207371756172696e6720746160648201526939b59036b0b730b3b2b960b11b608482015260a4015b60405180910390fd5b50565b604051631619718360e21b81523360048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635865c60c906024016040805180830381865afa158015610b30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b549190611f13565b8051909150610b6257600080fd5b3360009081526097602090815260409091208451610b82928601906117af565b503360009081526098602090815260409091208351610ba3928501906117af565b507f100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f338484604051610bd793929190611f43565b60405180910390a1505050565b60976020526000908152604090208054610bfd90611f83565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2990611f83565b8015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b505050505081565b60986020526000908152604090208054610bfd90611f83565b610c9f61160f565b610ca9600061172c565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cf35760405162461bcd60e51b8152600401610abe90611fbe565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061050b9085908590600401612036565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d895760405162461bcd60e51b8152600401610abe90611fbe565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610ded57600080fd5b505af1158015610e01573d6000803e3d6000fd5b5050505050565b6001600160a01b0381166000908152609760205260408120805460609283929091610e3290611f83565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5e90611f83565b8015610eab5780601f10610e8057610100808354040283529160200191610eab565b820191906000526020600020905b815481529060010190602001808311610e8e57829003601f168201915b5050506001600160a01b038716600090815260986020526040812080549495509093909250610eda9150611f83565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0690611f83565b8015610f535780601f10610f2857610100808354040283529160200191610f53565b820191906000526020600020905b815481529060010190602001808311610f3657829003601f168201915b505050505090506000825111610f7b5760405162461bcd60e51b8152600401610abe90612081565b6000815111610f9c5760405162461bcd60e51b8152600401610abe90612081565b6001600160a01b0385166000908152609760209081526040808320609890925290912081548290610fcc90611f83565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff890611f83565b80156110455780601f1061101a57610100808354040283529160200191611045565b820191906000526020600020905b81548152906001019060200180831161102857829003601f168201915b5050505050915080805461105890611f83565b80601f016020809104026020016040519081016040528092919081815260200182805461108490611f83565b80156110d15780601f106110a6576101008083540402835291602001916110d1565b820191906000526020600020905b8154815290600101906020018083116110b457829003601f168201915b50505050509050935093505050915091565b6110eb61160f565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610dd3908490600401611aec565b604051631619718360e21b81526001600160a01b03828116600483015260009182917f00000000000000000000000000000000000000000000000000000000000000001690635865c60c906024016040805180830381865afa1580156111a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c59190611f13565b5115159392505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611231573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112559190611e97565b60ff1690508061127357505060408051600081526020810190915290565b6000805b8281101561132857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa1580156112e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130a9190611e55565b6113149083611eba565b91508061132081611c93565b915050611277565b5060008167ffffffffffffffff81111561134457611344611943565b60405190808252806020026020018201604052801561136d578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f69190611e97565b60ff1681101561158f57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa15801561146a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061148e9190611e55565b905060005b8181101561157a576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015611508573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152c9190611ed2565b6000015185858151811061154257611542611c15565b6001600160a01b03909216602092830291909101909101528361156481611c93565b945050808061157290611c93565b915050611493565b5050808061158790611c93565b915050611374565b5090949350505050565b6115a161160f565b6001600160a01b0381166116065760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610abe565b610ac78161172c565b6033546001600160a01b03163314610ca95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610abe565b60606000806116778461177e565b61ffff1667ffffffffffffffff81111561169357611693611943565b6040519080825280601f01601f1916602001820160405280156116bd576020820181803683370190505b5090506000805b8251821080156116d5575061010081105b1561158f576001811b93508584161561171c578060f81b8383815181106116fe576116fe611c15565b60200101906001600160f81b031916908160001a9053508160010191505b61172581611c93565b90506116c4565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b82156117a9576117936001846120c8565b90921691806117a1816120df565b915050611782565b92915050565b8280546117bb90611f83565b90600052602060002090601f0160209004810192826117dd5760008555611823565b82601f106117f657805160ff1916838001178555611823565b82800160010185558215611823579182015b82811115611823578251825591602001919060010190611808565b5061182f929150611833565b5090565b5b8082111561182f5760008155600101611834565b6000806020838503121561185b57600080fd5b823567ffffffffffffffff8082111561187357600080fd5b818501915085601f83011261188757600080fd5b81358181111561189657600080fd5b8660208260051b85010111156118ab57600080fd5b60209290920196919550909350505050565b6001600160a01b0381168114610ac757600080fd5b6000602082840312156118e457600080fd5b81356118ef816118bd565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156119375783516001600160a01b031683529284019291840191600101611912565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561197c5761197c611943565b60405290565b6040805190810167ffffffffffffffff8111828210171561197c5761197c611943565b600067ffffffffffffffff808411156119c0576119c0611943565b604051601f8501601f19908116603f011681019082821181831017156119e8576119e8611943565b81604052809350858152868686011115611a0157600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611a2c57600080fd5b6118ef838335602085016119a5565b60008060408385031215611a4e57600080fd5b823567ffffffffffffffff80821115611a6657600080fd5b611a7286838701611a1b565b93506020850135915080821115611a8857600080fd5b50611a9585828601611a1b565b9150509250929050565b6000815180845260005b81811015611ac557602081850181015186830182015201611aa9565b81811115611ad7576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006118ef6020830184611a9f565b60008060408385031215611b1257600080fd5b8235611b1d816118bd565b9150602083013567ffffffffffffffff80821115611b3a57600080fd5b9084019060608287031215611b4e57600080fd5b611b56611959565b823582811115611b6557600080fd5b83019150601f82018713611b7857600080fd5b611b87878335602085016119a5565b815260208301356020820152604083013560408201528093505050509250929050565b604081526000611bbd6040830185611a9f565b8281036020840152611bcf8185611a9f565b95945050505050565b600060208284031215611bea57600080fd5b813567ffffffffffffffff811115611c0157600080fd5b611c0d84828501611a1b565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e19833603018112611c4157600080fd5b9190910192915050565b8035611c56816118bd565b919050565b600060208284031215611c6d57600080fd5b815180151581146118ef57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611ca757611ca7611c7d565b5060010190565b6bffffffffffffffffffffffff81168114610ac757600080fd5b8183526000602080850194508260005b85811015611d28578135611ceb816118bd565b6001600160a01b0316875281830135611d0381611cae565b6bffffffffffffffffffffffff16878401526040968701969190910190600101611cd8565b509495945050505050565b803563ffffffff81168114611c5657600080fd5b60208082528181018390526000906040808401600586901b8501820187855b88811015611e4757878303603f190184528135368b9003609e19018112611d8c57600080fd5b8a0160a0813536839003601e19018112611da557600080fd5b8201803567ffffffffffffffff811115611dbe57600080fd5b8060061b3603841315611dd057600080fd5b828752611de2838801828c8501611cc8565b92505050611df1888301611c4b565b6001600160a01b03168886015281870135878601526060611e13818401611d33565b63ffffffff16908601526080611e2a838201611d33565b63ffffffff16950194909452509285019290850190600101611d66565b509098975050505050505050565b600060208284031215611e6757600080fd5b5051919050565b600060208284031215611e8057600080fd5b81516001600160c01b03811681146118ef57600080fd5b600060208284031215611ea957600080fd5b815160ff811681146118ef57600080fd5b60008219821115611ecd57611ecd611c7d565b500190565b600060408284031215611ee457600080fd5b611eec611982565b8251611ef7816118bd565b81526020830151611f0781611cae565b60208201529392505050565b600060408284031215611f2557600080fd5b611f2d611982565b82518152602083015160038110611f0757600080fd5b6001600160a01b0384168152606060208201819052600090611f6790830185611a9f565b8281036040840152611f798185611a9f565b9695505050505050565b600181811c90821680611f9757607f821691505b60208210811415611fb857634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b038316815260406020820152600082516060604084015261206060a0840182611a9f565b90506020840151606084015260408401516080840152809150509392505050565b60208082526027908201527f4e6f2075726c207265676973746572656420666f7220726571756573746564206040820152666164647265737360c81b606082015260800190565b6000828210156120da576120da611c7d565b500390565b600061ffff808316818114156120f7576120f7611c7d565b600101939250505056fea26469706673582212205563f855462a838e4bdbdae781fb11322b7877e1b1f210feb3d0fba3e0a9d68c64736f6c634300080c0033",
}

// ContractIncredibleSquaringServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.ABI instead.
var ContractIncredibleSquaringServiceManagerABI = ContractIncredibleSquaringServiceManagerMetaData.ABI

// ContractIncredibleSquaringServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringServiceManagerMetaData.Bin instead.
var ContractIncredibleSquaringServiceManagerBin = ContractIncredibleSquaringServiceManagerMetaData.Bin

// DeployContractIncredibleSquaringServiceManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringServiceManager to it.
func DeployContractIncredibleSquaringServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _incredibleSquaringTaskManager common.Address) (common.Address, *types.Transaction, *ContractIncredibleSquaringServiceManager, error) {
	parsed, err := ContractIncredibleSquaringServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _incredibleSquaringTaskManager)
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

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.AvsDirectory(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.AvsDirectory(&_ContractIncredibleSquaringServiceManager.CallOpts)
}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) FetchOperatorUrls(opts *bind.CallOpts, operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "fetchOperatorUrls", operatorAddress)

	outstruct := new(struct {
		HttpUrl string
		RpcUrl  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HttpUrl = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.RpcUrl = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) FetchOperatorUrls(operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FetchOperatorUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, operatorAddress)
}

// FetchOperatorUrls is a free data retrieval call binding the contract method 0xa803be57.
//
// Solidity: function fetchOperatorUrls(address operatorAddress) view returns(string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) FetchOperatorUrls(operatorAddress common.Address) (struct {
	HttpUrl string
	RpcUrl  string
}, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.FetchOperatorUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, operatorAddress)
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

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) IsValidOperator(opts *bind.CallOpts, operatorAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "isValidOperator", operatorAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IsValidOperator(&_ContractIncredibleSquaringServiceManager.CallOpts, operatorAddress)
}

// IsValidOperator is a free data retrieval call binding the contract method 0xe2d53d7c.
//
// Solidity: function isValidOperator(address operatorAddress) view returns(bool)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) IsValidOperator(operatorAddress common.Address) (bool, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.IsValidOperator(&_ContractIncredibleSquaringServiceManager.CallOpts, operatorAddress)
}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) OperatorHttpUrls(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "operatorHttpUrls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) OperatorHttpUrls(arg0 common.Address) (string, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.OperatorHttpUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, arg0)
}

// OperatorHttpUrls is a free data retrieval call binding the contract method 0x60707ef6.
//
// Solidity: function operatorHttpUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) OperatorHttpUrls(arg0 common.Address) (string, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.OperatorHttpUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, arg0)
}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCaller) OperatorRpcUrls(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringServiceManager.contract.Call(opts, &out, "operatorRpcUrls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) OperatorRpcUrls(arg0 common.Address) (string, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.OperatorRpcUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, arg0)
}

// OperatorRpcUrls is a free data retrieval call binding the contract method 0x6fcb7a75.
//
// Solidity: function operatorRpcUrls(address ) view returns(string)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerCallerSession) OperatorRpcUrls(arg0 common.Address) (string, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.OperatorRpcUrls(&_ContractIncredibleSquaringServiceManager.CallOpts, arg0)
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

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.PayForRange(&_ContractIncredibleSquaringServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.PayForRange(&_ContractIncredibleSquaringServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) RegisterOperatorConsensusUrl(opts *bind.TransactOpts, httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "registerOperatorConsensusUrl", httpUrl, rpcUrl)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) RegisterOperatorConsensusUrl(httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorConsensusUrl(&_ContractIncredibleSquaringServiceManager.TransactOpts, httpUrl, rpcUrl)
}

// RegisterOperatorConsensusUrl is a paid mutator transaction binding the contract method 0x4fa7ab36.
//
// Solidity: function registerOperatorConsensusUrl(string httpUrl, string rpcUrl) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) RegisterOperatorConsensusUrl(httpUrl string, rpcUrl string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.RegisterOperatorConsensusUrl(&_ContractIncredibleSquaringServiceManager.TransactOpts, httpUrl, rpcUrl)
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

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSquaringServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSquaringServiceManager.TransactOpts, _metadataURI)
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

// ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator is returned from FilterOperatorUrlRegistered and is used to iterate over the raw logs and unpacked data for OperatorUrlRegistered events raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator struct {
	Event *ContractIncredibleSquaringServiceManagerOperatorUrlRegistered // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringServiceManagerOperatorUrlRegistered)
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
		it.Event = new(ContractIncredibleSquaringServiceManagerOperatorUrlRegistered)
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
func (it *ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringServiceManagerOperatorUrlRegistered represents a OperatorUrlRegistered event raised by the ContractIncredibleSquaringServiceManager contract.
type ContractIncredibleSquaringServiceManagerOperatorUrlRegistered struct {
	OperatorId common.Address
	HttpUrl    string
	RpcUrl     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorUrlRegistered is a free log retrieval operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) FilterOperatorUrlRegistered(opts *bind.FilterOpts) (*ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.FilterLogs(opts, "OperatorUrlRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringServiceManagerOperatorUrlRegisteredIterator{contract: _ContractIncredibleSquaringServiceManager.contract, event: "OperatorUrlRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUrlRegistered is a free log subscription operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) WatchOperatorUrlRegistered(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringServiceManagerOperatorUrlRegistered) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringServiceManager.contract.WatchLogs(opts, "OperatorUrlRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringServiceManagerOperatorUrlRegistered)
				if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OperatorUrlRegistered", log); err != nil {
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

// ParseOperatorUrlRegistered is a log parse operation binding the contract event 0x100388966d784f0d0a7df3e4926797840162cc1e2cf828c6ca6ec6adaeb4640f.
//
// Solidity: event OperatorUrlRegistered(address operatorId, string httpUrl, string rpcUrl)
func (_ContractIncredibleSquaringServiceManager *ContractIncredibleSquaringServiceManagerFilterer) ParseOperatorUrlRegistered(log types.Log) (*ContractIncredibleSquaringServiceManagerOperatorUrlRegistered, error) {
	event := new(ContractIncredibleSquaringServiceManagerOperatorUrlRegistered)
	if err := _ContractIncredibleSquaringServiceManager.contract.UnpackLog(event, "OperatorUrlRegistered", log); err != nil {
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
