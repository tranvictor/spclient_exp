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
	"github.com/ethereum/go-ethereum/rpc"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sort"
)

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

type testBlock struct {
	difficulty  *big.Int
	hashNoNonce common.Hash
	nonce       uint64
	mixDigest   common.Hash
	number      uint64
}

func (b *testBlock) Difficulty() *big.Int     { return b.difficulty }
func (b *testBlock) HashNoNonce() common.Hash { return b.hashNoNonce }
func (b *testBlock) Nonce() uint64            { return b.nonce }
func (b *testBlock) MixDigest() common.Hash   { return b.mixDigest }
func (b *testBlock) NumberU64() uint64        { return b.number }

func printShare(s share.Share) {
	fmt.Printf("	ParentHash: %s\n", s.BlockHeader().ParentHash.Hex())
	fmt.Printf("	UncleHash: %s\n", s.BlockHeader().UncleHash.Hex())
	fmt.Printf("	Coinbase: %s\n", s.BlockHeader().Coinbase.Hex())
	fmt.Printf("	Root: %s\n", s.BlockHeader().Root.Hex())
	fmt.Printf("	TxHash: %s\n", s.BlockHeader().TxHash.Hex())
	fmt.Printf("	ReceiptHash: %s\n", s.BlockHeader().ReceiptHash.Hex())
	fmt.Printf("	Bloom: %s\n", s.BlockHeader().Bloom)
	fmt.Printf("	Difficulty: 0x%s\n", s.BlockHeader().Difficulty.Text(16))
	fmt.Printf("	Number: %s\n", s.BlockHeader().Number)
	fmt.Printf("	GasLimit: 0x%s\n", s.BlockHeader().GasLimit.Text(16))
	fmt.Printf("	GasUsed: 0x%s\n", s.BlockHeader().GasUsed.Text(16))
	fmt.Printf("	Time: %v\n", s.BlockHeader().Time.Bytes())
	fmt.Printf("	Nonce: %v\n", s.BlockHeader().Nonce[:])
	fmt.Printf("	Extra: %v\n", s.BlockHeader().Extra)
	fmt.Printf("	Counter: %v\n", s.Counter().Bytes())
	fmt.Printf("	Corresponding Min-Max: 0x%s\n", s.Counter().Text(16))
	fmt.Printf("	Corresponding Hash: %s\n", s.Hash().Hex())
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
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	for i := 0; i < 5; i++ {
		s := *getShareFromBlock(client, i+8891)
		claim = append(claim[:], &s)
	}
	amt := mtree.NewAugTree()
	amt.RegisterIndex(2)
	sort.Sort(claim)
	for i, s := range claim[:] {
		fmt.Printf("Share index %d\n", i)
		printShare(*s)
		amt.Insert(*s, uint32(i))
	}
	amt.Finalize()
	root := amt.Root().(mtree.AugData)
	fmt.Printf("Root Hash: %s\n", root.Hash.Hex())
	fmt.Printf("Root Min: 0x%s\n", root.Min.(*big.Int).Text(16))
	fmt.Printf("Root Max: 0x%s\n", root.Max.(*big.Int).Text(16))
	counterArray := amt.CounterBranchArray()
	hashArray := amt.HashBranchArray()
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
}

func testVerifyShare() {
	block := &testBlock{
		// number:      22,
		number:      30000 * 11,
		hashNoNonce: common.HexToHash("372eca2454ead349c3df0ab5d00b0b706b23e49d469387db91811cee0358fc6d"),
		difficulty:  big.NewInt(132416),
		nonce:       0x495732e0ed7a801c,
		mixDigest:   common.HexToHash("2f74cdeb198af0b9abe65d22d372e22fb2d474371774a9583c1cc427a07939f6"),
	}
	eth := ethash.New()
	indices := eth.GetVerificationIndices(block)
	fmt.Printf("Indices: %v\n", indices)
	// getting the dag path
	fmt.Printf("Checking DAG file. Generate if needed...\n")
	ethash.MakeDAG(block.NumberU64(), "")
	seedHash, err := ethash.GetSeedHash(block.NumberU64())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	fmt.Printf("Path: %s\n", path)
	testDatasetMerkleTree(path, indices)
}

func testDatasetMerkleTree(datasetPath string, indices []uint32) {
	mt := mtree.NewDagTree()
	mt.RegisterIndex(indices...)
	processDuringRead(datasetPath, mt)
	mt.Finalize()
	result := mt.Root()
	fmt.Printf("Merkle Root: %s\n", result.(spcommon.SPHash).Hex())
	sproof := share.ShareProof{
		mt.AllDAGElements(),
		mt.AllBranchesArray(),
	}
	fmt.Printf("Element Array: %s\n", sproof.DAGElementArray())
	fmt.Printf("Proof Array: %v\n", sproof.DAGProofArray())
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
	// testDatasetMerkleTree(datasetPath, indicesFromYaron)
	// testVerifyShare()
	testAugMerkleTree()
}
