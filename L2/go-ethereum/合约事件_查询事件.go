package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

/*
	合约事件-查询事件
*/
const contractAddress = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"
const contractJson = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

func main() {
	//client
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}
	//address
	address := common.HexToAddress(contractAddress)
	//filterQuery 过滤条件
	query := ethereum.FilterQuery{
		//BlockHash
		FromBlock: big.NewInt(6920583),
		Addresses: []common.Address{
			address,
		},
		//Topics: [][]common.Hash{},
	}
	//logs
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	//contractABI
	contractABI, err := abi.JSON(strings.NewReader(contractJson))
	if err != nil {
		log.Fatal(err)
	}
	//遍历logs
	for _, vlog := range logs {
		fmt.Println(vlog.BlockNumber)
		fmt.Println(vlog.BlockHash.Hex())
		fmt.Println(vlog.TxHash.Hex())
		event := struct {
			key   [32]byte
			value [32]byte
		}{}
		//解析vlog.Data
		err = contractABI.UnpackIntoInterface(&event, "ItemSet", vlog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(common.Bytes2Hex(event.key[:]))
		fmt.Println(common.Bytes2Hex(event.value[:]))

		//topics
		var topics []string
		for i, _ := range vlog.Topics {
			topics = append(topics, vlog.Topics[i].Hex())
		}
		//topics的第一个是事件函数名(参数列表)的签名
		fmt.Println(topics[0])
		eventSignature := []byte("ItemSet(bytes32,bytes32)")
		hash := crypto.Keccak256Hash(eventSignature)
		fmt.Println(hash) //hash==topics[0]
		//topics其余的值是 参数列表中被 indexed声明的参数
		if len(topics) > 1 {
			fmt.Println(topics[1:])
		}

	}
}
