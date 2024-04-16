package web3

import (
	"context"
	"runtime"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/tracers"

	"github.com/ethereum/go-ethereum/rpc"
)

// AccountRange enumerates all accounts in the given block and start point in paging request
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) AccountRange(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, start hexutil.Bytes, maxResults int, nocode, nostorage, incompletes bool) (state.Dump, error) {
	var result state.Dump
	err := d.c.CallContext(ctx, &result, "debug_accountRange", blockNrOrHash, start, maxResults, nocode, nostorage, incompletes)
	return result, err
}

// PrintBlock retrieves a block and returns its pretty printed form.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) PrintBlock(ctx context.Context, number uint64) (string, error) {
	var result string
	err := d.c.CallContext(ctx, &result, "debug_printBlock", number)
	return result, err
}

// GetRawHeader retrieves the RLP encoding for a single header.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetRawHeader(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_getRawHeader", blockNrOrHash)
	return result, err
}

// GetRawBlock retrieves the RLP encoded for a single block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetRawBlock(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_getRawBlock", blockNrOrHash)
	return result, err
}

// GetRawReceipts retrieves the binary-encoded receipts of a single block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetRawReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]hexutil.Bytes, error) {
	var result []hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_getRawReceipts", blockNrOrHash)
	return result, err
}

// GetRawTransaction returns the bytes of the transaction for the given hash.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetRawTransaction(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_getRawTransaction", hash)
	return result, err
}

// SetHead rewinds the head of the blockchain to a previous block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) SetHead(ctx context.Context, number hexutil.Uint64) error {
	err := d.c.CallContext(ctx, nil, "debug_setHead", number)
	return err
}

// debug_seedHash
// todo

// DumpBlock retrieves the entire state of the database at a given block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) DumpBlock(ctx context.Context, blockNr rpc.BlockNumber) (state.Dump, error) {
	var result state.Dump
	err := d.c.CallContext(ctx, &result, "debug_dumpBlock", blockNr)
	return result, err
}

// ChaindbProperty returns leveldb properties of the key-value database.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) ChaindbProperty(ctx context.Context, property string) (string, error) {
	var result string
	err := d.c.CallContext(ctx, &result, "debug_chaindbProperty", property)
	return result, err
}

// ChaindbCompact flattens the entire key-value database into a single level,
// removing all unused slots and merging all keys.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) ChaindbCompact(ctx context.Context) error {
	err := d.c.CallContext(ctx, nil, "debug_chaindbCompact")
	return err
}

// debug_verbosity
// todo

// Verbosity sets the log verbosity ceiling. The verbosity of individual packages
// and source files can be raised using Vmodule.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) Verbosity(ctx context.Context, level int) error {
	err := d.c.CallContext(ctx, nil, "debug_verbosity", level)
	return err
}

// Vmodule sets the log verbosity pattern. See package log for details on the
// pattern syntax.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) Vmodule(ctx context.Context, pattern string) error {
	err := d.c.CallContext(ctx, nil, "debug_vmodule", pattern)
	return err
}

// debug_backtraceAt
// todo

// Stacks returns a printed representation of the stacks of all goroutines. It
// also permits the following optional filters to be used:
//   - filter: boolean expression of packages to filter for
//
// from HandlerT
// from web3ext.go
// method
func (d *Debug) Stacks(ctx context.Context, filter *string) (string, error) {
	var result string
	err := d.c.CallContext(ctx, &result, "debug_stacks", filter)
	return result, err
}

// FreeOSMemory forces a garbage collection.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) FreeOSMemory(ctx context.Context) error {
	err := d.c.CallContext(ctx, nil, "debug_freeOSMemory")
	return err
}

// SetGCPercent sets the garbage collection target percentage. It returns the previous
// setting. A negative value disables GC.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) SetGCPercent(ctx context.Context, v int) (int, error) {
	var result int
	err := d.c.CallContext(ctx, &result, "debug_setGCPercent", v)
	return result, err
}

// MemStats returns detailed runtime memory statistics.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) MemStats(ctx context.Context) (*runtime.MemStats, error) {
	var result *runtime.MemStats
	err := d.c.CallContext(ctx, &result, "debug_memStats")
	return result, err
}

