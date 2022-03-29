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

func (node *Astar) Init(ctx context.Context) (astar Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, ASTAR_RPC_ADDR)
	fmt.Println("Astar core init.")
	astar = node
	return
}
