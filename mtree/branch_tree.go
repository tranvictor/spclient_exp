package mtree

import "../common"

type BranchTree struct {
	RawData    common.Word
	HashedData common.SPHash
	Root       *BranchNode
}

func (t BranchTree) ToHashArray() []common.SPHash {
	return t.Root.ToHashArray()
}
