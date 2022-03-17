package dex

type ChainNet interface {
	LastBlock() int64
}
