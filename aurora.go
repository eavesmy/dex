package dex

import (
	"context"
	"fmt"

	"github.com/eavesmy/dex/net"
)

var AURORA_RPC_ADDR = "https://mainnet.aurora.dev"

type Aurora struct {
	*Client
	RpcAddr string
}

func (node *Aurora) Init(ctx context.Context, rpcAddrs ...string) (aurora Chain, err error) {

	node.Client = &Client{
		ctx: ctx,
	}

	addr := AURORA_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Aurora core init.")
	aurora = node
	return aurora, err
}
