package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var ETH_RPC_ADDR = "https://mainnet.infura.io/v3/45a55fdb8471465da9ef58e655537cb0"

type Eth struct {
	*Client
	RpcAddr string
}

func (node *Eth) Init(ctx context.Context, rpcAddrs ...string) (eth Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := ETH_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Eth core init.")
	eth = node
	return eth, err
}
