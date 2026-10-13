// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 反转字符串
contract ReverseString {
    /// @notice "abcde" => "edcba"
    function reverse(string calldata s) external pure returns (string memory) {
        bytes memory b = bytes(s);
        // 双指针首尾交换
        uint256 n = b.length;
        for (uint256 i = 0; i < n / 2; i++) {
            (b[i], b[n - 1 - i]) = (b[n - 1 - i], b[i]);
        }
        return string(b);
    }
}
