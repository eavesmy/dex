package main

import (
	"context"
	"fmt"

	"github.com/eavesmy/dex"
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

	num, _ = bsc.BlockNumber()
	fmt.Println("lastest block number: ", num)

	block, _ := bsc.GetBlockByNumber(num)
	fmt.Printf("block: %+v \n", block.Hash)

	block, _ = bsc.GetBlockByHash(block.Hash)
	fmt.Printf("block: %+v \n", block.Hash)
	fmt.Println("block size: ", block.Size, block.BaseFeeGas)

	// logs, _ := bsc.GetPastLogs(schema.LogQuery{BlockHash: block.Hash})
	// fmt.Println(len(logs), logs[0])

	// Create a new account from bsc.
	//wallet, _ := bsc.WalletCreate()

	//// Set privateKey and send transaction to self.
	//tx, err := bsc.SetPrivateKey(wallet.PrivateKey).SendTransaction(&schema.Transaction{To: wallet.Address, Value: new(big.Int).SetUint64(23937000000000)})
	//fmt.Println(tx, err)
}
