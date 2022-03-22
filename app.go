package dex

import (
	"context"
	"github.com/eavesmy/dex/schema"
	"math/big"
)

type ChainNet interface {
	BlockNumber() (uint64, error)
	GetBalance(string) (*big.Int, error)
	GetBalanceOf(string, string) (*big.Int, error)
	NetworkID() (string, error)
	GetTransaction(string) (*schema.Transaction, error)
}

type Client struct {
	core ChainNet
	ctx  context.Context
}

func (c *Client) LastBlock() (uint64, error) {
	return c.core.BlockNumber()
}

func (c *Client) GetBalance(addr string) (*big.Int, error) {
	return c.core.GetBalance(addr)
}

func (c *Client) GetBalanceOf(addr, contractAddr string) (*big.Int, error) {
	return c.core.GetBalanceOf(addr, contractAddr)
}

func (c *Client) NetWorkID() (string, error) {
	return c.core.NetworkID()
}

func (c *Client) GetTransaction(txHash string) (*schema.Transaction, error) {
	return c.core.GetTransaction(txHash)
}
