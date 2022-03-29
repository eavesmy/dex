package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var MATIC_RPC_ADDR = "https://polygon-rpc.com"

type Matic struct {
	*Client
	RpcAddr string
}

func (node *Matic) Init(ctx context.Context, rpcAddrs ...string) (matic Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := MATIC_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Matic core init.")
	return node, err
}
