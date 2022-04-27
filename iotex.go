package dex

import (
	"context"
	"fmt"

	"github.com/eavesmy/dex/net"
)

var IOTEX_RPC_ADDR = "https://babel-api.mainnet.iotex.io"

type Iotex struct {
	*Client
	RpcAddr string
}

func (node *Iotex) Init(ctx context.Context, rpcAddrs ...string) (eth Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := ETH_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Iotex core init.")
	eth = node
	return eth, err
}
