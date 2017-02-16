package rpc

import (
	spcommon "../common"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"time"
)

type jsonHeader struct {
	ParentHash  *common.Hash      `json:"parentHash"`
	UncleHash   *common.Hash      `json:"sha3Uncles"`
	Coinbase    *common.Address   `json:"miner"`
	Root        *common.Hash      `json:"stateRoot"`
	TxHash      *common.Hash      `json:"transactionsRoot"`
	ReceiptHash *common.Hash      `json:"receiptsRoot"`
	Bloom       *types.Bloom      `json:"logsBloom"`
	Difficulty  *hexutil.Big      `json:"difficulty"`
	Number      *hexutil.Big      `json:"number"`
	GasLimit    *hexutil.Big      `json:"gasLimit"`
	GasUsed     *hexutil.Big      `json:"gasUsed"`
	Time        *hexutil.Big      `json:"timestamp"`
	Extra       *hexutil.Bytes    `json:"extraData"`
	MixDigest   *common.Hash      `json:"mixHash"`
	Nonce       *types.BlockNonce `json:"nonce"`
}

var Geth = NewGethRPCClient()

// var ContractAddress = "0xe034afdcc2ba0441ff215ee9ba0da3e86450108d"
var ContractAddress = "0xa1a2a3a4a34598abcdeffed45902390854389043"
var ExtraData = "somethingextra"
var ShareDif = "0x186a0"

type geth struct {
	client *rpc.Client
}

func (g geth) GetPendingBlockHeader() *types.Header {
	header := jsonHeader{}
	err := g.client.Call(&header, "eth_getBlockByNumber", "pending", false)
	if err != nil {
		log.Fatal("Couldn't get latest block:", err)
		return nil
	}
	result := types.Header{}
	result.ParentHash = *header.ParentHash
	result.UncleHash = *header.UncleHash
	result.Root = *header.Root
	result.TxHash = *header.TxHash
	result.ReceiptHash = *header.ReceiptHash
	result.Difficulty = (*big.Int)(header.Difficulty)
	result.Number = (*big.Int)(header.Number)
	result.GasLimit = (*big.Int)(header.GasLimit)
	result.GasUsed = (*big.Int)(header.GasUsed)
	result.Time = (*big.Int)(header.Time)
	result.Coinbase = common.HexToAddress(ContractAddress)
	// result.Extra = []byte("0xd883010505846765746887676f312e372e348664617277696e")
	result.Extra = []byte(ExtraData)
	if header.Bloom == nil {
		result.Bloom = types.Bloom{}
	} else {
		result.Bloom = *header.Bloom
	}
	result.MixDigest = *header.MixDigest
	result.Nonce = types.BlockNonce{}
	return &result
}

func (g geth) GetBlockHeader(number int) *types.Header {
	header := types.Header{}
	err := g.client.Call(&header, "eth_getBlockByNumber", number, false)
	if err != nil {
		log.Fatal("Couldn't get latest block:", err)
		return nil
	}
	return &header
}

type gethWork [3]string

func (w gethWork) PoWHash() string { return w[0] }

func (g geth) GetWork() *spcommon.Work {
	w := gethWork{}
	var h *types.Header
	for {
		h = g.GetPendingBlockHeader()
		g.client.Call(&w, "eth_getWork")
		// waiting for pending block to be the same as
		// block we are going to pass to miner
		if w.PoWHash() != "" && w.PoWHash() == h.HashNoNonce().Hex() {
			break
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("Get inconsistent pending block header. Retry in 1s...\n")
	}
	return spcommon.NewWork(h, w[0], w[1])
}

func (g geth) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitHashrate", hashrate, id)
	return result
}

func (g geth) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitWork", nonce, hash, mixDigest)
	return result
}

func NewGethRPCClient() *geth {
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	return &geth{client}
}