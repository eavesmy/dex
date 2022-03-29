package dex

import (
	"context"
	"github.com/eavesmy/dex/schema"
	"math/big"
)

type TransferOption struct {
}

// Chain is implement by Client
type Chain interface {
	Init(context.Context) (Chain, error)
	WalletCreate() (*schema.Wallet, error)
	BlockNumber() (uint64, error)
	GetBalance(string) (*big.Int, error)
	// GetBalanceOf Query token balance.
	GetBalanceOf(string, string) (*big.Int, error)
	NetworkID() (*big.Int, error)
	Nonce(string) (uint64, error)
	GetTransaction(string) (*schema.Transaction, error)
	GetPastLogs(schema.LogQuery) ([]*schema.Log, error)
	// Call is call contract method.
	// ex.: `Call(contractAddr,"Transfer(address,address)",...params)`
	Call(string, string, ...interface{}) map[string]interface{}
	SetPrivateKey(string) Chain
	SendTransaction(*schema.Transaction) (*schema.Transaction, error)
}

// SetPrivateKey is implement by net.Eth or other chain net.
type ChainNet interface {
	WalletCreate() (*schema.Wallet, error)
	BlockNumber() (uint64, error)
	GetBalance(string) (*big.Int, error)
	GetBalanceOf(string, string) (*big.Int, error)
	NetworkID() (*big.Int, error)
	Nonce(string) (uint64, error)
	GetTransaction(string) (*schema.Transaction, error)
	GetPastLogs(schema.LogQuery) ([]*schema.Log, error)
	Call(string, string, ...interface{}) map[string]interface{}
	SendTransaction(*schema.Transaction, string) (*schema.Transaction, error)
}

type Client struct {
	core       ChainNet
	ctx        context.Context
	privateKey string
	node       Chain
}

func (c *Client) BlockNumber() (uint64, error) {
	return c.core.BlockNumber()
}

func (c *Client) WalletCreate() (wallet *schema.Wallet, err error) {
	return c.core.WalletCreate()
}

func (c *Client) GetBalance(addr string) (*big.Int, error) {
	return c.core.GetBalance(addr)
}

func (c *Client) GetBalanceOf(addr, contractAddr string) (*big.Int, error) {
	return c.core.GetBalanceOf(addr, contractAddr)
}

func (c *Client) NetworkID() (*big.Int, error) {
	return c.core.NetworkID()
}

func (c *Client) GetTransaction(txHash string) (*schema.Transaction, error) {
	return c.core.GetTransaction(txHash)
}

// SetPrivateKey Before calling an operation that requires verification of account privileges.
func (c *Client) SetPrivateKey(privateKey string) Chain {
	c.privateKey = privateKey
	return c.node
}

func (c *Client) SendTransaction(transaction *schema.Transaction) (*schema.Transaction, error) {
	return c.core.SendTransaction(transaction, c.privateKey)
}

// Call is call contract method.
func (c *Client) Call(to string, method string, params ...interface{}) map[string]interface{} {
	return c.core.Call(to, method, params)
}

func (c *Client) Nonce(addr string) (uint64, error) {
	return c.core.Nonce(addr)
}
func (c *Client) GetPastLogs(query schema.LogQuery) (logs []*schema.Log, err error) {
	return c.core.GetPastLogs(query)
}
