package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

var CRONOS_RPC_ADDR = "https://evm-cronos.crypto.org"

type Cronos struct {
	*Client
	RpcAddr string
}

func (node *Cronos) Init(ctx context.Context, rpcAddrs ...string) (Cronos Chain, err error) {
	node.Client = &Client{
		ctx: ctx,
	}

	addr := CRONOS_RPC_ADDR
	if len(rpcAddrs) > 0 {
		addr = rpcAddrs[0]
	}

	node.Client.core, err = new(net.Eth).Init(ctx, addr)
	fmt.Println("Cronos core init.")
	Cronos = node
	return Cronos, err
}
