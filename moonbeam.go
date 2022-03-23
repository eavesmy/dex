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

func (node *Moonbeam) Init(ctx context.Context) (moonbeam *Moonbeam, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, MOONBEAM_RPC_ADDR)
	fmt.Println("Moonbeam core init.")
	moonbeam = node
	return
}
