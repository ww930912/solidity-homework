require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("dotenv").config();

const { SEPOLIA_RPC_URL, PRIVATE_KEY, ETHERSCAN_API_KEY } = process.env;

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  // OpenZeppelin 5.x 用到了 mcopy 指令，需要 cancun 版本 EVM（Sepolia 已支持）
  solidity: { version: "0.8.28", settings: { evmVersion: "cancun" } },
  networks: {
    // 只有配置了环境变量才启用 sepolia，避免本地测试时报错
    ...(SEPOLIA_RPC_URL && PRIVATE_KEY
      ? { sepolia: { url: SEPOLIA_RPC_URL, accounts: [PRIVATE_KEY] } }
      : {}),
  },
  etherscan: { apiKey: ETHERSCAN_API_KEY || "" },
};
