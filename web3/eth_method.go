package web3

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
)

// ChainID retrieves the current chain ID for transaction replay protection.
// from Client
// from web3ext.go
// method
func (e *Eth) ChainID(ctx context.Context) (*big.Int, error) {
	var result *big.Int
	err := e.c.CallContext(ctx, &result, "eth_chainId")
	return result, err
}

// Sign calculates an ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message).
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The account associated with addr must be unlocked.
//
// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sign
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) Sign(ctx context.Context, addr common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_sign", addr, data)
	return result, err
}

// Resend accepts an existing transaction and a new gas price and limit. It will remove
// the given transaction from the pool and reinsert it with the new gas price and limit.
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) Resend(ctx context.Context, sendArgs TransactionArgs, gasPrice *hexutil.Big, gasLimit *hexutil.Uint64) (common.Hash, error) {
	var result common.Hash
	err := e.c.CallContext(ctx, &result, "eth_resend", sendArgs, gasPrice, gasLimit)
	return result, err
}

// SignTransactionResult represents a RLP encoded signed transaction.
type SignTransactionResult struct {
	Raw hexutil.Bytes      `json:"raw"`
	Tx  *types.Transaction `json:"tx"`
}

// SignTransaction will sign the given transaction with the from account.
// The node needs to have the private key of the account corresponding with
// the given from address and it needs to be unlocked.
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) SignTransaction(ctx context.Context, args TransactionArgs) (*SignTransactionResult, error) {
	var result *SignTransactionResult
	err := e.c.CallContext(ctx, &result, "eth_signTransaction", args)
	return result, err
}

// StateOverride is the collection of overridden accounts.
type StateOverride map[common.Address]OverrideAccount

// OverrideAccount indicates the overriding fields of account during the execution
// of a message call.
// Note, state and stateDiff can't be specified at the same time. If state is
// set, message execution will only use the data in the given state. Otherwise
// if statDiff is set, all diff will be applied first and then execute the call
// message.
type OverrideAccount struct {
	Nonce     *hexutil.Uint64              `json:"nonce"`
	Code      *hexutil.Bytes               `json:"code"`
	Balance   **hexutil.Big                `json:"balance"`
	State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff *map[common.Hash]common.Hash `json:"stateDiff"`
}

// EstimateGas returns the lowest possible gas limit that allows the transaction to run
// successfully at block `blockNrOrHash`, or the latest block if `blockNrOrHash` is unspecified. It
// returns error if the transaction would revert or if there are unexpected failures. The returned
// value is capped by both `args.Gas` (if non-nil & non-zero) and the backend's RPCGasCap
// configuration (if non-zero).
// Note: Required blob gas is not computed in this method.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) EstimateGas(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *StateOverride) (hexutil.Uint64, error) {
	var result hexutil.Uint64
	err := e.c.CallContext(ctx, &result, "eth_estimateGas", args, blockNrOrHash, overrides)
	return result, err
}

// SubmitTransaction is a helper function that submits tx to txPool and logs a message.
// from web3ext.go
// method
func (e *Eth) SubmitTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	var result common.Hash
	err := e.c.CallContext(ctx, &result, "eth_submitTransaction", tx)
	return result, err
}

// FillTransaction fills the defaults (nonce, gas, gasPrice or 1559 fields)
// on a given unsigned transaction, and returns it to the caller for further
// processing (signing + broadcast).
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) FillTransaction(ctx context.Context, args TransactionArgs) (*SignTransactionResult, error) {
	var result *SignTransactionResult
	err := e.c.CallContext(ctx, &result, "eth_fillTransaction", args)
	return result, err
}

// GetHeaderByNumber returns the requested canonical block header.
//   - When blockNr is -1 the chain pending header is returned.
//   - When blockNr is -2 the chain latest header is returned.
//   - When blockNr is -3 the chain finalized header is returned.
//   - When blockNr is -4 the chain safe header is returned.
//
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetHeaderByNumber(ctx context.Context, number rpc.BlockNumber) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := e.c.CallContext(ctx, &result, "eth_getHeaderByNumber", number)
	return result, err

}

// GetHeaderByHash returns the requested header by hash.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetHeaderByHash(ctx context.Context, hash common.Hash) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := e.c.CallContext(ctx, &result, "eth_getHeaderByHash", hash)
	return result, err

}

