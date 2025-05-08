package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

/*
	创建钱包
*/

func main() {
	//创建新钱包私钥
	//privateKey, err := crypto.GenerateKey()
	//根据私钥string 得到私钥对象
	privateKey, err := crypto.HexToECDSA("85690a1214049cf2565b6bbc173840ac683a4ee7d316ffe2a4a3f320a216da17")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)

	//将私钥转换为[]byte
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(privateKeyBytes)
	//转换为16进制字符串
	privateKeyString := hexutil.Encode(privateKeyBytes)
	fmt.Println(privateKeyString) //0x85690a1214049cf2565b6bbc173840ac683a4ee7d316ffe2a4a3f320a216da17

	//去掉前两位0x
	privateKeyStringSub := privateKeyString[2:]
	fmt.Println(privateKeyStringSub) //85690a1214049cf2565b6bbc173840ac683a4ee7d316ffe2a4a3f320a216da17

	fmt.Println("=========================================")

	//根据私钥获取公钥
	publicKey := privateKey.Public()
	fmt.Println(publicKey)
	//将公钥转换为[]byte
	publicKeyEcdsa := publicKey.(*ecdsa.PublicKey)
	publickeyBytes := crypto.FromECDSAPub(publicKeyEcdsa)
	fmt.Println(publickeyBytes)
	//将公钥转换为string
	publickeyString := hexutil.Encode(publickeyBytes)
	fmt.Println(publickeyString)
	//去掉前4位0x04
	publickeyStringSub := publickeyString[4:]
	fmt.Println(publickeyStringSub)
	//生成公共地址
	address := crypto.PubkeyToAddress(*publicKeyEcdsa)
	fmt.Println(address) //0x87e22B36cC1cc94c83ce8744a9a2701cBe3F8615

	fmt.Println("=========================================")

	//验证address生成方法
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publickeyBytes[1:])
	str := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println(str) //0x87e22B36cC1cc94c83ce8744a9a2701cBe3F8615

}
