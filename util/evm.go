package util

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/rpc"
)

func DefaultBlockContext(at *big.Int) vm.BlockContext {
	var randBytes []byte

	// read bytes from secure random
	rand.Read(randBytes)

	randomhash := common.BytesToHash(randBytes)

	return vm.BlockContext{
		BlockNumber: at,
		Coinbase:    common.HexToAddress("0x1"),
		CanTransfer: func(sd vm.StateDB, a common.Address, delta *big.Int) bool {
			return sd.GetBalance(a).Cmp(delta) == 1
		},

		Transfer: func(sd vm.StateDB, from, to common.Address, delta *big.Int) {
			sd.SubBalance(from, delta)
			sd.AddBalance(to, delta)
		},

		GetHash: func(u uint64) common.Hash {
			b := new(big.Int).SetUint64(u)
			return common.BigToHash(b)
		},
		Time:       uint64(time.Now().Unix()),
		GasLimit:   math.MaxUint64,
		Difficulty: big.NewInt(0),
		Random:     &randomhash,
	}
}

// NewEVM creates a new EVM instance with a storage backed by an RPC client.
func NewEVM(rpc *rpc.Client, at *big.Int) {
	// db := remotedb.New(rpc)

	// vm.NewEVM(DefaultBlockContext(at), vm.TxContext{
	// 	Origin:     common.HexToAddress("0x0"),
	// 	GasPrice:   big.NewInt(0),
	// 	BlobHashes: []common.Hash{},
	// }, db, params.MainnetChainConfig, vm.Config{})
}