// GcStats returns GC statistics.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) GcStats(ctx context.Context) (*debug.GCStats, error) {
	var result *debug.GCStats
	err := d.c.CallContext(ctx, &result, "debug_gcStats")
	return result, err
}

// CpuProfile turns on CPU profiling for nsec seconds and writes
// profile data to file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) CpuProfile(ctx context.Context, file string, nsec uint) error {
	err := d.c.CallContext(ctx, nil, "debug_cpuProfile", file)
	return err
}

// StartCPUProfile turns on CPU profiling, writing to the given file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) StartCPUProfile(ctx context.Context, file string) error {
	err := d.c.CallContext(ctx, nil, "debug_startCPUProfile", file)
	return err
}

// StopCPUProfile stops an ongoing CPU profile.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) StopCPUProfile(ctx context.Context) error {
	err := d.c.CallContext(ctx, nil, "debug_stopCPUProfile")
	return err
}

// GoTrace turns on tracing for nsec seconds and writes
// trace data to file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) GoTrace(ctx context.Context, file string, nsec uint) error {
	err := d.c.CallContext(ctx, nil, "debug_goTrace", file, nsec)
	return err
}

// StartGoTrace turns on tracing, writing to the given file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) StartGoTrace(ctx context.Context, file string) error {
	err := d.c.CallContext(ctx, nil, "debug_startGoTrace", file)
	return err
}

// StopGoTrace stops an ongoing trace.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) StopGoTrace(ctx context.Context) error {
	err := d.c.CallContext(ctx, nil, "debug_stopGoTrace")
	return err
}

// BlockProfile turns on goroutine profiling for nsec seconds and writes profile data to
// file. It uses a profile rate of 1 for most accurate information. If a different rate is
// desired, set the rate and write the profile manually.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) BlockProfile(ctx context.Context, file string, nsec uint) error {
	err := d.c.CallContext(ctx, nil, "debug_blockProfile", file, nsec)
	return err
}

// SetBlockProfileRate sets the rate of goroutine block profile data collection.
// rate 0 disables block profiling.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) SetBlockProfileRate(ctx context.Context, rate int) error {
	err := d.c.CallContext(ctx, nil, "debug_setBlockProfileRate", rate)
	return err
}

// WriteBlockProfile writes a goroutine blocking profile to the given file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) WriteBlockProfile(ctx context.Context, file string) error {
	err := d.c.CallContext(ctx, nil, "debug_writeBlockProfile", file)
	return err
}

// MutexProfile turns on mutex profiling for nsec seconds and writes profile data to file.
// It uses a profile rate of 1 for most accurate information. If a different rate is
// desired, set the rate and write the profile manually.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) MutexProfile(ctx context.Context, file string, nsec uint) error {
	err := d.c.CallContext(ctx, nil, "debug_mutexProfile", file, nsec)
	return err
}

// SetMutexProfileFraction sets the rate of mutex profiling.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) SetMutexProfileFraction(ctx context.Context, rate int) error {
	err := d.c.CallContext(ctx, nil, "debug_setMutexProfileFraction", rate)
	return err
}

// WriteMutexProfile writes a goroutine blocking profile to the given file.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) WriteMutexProfile(ctx context.Context, file string) error {
	err := d.c.CallContext(ctx, nil, "debug_writeMutexProfile", file)
	return err
}

// WriteMemProfile writes an allocation profile to the given file.
// Note that the profiling rate cannot be set through the API,
// it must be set on the command line.
// from HandlerT
// from web3ext.go
// method
func (d *Debug) WriteMemProfile(ctx context.Context, file string) error {
	err := d.c.CallContext(ctx, nil, "debug_writeMemProfile", file)
	return err
}

// TraceBlock returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceBlock(ctx context.Context, blob hexutil.Bytes, config *tracers.TraceConfig) ([]*TxTraceResult, error) {
	var result []*TxTraceResult
	err := d.c.CallContext(ctx, &result, "debug_traceBlock", blob, config)
	return result, err
}

// txTraceResult is the result of a single transaction trace.
type TxTraceResult struct {
	TxHash common.Hash `json:"txHash"`           // transaction hash
	Result interface{} `json:"result,omitempty"` // Trace results produced by the tracer
	Error  string      `json:"error,omitempty"`  // Trace failure produced by the tracer
}

