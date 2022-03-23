package dex

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex/net"
)

type Oasis struct {
	*Client
}

func (o *Oasis) Init(ctx context.Context) (oasis *Oasis, err error) {
	o.Client = &Client{
		ctx: ctx,
	}
	o.Client.core, err = new(net.Eth).Init(ctx, "ws://54.65.82.217:9944")
	// o.Client.core, err = new(net.Eth).Init(ctx, "https://emerald.oasis.dev/")

	oasis = o

	fmt.Println("Oasis core init.")
	return
}
