package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/rpc"
)

// Discard drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
// from API
// from web3 console clique
// method
func (c *Clique) Discard(ctx context.Context, address common.Address) error {
	err := c.c.CallContext(ctx, nil, "clique_discard")
	return err
}

type BlockNumberOrHashOrRLP struct {
	*rpc.BlockNumberOrHash
	RLP hexutil.Bytes `json:"rlp,omitempty"`
}

// GetSigner returns the signer for a specific clique block.
// Can be called with a block number, a block hash or a rlp encoded blob.
// The RLP encoded blob can either be a block or a header.
// from API
// from web3 console clique
// method
func (c *Clique) GetSigner(ctx context.Context, rlpOrBlockNr *BlockNumberOrHashOrRLP) (common.Address, error) {
	var result common.Address
	err := c.c.CallContext(ctx, &result, "clique_getSigner", rlpOrBlockNr)
	return result, err
}

// GetSigners retrieves the list of authorized signers at the specified block.
// from API
// from web3 console clique
// method
func (c *Clique) GetSigners(ctx context.Context, number *rpc.BlockNumber) ([]common.Address, error) {
	var result []common.Address
	err := c.c.CallContext(ctx, &result, "clique_getSigners", number)
	return result, err
}

// GetSignersAtHash retrieves the list of authorized signers at the specified block.
// from API
// from web3 console clique
// method
func (c *Clique) GetSignersAtHash(ctx context.Context, hash common.Hash) ([]common.Address, error) {
	var result []common.Address
	err := c.c.CallContext(ctx, &result, "clique_getSignersAtHash", hash)
	return result, err
}

// GetSnapshot retrieves the state snapshot at a given block.
// from API
// from web3 console clique
// method
func (c *Clique) GetSnapshot(ctx context.Context, number *rpc.BlockNumber) (*clique.Snapshot, error) {
	var result *clique.Snapshot
	err := c.c.CallContext(ctx, &result, "clique_getSnapshot", number)
	return result, err
}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
// from API
// from web3 console clique
// method
func (c *Clique) GetSnapshotAtHash(ctx context.Context, hash common.Hash) (*clique.Snapshot, error) {
	var result *clique.Snapshot
	err := c.c.CallContext(ctx, &result, "clique_getSnapshotAtHash", hash)
	return result, err
}

// Propose injects a new authorization proposal that the signer will attempt to
// push through.
// from API
// from web3 console clique
// method
func (c *Clique) Propose(ctx context.Context, address common.Address, auth bool) error {
	err := c.c.CallContext(ctx, nil, "clique_propose", address, auth)
	return err
}

type Status struct {
	InturnPercent float64                `json:"inturnPercent"`
	SigningStatus map[common.Address]int `json:"sealerActivity"`
	NumBlocks     uint64                 `json:"numBlocks"`
}

// Status returns the status of the last N blocks,
// - the number of active signers,
// - the number of signers,
// - the percentage of in-turn blocks
// from API
// from web3 console clique
// method
func (c *Clique) Status(ctx context.Context) (*Status, error) {
	var result *Status
	err := c.c.CallContext(ctx, &result, "clique_status")
	return result, err
}