// TraceBlockFromFile returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceBlockFromFile(ctx context.Context, file string, config *tracers.TraceConfig) ([]*TxTraceResult, error) {
	var result []*TxTraceResult
	err := d.c.CallContext(ctx, &result, "debug_traceBlockFromFile", file, config)
	return result, err
}

// TraceBadBlock returns the structured logs created during the execution of
// EVM against a block pulled from the pool of bad ones and returns them as a JSON
// object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceBadBlock(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) ([]*TxTraceResult, error) {
	var result []*TxTraceResult
	err := d.c.CallContext(ctx, &result, "debug_traceBadBlock", hash, config)
	return result, err
}

// StandardTraceBadBlockToFile dumps the structured logs created during the
// execution of EVM against a block pulled from the pool of bad ones to the
// local file system and returns a list of files to the caller.
// from API
// from web3ext.go
// method
func (d *Debug) StandardTraceBadBlockToFile(ctx context.Context, hash common.Hash, config *tracers.StdTraceConfig) ([]string, error) {
	var result []string
	err := d.c.CallContext(ctx, &result, "debug_standardTraceBadBlockToFile", hash, config)
	return result, err
}

// IntermediateRoots executes a block (bad- or canon- or side-), and returns a list
// of intermediate roots: the stateroot after each transaction.
// from API
// from web3ext.go
// method
func (d *Debug) IntermediateRoots(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) ([]common.Hash, error) {
	var result []common.Hash
	err := d.c.CallContext(ctx, &result, "debug_intermediateRoots", hash, config)
	return result, err
}

// StandardTraceBlockToFile dumps the structured logs created during the
// execution of EVM to the local file system and returns a list of files
// to the caller.
// from API
// from web3ext.go
// method
func (d *Debug) StandardTraceBlockToFile(ctx context.Context, hash common.Hash, config *tracers.StdTraceConfig) ([]string, error) {
	var result []string
	err := d.c.CallContext(ctx, &result, "debug_standardTraceBlockToFile", hash, config)
	return result, err
}

// TraceBlockByNumber returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *tracers.TraceConfig) ([]*TxTraceResult, error) {
	var result []*TxTraceResult
	err := d.c.CallContext(ctx, &result, "debug_traceBlockByNumber", number, config)
	return result, err
}

// TraceBlockByHash returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceBlockByHash(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) ([]*TxTraceResult, error) {
	var result []*TxTraceResult
	err := d.c.CallContext(ctx, &result, "debug_traceBlockByHash", hash, config)
	return result, err
}

// TraceTransaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
// from API
// from web3ext.go
// method
func (d *Debug) TraceTransaction(ctx context.Context, hash common.Hash, config *tracers.TraceConfig) (interface{}, error) {
	var result interface{}
	err := d.c.CallContext(ctx, &result, "debug_traceTransaction", hash, config)
	return result, err
}

// TransactionArgs represents the arguments to construct a new transaction
// or a message call.
type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`

	// We accept "data" and "input" for backwards-compatibility reasons.
	// "input" is the newer name and should be preferred by clients.
	// Issue detail: https://github.com/ethereum/go-ethereum/issues/15628
	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`

	// Introduced by AccessListTxType transaction.
	AccessList *types.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big      `json:"chainId,omitempty"`

	// For BlobTxType
	BlobFeeCap *hexutil.Big  `json:"maxFeePerBlobGas"`
	BlobHashes []common.Hash `json:"blobVersionedHashes,omitempty"`

	// For BlobTxType transactions with blob sidecar
	Blobs       []kzg4844.Blob       `json:"blobs"`
	Commitments []kzg4844.Commitment `json:"commitments"`
	Proofs      []kzg4844.Proof      `json:"proofs"`

	// This configures whether blobs are allowed to be passed.
	blobSidecarAllowed bool
}

