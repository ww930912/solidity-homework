// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 讨饭合约：接收捐赠，所有者提款
contract BeggingContract {
    address public owner;
    // 捐赠者 => 累计捐赠金额（wei）
    mapping(address => uint256) public donations;

    // 额外挑战 1：捐赠事件
    event Donation(address indexed donor, uint256 amount);

    // 额外挑战 2：捐赠排行榜，按累计金额从高到低
    address[3] private topDonors;

    // 额外挑战 3：允许捐赠的时间段 [startTime, endTime]（区块时间戳，秒）
    uint256 public startTime;
    uint256 public endTime;

    modifier onlyOwner() {
        require(msg.sender == owner, "not owner");
        _;
    }

    /// @dev 默认从部署时刻起开放 7 天，owner 可通过 setDonationPeriod 修改
    constructor() {
        owner = msg.sender;
        startTime = block.timestamp;
        endTime = block.timestamp + 7 days;
    }

    /// @notice 捐赠以太币（仅限捐赠时间段内）
    function donate() public payable {
        require(msg.value > 0, "amount must > 0");
        require(
            block.timestamp >= startTime && block.timestamp <= endTime,
            "not in donation period"
        );
        donations[msg.sender] += msg.value;
        _updateTopDonors(msg.sender);
        emit Donation(msg.sender, msg.value);
    }

    /// @notice owner 设置捐赠时间段
    function setDonationPeriod(
        uint256 _startTime,
        uint256 _endTime
    ) external onlyOwner {
        require(_startTime < _endTime, "invalid period");
        startTime = _startTime;
        endTime = _endTime;
    }

    /// @notice 捐赠金额最多的前 3 个地址及其金额
    function getTopDonors()
        external
        view
        returns (address[3] memory donors, uint256[3] memory amounts)
    {
        donors = topDonors;
        for (uint256 i = 0; i < 3; i++) {
            amounts[i] = donations[donors[i]];
        }
    }

    /// @dev 捐赠后更新排行榜。
    function _updateTopDonors(address donor) private {
        uint256 amount = donations[donor];

        // 1. 找 donor 是否已在榜上
        uint256 i = 0;
        while (i < 3 && topDonors[i] != donor) {
            i++;
        }

        // 2. 不在榜上：只有比第 3 名多才能上榜，直接替换第 3 名
        if (i == 3) {
            if (amount <= donations[topDonors[2]]) {
                return;
            }
            i = 2;
            topDonors[2] = donor;
        }

        // 3. 向前冒泡：比前一名多就交换，直到位置正确
        while (i > 0 && amount > donations[topDonors[i - 1]]) {
            (topDonors[i], topDonors[i - 1]) = (topDonors[i - 1], topDonors[i]);
            i--;
        }
    }

    /// @notice 直接向合约地址转账
    receive() external payable {
        donate();
    }

    /// @notice 所有者提取全部余额
    function withdraw() external onlyOwner {
        // 获取当前合约的所有余额
        uint256 amount = address(this).balance;
        // 使用 call 代替 transfer，并安全地处理返回值
        (bool success, ) = payable(owner).call{value: amount}("");
        require(success, "Transfer failed.");
    }

    /// @notice 查询某地址的累计捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }
}