// GetBlockByNumber returns the requested canonical block.
//   - When blockNr is -1 the chain pending block is returned.
//   - When blockNr is -2 the chain latest block is returned.
//   - When blockNr is -3 the chain finalized block is returned.
//   - When blockNr is -4 the chain safe block is returned.
//   - When fullTx is true all transactions in the block are returned, otherwise
//     only the transaction hash is returned.
//
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := e.c.CallContext(ctx, &result, "eth_getBlockByNumber", number, fullTx)
	return result, err
}

// GetBlockByHash returns the requested block. When fullTx is true all transactions in the block are returned in full
// detail, otherwise only the transaction hash is returned.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetBlockByHash(ctx context.Context, hash common.Hash, fullTx bool) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := e.c.CallContext(ctx, &result, "eth_getBlockByHash", hash, fullTx)
	return result, err

}

// GetRawTransaction returns the bytes of the transaction for the given hash.
// from DebugAPI
// from web3ext.go
// method
func (e *Eth) GetRawTransaction(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_getRawTransactionByHash", hash)
	return result, err
}
func (e *Eth) GetRawTransactionByHash(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_getRawTransactionByHash", hash)
	return result, err
}

// getRawTransactionFromBlock
// todo

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) GetRawTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_getRawTransactionByBlockHashAndIndex", blockHash, index)
	return result, err
}

// GetRawTransactionByBlockNumberAndIndex returns the bytes of the transaction for the given block number and index.
// from TransactionAPI
// from web3ext.go
// method
func (e *Eth) GetRawTransactionByBlockNumberAndIndex(ctx context.Context, blockNr rpc.BlockNumber, index hexutil.Uint) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_getRawTransactionByBlockNumberAndIndex", blockNr, index)
	return result, err
}

// AccountResult structs for GetProof
type AccountResult struct {
	Address      common.Address  `json:"address"`
	AccountProof []string        `json:"accountProof"`
	Balance      *hexutil.Big    `json:"balance"`
	CodeHash     common.Hash     `json:"codeHash"`
	Nonce        hexutil.Uint64  `json:"nonce"`
	StorageHash  common.Hash     `json:"storageHash"`
	StorageProof []StorageResult `json:"storageProof"`
}

type StorageResult struct {
	Key   string       `json:"key"`
	Value *hexutil.Big `json:"value"`
	Proof []string     `json:"proof"`
}

// GetProof returns the Merkle-proof for a given account and optionally some storage keys.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*AccountResult, error) {
	var result *AccountResult
	err := e.c.CallContext(ctx, &result, "eth_getProof", address, storageKeys, blockNrOrHash)
	return result, err
}

// accessListResult returns an optional accesslist
// It's the result of the `debug_createAccessList` RPC call.
// It contains an error if the transaction itself failed.
type AccessListResult struct {
	Accesslist *types.AccessList `json:"accessList"`
	Error      string            `json:"error,omitempty"`
	GasUsed    hexutil.Uint64    `json:"gasUsed"`
}

// CreateAccessList creates an EIP-2930 type AccessList for the given transaction.
// Reexec and BlockNrOrHash can be specified to create the accessList on top of a certain state.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) CreateAccessList(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash) (*AccessListResult, error) {
	var result *AccessListResult
	err := e.c.CallContext(ctx, &result, "eth_getProof", args, blockNrOrHash)
	return result, err
}

type FeeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

// FeeHistory returns the fee market history.
// from EthereumAPI
// from web3ext.go
// method
func (e *Eth) FeeHistory(ctx context.Context, blockCount math.HexOrDecimal64, lastBlock rpc.BlockNumber, rewardPercentiles []float64) (*FeeHistoryResult, error) {
	var result *FeeHistoryResult
	err := e.c.CallContext(ctx, &result, "eth_feeHistory", blockCount, lastBlock, rewardPercentiles)
	return result, err
}

// GetLogs returns logs matching the given argument that are stored within the state.
// from FilterAPI
// from web3ext.go
// method
func (e *Eth) GetLogs(ctx context.Context, crit filters.FilterCriteria) ([]*types.Log, error) {

	var result []*types.Log
	err := e.c.CallContext(ctx, &result, "eth_getLogs", crit)
	return result, err
}

