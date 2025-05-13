package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

/*
	合约事件-订阅事件
*/
const contractAddress = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"
const contractJson = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func main() {
	//client 订阅需要用到websocket RPC URL
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	//address
	address := common.HexToAddress(contractAddress)
	//fiterQuery
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			address,
		},
	}
	//channel
	logsCh := make(chan types.Log)
	//contractABI
	contractABI, err := abi.JSON(strings.NewReader(contractJson))
	if err != nil {
		log.Fatal(err)
	}
	//logs
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsCh)
	if err != nil {
		log.Fatal(err)
	}
	//监听channel
	for {
		select {
		case subErr := <-sub.Err():
			fmt.Println(subErr)
		case vlog := <-logsCh:
			fmt.Println(vlog.BlockNumber)
			fmt.Println(vlog.BlockHash.Hex())
			fmt.Println(vlog.TxHash.Hex())
			//vlog.data
			event := struct {
				key   [32]byte
				value [32]byte
			}{}
			err = contractABI.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(common.Bytes2Hex(event.key[:]))
			fmt.Println(common.Bytes2Hex(event.value[:]))

			//vlog.topics
			var topics []string
			for i, _ := range vlog.Topics {
				topics = append(topics, vlog.Topics[i].Hex())
			}
			fmt.Println(topics[0])
			if len(topics) > 1 {
				fmt.Println(topics[1:])
			}
		}
	}
}
