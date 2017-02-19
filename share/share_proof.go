package share

import (
	spcommon "../common"
	"math/big"
)

type ShareProof struct {
	DAGElements []spcommon.Word
	DAGProof    []spcommon.BranchElement
}

func (p ShareProof) DAGProofArray() []*big.Int {
	result := []*big.Int{}
	for _, be := range p.DAGProof {
		result = append(result, be.Big())
	}
	return result
}

func (p ShareProof) DAGElementArray() []*big.Int {
	result := []*big.Int{}
	for _, w := range p.DAGElements {
		result = append(result, w.ToUint256Array()...)
	}
	return result
}
