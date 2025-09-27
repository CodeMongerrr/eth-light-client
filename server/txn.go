package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ProcessTxHashes listens on the channel and fetches full txs
func ProcessTxHashes(client *ethclient.Client, txChan <-chan string) {
	for txHash := range txChan {
		fmt.Println("🔍 New pending tx hash:", txHash)
		tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Printf("❌ Failed to fetch tx %s: %v", txHash, err)
			continue
		}
		printTx(tx)
	}
}

func printTx(tx *types.Transaction) {
	fmt.Println("====================================")
	fmt.Println("Hash:", tx.Hash().Hex())
	fmt.Println("To:", tx.To())
	fmt.Println("Value:", tx.Value())
	fmt.Println("Gas:", tx.Gas())
	fmt.Println("GasPrice:", tx.GasPrice())
	fmt.Println("Data:", fmt.Sprintf("%x", tx.Data()[:8])) // first 4 bytes = method ID
	fmt.Println("====================================")
}
