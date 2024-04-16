package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type Personal struct {
	c *rpc.Client
}

func NewPersonal(c *rpc.Client) *Personal {
	e := &Personal{}
	e.c = c
	return e
}

// rawWallet is a JSON representation of an accounts.Wallet interface, with its
// data contents extracted into plain fields.
type RawWallet struct {
	URL      string             `json:"url"`
	Status   string             `json:"status"`
	Failure  string             `json:"failure,omitempty"`
	Accounts []accounts.Account `json:"accounts,omitempty"`
}

// ListWallets will return a list of wallets this node manages.
// from PersonalAccountAPI
// from web3ext.go
// property
func (p *Personal) ListWallets(ctx context.Context) ([]RawWallet, error) {
	var result []RawWallet
	err := p.c.CallContext(ctx, &result, "personal_listWallets")
	return result, err
}

// ListAccounts will return a list of addresses for accounts this node manages.
// from PersonalAccountAPI
// from web3.js
// property
func (p *Personal) ListAccounts(ctx context.Context) ([]common.Address, error) {
	var result []common.Address
	err := p.c.CallContext(ctx, &result, "personal_listAccounts")
	return result, err
}
