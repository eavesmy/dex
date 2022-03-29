package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var MOONRIVER_RPC_ADDR = "https://rpc.moonriver.moonbeam.network"

type Moonriver struct {
	*Client
	RpcAddr string
}

func (node *Moonriver) Init(ctx context.Context, rpcAddrs ...string) (moonriver Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := MOONRIVER_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Moonriver core init.")
	moonriver = node
	return
}
