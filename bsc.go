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

func (node *Bsc) Init(ctx context.Context) (bsc *Bsc, err error) {

	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, BSC_RPC_ADDR)
	fmt.Println("Bsc core init.")
	bsc = node
	return
}
