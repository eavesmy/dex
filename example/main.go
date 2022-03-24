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

	eth, _ := new(dex.Eth).Init(ctx)
	num, err := eth.BlockNumber()
	fmt.Println(num, err)

	dex.BSC_RPC_ADDR = "https://bsc-dataseed2.binance.org/"
	bsc, _ := new(dex.Bsc).Init(ctx)

	wallet, _ := bsc.WalletCreate()
	tx, err := bsc.SetPrivateKey(wallet.PrivateKey).SendTransaction(&schema.Transaction{To: wallet.Address, Value: new(big.Int).SetUint64(23937000000000)})
	fmt.Println(tx, err)

	klaytn, _ := new(dex.Klaytn).Init(ctx)
	num, err = klaytn.BlockNumber()
	fmt.Println(num, err)
}
