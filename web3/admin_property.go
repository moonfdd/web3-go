package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

type Admin struct {
	c *rpc.Client
}

func NewAdmin(c *rpc.Client) *Admin {
	admin := &Admin{}
	admin.c = c
	return admin
}

// NodeInfo retrieves all the information we know about the host node at the
// protocol granularity.
// from adminAPI
// from web3ext.go
// property
func (admin *Admin) NodeInfo(ctx context.Context) (*p2p.NodeInfo, error) {
	var result *p2p.NodeInfo
	err := admin.c.CallContext(ctx, &result, "admin_nodeInfo")
	return result, err
}

// Peers retrieves all the information we know about each individual peer at the
// protocol granularity.
// from adminAPI
// from web3ext.go
// property
func (admin *Admin) Peers(ctx context.Context) ([]*p2p.PeerInfo, error) {
	var result []*p2p.PeerInfo
	err := admin.c.CallContext(ctx, &result, "admin_peers")
	return result, err
}

// Datadir retrieves the current data directory the node is using.
// from adminAPI
// from web3ext.go
// property
func (admin *Admin) Datadir(ctx context.Context) (string, error) {
	var result string
	err := admin.c.CallContext(ctx, &result, "admin_datadir")
	return result, err
}
