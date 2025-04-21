// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
    众筹合约是一个募集资金的合约，在区块链上，我们是募集以太币，类似互联网业务的水滴筹。
    区块链早起的 ICO 就是类似业务。
    1.需求分析
        众筹合约分为两种角色：一个是受益人，一个是资助者。
     两种角色:
        受益人   beneficiary => address         => address 类型
        资助者   funders     => address:amount  => mapping 类型 或者 struct 类型

    状态变量按照众筹的业务：
    状态变量
         筹资目标数量    fundingGoal
         当前募集数量    fundingAmount
         资助者列表      funders
         资助者人数      fundersKey

    需要部署时候传入的数据:
          受益人
          筹资目标数量
*/
contract CrowdFunding{
    //声明两种角色
    address public immutable beneficiary;//受益人
    mapping(address => uint256) public funders;//资助人
    //声明状态变量
    uint256 public immutable fundingGoal;//筹资目标
    uint256 public fundingAmount;//筹集的当前金额
    mapping(address => bool) public fundersInserted;//资助人
    address[] public fundersKey;//用list.length表示资助者人数
    bool public AVAILABLED=true;//是否完成筹集的开关

    //构造函数
    constructor(address beneficiary_ , uint256 fundingGoal_){
        beneficiary=beneficiary_;
        fundingGoal=fundingGoal_;
    }

    //筹资函数
    function contribute() external payable {
        //先判断 众筹是否已关闭
        require(AVAILABLED,"CrowdFunding is closed");
        //预估筹集金额=当前资助金额+已筹集的金额
        uint256 preFundingAmount= fundingAmount+msg.value;
        //退还金额   当筹集金额>筹集目标  多余部分退还
        uint256 refundAmount=0;
        if(preFundingAmount>fundingGoal){
            refundAmount=preFundingAmount-fundingGoal;
        }
        //添加到资助人map中
        funders[msg.sender]+=msg.value-refundAmount;
        //更新已筹集金额
        fundingAmount+=msg.value-refundAmount;
        //第一次资助，则加入fundersKey
        if(!fundersInserted[msg.sender]){
            fundersInserted[msg.sender]=true;
            fundersKey.push(msg.sender);
        }
        //判断是否满足筹资金额，满足则 退还金额，并关闭众筹
        if(refundAmount>0){
            payable(beneficiary).transfer(refundAmount);
            closed();
        }
    }
    //关闭函数
    function closed()public returns(bool){
        if(fundingAmount<fundingGoal){
            return false;
        }
        //修改状态
        uint256 amount=fundingAmount;
        fundingAmount=0;
        AVAILABLED=false;
        //将筹集金额发送给受益人
        payable (beneficiary).transfer(amount);
        return true;
    }

    //获取资助人数
    function getFundersLength()public view returns(uint256){
        return fundersKey.length;
    }
}

