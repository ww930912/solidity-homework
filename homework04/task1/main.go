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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 任务 1：区块链读写（Sepolia）
//
// 环境变量：
//
//	ETH_RPC_URL        例如 https://sepolia.infura.io/v3/<API_KEY>
//	SENDER_PRIVATE_KEY 发送交易时需要（不带 0x 也可）
//
// 使用示例：
//
//	go run ./task1                                   # 查询最新区块
//	go run ./task1 -number 5000000                   # 查询指定区块
//	go run ./task1 -send -to 0xabc... -amount 0.001  # 发送 ETH 转账
//	go run ./task1 -tx 0x<hash>                      # 查询交易及回执
func main() {
	number := flag.Uint64("number", 0, "要查询的区块号（0 表示最新区块）")
	txHash := flag.String("tx", "", "要查询的交易哈希")
	send := flag.Bool("send", false, "发送转账交易")
	to := flag.String("to", "", "接收方地址（-send 时必填）")
	amount := flag.Float64("amount", 0, "转账金额，单位 ETH（-send 时必填）")
	flag.Parse()

	rpcURL := os.Getenv("ETH_RPC_URL")
	if rpcURL == "" {
		log.Fatal("ETH_RPC_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()

	if *send {
		if !common.IsHexAddress(*to) || *amount <= 0 {
			log.Fatal("send mode requires valid -to and -amount > 0")
		}
		sendTransaction(ctx, client, common.HexToAddress(*to), *amount)
		return
	}
	if *txHash != "" {
		queryTransaction(ctx, client, common.HexToHash(*txHash))
		return
	}
	queryBlock(ctx, client, *number)
}

// queryTransaction 查询交易本身和回执；pending 交易暂时没有回执。
func queryTransaction(ctx context.Context, client *ethclient.Client, hash common.Hash) {
	tx, isPending, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		log.Fatalf("failed to get transaction: %v", err)
	}

	fmt.Println("=== Transaction ===")
	fmt.Printf("Hash        : %s\n", tx.Hash().Hex())
	fmt.Printf("To          : %v\n", tx.To())
	fmt.Printf("Value (Wei) : %s\n", tx.Value().String())
	fmt.Printf("Nonce       : %d\n", tx.Nonce())
	fmt.Printf("Gas Limit   : %d\n", tx.Gas())
	fmt.Printf("Pending     : %v\n", isPending)
	if isPending {
		fmt.Println("交易还在交易池中等待打包，稍后再查。")
		return
	}

	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Fatalf("failed to get receipt: %v", err)
	}
	fmt.Println("=== Receipt ===")
	fmt.Printf("Status      : %d (1=成功, 0=失败)\n", receipt.Status)
	fmt.Printf("Block       : %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("Gas Used    : %d\n", receipt.GasUsed)
	if receipt.EffectiveGasPrice != nil {
		fmt.Printf("Gas Price   : %s Wei\n", receipt.EffectiveGasPrice.String())
	}
}

// queryBlock 查询区块并打印哈希、时间戳、交易数量等信息。
func queryBlock(ctx context.Context, client *ethclient.Client, number uint64) {
	var num *big.Int // nil 表示最新区块
	if number > 0 {
		num = new(big.Int).SetUint64(number)
	}

	block, err := client.BlockByNumber(ctx, num)
	if err != nil {
		log.Fatalf("failed to get block: %v", err)
	}

	fmt.Println("=== Block ===")
	fmt.Printf("Number      : %d\n", block.Number().Uint64())
	fmt.Printf("Hash        : %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash : %s\n", block.ParentHash().Hex())
	fmt.Printf("Timestamp   : %d (%s)\n", block.Time(), time.Unix(int64(block.Time()), 0).Format(time.RFC3339))
	fmt.Printf("Tx Count    : %d\n", len(block.Transactions()))
	fmt.Printf("Gas Used    : %d / %d\n", block.GasUsed(), block.GasLimit())
	fmt.Printf("Miner       : %s\n", block.Coinbase().Hex())
}

// sendTransaction 构造、签名并发送一笔 EIP-1559 ETH 转账，输出交易哈希。
func sendTransaction(ctx context.Context, client *ethclient.Client, toAddr common.Address, amountEth float64) {
	privKeyHex := os.Getenv("SENDER_PRIVATE_KEY")
	if privKeyHex == "" {
		log.Fatal("SENDER_PRIVATE_KEY is not set")
	}
	privKey, err := crypto.HexToECDSA(trim0x(privKeyHex))
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}
	fromAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// EIP-1559 费用：gasFeeCap = baseFee*2 + tip，给 baseFee 上涨留出余量
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("failed to get gas tip cap: %v", err)
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get header: %v", err)
	}
	gasFeeCap := new(big.Int).Add(new(big.Int).Mul(header.BaseFee, big.NewInt(2)), gasTipCap)

	value := ethToWei(amountEth)

	// 目标地址可能是合约，不能写死 21000，用 EstimateGas 估算
	gasLimit, err := client.EstimateGas(ctx, ethereum.CallMsg{From: fromAddr, To: &toAddr, Value: value})
	if err != nil {
		log.Fatalf("failed to estimate gas: %v", err)
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     value,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privKey)
	if err != nil {
		log.Fatalf("failed to sign tx: %v", err)
	}
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("failed to send tx: %v", err)
	}

	fmt.Println("=== Transaction Sent ===")
	fmt.Printf("From    : %s\n", fromAddr.Hex())
	fmt.Printf("To      : %s\n", toAddr.Hex())
	fmt.Printf("Value   : %s Wei\n", value.String())
	fmt.Printf("Nonce   : %d\n", nonce)
	fmt.Printf("Tx Hash : %s\n", signedTx.Hash().Hex())
	fmt.Println("\n交易已提交，可用以下任一方式查看状态：")
	fmt.Printf("  浏览器: https://sepolia.etherscan.io/tx/%s\n", signedTx.Hash().Hex())
	fmt.Printf("  命令行: go run ./task1 -tx %s\n", signedTx.Hash().Hex())
}

// ethToWei 把 ETH 数量转换为 wei。
func ethToWei(amountEth float64) *big.Int {
	wei, _ := new(big.Float).Mul(big.NewFloat(amountEth), big.NewFloat(1e18)).Int(nil)
	return wei
}

// trim0x 去掉首尾空白和 "0x"/"0X" 前缀。
// 私钥通常是从钱包或文档复制来的，容易带上空格、换行或大写前缀。
func trim0x(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}
	return s
}
