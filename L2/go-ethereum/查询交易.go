package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
	查询交易
	1、block.Transactions() 获取区块的交易信息
	2、client.ChainId(context)  获取链id
	3、types.Sender(types.NewEIP155Signer(ChainId),tx) 获取发送者信息
	4、client.TransactionReceipt(context,tx.Hash()) 获取交易的收据信息
*/

func main() {
	//创建client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	//1、根据区块高度 获取区块信息
	blockNumber := big.NewInt(543342)
	block, err2 := client.BlockByNumber(context.Background(), blockNumber)
	if err2 != nil {
		log.Fatal(err2)
	}
	//获取链id
	chainID, err3 := client.ChainID(context.Background())
	if err3 != nil {
		log.Fatal(err3)
	}
	//根据区块获取交易信息
	transactions := block.Transactions()
	//遍历交易信息
	for _, tx := range transactions {
		fmt.Println(tx.Hash().Hex(), tx.To().Hex(), tx.Time(), tx.Size())
		fmt.Println(tx.Nonce(), tx.GasPrice(), tx.Gas(), tx.Value(), tx.Data(), tx.To().Hex())
		//获取发送者信息
		sender, err4 := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err4 == nil {
			fmt.Println(sender.Hex())
		}
		//获取收据信息
		receipt, err4 := client.TransactionReceipt(context.Background(), tx.Hash())
		if err4 == nil {
			fmt.Println(receipt.Status, receipt.Logs)
		}
	}

	//2、根据区块hash获取交易信息
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err5 := client.TransactionCount(context.Background(), blockHash)
	if err5 != nil {
		log.Fatal(err5)
	}
	for i := uint(0); i < count; i++ {
		txx, err6 := client.TransactionInBlock(context.Background(), blockHash, i)
		if err6 == nil {
			fmt.Println(txx)
		}
	}

	//3、根据交易hash获取交易
	txHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	txxx, ispending, err7 := client.TransactionByHash(context.Background(), txHash)
	if err7 == nil {
		fmt.Println(txxx)
		fmt.Println(ispending)
	}
}
