package contract

import (
	"../params"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var DefaultContractClient *ContractClient

type ContractClient struct {
	// the contract implementation that holds all underlying
	// communication with Ethereum Contract
	contract   Contract
	transactor *bind.TransactOpts
}

func (cc ContractClient) SubmitClaim(
	numShares *big.Int,
	difficulty *big.Int,
	min *big.Int,
	max *big.Int,
	augMerkle *big.Int) (*types.Transaction, error) {
	return cc.contract.SubmitClaim(cc.transactor,
		numShares, difficulty, min, max, augMerkle)
}

func (cc ContractClient) VerifyClaim(
	rlpHeader []byte,
	nonce *big.Int,
	shareIndex *big.Int,
	dataSetLookup []*big.Int,
	witnessForLookup []*big.Int,
	augCountersBranch []*big.Int,
	augHashesBranch []*big.Int) (*types.Transaction, error) {
	return cc.contract.VerifyClaim(cc.transactor,
		rlpHeader, nonce, shareIndex, dataSetLookup,
		witnessForLookup, augCountersBranch, augHashesBranch)
}

func (cc ContractClient) Version() string {
	v, err := cc.contract.Version(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve pool version: %s\n", err)
		return ""
	} else {
		fmt.Printf("SmartPool version: %s\n", v)
		return v
	}
}

func getClient() (*ethclient.Client, error) {
	return ethclient.Dial(params.IPCPath)
}

func NewContractClient() (*ContractClient, error) {
	client, err := getClient()
	if err != nil {
		fmt.Printf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return nil, err
	}
	pool, err := NewTestPool(common.HexToAddress(params.ContractAddress), client)
	if err != nil {
		fmt.Printf("Couldn't get SmartPool information from Ethereum Blockchain. Error: %s\n", err)
		return nil, err
	}
	account := GetAccount()
	if account == nil {
		fmt.Printf("Couldn't get any account from key store.\n")
		return nil, err
	}
	fmt.Printf("Key: %s\n", account.KeyFile())
	keyio, err := os.Open(account.KeyFile())
	if err != nil {
		fmt.Printf("Failed to open key file: %s\n", err)
		return nil, err
	}
	fmt.Printf("Unlocking account...")
	auth, err := bind.NewTransactor(keyio, account.PassPhrase())
	if err != nil {
		fmt.Printf("Failed to create authorized transactor: %s\n", err)
		return nil, err
	}
	fmt.Printf("Done.\n")
	return &ContractClient{pool, auth}, nil
}
