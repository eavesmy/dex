package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var BSC_RPC_ADDR = "https://bsc-dataseed1.binance.org/"

type Bsc struct {
	*Client
	RpcAddr string
}

func (node *Bsc) Init(ctx context.Context, rpcAddrs ...string) (bsc Chain, err error) {

	node.Client = &Client{
		ctx:  ctx,
		node: node,
	}

	addr := BSC_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Bsc core init.")
	bsc = node
	return
}
