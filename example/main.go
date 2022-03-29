package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/eavesmy/dex"
	"github.com/eavesmy/dex/schema"
)

func main() {

	ctx := context.Background()

	var c dex.Chain

	// Init eth client.
	c, _ = new(dex.Eth).Init(ctx)
	// Get last block number.
	num, err := c.BlockNumber()
	fmt.Println(num, err)

	// Set bsc rpc address.
	// Of course you can set ws/wss protocal here.
	dex.BSC_RPC_ADDR = "https://bsc-dataseed2.binance.org/"
	bsc, _ := new(dex.Bsc).Init(ctx)

	// Create a new account from bsc.
	wallet, _ := bsc.WalletCreate()

	// Set privateKey and send transaction to self.
	tx, err := bsc.SetPrivateKey(wallet.PrivateKey).SendTransaction(&schema.Transaction{To: wallet.Address, Value: new(big.Int).SetUint64(23937000000000)})
	fmt.Println(tx, err)
}
