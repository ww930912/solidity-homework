// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 测试用的 Chainlink 喂价合约，只实现拍卖合约用到的两个函数
contract MockV3Aggregator {
    uint8 public decimals;
    int256 public answer;

    constructor(uint8 _decimals, int256 _answer) {
        decimals = _decimals;
        answer = _answer;
    }

    function latestRoundData() external view returns (uint80, int256, uint256, uint256, uint80) {
        return (0, answer, 0, block.timestamp, 0);
    }
}
