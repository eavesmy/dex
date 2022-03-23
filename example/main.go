package main

import (
	"context"
	"fmt"

	"github.com/eavesmy/dex"
)

func main() {

	ctx := context.Background()

	eth, _ := new(dex.Eth).Init(ctx)

	num, err := eth.LastBlock()
	fmt.Println(num, err)

	dex.BSC_RPC_ADDR = "https://bsc-dataseed2.binance.org/"
	bsc, _ := new(dex.Bsc).Init(ctx)

	num, err = bsc.LastBlock()

	fmt.Println(num, err)
}
