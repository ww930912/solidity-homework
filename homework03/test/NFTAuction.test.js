const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");
const { time, loadFixture } = require("@nomicfoundation/hardhat-network-helpers");

const ETH = ethers.ZeroAddress; // address(0) 代表 ETH
const usd = (n) => BigInt(n) * 10n ** 8n; // 美元金额，8 位小数
const usdc = (n) => BigInt(n) * 10n ** 6n; // MockERC20 为 6 位小数
const DAY = 24 * 60 * 60;

describe("NFTAuction", function () {
  // 公共部署流程：NFT + ERC20 + 两个喂价 + 拍卖代理，并给卖家铸造一个 NFT 上架
  async function deployFixture() {
    const [owner, seller, alice, bob] = await ethers.getSigners();

    const nft = await ethers.deployContract("MyNFT");
    const token = await ethers.deployContract("MockERC20");
    const ethFeed = await ethers.deployContract("MockV3Aggregator", [8, usd(2000)]); // 1 ETH = 2000 USD
    const tokenFeed = await ethers.deployContract("MockV3Aggregator", [8, usd(1)]); // 1 mUSDC = 1 USD

    const Auction = await ethers.getContractFactory("NFTAuction");
    const auction = await upgrades.deployProxy(Auction, [], { kind: "uups" });
    await auction.setPriceFeed(ETH, ethFeed.target);
    await auction.setPriceFeed(token.target, tokenFeed.target);

    await nft.mint(seller.address); // tokenId = 0
    await nft.connect(seller).approve(auction.target, 0);
    await auction.connect(seller).createAuction(nft.target, 0, usd(100), DAY); // auctionId = 0

    for (const u of [alice, bob]) {
      await token.mint(u.address, usdc(100000));
      await token.connect(u).approve(auction.target, ethers.MaxUint256);
    }
    return { nft, token, auction, owner, seller, alice, bob };
  }

  describe("NFT", function () {
    it("只有 owner 能铸造，持有人可转移", async function () {
      const { nft, seller, alice } = await loadFixture(deployFixture);
      await expect(nft.connect(alice).mint(alice.address)).to.be.revertedWithCustomError(nft, "OwnableUnauthorizedAccount");
      await nft.mint(alice.address); // tokenId = 1
      await nft.connect(alice).transferFrom(alice.address, seller.address, 1);
      expect(await nft.ownerOf(1)).to.equal(seller.address);
    });
  });

  describe("价格换算", function () {
    it("ETH 和 ERC20 都能换算成美元", async function () {
      const { auction, token } = await loadFixture(deployFixture);
      expect(await auction.getUsdValue(ETH, ethers.parseEther("1.5"))).to.equal(usd(3000));
      expect(await auction.getUsdValue(token.target, usdc(250))).to.equal(usd(250));
    });

    it("未配置喂价的币种会 revert", async function () {
      const { auction, alice } = await loadFixture(deployFixture);
      await expect(auction.getUsdValue(alice.address, 1)).to.be.revertedWith("no price feed");
    });
  });

  describe("拍卖流程", function () {
    it("创建拍卖后 NFT 托管在合约中", async function () {
      const { nft, auction, seller } = await loadFixture(deployFixture);
      expect(await nft.ownerOf(0)).to.equal(auction.target);
      expect((await auction.auctions(0)).seller).to.equal(seller.address);
    });

    it("低于起拍价或不高于当前最高价的出价被拒绝", async function () {
      const { auction, alice, bob, token } = await loadFixture(deployFixture);
      await expect(auction.connect(alice).bid(0, token.target, usdc(50))).to.be.revertedWith("bid too low");
      await auction.connect(alice).bid(0, token.target, usdc(200));
      // 0.1 ETH = 200 USD，与当前最高价相等，不算更高
      await expect(auction.connect(bob).bid(0, ETH, 0, { value: ethers.parseEther("0.1") })).to.be.revertedWith("bid too low");
    });

    it("ETH 与 ERC20 混合出价，按美元比较，被超越者自动退款", async function () {
      const { auction, token, alice, bob } = await loadFixture(deployFixture);
      await auction.connect(alice).bid(0, token.target, usdc(200)); // 200 USD
      const before = await token.balanceOf(alice.address);

      // bob 出 0.2 ETH = 400 USD，超过 alice
      await expect(auction.connect(bob).bid(0, ETH, 0, { value: ethers.parseEther("0.2") }))
        .to.emit(auction, "BidPlaced")
        .withArgs(0, bob.address, ETH, ethers.parseEther("0.2"), usd(400));
      expect(await token.balanceOf(alice.address)).to.equal(before + usdc(200));

      // alice 再用 ERC20 出 500 USD 超过 bob，bob 的 ETH 被退回
      await expect(auction.connect(alice).bid(0, token.target, usdc(500))).to.changeEtherBalance(bob, ethers.parseEther("0.2"));
      const a = await auction.auctions(0);
      expect(a.highestBidder).to.equal(alice.address);
      expect(a.highestBidUsd).to.equal(usd(500));
    });

    it("用 ERC20 出价时不能附带 ETH", async function () {
      const { auction, token, alice } = await loadFixture(deployFixture);
      await expect(auction.connect(alice).bid(0, token.target, usdc(200), { value: 1 })).to.be.revertedWith("do not send ETH");
    });

    it("到期前不能结束，到期后不能出价", async function () {
      const { auction, alice } = await loadFixture(deployFixture);
      await expect(auction.endAuction(0)).to.be.revertedWith("auction not over");
      await time.increase(DAY);
      await expect(auction.connect(alice).bid(0, ETH, 0, { value: ethers.parseEther("1") })).to.be.revertedWith("auction ended");
    });

    it("结束拍卖（ETH 成交）：NFT 给赢家，扣 2% 手续费后 ETH 给卖家，不可重复结束", async function () {
      const { nft, auction, owner, seller, bob } = await loadFixture(deployFixture);
      await auction.connect(bob).bid(0, ETH, 0, { value: ethers.parseEther("1") }); // $2,000 落在 2% 档
      await time.increase(DAY);
      // 由 bob 调用 endAuction，避免 owner 自己付 gas 干扰余额断言
      await expect(auction.connect(bob).endAuction(0)).to.changeEtherBalances(
        [seller, owner],
        [ethers.parseEther("0.98"), ethers.parseEther("0.02")]
      );
      expect(await nft.ownerOf(0)).to.equal(bob.address);
      await expect(auction.endAuction(0)).to.be.revertedWith("already ended");
    });

    it("结束拍卖（ERC20 成交）：扣 3% 手续费后代币给卖家", async function () {
      const { nft, token, auction, owner, seller, alice } = await loadFixture(deployFixture);
      await auction.connect(alice).bid(0, token.target, usdc(300)); // $300 落在 3% 档
      await time.increase(DAY);
      await expect(auction.endAuction(0))
        .to.emit(auction, "AuctionEnded")
        .withArgs(0, alice.address, token.target, usdc(300), usdc(9));
      expect(await nft.ownerOf(0)).to.equal(alice.address);
      expect(await token.balanceOf(seller.address)).to.equal(usdc(291));
      expect(await token.balanceOf(owner.address)).to.equal(usdc(9));
    });

    it("高价成交（>= $10,000）只收 1% 手续费", async function () {
      const { auction, owner, seller, bob } = await loadFixture(deployFixture);
      await auction.connect(bob).bid(0, ETH, 0, { value: ethers.parseEther("6") }); // $12,000
      await time.increase(DAY);
      await expect(auction.connect(bob).endAuction(0)).to.changeEtherBalances(
        [seller, owner],
        [ethers.parseEther("5.94"), ethers.parseEther("0.06")]
      );
    });

    it("无人出价时 NFT 退回卖家", async function () {
      const { nft, auction, seller } = await loadFixture(deployFixture);
      await time.increase(DAY);
      await auction.endAuction(0);
      expect(await nft.ownerOf(0)).to.equal(seller.address);
    });

    it("不存在的拍卖会 revert", async function () {
      const { auction, alice } = await loadFixture(deployFixture);
      await expect(auction.connect(alice).bid(99, ETH, 0, { value: 1 })).to.be.revertedWith("auction not exist");
      await expect(auction.endAuction(99)).to.be.revertedWith("auction not exist");
    });
  });

  describe("动态手续费", function () {
    it("按成交价分档，边界值归入更低费率档", async function () {
      const { auction } = await loadFixture(deployFixture);
      expect(await auction.getFeeBps(usd(999))).to.equal(300);
      expect(await auction.getFeeBps(usd(1000))).to.equal(200); // 恰好 $1,000 进入 2% 档
      expect(await auction.getFeeBps(usd(9999))).to.equal(200);
      expect(await auction.getFeeBps(usd(10000))).to.equal(100); // 恰好 $10,000 进入 1% 档
    });
  });

  describe("UUPS 升级", function () {
    it("升级到 V2 后数据保留、新函数可用", async function () {
      const { auction, alice } = await loadFixture(deployFixture);
      await auction.connect(alice).bid(0, ETH, 0, { value: ethers.parseEther("1") });

      const V2 = await ethers.getContractFactory("NFTAuctionV2");
      const upgraded = await upgrades.upgradeProxy(auction.target, V2);
      expect(upgraded.target).to.equal(auction.target); // 代理地址不变
      expect(await upgraded.version()).to.equal("v2");
      expect((await upgraded.auctions(0)).highestBidder).to.equal(alice.address); // 存储数据还在
    });

    it("非 owner 不能升级、不能设置喂价", async function () {
      const { auction, alice } = await loadFixture(deployFixture);
      const V2 = await ethers.getContractFactory("NFTAuctionV2", alice);
      await expect(upgrades.upgradeProxy(auction.target, V2)).to.be.revertedWithCustomError(auction, "OwnableUnauthorizedAccount");
      await expect(auction.connect(alice).setPriceFeed(ETH, alice.address)).to.be.revertedWithCustomError(auction, "OwnableUnauthorizedAccount");
    });
  });
});
