# Solidity 智能合约作业

MetaNode 学院 Web3 学习路线 · 智能合约方向的三次作业，包含 Solidity 基础练习、部署在测试网上的讨饭合约，以及一个基于 Hardhat 的 NFT 拍卖市场。

| 作业 | 内容 | 技术栈 | 测试网部署 | 详细文档 |
| --- | --- | --- | --- | --- |
| [作业 1](./homework01) | 投票合约 + 5 道算法题 | Solidity | — | 见下文 |
| [作业 2](./homework02) | 讨饭合约 BeggingContract | Solidity · Foundry | ✅ Sepolia | [homework02/README.md](./homework02/README.md) |
| [作业 3](./homework03) | NFT 拍卖市场 | Hardhat · OpenZeppelin · Chainlink · UUPS | ✅ Sepolia | [homework03/README.md](./homework03/README.md) |

## 目录结构

```
.
├── homework01/                 作业 1：Solidity 基础
│   ├── Voting.sol
│   ├── ReverseString.sol
│   ├── IntToRoman.sol
│   ├── RomanToInt.sol
│   ├── MergeSortedArray.sol
│   └── BinarySearch.sol
├── homework02/                 作业 2：讨饭合约
│   ├── BeggingContract.sol
│   ├── BeggingContract.standard-input.json   Etherscan 源码验证文件
│   ├── img/                    测试截图
│   └── README.md
└── homework03/                 作业 3：NFT 拍卖市场（Hardhat 项目）
    ├── contracts/              MyNFT、NFTAuction（UUPS）、NFTAuctionV2、mocks
    ├── scripts/deploy.js       Sepolia 部署脚本
    ├── test/                   单元测试 + 集成测试
    ├── .openzeppelin/          代理部署记录（升级时使用）
    ├── img/                    测试截图
    └── README.md
```

---

## 作业 1：Solidity 基础

每个合约都是独立文件，可以直接在 [Remix](https://remix.ethereum.org) 里编译和调用。

| 合约 | 题目 | 实现思路 | 示例 |
| --- | --- | --- | --- |
| [Voting.sol](./homework01/Voting.sol) | 投票合约：`vote` / `getVotes` / `resetVotes` | mapping 记录票数；mapping 不能遍历，所以另外用一个数组记录候选人，重置时逐个清零 | 投 a 两票 → `getVotes("a") = 2` → 重置后为 0 |
| [ReverseString.sol](./homework01/ReverseString.sol) | 反转字符串 | 转成 bytes 后，用双指针首尾交换 | `"abcde"` → `"edcba"` |
| [IntToRoman.sol](./homework01/IntToRoman.sol) | 整数转罗马数字（1～3999） | 贪心：按 1000、900、500…1 从大到小依次减去，同时拼上对应符号 | `1994` → `"MCMXCIV"` |
| [RomanToInt.sol](./homework01/RomanToInt.sol) | 罗马数字转整数 | **从右往左**遍历，比右边的字母小就减、否则加。不从左往右是因为 uint 不能为负，读 `"IV"` 时先减 1 会下溢 | `"MCMXCIV"` → `1994` |
| [MergeSortedArray.sol](./homework01/MergeSortedArray.sol) | 合并两个有序数组 | 双指针，每次取两个数组里较小的元素 | `[1,3,5] + [2,4]` → `[1,2,3,4,5]` |
| [BinarySearch.sol](./homework01/BinarySearch.sol) | 二分查找 | 左闭右开区间 `[left, right)`，避免计算 `mid - 1` 时下溢；找不到返回 -1 | 在 `[1..5]` 中找 `5` → `4` |

> 题目文档里第 3、4 题的 LeetCode 链接放反了，这里按题目标题实现，两个方向都完成了。

---

## 作业 2：讨饭合约

用户可以向合约捐赠 ETH，合约记录每个捐赠者的累计金额，所有者可以提取全部资金。三项额外挑战都已完成：

- ✅ 捐赠事件 `Donation`
- ✅ 捐赠排行榜：前 3 名
- ✅ 捐赠时间限制

| 项目 | 内容 |
| --- | --- |
| 合约地址（Sepolia） | [`0x7fb2b27ebDa5e371539EEb3325E14f779Be533Dd`](https://sepolia.etherscan.io/address/0x7fb2b27ebDa5e371539EEb3325E14f779Be533Dd) |
| 源码验证 | Etherscan ✅ |
| 测试 | 三个账户分别捐赠 1 / 2 / 0.5 ETH，排行榜按 2 > 1 > 0.5 排序 |

合约说明、部署方式和测试截图见 **[homework02/README.md](./homework02/README.md)**。

---

## 作业 3：NFT 拍卖市场

使用 Hardhat 开发的 NFT 拍卖市场：

- 出价可以用 ETH，也可以用 ERC20 代币；
- 通过 Chainlink 喂价把不同币种的出价统一换算成美元，再比较高低；
- 拍卖合约使用 UUPS 代理模式，部署后可以升级。

| 要求 | 完成情况 |
| --- | --- |
| ERC721 NFT：铸造和转移 | ✅ `MyNFT` |
| 创建拍卖 / 出价（ETH 或 ERC20）/ 结束拍卖 | ✅ `NFTAuction`，被超过的出价自动退款 |
| Chainlink 喂价换算美元 | ✅ ETH/USD、USDC/USD |
| UUPS 合约升级 | ✅ 有升级到 V2 的测试 |
| 单元测试和集成测试 | ✅ 16 个测试全部通过，行覆盖率 100% |
| 部署到 Sepolia | ✅ 已在测试网走完完整拍卖流程 |
| 额外挑战：动态手续费 | ✅ 按成交价分档：3% / 2% / 1% |

| 合约 | 地址（Sepolia） |
| --- | --- |
| MyNFT | [`0x7f25CceE06E0e6E4A43bB2D05Aa748707343c495`](https://sepolia.etherscan.io/address/0x7f25CceE06E0e6E4A43bB2D05Aa748707343c495) |
| **NFTAuction 代理**（用户交互地址） | [`0xD26429D7E3CeeA027A19d0A37F06429026257D5D`](https://sepolia.etherscan.io/address/0xD26429D7E3CeeA027A19d0A37F06429026257D5D) |
| NFTAuction 实现合约 | [`0x86aAC6602a715819E8393B8dCB5B952595627784`](https://sepolia.etherscan.io/address/0x86aAC6602a715819E8393B8dCB5B952595627784) |

测试网上的实际流程：

1. 卖家上架 NFT #0；
2. 一个账户用 0.01 ETH 出价，折合 $25；
3. 另一个账户用 50 USDT 出价，折合 $49.99，超过前一个出价，ETH 自动退回；
4. 拍卖结束：NFT 转给赢家，卖家收到 48.5 USDT，平台收到 1.5 USDT 手续费（3%）。

本地运行：

```bash
cd homework03
npm install
npx hardhat test        # 运行测试
npx hardhat coverage    # 覆盖率报告
```

项目结构、功能说明、部署步骤、完整测试报告和截图见 **[homework03/README.md](./homework03/README.md)**。

---

## 说明

- 所有合约只部署在 **Sepolia 测试网**，仅用于学习，没有经过安全审计。
- 作业 3 部署时需要的私钥和 API Key 放在 `.env` 文件里（已在 `.gitignore` 中排除）。配置方法参考 [homework03/.env.example](./homework03/.env.example)。
