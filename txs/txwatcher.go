package txs

import (
	"../client"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// transaction pool keeps track of pending transactions
// and acknowledge corresponding channel when a transaction is
// confirmed
type TxWatcher struct {
	tx      *types.Transaction
	verChan chan bool
}

func (tw *TxWatcher) isVerified() bool {
	return client.DefaultGethClient.IsVerified(tw.tx.Hash())
}

// loop to check transactions verification
// if a transaction is verified, send it to verChan
func (tw *TxWatcher) loop() {
	for {
		if tw.isVerified() {
			tw.verChan <- true
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func (tw *TxWatcher) Wait() {
	go tw.loop()
	<-tw.verChan
}

func NewTxWatcher(tx *types.Transaction) *TxWatcher {
	return &TxWatcher{tx, make(chan bool)}
}
