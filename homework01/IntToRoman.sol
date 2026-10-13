// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 整数转罗马数字（1 ~ 3999）
contract IntToRoman {
    function intToRoman(uint256 num) external pure returns (string memory) {
        require(num > 0 && num < 4000, "out of range");

        // 从大到小排列，包含 900/400/90/40/9/4 这类特殊组合，直接贪心即可
        uint16[13] memory values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];

        bytes memory result;
        for (uint256 i = 0; i < 13; i++) {
            while (num >= values[i]) {
                result = abi.encodePacked(result, symbols[i]);
                num -= values[i];
            }
        }
        return string(result);
    }
}
