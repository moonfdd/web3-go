package web3

import (
	"context"

	"github.com/ethereum/go-ethereum/log"
)

// AddPeer requests connecting to a remote node, and also maintaining the new
// connection at all times, even reconnecting if it is lost.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) AddPeer(ctx context.Context, url string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_addPeer", url)
	return result, err
}

// RemovePeer disconnects from a remote node if the connection exists
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) RemovePeer(ctx context.Context, url string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_removePeer", url)
	return result, err
}

// AddTrustedPeer allows a remote node to always connect, even if slots are full
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) AddTrustedPeer(ctx context.Context, url string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_addTrustedPeer", url)
	return result, err
}

// RemoveTrustedPeer removes a remote node from the trusted peer set, but it
// does not disconnect it automatically.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) RemoveTrustedPeer(ctx context.Context, url string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_removeTrustedPeer", url)
	return result, err
}

// ExportChain exports the current blockchain into a local file,
// or a range of blocks if first and last are non-nil.
// from AdminAPI
// from web3ext.go
// method
func (admin *Admin) ExportChain(ctx context.Context, file string, first *uint64, last *uint64) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_exportChain", file, first, last)
	return result, err
}

// ImportChain imports a blockchain from a local file.
// from AdminAPI
// from web3ext.go
// method
func (admin *Admin) ImportChain(ctx context.Context, file string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_importChain", file)
	return result, err
}

// admin_sleepBlocks
// from web3ext.go
// method
// todo

// StartHTTP starts the HTTP RPC API server.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StartHTTP(ctx context.Context, host *string, port *int, cors *string, apis *string, vhosts *string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_startHTTP", host, port, cors, apis, vhosts)
	return result, err
}

// StopHTTP shuts down the HTTP server.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StopHTTP(ctx context.Context) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_stopHTTP")
	return result, err
}

// StartRPC starts the HTTP RPC API server.
// Deprecated: use StartHTTP instead.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StartRPC(ctx context.Context, host *string, port *int, cors *string, apis *string, vhosts *string) (bool, error) {
	log.Warn("Deprecation warning", "method", "admin.StartRPC", "use-instead", "admin.StartHTTP")
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_startRPC", host, port, cors, apis, vhosts)
	return result, err
}

// StopRPC shuts down the HTTP server.
// Deprecated: use StopHTTP instead.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StopRPC(ctx context.Context) (bool, error) {
	log.Warn("Deprecation warning", "method", "admin.StopRPC", "use-instead", "admin.StopHTTP")
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_stopRPC")
	return result, err
}

// StartWS starts the websocket RPC API server.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StartWS(ctx context.Context, host *string, port *int, allowedOrigins *string, apis *string) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_startWS", host, port, allowedOrigins, apis)
	return result, err
}

// StopWS terminates all WebSocket servers.
// from adminAPI
// from web3ext.go
// method
func (admin *Admin) StopWS(ctx context.Context) (bool, error) {
	var result bool
	err := admin.c.CallContext(ctx, &result, "admin_stopWS")
	return result, err
}
