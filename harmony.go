package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var HARMONY_RPC_ADDR = "https://api.s0.t.hmny.io/"

type Harmony struct {
	*Client
	RpcAddr string
}

func (node *Harmony) Init(ctx context.Context, rpcAddrs ...string) (eth Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := HARMONY_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Harmony core init.")
	return node, err
}
