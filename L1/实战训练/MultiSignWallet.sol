// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*多签钱包的功能: 合约有多个 owner，一笔交易发出后，需要多个 owner 确认，确认数达到最低要求数之后，才可以真正的执行。
### 1.原理
- 部署时候传入地址参数和需要的签名数
  - 多个 owner 地址
  - 发起交易的最低签名数
- 有接受 ETH 主币的方法，
- 除了存款外，其他所有方法都需要 owner 地址才可以触发
- 发送前需要检测是否获得了足够的签名数
- 使用发出的交易数量值作为签名的凭据 ID（类似上么）
- 每次修改状态变量都需要抛出事件
- 允许批准的交易，在没有真正执行前取消。
- 足够数量的 approve 后，才允许真正执行。*/
contract MultiSignWallet{
    //多个owner
    address[] public owners;
    uint256 public immutable required;
    mapping (address => bool) public isOwner;//判断owners是否重复
    //交易结构体
    struct Transaction{
        address toAds;//收款地址
        uint256 amount;//金额
        bytes data;//交易数据（备注）
        bool excuted;//执行状态
    }
    Transaction[] public transactions;
    mapping (uint256 => mapping(address => bool)) public approved;

    //事件
    event Deposit(address indexed  ads,uint256 amount);
    event Submit(uint256 indexed _txId);
    event Approve(uint256 indexed _txId,address indexed _owner);
    event Revoke(uint256 indexed _txId,address indexed _owner);
    event Excute(uint256 indexed _txId);

    //函数修饰器
    modifier onlyOwner(){
        require(isOwner[msg.sender],"the sender is not a owner");
        _;
    }
    modifier txIsExists(uint256 _txId){
        require(_txId<transactions.length,"_txId is not exists");
        _;
    }
    modifier notApproved(uint256 _txId){
        require(!approved[_txId][msg.sender],"the transaction is approved");
        _;
    }
    modifier notExcuted(uint256 _txId){
        require(!transactions[_txId].excuted,"the transaction is excuted");
        _;
    }

    //构造函数
    constructor(address[] memory _owner,uint256 _required){
        require(_owner.length>0,"owner required");
        require(_required>0 && _required<owners.length,"invalid required number of owners");
        for(uint256 i=0;i<_owner.length;i++){
            address owner=_owner[i];
            require(!isOwner[owner],"the owner is exits");
            isOwner[owner]=true;
            owners.push(owner);
        }
        required=_required;
    }
    //存款
    receive () external payable{
        emit Deposit(msg.sender, msg.value);
    }
    //获取余额
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    //提交交易,并将交易数组下标作为id
    function submit(address _toAds,uint256 _amount,bytes memory data)external onlyOwner returns(uint256){
        transactions.push(Transaction(_toAds,_amount,data,false));
        emit Submit(transactions.length-1);
        return transactions.length-1;

    }
    //审批交易,不同的owner来调用此方法进行审批
    function approve(uint256 _txId) external onlyOwner txIsExists(_txId) notApproved(_txId) notExcuted(_txId){
        approved[_txId][msg.sender]=true;
        emit Approve(_txId,msg.sender);
    }

    //执行交易
    function excute(uint256 _txId) external onlyOwner txIsExists(_txId) notExcuted(_txId){
        //先判断审批人数是否满足
        require(getApprovedNum(_txId)>=required,"approvals < required");
        //获取交易
        Transaction storage tran=transactions[_txId];
        tran.excuted=true;
        (bool success,)=tran.toAds.call{value:tran.amount}(
            tran.data
        );
        require(success,"tx is faild");
        emit Excute(_txId);
    }
    //获取审批人数
    function getApprovedNum(uint256 _txId) public view returns(uint256 count){
        for(uint256 i=0;i<owners.length;i++){
            if(approved[_txId][owners[i]]){
                count+=1;
            }
        }
    } 

    //取消审批
    function revoke(uint256 _txId) external onlyOwner txIsExists(_txId) notExcuted(_txId){ 
        //先判断交易被审批，才能取消审批 
        require(approved[_txId][msg.sender],"tx not approve");
        //更新状态
        approved[_txId][msg.sender]=false;
        emit Approve(_txId, msg.sender);
    }



}