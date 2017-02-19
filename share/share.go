package share

import (
	spcommon "../common"
	"../ethash"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type Share struct {
	blockHeader     *types.Header
	nonce           types.BlockNonce
	mixDigest       common.Hash
	ShareDifficulty *big.Int
	SolutionState   int
}

func (s Share) Difficulty() *big.Int     { return s.blockHeader.Difficulty }
func (s Share) HashNoNonce() common.Hash { return s.blockHeader.HashNoNonce() }
func (s Share) Nonce() uint64            { return s.nonce.Uint64() }
func (s Share) MixDigest() common.Hash   { return s.mixDigest }
func (s Share) NumberU64() uint64        { return s.blockHeader.Number.Uint64() }
func (s Share) NonceBig() *big.Int {
	n := new(big.Int)
	n.SetBytes(s.nonce[:])
	return n
}

func (s *Share) AcceptSolution(nonce types.BlockNonce, mixDigest common.Hash) {
	s.nonce = nonce
	s.mixDigest = mixDigest
	eth := ethash.New()
	s.SolutionState = eth.SolutionState(s, s.ShareDifficulty)
}

func (s Share) BlockHeader() *types.Header {
	return s.blockHeader
}

func (s Share) RlpHeaderWithoutNonce() ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := rlp.Encode(buffer, []interface{}{
		s.BlockHeader().ParentHash,
		s.BlockHeader().UncleHash,
		s.BlockHeader().Coinbase,
		s.BlockHeader().Root,
		s.BlockHeader().TxHash,
		s.BlockHeader().ReceiptHash,
		s.BlockHeader().Bloom,
		s.BlockHeader().Difficulty,
		s.BlockHeader().Number,
		s.BlockHeader().GasLimit,
		s.BlockHeader().GasUsed,
		s.BlockHeader().Time,
		s.BlockHeader().Extra,
	})
	return buffer.Bytes(), err
}

func (s Share) Timestamp() big.Int {
	return *s.blockHeader.Time
}

// We use concatenation of timestamp and nonce
// as share counter
// Nonce in ethereum is 8 bytes so counter = timestamp << 64 + nonce
func (s Share) Counter() *big.Int {
	t := s.Timestamp()
	t.Lsh(&t, 64)
	n := big.NewInt(0).SetBytes(s.nonce[:])
	return t.Add(&t, n)
}

func (s Share) Hash() (result spcommon.SPHash) {
	h := s.blockHeader.HashNoNonce()
	copy(result[:spcommon.HashLength], h[spcommon.HashLength:])
	return
}

func (s Share) PrintInfo() {
	fmt.Printf("	ParentHash: %s\n", s.BlockHeader().ParentHash.Hex())
	fmt.Printf("	UncleHash: %s\n", s.BlockHeader().UncleHash.Hex())
	fmt.Printf("	Coinbase: %s\n", s.BlockHeader().Coinbase.Hex())
	fmt.Printf("	Root: %s\n", s.BlockHeader().Root.Hex())
	fmt.Printf("	TxHash: %s\n", s.BlockHeader().TxHash.Hex())
	fmt.Printf("	ReceiptHash: %s\n", s.BlockHeader().ReceiptHash.Hex())
	fmt.Printf("	Bloom: %s\n", s.BlockHeader().Bloom)
	fmt.Printf("	Difficulty: 0x%s\n", s.BlockHeader().Difficulty.Text(16))
	fmt.Printf("	Number: %s\n", s.BlockHeader().Number)
	fmt.Printf("	GasLimit: 0x%s\n", s.BlockHeader().GasLimit.Text(16))
	fmt.Printf("	GasUsed: 0x%s\n", s.BlockHeader().GasUsed.Text(16))
	fmt.Printf("	Time: %v\n", s.BlockHeader().Time.Bytes())
	fmt.Printf("	Nonce: 0x%s\n", hex.EncodeToString(s.BlockHeader().Nonce[:]))
	fmt.Printf("	Extra: %v\n", s.BlockHeader().Extra)
	fmt.Printf("	Counter: %v\n", s.Counter().Bytes())
	fmt.Printf("	Corresponding Min-Max: 0x%s\n", s.Counter().Text(16))
	fmt.Printf("	Corresponding Hash: %s\n", s.Hash().Hex())
	rlpEncoded, _ := s.RlpHeaderWithoutNonce()
	fmt.Printf("	RlpEncode: 0x%s\n", hex.EncodeToString(rlpEncoded))
}

func NewShare(h *types.Header, dif *big.Int) *Share {
	return &Share{
		h,
		types.BlockNonce{},
		common.Hash{},
		dif,
		0,
	}
}
