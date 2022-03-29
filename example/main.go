package main

import (
	"context"
	"fmt"
	"github.com/eavesmy/dex"
	"github.com/eavesmy/dex/schema"
	"math/big"
)

func main() {

	ctx := context.Background()

	// Init eth client.
	eth, _ := new(dex.Eth).Init(ctx)
	// Get last block number.
	num, err := eth.BlockNumber()
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
