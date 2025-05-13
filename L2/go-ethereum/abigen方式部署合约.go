package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-study/store"
	"log"
	"math/big"
)

/*
	abigen方式部署合约
*/

func main() {
	//创建client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	//privatekey
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//publickey
	publickey := privateKey.Public()
	publicKeyEcdsa := publickey.(*ecdsa.PublicKey)
	//address
	address := crypto.PubkeyToAddress(*publicKeyEcdsa)

	//nonce
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	//chainId
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//gaslimit
	gasLimit := uint64(300000)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasLimit = gasLimit
	transactOpts.GasPrice = gasPrice
	transactOpts.Value = big.NewInt(0)
	vserion := "1.0"
	deployAddress, tx, instance, err := store.DeployStore(transactOpts, client, vserion)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(deployAddress.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
