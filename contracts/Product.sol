// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import "./Shop.sol";

contract Product is Shop {
    struct ProductStruct {
        uint256 id;
        string sku;
        address seller;
        string name;
        string imageURL;
        string description;
        uint256 price;
        uint256 timestamp;
        bool deleted;
        uint256 stock;
    }

    uint256 public fee;
    ProductStruct[] products;

    mapping(uint256 => bool) public productExist;

    event NewProduct(string name, string description, string imageURL);

    modifier createProductValidate(
        string memory _name,
        string memory _sku,
        string memory _description,
        string memory _imageURL,
        uint256 _price,
        uint256 _stock
    ) {
        require(msg.value >= fee, "Insufficient fund");
        require(bytes(_name).length > 0, "name can't be empty");
        require(bytes(_sku).length > 0, "sku can't be empty");
        require(bytes(_description).length > 0, "description can't be empty");
        require(bytes(_imageURL).length > 0, "image URL can't be empty");
        require(_price > 0, "price can't be zero");
        require(_stock > 0, "stock can't be zero");
        _;
    }

    function createProduct(
        string memory _name,
        string memory _sku,
        string memory _description,
        string memory _imageURL,
        uint256 _price,
        uint256 _stock
    )
        public
        payable
        createProductValidate(
            _name,
            _sku,
            _description,
            _imageURL,
            _price,
            _stock
        )
    {
        productExist[shopStats.productCount] = true;
        shopStatsOf[msg.sender].productCount += 1;
        shopStats.sellerCount += 1;

        ProductStruct memory product;
        product.id = shopStats.productCount += 1;
        product.name = _name;
        product.sku = _sku;
        product.description = _description;
        product.imageURL = _imageURL;
        product.price = _price;
        product.stock = _stock;
        product.seller = msg.sender;
        product.timestamp = block.timestamp;
        products.push(product);

        emit NewProduct(_name, _description, _imageURL);
    }
}
