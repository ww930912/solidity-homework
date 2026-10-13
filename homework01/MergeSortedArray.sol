// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 合并两个有序数组
contract MergeSortedArray {
    /// @notice [1,3,5] + [2,4] => [1,2,3,4,5]
    function merge(uint256[] calldata a, uint256[] calldata b) external pure returns (uint256[] memory) {
        uint256[] memory result = new uint256[](a.length + b.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;
        // 双指针：每次取两个数组当前较小的元素
        while (i < a.length && j < b.length) {
            result[k++] = a[i] <= b[j] ? a[i++] : b[j++];
        }
        // 把剩余部分直接拷过去
        while (i < a.length) result[k++] = a[i++];
        while (j < b.length) result[k++] = b[j++];
        return result;
    }
}
