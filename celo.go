package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var CELO_RPC_ADDR = "https://forno.celo.org"

type Celo struct {
	*Client
	RpcAddr string
}

func (node *Celo) Init(ctx context.Context, rpcAddrs ...string) (celo Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := CELO_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Celo core init.")
	celo = node
	return
}
