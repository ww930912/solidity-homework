package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"homework04/task2/nft"
	"homework04/task2/token"
)

// 任务 2：使用 abigen 生成的绑定代码与计数器合约交互
//
// 环境变量：ETH_RPC_URL、SENDER_PRIVATE_KEY
//
// 使用示例：
//
//	go run ./task2 -type erc20 -deploy           # 部署 CounterToken（ERC20）并调用一次 increment
//	go run ./task2 -type erc721 -deploy          # 部署 CounterNFT（ERC721，带 logo）并调用一次 increment
//	go run ./task2 -type erc20 -addr 0xabc...    # 对已部署的合约调用 increment
func main() {
	contractType := flag.String("type", "", "合约类型：erc20 或 erc721")
	deploy := flag.Bool("deploy", false, "先部署一个新合约")
	addrHex := flag.String("addr", "", "已部署的合约地址")
	flag.Parse()

	if *contractType != "erc20" && *contractType != "erc721" {
		log.Fatal("-type must be erc20 or erc721")
	}
	if !*deploy && !common.IsHexAddress(*addrHex) {
		log.Fatal("use -deploy or -addr <contract address>")
	}

	rpcURL := os.Getenv("ETH_RPC_URL")
	// 兼容复制私钥时带上的空白和大小写 0x 前缀
	privKeyHex := strings.TrimSpace(os.Getenv("SENDER_PRIVATE_KEY"))
	privKeyHex = strings.TrimPrefix(strings.TrimPrefix(privKeyHex, "0x"), "0X")
	if rpcURL == "" || privKeyHex == "" {
		log.Fatal("ETH_RPC_URL and SENDER_PRIVATE_KEY must be set")
	}

	// 部署和等待上链比较慢，超时给宽一点
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	privKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}
	// TransactOpts 负责签名，nonce / gas 留空时由绑定代码自动估算
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("failed to create transactor: %v", err)
	}
	auth.Context = ctx
	callOpts := &bind.CallOpts{Context: ctx}

	var c counter
	if *contractType == "erc20" {
		c = setupToken(ctx, client, auth, callOpts, *deploy, *addrHex)
	} else {
		c = setupNFT(ctx, client, auth, callOpts, *deploy, *addrHex)
	}
	runIncrement(ctx, client, auth, callOpts, c)
}

// counter 是两个合约共有的计数器方法。
// abigen 为 CounterToken 和 CounterNFT 生成的这两个方法签名一致，因此可以用同一套流程调用。
type counter interface {
	Count(opts *bind.CallOpts) (*big.Int, error)
	Increment(opts *bind.TransactOpts) (*types.Transaction, error)
}

// setupToken 部署或绑定 CounterToken，并打印 ERC20 基本信息。
func setupToken(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, callOpts *bind.CallOpts, deploy bool, addrHex string) *token.CounterToken {
	var (
		addr     common.Address
		instance *token.CounterToken
		err      error
	)
	if deploy {
		var tx *types.Transaction
		addr, tx, instance, err = token.DeployCounterToken(auth, client)
		if err != nil {
			log.Fatalf("failed to deploy: %v", err)
		}
		waitDeployed(ctx, client, tx, addr)
		fmt.Printf("代币页面: https://sepolia.etherscan.io/token/%s\n", addr.Hex())
	} else {
		addr = common.HexToAddress(addrHex)
		if instance, err = token.NewCounterToken(addr, client); err != nil {
			log.Fatalf("failed to bind contract: %v", err)
		}
	}

	name, err := instance.Name(callOpts)
	if err != nil {
		log.Fatalf("failed to read name: %v", err)
	}
	symbol, err := instance.Symbol(callOpts)
	if err != nil {
		log.Fatalf("failed to read symbol: %v", err)
	}
	supply, err := instance.TotalSupply(callOpts)
	if err != nil {
		log.Fatalf("failed to read totalSupply: %v", err)
	}
	balance, err := instance.BalanceOf(callOpts, auth.From)
	if err != nil {
		log.Fatalf("failed to read balance: %v", err)
	}
	fmt.Printf("代币     : %s (%s)\n", name, symbol)
	fmt.Printf("总供应量 : %s（最小单位，18 位小数）\n", supply)
	fmt.Printf("我的余额 : %s\n", balance)
	return instance
}

// setupNFT 部署或绑定 CounterNFT，打印 #1 号 NFT 信息并把链上 logo 保存为文件。
func setupNFT(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, callOpts *bind.CallOpts, deploy bool, addrHex string) *nft.CounterNFT {
	var (
		addr     common.Address
		instance *nft.CounterNFT
		err      error
	)
	if deploy {
		var tx *types.Transaction
		addr, tx, instance, err = nft.DeployCounterNFT(auth, client)
		if err != nil {
			log.Fatalf("failed to deploy: %v", err)
		}
		waitDeployed(ctx, client, tx, addr)
		// 构造函数给部署者铸造了 #1 号 NFT，Etherscan 的 NFT 页面会通过 tokenURI 渲染 logo
		fmt.Printf("NFT 页面: https://sepolia.etherscan.io/nft/%s/1\n", addr.Hex())
	} else {
		addr = common.HexToAddress(addrHex)
		if instance, err = nft.NewCounterNFT(addr, client); err != nil {
			log.Fatalf("failed to bind contract: %v", err)
		}
	}

	owner, err := instance.OwnerOf(callOpts, big.NewInt(1))
	if err != nil {
		log.Fatalf("failed to read owner: %v", err)
	}
	logo, err := instance.Logo(callOpts)
	if err != nil {
		log.Fatalf("failed to read logo: %v", err)
	}
	// 把链上读回的 SVG 落盘，用浏览器打开即可看到 logo
	if err := os.WriteFile("logo_from_chain.svg", []byte(logo), 0o644); err != nil {
		log.Fatalf("failed to save logo: %v", err)
	}
	fmt.Printf("NFT #1 持有者: %s\n", owner.Hex())
	fmt.Printf("Logo     : 已保存到 logo_from_chain.svg（%d 字节）\n", len(logo))
	return instance
}

// waitDeployed 等待部署交易上链。
// 必须等部署完成，否则紧接着读合约会因为地址上还没有代码而失败。
func waitDeployed(ctx context.Context, client *ethclient.Client, tx *types.Transaction, addr common.Address) {
	fmt.Printf("部署交易: %s\n等待上链...\n", tx.Hash().Hex())
	if _, err := bind.WaitDeployed(ctx, client, tx); err != nil {
		log.Fatalf("deploy not mined: %v", err)
	}
	fmt.Printf("合约地址: %s\n", addr.Hex())
}

// runIncrement 读取 count，调用 increment 并等待上链，再读取新的 count。
func runIncrement(ctx context.Context, client *ethclient.Client, auth *bind.TransactOpts, callOpts *bind.CallOpts, c counter) {
	before, err := c.Count(callOpts)
	if err != nil {
		log.Fatalf("failed to read count: %v", err)
	}
	fmt.Printf("调用前 count: %s\n", before)

	tx, err := c.Increment(auth)
	if err != nil {
		log.Fatalf("failed to call increment: %v", err)
	}
	fmt.Printf("increment 交易: %s\n等待上链...\n", tx.Hash().Hex())
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("increment not mined: %v", err)
	}

	after, err := c.Count(callOpts)
	if err != nil {
		log.Fatalf("failed to read count: %v", err)
	}
	fmt.Printf("区块号   : %d, 状态: %d (1=成功)\n", receipt.BlockNumber.Uint64(), receipt.Status)
	fmt.Printf("调用后 count: %s\n", after)
}
