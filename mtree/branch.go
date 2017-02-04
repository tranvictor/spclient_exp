package mtree

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

type Branch struct {
	Hash             common.Hash
	Left             *Branch
	Right            *Branch
	ElementOnTheLeft bool
}

func (b Branch) ToProofArray() []common.Hash {
	if b.Left == nil && b.Right == nil {
		return []common.Hash{b.Hash}
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
func AcceptLeftSibling(b *Branch, h common.Hash) *Branch {
	return &Branch{
		Hash:             common.Hash{},
		Left:             &Branch{h, nil, nil, false},
		Right:            b,
		ElementOnTheLeft: false,
	}
}

func AcceptRightSibling(b *Branch, h common.Hash) *Branch {
	return &Branch{
		Hash:             common.Hash{},
		Right:            &Branch{h, nil, nil, false},
		Left:             b,
		ElementOnTheLeft: true,
	}
}
