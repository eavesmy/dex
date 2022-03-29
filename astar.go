package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var ASTAR_RPC_ADDR = "https://astar.api.onfinality.io/public"

type Astar struct {
	*Client
	RpcAddr string
}

func (node *Astar) Init(ctx context.Context, rpcAddrs ...string) (astar Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := ASTAR_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Astar core init.")
	astar = node
	return
}
