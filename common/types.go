package common

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
)

func BytesToBig(data []byte) *big.Int {
	n := new(big.Int)
	n.SetBytes(data)

	return n
}

func rev(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func (w Word) ToUint256Array() []big.Int {
	result := []big.Int{}
	for i := 0; i < WordLength/32; i++ {
		z := big.NewInt(0)
		// reverse the bytes because contract expects
		// big Int is construct in little endian
		z.SetBytes(rev(w[i*32 : (i+1)*32]))
		result = append(result, *z)
	}
	return result
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
