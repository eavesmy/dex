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

	txs, err := eth.GetTransaction("0xf4a99f1e6560b3f0a075c598e9d0fdbf8d1eee3d4a51113b1b32b697191a06c4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txs)

	// balance, err := bsc.GetBalance("0x9Be835d10A7073CA52C5002C88adE89397B2330e")
	// fmt.Println("balance: ", balance, err)

	// balance, err := oasis.GetBalanceOf("0xC6C208F5B0Aa905CE2850952CeE8d584aD3511AA", "0xe9b38ed157429483ebf87cf6c002ceca5fd66783")
	// fmt.Println("usdt balance: ", balance, err)

	// id, err := bsc.NetWorkID()
	// fmt.Println("network id: ", id, err)

	// tx, err := oasis.GetTransaction("0xe054b3fcdca77536c555e40ea2ff541a2358a8adb6c9889572fb584b0e33f306")
	// fmt.Println(tx, err)

}
