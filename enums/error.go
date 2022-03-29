package enums

import "errors"

var (
	RequestFailed      = errors.New("Request failed")
	CreateWalletError  = errors.New("Create wallet error")
	PrivateKeyRequired = errors.New("Set private key before send transaction")
	InvalidPrivateKey  = errors.New("Invalid private key")
	ParamToRequired    = errors.New("Param to required")
	ParamValueRequired = errors.New("Param value required")
	NotSupported       = errors.New("Not supported")
)
