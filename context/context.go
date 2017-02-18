package context

// Context holds all of the SmartPool client data that can be
// accessed from multiple concurrent rpc requests
// Context should contain information about:
//	1. IPC path
//	2. Keystore directory
//	3. Pending transaction pool
//	4. Geth client
//	5. RPC server
//	6. Contract client
//	7. Ethash instance
//	8. Claim repository
//	9. Claim channel
//	10. Event channel
//	-- setting --
//	11. Number of share per claim
//	12. Difficulty for a share if it's fixed (need more discussion)

import (
	"../contract"
	"../ethash"
	"../rpc"
	"../share"
	"math/big"
)

type Context struct {
	IPCPath        string
	KeystorePath   string
	GethClient     *rpc.GethClient
	Server         *rpc.Server
	Contract       *contract.ContractClient
	Ethash         *ethash.Ethash
	ClaimRepo      *share.ClaimRepo
	ReadyClaimChan chan *share.Claim
	// Setting
	NoSharePerClaim uint32
	ShareDifficulty *big.Int
	// TODO: event channel
	// PendingTxs     *PendingTxPool
}
