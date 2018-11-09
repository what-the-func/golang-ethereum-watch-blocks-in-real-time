package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Block Hash: %v\n", block.Hash().Hex())
			fmt.Printf("Block #: %v\n", block.Number().Uint64())
			fmt.Printf("Block Time: %v\n", block.Time().Uint64())
			fmt.Printf("Nonce: %v\n", block.Nonce())
			fmt.Printf("Tx Count: %v\n", len(block.Transactions()))
		}
	}
}
