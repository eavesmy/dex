package schema

import "math/big"

type Transaction struct {
	Hash        string
	BlockNumber uint64
	To          string
	ChainId     *big.Int
	Gas         uint64
	Cost        *big.Int
	Data        []byte
	GasPrice    *big.Int
	Nonce       uint64

	IsPending bool
}
