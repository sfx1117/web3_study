// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
/*  - 所有人都可以存钱-----------------------写个receive方法
        - ETH
    - 只有合约 owner 才可以取钱 -------------写个取钱方法
    - 只要取钱，合约就销毁掉 selfdestruct
    - 扩展：支持主币以外的资产
    - ERC20
    - ERC721
*/
contract Bank{
    //声明 账户地址
    address public immutable owner;
    //构造函数
    constructor(){
        owner=payable(msg.sender);
    }
    //存钱事件
    event SaveMoney(address addr,uint256 amount);
    //取钱事件
    event WithdrawMoney(uint256 amount);

    //receive方法 存钱
    receive() external payable{
        emit SaveMoney(msg.sender, msg.value);
    }

    //取钱方法
    function withdrawMoney() external {
        require(msg.sender==owner,"Not owner");
        emit WithdrawMoney(address(this).balance);
        selfdestruct(payable(msg.sender));
    }
    //获取账户余额
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }

}