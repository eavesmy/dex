package schema

import (
	"math/big"
)

type Transaction struct {
	Hash        string
	BlockNumber uint64
	From        string
	To          string
	Value       *big.Int
	ChainId     *big.Int

	// Gas amount
	Gas uint64

	// Total gas price
	Cost *big.Int

	// Per gas price
	GasPrice *big.Int

	Data  []byte
	Nonce uint64

	IsPending bool
}
