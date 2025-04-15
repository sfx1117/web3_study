// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MappingExample{

    mapping(string account => uint amount) public myMapping;

    //mapping 赋值只能通过函数
    function setMyMapping(string memory key,uint value) public {
        myMapping[key]=value;
    }

    //多级mapping  
    mapping (string key => mapping (string key1=> uint amount)) public  mulitMapping;

    function setMulitMapping(string memory key,string memory key1,uint amount) public {
        mulitMapping[key][key1]=amount;
    }
    //mapping 不能作为返回值类型
    // 原因：EVM  没有办法 将mapping全量数据从以太链中拉下来  
    function getMulitMapping(string memory key,string memory key1) public view returns (uint amount){
        return mulitMapping[key][key1];
    }
    //mapping 不能整体被删除，只能根据key删除value
    function delMulitMapping(string memory key,string memory key1) public{
        delete mulitMapping[key][key1];
    }



}