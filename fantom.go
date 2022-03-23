package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var FANTOM_RPC_ADDR = "https://rpc.ftm.tools/"

type Fantom struct {
	*Client
	RpcAddr string
}

func (node *Fantom) Init(ctx context.Context) (fantom *Fantom, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, FANTOM_RPC_ADDR)
	fmt.Println("Fantom core init.")
	fantom = node
	return
}
