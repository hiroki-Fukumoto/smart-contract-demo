// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Shop {
    struct ShopStatsStruct {
        uint256 productCount;
        uint256 orderCount;
        uint256 sellerCount;
        uint256 salesCount;
        uint256 paid;
        uint256 balance;
    }

    ShopStatsStruct public shopStats;

    mapping(address => ShopStatsStruct) public shopStatsOf;
}
