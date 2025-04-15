// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract arrayContract{
    uint256 public a=10;

    /* 动态数组*/
    uint256[] public arr1;//长度为0

    uint256[] public arr2=[1,2,3];//长度为3 并初始化

    uint256[] public arr3=new uint256[](5); //new关键字 长度为5 全部为0

    function readArray1() public view returns (uint256[] memory data,uint256 len){
        data=arr1;
        len=data.length;
    }
    function readArray2() public view returns (uint256[] memory data,uint256 len){
        data=arr2;
        len=data.length;
    }
    function readArray3() public view returns (uint256[] memory data,uint256 len){
        data=arr3;
        len=data.length;
    }

    /*静态数组*/
    uint256[3] public arr4;//长度为3 全部都为0
    uint256[3] public arr5=[1,2,3];//长度为3 并初始化
    function readStaticArray1() public view returns (uint256[3] memory data,uint256 len){
        data=arr4;
        len=data.length;
    }
    function readStaticArray2() public view returns (uint256[3] memory data,uint256 len){
        data=arr5;
        len=data.length;
    }

    /*特殊数组 bytes和string*/
    bytes public bs="abc\x22\x22";
    bytes public bs1=new bytes(10); //创建一个长度为10 的字节数组
    //bytes可以通过下标进行读取
    function readBytes(uint i) public view returns(bytes1){
        return bs[i];
    }
    //bytes 切片读取
    function readBytes1(bytes calldata data,uint start,uint end) public pure returns(bytes memory){
        return data[start:end];
    }

    //string 不可以通过下标进行读取
    string public str1;
    string public str2="abc\x22\x22";






    
}
