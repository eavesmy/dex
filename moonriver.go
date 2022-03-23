package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var MOONRIVER_RPC_ADDR = "https://rpc.moonriver.moonbeam.network"

type Moonriver struct {
	*Client
	RpcAddr string
}

func (node *Moonriver) Init(ctx context.Context) (moonriver *Moonriver, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, MOONRIVER_RPC_ADDR)
	fmt.Println("Moonriver core init.")
	moonriver = node
	return
}
