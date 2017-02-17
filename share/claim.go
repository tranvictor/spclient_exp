package share

import (
	"fmt"
)

type Claim []*Share

func (c Claim) Len() int      { return len(c) }
func (c Claim) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Claim) Less(i, j int) bool {
	return c[i].Counter().Cmp(c[j].Counter()) == -1
}

var ClaimRepo = LoadClaimRepo()

type claimrepo struct {
	claims         []Claim
	cIndex         uint64
	shareThreshold uint64
}

func (cr *claimrepo) AddShare(s *Share) {
	currentClaim := cr.CurrentClaim()
	if uint64(len(currentClaim)) >= cr.shareThreshold {
		fmt.Printf("================\n")
		fmt.Printf("  Got enough shares to construct augmented merkle tree and submit to contract\n")
		fmt.Printf("  Start new claim\n")
		fmt.Printf("================\n")
		cr.cIndex++
		cr.claims = append(cr.claims, Claim{})
	}
	cr.claims[cr.cIndex] = append(cr.claims[cr.cIndex][:], s)
}

func (cr claimrepo) CurrentClaim() Claim {
	return cr.claims[cr.cIndex]
}

func LoadClaimRepo() *claimrepo {
	// TODO: load from persistent storage
	return &claimrepo{[]Claim{Claim{}}, 0, 16}
}
