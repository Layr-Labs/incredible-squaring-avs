package common

import (
	"context"
	_ "embed"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:embed abis/IncredibleSquaringTaskManager.json
var IncredibleSquaringTaskManagerAbi []byte

type EthClientInterface interface {
	wallet.EthBackend
	TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error)
	ChainID(ctx context.Context) (*big.Int, error)

	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	BlockNumber(ctx context.Context) (uint64, error)
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
}
