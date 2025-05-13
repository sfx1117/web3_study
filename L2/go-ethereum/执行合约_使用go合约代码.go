package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-study/store"
	"log"
)

/*
	执行合约--使用go合约代码
*/
const contraceAddress = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"

func main() {
	//client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contraceAddress)
	//合约实例
	storeContract, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//privatekey
	privateKey, err := crypto.GenerateKey()
	//chainId
	chainID, err := client.ChainID(context.Background())
	//*bind.TransactOpts
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	//key
	var key [32]byte
	//value
	var value [32]byte
	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))
	//调用合约方法
	tx, err := storeContract.SetItem(opts, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	//根据key  获取value
	callOpts := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpts, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valueInContract == value)
}
