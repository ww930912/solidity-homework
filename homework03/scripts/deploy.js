const { ethers, upgrades } = require("hardhat");

// Sepolia 上的 Chainlink 喂价地址（已在链上核实 description()），来自 https://docs.chain.link/data-feeds/price-feeds/addresses
const SEPOLIA_ETH_USD = "0x694AA1769357215DE4FAC081bf1f309aDC325306";
const SEPOLIA_USDC_USD = "0xA2F78ab2355fe2f984D808B5CeE7FD0A93D5270E";

// Sepolia 上的公共测试 USDT（6 位小数，非 Tether 官方）
const SEPOLIA_USDT = "0xaA8E23Fb1079EA71e0a56F48a2aA51851D8433D0";

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("部署账户:", deployer.address);

  const nft = await ethers.deployContract("MyNFT");
  await nft.waitForDeployment();
  console.log("MyNFT:", nft.target);

  // deployProxy 会部署实现合约 + ERC1967 代理，并调用 initialize()
  const auction = await upgrades.deployProxy(await ethers.getContractFactory("NFTAuction"), [], { kind: "uups" });
  await auction.waitForDeployment();
  console.log("NFTAuction 代理:", auction.target);
  console.log("NFTAuction 实现:", await upgrades.erc1967.getImplementationAddress(auction.target));

  await (await auction.setPriceFeed(ethers.ZeroAddress, SEPOLIA_ETH_USD)).wait();
  console.log("已设置 ETH/USD 喂价");
  // Sepolia 没有 USDT/USD 喂价，用同为 1 美元锚定的 USDC/USD 近似
  await (await auction.setPriceFeed(SEPOLIA_USDT, SEPOLIA_USDC_USD)).wait();
  console.log("已设置 USDT 喂价（使用 USDC/USD）");

  // 自检：读一次换算结果，确认喂价配置正确（结果为 8 位小数美元）
  const ethUsd = await auction.getUsdValue(ethers.ZeroAddress, ethers.parseEther("1"));
  const usdtUsd = await auction.getUsdValue(SEPOLIA_USDT, ethers.parseUnits("100", 6));
  console.log("自检 1 ETH   =", ethers.formatUnits(ethUsd, 8), "USD");
  console.log("自检 100 USDT =", ethers.formatUnits(usdtUsd, 8), "USD");
}

main().catch((e) => {
  console.error(e);
  process.exitCode = 1;
});
