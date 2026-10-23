# 任务 2：abigen 合约绑定 · Sepolia 测试结果

测试时间：2026-10-23，网络：Sepolia 测试网，运行目录：`homework04/`

结论：两个合约（ERC20 的 `CounterToken`、ERC721 的 `CounterNFT`）的编译、abigen 生成绑定、部署、调用 `increment` 全部成功，两个计数器在链上都从 0 → 1 → 2。

## 1. 编译合约，生成 ABI 和字节码

```bash
"$HOME/Library/Application Support/svm/0.8.37/solc-0.8.37" --base-path . --include-path task2/node_modules \
  --abi --bin --overwrite -o task2/build task2/contracts/CounterToken.sol task2/contracts/CounterNFT.sol
```

```
Compiler run successful. Artifact(s) can be found in directory "task2/build".
```

用到的产物：`CounterToken.abi/.bin`、`CounterNFT.abi/.bin`（`build/` 下其余文件是 OpenZeppelin 依赖合约，忽略即可）。

## 2. abigen 生成 Go 绑定代码

```bash
mkdir -p task2/token task2/nft   # abigen 不会自动创建输出目录
abigen --abi task2/build/CounterToken.abi --bin task2/build/CounterToken.bin \
       --pkg token --type CounterToken --out task2/token/token.go
abigen --abi task2/build/CounterNFT.abi --bin task2/build/CounterNFT.bin \
       --pkg nft --type CounterNFT --out task2/nft/nft.go
```

执行成功无输出，生成两个 Go 包：

| 包 | 主要方法 |
|---|---|
| `task2/token` | `DeployCounterToken`、`NewCounterToken`、`Name()`、`Symbol()`、`TotalSupply()`、`BalanceOf()`、`Count()`、`Increment()` |
| `task2/nft` | `DeployCounterNFT`、`NewCounterNFT`、`OwnerOf()`、`TokenURI()`、`Logo()`、`Count()`、`Increment()` |

## 3. ERC20：CounterToken

### 3.1 部署并调用 increment

```bash
go run ./task2 -type erc20 -deploy
```

```
部署交易: 0xbdd4783c71fcdd5505038778646403a6e9828ffc3fefe11cb1ae9e564a0cff20
等待上链...
合约地址: 0xC9fB6CA5265DacB5400215498F1D8086A1F0Aa6E
代币页面: https://sepolia.etherscan.io/token/0xC9fB6CA5265DacB5400215498F1D8086A1F0Aa6E
代币     : Counter Token (CNT)
总供应量 : 1000000000000000000000（最小单位，18 位小数）
我的余额 : 1000000000000000000000
调用前 count: 0
increment 交易: 0x5e8e7cc223d3c462d0c2c0f844330e8d114124247cfd8d4474262a928d5f5442
等待上链...
区块号   : 11975695, 状态: 1 (1=成功)
调用后 count: 1
```

### 3.2 对已部署合约再次调用

```bash
go run ./task2 -type erc20 -addr 0xC9fB6CA5265DacB5400215498F1D8086A1F0Aa6E
```

```
代币     : Counter Token (CNT)
总供应量 : 1000000000000000000000（最小单位，18 位小数）
我的余额 : 1000000000000000000000
调用前 count: 1
increment 交易: 0xe1764a06a1c6fb82930052d1d7a6cd57d03ec9d8e572c236ae2ac92b5d4df54f
等待上链...
区块号   : 11975699, 状态: 1 (1=成功)
调用后 count: 2
```

## 4. ERC721：CounterNFT（带 logo）

### 4.1 部署并调用 increment

```bash
go run ./task2 -type erc721 -deploy
```

```
部署交易: 0xee3ba4c093fe6c7f171f735947abd30061abcd6b1538e56c6631e8957ead1220
等待上链...
合约地址: 0x2d603Ff01E88F5dcF4f0e8F23750010700a33c01
NFT 页面: https://sepolia.etherscan.io/nft/0x2d603Ff01E88F5dcF4f0e8F23750010700a33c01/1
NFT #1 持有者: 0x83Da6626Ef0721dA9D452a7eeDc4F291A3f9d88C
Logo     : 已保存到 logo_from_chain.svg（214 字节）
调用前 count: 0
increment 交易: 0x7a2ecdefe8505ac3a4d265a9b8daf9d84cd243c8ab57283d55f2c4b629ee9079
等待上链...
区块号   : 11975702, 状态: 1 (1=成功)
调用后 count: 1
```

