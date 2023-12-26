// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSquaringTaskManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSquaringTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTask struct {
	NumberToBeSquared         *big.Int
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSquaringTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

// IIncredibleSquaringTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSquaringTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSquaringTaskManagerMetaData contains all meta data concerning the ContractIncredibleSquaringTaskManager contract.
var ContractIncredibleSquaringTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"NewTaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"StaleStakesForbiddenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedSuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedUnsuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"numberSquared\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"}],\"name\":\"TaskResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"createNewTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getCheckSignaturesIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskResponseWindowBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_generator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTaskNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"numberSquared\",\"type\":\"uint256\"}],\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\"}],\"name\":\"raiseAndResolveChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberToBeSquared\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"numberSquared\",\"type\":\"uint256\"}],\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"respondToTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setStaleStakesForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleStakesForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"taskSuccesfullyChallenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005976380380620059768339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e0516101005161567f620002f7600039600081816102720152818161057b015261177c01526000818161054401526127e70152600081816103f301528181611ea301526129d101526000818161041a01528181612ba70152612d6901526000818161046701528181610ccd015281816124e40152818161265d015261288b015261567f6000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636b532e9e116101255780638da5cb5b116100ad578063f2fde38b1161007c578063f2fde38b14610566578063f5c9899d14610579578063f63c5bab1461059f578063f8c8765e146105a7578063fabc1cbc146105ba57600080fd5b80638da5cb5b14610500578063b98d090814610511578063cefdc1d41461051e578063df5cf7231461053f57600080fd5b8063715018a6116100f4578063715018a6146104aa57806372d18e8d146104b25780637afa1eed146104c5578063886f1195146104d85780638b00ce7c146104eb57600080fd5b80636b532e9e1461043c5780636b92787e1461044f5780636d14a987146104625780636efb46361461048957600080fd5b8063416c7e5e116101a85780635baec9a0116101775780635baec9a0146103b05780635c975abb146103c35780635decc3f5146103cb5780635df45946146103ee578063683048351461041557600080fd5b8063416c7e5e146103425780634f739f7414610355578063595c6a67146103755780635ac86ab71461037d57600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b61022961022436600461427b565b6105cd565b005b610229610239366004614298565b610689565b61025161024c366004614416565b6107c8565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b609b546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e2366004614484565b60996020526000908152604090205481565b604051908152602001610264565b6102f4610310366004614484565b60986020526000908152604090205481565b6103356103303660046144a1565b610952565b60405161026491906145e8565b610229610350366004614610565b610ccb565b610368610363366004614675565b610e40565b6040516102649190614779565b6102296114c4565b6103a061038b366004614843565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6102296103be366004614b23565b61158b565b6066546102f4565b6103a06103d9366004614484565b609a6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61022961044a366004614b97565b611a0a565b61022961045d366004614c1d565b611fdc565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61049c610497366004614c78565b612192565b604051610264929190614d38565b61022961301e565b609754610100900463ffffffff16610294565b609c546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60975461029490610100900463ffffffff1681565b6033546001600160a01b03166102bc565b6097546103a09060ff1681565b61053161052c366004614d81565b613032565b604051610264929190614dc3565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61022961057436600461427b565b6131c4565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b5366004614de4565b61323a565b6102296105c8366004614298565b61338b565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610620573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106449190614e40565b6001600160a01b0316336001600160a01b03161461067d5760405162461bcd60e51b815260040161067490614e5d565b60405180910390fd5b610686816134e7565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f59190614ea7565b6107115760405162461bcd60e51b815260040161067490614ec4565b6066548181161461078a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081057610810614f0c565b60200201518951600160200201518a6020015160006002811061083557610835614f0c565b60200201518b6020015160016002811061085157610851614f0c565b602090810291909101518c518d8301516040516108ae9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d19190614f22565b90506109446108ea6108e388846135de565b8690613675565b6108f2613709565b61093a61092b85610925604080518082018252600080825260209182015281518083019092526001825260029082015290565b906135de565b6109348c6137c9565b90613675565b886201d4c0613859565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610994573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b89190614e40565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1e9190614e40565b9050600085516001600160401b03811115610a3b57610a3b6142b1565b604051908082528060200260200182016040528015610a6e57816020015b6060815260200190600190039081610a595790505b50905060005b8651811015610cc0576000878281518110610a9157610a91614f0c565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b03861690638902624590604401600060405180830381865afa158015610af2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b1a9190810190614f44565b905080516001600160401b03811115610b3557610b356142b1565b604051908082528060200260200182016040528015610b7a57816020015b6040805180820190915260008082526020820152815260200190600190039081610b535790505b50848481518110610b8d57610b8d614f0c565b602002602001018190525060005b8151811015610caa576000828281518110610bb857610bb8614f0c565b602090810291909101810151604080518082018252828152905163fa28c62760e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b169063fa28c62790606401602060405180830381865afa158015610c30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c549190614fd4565b6001600160601b0316815250868681518110610c7257610c72614f0c565b60200260200101518381518110610c8b57610c8b614f0c565b6020026020010181905250508080610ca290615013565b915050610b9b565b5050508080610cb890615013565b915050610a74565b509695505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4d9190614e40565b6001600160a01b0316336001600160a01b031614610df95760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610674565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610e6b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf9190614e40565b9050610efc6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90610f2c908b908990899060040161502e565b600060405180830381865afa158015610f49573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f719190810190615078565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290610fa3908b908b908b9060040161512f565b600060405180830381865afa158015610fc0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fe89190810190615078565b6040820152856001600160401b03811115611005576110056142b1565b60405190808252806020026020018201604052801561103857816020015b60608152602001906001900390816110235790505b50606082015260005b60ff81168711156113d5576000856001600160401b03811115611066576110666142b1565b60405190808252806020026020018201604052801561108f578160200160208202803683370190505b5083606001518360ff16815181106110a9576110a9614f0c565b602002602001018190525060005b868110156112d55760008c6001600160a01b03166304ec63518a8a858181106110e2576110e2614f0c565b905060200201358e8860000151868151811061110057611100614f0c565b60200260200101516040518463ffffffff1660e01b815260040161113d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561115a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117e9190615158565b90508a8a8560ff1681811061119557611195614f0c565b6001600160c01b03841692013560f81c9190911c6001908116141590506112c257856001600160a01b031663dd9846b98a8a858181106111d7576111d7614f0c565b905060200201358d8d8860ff168181106111f3576111f3614f0c565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126d9190615181565b85606001518560ff168151811061128657611286614f0c565b6020026020010151848151811061129f5761129f614f0c565b63ffffffff90921660209283029190910190910152826112be81615013565b9350505b50806112cd81615013565b9150506110b7565b506000816001600160401b038111156112f0576112f06142b1565b604051908082528060200260200182016040528015611319578160200160208202803683370190505b50905060005b8281101561139a5784606001518460ff168151811061134057611340614f0c565b6020026020010151818151811061135957611359614f0c565b602002602001015182828151811061137357611373614f0c565b63ffffffff909216602092830291909101909101528061139281615013565b91505061131f565b508084606001518460ff16815181106113b5576113b5614f0c565b6020026020010181905250505080806113cd9061519e565b915050611041565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611416573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143a9190614e40565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061146d908b908b908e906004016151be565b600060405180830381865afa15801561148a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114b29190810190615078565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561150c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115309190614ea7565b61154c5760405162461bcd60e51b815260040161067490614ec4565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b609b546001600160a01b031633146115e55760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610674565b60006115f76040850160208601614484565b905036600061160960408701876151e8565b909250905060006116206080880160608901614484565b9050609860006116336020890189614484565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161165f919061522e565b60405160208183030381529060405280519060200120146116e85760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b60006099816116fa60208a018a614484565b63ffffffff1663ffffffff16815260200190815260200160002054146117775760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610674565b6117a17f0000000000000000000000000000000000000000000000000000000000000000856152cf565b63ffffffff164363ffffffff1611156118125760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610674565b6000866040516020016118259190615315565b60405160208183030381529060405280519060200120905060008061184d8387878a8c612192565b9150915060005b8581101561194c578460ff168360200151828151811061187657611876614f0c565b60200260200101516118889190615323565b6001600160601b03166064846000015183815181106118a9576118a9614f0c565b60200260200101516001600160601b03166118c49190615352565b101561193a576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610674565b8061194481615013565b915050611854565b5060408051808201825263ffffffff43168152602080820184905291519091611979918c91849101615371565b60405160208183030381529060405280519060200120609960008c60000160208101906119a69190614484565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a826040516119f5929190615371565b60405180910390a15050505050505050505050565b6000611a196020850185614484565b63ffffffff8116600090815260996020526040902054909150853590611a8b5760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610674565b8484604051602001611a9e92919061539d565b60408051601f19818403018152918152815160209283012063ffffffff85166000908152609990935291205414611b3d5760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b63ffffffff82166000908152609a602052604090205460ff1615611bd55760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610674565b6064611be46020860186614484565b611bee91906152cf565b63ffffffff164363ffffffff161115611c6f5760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610674565b6000611c7b8280615352565b9050602086013581146001811415611cc957604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350505050611fd6565b600085516001600160401b03811115611ce457611ce46142b1565b604051908082528060200260200182016040528015611d0d578160200160208202803683370190505b50905060005b8651811015611d7f57611d50878281518110611d3157611d31614f0c565b6020026020010151805160009081526020918201519091526040902090565b828281518110611d6257611d62614f0c565b602090810291909101015280611d7781615013565b915050611d13565b506000611d9260408b0160208c01614484565b82604051602001611da49291906153d3565b60405160208183030381529060405280519060200120905087602001358114611e4e5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610674565b600087516001600160401b03811115611e6957611e696142b1565b604051908082528060200260200182016040528015611e92578160200160208202803683370190505b50905060005b8851811015611f85577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110611ee257611ee2614f0c565b60200260200101516040518263ffffffff1660e01b8152600401611f0891815260200190565b602060405180830381865afa158015611f25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f499190614e40565b828281518110611f5b57611f5b614f0c565b6001600160a01b039092166020928302919091019091015280611f7d81615013565b915050611e98565b5063ffffffff87166000818152609a6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505b50505050565b609c546001600160a01b031633146120405760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610674565b612077604051806080016040528060008152602001600063ffffffff16815260200160608152602001600063ffffffff1681525090565b84815263ffffffff438116602080840191909152908516606083015260408051601f8501839004830281018301909152838152908490849081908401838280828437600092019190915250505050604080830191909152516120dd908290602001615447565b60408051601f1981840301815282825280516020918201206097805463ffffffff610100918290048116600090815260989095529490932091909155540416907f1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c489061214a908490615447565b60405180910390a260975461216b90610100900463ffffffff1660016152cf565b609760016101000a81548163ffffffff021916908363ffffffff1602179055505050505050565b604080518082018252606080825260208201529082015151600090851480156121bf575060a08301515185145b80156121cf575060c08301515185145b80156121df575060e08301515185145b6122495760405162461bcd60e51b8152602060048201526041602482015260008051602061562a83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610674565b825151602084015151146122c15760405162461bcd60e51b81526020600482015260446024820181905260008051602061562a833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610674565b4363ffffffff168463ffffffff1611156123315760405162461bcd60e51b815260206004820152603c602482015260008051602061562a83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610674565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115612372576123726142b1565b60405190808252806020026020018201604052801561239b578160200160208202803683370190505b506020820152866001600160401b038111156123b9576123b96142b1565b6040519080825280602002602001820160405280156123e2578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115612416576124166142b1565b60405190808252806020026020018201604052801561243f578160200160208202803683370190505b5081526020860151516001600160401b0381111561245f5761245f6142b1565b604051908082528060200260200182016040528015612488578160200160208202803683370190505b508160200181905250600061255a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612531573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061255591906154ac565b613a7d565b905060005b8760200151518110156127d65761258588602001518281518110611d3157611d31614f0c565b8360200151828151811061259b5761259b614f0c565b6020908102919091010152801561265b5760208301516125bc6001836154c9565b815181106125cc576125cc614f0c565b602002602001015160001c836020015182815181106125ed576125ed614f0c565b602002602001015160001c1161265b576040805162461bcd60e51b815260206004820152602481019190915260008051602061562a83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec6351846020015183815181106126a0576126a0614f0c565b60200260200101518b8b6000015185815181106126bf576126bf614f0c565b60200260200101516040518463ffffffff1660e01b81526004016126fc9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015612719573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061273d9190615158565b6001600160c01b03168360000151828151811061275c5761275c614f0c565b6020026020010181815250506127c26108e3612796848660000151858151811061278857612788614f0c565b602002602001015116613b38565b8a6020015184815181106127ac576127ac614f0c565b6020026020010151613b6390919063ffffffff16565b9450806127ce81615013565b91505061255f565b50506127e183613c47565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612843573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061286791906154e0565b60975490915060ff1660005b8a811015612eed5781156129cf578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106128ca576128ca614f0c565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa15801561290a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061292e91906154e0565b61293891906154f9565b10156129cf5760405162461bcd60e51b8152602060048201526066602482015260008051602061562a83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612a1057612a10614f0c565b9050013560f81c60f81b60f81c8c8c60a001518581518110612a3457612a34614f0c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612a90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ab49190615511565b6001600160401b031916612ad78a604001518381518110611d3157611d31614f0c565b67ffffffffffffffff191614612b735760405162461bcd60e51b8152602060048201526061602482015260008051602061562a83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610674565b612ba389604001518281518110612b8c57612b8c614f0c565b60200260200101518761367590919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612be657612be6614f0c565b9050013560f81c60f81b60f81c8c8c60c001518581518110612c0a57612c0a614f0c565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612c66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c8a9190614fd4565b85602001518281518110612ca057612ca0614f0c565b6001600160601b03909216602092830291909101820152850151805182908110612ccc57612ccc614f0c565b602002602001015185600001518281518110612cea57612cea614f0c565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015612ed857612d6286600001518281518110612d3457612d34614f0c565b60200260200101518f8f86818110612d4e57612d4e614f0c565b600192013560f81c9290921c811614919050565b15612ec6577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f86818110612da857612da8614f0c565b9050013560f81c60f81b60f81c8e89602001518581518110612dcc57612dcc614f0c565b60200260200101518f60e001518881518110612dea57612dea614f0c565b60200260200101518781518110612e0357612e03614f0c565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e8b9190614fd4565b8751805185908110612e9f57612e9f614f0c565b60200260200101818151612eb3919061553c565b6001600160601b03169052506001909101905b80612ed081615013565b915050612d0e565b50508080612ee590615013565b915050612873565b505050600080612f078c868a606001518b608001516107c8565b9150915081612f785760405162461bcd60e51b8152602060048201526043602482015260008051602061562a83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610674565b80612fd95760405162461bcd60e51b8152602060048201526039602482015260008051602061562a83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610674565b50506000878260200151604051602001612ff49291906153d3565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b613026613ce2565b6130306000613d3c565b565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061306d5761306d614f0c565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906130a99088908690600401615564565b600060405180830381865afa1580156130c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526130ee9190810190615078565b60008151811061310057613100614f0c565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561316c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131909190615158565b6001600160c01b0316905060006131a682613d8e565b9050816131b48a838a610952565b9550955050505050935093915050565b6131cc613ce2565b6001600160a01b0381166132315760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610674565b61068681613d3c565b600054610100900460ff161580801561325a5750600054600160ff909116105b806132745750303b158015613274575060005460ff166001145b6132d75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610674565b6000805460ff1916600117905580156132fa576000805461ff0019166101001790555b613305856000613deb565b61330e84613d3c565b609b80546001600160a01b038086166001600160a01b031992831617909255609c8054928516929091169190911790558015613384576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134029190614e40565b6001600160a01b0316336001600160a01b0316146134325760405162461bcd60e51b815260040161067490614e5d565b6066541981196066541916146134b05760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107bd565b6001600160a01b0381166135755760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610674565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b60408051808201909152600080825260208201526135fa61418c565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561362d5761362f565bfe5b508061366d5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610674565b505092915050565b60408051808201909152600080825260208201526136916141aa565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561362d57508061366d5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610674565b6137116141c8565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820190915260008082526020820152600080806137f960008051602061560a83398151915286614f22565b90505b61380581613ed5565b909350915060008051602061560a83398151915282830983141561383f576040805180820190915290815260208101919091529392505050565b60008051602061560a8339815191526001820890506137fc565b60408051808201825286815260208082018690528251808401909352868352820184905260009182919061388b6141ed565b60005b6002811015613a505760006138a4826006615352565b90508482600281106138b8576138b8614f0c565b602002015151836138ca8360006154f9565b600c81106138da576138da614f0c565b60200201528482600281106138f1576138f1614f0c565b6020020151602001518382600161390891906154f9565b600c811061391857613918614f0c565b602002015283826002811061392f5761392f614f0c565b60200201515151836139428360026154f9565b600c811061395257613952614f0c565b602002015283826002811061396957613969614f0c565b60200201515160016020020151836139828360036154f9565b600c811061399257613992614f0c565b60200201528382600281106139a9576139a9614f0c565b6020020151602001516000600281106139c4576139c4614f0c565b6020020151836139d58360046154f9565b600c81106139e5576139e5614f0c565b60200201528382600281106139fc576139fc614f0c565b602002015160200151600160028110613a1757613a17614f0c565b602002015183613a288360056154f9565b600c8110613a3857613a38614f0c565b60200201525080613a4881615013565b91505061388e565b50613a5961420c565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613a8984613f57565b90508015613b2f578260ff168460018651613aa491906154c9565b81518110613ab457613ab4614f0c565b016020015160f81c10613b2f5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610674565b90505b92915050565b6000805b8215613b3257613b4d6001846154c9565b9092169180613b5b816155b8565b915050613b3c565b60408051808201909152600080825260208201526102008261ffff1610613bbf5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610674565b8161ffff1660011415613bd3575081613b32565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613c3c57600161ffff871660ff83161c81161415613c1f57613c1c8484613675565b93505b613c298384613675565b92506201fffe600192831b169101613bef565b509195945050505050565b60408051808201909152600080825260208201528151158015613c6c57506020820151155b15613c8a575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061560a8339815191528460200151613cbd9190614f22565b613cd59060008051602061560a8339815191526154c9565b905292915050565b919050565b6033546001600160a01b031633146130305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610674565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613de4576001811b915083821615613dd457828160f81b604051602001613dc29291906155da565b60405160208183030381529060405292505b613ddd81615013565b9050613d94565b5050919050565b6065546001600160a01b0316158015613e0c57506001600160a01b03821615155b613e8e5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613ed1826134e7565b5050565b6000808060008051602061560a833981519152600360008051602061560a8339815191528660008051602061560a833981519152888909090890506000613f4b827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061560a8339815191526140e4565b91959194509092505050565b600061010082511115613fe05760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610674565b8151613fee57506000919050565b6000808360008151811061400457614004614f0c565b0160200151600160f89190911c81901b92505b84518110156140db5784818151811061403257614032614f0c565b0160200151600160f89190911c1b91508282116140c75760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610674565b918117916140d481615013565b9050614017565b50909392505050565b6000806140ef61420c565b6140f761422a565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561362d5750826141815760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610674565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806141db614248565b81526020016141e8614248565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068657600080fd5b60006020828403121561428d57600080fd5b8135613b2f81614266565b6000602082840312156142aa57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142e9576142e96142b1565b60405290565b60405161010081016001600160401b03811182821017156142e9576142e96142b1565b604051601f8201601f191681016001600160401b038111828210171561433a5761433a6142b1565b604052919050565b60006040828403121561435457600080fd5b61435c6142c7565b9050813581526020820135602082015292915050565b600082601f83011261438357600080fd5b604051604081018181106001600160401b03821117156143a5576143a56142b1565b80604052508060408401858111156143bc57600080fd5b845b81811015613c3c5780358352602092830192016143be565b6000608082840312156143e857600080fd5b6143f06142c7565b90506143fc8383614372565b815261440b8360408401614372565b602082015292915050565b600080600080610120858703121561442d57600080fd5b8435935061443e8660208701614342565b925061444d86606087016143d6565b915061445c8660e08701614342565b905092959194509250565b63ffffffff8116811461068657600080fd5b8035613cdd81614467565b60006020828403121561449657600080fd5b8135613b2f81614467565b6000806000606084860312156144b657600080fd5b83356144c181614266565b92506020848101356001600160401b03808211156144de57600080fd5b818701915087601f8301126144f257600080fd5b813581811115614504576145046142b1565b614516601f8201601f19168501614312565b9150808252888482850101111561452c57600080fd5b808484018584013760008482840101525080945050505061454f60408501614479565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156145da578385038a52825180518087529087019087870190845b818110156145c5578351805184528a01516001600160601b03168a84015292890192604090920191600101614595565b50509a87019a95505091850191600101614577565b509298975050505050505050565b6020815260006145fb6020830184614558565b9392505050565b801515811461068657600080fd5b60006020828403121561462257600080fd5b8135613b2f81614602565b60008083601f84011261463f57600080fd5b5081356001600160401b0381111561465657600080fd5b60208301915083602082850101111561466e57600080fd5b9250929050565b6000806000806000806080878903121561468e57600080fd5b863561469981614266565b955060208701356146a981614467565b945060408701356001600160401b03808211156146c557600080fd5b6146d18a838b0161462d565b909650945060608901359150808211156146ea57600080fd5b818901915089601f8301126146fe57600080fd5b81358181111561470d57600080fd5b8a60208260051b850101111561472257600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561476e57815163ffffffff168752958201959082019060010161474c565b509495945050505050565b60006020808352835160808285015261479560a0850182614738565b905081850151601f19808684030160408701526147b28383614738565b925060408701519150808684030160608701526147cf8383614738565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156148265784878303018452614814828751614738565b958801959388019391506001016147fa565b509998505050505050505050565b60ff8116811461068657600080fd5b60006020828403121561485557600080fd5b8135613b2f81614834565b60006080828403121561487257600080fd5b50919050565b60006040828403121561487257600080fd5b60006001600160401b038211156148a3576148a36142b1565b5060051b60200190565b600082601f8301126148be57600080fd5b813560206148d36148ce8361488a565b614312565b82815260059290921b840181019181810190868411156148f257600080fd5b8286015b84811015610cc057803561490981614467565b83529183019183016148f6565b600082601f83011261492757600080fd5b813560206149376148ce8361488a565b82815260069290921b8401810191818101908684111561495657600080fd5b8286015b84811015610cc05761496c8882614342565b83529183019160400161495a565b600082601f83011261498b57600080fd5b8135602061499b6148ce8361488a565b82815260059290921b840181019181810190868411156149ba57600080fd5b8286015b84811015610cc05780356001600160401b038111156149dd5760008081fd5b6149eb8986838b01016148ad565b8452509183019183016149be565b60006101808284031215614a0c57600080fd5b614a146142ef565b905081356001600160401b0380821115614a2d57600080fd5b614a39858386016148ad565b83526020840135915080821115614a4f57600080fd5b614a5b85838601614916565b60208401526040840135915080821115614a7457600080fd5b614a8085838601614916565b6040840152614a9285606086016143d6565b6060840152614aa48560e08601614342565b6080840152610120840135915080821115614abe57600080fd5b614aca858386016148ad565b60a0840152610140840135915080821115614ae457600080fd5b614af0858386016148ad565b60c0840152610160840135915080821115614b0a57600080fd5b50614b178482850161497a565b60e08301525092915050565b600080600060808486031215614b3857600080fd5b83356001600160401b0380821115614b4f57600080fd5b614b5b87838801614860565b9450614b6a8760208801614878565b93506060860135915080821115614b8057600080fd5b50614b8d868287016149f9565b9150509250925092565b60008060008060c08587031215614bad57600080fd5b84356001600160401b0380821115614bc457600080fd5b614bd088838901614860565b9550614bdf8860208901614878565b9450614bee8860608901614878565b935060a0870135915080821115614c0457600080fd5b50614c1187828801614916565b91505092959194509250565b60008060008060608587031215614c3357600080fd5b843593506020850135614c4581614467565b925060408501356001600160401b03811115614c6057600080fd5b614c6c8782880161462d565b95989497509550505050565b600080600080600060808688031215614c9057600080fd5b8535945060208601356001600160401b0380821115614cae57600080fd5b614cba89838a0161462d565b909650945060408801359150614ccf82614467565b90925060608701359080821115614ce557600080fd5b50614cf2888289016149f9565b9150509295509295909350565b600081518084526020808501945080840160005b8381101561476e5781516001600160601b031687529582019590820190600101614d13565b6040815260008351604080840152614d536080840182614cff565b90506020850151603f19848303016060850152614d708282614cff565b925050508260208301529392505050565b600080600060608486031215614d9657600080fd5b8335614da181614266565b9250602084013591506040840135614db881614467565b809150509250925092565b828152604060208201526000614ddc6040830184614558565b949350505050565b60008060008060808587031215614dfa57600080fd5b8435614e0581614266565b93506020850135614e1581614266565b92506040850135614e2581614266565b91506060850135614e3581614266565b939692955090935050565b600060208284031215614e5257600080fd5b8151613b2f81614266565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614eb957600080fd5b8151613b2f81614602565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614f3f57634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614f5757600080fd5b82516001600160401b03811115614f6d57600080fd5b8301601f81018513614f7e57600080fd5b8051614f8c6148ce8261488a565b81815260059190911b82018301908381019087831115614fab57600080fd5b928401925b82841015614fc957835182529284019290840190614fb0565b979650505050505050565b600060208284031215614fe657600080fd5b81516001600160601b0381168114613b2f57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561502757615027614ffd565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561505b57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561508b57600080fd5b82516001600160401b038111156150a157600080fd5b8301601f810185136150b257600080fd5b80516150c06148ce8261488a565b81815260059190911b820183019083810190878311156150df57600080fd5b928401925b82841015614fc95783516150f781614467565b825292840192908401906150e4565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff8416815260406020820152600061514f604083018486615106565b95945050505050565b60006020828403121561516a57600080fd5b81516001600160c01b0381168114613b2f57600080fd5b60006020828403121561519357600080fd5b8151613b2f81614467565b600060ff821660ff8114156151b5576151b5614ffd565b60010192915050565b6040815260006151d2604083018587615106565b905063ffffffff83166020830152949350505050565b6000808335601e198436030181126151ff57600080fd5b8301803591506001600160401b0382111561521957600080fd5b60200191503681900382131561466e57600080fd5b60208152813560208201526000602083013561524981614467565b63ffffffff81166040840152506040830135601e1984360301811261526d57600080fd5b830180356001600160401b0381111561528557600080fd5b80360385131561529457600080fd5b608060608501526152ac60a085018260208501615106565b9150506152bb60608501614479565b63ffffffff81166080850152509392505050565b600063ffffffff8083168185168083038211156152ee576152ee614ffd565b01949350505050565b803561530281614467565b63ffffffff168252602090810135910152565b60408101613b3282846152f7565b60006001600160601b038083168185168183048111821515161561534957615349614ffd565b02949350505050565b600081600019048311821515161561536c5761536c614ffd565b500290565b6080810161537f82856152f7565b63ffffffff8351166040830152602083015160608301529392505050565b608081016153ab82856152f7565b82356153b681614467565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561540e578151855293820193908201906001016153f2565b5092979650505050505050565b60005b8381101561543657818101518382015260200161541e565b83811115611fd65750506000910152565b60208152815160208201526000602083015163ffffffff8082166040850152604085015191506080606085015281518060a086015261548d8160c087016020860161541b565b60609590950151166080840152505060c0601f909201601f1916010190565b6000602082840312156154be57600080fd5b8151613b2f81614834565b6000828210156154db576154db614ffd565b500390565b6000602082840312156154f257600080fd5b5051919050565b6000821982111561550c5761550c614ffd565b500190565b60006020828403121561552357600080fd5b815167ffffffffffffffff1981168114613b2f57600080fd5b60006001600160601b038381169083168181101561555c5761555c614ffd565b039392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156155ab5784518352938301939183019160010161558f565b5090979650505050505050565b600061ffff808316818114156155d0576155d0614ffd565b6001019392505050565b600083516155ec81846020880161541b565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a26469706673582212201b205cefea0273eaf432dacdab393a02db0139437c7cfc88dc32c4685e46309f64736f6c634300080c0033",
}

// ContractIncredibleSquaringTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.ABI instead.
var ContractIncredibleSquaringTaskManagerABI = ContractIncredibleSquaringTaskManagerMetaData.ABI

// ContractIncredibleSquaringTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSquaringTaskManagerMetaData.Bin instead.
var ContractIncredibleSquaringTaskManagerBin = ContractIncredibleSquaringTaskManagerMetaData.Bin

// DeployContractIncredibleSquaringTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSquaringTaskManager to it.
func DeployContractIncredibleSquaringTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSquaringTaskManager, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSquaringTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSquaringTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManager struct {
	ContractIncredibleSquaringTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSquaringTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSquaringTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSquaringTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSquaringTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSquaringTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSquaringTaskManagerSession struct {
	Contract     *ContractIncredibleSquaringTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSquaringTaskManagerCallerSession struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSquaringTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSquaringTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// ContractIncredibleSquaringTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerRaw struct {
	Contract *ContractIncredibleSquaringTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSquaringTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSquaringTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSquaringTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSquaringTaskManager creates a new instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSquaringTaskManager, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManager{ContractIncredibleSquaringTaskManagerCaller: ContractIncredibleSquaringTaskManagerCaller{contract: contract}, ContractIncredibleSquaringTaskManagerTransactor: ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, ContractIncredibleSquaringTaskManagerFilterer: ContractIncredibleSquaringTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSquaringTaskManagerCaller creates a new read-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSquaringTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerTransactor creates a new write-only instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSquaringTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSquaringTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSquaringTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSquaringTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSquaringTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSquaringTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSquaringTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSquaringTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.ContractIncredibleSquaringTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Aggregator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CheckSignatures(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Delegation(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Generator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Owner(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused(&_ContractIncredibleSquaringTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Paused0(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauserRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StakeRegistry(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskNumber(&_ContractIncredibleSquaringTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSquaringTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "createNewTask", numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0x6b92787e.
//
// Solidity: function createNewTask(uint256 numberToBeSquared, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) CreateNewTask(numberToBeSquared *big.Int, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.CreateNewTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, numberToBeSquared, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Initialize(&_ContractIncredibleSquaringTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Pause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.PauseAll(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x6b532e9e.
//
// Solidity: function raiseAndResolveChallenge((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, taskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0x5baec9a0.
//
// Solidity: function respondToTask((uint256,uint32,bytes,uint32) task, (uint32,uint256) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) RespondToTask(task IIncredibleSquaringTaskManagerTask, taskResponse IIncredibleSquaringTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.RespondToTask(&_ContractIncredibleSquaringTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSquaringTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TransferOwnership(&_ContractIncredibleSquaringTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.Unpause(&_ContractIncredibleSquaringTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSquaringTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerInitialized)
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
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerInitializedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerInitialized)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSquaringTaskManagerInitialized, error) {
	event := new(ContractIncredibleSquaringTaskManagerInitialized)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
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
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSquaringTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint256,uint32,bytes,uint32) task)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSquaringTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSquaringTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSquaringTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
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
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSquaringTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSquaringTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerPaused)
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
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPaused represents a Paused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSquaringTaskManagerPaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerPaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleSquaringTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
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
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleSquaringTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleSquaringTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSquaringTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskCompleted)
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
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSquaringTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskCompleted)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerTaskResponded)
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
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSquaringTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSquaringTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0x349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a.
//
// Solidity: event TaskResponded((uint32,uint256) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSquaringTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSquaringTaskManagerTaskResponded)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerUnpaused)
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
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSquaringTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerUnpausedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerUnpaused)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSquaringTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSquaringTaskManagerUnpaused)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
