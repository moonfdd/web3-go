package web3

import "github.com/ethereum/go-ethereum/rpc"

type Web3 struct {
	c        *rpc.Client
	Admin    *Admin
	Clique   *Clique
	Debug    *Debug
	Eth      *Eth
	Miner    *Miner
	Net      *Net
	Personal *Personal
	Rpc      *Rpc
	TxPool   *TxPool
}

func NewWeb3(c *rpc.Client) *Web3 {
	web3 := &Web3{}
	web3.c = c
	web3.Admin = NewAdmin(c)
	web3.Clique = NewClique(c)
	web3.Debug = NewDebug(c)
	web3.Eth = NewEth(c)
	web3.Miner = NewMiner(c)
	web3.Net = NewNet(c)
	web3.Personal = NewPersonal(c)
	web3.Rpc = NewRpc(c)
	web3.TxPool = NewTxPool(c)
	return web3
}
