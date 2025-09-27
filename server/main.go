package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/holiman/uint256"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/19f8fd6dddaf4a52b9252e4bcebd24c8")
	rpc, err := rpc.Dial("wss://mainnet.infura.io/ws/v3/19f8fd6dddaf4a52b9252e4bcebd24c8")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("Connected to Ethereum client")

	header, err := BlockHeaderByNumber(client, uint256.NewInt(12345678))
	if err != nil {
		log.Fatalf("Failed to retrieve block header: %v", err)
	}
	txChan := make(chan string)
	StartPendingTxSub(rpc, txChan)
	go ProcessTxHashes(client, txChan)
	fmt.Printf("Block Header: %+v\n", header)
	defer client.Close()
}
