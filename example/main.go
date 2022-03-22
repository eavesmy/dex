package main

import (
	"context"
	"dex"
	"fmt"
)

func main() {

	ctx := context.Background()

	bsc, _ := new(dex.Bsc).Init(ctx)

	// num, e := bsc.LastBlock()
	// fmt.Println(num, err, e)

	// balance, err := bsc.GetBalance("0x9Be835d10A7073CA52C5002C88adE89397B2330e")
	// fmt.Println("balance: ", balance, err)

	// balance, err := bsc.GetBalanceOf("0x9Be835d10A7073CA52C5002C88adE89397B2330e", "0x55d398326f99059fF775485246999027B3197955")
	// fmt.Println("usdt balance: ", balance, err)

	// id, err := bsc.NetWorkID()
	// fmt.Println("network id: ", id, err)

	tx, err := bsc.GetTransaction("0xe054b3fcdca77536c555e40ea2ff541a2358a8adb6c9889572fb584b0e33f306")
	fmt.Println(tx, err)

}
