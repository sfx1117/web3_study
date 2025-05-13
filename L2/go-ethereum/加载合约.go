package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-study/store"
	"log"
)

/*
	加载合约===生成合约实例
*/
const contraceAddress = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"

func main() {
	//client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contraceAddress)
	storeContract, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = storeContract
}
