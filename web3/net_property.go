package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type Net struct {
	c *rpc.Client
}

func NewNet(c *rpc.Client) *Net {
	e := &Net{}
	e.c = c
	return e
}

// Returns the external api version. This method does not require user acceptance. Available methods are
// available via enumeration anyway, and this info does not contain user-specific data
// from SignerAPI
// from web3ext.go
// property
func (n *Net) Version(ctx context.Context) (string, error) {
	var result string
	err := n.c.CallContext(ctx, &result, "net_version")
	return result, err
}

// Listening returns an indication if the node is listening for network connections.
// from NetAPI
// from web3.js
// property
func (n *Net) Listening(ctx context.Context) (bool, error) {
	var result bool
	err := n.c.CallContext(ctx, &result, "net_listening")
	return result, err
}

// PeerCount returns the number of connected peers
// from NetAPI
// from web3.js
// property
func (n *Net) PeerCount(ctx context.Context) (hexutil.Uint, error) {
	var result hexutil.Uint
	err := n.c.CallContext(ctx, &result, "net_peerCount")
	return result, err
}
