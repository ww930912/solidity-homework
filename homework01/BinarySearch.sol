// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 二分查找
contract BinarySearch {
    /// @notice 在升序数组中查找 target，找到返回下标，找不到返回 -1
    function search(
        uint256[] calldata nums,
        uint256 target
    ) external pure returns (int256) {
        uint256 left = 0;
        uint256 right = nums.length;
        while (left < right) {
            uint256 mid = left + (right - left) / 2;
            if (nums[mid] == target) return int256(mid);
            if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return -1;
    }
}
