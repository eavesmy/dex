package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var BSC_RPC_ADDR = "https://emerald.oasis.dev/"

type Bsc struct {
	*Client
	RpcAddr string
}

func (node *Bsc) Init(ctx context.Context) (eth *Bsc, err error) {

	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, BSC_RPC_ADDR)
	fmt.Println("Bsc core init.")
	eth = node
	return eth, err
}
