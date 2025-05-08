package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

/*
	查询收据

*/

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy/n1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	//1、根据blockhash
	receiptsHash, err2 := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, receipt := range receiptsHash {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.ContractAddress.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.BlockHash)
	}
	//2、根据blockNumber
	receiptsNum, err3 := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err3 != nil {
		log.Fatal(err3)
	}
	for _, receipt := range receiptsNum {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.ContractAddress.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println(receipt.BlockHash)
	}
	//3、根据txHash
	txHash := common.HexToHash("")
	receiptTxHash, err := client.TransactionReceipt(context.Background(), txHash)
	fmt.Println(receiptTxHash.Status)
}
