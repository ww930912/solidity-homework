// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

/// @title NFT 拍卖市场（UUPS 可升级）
/// @notice 出价可用 ETH 或 ERC20，统一按 Chainlink 价格换算成美元比较高低
contract NFTAuction is Initializable, UUPSUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    struct Auction {
        address seller;
        address nft;
        uint256 tokenId;
        uint256 endTime;
        uint256 startPriceUsd; // 起拍价（美元，8 位小数，与 Chainlink 一致）
        address highestBidder;
        address bidToken; // 最高出价所用币种，address(0) 表示 ETH
        uint256 highestBid; // 最高出价的币数量
        uint256 highestBidUsd; // 最高出价折合美元（8 位小数）
        bool ended;
    }

    // 拍卖 id => 拍卖信息
    mapping(uint256 => Auction) public auctions;
    uint256 public nextAuctionId;
    // 币种 => Chainlink 喂价合约（address(0) 为 ETH/USD）
    mapping(address => AggregatorV3Interface) public priceFeeds;

    event AuctionCreated(uint256 indexed auctionId, address indexed seller, address nft, uint256 tokenId, uint256 endTime);
    event BidPlaced(uint256 indexed auctionId, address indexed bidder, address token, uint256 amount, uint256 usdValue);
    event AuctionEnded(uint256 indexed auctionId, address winner, address token, uint256 amount, uint256 fee);

    // 额外挑战：动态手续费，按成交价（美元）分档，成交价越高费率越低。
    // 费率单位为基点（1 bp = 0.01%），用常量而非存储变量：不占存储槽、读取零成本，
    // 需要调整费率时通过 UUPS 升级修改即可。
    uint256 private constant TIER1_USD = 1_000 * 1e8; // $1,000（8 位小数）
    uint256 private constant TIER2_USD = 10_000 * 1e8; // $10,000
    uint256 private constant FEE_BPS_LOW = 300; // < $1,000：3%
    uint256 private constant FEE_BPS_MID = 200; // $1,000 ~ $10,000：2%
    uint256 private constant FEE_BPS_HIGH = 100; // >= $10,000：1%

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        // 锁住实现合约本身，防止有人直接初始化实现合约
        _disableInitializers();
    }

    /// @notice 代理合约的“构造函数”，部署时只调用一次
    function initialize() public initializer {
        __Ownable_init(msg.sender);
    }

    /// @notice 设置某币种的喂价合约，token = address(0) 表示 ETH
    function setPriceFeed(address token, address feed) external onlyOwner {
        priceFeeds[token] = AggregatorV3Interface(feed);
    }

    /// @notice 把 amount 个 token 换算成美元（结果 8 位小数）
    function getUsdValue(address token, uint256 amount) public view returns (uint256) {
        AggregatorV3Interface feed = priceFeeds[token];
        require(address(feed) != address(0), "no price feed");
        (, int256 price,,,) = feed.latestRoundData();
        require(price > 0, "invalid price");
        // 喂价本身带 feed.decimals() 位小数，这里统一折算成 8 位，方便不同币种直接比较
        uint8 tokenDecimals = token == address(0) ? 18 : IERC20Metadata(token).decimals();
        return amount * uint256(price) * 1e8 / (10 ** tokenDecimals) / (10 ** feed.decimals());
    }

    /// @notice 上架拍卖：NFT 先托管到本合约，需提前 approve
    function createAuction(address nft, uint256 tokenId, uint256 startPriceUsd, uint256 duration)
        external
        returns (uint256 auctionId)
    {
        require(duration > 0, "duration must > 0");
        IERC721(nft).transferFrom(msg.sender, address(this), tokenId);

        auctionId = nextAuctionId++;
        Auction storage a = auctions[auctionId];
        a.seller = msg.sender;
        a.nft = nft;
        a.tokenId = tokenId;
        a.endTime = block.timestamp + duration;
        a.startPriceUsd = startPriceUsd;
        emit AuctionCreated(auctionId, msg.sender, nft, tokenId, a.endTime);
    }

    /// @notice 出价。token = address(0) 时用 ETH（amount 取 msg.value），否则用 ERC20（需提前 approve）
    function bid(uint256 auctionId, address token, uint256 amount) external payable {
        Auction storage a = auctions[auctionId];
        require(a.seller != address(0), "auction not exist");
        require(block.timestamp < a.endTime, "auction ended");

        if (token == address(0)) {
            amount = msg.value;
        } else {
            require(msg.value == 0, "do not send ETH");
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        }

        uint256 usd = getUsdValue(token, amount);
        require(usd >= a.startPriceUsd && usd > a.highestBidUsd, "bid too low");

        // 先记下上一个出价者，更新状态后再退款（先改状态后转账，防重入）
        address prevBidder = a.highestBidder;
        address prevToken = a.bidToken;
        uint256 prevAmount = a.highestBid;

        a.highestBidder = msg.sender;
        a.bidToken = token;
        a.highestBid = amount;
        a.highestBidUsd = usd;

        if (prevBidder != address(0)) {
            _pay(prevToken, prevBidder, prevAmount);
        }
        emit BidPlaced(auctionId, msg.sender, token, amount, usd);
    }

    /// @notice 结束拍卖：任何人都可在到期后调用。NFT 给最高出价者，钱给卖家；无人出价则 NFT 退回卖家
    function endAuction(uint256 auctionId) external {
        Auction storage a = auctions[auctionId];
        require(a.seller != address(0), "auction not exist");
        require(block.timestamp >= a.endTime, "auction not over");
        require(!a.ended, "already ended");
        a.ended = true;

        uint256 fee;
        if (a.highestBidder == address(0)) {
            IERC721(a.nft).transferFrom(address(this), a.seller, a.tokenId);
        } else {
            IERC721(a.nft).transferFrom(address(this), a.highestBidder, a.tokenId);
            // 手续费按成交价的美元价值定档，但用出价币种本身扣除（ETH 成交就扣 ETH）
            fee = a.highestBid * getFeeBps(a.highestBidUsd) / 10_000;
            if (fee > 0) {
                _pay(a.bidToken, owner(), fee); // 手续费归平台（合约 owner）
            }
            _pay(a.bidToken, a.seller, a.highestBid - fee);
        }
        emit AuctionEnded(auctionId, a.highestBidder, a.bidToken, a.highestBid, fee);
    }

    /// @notice 根据成交价（美元，8 位小数）返回手续费率（基点）
    function getFeeBps(uint256 usdValue) public pure returns (uint256) {
        if (usdValue < TIER1_USD) return FEE_BPS_LOW;
        if (usdValue < TIER2_USD) return FEE_BPS_MID;
        return FEE_BPS_HIGH;
    }

    /// @dev 统一的付款函数：ETH 或 ERC20
    function _pay(address token, address to, uint256 amount) private {
        if (token == address(0)) {
            (bool ok,) = to.call{value: amount}("");
            require(ok, "ETH transfer failed");
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    /// @dev UUPS 要求：限制只有 owner 能升级
    function _authorizeUpgrade(address) internal override onlyOwner {}
}
