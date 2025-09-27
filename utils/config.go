package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	rpcClient, err := rpc.Dial("wss://mainnet.infura.io/ws/v3/19f8fd6dddaf4a52b9252e4bcebd24c8")
	if err != nil {
		log.Fatalf("Failed to connect to Infura WebSocket: %v", err)
	}
	defer rpcClient.Close()
	fmt.Println("✅ Connected to Infura WebSocket")

	// Channel to receive tx hashes
	txHashes := make(chan string)

	// Subscribe to pending transactions
	sub, err := rpcClient.EthSubscribe(context.Background(), txHashes, "newPendingTransactions")
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("📡 Subscribed to pending transactions...")

	// Listen for results
	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription error:", err)
		case txHash := <-txHashes:
			fmt.Println("📝 Pending TX:", txHash)
			// if you want to fetch tx details, use eth_getTransactionByHash
			// or wrap ethclient.NewClient(rpcClient) and call client.TransactionByHash
		}
	}
}
