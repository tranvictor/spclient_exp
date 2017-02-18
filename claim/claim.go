package claim

import (
	"../contract"
	"../mtree"
	"../share"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sort"
)

type Claim []*share.Share

func (c Claim) Len() int      { return len(c) }
func (c Claim) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Claim) Less(i, j int) bool {
	return c[i].Counter().Cmp(c[j].Counter()) == -1
}

func (c Claim) MinDifficulty() *big.Int {
	var m *big.Int
	if len(c) > 0 {
		m = c[0].ShareDifficulty
		for _, s := range c {
			if m.Cmp(s.ShareDifficulty) >= 0 {
				m = s.ShareDifficulty
			}
		}
	}
	return m
}

func (c *Claim) SubmitToContract(_client *contract.ContractClient) (*types.Transaction, error) {
	sort.Sort(c)
	amt := mtree.NewAugTree()
	for i, s := range *c {
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	return _client.SubmitClaim(
		big.NewInt(int64(len(*c))),
		c.MinDifficulty(),
		amt.RootMin(),
		amt.RootMax(),
		amt.RootHash().Big(),
	)
}

var Repo *ClaimRepo

type ClaimRepo struct {
	claims         []Claim
	cIndex         uint64
	shareThreshold uint64
	contract       *contract.ContractClient
}

func (cr *ClaimRepo) AddShare(s *share.Share) {
	currentClaim := cr.CurrentClaim()
	if uint64(len(currentClaim)) >= cr.shareThreshold {
		fmt.Printf("================\n")
		fmt.Printf("  Got enough shares to construct augmented merkle tree and submit to contract\n")
		tx, err := currentClaim.SubmitToContract(cr.contract)
		fmt.Printf("  Submitted by pending tx: 0x%x\n", tx.Hash())
		if err != nil {
			panic(err)
		}
		fmt.Printf("  Start new claim\n")
		fmt.Printf("================\n")
		cr.cIndex++
		cr.claims = append(cr.claims, Claim{})
	}
	cr.claims[cr.cIndex] = append(cr.claims[cr.cIndex][:], s)
}

func (cr ClaimRepo) CurrentClaim() Claim {
	return cr.claims[cr.cIndex]
}

func LoadClaimRepo(cc *contract.ContractClient) *ClaimRepo {
	// TODO: load from persistent storage
	// TODO: this is currently not safe for multiple go routines
	if Repo == nil {
		Repo = &ClaimRepo{[]Claim{Claim{}}, 0, 16, cc}
	}
	return Repo
}
