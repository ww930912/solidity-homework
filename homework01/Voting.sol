// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 简单投票合约
contract Voting {
    uint256 public round;
    // 轮次 => 候选人 => 得票数
    mapping(uint256 => mapping(string => uint256)) private _votes;

    /// @notice 给候选人投一票
    function vote(string calldata candidate) external {
        unchecked {
            ++_votes[round][candidate];
        }
    }

    /// @notice 查询候选人在当前轮次的得票数
    function getVotes(
        string calldata candidate
    ) external view returns (uint256) {
        return _votes[round][candidate];
    }

    /// @notice 重置所有候选人的得票数：只改一个存储槽
    function resetVotes() external {
        unchecked {
            ++round;
        }
    }
}
