package main

import (
	"fmt"
	"log"

	"github.com/ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func getPendingTxns(client *ethclient.Client) {
	rpc, err := rpc.Dial("wss://mainnet.infura.io/ws/v3/19f8fd6dddaf4a52b9252e4bcebd24c8")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	client = ethclient.NewClient(rpc)

	txns := make(chan *types.Transaction)

	sub, err := client.SubscribeFullPendingTransactions(context.Background(), txs)

	if err != nil {
		log.Fatalf("Failed to subscribe to pending transactions: %v", err)
	}
	fmt.Println("Subscribed to pending transactions")
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case tx := <-txns:
			fmt.Printf("Pending Transaction: %+v\n", tx)
		}
	}

}
