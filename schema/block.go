package schema

import "time"

type Block struct {
	Uncle  string
	Parent string

	Hash         string
	Number       uint64
	Nonce        uint64
	Transactions []*Transaction

	ReceiveAt time.Time
	Time      uint64

	Difficulty uint64
	Size       string
	GasUsed    uint64
	GasLimit   uint64
	BaseFeeGas uint64
	ExtraData  []byte
}
