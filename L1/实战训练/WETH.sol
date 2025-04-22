// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
/*
    WETH 是包装 ETH 主币，作为 ERC20 的合约。
    标准的 ERC20 合约包括如下几个
    - 3 个查询
        - balanceOf: 查询指定地址的 Token 数量
        - allowance: 查询指定地址对另外一个地址的剩余授权额度
        - totalSupply: 查询当前合约的 Token 总量
    - 2 个交易
        - transfer: 从当前调用者地址发送指定数量的 Token 到指定地址。
            - 这是一个写入方法，所以还会抛出一个 Transfer 事件。
        - transferFrom: 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中。
    - 2 个事件
        - Transfer
        - Approval
    - 1 个授权
        - approve: 授权指定地址可以操作调用者的最大 Token 数量。
*/
contract WETH{
    string public name="Wrapped Ether";
    string public symbol="Weth";
    uint8 public decimals=18;

    //声明两个map
    mapping (address => uint256) public balanceOf;
    mapping (address => mapping(address => uint256)) public allwance;


    //事件===打印日志
    event Dopesit(address indexed src,uint256 amount);
    event Withdraw(address indexed src, uint256 amount);
    event Approve(address indexed src,address indexed toAds,uint256 amount);
    event Transfer (address indexed src, address indexed toAds, uint256 amount);


    //存钱
    function deposit() public payable{
        balanceOf[msg.sender]+=msg.value;
        emit Dopesit(msg.sender, msg.value);
    }

    //取钱
    function withdraw(uint256 amount)public {
        require(balanceOf[msg.sender] >= amount,"the balance is not enough");
        balanceOf[msg.sender]-=amount;
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender,amount);
    }

    //查询当前合约的 Token 总量totalSupply
    function totalSupply() public view returns (uint256){
        return address(this).balance;
    }

    //给另一个地址授权Token
    function approve(address toAds,uint256 amount) public returns(bool){
        allwance[msg.sender][toAds]=amount;
        emit Approve(msg.sender,toAds,amount);
        return true;
    }

    // 从当前调用者地址发送指定数量的 Token 到指定地址。
    function transfer(address toAds,uint256 amount) public payable returns(bool){
       return transferFrom(msg.sender, toAds, amount);
    }

    //transferFrom: 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中。
    function transferFrom(address fromAds, address toAds, uint256 amount) public returns(bool){
        require(balanceOf[fromAds]>=amount,"the balance is not enough");
        //如果msg.sender != fromAds，则代表是从授权账户转账
        if(msg.sender != fromAds){
            require(allwance[fromAds][msg.sender]>=amount);
            allwance[fromAds][msg.sender]-=amount;
        }
        balanceOf[fromAds]-=amount;
        balanceOf[toAds]+=amount;
        emit Transfer(fromAds,toAds,amount);
        return true;
    }

    receive() external payable { 
        deposit();
    }
    fallback() external payable { 
        deposit();
    }

}



