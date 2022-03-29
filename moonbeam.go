package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var MOONBEAM_RPC_ADDR = "https://rpc.api.moonbeam.network"

type Moonbeam struct {
	*Client
	RpcAddr string
}

func (node *Moonbeam) Init(ctx context.Context, rpcAddrs ...string) (moonbeam Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := MOONBEAM_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Moonbeam core init.")
	moonbeam = node
	return
}
