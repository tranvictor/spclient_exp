package mtree

import (
	"../common"
	"container/list"
	"github.com/ethereum/go-ethereum/crypto"
)

type DagData common.SPHash

func (dd DagData) Copy() NodeData {
	result := DagData{}
	copy(result[:], dd[:])
	return result
}

type DagTree struct {
	MerkleTree
}

func _elementHash(data ElementData) NodeData {
	// insert data into the mtbuf and aggregate the
	// hashes
	// because contract side is expecting the bytes
	// to be reversed each 32 bytes on leaf nodes
	first, second := conventionalWord(data.(common.Word))
	keccak := crypto.Keccak256(first, second)
	result := DagData{}
	copy(result[:common.HashLength], keccak[common.HashLength:])
	return result
}

func _hash(a, b NodeData) NodeData {
	var keccak []byte
	left := a.(DagData)
	right := b.(DagData)
	keccak = crypto.Keccak256(
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, left[:]...),
		append([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, right[:]...),
	)
	result := DagData{}
	copy(result[:common.HashLength], keccak[common.HashLength:])
	return result
}

func _modifier(data NodeData) {}

func NewDagTree() *DagTree {
	mtbuf := list.New()
	return &DagTree{
		MerkleTree{
			mtbuf,
			_hash,
			_elementHash,
			_modifier,
			false,
			map[uint32]bool{},
			[]uint32{},
		},
	}
}

func (dt DagTree) RootHash() common.SPHash {
	if dt.finalized {
		return common.SPHash(dt.Root().(DagData))
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

// return only one array with necessary hashes for each
// index in order. Element's hash and root are not included
// eg. registered indexes are 1, 2, each needs 2 hashes
// then the function return an array of 4 hashes [a1, a2, b1, b2]
// where a1, a2 are proof branch for element at index 1
// b1, b2 are proof branch for element at index 2
func (dt DagTree) AllBranchesArray() []common.BranchElement {
	if dt.finalized {
		result := []common.BranchElement{}
		branches := dt.Branches()
		for _, k := range dt.Indices() {
			// p := proofs[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			hashes := branches[k].ToNodeArray()[1:]
			// fmt.Printf("Len proofs: %s\n", len(pfs))
			for i := 0; i*2 < len(hashes); i++ {
				// for anyone who is courious why i*2 + 1 comes before i * 2
				// it's agreement between client side and contract side
				if i*2+1 >= len(hashes) {
					result = append(result,
						common.BranchElementFromHash(
							common.SPHash(DagData{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
							common.SPHash(hashes[i*2].(DagData))))
				} else {
					result = append(result,
						common.BranchElementFromHash(
							common.SPHash(hashes[i*2+1].(DagData)),
							common.SPHash(hashes[i*2].(DagData))))
				}
			}
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}

func (dt DagTree) AllDAGElements() []common.Word {
	if dt.finalized {
		result := []common.Word{}
		branches := dt.Branches()
		for _, k := range dt.Indices() {
			// p := branches[k]
			// fmt.Printf("Index: %d\nRawData: %s\nHashedData: %s\n", k, hex.EncodeToString(p.RawData[:]), proofs[k].HashedData.Hex())
			result = append(result, branches[k].RawData.(common.Word))
		}
		return result
	}
	panic("SP Merkle tree needs to be finalized by calling mt.Finalize()")
}
