package schema

type LogQuery struct {
	BlockHash string
	FromBlock uint64
	ToBlock   uint64
	Addresses []string
	Topics    [][]string
}

type Log struct {
	Address     string
	Topics      []string
	Data        []byte
	Index       uint
	BlockNumber uint64
	BlockHash   string
	TxHash      string
	TxIndex     uint
	Removed     bool // 应对回滚等异常情况，这个字段保留
}
