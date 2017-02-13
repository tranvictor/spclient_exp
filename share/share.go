package share

import (
	spcommon "../common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Claim []*Share

func (c Claim) Len() int      { return len(c) }
func (c Claim) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Claim) Less(i, j int) bool {
	return c[i].Counter().Cmp(c[j].Counter()) == -1
}

type Share struct {
	blockHeader *types.Header
	difficulty  *big.Int
}

func NewShare() *Share {
	return &Share{
		&types.Header{},
		big.NewInt(0),
	}
}

func (s Share) BlockHeader() *types.Header {
	return s.blockHeader
}

func (s Share) Timestamp() big.Int {
	return *s.blockHeader.Time
}

func (s Share) Nonce() []byte {
	return s.blockHeader.Nonce[:]
}

// We use concatenation of timestamp and nonce
// as share counter
// Nonce in ethereum is 8 bytes so counter = timestamp << 64 + nonce
func (s Share) Counter() *big.Int {
	t := s.Timestamp()
	t.Lsh(&t, 64)
	n := big.NewInt(0).SetBytes(s.Nonce())
	return t.Add(&t, n)
}

func (s Share) Hash() (result spcommon.SPHash) {
	h := s.blockHeader.HashNoNonce()
	copy(result[:spcommon.HashLength], h[spcommon.HashLength:])
	return
}
