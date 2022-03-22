package net

import (
	"context"
	"dex/schema"
	"dex/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"math/big"
)

type Eth struct {
	Ctx    context.Context
	Rpc    string
	client *ethclient.Client
}

// Init eth chain net.
func (e *Eth) Init(ctx context.Context, rpc string) (c *Eth, err error) {
	e.Ctx = ctx
	e.Rpc = rpc

	client, err := ethclient.Dial(e.Rpc)
	if err != nil {
		return
	}
	e.client = client
	return e, nil
}

func (e *Eth) BlockNumber() (uint64, error) {
	return e.client.BlockNumber(e.Ctx)
}

func (e *Eth) GetBalance(addr string) (balance *big.Int, err error) {
	num, err := e.BlockNumber()
	if err != nil {
		return
	}
	balance, err = e.client.BalanceAt(e.Ctx, common.HexToAddress(addr), big.NewInt(int64(num)))
	return
}

func (e *Eth) GetBalanceOf(addr, contractAddr string) (balance *big.Int, err error) {
	num, err := e.BlockNumber()
	if err != nil {
		return
	}

	method := "balanceOf(address)"
	msg := e.callContract(contractAddr, method, addr)

	data, err := e.client.CallContract(e.Ctx, msg, big.NewInt(int64(num)))

	if err != nil {
		return
	}

	balance = new(big.Int).SetBytes(data)
	return
}

func (e *Eth) NetworkID() (id string, err error) {
	big_id, err := e.client.NetworkID(e.Ctx)
	id = big_id.String()
	return
}

func (e *Eth) GetPastLogs() {
	// e.client.FilterLogs()
}

func (e *Eth) GetTransaction(txHash string) (transaction *schema.Transaction, err error) {

	hash := common.HexToHash(txHash)

	tx, isPending, err := e.client.TransactionByHash(e.Ctx, hash)

	transaction = &schema.Transaction{
		Hash:      tx.Hash().String(),
		ChainId:   tx.ChainId(),
		To:        tx.To().String(),
		Gas:       tx.Gas(),
		GasPrice:  tx.GasPrice(),
		Cost:      tx.Cost(),
		Data:      tx.Data(),
		Nonce:     tx.Nonce(),
		IsPending: isPending,
	}

	return
}

func (e *Eth) GetTransactionCount(blockHash string) (count int, err error) {
	hash := common.HexToHash(blockHash)
	num, err := e.client.TransactionCount(e.Ctx, hash)
	count = int(num)
	return
}

func (e *Eth) callContract(to, method string, params ...interface{}) ethereum.CallMsg {

	hex_to := common.HexToAddress(to)
	data := []byte{}

	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(method))
	methodId := hash.Sum(nil)[:4]

	data = append(data, methodId...)

	for _, param := range params {
		if typedP, ok := param.(string); ok {
			p := common.HexToAddress(typedP)
			data = append(data, utils.LeftPadding(p.Bytes(), 32)...)
		}
		if typedP, ok := param.(*big.Int); ok {
			data = append(data, utils.LeftPadding(typedP.Bytes(), 32)...)
		}
	}

	return ethereum.CallMsg{
		To:   &hex_to,
		Data: data,
	}
}
