package dex

import (
	"context"
	"dex/net"
	"fmt"
)

type Near struct {
	*Client
}

func (n *Near) Init(ctx context.Context) (near *Near, err error) {
	n.Client = &Client{ctx: ctx}
	n.Client.core, err = new(net.Near).Init(ctx, "https://rpc.mainnet.near.org")
	fmt.Println("Near core init.")
	near = n
	return
}
