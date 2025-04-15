// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//全局变量
//区块信息  block
//交易信息  tx
//消息信息  msg
contract GlobalVar {
  
    function test() public {
    // block.number;
    // block.timestamp;

    // tx.origin;
    // tx.gas;

    // msg.data;
    // msg.sender;
    // msg.gas;
    // msg.value;

    }
    //获取区块信息  
    function getBlockDetail() public view returns (uint256,uint256){
        return(
            block.number,
            block.timestamp
        );
    }

}


contract Caller{
    //0x5B38Da6a701c568545dCfcB03FcB875f56beddC4   钱包地址
    function testMsg() public view returns(address){
        return msg.sender;
    }
    //0x5B38Da6a701c568545dCfcB03FcB875f56beddC4   钱包地址
    function testTx() public view returns(address){
        return tx.origin;
    }
}

contract Callee{
    Caller caller;
    constructor(){
        caller=new Caller();
    }
    //0xDA07165D4f7c84EEEfa7a4Ff439e039B7925d3dF  Callee此合约的地址
    function testMsg() public view returns(address){
        return caller.testMsg();
    }
    //0x5B38Da6a701c568545dCfcB03FcB875f56beddC4   钱包地址
    function testTx() public view returns(address){
        return caller.testTx();
    }

}