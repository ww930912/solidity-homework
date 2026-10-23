// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// @title CounterToken - 带计数器的 ERC20 代币
/// @notice 代币图标不写在合约里：钱包和浏览器从 Etherscan、Token List 等链下渠道获取。
contract CounterToken is ERC20 {
    uint256 public count;

    event Incremented(address indexed by, uint256 newCount);

    /// @notice 部署时给部署者铸造 1000 枚代币（decimals 默认 18）
    constructor() ERC20("Counter Token", "CNT") {
        _mint(msg.sender, 1000 * 10 ** decimals());
    }

    /// @notice 计数器加 1
    function increment() external {
        count += 1;
        emit Incremented(msg.sender, count);
    }
}
