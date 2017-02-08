package mtree

import (
	"container/list"
	// "encoding/hex"
	// "fmt"
	"../common"
	"github.com/ethereum/go-ethereum/crypto"
)

type node struct {
	Data      common.SPHash
	NodeCount uint32
	Branches  *map[uint32]BranchTree
}

type hashFunc func([]byte, []byte) common.SPHash

type SPMerkleTree struct {
	mtbuf          *list.List
	h              hashFunc
	finalized      bool
	indexes        map[uint32]bool
	orderedIndexes []uint32
}

// register indexes to build branches
func (mt *SPMerkleTree) RegisterIndex(indexes ...uint32) {
	for _, i := range indexes {
		mt.indexes[i] = true
		mt.orderedIndexes = append(mt.orderedIndexes, i)
	}
}

func (mt *SPMerkleTree) SetHashFunction(_h hashFunc) {
	mt.h = _h
}

func (mt *SPMerkleTree) Insert(data common.Word, index uint32) {
	// insert data into the mtbuf and aggregate the
	// hashes
	// because contract side is expecting the bytes
	// to be reversed each 32 bytes on leaf nodes
	first, second := conventionalWord(data)
	_node := node{mt.h(first, second), 1, &map[uint32]BranchTree{}}
	// fmt.Printf("Inserted node for word (%s): %4s\n", hex.EncodeToString(data[:]), hex.EncodeToString(_node.Data[:]))
	if mt.indexes[index] {
		(*_node.Branches)[index] = BranchTree{
			RawData:    data,
			HashedData: _node.Data,
			Root: &BranchNode{
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
		if prevNode.Branches != nil {
			// fmt.Printf("Accepting right sibling\n")
			for k, v := range *prevNode.Branches {
				v.Root = AcceptRightSibling(v.Root, cNode.Data)
				(*prevNode.Branches)[k] = v
				// fmt.Printf("Proof: %v\n", v.String())
			}
		}
		if cNode.Branches != nil {
			// fmt.Printf("Accepting left sibling\n")
			for k, v := range *cNode.Branches {
				v.Root = AcceptLeftSibling(v.Root, prevNode.Data)
				(*prevNode.Branches)[k] = v
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

func _hash(a, b []byte) common.SPHash {
	result := [common.HashLength]byte{}
	var keccak []byte
	if len(a) == 16 {
		keccak = crypto.Keccak256(
			append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, a...),
			append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, b...),
		)
	} else {
		keccak = crypto.Keccak256(a, b)
	}
	copy(result[:common.HashLength], keccak[common.HashLength:])
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
			dupNode.Branches = &map[uint32]BranchTree{}
			mt.insertNode(dupNode)
			if mt.mtbuf.Len() == 1 {
				break
			}
		}
	}
	mt.finalized = true
}

func (mt SPMerkleTree) Root() common.SPHash {
	if mt.finalized {
		return mt.mtbuf.Front().Value.(node).Data
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (mt SPMerkleTree) Branches() map[uint32]BranchTree {
	if mt.finalized {
		return *(mt.mtbuf.Front().Value.(node).Branches)
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

// return only one array with necessary hashes for each
// index in order. Element's hash and root are not included
// eg. registered indexes are 1, 2, each needs 2 hashes
// then the function return an array of 4 hashes [a1, a2, b1, b2]
// where a1, a2 are proof branch for element at index 1
// b1, b2 are proof branch for element at index 2
func (mt SPMerkleTree) AllBranchesArray() []common.BranchElement {
	if mt.finalized {
		result := []common.BranchElement{}
		branches := mt.Branches()
		for _, k := range mt.orderedIndexes {
			// p := proofs[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			hashes := branches[k].ToHashArray()[1:]
			// fmt.Printf("Len proofs: %s\n", len(pfs))
			for i := 0; i*2 < len(hashes); i++ {
				// for anyone who is courious why i*2 + 1 comes before i * 2
				// it's agreement between client side and contract side
				if i*2+1 >= len(hashes) {
					result = append(result, common.BranchElementFromHash(
						common.SPHash{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, hashes[i*2]))
				} else {
					result = append(result, common.BranchElementFromHash(hashes[i*2+1], hashes[i*2]))
				}
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (mt SPMerkleTree) AllDAGElements() []common.Word {
	if mt.finalized {
		result := []common.Word{}
		branches := mt.Branches()
		for _, k := range mt.orderedIndexes {
			// p := branches[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			result = append(result, branches[k].RawData)
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}
