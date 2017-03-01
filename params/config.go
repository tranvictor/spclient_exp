package params

import (
	"math/big"
	"time"
)

var (
	IPCPath         string
	KeystorePath    string
	NoSharePerClaim uint32
	ShareDifficulty *big.Int
	SubmitInterval  time.Duration
	ContractAddress string
	MinerAddress    string
	ExtraData       string
)
