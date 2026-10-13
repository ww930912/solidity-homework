// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./NFTAuction.sol";

/// @title 用于演示升级的 V2 版本：只新增一个函数，不改存储布局
/// @custom:oz-upgrades-from NFTAuction
contract NFTAuctionV2 is NFTAuction {
    function version() external pure returns (string memory) {
        return "v2";
    }
}
