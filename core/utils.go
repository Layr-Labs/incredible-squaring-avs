package core

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	cstaskmanager "github.com/Layr-Labs/incredible-squaring-avs/contracts/bindings/IncredibleSquaringTaskManager"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

type PriceUpdateTaskResponse struct {
	Price    uint32
	Decimals uint32
	TaskId   uint32
	Source   string
}

// this hardcodes abi.encode() for cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
// unclear why abigen doesn't provide this out of the box...
func AbiEncodeTaskResponse(price uint32, source string, taskId uint32, decimals uint32) ([]byte, error) {
	priceUpdate := &PriceUpdateTaskResponse{Price: price, Source: source, TaskId: taskId}
	// The order here has to match the field ordering of cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{
			Name: "price",
			Type: "uint32",
		},
		{
			Name: "decimals",
			Type: "uint32",
		},
		{
			Name: "taskId",
			Type: "uint32",
		},
		{
			Name: "source",
			Type: "string",
		},
	})
	if err != nil {
		return nil, err
	}
	arguments := abi.Arguments{
		{
			Type: taskResponseType,
		},
	}

	bytes, err := arguments.Pack(priceUpdate)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// GetTaskResponseDigest returns the hash of the TaskResponse, which is what operators sign over
func GetTaskResponseDigest(price uint32, source string, taskId uint32, decimals uint32) ([32]byte, error) {

	encodeTaskResponseByte, err := AbiEncodeTaskResponse(price, source, taskId, decimals)
	if err != nil {
		return [32]byte{}, err
	}

	var taskResponseDigest [32]byte
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodeTaskResponseByte)
	copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

	return taskResponseDigest, nil
}

// BINDING UTILS - conversion from contract structs to golang structs

// BN254.sol is a library, so bindings for G1 Points and G2 Points are only generated
// in every contract that imports that library. Thus the output here will need to be
// type casted if G1Point is needed to interface with another contract (eg: BLSPublicKeyCompendium.sol)
func ConvertToBN254G1Point(input *bls.G1Point) cstaskmanager.BN254G1Point {
	output := cstaskmanager.BN254G1Point{
		X: input.X.BigInt(big.NewInt(0)),
		Y: input.Y.BigInt(big.NewInt(0)),
	}
	return output
}

func ConvertToBN254G2Point(input *bls.G2Point) cstaskmanager.BN254G2Point {
	output := cstaskmanager.BN254G2Point{
		X: [2]*big.Int{input.X.A1.BigInt(big.NewInt(0)), input.X.A0.BigInt(big.NewInt(0))},
		Y: [2]*big.Int{input.Y.A1.BigInt(big.NewInt(0)), input.Y.A0.BigInt(big.NewInt(0))},
	}
	return output
}
