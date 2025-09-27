package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

// StartPendingTxSub subscribes to pending txs and sends them to the channel
func StartPendingTxSub(rpcClient *rpc.Client, txChan chan string) {
	txHash := make(chan<- string)
	sub, err := rpcClient.EthSubscribe(context.Background(), txHash, "newPendingTransactions")
	fmt.Println(sub)
	if err != nil {
		fmt.Printf("Failed to subscribe: %v\n", err)
		return
	}
	fmt.Println("📡 Subscribed to pending transactions...")

	// Listen for subscription errors in the background
	go func() {
		for {
			select {
			case txChan <- txHash:
				fmt.Println("📝 Pending TX:", txHash)
			case err := <-sub.Err():
				fmt.Println("Subscription error:", err)
			}
		}
	}()
}
