package mtree

import (
	"../common"
	"../share"
	"container/list"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Counter interface{}

type AugData struct {
	Min  Counter
	Max  Counter
	Hash common.SPHash
}

func (ad AugData) MaxCounterBytes() []byte {
	return msbPadding(ad.Max.(*big.Int).Bytes(), 32)
}

func (ad AugData) MinCounterBytes() []byte {
	return msbPadding(ad.Min.(*big.Int).Bytes(), 32)
}

type AugTree struct {
	MerkleTree
}

func _min(a, b Counter) Counter {
	left := a.(*big.Int)
	right := b.(*big.Int)
	if left.Cmp(right) == -1 {
		return left
	} else {
		return right
	}
}

func _max(a, b Counter) Counter {
	left := a.(*big.Int)
	right := b.(*big.Int)
	if left.Cmp(right) == 1 {
		return left
	} else {
		return right
	}
}

func _augElementHash(data ElementData) NodeData {
	s := data.(share.Share)
	fmt.Printf("Constructing node:\n")
	fmt.Printf("	Min: %v\n", s.Counter())
	fmt.Printf("	Max: %v\n", s.Counter())
	fmt.Printf("	Hash: %s\n", s.Hash().Hex())
	return AugData{
		Min:  s.Counter(),
		Max:  s.Counter(),
		Hash: s.Hash(),
	}
}

func _augHash(a, b NodeData) NodeData {
	left := a.(AugData)
	right := b.(AugData)
	h := common.SPHash{}
	keccak := crypto.Keccak256(
		left.MinCounterBytes(),
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, left.Hash[:]...),
		right.MaxCounterBytes(),
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, right.Hash[:]...),
	)
	copy(h[:common.HashLength], keccak[common.HashLength:])
	fmt.Printf("Prepare to construct node: \n")
	fmt.Printf("--> left_counter: 0x%s\n", hex.EncodeToString(left.MinCounterBytes()))
	fmt.Printf("--> left_hash: 0x%s\n", hex.EncodeToString(append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, left.Hash[:]...)))
	fmt.Printf("--> right_counter: 0x%s\n", hex.EncodeToString(right.MaxCounterBytes()))
	fmt.Printf("--> right_hash: 0x%s\n", hex.EncodeToString(append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, right.Hash[:]...)))
	fmt.Printf("Constructing node:\n")
	fmt.Printf("	Left: %s\n", left.Hash.Hex())
	fmt.Printf("	Right: %s\n", right.Hash.Hex())
	fmt.Printf("	Min: %v\n", _min(left.Min, right.Min).(*big.Int).Text(16))
	fmt.Printf("	Max: %v\n", _max(left.Max, right.Max).(*big.Int).Text(16))
	fmt.Printf("	Hash: %s\n", h.Hex())
	fmt.Printf("	Keccak: 0x%s\n", hex.EncodeToString(keccak))
	return AugData{
		Min:  _min(left.Min, right.Min),
		Max:  _max(left.Max, right.Max),
		Hash: h,
	}
}

func NewAugTree() *AugTree {
	mtbuf := list.New()
	return &AugTree{
		MerkleTree{
			mtbuf,
			_augHash,
			_augElementHash,
			false,
			map[uint32]bool{},
			[]uint32{},
		},
	}
}

func (amt AugTree) CounterBranchArray() []common.BranchElement {
	if amt.finalized {
		result := []common.BranchElement{}
		branches := amt.Branches()
		var node AugData
		for _, k := range amt.Indices() {
			// p := branches[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			nodes := branches[k].ToNodeArray()[1:]
			// fmt.Printf("Len proofs: %s\n", len(pfs))
			for _, n := range nodes {
				node = n.(AugData)
				// fmt.Printf("node %v\n", node)
				be := common.BranchElement{}
				if k%2 == 0 {
					copy(be[:], node.MaxCounterBytes())
				} else {
					copy(be[:], node.MinCounterBytes())
				}
				result = append(result, be)
				k >>= 1
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (amt AugTree) HashBranchArray() []common.BranchElement {
	if amt.finalized {
		result := []common.BranchElement{}
		branches := amt.Branches()
		var node AugData
		for _, k := range amt.Indices() {
			// p := branches[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			nodes := branches[k].ToNodeArray()[1:]
			// fmt.Printf("Len proofs: %s\n", len(pfs))
			for _, n := range nodes {
				node = n.(AugData)
				be := common.BranchElement{}
				copy(be[:], msbPadding(node.Hash[:], 32))
				result = append(result, be)
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}
