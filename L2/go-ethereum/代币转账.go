package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

/*
	代币转账
*/

func main() {
	//创建client
	client, err1 := ethclient.Dial("")
	if err1 != nil {
		log.Fatal(err1)
	}
	//私钥
	privatekey, err2 := crypto.GenerateKey()
	if err2 != nil {
		log.Fatal(err2)
	}
	//公钥
	publickey := privatekey.Public()
	publickeyEcdsa := publickey.(*ecdsa.PublicKey)
	//fromaddress
	fromAddress := crypto.PubkeyToAddress(*publickeyEcdsa)
	//nonce
	nonce, err3 := client.PendingNonceAt(context.Background(), fromAddress)
	if err3 != nil {
		log.Fatal(err3)
	}
	//toaddress
	toAddress := common.HexToAddress("")    //真实接收方地址
	tokenAddress := common.HexToAddress("") //合约地址
	//value
	value := big.NewInt(0)
	//gasprice
	gasprice, err4 := client.SuggestGasPrice(context.Background())
	if err4 != nil {
		log.Fatal(err4)
	}

	/**
	data 包含3个部分，methodId, toaddress, amount
	*/
	//methodId
	var transferMethod = []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferMethod)
	methodId := hash.Sum(nil)[:4]

	//toaddress  左边补齐32位
	padAdress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//amount 实际转账金额 左边补齐32
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	padAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodId...)
	data = append(data, padAdress...)
	data = append(data, padAmount...)
	//gaslimit  需要根据data来定义
	gaslimit, err5 := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err5 != nil {
		log.Fatal(err5)
	}
	//tx
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    value,
		Gas:      gaslimit,
		GasPrice: gasprice,
		Data:     data,
	})
	//ChainId
	chainID, err6 := client.ChainID(context.Background())
	if err6 != nil {
		log.Fatal(err6)
	}
	//signtx
	signTx, err7 := types.SignTx(tx, types.NewEIP155Signer(chainID), privatekey)
	if err7 != nil {
		log.Fatal(err7)
	}
	//sendtx
	err8 := client.SendTransaction(context.Background(), signTx)
	if err8 != nil {
		log.Fatal(err8)
	}
	fmt.Println(signTx.Hash().Hex())
}
