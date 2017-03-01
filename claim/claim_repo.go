package claim

import (
	"../contract"
	"../params"
	"../share"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

var DefaultClaimRepo *ClaimRepo

type ClaimRepo struct {
	claims         map[int]Claim
	cClaimNumber   uint64
	shareThreshold uint64
	watcherStarted bool
	ticker         <-chan time.Time
	contract       *contract.ContractClient
}

func LoadClaimRepo(cc *contract.ContractClient) *ClaimRepo {
	// TODO: load from persistent storage
	// TODO: this is currently not safe for multiple go routines
	repo := &ClaimRepo{
		map[int]Claim{0: Claim{}},
		0,
		13,
		false,
		time.Tick(params.SubmitInterval),
		cc,
	}
	repo.StartWatcher()
	return repo
}

func (cr *ClaimRepo) actOnTick() {
	for t := range cr.ticker {
		currentClaim := cr.CurrentClaim()
		if uint64(len(currentClaim)) >= cr.shareThreshold {
			fmt.Printf("\n================\n")
			fmt.Printf("  It's time (%s) to collect submitted shares to construct augmented merkle tree and submit to contract\n", t)
			tx, err := currentClaim.SubmitToContract(cr.contract)
			if err != nil {
				panic(err)
			}
			fmt.Printf("  Submitted by pending tx: 0x%x\n", tx.Hash())
			fmt.Printf("  Starting new claim\n")
			fmt.Printf("  Getting claim number: ")
			cr.cClaimNumber = cr.NextClaimNumber()
			fmt.Printf("%d\n", cr.cClaimNumber)
			cr.claims[int(cr.cClaimNumber)] = Claim{}
			tx, err = cr.VerifyClaim()
			if err != nil {
				panic(err)
			}
			if tx != nil {
				fmt.Printf("  Verification submitted by pending tx: 0x%x\n", tx.Hash())
			}
			fmt.Printf("================\n")
		}
	}
}

func (cr *ClaimRepo) StartWatcher() {
	if cr.watcherStarted {
		fmt.Printf("Warning: calling ClaimRepo.StatWatcher multiple times\n")
		return
	}
	go cr.actOnTick()
	cr.watcherStarted = true
}

func (cr *ClaimRepo) AddShare(s *share.Share) {
	cr.claims[int(cr.cClaimNumber)] = append(cr.claims[int(cr.cClaimNumber)][:], s)
}

func (cr *ClaimRepo) GetClaim(number int) Claim {
	return cr.claims[number]
}

func (cr *ClaimRepo) getClaimToVerify() Claim {
	// TODO: get the oldest unverified claim
	// right now, we just get the claim that has
	// claim number of current number - 3
	return cr.GetClaim(int(cr.cClaimNumber) - 1)
}

func (cr *ClaimRepo) VerifyClaim() (*types.Transaction, error) {
	// TODO: Get seed from contract
	index := 8
	claim := cr.getClaimToVerify()
	if claim != nil {
		return claim.SubmitProof(cr.contract, index)
	} else {
		return nil, nil
	}
}

func (cr *ClaimRepo) NextClaimNumber() uint64 {
	return cr.cClaimNumber + 1
}

func (cr *ClaimRepo) CurrentClaim() Claim {
	return cr.claims[int(cr.cClaimNumber)]
}
