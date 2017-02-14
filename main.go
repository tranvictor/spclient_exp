package main

import (
	spcommon "./common"
	"./ethash"
	"./mtree"
	"./share"
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sort"
)

type InputForYaron struct {
	NumShare             int
	ShareIndex           int
	RlpHeader            []byte
	Difficulty           uint64
	Epoch                uint64
	Nonce                types.BlockNonce
	AugMerkleRoot        spcommon.SPHash
	MinCounter           *big.Int
	MaxCounter           *big.Int
	AugTreeCounterBranch []spcommon.BranchElement
	AugTreeHashBranch    []spcommon.BranchElement
	EthashCacheRoot      spcommon.SPHash
	CacheElements        []big.Int
	CacheBranch          []spcommon.BranchElement
	CacheNumberOfElement uint64
	BranchDepth          int
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
		mt.Insert(spcommon.Word(buf), i)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		i++
	}
}

func getShareFromBlock(client *rpc.Client, number int) *share.Share {
	s := share.NewShare()
	err := client.Call(s.BlockHeader(), "eth_getBlockByNumber", number, false)
	if err != nil {
		log.Fatal("Couldn't get latest block:", err)
	}
	return s
}

func testAugMerkleTree() {
	claim := share.Claim{}
	input := InputForYaron{}
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	input.NumShare = 8
	input.ShareIndex = 0
	for i := 0; i < 8; i++ {
		s := *getShareFromBlock(client, 1206-i)
		claim = append(claim[:], &s)
	}
	amt := mtree.NewAugTree()
	requestedIndex := uint32(input.ShareIndex)
	amt.RegisterIndex(requestedIndex)
	var requestedShare *share.Share
	sort.Sort(claim)
	for i, s := range claim[:] {
		if uint32(i) == requestedIndex {
			fmt.Printf("Share index %d\n", i)
			s.PrintInfo()
			requestedShare = s
			_rlpHeader, _ := s.RlpHeaderWithoutNonce()
			input.RlpHeader = _rlpHeader
			input.Nonce = s.BlockHeader().Nonce
			input.Difficulty = s.BlockHeader().Difficulty.Uint64()
			input.Epoch = s.BlockHeader().Number.Uint64() / 30000
		}
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	root := amt.Root().(mtree.AugData)
	input.AugMerkleRoot = root.Hash
	input.MinCounter = root.Min.(*big.Int)
	input.MaxCounter = root.Max.(*big.Int)
	fmt.Printf("Root Hash: %s\n", root.Hash.Hex())
	fmt.Printf("Root Min: 0x%s\n", root.Min.(*big.Int).Text(16))
	fmt.Printf("Root Max: 0x%s\n", root.Max.(*big.Int).Text(16))
	counterArray := amt.CounterBranchArray()
	input.AugTreeCounterBranch = counterArray
	hashArray := amt.HashBranchArray()
	input.AugTreeHashBranch = hashArray
	fmt.Printf("Counter Array: [")
	for _, c := range counterArray {
		fmt.Printf("%s, ", c.Hex())
	}
	fmt.Printf("]\n")
	fmt.Printf("Hash Array: [")
	for _, h := range hashArray {
		fmt.Printf("%s, ", h.Hex())
	}
	fmt.Printf("]\n")
	testVerifyShare(requestedShare, &input)
	fmt.Printf("// ==========================\n")
	fmt.Printf("epoch_params = [%s, %d, %d, %d]\n",
		input.EthashCacheRoot.Hex(),
		input.CacheNumberOfElement,
		input.BranchDepth,
		input.Epoch,
	)
	fmt.Printf("submit_claim_params = [%d, %d, 0x%s, 0x%s, %s]\n",
		input.NumShare,
		input.Difficulty,
		input.MinCounter.Text(16),
		input.MaxCounter.Text(16),
		input.AugMerkleRoot.Hex(),
	)
	fmt.Printf("cache_elements = [")
	for _, bint := range input.CacheElements {
		fmt.Printf("0x%s, ", bint.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("cache_branch = [")
	for _, e := range input.CacheBranch {
		fmt.Printf("%s, ", e.Hex())
	}
	fmt.Printf("]\n")
	fmt.Printf("aug_tree_counters_branch = [")
	for _, c := range input.AugTreeCounterBranch {
		fmt.Printf("%s, ", c.Hex())
	}
	fmt.Printf("]\n")
	fmt.Printf("aug_tree_hashes_branch = [")
	for _, h := range input.AugTreeHashBranch {
		fmt.Printf("%s, ", h.Hex())
	}
	fmt.Printf("]\n")
	fmt.Printf("rlp_header = [")
	for _, b := range input.RlpHeader {
		fmt.Printf("%d, ", b)
	}
	fmt.Printf("]\n")
	fmt.Printf("verify_claim_params = [bytes(bytearray(rlp_header)), 0x%s, %d, cache_elements, cache_branch, aug_tree_counters_branch, aug_tree_hashes_branch]\n",
		hex.EncodeToString(input.Nonce[:]),
		input.ShareIndex,
	)
	fmt.Printf("// ==========================\n")
}

type testBlock struct {
	difficulty  *big.Int
	hashNoNonce common.Hash
	nonce       types.BlockNonce
	mixDigest   common.Hash
	number      *big.Int
}

func (b *testBlock) Difficulty() *big.Int     { return b.difficulty }
func (b *testBlock) HashNoNonce() common.Hash { return b.hashNoNonce }
func (b *testBlock) Nonce() uint64            { return b.nonce.Uint64() }
func (b *testBlock) MixDigest() common.Hash   { return b.mixDigest }
func (b *testBlock) NumberU64() uint64        { return b.number.Uint64() }

func testVerifyShare(s *share.Share, input *InputForYaron) {
	block := &testBlock{
		// number:      22,
		number:      s.BlockHeader().Number,
		hashNoNonce: s.BlockHeader().HashNoNonce(),
		difficulty:  s.BlockHeader().Difficulty,
		nonce:       s.BlockHeader().Nonce,
		mixDigest:   s.BlockHeader().MixDigest,
	}
	eth := ethash.New()
	indices := eth.GetVerificationIndices(block)
	fmt.Printf("Indices: %v\n", indices)
	// getting the dag path
	fmt.Printf("Block number: %d\n", block.NumberU64())
	fmt.Printf("Checking DAG file. Generate if needed...\n")
	fullSize, _ := ethash.MakeDAGWithSize(block.NumberU64(), "")
	input.CacheNumberOfElement = fullSize / 128
	seedHash, err := ethash.GetSeedHash(block.NumberU64())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	fmt.Printf("Path: %s\n", path)
	testDatasetMerkleTree(path, indices, input)
}

func testDatasetMerkleTree(datasetPath string, indices []uint32, input *InputForYaron) {
	mt := mtree.NewDagTree()
	mt.RegisterIndex(indices...)
	processDuringRead(datasetPath, mt)
	mt.Finalize()
	result := mt.Root()
	input.EthashCacheRoot = spcommon.SPHash(result.(mtree.DagData))
	fmt.Printf("Dag Merkle Root: %s\n", spcommon.SPHash(result.(mtree.DagData)).Hex())
	sproof := share.ShareProof{
		mt.AllDAGElements(),
		mt.AllBranchesArray(),
	}
	input.CacheElements = sproof.DAGElementArray()
	input.CacheBranch = sproof.DAGProofArray()
	input.BranchDepth = 23
	fmt.Printf("Element Array: [")
	for _, bint := range sproof.DAGElementArray() {
		fmt.Printf("0x%s, ", bint.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("Proof Array: [")
	for _, e := range sproof.DAGProofArray() {
		fmt.Printf("%s, ", e.Hex())
	}
	fmt.Printf("]\n")
}

func main() {
	// compute merkle root of dataset
	// datasetPath := "/Users/victor/.ethash/test"
	// datasetPath := "/Users/victor/.ethash/full-R23-afeb5e4f7c8312e3"
	// indicesFromYaron := []uint32{
	// 	13282552, 1105031, 11812463, 2790415,
	// 	2625720, 4539816, 5187220, 7735247, 12827669, 8220447,
	// 	3771673, 6107320, 4322584, 499202, 9249127, 10483756,
	// 	3398027, 3569374, 9182293, 3054465, 12067048, 5155926,
	// 	12645521, 10530848, 434740, 8209194, 10983812, 10821517,
	// 	2058423, 4629979, 11416915, 8357745, 1421006, 5312874,
	// 	9603835, 1436343, 10252321, 6548335, 5237163, 8705311,
	// 	4940987, 374080, 1865848, 2998453, 12031173, 2455677,
	// 	3294052, 11569114, 4610178, 7289900, 8507270, 1839564,
	// 	5626595, 5680798, 12309, 6314194, 11400756, 3646046,
	// 	552207, 1118353, 12823889, 11905227, 7079429, 3667145,
	// }
	// input := InputForYaron{}
	// testDatasetMerkleTree(datasetPath, indicesFromYaron, &input)
	// testVerifyShare()
	testAugMerkleTree()
}
