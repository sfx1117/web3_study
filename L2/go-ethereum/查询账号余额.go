package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

/*
	查询账号余额
*/

func main() {
	//创建client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	//账户地址
	accAdr := common.HexToAddress("")
	//查询账号余额
	balance, err := client.BalanceAt(context.Background(), accAdr, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	//查询指定区块余额
	blockNum := big.NewInt(54722)
	balanceAt, err := client.BalanceAt(context.Background(), accAdr, blockNum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)

	//将余额转换为eth
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethVal := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethVal)

	//查询待处理的余额
	pendBalance, err := client.PendingBalanceAt(context.Background(), accAdr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pendBalance)

}
