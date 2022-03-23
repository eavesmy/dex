package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var CELO_RPC_ADDR = "https://forno.celo.org"

type Celo struct {
	*Client
	RpcAddr string
}

func (node *Celo) Init(ctx context.Context) (celo *Celo, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, CELO_RPC_ADDR)
	fmt.Println("Celo core init.")
	celo = node
	return
}