// BlockOverrides is a set of header fields to override.
type BlockOverrides struct {
	Number      *hexutil.Big
	Difficulty  *hexutil.Big
	Time        *hexutil.Uint64
	GasLimit    *hexutil.Uint64
	Coinbase    *common.Address
	Random      *common.Hash
	BaseFee     *hexutil.Big
	BlobBaseFee *hexutil.Big
}

// Call executes the given transaction on the state for the given block number.
//
// Additionally, the caller can specify a batch of contract for fields overriding.
//
// Note, this function doesn't make and changes in the state/blockchain and is
// useful to execute and retrieve values.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) Call(ctx context.Context, args TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *StateOverride, blockOverrides *BlockOverrides) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := e.c.CallContext(ctx, &result, "eth_call", args, blockNrOrHash, overrides, blockOverrides)
	return result, err
}

// GetBlockReceipts returns the block receipts for the given block hash or number or tag.
// from BlockChainAPI
// from web3ext.go
// method
func (e *Eth) GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]map[string]interface{}, error) {

	var result []map[string]interface{}
	err := e.c.CallContext(ctx, &result, "eth_getBlockReceipts", blockNrOrHash)
	return result, err
}

// TransactionSender wraps transaction sending. The SendTransaction method injects a
// signed transaction into the pending transaction pool for execution. If the transaction
// was a contract creation, the TransactionReceipt method can be used to retrieve the
// contract address after the transaction has been mined.
//
// The transaction must be signed and have a valid nonce to be included. Consumers of the
// API can use package accounts to maintain local private keys and need can retrieve the
// next available nonce using PendingNonceAt.
// from TransactionSender
// from web3.js
// method
func (e *Eth) SendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	var result common.Hash
	err := e.c.CallContext(ctx, &result, "eth_sendTransaction", tx)
	return result, err
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

// NewBlockFilter creates a filter that fetches blocks that are imported into the chain.
// It is part of the filter package since polling goes with eth_getFilterChanges.
// from FilterAPI
// from web3.js
// method
func (e *Eth) NewBlockFilter(ctx context.Context) (rpc.ID, error) {
	var result rpc.ID
	err := e.c.CallContext(ctx, &result, "eth_newBlockFilter")
	return result, err
}

// NewPendingTransactionFilter creates a filter that fetches pending transactions
// as transactions enter the pending state.
//
// It is part of the filter package because this filter can be used through the
// `eth_getFilterChanges` polling method that is also used for log filters.
// from FilterAPI
// from web3.js
// method
func (e *Eth) NewPendingTransactionFilter(ctx context.Context, fullTx *bool) (rpc.ID, error) {
	var result rpc.ID
	err := e.c.CallContext(ctx, &result, "eth_newPendingTransactionFilter", fullTx)
	return result, err
}

// NewFilter creates a new filter and returns the filter id. It can be
// used to retrieve logs when the state changes. This method cannot be
// used to fetch logs that are already stored in the state.
//
// Default criteria for the from and to block are "latest".
// Using "latest" as block number will return logs for mined blocks.
// Using "pending" as block number returns logs for not yet mined (pending) blocks.
// In case logs are removed (chain reorg) previously returned logs are returned
// again but with the removed property set to true.
//
// In case "fromBlock" > "toBlock" an error is returned.
// from FilterAPI
// from web3.js
// method
func (e *Eth) NewFilter(ctx context.Context, crit filters.FilterCriteria) (rpc.ID, error) {
	var result rpc.ID
	err := e.c.CallContext(ctx, &result, "eth_newFilter", crit)
	return result, err
}

// UninstallFilter removes the filter with the given filter id.
// from FilterAPI
// from web3.js
// method
func (e *Eth) UninstallFilter(ctx context.Context, id rpc.ID) (bool, error) {
	var result bool
	err := e.c.CallContext(ctx, &result, "eth_uninstallFilter", id)
	return result, err
}

// GetFilterLogs returns the logs for the filter with the given id.
// If the filter could not be found an empty array of logs is returned.
// from FilterAPI
// from web3.js
// method
func (e *Eth) GetFilterLogs(ctx context.Context, id rpc.ID) ([]*types.Log, error) {
	var result []*types.Log
	err := e.c.CallContext(ctx, &result, "eth_getFilterLogs", id)
	return result, err
}
