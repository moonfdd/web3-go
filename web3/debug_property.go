package web3

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type Debug struct {
	c *rpc.Client
}

func NewDebug(c *rpc.Client) *Debug {
	d := &Debug{}
	d.c = c
	return d
}
