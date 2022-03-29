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

func (node *Klaytn) Init(ctx context.Context, rpcAddrs ...string) (klaytn Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := KLAYTN_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Klaytn core init.")
	klaytn = node
	return
}
