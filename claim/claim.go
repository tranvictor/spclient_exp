package claim

import (
	"../common"
	"../contract"
	"../ethash"
	"../mtree"
	"../share"
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
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

func processDuringRead(
	datasetPath string, mt *mtree.DagTree) {

	f, err := os.Open(datasetPath)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	buf := [128]byte{}
	// ignore first 8 bytes magic number at the beginning
	// of dataset. See more at https://github.com/ethereum/wiki/wiki/Ethash-DAG-Disk-Storage-Format
	_, err = io.ReadFull(r, buf[:8])
	if err != nil {
		log.Fatal(err)
	}
	var i uint32 = 0
	for {
		n, err := io.ReadFull(r, buf[:128])
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n != 128 {
			fmt.Println(n)
			log.Fatal("Malformed dataset")
		}
		mt.Insert(common.Word(buf), i)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		i++
	}
}

// TODO: remove this
func (c *Claim) SubmitProof_debug(_client *contract.ContractClient, index int) (*big.Int, error) {
	sort.Sort(c)
	amt := mtree.NewAugTree()
	amt.RegisterIndex(uint32(index))
	for i, s := range *c {
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	requestedShare := (*c)[index]
	rlpHeader, _ := requestedShare.RlpHeaderWithoutNonce()
	nonce := requestedShare.NonceBig()
	shareIndex := big.NewInt(int64(index))
	augCountersBranch := amt.CounterBranchArray()
	augHashesBranch := amt.HashBranchArray()

	eth := ethash.New()
	indices := eth.GetVerificationIndices(requestedShare)
	seedHash, err := ethash.GetSeedHash(requestedShare.NumberU64())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	mt := mtree.NewDagTree()
	mt.RegisterIndex(indices...)
	processDuringRead(path, mt)
	mt.Finalize()
	sproof := share.ShareProof{
		mt.AllDAGElements(),
		mt.AllBranchesArray(),
	}
	dataSetLookup := sproof.DAGElementArray()
	witnessForLookup := sproof.DAGProofArray()
	return _client.VerifyClaim_debug(
		rlpHeader,
		nonce,
		shareIndex,
		dataSetLookup,
		witnessForLookup,
		augCountersBranch,
		augHashesBranch,
	)
}

// TODO: should break this function into smaller meaningful ones
func (c *Claim) SubmitProof(_client *contract.ContractClient, index int) (*types.Transaction, error) {
	sort.Sort(c)
	amt := mtree.NewAugTree()
	amt.RegisterIndex(uint32(index))
	for i, s := range *c {
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	requestedShare := (*c)[index]
	rlpHeader, _ := requestedShare.RlpHeaderWithoutNonce()
	nonce := requestedShare.NonceBig()
	shareIndex := big.NewInt(int64(index))
	augCountersBranch := amt.CounterBranchArray()
	augHashesBranch := amt.HashBranchArray()

	eth := ethash.New()
	indices := eth.GetVerificationIndices(requestedShare)
	seedHash, err := ethash.GetSeedHash(requestedShare.NumberU64())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	mt := mtree.NewDagTree()
	mt.RegisterIndex(indices...)
	processDuringRead(path, mt)
	mt.Finalize()
	sproof := share.ShareProof{
		mt.AllDAGElements(),
		mt.AllBranchesArray(),
	}
	dataSetLookup := sproof.DAGElementArray()
	witnessForLookup := sproof.DAGProofArray()
	return _client.VerifyClaim(
		rlpHeader,
		nonce,
		shareIndex,
		dataSetLookup,
		witnessForLookup,
		augCountersBranch,
		augHashesBranch,
	)
}

func (c *Claim) SubmitToContract(_client *contract.ContractClient) (*types.Transaction, error) {
	sort.Sort(c)
	amt := mtree.NewAugTree()
	for i, s := range *c {
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	fmt.Printf("  Submitting %d shares to contract\n", len(*c))
	return _client.SubmitClaim(
		big.NewInt(int64(len(*c))),
		c.MinDifficulty(),
		amt.RootMin(),
		amt.RootMax(),
		amt.RootHash().Big(),
	)
}
