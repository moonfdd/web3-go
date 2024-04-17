package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type Eth struct {
	c *rpc.Client
}

// eth_compileSolidity
// todo

// eth_compileLLL
// todo

// eth_compileSerpent
// todo

// eth_submitWork
// todo

// eth_getWork
// todo

// defaultAccount 属性
// defaultBlock
// eth_protocolVersion
// eth_contract
// createAccessList
// filter
// getBalance
// getBlock
// getBlockTransactionCount
// getBlockUncleCount
// getCode
// getCompilers
// getRawTransactionFromBlock
// getStorageAt
// getTransaction
// getTransactionCount
// getTransactionFromBlock
// getTransactionReceipt
// getUncle
// iban
// icapNamereg
// namereg
// sendIBANTransaction
// sendRawTransaction
// submitWork

func NewEth(c *rpc.Client) *Eth {
	e := &Eth{}
	e.c = c
	return e
}

// RPCTransaction represents a transaction that will serialize to the RPC representation of a transaction
type RPCTransaction struct {
	BlockHash           *common.Hash      `json:"blockHash"`
	BlockNumber         *hexutil.Big      `json:"blockNumber"`
	From                common.Address    `json:"from"`
	Gas                 hexutil.Uint64    `json:"gas"`
	GasPrice            *hexutil.Big      `json:"gasPrice"`
	GasFeeCap           *hexutil.Big      `json:"maxFeePerGas,omitempty"`
	GasTipCap           *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
	MaxFeePerBlobGas    *hexutil.Big      `json:"maxFeePerBlobGas,omitempty"`
	Hash                common.Hash       `json:"hash"`
	Input               hexutil.Bytes     `json:"input"`
	Nonce               hexutil.Uint64    `json:"nonce"`
	To                  *common.Address   `json:"to"`
	TransactionIndex    *hexutil.Uint64   `json:"transactionIndex"`
	Value               *hexutil.Big      `json:"value"`
	Type                hexutil.Uint64    `json:"type"`
	Accesses            *types.AccessList `json:"accessList,omitempty"`
	ChainID             *hexutil.Big      `json:"chainId,omitempty"`
	BlobVersionedHashes []common.Hash     `json:"blobVersionedHashes,omitempty"`
	V                   *hexutil.Big      `json:"v"`
	R                   *hexutil.Big      `json:"r"`
	S                   *hexutil.Big      `json:"s"`
	YParity             *hexutil.Uint64   `json:"yParity,omitempty"`
}

// PendingTransactions returns the transactions that are in the transaction pool
// and have a from address that is one of the accounts this node manages.
// from TransactionAPI
// from web3ext.go
// property
func (e *Eth) PendingTransactions(ctx context.Context) ([]*RPCTransaction, error) {
	var result []*RPCTransaction
	err := e.c.CallContext(ctx, &result, "eth_pendingTransactions")
	return result, err
}

// MaxPriorityFeePerGas returns a suggestion for a gas tip cap for dynamic fee transactions.
// from EthereumAPI
// from web3ext.go
// property
func (e *Eth) MaxPriorityFeePerGas(ctx context.Context) (*hexutil.Big, error) {
	var result *hexutil.Big
	err := e.c.CallContext(ctx, &result, "eth_maxPriorityFeePerGas")
	return result, err
}

// Coinbase is the address that mining rewards will be sent to (alias for Etherbase).
// from EthereumAPI
// from web3.js
// property
func (e *Eth) Coinbase(ctx context.Context) (common.Address, error) {
	var result common.Address
	err := e.c.CallContext(ctx, &result, "eth_coinbase")
	return result, err
}

// Mining returns an indication if this node is currently mining.
// from EthereumAPI
// from web3.js
// property
func (e *Eth) Mining(ctx context.Context) (bool, error) {
	var result bool
	err := e.c.CallContext(ctx, &result, "eth_mining")
	return result, err
}

// Hashrate returns the POW hashrate.
// from EthereumAPI
// from web3.js
// property
func (e *Eth) Hashrate(ctx context.Context) (hexutil.Uint64, error) {
	var result hexutil.Uint64
	err := e.c.CallContext(ctx, &result, "eth_hashrate")
	return result, err
}

// from EthereumAPI
// from web3.js
// property
func (e *Eth) Syncing(ctx context.Context) (interface{}, error) {
	var result interface{}
	err := e.c.CallContext(ctx, &result, "eth_syncing")
	return result, err
}

// GasPrice returns a suggestion for a gas price for legacy transactions.
// from EthereumAPI
// from web3.js
// property
func (e *Eth) GasPrice(ctx context.Context) (*hexutil.Big, error) {
	var result *hexutil.Big
	err := e.c.CallContext(ctx, &result, "eth_gasPrice")
	return result, err
}

// Accounts returns the collection of accounts this node manages.
// from EthereumAccountAPI
// from web3.js
// property
func (e *Eth) Accounts(ctx context.Context) ([]common.Address, error) {
	var result []common.Address
	err := e.c.CallContext(ctx, &result, "eth_accounts")
	return result, err
}

// BlockNumberReader provides access to the current block number.
// from BlockNumberReader
// from web3.js
// property

func (e *Eth) BlockNumber(ctx context.Context) (uint64, error) {
	var result uint64
	err := e.c.CallContext(ctx, &result, "eth_blockNumber")
	return result, err
}

//eth_protocolVersion
//todo
