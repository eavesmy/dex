package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var OASIS_RPC_ADDR = "https://emerald.oasis.dev/"

type Oasis struct {
	*Client
}

func (node *Oasis) Init(ctx context.Context) (oasis *Oasis, err error) {
	node.Client = &Client{
		ctx: ctx,
	}
	node.Client.core, err = new(net.Eth).Init(ctx, OASIS_RPC_ADDR)
	fmt.Println("Oasis core init.")
	oasis = node
	return
}