### 4.2 对已部署合约再次调用

```bash
go run ./task2 -type erc721 -addr 0x2d603Ff01E88F5dcF4f0e8F23750010700a33c01
```

```
NFT #1 持有者: 0x83Da6626Ef0721dA9D452a7eeDc4F291A3f9d88C
Logo     : 已保存到 logo_from_chain.svg（214 字节）
调用前 count: 1
increment 交易: 0xfada59e1e0cb9249ed1dcb6f5850d778648187f64e70dbf7ac3492228f228b88
等待上链...
区块号   : 11975706, 状态: 1 (1=成功)
调用后 count: 2
```

## 链上记录（Etherscan）

| 项目 | 链接 |
|---|---|
| CounterToken 代币页面 | <https://sepolia.etherscan.io/token/0xC9fB6CA5265DacB5400215498F1D8086A1F0Aa6E> |
| CounterToken 部署交易 | <https://sepolia.etherscan.io/tx/0xbdd4783c71fcdd5505038778646403a6e9828ffc3fefe11cb1ae9e564a0cff20> |
| CounterToken increment #1 | <https://sepolia.etherscan.io/tx/0x5e8e7cc223d3c462d0c2c0f844330e8d114124247cfd8d4474262a928d5f5442> |
| CounterToken increment #2 | <https://sepolia.etherscan.io/tx/0xe1764a06a1c6fb82930052d1d7a6cd57d03ec9d8e572c236ae2ac92b5d4df54f> |
| CounterNFT #1 页面（logo） | <https://sepolia.etherscan.io/nft/0x2d603Ff01E88F5dcF4f0e8F23750010700a33c01/1> |
| CounterNFT 部署交易 | <https://sepolia.etherscan.io/tx/0xee3ba4c093fe6c7f171f735947abd30061abcd6b1538e56c6631e8957ead1220> |
| CounterNFT increment #1 | <https://sepolia.etherscan.io/tx/0x7a2ecdefe8505ac3a4d265a9b8daf9d84cd243c8ab57283d55f2c4b629ee9079> |
| CounterNFT increment #2 | <https://sepolia.etherscan.io/tx/0xfada59e1e0cb9249ed1dcb6f5850d778648187f64e70dbf7ac3492228f228b88> |

## 结果解读

- **状态持久化**：第二次运行是全新进程，`调用前 count` 读到的是 1，说明计数存在合约 storage 里，而不是程序内存里。
- **ERC20 供应量**：`1000000000000000000000` = 1000 × 10^18，即 1000 枚 CNT（decimals = 18）。链上只存整数，"1000 枚"是前端按 decimals 换算后的显示值。
- **ERC721 的 logo**：部署时铸造给部署者 `0x83Da…d88C` 的 #1 号 NFT，`tokenURI(1)` 返回的元数据里 `image` 字段就是合约内的 SVG logo；程序同时把 `logo()` 读回的 214 字节 SVG 保存为 `logo_from_chain.svg`。
- **读写两类调用**：`Name()`、`Count()`、`OwnerOf()` 等是 view 调用（`eth_call`），不发交易、不花 gas；`Increment()` 是写操作，要签名发交易、付 gas，并用 `bind.WaitMined` 等待打包后才能读到新值。
- **同一套流程调两个合约**：abigen 为两个合约生成的 `Count`/`Increment` 签名一致，`main.go` 用一个小接口统一调用，`-type` 只决定绑定哪个合约。

> 历史记录：第一版单文件 `Counter.sol`（不继承任何标准）曾部署在 [0xaacC3D125053D969dfF25732205C1765DF6a7557](https://sepolia.etherscan.io/address/0xaacC3D125053D969dfF25732205C1765DF6a7557)，现已被上面两个合约取代。
