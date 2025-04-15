// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
//应用二进制接口
contract ABI{
    //1.编码
    // 入参：sfx,17
    function encode(string memory text,uint256 num)public pure returns(bytes memory,bytes memory){
        return(
            abi.encode(text,num),//正常编码
            //0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001100000000000000000000000000000000000000000000000000000000000000037366780000000000000000000000000000000000000000000000000000000000
            abi.encodePacked(text,num)//编码压缩
            //0x7366780000000000000000000000000000000000000000000000000000000000000011
        );
    }
    //2.解码  编码压缩的无法正常解码，需要补全
    // 0x0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000001100000000000000000000000000000000000000000000000000000000000000037366780000000000000000000000000000000000000000000000000000000000
    function decode(bytes memory encodeData) public pure returns(string memory,uint256){
        return abi.decode(encodeData, (string,uint256));
    }

    //3.获取当前函数签名 msg.sig
    function getSign() public pure returns(bytes4){
        return msg.sig;
        //0xcf504e0a
    }

    //4.计算函数选择器  调用函数生成msg.sig
    //入参：getSign()  结果和函数3的结果一样
    function computeSign(string memory func) public pure returns(bytes4){
        return bytes4(keccak256(bytes(func)));
        //0xcf504e0a
    }

    //5.获取当前函数的数据 msg.data
    //入参：
    //0xd4343bd30000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc40000000000000000000000000000000000000000000000000000000000000012
    function getMsgData(address addr,uint8 amount) public pure returns(bytes memory){
        return msg.data;
    }

    //6.调用函数生成 msg.data
    //0xd4343bd30000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc40000000000000000000000000000000000000000000000000000000000000012
    function encodeFunCall() public pure returns(bytes memory){
        return abi.encodeWithSignature("getMsgData(address,uint8)", 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4,18);
    }


    //7.哈希算法
    //入参：sfx
    function hashFunction(string memory input) public pure returns(bytes32,bytes32,bytes32){
        bytes memory data=abi.encodePacked(input);
        return(
            keccak256(data),//0xe4c2f43801fbf3ec42e65bef409709fd8d4b1b98d39b93b216bcf93c44404eae
            sha256(data),//0x001d5b73f5a97ade16cfd9bbbcf1a7b6e37ad954cbef0f47aa44da7c55e63c9e
            ripemd160(data)//0xcf142f8067f6811edd6d1a9987baeee657b9e3e8000000000000000000000000
        );
    }

    //8.数学运算
    // addmod   加法取模 计算 (x + y) % k，但使用数学上的安全计算方式，确保中间结果 (x + y) 即使溢出也能正确计算
    // mulmod   乘法取模 计算 (x * y) % k，同样避免中间结果 (x * y) 溢出导致的问题。
    function modualMath(uint256 x,uint256 y,uint256 z) public pure returns(uint256,uint256){
        return (
            addmod(x,y,z),
            mulmod(x,y,z)
        );
    }

    //9.椭圆曲线恢复公钥
    function recoverAddress(bytes32 hash,uint8 v,bytes32 r,bytes32 z) public pure returns(address){
        return ecrecover(hash,v,r,z);
    }

}