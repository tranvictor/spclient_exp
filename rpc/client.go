package rpc

import (
	spcommon "../common"
	"../params"
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

var DefaultGethClient *GethClient

type GethClient struct {
	client *rpc.Client
}

func (g GethClient) GetPendingBlockHeader() *types.Header {
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
	result.Coinbase = common.HexToAddress(params.ContractAddress)
	// result.Extra = []byte("0xd883010505846765746887676f312e372e348664617277696e")
	result.Extra = []byte(params.ExtraData)
	if header.Bloom == nil {
		result.Bloom = types.Bloom{}
	} else {
		result.Bloom = *header.Bloom
	}
	result.MixDigest = *header.MixDigest
	result.Nonce = types.BlockNonce{}
	return &result
}

func (g GethClient) GetBlockHeader(number int) *types.Header {
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

func (g GethClient) GetWork() *spcommon.Work {
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

func (g GethClient) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitHashrate", hashrate, id)
	return result
}

func (g GethClient) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitWork", nonce, hash, mixDigest)
	return result
}

func NewGethRPCClient() (*GethClient, error) {
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		return nil, err
	}
	return &GethClient{client}, nil
}
