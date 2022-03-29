package main

import (
	"context"
	"fmt"

	"github.com/eavesmy/dex"
	"github.com/eavesmy/dex/schema"
)

func main() {

	ctx := context.Background()

	var c dex.Chain
	var err error

	// Init eth client.

	// Set bsc rpc address.
	// Of course you can set ws/wss protocal here.
	// dex.BSC_RPC_ADDR = "https://bsc-dataseed2.binance.org/"
	c, _ = new(dex.Oasis).Init(ctx)

	num, _ := c.BlockNumber()
	fmt.Println("lastest block number: ", num)

	block, _ := c.GetBlockByNumber(num)
	fmt.Printf("block: %+v \n", block.Hash)
	fmt.Println("block size: ", block.Size, block.BaseFeeGas)

	if _, err = c.GetPastLogs(schema.LogQuery{FromBlock: 773609, ToBlock: 773609}); err != nil {
		fmt.Println("logs error:", err)
	}

	block, err = c.GetBlockByHash(block.Hash)
	fmt.Printf("block: %+v error: %s \n ", block, err.Error())

	// logs, _ := bsc.GetPastLogs(schema.LogQuery{BlockHash: block.Hash})
	// fmt.Println(len(logs), logs[0])

	// Create a new account from bsc.
	//wallet, _ := bsc.WalletCreate()

	//// Set privateKey and send transaction to self.
	//tx, err := bsc.SetPrivateKey(wallet.PrivateKey).SendTransaction(&schema.Transaction{To: wallet.Address, Value: new(big.Int).SetUint64(23937000000000)})
	//fmt.Println(tx, err)
}
