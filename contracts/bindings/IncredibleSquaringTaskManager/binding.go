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
	TaskRespondedBlock uint32
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
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSquaringTaskManagerMetaData contains all meta data concerning the ContractIncredibleSquaringTaskManager contract.
var ContractIncredibleSquaringTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAggregator\",\"inputs\":[{\"name\":\"newAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGenerator\",\"inputs\":[{\"name\":\"newGenerator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccessfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AggregatorUpdated\",\"inputs\":[{\"name\":\"oldAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GeneratorUpdated\",\"inputs\":[{\"name\":\"oldGenerator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGenerator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.Task\",\"components\":[{\"name\":\"numberToBeSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numberSquared\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSquaringTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskRespondedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006218380380620062188339810160408190526200003591620001ea565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b5919062000231565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000133919062000231565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b3919062000231565b6001600160a01b031660e0525063ffffffff16610100525062000258565b6001600160a01b0381168114620001e757600080fd5b50565b60008060408385031215620001fe57600080fd5b82516200020b81620001d1565b602084015190925063ffffffff811681146200022657600080fd5b809150509250929050565b6000602082840312156200024457600080fd5b81516200025181620001d1565b9392505050565b60805160a05160c05160e05161010051615f2e620002ea60003960008181610239015281816105970152611a870152600081816105600152612cb3015260008181610422015281816123470152612e950152600081816104490152818161305101526131fa01526000818161049601528181610ed9015281816129b401528181612b2d0152612d500152615f2e6000f3fe608060405234801561001057600080fd5b50600436106101d85760003560e01c806310d67a2f146101dd578063136439dd146101f2578063171f1d5b146102055780631ad43189146102345780631ce8a3e314610270578063245a7bfc146102a35780632cb223d5146102c35780632d89f6fc146102f157806331b36bd9146103115780633563b0d114610331578063416c7e5e146103515780634a7c7e4b146103645780634d2b57fe146103775780634f739f7414610397578063595c6a67146103b75780635ac86ab7146103bf5780635baec9a0146103e25780635c155662146103f55780635c975abb146104155780635df459461461041d57806368304835146104445780636b532e9e1461046b5780636b92787e1461047e5780636d14a987146104915780636efb4636146104b8578063715018a6146104d957806372d18e8d146104e15780637afa1eed146104ef578063886f1195146105025780638b00ce7c146105155780638da5cb5b14610525578063b98d09081461052d578063cefdc1d41461053a578063df5cf7231461055b578063f2fde38b14610582578063f5c9899d14610595578063f63c5bab146105bb578063f8c8765e146105c3578063f9120af6146105d6578063fabc1cbc146105e9575b600080fd5b6101f06101eb366004614832565b6105fc565b005b6101f061020036600461484f565b6106b8565b6102186102133660046149cd565b6107e5565b6040805192151583529015156020830152015b60405180910390f35b61025b7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161022b565b61029361027e366004614a3b565b60cc6020526000908152604090205460ff1681565b604051901515815260200161022b565b60cd546102b6906001600160a01b031681565b60405161022b9190614a58565b6102e36102d1366004614a3b565b60cb6020526000908152604090205481565b60405190815260200161022b565b6102e36102ff366004614a3b565b60ca6020526000908152604090205481565b61032461031f366004614a8f565b61094b565b60405161022b9190614b7d565b61034461033f366004614b97565b610a58565b60405161022b9190614cf2565b6101f061035f366004614d13565b610ed7565b6101f0610372366004614832565b61100d565b61038a610385366004614d96565b61101e565b60405161022b9190614de5565b6103aa6103a5366004614e7a565b611133565b60405161022b9190614f73565b6101f06117ee565b6102936103cd36600461503d565b606654600160ff9092169190911b9081161490565b6101f06103f03660046152f5565b6118a8565b610408610403366004615369565b611d15565b60405161022b91906153cc565b6066546102e3565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6101f0610479366004615404565b611ec6565b6101f061048c36600461548a565b612480565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6104cb6104c63660046154e5565b612621565b60405161022b9291906155a5565b6101f06134ab565b60c95463ffffffff1661025b565b60ce546102b6906001600160a01b031681565b6065546102b6906001600160a01b031681565b60c95461025b9063ffffffff1681565b6102b66134bf565b6097546102939060ff1681565b61054d6105483660046155ee565b6134ce565b60405161022b929190615625565b6102b67f000000000000000000000000000000000000000000000000000000000000000081565b6101f0610590366004614832565b613660565b7f000000000000000000000000000000000000000000000000000000000000000061025b565b61025b606481565b6101f06105d1366004615646565b6136d6565b6101f06105e4366004614832565b613809565b6101f06105f736600461484f565b61381a565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561064f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067391906156a2565b6001600160a01b0316336001600160a01b0316146106ac5760405162461bcd60e51b81526004016106a3906156bf565b60405180910390fd5b6106b581613971565b50565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e906106e8903390600401614a58565b602060405180830381865afa158015610705573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072991906156f7565b6107455760405162461bcd60e51b81526004016106a390615714565b606654818116146107b95760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d707420604482015277746f20756e70617573652066756e6374696f6e616c69747960401b60648201526084016106a3565b60668190556040518181523390600080516020615e39833981519152906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061082d5761082d61574a565b60200201518951600160200201518a602001516000600281106108525761085261574a565b60200201518b6020015160016002811061086e5761086e61574a565b602090810291909101518c518d8301516040516108cb9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108ee9190615760565b905061093d6109076109008884613a68565b8690613af3565b61090f613b7b565b6109336109248561091e613c3b565b90613a68565b61092d8c613c5c565b90613af3565b886201d4c0613ce0565b909890975095505050505050565b606081516001600160401b0381111561096657610966614868565b60405190808252806020026020018201604052801561098f578160200160208202803683370190505b50905060005b8251811015610a5157836001600160a01b03166313542a4e8483815181106109bf576109bf61574a565b60200260200101516040518263ffffffff1660e01b81526004016109e39190614a58565b602060405180830381865afa158015610a00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a249190615782565b828281518110610a3657610a3661574a565b6020908102919091010152610a4a816157b1565b9050610995565b5092915050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abe91906156a2565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2491906156a2565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8a91906156a2565b9050600086516001600160401b03811115610ba757610ba7614868565b604051908082528060200260200182016040528015610bda57816020015b6060815260200190600190039081610bc55790505b50905060005b8751811015610ecb576000888281518110610bfd57610bfd61574a565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610c5e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c8691908101906157cc565b905080516001600160401b03811115610ca157610ca1614868565b604051908082528060200260200182016040528015610cec57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cbf5790505b50848481518110610cff57610cff61574a565b602002602001018190525060005b8151811015610eb5576040518060600160405280876001600160a01b03166347b314e8858581518110610d4257610d4261574a565b60200260200101516040518263ffffffff1660e01b8152600401610d6891815260200190565b602060405180830381865afa158015610d85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da991906156a2565b6001600160a01b03168152602001838381518110610dc957610dc961574a565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610df757610df761574a565b6020026020010151878f6040518463ffffffff1660e01b8152600401610e1f9392919061585c565b602060405180830381865afa158015610e3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e60919061587b565b6001600160601b0316815250858581518110610e7e57610e7e61574a565b60200260200101518281518110610e9757610e9761574a565b60200260200101819052508080610ead906157b1565b915050610d0d565b5050508080610ec3906157b1565b915050610be0565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5991906156a2565b6001600160a01b0316336001600160a01b0316146110045760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527b391037b3103a3432903932b3b4b9ba393ca1b7b7b93234b730ba37b960211b608482015260a4016106a3565b6106b581613f04565b611015613f4b565b6106b581613faa565b606081516001600160401b0381111561103957611039614868565b604051908082528060200260200182016040528015611062578160200160208202803683370190505b50905060005b8251811015610a5157836001600160a01b031663296bb0648483815181106110925761109261574a565b60200260200101516040518263ffffffff1660e01b81526004016110b891815260200190565b602060405180830381865afa1580156110d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f991906156a2565b82828151811061110b5761110b61574a565b6001600160a01b039092166020928302919091019091015261112c816157b1565b9050611068565b61113b6146e7565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561117b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119f91906156a2565b90506111a96146e7565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e906111d9908b90899089906004016158a4565b600060405180830381865afa1580156111f6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261121e91908101906158ee565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290611250908b908b908b906004016159a5565b600060405180830381865afa15801561126d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261129591908101906158ee565b6040820152856001600160401b038111156112b2576112b2614868565b6040519080825280602002602001820160405280156112e557816020015b60608152602001906001900390816112d05790505b50606082015260005b60ff81168711156116ff576000856001600160401b0381111561131357611313614868565b60405190808252806020026020018201604052801561133c578160200160208202803683370190505b5083606001518360ff16815181106113565761135661574a565b602002602001018190525060005b868110156115ff5760008c6001600160a01b03166304ec63518a8a8581811061138f5761138f61574a565b905060200201358e886000015186815181106113ad576113ad61574a565b60200260200101516040518463ffffffff1660e01b81526004016113d3939291906159ce565b602060405180830381865afa1580156113f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141491906159ea565b90506001600160c01b0381166114b75760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527b3132903932b3b4b9ba32b932b21030ba10313637b1b5b73ab6b132b960211b608482015260a4016106a3565b8a8a8560ff168181106114cc576114cc61574a565b6001600160c01b03841692013560f81c9190911c6001908116141590506115ec57856001600160a01b031663dd9846b98a8a8581811061150e5761150e61574a565b905060200201358d8d8860ff1681811061152a5761152a61574a565b9050013560f81c60f81b60f81c8f6040518463ffffffff1660e01b81526004016115569392919061585c565b602060405180830381865afa158015611573573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115979190615a13565b85606001518560ff16815181106115b0576115b061574a565b602002602001015184815181106115c9576115c961574a565b63ffffffff90921660209283029190910190910152826115e8816157b1565b9350505b50806115f7816157b1565b915050611364565b506000816001600160401b0381111561161a5761161a614868565b604051908082528060200260200182016040528015611643578160200160208202803683370190505b50905060005b828110156116c45784606001518460ff168151811061166a5761166a61574a565b602002602001015181815181106116835761168361574a565b602002602001015182828151811061169d5761169d61574a565b63ffffffff90921660209283029190910190910152806116bc816157b1565b915050611649565b508084606001518460ff16815181106116df576116df61574a565b6020026020010181905250505080806116f790615a30565b9150506112ee565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611740573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176491906156a2565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90611797908b908b908e90600401615a50565b600060405180830381865afa1580156117b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117dc91908101906158ee565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81526001600160a01b03909116906346fbf68e9061181e903390600401614a58565b602060405180830381865afa15801561183b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061185f91906156f7565b61187b5760405162461bcd60e51b81526004016106a390615714565b60001960668190556040519081523390600080516020615e398339815191529060200160405180910390a2565b60cd546001600160a01b031633146119025760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c657200000060448201526064016106a3565b60006119146040850160208601614a3b565b90503660006119266040870187615a7a565b9092509050600061193d6080880160608901614a3b565b905060ca60006119506020890189614a3b565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161197c9190615ac0565b60405160208183030381529060405280519060200120146119f35760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d61746368207468656044820152600080516020615e7983398151915260648201526084016106a3565b600060cb81611a0560208a018a614a3b565b63ffffffff1663ffffffff1681526020019081526020016000205414611a825760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b60648201526084016106a3565b611aac7f000000000000000000000000000000000000000000000000000000000000000085615b61565b63ffffffff164363ffffffff161115611b1d5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b60648201526084016106a3565b600086604051602001611b309190615ba7565b604051602081830303815290604052805190602001209050600080611b588387878a8c612621565b9150915060005b85811015611c57578460ff1683602001518281518110611b8157611b8161574a565b6020026020010151611b939190615bb5565b6001600160601b0316606484600001518381518110611bb457611bb461574a565b60200260200101516001600160601b0316611bcf9190615be4565b1015611c45576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d60648201526084016106a3565b80611c4f816157b1565b915050611b5f565b5060408051808201825263ffffffff43168152602080820184905291519091611c84918c91849101615c03565b6040516020818303038152906040528051906020012060cb60008c6000016020810190611cb19190614a3b565b63ffffffff1663ffffffff168152602001908152602001600020819055507f349c1ee60e4e8972ee9dba642c1774543d5c4136879b7f4caaf04bf81a487a2a8a82604051611d00929190615c03565b60405180910390a15050505050505050505050565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401611d47929190615c2f565b600060405180830381865afa158015611d64573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611d8c91908101906158ee565b9050600084516001600160401b03811115611da957611da9614868565b604051908082528060200260200182016040528015611dd2578160200160208202803683370190505b50905060005b8551811015611ebc57866001600160a01b03166304ec6351878381518110611e0257611e0261574a565b602002602001015187868581518110611e1d57611e1d61574a565b60200260200101516040518463ffffffff1660e01b8152600401611e43939291906159ce565b602060405180830381865afa158015611e60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e8491906159ea565b6001600160c01b0316828281518110611e9f57611e9f61574a565b602090810291909101015280611eb4816157b1565b915050611dd8565b5095945050505050565b6000611ed56020850185614a3b565b63ffffffff8116600090815260cb6020526040902054909150853590611f475760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b60648201526084016106a3565b8484604051602001611f5a929190615c4e565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260cb90935291205414611fe75760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d61746368207468656044820152600080516020615e7983398151915260648201526084016106a3565b63ffffffff8216600090815260cc602052604090205460ff161561207f5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a4016106a3565b606461208e6020860186614a3b565b6120989190615b61565b63ffffffff164363ffffffff1611156121135760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527639b5903430b99030b63932b0b23c9032bc3834b932b21760491b60648201526084016106a3565b600061211f8280615be4565b905060208601358114600181141561216d57604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505061247a565b600085516001600160401b0381111561218857612188614868565b6040519080825280602002602001820160405280156121b1578160200160208202803683370190505b50905060005b8651811015612223576121f48782815181106121d5576121d561574a565b6020026020010151805160009081526020918201519091526040902090565b8282815181106122065761220661574a565b60209081029190910101528061221b816157b1565b9150506121b7565b50600061223660408b0160208c01614a3b565b82604051602001612248929190615c84565b604051602081830303815290604052805190602001209050876020013581146122f25760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a4016106a3565b600087516001600160401b0381111561230d5761230d614868565b604051908082528060200260200182016040528015612336578160200160208202803683370190505b50905060005b8851811015612429577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae68583815181106123865761238661574a565b60200260200101516040518263ffffffff1660e01b81526004016123ac91815260200190565b602060405180830381865afa1580156123c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123ed91906156a2565b8282815181106123ff576123ff61574a565b6001600160a01b039092166020928302919091019091015280612421816157b1565b91505061233c565b5063ffffffff8716600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505b50505050565b60ce546001600160a01b031633146124e45760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b60648201526084016106a3565b61251b604051806080016040528060008152602001600063ffffffff16815260200160608152602001600063ffffffff1681525090565b84815263ffffffff438116602080840191909152908516606083015260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251612581908290602001615ccc565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f1695b8d06ec800b4615e745cfb5bd00c1f2875615d42925c3b5afa543bb24c48906125e4908490615ccc565b60405180910390a260c9546126009063ffffffff166001615b61565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b61262961470f565b6000846126865760405162461bcd60e51b81526020600482015260376024820152600080516020615ed98339815191526044820152761c995cce88195b5c1d1e481c5d5bdc9d5b481a5b9c1d5d604a1b60648201526084016106a3565b6040830151518514801561269e575060a08301515185145b80156126ae575060c08301515185145b80156126be575060e08301515185145b6127285760405162461bcd60e51b81526020600482015260416024820152600080516020615ed983398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a4016106a3565b825151602084015151146127a05760405162461bcd60e51b815260206004820152604460248201819052600080516020615ed9833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a4016106a3565b4363ffffffff168463ffffffff161061280e5760405162461bcd60e51b815260206004820152603c6024820152600080516020615ed983398151915260448201527b7265733a20696e76616c6964207265666572656e636520626c6f636b60201b60648201526084016106a3565b604080518082019091526000808252602082015261282a61470f565b866001600160401b0381111561284257612842614868565b60405190808252806020026020018201604052801561286b578160200160208202803683370190505b506020820152866001600160401b0381111561288957612889614868565b6040519080825280602002602001820160405280156128b2578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b038111156128e6576128e6614868565b60405190808252806020026020018201604052801561290f578160200160208202803683370190505b5081526020860151516001600160401b0381111561292f5761292f614868565b604051908082528060200260200182016040528015612958578160200160208202803683370190505b5081602001819052506000612a2a8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015612a01573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a259190615d57565b613ffc565b905060005b876020015151811015612c8f57612a55886020015182815181106121d5576121d561574a565b83602001518281518110612a6b57612a6b61574a565b60209081029190910101528015612b2b576020830151612a8c600183615d74565b81518110612a9c57612a9c61574a565b602002602001015160001c83602001518281518110612abd57612abd61574a565b602002602001015160001c11612b2b576040805162461bcd60e51b8152602060048201526024810191909152600080516020615ed983398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f7274656460648201526084016106a3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110612b7057612b7061574a565b60200260200101518b8b600001518581518110612b8f57612b8f61574a565b60200260200101516040518463ffffffff1660e01b8152600401612bb5939291906159ce565b602060405180830381865afa158015612bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf691906159ea565b6001600160c01b031683600001518281518110612c1557612c1561574a565b602002602001018181525050612c7b610900612c4f8486600001518581518110612c4157612c4161574a565b60200260200101511661407d565b8a602001518481518110612c6557612c6561574a565b60200260200101516140a890919063ffffffff16565b945080612c87816157b1565b915050612a2f565b5050612c9a83614180565b60975490935060ff16600081612cb1576000612d33565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d339190615782565b905060005b8a81101561337e578215612e93578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110612d8f57612d8f61574a565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612dcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df39190615782565b612dfd9190615d8b565b11612e935760405162461bcd60e51b81526020600482015260666024820152600080516020615ed983398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c4016106a3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110612ed457612ed461574a565b9050013560f81c60f81b60f81c8c8c60a001518581518110612ef857612ef861574a565b60200260200101516040518463ffffffff1660e01b8152600401612f1e93929190615da3565b602060405180830381865afa158015612f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f5f9190615dc4565b6001600160401b031916612f828a6040015183815181106121d5576121d561574a565b6001600160401b0319161461301d5760405162461bcd60e51b81526020600482015260616024820152600080516020615ed983398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c4016106a3565b61304d896040015182815181106130365761303661574a565b602002602001015187613af390919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106130905761309061574a565b9050013560f81c60f81b60f81c8c8c60c0015185815181106130b4576130b461574a565b60200260200101516040518463ffffffff1660e01b81526004016130da93929190615da3565b602060405180830381865afa1580156130f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061311b919061587b565b856020015182815181106131315761313161574a565b6001600160601b0390921660209283029190910182015285015180518290811061315d5761315d61574a565b60200260200101518560000151828151811061317b5761317b61574a565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a6020015151811015613369576131f3866000015182815181106131c5576131c561574a565b60200260200101518f8f868181106131df576131df61574a565b600192013560f81c9290921c811614919050565b15613357577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106132395761323961574a565b9050013560f81c60f81b60f81c8e8960200151858151811061325d5761325d61574a565b60200260200101518f60e00151888151811061327b5761327b61574a565b602002602001015187815181106132945761329461574a565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa1580156132f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061331c919061587b565b87518051859081106133305761333061574a565b602002602001018181516133449190615dee565b6001600160601b03169052506001909101905b80613361816157b1565b91505061319f565b50508080613376906157b1565b915050612d38565b5050506000806133988c868a606001518b608001516107e5565b91509150816134095760405162461bcd60e51b81526020600482015260436024820152600080516020615ed983398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a4016106a3565b806134665760405162461bcd60e51b81526020600482015260396024820152600080516020615ed98339815191526044820152781c995cce881cda59db985d1d5c99481a5cc81a5b9d985b1a59603a1b60648201526084016106a3565b50506000878260200151604051602001613481929190615c84565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b6134b3613f4b565b6134bd600061420f565b565b6033546001600160a01b031690565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106135095761350961574a565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906135459088908690600401615c2f565b600060405180830381865afa158015613562573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261358a91908101906158ee565b60008151811061359c5761359c61574a565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015613608573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061362c91906159ea565b6001600160c01b03169050600061364282614261565b9050816136508a838a610a58565b9550955050505050935093915050565b613668613f4b565b6001600160a01b0381166136cd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016106a3565b6106b58161420f565b600054610100900460ff16158080156136f65750600054600160ff909116105b806137105750303b158015613710575060005460ff166001145b6137735760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106a3565b6000805460ff191660011790558015613796576000805461ff0019166101001790555b6137a185600061432d565b6137aa8461420f565b6137b383614405565b6137bc82613faa565b8015613802576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b613811613f4b565b6106b581614405565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561386d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061389191906156a2565b6001600160a01b0316336001600160a01b0316146138c15760405162461bcd60e51b81526004016106a3906156bf565b60665419811960665419161461393a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d706044820152777420746f2070617573652066756e6374696f6e616c69747960401b60648201526084016106a3565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107da565b6001600160a01b0381166139ff5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a4016106a3565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b613a70614729565b613a78614743565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613aab57613aad565bfe5b5080613aeb5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016106a3565b505092915050565b613afb614729565b613b03614761565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613aab575080613aeb5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016106a3565b613b8361477f565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b613c43614729565b5060408051808201909152600181526002602082015290565b613c64614729565b60008080613c80600080516020615e5983398151915286615760565b90505b613c8c81614457565b9093509150600080516020615e59833981519152828309831415613cc6576040805180820190915290815260208101919091529392505050565b600080516020615e59833981519152600182089050613c83565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613d126147a4565b60005b6002811015613ed7576000613d2b826006615be4565b9050848260028110613d3f57613d3f61574a565b60200201515183613d51836000615d8b565b600c8110613d6157613d6161574a565b6020020152848260028110613d7857613d7861574a565b60200201516020015183826001613d8f9190615d8b565b600c8110613d9f57613d9f61574a565b6020020152838260028110613db657613db661574a565b6020020151515183613dc9836002615d8b565b600c8110613dd957613dd961574a565b6020020152838260028110613df057613df061574a565b6020020151516001602002015183613e09836003615d8b565b600c8110613e1957613e1961574a565b6020020152838260028110613e3057613e3061574a565b602002015160200151600060028110613e4b57613e4b61574a565b602002015183613e5c836004615d8b565b600c8110613e6c57613e6c61574a565b6020020152838260028110613e8357613e8361574a565b602002015160200151600160028110613e9e57613e9e61574a565b602002015183613eaf836005615d8b565b600c8110613ebf57613ebf61574a565b60200201525080613ecf816157b1565b915050613d15565b50613ee06147c3565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b33613f546134bf565b6001600160a01b0316146134bd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106a3565b60ce80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b90600090a35050565b600080614008846144d9565b9050808360ff166001901b116140745760405162461bcd60e51b815260206004820152603f6024820152600080516020615eb983398151915260448201527f69746d61703a206269746d61702065786365656473206d61782076616c75650060648201526084016106a3565b90505b92915050565b6000805b821561407757614092600184615d74565b90921691806140a081615e16565b915050614081565b6140b0614729565b6102008261ffff16106140f85760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b60448201526064016106a3565b8161ffff166001141561410c575081614077565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061417557600161ffff871660ff83161c81161415614158576141558484613af3565b93505b6141628384613af3565b92506201fffe600192831b169101614128565b509195945050505050565b614188614729565b815115801561419957506020820151155b156141b7575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615e5983398151915284602001516141ea9190615760565b61420290600080516020615e59833981519152615d74565b905292915050565b919050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606060008061426f8461407d565b61ffff166001600160401b0381111561428a5761428a614868565b6040519080825280601f01601f1916602001820160405280156142b4576020820181803683370190505b5090506000805b8251821080156142cc575061010081105b15614323576001811b935085841615614313578060f81b8383815181106142f5576142f561574a565b60200101906001600160f81b031916908160001a9053508160010191505b61431c816157b1565b90506142bb565b5090949350505050565b6065546001600160a01b031615801561434e57506001600160a01b03821615155b6143d05760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a4016106a3565b60668190556040518181523390600080516020615e398339815191529060200160405180910390a261440182613971565b5050565b60cd80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f290600090a35050565b60008080600080516020615e598339815191526003600080516020615e5983398151915286600080516020615e598339815191528889090908905060006144cd827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615e59833981519152614642565b91959194509092505050565b6000610100825111156145505760405162461bcd60e51b815260206004820152604460248201819052600080516020615eb9833981519152908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a4016106a3565b815161455e57506000919050565b600080836000815181106145745761457461574a565b0160200151600160f89190911c81901b92505b8451811015614639578481815181106145a2576145a261574a565b0160200151600160f89190911c1b91508282116146255760405162461bcd60e51b81526020600482015260476024820152600080516020615eb983398151915260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a4016106a3565b91811791614632816157b1565b9050614587565b50909392505050565b60008061464d6147c3565b6146556147e1565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613aab5750826146dc5760405162461bcd60e51b815260206004820152601a602482015279424e3235342e6578704d6f643a2063616c6c206661696c75726560301b60448201526064016106a3565b505195945050505050565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060608152602001606081525090565b604051806040016040528060008152602001600081525090565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806147926147ff565b815260200161479f6147ff565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106b557600080fd5b60006020828403121561484457600080fd5b81356140748161481d565b60006020828403121561486157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156148a0576148a0614868565b60405290565b60405161010081016001600160401b03811182821017156148a0576148a0614868565b604051601f8201601f191681016001600160401b03811182821017156148f1576148f1614868565b604052919050565b60006040828403121561490b57600080fd5b61491361487e565b9050813581526020820135602082015292915050565b600082601f83011261493a57600080fd5b604080519081016001600160401b038111828210171561495c5761495c614868565b806040525080604084018581111561497357600080fd5b845b81811015614175578035835260209283019201614975565b60006080828403121561499f57600080fd5b6149a761487e565b90506149b38383614929565b81526149c28360408401614929565b602082015292915050565b60008060008061012085870312156149e457600080fd5b843593506149f586602087016148f9565b9250614a04866060870161498d565b9150614a138660e087016148f9565b905092959194509250565b63ffffffff811681146106b557600080fd5b803561420a81614a1e565b600060208284031215614a4d57600080fd5b813561407481614a1e565b6001600160a01b0391909116815260200190565b60006001600160401b03821115614a8557614a85614868565b5060051b60200190565b60008060408385031215614aa257600080fd5b8235614aad8161481d565b91506020838101356001600160401b03811115614ac957600080fd5b8401601f81018613614ada57600080fd5b8035614aed614ae882614a6c565b6148c9565b81815260059190911b82018301908381019088831115614b0c57600080fd5b928401925b82841015614b33578335614b248161481d565b82529284019290840190614b11565b80955050505050509250929050565b600081518084526020808501945080840160005b83811015614b7257815187529582019590820190600101614b56565b509495945050505050565b602081526000614b906020830184614b42565b9392505050565b600080600060608486031215614bac57600080fd5b8335614bb78161481d565b92506020848101356001600160401b0380821115614bd457600080fd5b818701915087601f830112614be857600080fd5b813581811115614bfa57614bfa614868565b614c0c601f8201601f191685016148c9565b91508082528884828501011115614c2257600080fd5b8084840185840137600084828401015250809450505050614c4560408501614a30565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614ce4578385038a52825180518087529087019087870190845b81811015614ccf57835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614c8b565b50509a87019a95505091850191600101614c6d565b509298975050505050505050565b602081526000614b906020830184614c4e565b80151581146106b557600080fd5b600060208284031215614d2557600080fd5b813561407481614d05565b600082601f830112614d4157600080fd5b81356020614d51614ae883614a6c565b82815260059290921b84018101918181019086841115614d7057600080fd5b8286015b84811015614d8b5780358352918301918301614d74565b509695505050505050565b60008060408385031215614da957600080fd5b8235614db48161481d565b915060208301356001600160401b03811115614dcf57600080fd5b614ddb85828601614d30565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614e265783516001600160a01b031683529284019291840191600101614e01565b50909695505050505050565b60008083601f840112614e4457600080fd5b5081356001600160401b03811115614e5b57600080fd5b602083019150836020828501011115614e7357600080fd5b9250929050565b60008060008060008060808789031215614e9357600080fd5b8635614e9e8161481d565b95506020870135614eae81614a1e565b945060408701356001600160401b0380821115614eca57600080fd5b614ed68a838b01614e32565b90965094506060890135915080821115614eef57600080fd5b818901915089601f830112614f0357600080fd5b813581811115614f1257600080fd5b8a60208260051b8501011115614f2757600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614b7257815163ffffffff1687529582019590820190600101614f51565b600060208083528351608082850152614f8f60a0850182614f3d565b905081850151601f1980868403016040870152614fac8383614f3d565b92506040870151915080868403016060870152614fc98383614f3d565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015615020578487830301845261500e828751614f3d565b95880195938801939150600101614ff4565b509998505050505050505050565b60ff811681146106b557600080fd5b60006020828403121561504f57600080fd5b81356140748161502e565b60006080828403121561506c57600080fd5b50919050565b60006040828403121561506c57600080fd5b600082601f83011261509557600080fd5b813560206150a5614ae883614a6c565b82815260059290921b840181019181810190868411156150c457600080fd5b8286015b84811015614d8b5780356150db81614a1e565b83529183019183016150c8565b600082601f8301126150f957600080fd5b81356020615109614ae883614a6c565b82815260069290921b8401810191818101908684111561512857600080fd5b8286015b84811015614d8b5761513e88826148f9565b83529183019160400161512c565b600082601f83011261515d57600080fd5b8135602061516d614ae883614a6c565b82815260059290921b8401810191818101908684111561518c57600080fd5b8286015b84811015614d8b5780356001600160401b038111156151af5760008081fd5b6151bd8986838b0101615084565b845250918301918301615190565b600061018082840312156151de57600080fd5b6151e66148a6565b905081356001600160401b03808211156151ff57600080fd5b61520b85838601615084565b8352602084013591508082111561522157600080fd5b61522d858386016150e8565b6020840152604084013591508082111561524657600080fd5b615252858386016150e8565b6040840152615264856060860161498d565b60608401526152768560e086016148f9565b608084015261012084013591508082111561529057600080fd5b61529c85838601615084565b60a08401526101408401359150808211156152b657600080fd5b6152c285838601615084565b60c08401526101608401359150808211156152dc57600080fd5b506152e98482850161514c565b60e08301525092915050565b60008060006080848603121561530a57600080fd5b83356001600160401b038082111561532157600080fd5b61532d8783880161505a565b945061533c8760208801615072565b9350606086013591508082111561535257600080fd5b5061535f868287016151cb565b9150509250925092565b60008060006060848603121561537e57600080fd5b83356153898161481d565b925060208401356001600160401b038111156153a457600080fd5b6153b086828701614d30565b92505060408401356153c181614a1e565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015614e26578351835292840192918401916001016153e8565b60008060008060c0858703121561541a57600080fd5b84356001600160401b038082111561543157600080fd5b61543d8883890161505a565b955061544c8860208901615072565b945061545b8860608901615072565b935060a087013591508082111561547157600080fd5b5061547e878288016150e8565b91505092959194509250565b600080600080606085870312156154a057600080fd5b8435935060208501356154b281614a1e565b925060408501356001600160401b038111156154cd57600080fd5b6154d987828801614e32565b95989497509550505050565b6000806000806000608086880312156154fd57600080fd5b8535945060208601356001600160401b038082111561551b57600080fd5b61552789838a01614e32565b90965094506040880135915061553c82614a1e565b9092506060870135908082111561555257600080fd5b5061555f888289016151cb565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614b725781516001600160601b031687529582019590820190600101615580565b60408152600083516040808401526155c0608084018261556c565b90506020850151603f198483030160608501526155dd828261556c565b925050508260208301529392505050565b60008060006060848603121561560357600080fd5b833561560e8161481d565b92506020840135915060408401356153c181614a1e565b82815260406020820152600061563e6040830184614c4e565b949350505050565b6000806000806080858703121561565c57600080fd5b84356156678161481d565b935060208501356156778161481d565b925060408501356156878161481d565b915060608501356156978161481d565b939692955090935050565b6000602082840312156156b457600080fd5b81516140748161481d565b6020808252602a90820152600080516020615e9983398151915260408201526939903ab73830bab9b2b960b11b606082015260800190565b60006020828403121561570957600080fd5b815161407481614d05565b6020808252602890820152600080516020615e9983398151915260408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261577d57634e487b7160e01b600052601260045260246000fd5b500690565b60006020828403121561579457600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006000198214156157c5576157c561579b565b5060010190565b600060208083850312156157df57600080fd5b82516001600160401b038111156157f557600080fd5b8301601f8101851361580657600080fd5b8051615814614ae882614a6c565b81815260059190911b8201830190838101908783111561583357600080fd5b928401925b8284101561585157835182529284019290840190615838565b979650505050505050565b92835260ff91909116602083015263ffffffff16604082015260600190565b60006020828403121561588d57600080fd5b81516001600160601b038116811461407457600080fd5b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156158d157600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561590157600080fd5b82516001600160401b0381111561591757600080fd5b8301601f8101851361592857600080fd5b8051615936614ae882614a6c565b81815260059190911b8201830190838101908783111561595557600080fd5b928401925b8284101561585157835161596d81614a1e565b8252928401929084019061595a565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006159c560408301848661597c565b95945050505050565b92835263ffffffff918216602084015216604082015260600190565b6000602082840312156159fc57600080fd5b81516001600160c01b038116811461407457600080fd5b600060208284031215615a2557600080fd5b815161407481614a1e565b600060ff821660ff811415615a4757615a4761579b565b60010192915050565b604081526000615a6460408301858761597c565b905063ffffffff83166020830152949350505050565b6000808335601e19843603018112615a9157600080fd5b8301803591506001600160401b03821115615aab57600080fd5b602001915036819003821315614e7357600080fd5b602081528135602082015260006020830135615adb81614a1e565b63ffffffff81166040840152506040830135601e19843603018112615aff57600080fd5b830180356001600160401b03811115615b1757600080fd5b803603851315615b2657600080fd5b60806060850152615b3e60a08501826020850161597c565b915050615b4d60608501614a30565b63ffffffff81166080850152509392505050565b600063ffffffff808316818516808303821115615b8057615b8061579b565b01949350505050565b8035615b9481614a1e565b63ffffffff168252602090810135910152565b604081016140778284615b89565b60006001600160601b0382811684821681151582840482111615615bdb57615bdb61579b565b02949350505050565b6000816000190483118215151615615bfe57615bfe61579b565b500290565b60808101615c118285615b89565b63ffffffff8351166040830152602083015160608301529392505050565b63ffffffff8316815260406020820152600061563e6040830184614b42565b60808101615c5c8285615b89565b8235615c6781614a1e565b63ffffffff16604083015260209290920135606090910152919050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b83811015615cbf57815185529382019390820190600101615ca3565b5092979650505050505050565b6000602080835283518184015263ffffffff8185015116604084015260408401516080606085015280518060a086015260005b81811015615d1b5782810184015186820160c001528301615cff565b81811115615d2d57600060c083880101525b50606086015163ffffffff811660808701529250601f01601f19169390930160c001949350505050565b600060208284031215615d6957600080fd5b81516140748161502e565b600082821015615d8657615d8661579b565b500390565b60008219821115615d9e57615d9e61579b565b500190565b60ff93909316835263ffffffff918216602084015216604082015260600190565b600060208284031215615dd657600080fd5b81516001600160401b03198116811461407457600080fd5b60006001600160601b0383811690831681811015615e0e57615e0e61579b565b039392505050565b600061ffff80831681811415615e2e57615e2e61579b565b600101939250505056feab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47206f6e65207265636f7264656420696e2074686520636f6e74726163740000006d73672e73656e646572206973206e6f74207065726d697373696f6e656420614269746d61705574696c732e6f72646572656442797465734172726179546f42424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220407bb832b2840c9f93930e69fd2e8c4dff4b81156cf70cf4976227f29698bf3464736f6c634300080c0033",
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

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorFromId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetBatchOperatorId(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operators)
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
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
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
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
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
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSquaringTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
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

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCaller) TaskSuccessfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSquaringTaskManager.contract.Call(opts, &out, "taskSuccessfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
}

