package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// ImportRawKey stores the given hex encoded ECDSA key into the key directory,
// encrypting it with the passphrase.
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) ImportRawKey(ctx context.Context, privkey string, password string) (common.Address, error) {
	var result common.Address
	err := p.c.CallContext(ctx, &result, "personal_importRawKey", privkey, password)
	return result, err
}

// Sign calculates an Ethereum ECDSA signature for:
// keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The key used to calculate the signature is decrypted with the given password.
//
// https://geth.ethereum.org/docs/interacting-with-geth/rpc/ns-personal#personal-sign
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) Sign(ctx context.Context, data hexutil.Bytes, addr common.Address, passwd string) (hexutil.Bytes, error) {
	var result hexutil.Bytes
	err := p.c.CallContext(ctx, &result, "personal_sign", data, addr, passwd)
	return result, err
}

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be 27 or 28 for legacy reasons.
//
// https://geth.ethereum.org/docs/interacting-with-geth/rpc/ns-personal#personal-ecrecover
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) EcRecover(ctx context.Context, data, sig hexutil.Bytes) (common.Address, error) {
	var result common.Address
	err := p.c.CallContext(ctx, &result, "personal_ecRecover", data, sig)
	return result, err
}

// OpenWallet initiates a hardware wallet opening procedure, establishing a USB
// connection and attempting to authenticate via the provided passphrase. Note,
// the method may return an extra challenge requiring a second open (e.g. the
// Trezor PIN matrix challenge).
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) OpenWallet(ctx context.Context, url string, passphrase *string) error {
	err := p.c.CallContext(ctx, nil, "personal_openWallet", url, passphrase)
	return err
}

// DeriveAccount requests an HD wallet to derive a new account, optionally pinning
// it for later reuse.
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) DeriveAccount(ctx context.Context, url string, path string, pin *bool) (accounts.Account, error) {
	var result accounts.Account
	err := p.c.CallContext(ctx, &result, "personal_deriveAccount", url, path, pin)
	return result, err
}

// SignTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.From. If the given passwd isn't
// able to decrypt the key it fails. The transaction is returned in RLP-form, not broadcast
// to other nodes
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) SignTransaction(ctx context.Context, args TransactionArgs, passwd string) (*SignTransactionResult, error) {
	var result *SignTransactionResult
	err := p.c.CallContext(ctx, &result, "personal_signTransaction", args, passwd)
	return result, err
}

// Unpair deletes a pairing between wallet and geth.
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) Unpair(ctx context.Context, url string, pin string) error {
	err := p.c.CallContext(ctx, nil, "personal_unpair", url, pin)
	return err
}

// InitializeWallet initializes a new wallet at the provided URL, by generating and returning a new private key.
// from PersonalAccountAPI
// from web3ext.go
// method
func (p *Personal) InitializeWallet(ctx context.Context, url string) (string, error) {
	var result string
	err := p.c.CallContext(ctx, &result, "personal_initializeWallet", url)
	return result, err
}

// NewAccount will create a new account and returns the address for the new account.
// from PersonalAccountAPI
// from web3.js
// method
func (p *Personal) NewAccount(ctx context.Context, password string) (common.AddressEIP55, error) {
	var result common.AddressEIP55
	err := p.c.CallContext(ctx, &result, "personal_newAccount", password)
	return result, err
}

// UnlockAccount will unlock the account associated with the given address with
// the given password for duration seconds. If duration is nil it will use a
// default of 300 seconds. It returns an indication if the account was unlocked.
// from PersonalAccountAPI
// from web3.js
// method
func (p *Personal) UnlockAccount(ctx context.Context, addr common.Address, password string, duration *uint64) (bool, error) {
	var result bool
	err := p.c.CallContext(ctx, &result, "personal_unlockAccount", addr, password, duration)
	return result, err
}

// SendTransaction will create a transaction from the given arguments and
// tries to sign it with the key associated with args.From. If the given
// passwd isn't able to decrypt the key it fails.
// from PersonalAccountAPI
// from web3.js
// method
func (p *Personal) SendTransaction(ctx context.Context, args TransactionArgs, passwd string) (common.Hash, error) {
	var result common.Hash
	err := p.c.CallContext(ctx, &result, "personal_sendTransaction", args, passwd)
	return result, err
}

// LockAccount will lock the account associated with the given address when it's unlocked.
// from PersonalAccountAPI
// from web3.js
// method
func (p *Personal) LockAccount(ctx context.Context, addr common.Address) (bool, error) {
	var result bool
	err := p.c.CallContext(ctx, &result, "personal_lockAccount", addr)
	return result, err
}
