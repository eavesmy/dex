package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var KLAYTN_RPC_ADDR = "https://public-node-api.klaytnapi.com/v1/cypress"

type Klaytn struct {
	*Client
	RpcAddr string
}

func (node *Klaytn) Init(ctx context.Context) (klaytn Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, KLAYTN_RPC_ADDR)
	fmt.Println("Klaytn core init.")
	klaytn = node
	return
}
