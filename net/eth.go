package net

import (
	"context"
	"crypto/ecdsa"
	"github.com/eavesmy/dex/enums"
	"github.com/eavesmy/dex/schema"
	"github.com/eavesmy/dex/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"math/big"
	"time"
)

type Eth struct {
	chainID *big.Int
	Ctx     context.Context
	Rpc     string
	client  *ethclient.Client
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

func (c *Eth) WalletCreate() (wallet *schema.Wallet, err error) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		err = enums.CreateWalletError
		return
	}

	wallet = new(schema.Wallet)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	wallet.PrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()

	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	wallet.PublicKey = hexutil.Encode(publicKeyBytes)[4:]

	wallet.Address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

func (c *Eth) WalletCreateByPrivateKey(privateKey string) (wallet *schema.Wallet, err error) {

	wallet = new(schema.Wallet)

	wallet.PrivateKey = privateKey

	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return
	}

	privateKeyBytes := crypto.FromECDSA(priKey)
	wallet.PrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := priKey.Public()

	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	wallet.PublicKey = hexutil.Encode(publicKeyBytes)[4:]

	wallet.Address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
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
	msg := e.Call(contractAddr, method, addr)

	callMsg := ethereum.CallMsg{
		To:   msg["To"].(*common.Address),
		Data: msg["Data"].([]byte),
	}

	data, err := e.client.CallContract(e.Ctx, callMsg, big.NewInt(int64(num)))

	if err != nil {
		return
	}

	balance = new(big.Int).SetBytes(data)
	return
}

func (e *Eth) NetworkID() (chainID *big.Int, err error) {
	chainID, err = e.client.ChainID(e.Ctx)
	e.chainID = chainID
	return
}

func (e *Eth) GetPastLogs() {
	// e.client.FilterLogs()
}

func (e *Eth) GetTransaction(txHash string) (transaction *schema.Transaction, err error) {

	hash := common.HexToHash(txHash)

	tx, isPending, err := e.client.TransactionByHash(e.Ctx, hash)

	transaction = &schema.Transaction{
		Hash:    tx.Hash().String(),
		ChainId: tx.ChainId(),
		To:      tx.To().String(),
		// Gas:       tx.Gas(),
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

func (e *Eth) Call(to, method string, params ...interface{}) map[string]interface{} {

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

	return map[string]interface{}{"To": &hex_to, "Data": data}
}

func (e *Eth) Nonce(from string) (nonce uint64, err error) {
	addr := common.HexToAddress(from)
	return e.client.PendingNonceAt(e.Ctx, addr)
}

func (e *Eth) GetGasPrice() (gasPrice *big.Int, err error) {
	return e.client.SuggestGasPrice(e.Ctx)
}

func (e *Eth) SendTransaction(transaction *schema.Transaction, privateKey string) (ret_transaction *schema.Transaction, err error) {

	// to
	var to common.Address
	if transaction.To == "" {
		err = enums.ParamToRequired
		return
	}
	to = common.HexToAddress(transaction.To)

	// value
	var value *big.Int
	if transaction.Value == nil {
		err = enums.ParamValueRequired
		return
	}
	value = transaction.Value

	// privateKey
	var priKey *ecdsa.PrivateKey
	if privateKey == "" {
		err = enums.PrivateKeyRequired
		return
	}

	var wallet *schema.Wallet
	if wallet, err = e.WalletCreateByPrivateKey(privateKey); err != nil {
		err = enums.InvalidPrivateKey
		return
	}
	priKey, _ = crypto.HexToECDSA(wallet.PrivateKey)
	transaction.From = wallet.Address

	// chainID
	var chainID = e.chainID
	if chainID == nil {
		if chainID, err = e.NetworkID(); err != nil {
			err = enums.RequestFailed
			return
		}
	}
	transaction.ChainId = chainID

	// nonce
	var nonce uint64
	if transaction.Nonce == 0 {
		if nonce, err = e.Nonce(transaction.From); err != nil {
			err = enums.RequestFailed
			return
		}
		transaction.Nonce = nonce
	}
	nonce = transaction.Nonce

	// gas
	var gas uint64
	if transaction.Gas == 0 {
		transaction.Gas = uint64(21000)
	}
	gas = transaction.Gas

	// gasprice
	var gasPrice *big.Int
	if transaction.GasPrice == nil {
		if gasPrice, err = e.GetGasPrice(); err != nil {
			err = enums.RequestFailed
			return
		}
		transaction.GasPrice = gasPrice
	}
	gasPrice = transaction.GasPrice

	// data
	var data []byte
	if transaction.Data == nil {
		transaction.Data = []byte{}
	}
	data = transaction.Data

	tx := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gas,
		To:       &to,
		Value:    value,
		Data:     data,
	}

	signedTx, err := types.SignTx(types.NewTx(tx), types.NewEIP155Signer(chainID), priKey)

	transaction.Hash = signedTx.Hash().Hex()
	transaction.IsPending = true
	transaction.CreatedAt = time.Now()

	return transaction, e.client.SendTransaction(e.Ctx, signedTx)
}
