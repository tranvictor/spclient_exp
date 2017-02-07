package mtree

import (
	"encoding/hex"
	"fmt"
)

type Branch struct {
	Hash             SPHash
	Left             *Branch
	Right            *Branch
	ElementOnTheLeft bool
}

func (b Branch) ToProofArray() []SPHash {
	if b.Left == nil && b.Right == nil {
		return []SPHash{b.Hash}
	}
	left := b.Left.ToProofArray()
	right := b.Right.ToProofArray()
	if b.ElementOnTheLeft {
		return append(left, right...)
	} else {
		return append(right, left...)
	}
}

func (b Branch) InOrderTraversal() string {
	if b.Left == nil && b.Right == nil {
		return hex.EncodeToString(b.Hash[:])
	}
	return fmt.Sprintf(
		"(%s,%s)",
		b.Left.InOrderTraversal(),
		b.Right.InOrderTraversal(),
	)
}

// explain the operation
func AcceptLeftSibling(b *Branch, h SPHash) *Branch {
	return &Branch{
		Hash:             SPHash{},
		Left:             &Branch{h, nil, nil, false},
		Right:            b,
		ElementOnTheLeft: false,
	}
}

func AcceptRightSibling(b *Branch, h SPHash) *Branch {
	return &Branch{
		Hash:             SPHash{},
		Right:            &Branch{h, nil, nil, false},
		Left:             b,
		ElementOnTheLeft: true,
	}
}
