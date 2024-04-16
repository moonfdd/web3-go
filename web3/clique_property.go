package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type Clique struct {
	c *rpc.Client
}

func NewClique(c *rpc.Client) *Clique {
	admin := &Clique{}
	admin.c = c
	return admin
}

// Proposals returns the current proposals the node tries to uphold and vote on.
// from API
// from web3 console clique
// property
func (c *Clique) Proposals(ctx context.Context) (map[common.Address]bool, error) {
	var result map[common.Address]bool
	err := c.c.CallContext(ctx, &result, "clique_proposals")
	return result, err
}
