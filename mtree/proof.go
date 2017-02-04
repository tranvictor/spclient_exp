package mtree

import (
	"github.com/ethereum/go-ethereum/common"
)

type proof struct {
	RawData    Word
	HashedData common.Hash
	Branch     *Branch
}

func (p proof) String() string {
	return p.Branch.InOrderTraversal()
}

func (p proof) ToArray() []common.Hash {
	return p.Branch.ToProofArray()
}
