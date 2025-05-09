package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

/*
	订阅区块
*/

func main() {
	//client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	//headers channel
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//监听headers channel
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
			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number())
			fmt.Println(block.Nonce())
			fmt.Println(block.Time())
			fmt.Println(block.Transactions())
		}
	}
}
