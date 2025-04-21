// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
    这一个实战主要是加深大家对 3 个取钱方法的使用。
    - 任何人都可以发送金额到合约
    - 只有 owner 可以取款
    - 3 种取钱方式:
    transfer()
    send()
    call()
*/

contract EtherWallet{

    address payable public  immutable owner;
    constructor(){
        owner=payable(msg.sender);
    }

    event Log(string funName,address ads,uint256 value);
    //收款
    receive() external payable { 
        emit Log("receive",msg.sender,msg.value);
    }
    //取钱
    function witherdraw()external {
        require(owner==msg.sender,"Not owner");
        //owner.transfer比 payable(msg.sender).transfer更消耗gas
        //owner.transfer(address(this).balance);
        payable(msg.sender).transfer(address(this).balance);
    }
    function witherdraw2() external {
        require(owner==msg.sender,"Not owner");
        bool sucess=payable(msg.sender).send(address(this).balance);
        require(sucess,"send faild");
    }
    function  witherdraw3() external{
        require(owner==msg.sender,"Not owner");
        (bool sucess,)=msg.sender.call{value:address(this).balance}("");
        require(sucess,"Call faild");
    }

    //查询账号余额
    function getBalance() public view returns (uint256){
        return address(this).balance;
    }    

}
