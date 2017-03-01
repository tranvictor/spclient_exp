package contract

import (
	"../params"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

type MinerAccount struct {
	keyFile    string
	passphrase string
}

func (ma MinerAccount) KeyFile() string    { return ma.keyFile }
func (ma MinerAccount) PassPhrase() string { return ma.passphrase }

// Get the first account in key store
// Return nil if there's no account
func GetAccount() *MinerAccount {
	keys := keystore.NewKeyStore(
		params.KeystorePath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	if len(keys.Accounts()) == 0 {
		return nil
	}
	acc := keys.Accounts()[0]
	keyFile := acc.URL.Path
	passphrase, err := promptUserPassPhrase(acc.Address.Hex())
	if err != nil {
		return nil
	}
	return &MinerAccount{
		keyFile,
		passphrase,
	}
}

func promptUserPassPhrase(acc string) (string, error) {
	fmt.Printf("Using account address: %s\n", acc)
	fmt.Printf("Please enter passphrase: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")
	if err != nil {
		return "", err
	} else {
		return string(bytePassword), nil
	}
}
