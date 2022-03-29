package dex

import (
	"context"
	"errors"
	"fmt"

	"github.com/eavesmy/dex/net"
	"github.com/eavesmy/dex/schema"
)

var OASIS_RPC_ADDR = "https://emerald.oasis.dev/"

type Oasis struct {
	*Client
}

func (node *Oasis) Init(ctx context.Context, rpcAddrs ...string) (oasis Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := OASIS_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Oasis core init.")
	oasis = node
	return
}

func (node *Oasis) GetBlockByHash(string) (block *schema.Block, err error) {
	err = errors.New("Not supported. Please use `GetBlockByNumber` method instead.")
	return
}
