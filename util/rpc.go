package util

import "github.com/ethereum/go-ethereum/ethclient"

type RPCDB struct {
	client *ethclient.Client
}
