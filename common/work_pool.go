package common

import (
	"github.com/ethereum/go-ethereum/common"
)

type workpool map[common.Hash]*Work

var WorkPool = workpool{}

const (
	FullBlockSolution int = 2
	ValidShare        int = 1
	InvalidShare      int = 0
)
