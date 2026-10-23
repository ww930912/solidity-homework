# 任务 1：区块链读写 · Sepolia 测试结果

测试时间：2026-10-23，网络：Sepolia 测试网，运行目录：`homework04/`

结论：查询区块、发送转账、查询交易三项功能都正常，转账已在 Sepolia 上链并执行成功（Status = 1）。

## 1. 查询最新区块

```bash
go run ./task1
```

```
=== Block ===
Number      : 11975486
Hash        : 0xe9f630a4cf88ff077fa89fae29df9013403958983edf6dc8896725d476b96cd0
Parent Hash : 0x251747e1b236149a517aeb96ff13548ce830d2fbff743331c7a23ffdc6723b3d
Timestamp   : 1792744284 (2026-10-23T16:31:24+08:00)
Tx Count    : 152
Gas Used    : 72035212 / 199999809
Miner       : 0x8f614CdB61E37B6f18d9705942c93FA86B04D711
```

## 2. 查询指定区块

```bash
go run ./task1 -number 11959386
```

```
=== Block ===
Number      : 11959386
Hash        : 0x208a112fd2c9a5d6872367c4587b2d15f6ed635b4c24f2ba5b5a50add43aa35a
Parent Hash : 0x46c59b9e94825cec0055001ce382f6925f3432d6233fe297b3382aa67f477898
Timestamp   : 1792548996 (2026-10-21T10:16:36+08:00)
Tx Count    : 130
Gas Used    : 74997530 / 200000000
Miner       : 0x8f614CdB61E37B6f18d9705942c93FA86B04D711
```

## 3. 发送转账（0.01 ETH）

```bash
go run ./task1 -send -to 0xF78c997295AE31E35a4E9dEbAA773D1Df9F74Bf2 -amount 0.01
```

```
=== Transaction Sent ===
From    : 0x83Da6626Ef0721dA9D452a7eeDc4F291A3f9d88C
To      : 0xF78c997295AE31E35a4E9dEbAA773D1Df9F74Bf2
Value   : 10000000000000000 Wei
Nonce   : 35
Tx Hash : 0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a

交易已提交，可用以下任一方式查看状态：
  浏览器: https://sepolia.etherscan.io/tx/0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a
  命令行: go run ./task1 -tx 0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a
```

Etherscan：<https://sepolia.etherscan.io/tx/0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a>

## 4. 查询交易及回执

```bash
go run ./task1 -tx 0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a
```

```
=== Transaction ===
Hash        : 0x0c411df9290d51c9f2d38a405ceedb91972b93e202b0ebf02274f3d4ccf96b2a
To          : 0xF78c997295AE31E35a4E9dEbAA773D1Df9F74Bf2
Value (Wei) : 10000000000000000
Nonce       : 35
Gas Limit   : 24393
Pending     : false
=== Receipt ===
Status      : 1 (1=成功, 0=失败)
Block       : 11975492
Gas Used    : 24055
Gas Price   : 18 Wei
```

## 结果解读

- **转账成功**：`Status = 1`，交易打包在区块 11975492，发送后约 6 个区块（约 1 分钟多）即上链。
- **Gas Used 是 24055，不是 21000**：普通 ETH 转账到普通地址固定消耗 21000 gas。多出来的部分说明接收地址上有代码（合约或 EIP-7702 委托账户），收款时执行了额外逻辑。这正是代码里用 `EstimateGas` 估算、而不是写死 21000 的原因，写死会导致交易因 out of gas 失败。
- **Gas Price 只有 18 Wei**：Sepolia 测试网负载低，baseFee 很小，所以这笔交易的手续费几乎可以忽略（24055 × 18 Wei）。
- **Gas Used 约占区块 Gas Limit 的 36%～37%**：低于 50% 的目标值，按 EIP-1559 规则下一个区块的 baseFee 会下降。
