package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var AURORA_RPC_ADDR = ""

type Aurora struct {
	*Client
	RpcAddr string
}

func (node *Aurora) Init(ctx context.Context) (aurora *Aurora, err error) {

	if AURORA_RPC_ADDR == "" {
		panic("Variable 'AURORA_RPC_ADDR' required")
	}

	node.Client = &Client{
		ctx: ctx,
	}

	node.Client.core, err = new(net.Eth).Init(ctx, AURORA_RPC_ADDR)
	fmt.Println("Eth core init.")
	aurora = node
	return aurora, err
}
