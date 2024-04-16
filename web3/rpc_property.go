package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/rpc"
)

type Rpc struct {
	c *rpc.Client
}

func NewRpc(c *rpc.Client) *Rpc {
	e := &Rpc{}
	e.c = c
	return e
}

// Modules returns the list of RPC services with their version number
// from RPCService
// from web3ext.go
// property
func (r *Rpc) Modules(ctx context.Context) (map[string]string, error) {
	var result map[string]string
	err := r.c.CallContext(ctx, &result, "rpc_modules")
	return result, err
}
