package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Start starts the miner with the given number of threads. If threads is nil,
// the number of workers started is equal to the number of logical CPUs that are
// usable by this process. If mining is already running, this method adjust the
// number of threads allowed to use and updates the minimum price required by the
// transaction pool.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) Start(ctx context.Context) error {

	err := m.c.CallContext(ctx, nil, "miner_start")
	return err
}

// Stop terminates the miner, both at the consensus engine level as well as at
// the block creation level.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) Stop(ctx context.Context) error {
	err := m.c.CallContext(ctx, nil, "miner_stop")
	return err
}

// SetEtherbase sets the etherbase of the miner.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) SetEtherbase(ctx context.Context, etherbase common.Address) (bool, error) {
	var result bool
	err := m.c.CallContext(ctx, &result, "miner_setEtherbase", etherbase)
	return result, err
}

// SetExtra sets the extra data string that is included when this miner mines a block.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) SetExtra(ctx context.Context, extra string) (bool, error) {
	var result bool
	err := m.c.CallContext(ctx, &result, "miner_setExtra", extra)
	return result, err
}

// SetGasPrice sets the minimum accepted gas price for the miner.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) SetGasPrice(ctx context.Context, gasPrice hexutil.Big) (bool, error) {
	var result bool
	err := m.c.CallContext(ctx, &result, "miner_setGasPrice", gasPrice)
	return result, err
}

// SetGasLimit sets the gaslimit to target towards during mining.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) SetGasLimit(ctx context.Context, gasLimit hexutil.Uint64) (bool, error) {
	var result bool
	err := m.c.CallContext(ctx, &result, "miner_setGasLimit", gasLimit)
	return result, err
}

// SetRecommitInterval updates the interval for miner sealing work recommitting.
// from MinerAPI
// from web3ext.go
// method
func (m *Miner) SetRecommitInterval(ctx context.Context, interval int) error {
	err := m.c.CallContext(ctx, nil, "miner_setRecommitInterval")
	return err
}

// Hashrate returns the POW hashrate.
// from EthereumAPI
// from web3ext.go
// method
func (m *Miner) GetHashrate(ctx context.Context) (hexutil.Uint64, error) {
	var result hexutil.Uint64
	err := m.c.CallContext(ctx, &result, "miner_getHashrate")
	return result, err
}
