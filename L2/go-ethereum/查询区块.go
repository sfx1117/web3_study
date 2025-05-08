package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
	查询区块
	1、client.HeaderByNumber()获取区块头信息
	2、client.BlockByNumber()获取完整区块信息
*/

func main() {
	//创建客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGk0ZDMbZ1AkVvIRu9N8PyuX0Z")
	if err != nil {
		log.Fatal(err)
	}
	//区块高度
	blockNum := big.NewInt(5671744)
	//1、client.HeaderByNumber()获取区块头信息
	header, err1 := client.HeaderByNumber(context.Background(), blockNum)
	fmt.Println(header)
	if err1 != nil {
		log.Fatal(err1)
	}
	number := header.Number         //区块高度
	difficulty := header.Difficulty //区块难度
	hash := header.Hash().Hex()     //区块hash
	time := header.Time             //区块时间
	fmt.Println(number, difficulty, hash, time)

	//2、通过client.BlockByNumber 获取区块信息
	block, err2 := client.BlockByNumber(context.Background(), blockNum)
	if err2 != nil {
		log.Fatal(err2)
	}

	num := block.Number()
	hex := block.Hash().Hex()
	diff := block.Difficulty()
	tt := block.Time()
	transactions := block.Transactions() //区块包含的交易列表信息
	//transaction := block.Transaction()//入参应该是交易hash，单独查询某条交易信息使用
	fmt.Println(num, hex, diff, tt, transactions)
	len := len(transactions) //区块交易总数量
	fmt.Println(len)

	//3、client.TransactionCount 获取某一个区块交易个数
	count, err3 := client.TransactionCount(context.Background(), block.Hash())
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(count)

}
