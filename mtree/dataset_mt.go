package mtree

import (
	// "fmt"
	"container/list"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "encoding/hex"
)

type Word [128]byte

type hashFunc func([]byte, []byte) common.Hash

type node struct {
	Data      common.Hash
	NodeCount uint32
	Proofs    *map[uint32]proof
}

type SPMerkleTree struct {
	mtbuf          *list.List
	h              hashFunc
	finalized      bool
	indexes        map[uint32]bool
	orderedIndexes []uint32
}

// register indexes to build proof branches
func (mt *SPMerkleTree) RegisterIndex(indexes ...uint32) {
	for _, i := range indexes {
		mt.indexes[i] = true
		mt.orderedIndexes = append(mt.orderedIndexes, i)
	}
}

func (mt *SPMerkleTree) SetHashFunction(_h hashFunc) {
	mt.h = _h
}

func (mt *SPMerkleTree) Insert(data Word, index uint32) {
	// insert data into the mtbuf and aggregate the
	// hashes
	// because contract side is expecting the bytes
	// to be reversed each 32 bytes on leaf nodes
	first, second := conventionalWord(data)
	_node := node{mt.h(first, second), 1, &map[uint32]proof{}}
	// fmt.Printf("Inserted node for word (%s): %4s\n", hex.EncodeToString(data[:]), hex.EncodeToString(_node.Data[:]))
	if mt.indexes[index] {
		(*_node.Proofs)[index] = proof{
			RawData:    data,
			HashedData: _node.Data,
			Branch: &Branch{
				Hash:  _node.Data,
				Left:  nil,
				Right: nil,
			},
		}
	}
	mt.insertNode(_node)
}

func (mt *SPMerkleTree) insertNode(_node node) {
	var e, prev *list.Element
	var cNode, prevNode node
	e = mt.mtbuf.PushBack(_node)
	for {
		prev = e.Prev()
		cNode = e.Value.(node)
		if prev == nil {
			break
		}
		prevNode = prev.Value.(node)
		if cNode.NodeCount != prevNode.NodeCount {
			break
		}
		if prevNode.Proofs != nil {
			// fmt.Printf("Accepting right sibling\n")
			for k, v := range *prevNode.Proofs {
				v.Branch = AcceptRightSibling(v.Branch, cNode.Data)
				(*prevNode.Proofs)[k] = v
				// fmt.Printf("Proof: %v\n", v.String())
			}
		}
		if cNode.Proofs != nil {
			// fmt.Printf("Accepting left sibling\n")
			for k, v := range *cNode.Proofs {
				v.Branch = AcceptLeftSibling(v.Branch, prevNode.Data)
				(*prevNode.Proofs)[k] = v
				// fmt.Printf("Proof: %v\n", v.String())
			}
		}
		// fmt.Printf("Creating new Node: h(%4s, %4s) ", hex.EncodeToString(prevNode.Data[:]), hex.EncodeToString(cNode.Data[:]))
		prevNode.Data = mt.h(prevNode.Data[:], cNode.Data[:])
		// fmt.Printf("=> %4s\n", hex.EncodeToString(prevNode.Data[:]))
		prevNode.NodeCount = cNode.NodeCount*2 + 1

		mt.mtbuf.Remove(e)
		mt.mtbuf.Remove(prev)
		e = mt.mtbuf.PushBack(prevNode)
	}
}

func _hash(a, b []byte) common.Hash {
	result := [32]byte{}
	keccak := crypto.Keccak256(a, b)
	copy(result[:32], keccak[:])
	return result
}

func NewSPMerkleTree() *SPMerkleTree {
	mtbuf := list.New()
	return &SPMerkleTree{
		mtbuf,
		_hash,
		false,
		map[uint32]bool{},
		[]uint32{},
	}
}

func (mt *SPMerkleTree) Finalize() {
	if !mt.finalized && mt.mtbuf.Len() > 1 {
		for {
			dupNode := mt.mtbuf.Back().Value.(node)
			dupNode.Proofs = &map[uint32]proof{}
			mt.insertNode(dupNode)
			if mt.mtbuf.Len() == 1 {
				break
			}
		}
	}
	mt.finalized = true
}

func (mt SPMerkleTree) Root() common.Hash {
	if mt.finalized {
		return mt.mtbuf.Front().Value.(node).Data
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (mt SPMerkleTree) Proofs() map[uint32]proof {
	if mt.finalized {
		return *(mt.mtbuf.Front().Value.(node).Proofs)
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

// return only one array with necessary hashes for each
// index in order. Element's hash and root are not included
// eg. registered indexes are 1, 2, each needs 2 hashes
// then the function return an array of 4 hashes [a1, a2, b1, b2]
// where a1, a2 are proof branch for element at index 1
// b1, b2 are proof branch for element at index 2
func (mt SPMerkleTree) ProofArray() []common.Hash {
	if mt.finalized {
		result := []common.Hash{}
		proofs := mt.Proofs()
		for _, k := range mt.orderedIndexes {
			result = append(result, proofs[k].ToArray()[1:]...)
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}
