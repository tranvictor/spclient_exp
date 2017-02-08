package mtree

import (
	"../common"
)

type BranchNode struct {
	Hash             common.SPHash
	Left             *BranchNode
	Right            *BranchNode
	ElementOnTheLeft bool
}

func (b BranchNode) ToHashArray() []common.SPHash {
	if b.Left == nil && b.Right == nil {
		return []common.SPHash{b.Hash}
	}
	left := b.Left.ToHashArray()
	right := b.Right.ToHashArray()
	if b.ElementOnTheLeft {
		return append(left, right...)
	} else {
		return append(right, left...)
	}
}

// explain the operation
func AcceptLeftSibling(b *BranchNode, h common.SPHash) *BranchNode {
	return &BranchNode{
		Hash:             common.SPHash{},
		Left:             &BranchNode{h, nil, nil, false},
		Right:            b,
		ElementOnTheLeft: false,
	}
}

func AcceptRightSibling(b *BranchNode, h common.SPHash) *BranchNode {
	return &BranchNode{
		Hash:             common.SPHash{},
		Right:            &BranchNode{h, nil, nil, false},
		Left:             b,
		ElementOnTheLeft: true,
	}
}
