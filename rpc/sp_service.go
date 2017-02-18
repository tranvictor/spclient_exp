package rpc

import (
	"../claim"
	spcommon "../common"
	"../share"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type SmartPoolService struct{}

func (SmartPoolService) GetWork() ([3]string, error) {
	var res [3]string
	w := Geth.GetWork()
	spcommon.WorkPool[w.PoWHash()] = w
	// w.PrintInfo()
	res[0] = w.PoWHash().Hex()
	res[1] = w.SeedHash()
	n := big.NewInt(1)
	n.Lsh(n, 255)
	n.Div(n, w.ShareDifficulty())
	n.Lsh(n, 1)
	res[2] = common.BytesToHash(n.Bytes()).Hex()
	return res, nil
}

func (SmartPoolService) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	return Geth.SubmitHashrate(hashrate, id)
}

func (SmartPoolService) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	// Make sure the work submitted is present
	work := spcommon.WorkPool[hash]
	if work == nil {
		fmt.Printf("Work was submitted for %x but no pending work found\n", hash)
		return false
	}
	// fmt.Printf("Work submitted with: nonce(%v) mixDigest(%v) hash(%s)\n", nonce, mixDigest, hash.Hex())
	fmt.Printf(".")
	s := share.NewShare(work.BlockHeader(), work.ShareDifficulty())
	s.AcceptSolution(nonce, mixDigest)
	if s.SolutionState == spcommon.FullBlockSolution {
		if Geth.SubmitWork(nonce, hash, mixDigest) {
			fmt.Printf("\n==========YAY found a full solution==========\n")
			delete(spcommon.WorkPool, hash)
			return true
		} else {
			return false
		}
	} else if s.SolutionState == spcommon.ValidShare {
		claim.Repo.AddShare(s)
		return true
	} else {
		return false
	}
	return true
}
