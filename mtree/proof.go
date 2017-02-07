package mtree

type proof struct {
	RawData    Word
	HashedData SPHash
	Branch     *Branch
}

func (p proof) String() string {
	return p.Branch.InOrderTraversal()
}

func (p proof) ToProofArray() []SPHash {
	return p.Branch.ToProofArray()
}
