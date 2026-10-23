// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Base64} from "@openzeppelin/contracts/utils/Base64.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

/// @title CounterNFT - 带 NFT logo 的计数器合约（ERC721）
/// @notice 部署时给部署者铸造 #1 号 NFT，其图片就是合约 logo，
///         Etherscan / 钱包会通过 tokenURI 渲染出这张图。
contract CounterNFT is ERC721 {
    // 合约 logo（SVG）。
    // 用 constant 而不是 storage 变量：常量编译进字节码，只在部署时付一次费用，读取不消耗 SLOAD。
    // SVG 属性用单引号，避免和 Solidity 字符串、JSON 的双引号冲突。
    string public constant logo =
        "<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 64 64'><circle cx='32' cy='32' r='30' fill='#6366f1'/><text x='32' y='42' font-size='28' font-family='monospace' text-anchor='middle' fill='#fff'>C6</text></svg>";

    uint256 public count;

    event Incremented(address indexed by, uint256 newCount);

    constructor() ERC721("Counter NFT", "CNFT") {
        _mint(msg.sender, 1);
    }

    /// @notice 计数器加 1
    function increment() external {
        count += 1;
        emit Incremented(msg.sender, count);
    }

    /// @notice 返回 NFT 元数据，完全存在链上，不依赖 IPFS 或服务器。
    /// @dev 格式：data:application/json;base64,<JSON>，JSON 的 image 字段又是一个 base64 的 SVG。
    ///      用 base64 而不是直接拼原文，是因为 SVG 里的 '#'、'<' 等字符放进 URI 需要转义，容易出错。
    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        _requireOwned(tokenId);
        string memory json = string.concat(
            '{"name":"Counter #',
            Strings.toString(tokenId),
            '","description":"Counter contract logo NFT","image":"data:image/svg+xml;base64,',
            Base64.encode(bytes(logo)),
            '","attributes":[{"trait_type":"count","value":',
            Strings.toString(count),
            "}]}"
        );
        return string.concat("data:application/json;base64,", Base64.encode(bytes(json)));
    }
}
