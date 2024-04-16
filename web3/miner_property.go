package web3

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type Miner struct {
	c *rpc.Client
}

func NewMiner(c *rpc.Client) *Miner {
	e := &Miner{}
	e.c = c
	return e
}
