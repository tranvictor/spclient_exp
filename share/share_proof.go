package share

import (
	spcommon "../common"
	"math/big"
)

type ShareProof struct {
	DAGElements []spcommon.Word
	DAGProof    []spcommon.BranchElement
}

func (p ShareProof) DAGProofArray() []spcommon.BranchElement {
	return p.DAGProof
}

func (p ShareProof) DAGElementArray() []big.Int {
	result := []big.Int{}
	for _, w := range p.DAGElements {
		result = append(result, w.ToUint256Array()...)
	}
	return result
}
