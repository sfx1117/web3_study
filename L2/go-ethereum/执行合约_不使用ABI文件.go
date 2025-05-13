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
	"log"
	"math/big"
	"time"
)

/*
	执行合约--不使用ABI文件
*/
const contractAddress = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"
const contractJson = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

func main() {
	//client
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
	publicKey := privateKey.Public()
	publickeyEcdsa := publicKey.(*ecdsa.PublicKey)
	//fromAddress
	fromAddress := crypto.PubkeyToAddress(*publickeyEcdsa)

	//chainId
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//toAddress
	toAddress := common.HexToAddress(contractAddress)
	//gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	//gaslimit
	gaslimit := uint64(2100)
	//data
	setItem := []byte("setItem(bytes32,bytes32)")
	setItemId := crypto.Keccak256(setItem)[:4]
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	var data []byte
	data = append(data, setItemId...)
	data = append(data, key[:]...)
	data = append(data, value[:]...)

	//transaction
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gaslimit, gasPrice, data)

	//signTx
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//client.sendTransaction
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signTx.Hash().Hex())
	//等待执行成功，查询收据
	waitForReceipt(client, signTx.Hash())

	//根据key查询value
	items := []byte("items(bytes32)")
	itemsId := crypto.Keccak256(items)[:4]

	var callInput []byte
	callInput = append(callInput, itemsId...)
	callInput = append(callInput, key[:]...)
	callMsg := ethereum.CallMsg{
		To:   &toAddress,
		Data: callInput,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func waitForReceipt(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		time.Sleep(time.Second)
	}
}
