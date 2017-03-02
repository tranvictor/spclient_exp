package main

import (
	"./claim"
	"./client"
	spcommon "./common"
	"./contract"
	"./ethash"
	"./mtree"
	"./params"
	"./server"
	"./share"
	"./txs"
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"time"
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
	AugTreeCounterBranch []*big.Int
	AugTreeHashBranch    []*big.Int
	EthashCacheRoot      spcommon.SPHash
	CacheElements        []*big.Int
	CacheBranch          []*big.Int
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

func buildExtraData(address common.Address, diff *big.Int) string {
	// TODO: get default address from local environment
	// id = address % (26+26+10)**11
	base := big.NewInt(0)
	base.Exp(big.NewInt(62), big.NewInt(11), nil)
	id := big.NewInt(0)
	id.Mod(address.Big(), base)
	return fmt.Sprintf("SmartPool-%s%s", spcommon.BigToBase62(id), spcommon.BigToBase62(diff))
}

func Initialize() bool {
	// Setting
	params.NoSharePerClaim = uint32(13)
	params.ShareDifficulty = big.NewInt(100000)
	params.SubmitInterval = 1 * time.Minute
	params.ContractAddress = "0x9e7a1925fa43d5f47b36e2e27f84adae95ddd845"
	// TODO: Need better way to get ipc file
	// params.IPCPath = "/Users/victor/Library/Ethereum/testnet/geth.ipc"
	params.IPCPath = "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/geth.ipc"
	// TODO: Need better way to get keystore path
	// params.KeystorePath = "/Users/victor/Library/Ethereum/testnet/keystore"
	params.KeystorePath = "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/keystore"
	// TODO: Need better way to get default address for miner
	params.MinerAddress = "0xad42beeb07db31149f5d2c4bd33d01c6d7c34116"
	address := common.HexToAddress(params.MinerAddress)
	params.ExtraData = buildExtraData(address, big.NewInt(100000))

	// Share instances
	var err error
	client.DefaultGethClient, err = client.NewGethRPCClient()
	if err != nil {
		fmt.Printf("Geth RPC server is unavailable.\n")
		fmt.Printf("Make sure you have Geth installed. If you do, you can run geth by following command (Note: --etherbase and --extradata params are required.):\n")
		fmt.Printf(
			"geth --rpc --etherbase \"%s\" --extradata \"%s\"\n",
			params.ContractAddress, params.ExtraData)
		return false
	}
	server.DefaultServer = server.NewRPCServer()
	contract.DefaultContractClient, err = contract.NewContractClient()
	if err != nil {
		fmt.Printf("Geth RPC server is unavailable.\n")
		fmt.Printf("Make sure you have Geth installed. If you do, you can run geth by following command:\n")
		fmt.Printf(
			"geth --rpc --etherbase \"%s\" --extradata \"%s\"\n",
			params.ContractAddress, params.ExtraData)
		return false
	}
	claim.DefaultClaimRepo = claim.LoadClaimRepo(contract.DefaultContractClient)
	// TODO: check current geth setup to see if coinbase address
	// and extradata is set properly
	return registerToPool(address)
}

func registerToPool(address common.Address) bool {
	if !contract.DefaultContractClient.IsRegistered() {
		if contract.DefaultContractClient.CanRegister() {
			tx, err := contract.DefaultContractClient.Register(address)
			if err != nil {
				fmt.Printf("Unable to register to the pool: %s\n", err)
				return false
			}
			fmt.Printf("Registering to the pool. Please wait...")
			txs.NewTxWatcher(tx).Wait()
			if !contract.DefaultContractClient.IsRegistered() {
				fmt.Printf("Unable to register to the pool. You might try again.")
				return false
			}
			fmt.Printf("Done.\n")
			return true
		} else {
			fmt.Printf("Your etherbase address couldn't register to the pool. You need to try another address.\n")
			return false
		}
	}
	fmt.Printf("The address is already registered to the pool. Good to go.\n")
	return true
}

func testAugMerkleTree() {
	input := InputForYaron{}
	input.NumShare = 1
	input.ShareIndex = 0
	cl := claim.Claim{}
	for i := 0; i < input.NumShare; i++ {
		h := client.DefaultGethClient.GetBlockHeader(3141592 - i)
		s := share.NewShare(h, h.Difficulty)
		cl = append(cl[:], s)
	}
	amt := mtree.NewAugTree()
	requestedIndex := uint32(input.ShareIndex)
	amt.RegisterIndex(requestedIndex)
	var requestedShare *share.Share
	sort.Sort(cl)
	for i, s := range cl[:] {
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
		fmt.Printf("0x%s, ", c.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("Hash Array: [")
	for _, h := range hashArray {
		fmt.Printf("0x%s, ", h.Text(16))
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
		fmt.Printf("0x%s, ", e.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("aug_tree_counters_branch = [")
	for _, c := range input.AugTreeCounterBranch {
		fmt.Printf("0x%s, ", c.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("aug_tree_hashes_branch = [")
	for _, h := range input.AugTreeHashBranch {
		fmt.Printf("0x%s, ", h.Text(16))
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
		fmt.Printf("0x%s, ", e.Text(16))
	}
	fmt.Printf("]\n")
}

func testGetWork() {
	w := client.DefaultGethClient.GetWork()
	w.PrintInfo()
}

func testRPCServer() {
	server := server.NewRPCServer()
	server.Start()
}

func testInteractWithContract() {
	server.DefaultServer.Start()
}

func testExtraData() {
	updaterClient := contract.NewUpdaterClient()
	client.DefaultGethClient.GetWork()
	pendingBlock := client.DefaultGethClient.GetPendingBlockHeader()
	diff := big.NewInt(100000)
	extraData := pendingBlock.Extra
	minerAddress := common.HexToAddress(params.MinerAddress)
	base := big.NewInt(0)
	base.Exp(big.NewInt(62), big.NewInt(11), nil)
	id := big.NewInt(0)
	id.Mod(minerAddress.Big(), base)
	encodedID := spcommon.BigToBase62(id)
	encodedDiff := spcommon.BigToBase62(diff)
	fmt.Printf("minerID: %v\n", []byte(encodedID))
	fmt.Printf("diff: %v\n", []byte(encodedDiff))
	extra32 := [32]byte{}
	id32 := [32]byte{}
	copy(extra32[:], extraData[:])
	fmt.Printf("extra32: %v\n", extra32)
	copy(id32[21:], []byte(encodedID)[:])
	ok, _ := updaterClient.VerifyExtraData(extra32, id32, diff)
	fmt.Printf("Checking extra data: %s\n", ok)
	encoded, _ := updaterClient.To62Encoding(diff, big.NewInt(11))
	fmt.Printf("Diff encoded by contract: %v\n", encoded)
}

func testSubmitEpochInfo(blockNumber uint64) {
	updaterClient := contract.NewUpdaterClient()
	fmt.Printf("Block number: %d\n", blockNumber)
	fmt.Printf("Checking DAG file. Generate if needed...\n")
	fullSize, _ := ethash.MakeDAGWithSize(blockNumber, "")
	fullSizeIn128Resolution := fullSize / 128
	seedHash, err := ethash.GetSeedHash(blockNumber)
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	mt := mtree.NewDagTree()
	processDuringRead(path, mt)
	mt.Finalize()
	merkleRoot := mt.RootHash()
	epoch := int64(blockNumber) / 30000
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	tx, err := updaterClient.SetEpochData(
		merkleRoot.Big(),
		fullSizeIn128Resolution,
		uint64(branchDepth),
		big.NewInt(epoch),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	txs.NewTxWatcher(tx).Wait()
	fmt.Printf("Verified: 0x%x\n", tx.Hash())
}

func main() {
	if !Initialize() {
		return
	}
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
	// testAugMerkleTree()
	// testGetWork()
	// testRPCServer()
	testInteractWithContract()
	// testExtraData()
	// testSubmitEpochInfo(20)
}
