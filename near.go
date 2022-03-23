package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var NEAR_RPC_ADDR = "https://rpc.mainnet.near.org"

type Near struct {
	*Client
	RpcAddr string
}

func (node *Near) Init(ctx context.Context) (near *Near, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Near).Init(ctx, NEAR_RPC_ADDR)
	fmt.Println("Near core init.")
	near = node
	return
}
