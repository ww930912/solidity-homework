// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title 罗马数字转整数
contract RomanToInt {
    function romanToInt(string calldata s) external pure returns (uint256) {
        bytes memory b = bytes(s);
        uint256 total = 0;
        uint256 prev = 0;
        for (uint256 i = b.length; i > 0; i--) {
            uint256 cur = _value(b[i - 1]);
            if (cur < prev) {
                total -= cur;
            } else {
                total += cur;
            }
            prev = cur;
        }
        return total;
    }

    function _value(bytes1 c) private pure returns (uint256) {
        if (c == "I") return 1;
        if (c == "V") return 5;
        if (c == "X") return 10;
        if (c == "L") return 50;
        if (c == "C") return 100;
        if (c == "D") return 500;
        if (c == "M") return 1000;
        revert("invalid char");
    }
}
