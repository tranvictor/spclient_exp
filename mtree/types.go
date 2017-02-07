package mtree

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

const (
	HashLength          = 16
	WordLength          = 128
	BranchElementLength = 32
)

type (
	Word          [WordLength]byte
	SPHash        [HashLength]byte
	BranchElement [BranchElementLength]byte
	hashFunc      func([]byte, []byte) SPHash
)

func BytesToBig(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

func (h SPHash) Str() string   { return string(h[:]) }
func (h SPHash) Bytes() []byte { return h[:] }
func (h SPHash) Big() *big.Int { return BytesToBig(h[:]) }
func (h SPHash) Hex() string   { return hexutil.Encode(h[:]) }

func (h BranchElement) Str() string   { return string(h[:]) }
func (h BranchElement) Bytes() []byte { return h[:] }
func (h BranchElement) Big() *big.Int { return BytesToBig(h[:]) }
func (h BranchElement) Hex() string   { return hexutil.Encode(h[:]) }

func BranchElementFromHash(a, b SPHash) BranchElement {
	result := BranchElement{}
	copy(result[:], append(a[:], b[:]...)[:BranchElementLength])
	return result
}

type node struct {
	Data      SPHash
	NodeCount uint32
	Proofs    *map[uint32]proof
}
