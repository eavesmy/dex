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

	/*
		c, _ = new(dex.Oasis).Init(ctx)
		// ret, err := c.Call("0xE9b38eD157429483EbF87Cf6C002cecA5fd66783", "name()")

		c, _ = new(dex.Cronos).Init(ctx)
		blockNumber, err := c.BlockNumber()
		fmt.Printf("cronos height: %d\n", blockNumber)

		dex.HARMONY_RPC_ADDR = "wss://ws.s0.t.hmny.io"
		c, _ = new(dex.Harmony).Init(ctx)

		blockNumber, err = c.BlockNumber()
		fmt.Printf("harmony height: %d\n", blockNumber)
	*/

	c, _ = new(dex.Iotex).Init(ctx)
	blockNumber, err := c.BlockNumber()
	fmt.Println(blockNumber, err)

	block, err := c.GetBlockByNumber(blockNumber)
	fmt.Println("block", block.Transactions[0].From)

	c, _ = new(dex.Aurora).Init(ctx)

	num, _ := c.BlockNumber()
	fmt.Println("lastest block number: ", num)

	block, err = c.GetBlockByNumber(num)
	fmt.Println("block err: ", block, err)

	block, err = c.GetBlockByNumber(63060094)

	if err != nil {
		fmt.Printf("Get block by number error: %+v", err)
	} else {
		fmt.Printf("block: %+v \n", block.Hash)
		fmt.Println("block size: ", block.Size, block.BaseFeeGas)

		block, err = c.GetBlockByHash(block.Hash)
		fmt.Printf("block: %+v error: %s \n ", block, err.Error())
	}

	if _, err = c.GetPastLogs(schema.LogQuery{FromBlock: 773609, ToBlock: 773609}); err != nil {
		fmt.Println("logs error:", err)
	}

	// logs, _ := bsc.GetPastLogs(schema.LogQuery{BlockHash: block.Hash})
	// fmt.Println(len(logs), logs[0])

	// Create a new account from bsc.
	//wallet, _ := bsc.WalletCreate()

	//// Set privateKey and send transaction to self.
	//tx, err := bsc.SetPrivateKey(wallet.PrivateKey).SendTransaction(&schema.Transaction{To: wallet.Address, Value: new(big.Int).SetUint64(23937000000000)})
	//fmt.Println(tx, err)
}
