package contract

import (
	"../params"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"
)

type UpdaterClient struct {
	contract   Updater
	transactor *bind.TransactOpts
}

func (uc UpdaterClient) SetEpochData(merkleRoot *big.Int, fullSizeIn128Resolution uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return uc.contract.SetEpochData(uc.transactor, merkleRoot, fullSizeIn128Resolution, branchDepth, epoch)
}

func (uc UpdaterClient) VerifyExtraData(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (bool, error) {
	return uc.contract.VerifyExtraData(nil, extraData, minerId, difficulty)
}

func (uc UpdaterClient) VerifyExtraData_debug(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (*big.Int, error) {
	return uc.contract.VerifyExtraData_debug(nil, extraData, minerId, difficulty)
}

func (uc UpdaterClient) VerifyClaim_debug(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*big.Int, error) {
	return uc.contract.VerifyClaim_debug(nil, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

func (uc UpdaterClient) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return uc.contract.To62Encoding(nil, id, numChars)
}

func NewUpdaterClient() *UpdaterClient {
	client, err := getClient()
	if err != nil {
		log.Fatalf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return nil
	}
	pool, err := NewTestPool(common.HexToAddress(params.ContractAddress), client)
	if err != nil {
		log.Fatalf("Couldn't get SmartPool information from Ethereum Blockchain. Error: %s\n", err)
		return nil
	}
	account := GetAccount()
	if account == nil {
		log.Fatalf("Couldn't get any account from key store.\n")
		return nil
	}
	fmt.Printf("Key: %s\n", account.KeyFile())
	keyio, err := os.Open(account.KeyFile())
	if err != nil {
		log.Fatalf("Failed to open key file: %s\n", err)
		return nil
	}
	auth, err := bind.NewTransactor(keyio, account.PassPhrase())
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %s\n", err)
		return nil
	}
	return &UpdaterClient{pool, auth}
}
