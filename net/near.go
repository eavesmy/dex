package net

import (
	"context"
	"dex/schema"
	"encoding/json"
	"github.com/aurora-is-near/near-api-go"
	"math/big"
)

type Near struct {
	Ctx    context.Context
	Rpc    string
	client *near.Connection
}

func (n *Near) Init(ctx context.Context, rpc string) (inst *Near, err error) {

	n.client = near.NewConnection(rpc)

	inst = n
	return
}

func (n *Near) BlockNumber() (height uint64, err error) {

	block, err := n.client.Block()
	if err != nil {
		return
	}

	header := block["header"].(map[string]interface{})
	h, err := header["height"].(json.Number).Int64()
	if err != nil {
		return
	}
	height = uint64(h)
	return
}

// GetBalance not implemented!
func (n *Near) GetBalance(addr string) (balance *big.Int, err error) {
	return
}

func (n *Near) GetBalanceOf(addr, contract string) (balance *big.Int, err error) {
	return
}

func (n *Near) NetworkID() (id string, err error) {
	return
}
func (n *Near) GetTransaction(txHash string) (transaction *schema.Transaction, err error) {
	return
}
