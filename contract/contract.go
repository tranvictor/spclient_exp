package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Contract interface {
	Version(opts *bind.CallOpts) (string, error)
	IsRegistered(opts *bind.CallOpts) (bool, error)
	CanRegister(opts *bind.CallOpts) (bool, error)
	Register(opts *bind.TransactOpts, paymentAddress common.Address) (*types.Transaction, error)
	GetClaimSeed(opts *bind.CallOpts) (*big.Int, error)
	SubmitClaim(
		opts *bind.TransactOpts,
		numShares *big.Int,
		difficulty *big.Int,
		min *big.Int,
		max *big.Int,
		augMerkle *big.Int) (*types.Transaction, error)
	VerifyClaim(
		opts *bind.TransactOpts,
		rlpHeader []byte,
		nonce *big.Int,
		shareIndex *big.Int,
		dataSetLookup []*big.Int,
		witnessForLookup []*big.Int,
		augCountersBranch []*big.Int,
		augHashesBranch []*big.Int) (*types.Transaction, error)
	VerifyClaim_debug(
		opts *bind.CallOpts,
		rlpHeader []byte,
		nonce *big.Int,
		shareIndex *big.Int,
		dataSetLookup []*big.Int,
		witnessForLookup []*big.Int,
		augCountersBranch []*big.Int,
		augHashesBranch []*big.Int) (*big.Int, error)
}
