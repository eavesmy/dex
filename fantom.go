package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var FANTOM_RPC_ADDR = "https://rpc.ftm.tools/"

type Fantom struct {
	*Client
	RpcAddr string
}

func (node *Fantom) Init(ctx context.Context, rpcAddrs ...string) (fantom Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := FANTOM_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Fantom core init.")
	fantom = node
	return
}
