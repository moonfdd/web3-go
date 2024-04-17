package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/moonfdd/web3-go/web3"
)

func main() {
	c, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	web3 := web3.NewWeb3(c)
	fmt.Println(web3)
	if false {
		fmt.Println("NodeInfo------------")
		fmt.Println(web3.Admin.NodeInfo(context.Background()))
		fmt.Println("Datadir------------")
		fmt.Println(web3.Admin.Datadir(context.Background()))
		fmt.Println("AddPeer------------")
		fmt.Println(web3.Admin.AddPeer(context.Background(), "enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303"))
		fmt.Println("Peers------------")
		fmt.Println(web3.Admin.Peers(context.Background()))
		return
	}
	if false {
		fmt.Println("GetTrieFlushInterval------------")
		fmt.Println(web3.Debug.GetTrieFlushInterval(context.Background()))
	}
	if false {
		fmt.Println("MaxPriorityFeePerGas------------")
		fmt.Println(web3.Eth.MaxPriorityFeePerGas(context.Background()))
		fmt.Println("PendingTransactions------------")
		fmt.Println(web3.Eth.PendingTransactions(context.Background()))
	}
	if false {
		fmt.Println("Start------------")
		fmt.Println(web3.Miner.Start(context.Background()))
		fmt.Println("GetHashrate------------")
		fmt.Println(web3.Miner.GetHashrate(context.Background()))
	}
	if false {
		fmt.Println("Version------------")
		fmt.Println(web3.Net.Version(context.Background()))
		fmt.Println("Listening------------")
		fmt.Println(web3.Net.Listening(context.Background()))
		fmt.Println("PeerCount------------")
		fmt.Println(web3.Net.PeerCount(context.Background()))
	}
	if false {
		fmt.Println("ListWallets------------")
		fmt.Println(web3.Personal.ListWallets(context.Background()))
	}
	if false {
		fmt.Println("Modules------------")
		fmt.Println(web3.Rpc.Modules(context.Background()))
	}
	if true {
		bin := `// SPDX-License-Identifier: GPL-3.0

		pragma solidity >=0.8.2 <0.9.0;
		
		/**
		 * @title Storage
		 * @dev Store & retrieve value in a variable
		 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
		 */
		contract Storage {
		
			uint256 number;
		
			/**
			 * @dev Store value in variable
			 * @param num value to store
			 */
			function store(uint256 num) public {
				number = num;
			}
		
			/**
			 * @dev Return value 
			 * @return value of 'number'
			 */
			function retrieve() public view returns (uint256){
				return number;
			}
		}
		`
		_ = bin
		fmt.Println("Test------------")
		fmt.Println(web3.TxPool.Test(context.Background()))
		return
		fmt.Println("Content------------")
		fmt.Println(web3.TxPool.Content(context.Background()))
		fmt.Println("Inspect------------")
		fmt.Println(web3.TxPool.Inspect(context.Background()))
		fmt.Println("Status------------")
		fmt.Println(web3.TxPool.Status(context.Background()))
	}
	if false {
		fmt.Println("Proposals------------")
		fmt.Println(web3.Clique.Proposals(context.Background()))
		fmt.Println("Status------------")
		fmt.Println(web3.Clique.Status(context.Background()))
	}
}
