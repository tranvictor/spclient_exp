package main

import (
	"./ethash"
	"./mtree"
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"math/big"
	"os"
)

func processDuringRead(
	datasetPath string, mt *mtree.SPMerkleTree) {

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
		mt.Insert(mtree.Word(buf), i)
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

func testVerifyShare() {
	block := &testBlock{
		number:      60000,
		hashNoNonce: common.HexToHash("372eca2454ead349c3df0ab5d00b0b706b23e49d469387db91811cee0358fc6d"),
		difficulty:  big.NewInt(132416),
		nonce:       0x495732e0ed7a801c,
		mixDigest:   common.HexToHash("2f74cdeb198af0b9abe65d22d372e22fb2d474371774a9583c1cc427a07939f5"),
	}
	eth := ethash.New()
	fmt.Printf("%s\n", eth.Verify(block))
}

func testDatasetMerkleTree(datasetPath string) {
	mt := mtree.NewSPMerkleTree()
	mt.RegisterIndex(13282552,
		1105031, 11812463, 2790415,
		2625720, 4539816, 5187220, 7735247, 12827669, 8220447,
		3771673, 6107320, 4322584, 499202, 9249127, 10483756,
		3398027, 3569374, 9182293, 3054465, 12067048, 5155926,
		12645521, 10530848, 434740, 8209194, 10983812, 10821517,
		2058423, 4629979, 11416915, 8357745, 1421006, 5312874,
		9603835, 1436343, 10252321, 6548335, 5237163, 8705311,
		4940987, 374080, 1865848, 2998453, 12031173, 2455677,
		3294052, 11569114, 4610178, 7289900, 8507270, 1839564,
		5626595, 5680798, 12309, 6314194, 11400756, 3646046,
		552207, 1118353, 12823889, 11905227, 7079429, 3667145)
	processDuringRead(datasetPath, mt)
	mt.Finalize()
	result := mt.Root()
	fmt.Printf("Merkle Root: %s\n", result.Hex())
	// get proof for list of indexes
	fmt.Printf("[")
	for _, k := range mt.ProofArray() {
		fmt.Printf("\"%s\", ", k.Hex())
	}
	fmt.Printf("]\n")
}

func main() {
	// compute merkle root of dataset
	// datasetPath := "/Users/victor/.ethash/full-R23-afeb5e4f7c8312e3"
	// datasetPath := "/Users/victor/.ethash/test"
	// testDatasetMerkleTree(datasetPath)
	testVerifyShare()
}
