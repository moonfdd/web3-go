package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

// ContentFrom returns the transactions contained within the transaction pool.
// from TxPoolAPI
// from web3ext.go
// method
func (t *TxPool) ContentFrom(ctx context.Context, addr common.Address) (map[string]map[string]*RPCTransaction, error) {
	var result map[string]map[string]*RPCTransaction
	err := t.c.CallContext(ctx, &result, "txpool_contentFrom", addr)
	return result, err
}
