package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
)

func BlockHeaderByNumber(client *ethclient.Client, num *uint256.Int) (*types.Header, error) {
	header, err := client.HeaderByNumber(context.Background(), num.ToBig())
	fmt.Println(header)
	return header, err
}
