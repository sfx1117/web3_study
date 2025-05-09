package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/sfx/go_test/erc20"
	"log"
	"math"
	"math/big"
)

/*
	查询代币余额
*/

func main() {
	//client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	//tokenaddress
	tokenAddress := common.HexToAddress("")
	//instance
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	//用户钱包地址
	address := common.HexToAddress("")
	//余额
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	//name 代币名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	//symbol 代币简称
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	//decimals 代币单位
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	fmt.Println(name)
	fmt.Println(symbol)
	fmt.Println(decimals)

	//将余额转换为eth
	fbal := new(big.Float)
	fbal.SetString(balance.String())
	ethval := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Println(ethval)
}
