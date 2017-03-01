package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Updater interface {
	SetEpochData(opts *bind.TransactOpts, merkleRoot *big.Int, fullSizeIn128Resolution uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error)
	VerifyExtraData(opts *bind.CallOpts, extraData [32]byte, minerId [32]byte, difficulty *big.Int) (bool, error)
	To62Encoding(opts *bind.CallOpts, id *big.Int, numChars *big.Int) ([32]byte, error)
}
