// SPDX-License-Identifier: MIT

pragma solidity ^0.8.31;


/**
✅ 创建一个名为Voting的合约，包含以下功能：
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
**/
contract Voting_1 {

    mapping(address => uint) voteMap;
    //记录投递，防止重复投
    mapping (address => mapping (address=>bool)) onlyOneVoteMap;
    //记录vote
    address[]  voteUserArray;

    function getOnlyOneVoteMap(address _address) public view returns (bool isVote) {
        isVote = onlyOneVoteMap[msg.sender][_address];
    }

    //投票，不能重复投
    function vote(address _address) public{
        require(!onlyOneVoteMap[msg.sender][_address], "already vote!!!!!!!!!!");
        onlyOneVoteMap[msg.sender][_address] = true;
        voteMap[_address] += 1;
        if (voteMap[_address] == 1) {
            voteUserArray.push(_address);
        }
    }

    //返回候选人得票数
    function getVotes(address _address) public view returns (uint _voteNum) {
        _voteNum = voteMap[_address];
    }

    // 重置所有候选人的得票数
    function resetVotes() external {
        for(uint i = 0; i < voteUserArray.length; i++) {
            voteMap[voteUserArray[i]] = 0;
        }
    }

//用户
//0x5B38Da6a701c568545dCfcB03FcB875f56beddC4
//0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2
//0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db

//候选人
//0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB
//0x17F6AD8Ef982297579C203069C1DbfFE4348c372

}