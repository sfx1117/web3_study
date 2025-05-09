package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
	ETH转账
*/

func main() {
	//1、创建client
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	//私钥
	//privatekey, err := crypto.GenerateKey()
	privatekey, err := crypto.HexToECDSA("")
	//公钥
	publickey := privatekey.Public()
	publicKeyEcdsa := publickey.(*ecdsa.PublicKey)
	//发送方地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyEcdsa)
	//nonce  交易序列号
	nonce, err2 := client.PendingNonceAt(context.Background(), fromAddress)
	if err2 != nil {
		log.Fatal(err2)
	}
	//value
	value := big.NewInt(1000000000000000000)
	//gaslimit
	gaslimit := uint64(2100)
	//gasprice
	gasprice, err3 := client.SuggestGasPrice(context.Background())
	if err3 != nil {
		log.Fatal(err3)
	}
	//toaddress
	toAdress := common.HexToAddress("")
	//data
	var data []byte
	//生成交易 tx
	//tx := types.NewTransaction(nonce, toAdress, value, gaslimit, gasprice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAdress,
		Gas:      gaslimit,
		GasPrice: gasprice,
		Data:     data,
		Value:    value,
	})
	//chainId
	chainId, err4 := client.ChainID(context.Background())
	if err4 != nil {
		log.Fatal(err4)
	}
	//生成签名
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privatekey)
	//发送交易
	err5 := client.SendTransaction(context.Background(), signTx)
	if err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println(signTx.Hash().Hex())
}
