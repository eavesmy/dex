package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var AURORA_RPC_ADDR = ""

type Aurora struct {
	*Client
	RpcAddr string
}

func (node *Aurora) Init(ctx context.Context, rpcAddrs ...string) (aurora Chain, err error) {

	if AURORA_RPC_ADDR == "" {
		panic("Variable 'AURORA_RPC_ADDR' required")
	}

	node.Client = &Client{
		ctx: ctx,
	}

	addr := AURORA_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Eth core init.")
	aurora = node
	return aurora, err
}
