package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type TxPool struct {
	c *rpc.Client
}

func NewTxPool(c *rpc.Client) *TxPool {
	e := &TxPool{}
	e.c = c
	return e
}

// Content returns the transactions contained within the transaction pool.
// from TxPoolAPI
// from web3ext.go
// property
func (t *TxPool) Content(ctx context.Context) (map[string]map[string]map[string]*RPCTransaction, error) {
	var result map[string]map[string]map[string]*RPCTransaction
	err := t.c.CallContext(ctx, &result, "txpool_content")
	return result, err
}

// Inspect retrieves the content of the transaction pool and flattens it into an
// easily inspectable list.
// from TxPoolAPI
// from web3ext.go
// property
func (t *TxPool) Inspect(ctx context.Context) (map[string]map[string]map[string]string, error) {
	var result map[string]map[string]map[string]string
	err := t.c.CallContext(ctx, &result, "txpool_inspect")
	return result, err
}

// Status returns the number of pending and queued transaction in the pool.
// from TxPoolAPI
// from web3ext.go
// property
func (t *TxPool) Status(ctx context.Context) (map[string]hexutil.Uint, error) {
	var result map[string]hexutil.Uint
	err := t.c.CallContext(ctx, &result, "txpool_status")
	return result, err
}

// this is test method
func (t *TxPool) Test(ctx context.Context, a ...interface{}) (interface{}, error) {
	var result interface{}
	err := t.c.CallContext(ctx, &result, "web3_version_network")
	// err := t.c.CallContext(ctx, &result, "eth_coinbase")
	return result, err
}
