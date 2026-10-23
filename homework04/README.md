## 任务 1：区块链读写 任务目标

使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
具体任务

1. 环境搭建
   - 安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
   - 注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
2. 查询区块
   - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
   - 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
   - 输出查询结果到控制台。
3. 发送交易
   - 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
   - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
   - 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
   - 对交易进行签名，并将签名后的交易发送到网络。
   - 输出交易的哈希值。

## 任务 2：合约代码生成 任务目标

使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
具体任务

1. 编写智能合约
   - 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
   - 计数器合约需要一个logo，可自创一个。
   - 编译智能合约，生成 ABI 和字节码文件。
2. 使用 abigen 生成 Go 绑定代码
   - 安装 abigen 工具。
   - 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
3. 使用生成的 Go 绑定代码与合约交互
   - 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
   - 调用合约的方法，例如增加计数器的值。
   - 输出调用结果。

---

## 作业完成说明

### 目录结构

```
homework04/
├── go.mod / go.sum            # 单一 Go module（go-ethereum v1.17.8）
├── task1/main.go              # 任务 1：查询区块 + 发送转账
└── task2/
    ├── contracts/
    │   ├── CounterToken.sol   # ERC20 + 计数器，不带 logo（代币图标走链下渠道）
    │   └── CounterNFT.sol     # ERC721 + 计数器，#1 号 NFT 的图片即 logo（内嵌 SVG，tokenURI 返回）
    ├── package.json           # 只用来安装 @openzeppelin/contracts（node_modules 已被 .gitignore 忽略）
    ├── build/                 # solc 编译产物，只用 CounterToken/CounterNFT 的 .abi/.bin，其余是依赖合约
    ├── token/token.go         # abigen 生成的 ERC20 绑定（不要手改）
    ├── nft/nft.go             # abigen 生成的 ERC721 绑定（不要手改）
    └── main.go                # 任务 2：-type erc20|erc721 选择合约，部署 / 调用
```

### 0. 准备环境变量（两个任务共用）

```bash
cd homework04
export ETH_RPC_URL=https://sepolia.infura.io/v3/<你的 Infura API Key>
export SENDER_PRIVATE_KEY=<测试账户私钥>   # 只用测试号！不要提交到 git
```

> 测试账户需要先在 Sepolia 水龙头领一点 ETH（如 https://sepoliafaucet.com 、Google Cloud Web3 Faucet）。

### 任务 1：区块链读写

```bash
go run ./task1                     # 查询最新区块
go run ./task1 -number 11959386     # 查询指定区块：哈希、时间戳、交易数量等
go run ./task1 -send -to <接收地址> -amount 0.001   # 发送 0.001 ETH，输出交易哈希
go run ./task1 -tx <交易哈希>        # 查询交易详情和回执（是否成功、所在区块、Gas 消耗）
```

发送成功后会打印 `Tx Hash`，可打开 Etherscan 链接，或用 `-tx` 在命令行查询交易状态。

### 任务 2：abigen 生成合约绑定

**第 1 步：安装 OpenZeppelin，编译合约生成 ABI 和字节码**

```bash
npm install --prefix task2   # 安装 @openzeppelin/contracts 到 task2/node_modules

# --include-path 让 solc 能找到 import "@openzeppelin/..."（任选一种 solc）
# 方式 A：已装 solc（brew install solidity）
solc --base-path . --include-path task2/node_modules --abi --bin --overwrite -o task2/build \
     task2/contracts/CounterToken.sol task2/contracts/CounterNFT.sol
# 方式 B：已装 Foundry，用 svm 下载的 solc：
# "$HOME/Library/Application Support/svm/0.8.37/solc-0.8.37" --base-path . --include-path task2/node_modules \
#      --abi --bin --overwrite -o task2/build task2/contracts/CounterToken.sol task2/contracts/CounterNFT.sol
```

**第 2 步：安装 abigen 并生成 Go 代码**

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@v1.17.8
mkdir -p task2/token task2/nft   # abigen 不会自动创建输出目录

# --pkg 是 Go 包名（要和目录名一致），--type 是生成的 Go 结构体名
abigen --abi task2/build/CounterToken.abi --bin task2/build/CounterToken.bin \
       --pkg token --type CounterToken --out task2/token/token.go
abigen --abi task2/build/CounterNFT.abi --bin task2/build/CounterNFT.bin \
       --pkg nft --type CounterNFT --out task2/nft/nft.go
```

> 仓库里已经带了生成好的 `build/`、`token/`、`nft/`，修改合约后重新执行这两步即可。

**第 3 步：部署并调用合约**（`-type` 选择操作哪个合约）

```bash
go run ./task2 -type erc20 -deploy             # 部署 ERC20 → 读名称/供应量/余额 → increment
go run ./task2 -type erc721 -deploy            # 部署 ERC721 → 读 NFT 持有者/保存 logo → increment
go run ./task2 -type erc20 -addr <合约地址>     # 对已部署的合约再 +1（erc721 同理）
```

ERC20 预期输出：

```
合约地址: 0x...
代币页面: https://sepolia.etherscan.io/token/0x...
代币     : Counter Token (CNT)
总供应量 : 1000000000000000000000（最小单位，18 位小数）
我的余额 : 1000000000000000000000
调用前 count: 0
increment 交易: 0x...
区块号   : ..., 状态: 1 (1=成功)
调用后 count: 1
```

ERC721 预期输出：

```
合约地址: 0x...
NFT 页面: https://sepolia.etherscan.io/nft/0x.../1
NFT #1 持有者: 0x...
Logo     : 已保存到 logo_from_chain.svg（214 字节）
调用前 count: 0
increment 交易: 0x...
区块号   : ..., 状态: 1 (1=成功)
调用后 count: 1
```

### 本地快速验证（不想花测试币时）

```bash
anvil                                  # 另开终端，启动本地链（Foundry 自带）
export ETH_RPC_URL=http://127.0.0.1:8545
export SENDER_PRIVATE_KEY=<anvil 启动时打印的 (0) 号私钥>
go run ./task1 -send -to 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -amount 0.01
go run ./task2 -type erc721 -deploy
```
