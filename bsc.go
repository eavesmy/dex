package dex

import (
	"context"
	"dex/net"
	"fmt"
)

type Bsc struct {
	*Client
	core *net.Eth
}

func (bsc *Bsc) Init(ctx context.Context) (b *Bsc, err error) {
	bsc.Client = &Client{
		ctx: ctx,
	}
	bsc.Client.core, err = new(net.Eth).Init(ctx, "https://bsc-dataseed3.binance.org/")
	fmt.Println("Bsc core init.")
	return bsc, err
}