// TaskSuccessfullyChallenged is a free data retrieval call binding the contract method 0x1ce8a3e3.
//
// Solidity: function taskSuccessfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerCallerSession) TaskSuccessfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.TaskSuccessfullyChallenged(&_ContractIncredibleSquaringTaskManager.CallOpts, arg0)
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

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetAggregator(opts *bind.TransactOpts, newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setAggregator", newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetAggregator(newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetAggregator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newAggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address newAggregator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetAggregator(newAggregator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetAggregator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newAggregator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactor) SetGenerator(opts *bind.TransactOpts, newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.contract.Transact(opts, "setGenerator", newGenerator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerSession) SetGenerator(newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetGenerator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newGenerator)
}

// SetGenerator is a paid mutator transaction binding the contract method 0x4a7c7e4b.
//
// Solidity: function setGenerator(address newGenerator) returns()
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerTransactorSession) SetGenerator(newGenerator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSquaringTaskManager.Contract.SetGenerator(&_ContractIncredibleSquaringTaskManager.TransactOpts, newGenerator)
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

// ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
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
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerAggregatorUpdated represents a AggregatorUpdated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts, oldAggregator []common.Address, newAggregator []common.Address) (*ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerAggregatorUpdatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerAggregatorUpdated, oldAggregator []common.Address, newAggregator []common.Address) (event.Subscription, error) {

	var oldAggregatorRule []interface{}
	for _, oldAggregatorItem := range oldAggregator {
		oldAggregatorRule = append(oldAggregatorRule, oldAggregatorItem)
	}
	var newAggregatorRule []interface{}
	for _, newAggregatorItem := range newAggregator {
		newAggregatorRule = append(newAggregatorRule, newAggregatorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "AggregatorUpdated", oldAggregatorRule, newAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address indexed oldAggregator, address indexed newAggregator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseAggregatorUpdated(log types.Log) (*ContractIncredibleSquaringTaskManagerAggregatorUpdated, error) {
	event := new(ContractIncredibleSquaringTaskManagerAggregatorUpdated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator is returned from FilterGeneratorUpdated and is used to iterate over the raw logs and unpacked data for GeneratorUpdated events raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator struct {
	Event *ContractIncredibleSquaringTaskManagerGeneratorUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
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
		it.Event = new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
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
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSquaringTaskManagerGeneratorUpdated represents a GeneratorUpdated event raised by the ContractIncredibleSquaringTaskManager contract.
type ContractIncredibleSquaringTaskManagerGeneratorUpdated struct {
	OldGenerator common.Address
	NewGenerator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGeneratorUpdated is a free log retrieval operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) FilterGeneratorUpdated(opts *bind.FilterOpts, oldGenerator []common.Address, newGenerator []common.Address) (*ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator, error) {

	var oldGeneratorRule []interface{}
	for _, oldGeneratorItem := range oldGenerator {
		oldGeneratorRule = append(oldGeneratorRule, oldGeneratorItem)
	}
	var newGeneratorRule []interface{}
	for _, newGeneratorItem := range newGenerator {
		newGeneratorRule = append(newGeneratorRule, newGeneratorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.FilterLogs(opts, "GeneratorUpdated", oldGeneratorRule, newGeneratorRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSquaringTaskManagerGeneratorUpdatedIterator{contract: _ContractIncredibleSquaringTaskManager.contract, event: "GeneratorUpdated", logs: logs, sub: sub}, nil
}

// WatchGeneratorUpdated is a free log subscription operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) WatchGeneratorUpdated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSquaringTaskManagerGeneratorUpdated, oldGenerator []common.Address, newGenerator []common.Address) (event.Subscription, error) {

	var oldGeneratorRule []interface{}
	for _, oldGeneratorItem := range oldGenerator {
		oldGeneratorRule = append(oldGeneratorRule, oldGeneratorItem)
	}
	var newGeneratorRule []interface{}
	for _, newGeneratorItem := range newGenerator {
		newGeneratorRule = append(newGeneratorRule, newGeneratorItem)
	}

	logs, sub, err := _ContractIncredibleSquaringTaskManager.contract.WatchLogs(opts, "GeneratorUpdated", oldGeneratorRule, newGeneratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
				if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
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

// ParseGeneratorUpdated is a log parse operation binding the contract event 0x0ddfab8a635d71f15d72e2d2dff55d32119d13270d2ea4c3dc0043b66c2c476b.
//
// Solidity: event GeneratorUpdated(address indexed oldGenerator, address indexed newGenerator)
func (_ContractIncredibleSquaringTaskManager *ContractIncredibleSquaringTaskManagerFilterer) ParseGeneratorUpdated(log types.Log) (*ContractIncredibleSquaringTaskManagerGeneratorUpdated, error) {
	event := new(ContractIncredibleSquaringTaskManagerGeneratorUpdated)
	if err := _ContractIncredibleSquaringTaskManager.contract.UnpackLog(event, "GeneratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