// TraceCall lets you trace a given eth_call. It collects the structured logs
// created during the execution of EVM if the given transaction was added on
// top of the provided block and returns them as a JSON object.
// If no transaction index is specified, the trace will be conducted on the state
// after executing the specified block. However, if a transaction index is provided,
// the trace will be conducted on the state after executing the specified transaction
// within the specified block.
// from API
// from web3ext.go
// method
func (d *Debug) TraceCall(ctx context.Context, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, config *tracers.TraceCallConfig) (interface{}, error) {
	var result interface{}
	err := d.c.CallContext(ctx, &result, "debug_traceCall", args, blockNrOrHash, config)
	return result, err
}

// Preimage is a debug API function that returns the preimage for a sha3 hash, if known.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_preimage", hash)
	return result, err
}

// GetBadBlocks returns a list of the last 'bad blocks' that the client has seen on the network
// and returns them as a JSON list of block hashes.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetBadBlocks(ctx context.Context) ([]*eth.BadBlockArgs, error) {
	var result []*eth.BadBlockArgs
	err := d.c.CallContext(ctx, &result, "debug_getBadBlocks")
	return result, err
}

// StorageRangeAt returns the storage at the given block height and transaction index.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) StorageRangeAt(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, txIndex int, contractAddress common.Address, keyStart hexutil.Bytes, maxResult int) (eth.StorageRangeResult, error) {
	var result eth.StorageRangeResult
	err := d.c.CallContext(ctx, &result, "debug_storageRangeAt", blockNrOrHash, txIndex, contractAddress, keyStart, maxResult)
	return result, err
}

// GetModifiedAccountsByNumber returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetModifiedAccountsByNumber(ctx context.Context, startNum uint64, endNum *uint64) ([]common.Address, error) {
	var result []common.Address
	err := d.c.CallContext(ctx, &result, "debug_getModifiedAccountsByNumber", startNum, endNum)
	return result, err
}

// GetModifiedAccountsByHash returns all accounts that have changed between the
// two blocks specified. A change is defined as a difference in nonce, balance,
// code hash, or storage hash.
//
// With one parameter, returns the list of accounts modified in the specified block.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetModifiedAccountsByHash(ctx context.Context, startHash common.Hash, endHash *common.Hash) ([]common.Address, error) {
	var result []common.Address
	err := d.c.CallContext(ctx, &result, "debug_getModifiedAccountsByHash", startHash, endHash)
	return result, err
}

// debug_freezeClient
// todo

// GetAccessibleState returns the first number where the node has accessible
// state on disk. Note this being the post-state of that block and the pre-state
// of the next block.
// The (from, to) parameters are the sequence of blocks to search, which can go
// either forwards or backwards
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetAccessibleState(ctx context.Context, from, to rpc.BlockNumber) (uint64, error) {
	var result uint64
	err := d.c.CallContext(ctx, &result, "debug_getAccessibleState", from, to)
	return result, err
}

// DbGet returns the raw value of a key stored in the database.
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) DbGet(ctx context.Context, key string) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_dbGet", key)
	return result, err
}

// DbAncient retrieves an ancient binary blob from the append-only immutable files.
// It is a mapping to the `AncientReaderOp.Ancient` method
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) DbAncient(ctx context.Context, kind string, number uint64) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := d.c.CallContext(ctx, &result, "debug_dbAncient", kind, number)
	return result, err
}

// DbAncients returns the ancient item numbers in the ancient store.
// It is a mapping to the `AncientReaderOp.Ancients` method
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) DbAncients(ctx context.Context) (uint64, error) {
	var result uint64
	err := d.c.CallContext(ctx, &result, "debug_dbAncients")
	return result, err
}

// SetTrieFlushInterval configures how often in-memory tries are persisted
// to disk. The value is in terms of block processing time, not wall clock.
// If the value is shorter than the block generation time, or even 0 or negative,
// the node will flush trie after processing each block (effectively archive mode).
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) SetTrieFlushInterval(ctx context.Context, interval string) error {

	err := d.c.CallContext(ctx, nil, "debug_setTrieFlushInterval", interval)
	return err
}

// GetTrieFlushInterval gets the current value of in-memory trie flush interval
// from DebugAPI
// from web3ext.go
// method
func (d *Debug) GetTrieFlushInterval(ctx context.Context) (string, error) {
	var result string
	err := d.c.CallContext(ctx, &result, "debug_getTrieFlushInterval")
	return result, err
}
