// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package product

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProductMetaData contains all meta data concerning the Product contract.
var ProductMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURL\",\"type\":\"string\"}],\"name\":\"NewProduct\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sku\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stock\",\"type\":\"uint256\"}],\"name\":\"createProduct\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"productExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shopStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"productCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shopStatsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"productCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109c6806100206000396000f3fe60806040526004361061004a5760003560e01c806306f862cb1461004f57806338e3a2eb146100d1578063c39e0b2b146100fc578063ddca3f4314610111578063e77a67ea14610135575b600080fd5b34801561005b57600080fd5b5061009f61006a3660046105ec565b600660205260009081526040902080546001820154600283015460038401546004850154600590950154939492939192909186565b604080519687526020870195909552938501929092526060840152608083015260a082015260c0015b60405180910390f35b3480156100dd57600080fd5b5060005460015460025460035460045460055461009f95949392919086565b61010f61010a3660046106bf565b610175565b005b34801561011d57600080fd5b5061012760075481565b6040519081526020016100c8565b34801561014157600080fd5b5061016561015036600461077e565b60096020526000908152604090205460ff1681565b60405190151581526020016100c8565b8585858585856007543410156101c65760405162461bcd60e51b8152602060048201526011602482015270125b9cdd59999a58da595b9d08199d5b99607a1b60448201526064015b60405180910390fd5b600086511161020d5760405162461bcd60e51b81526020600482015260136024820152726e616d652063616e277420626520656d70747960681b60448201526064016101bd565b60008551116102535760405162461bcd60e51b8152602060048201526012602482015271736b752063616e277420626520656d70747960701b60448201526064016101bd565b60008451116102a45760405162461bcd60e51b815260206004820152601a60248201527f6465736372697074696f6e2063616e277420626520656d70747900000000000060448201526064016101bd565b60008351116102f55760405162461bcd60e51b815260206004820152601860248201527f696d6167652055524c2063616e277420626520656d707479000000000000000060448201526064016101bd565b6000821161033b5760405162461bcd60e51b815260206004820152601360248201527270726963652063616e2774206265207a65726f60681b60448201526064016101bd565b600081116103815760405162461bcd60e51b815260206004820152601360248201527273746f636b2063616e2774206265207a65726f60681b60448201526064016101bd565b600080548152600960209081526040808320805460ff19166001908117909155338452600690925282208054919290916103bc908490610797565b909155505060028054600191906000906103d7908490610797565b9250508190555061043f604051806101400160405280600081526020016060815260200160006001600160a01b031681526020016060815260200160608152602001606081526020016000815260200160008152602001600015158152602001600081525090565b60016000800160008282546104549190610797565b9182905550825250606081018d9052602081018c815260a082018c9052608082018b905260c082018a905261012082018990523360408301524260e0830152600880546001810182556000919091528251600a9091027ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee38101918255915183927ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee401906105019082610847565b5060408201516002820180546001600160a01b0319166001600160a01b039092169190911790556060820151600382019061053c9082610847565b50608082015160048201906105519082610847565b5060a082015160058201906105669082610847565b5060c0820151600682015560e0820151600782015561010082015160088201805460ff1916911515919091179055610120909101516009909101556040517fa7bfe4e266155b451ea2f70cfe3c91628080e7dc35c6aec52839101d5002a105906105d5908f908e908e9061094d565b60405180910390a150505050505050505050505050565b6000602082840312156105fe57600080fd5b81356001600160a01b038116811461061557600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261064357600080fd5b813567ffffffffffffffff8082111561065e5761065e61061c565b604051601f8301601f19908116603f011681019082821181831017156106865761068661061c565b8160405283815286602085880101111561069f57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060008060c087890312156106d857600080fd5b863567ffffffffffffffff808211156106f057600080fd5b6106fc8a838b01610632565b9750602089013591508082111561071257600080fd5b61071e8a838b01610632565b9650604089013591508082111561073457600080fd5b6107408a838b01610632565b9550606089013591508082111561075657600080fd5b5061076389828a01610632565b9350506080870135915060a087013590509295509295509295565b60006020828403121561079057600080fd5b5035919050565b808201808211156107b857634e487b7160e01b600052601160045260246000fd5b92915050565b600181811c908216806107d257607f821691505b6020821081036107f257634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561084257600081815260208120601f850160051c8101602086101561081f5750805b601f850160051c820191505b8181101561083e5782815560010161082b565b5050505b505050565b815167ffffffffffffffff8111156108615761086161061c565b6108758161086f84546107be565b846107f8565b602080601f8311600181146108aa57600084156108925750858301515b600019600386901b1c1916600185901b17855561083e565b600085815260208120601f198616915b828110156108d9578886015182559484019460019091019084016108ba565b50858210156108f75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000815180845260005b8181101561092d57602081850181015186830182015201610911565b506000602082860101526020601f19601f83011685010191505092915050565b6060815260006109606060830186610907565b82810360208401526109728186610907565b905082810360408401526109868185610907565b969550505050505056fea2646970667358221220e9a33823565832362f5bdfe463272959c96c0be57b4f33fd58a9f30fbd70d7ca64736f6c63430008110033",
}

// ProductABI is the input ABI used to generate the binding from.
// Deprecated: Use ProductMetaData.ABI instead.
var ProductABI = ProductMetaData.ABI

// ProductBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProductMetaData.Bin instead.
var ProductBin = ProductMetaData.Bin

// DeployProduct deploys a new Ethereum contract, binding an instance of Product to it.
func DeployProduct(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Product, error) {
	parsed, err := ProductMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProductBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Product{ProductCaller: ProductCaller{contract: contract}, ProductTransactor: ProductTransactor{contract: contract}, ProductFilterer: ProductFilterer{contract: contract}}, nil
}

// Product is an auto generated Go binding around an Ethereum contract.
type Product struct {
	ProductCaller     // Read-only binding to the contract
	ProductTransactor // Write-only binding to the contract
	ProductFilterer   // Log filterer for contract events
}

// ProductCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProductCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProductTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProductFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProductSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProductSession struct {
	Contract     *Product          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProductCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProductCallerSession struct {
	Contract *ProductCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProductTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProductTransactorSession struct {
	Contract     *ProductTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProductRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProductRaw struct {
	Contract *Product // Generic contract binding to access the raw methods on
}

// ProductCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProductCallerRaw struct {
	Contract *ProductCaller // Generic read-only contract binding to access the raw methods on
}

// ProductTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProductTransactorRaw struct {
	Contract *ProductTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProduct creates a new instance of Product, bound to a specific deployed contract.
func NewProduct(address common.Address, backend bind.ContractBackend) (*Product, error) {
	contract, err := bindProduct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Product{ProductCaller: ProductCaller{contract: contract}, ProductTransactor: ProductTransactor{contract: contract}, ProductFilterer: ProductFilterer{contract: contract}}, nil
}

// NewProductCaller creates a new read-only instance of Product, bound to a specific deployed contract.
func NewProductCaller(address common.Address, caller bind.ContractCaller) (*ProductCaller, error) {
	contract, err := bindProduct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProductCaller{contract: contract}, nil
}

// NewProductTransactor creates a new write-only instance of Product, bound to a specific deployed contract.
func NewProductTransactor(address common.Address, transactor bind.ContractTransactor) (*ProductTransactor, error) {
	contract, err := bindProduct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProductTransactor{contract: contract}, nil
}

// NewProductFilterer creates a new log filterer instance of Product, bound to a specific deployed contract.
func NewProductFilterer(address common.Address, filterer bind.ContractFilterer) (*ProductFilterer, error) {
	contract, err := bindProduct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProductFilterer{contract: contract}, nil
}

// bindProduct binds a generic wrapper to an already deployed contract.
func bindProduct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProductABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Product *ProductRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Product.Contract.ProductCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Product *ProductRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Product.Contract.ProductTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Product *ProductRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Product.Contract.ProductTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Product *ProductCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Product.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Product *ProductTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Product.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Product *ProductTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Product.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Product *ProductCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Product *ProductSession) Fee() (*big.Int, error) {
	return _Product.Contract.Fee(&_Product.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Product *ProductCallerSession) Fee() (*big.Int, error) {
	return _Product.Contract.Fee(&_Product.CallOpts)
}

// ProductExist is a free data retrieval call binding the contract method 0xe77a67ea.
//
// Solidity: function productExist(uint256 ) view returns(bool)
func (_Product *ProductCaller) ProductExist(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "productExist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductExist is a free data retrieval call binding the contract method 0xe77a67ea.
//
// Solidity: function productExist(uint256 ) view returns(bool)
func (_Product *ProductSession) ProductExist(arg0 *big.Int) (bool, error) {
	return _Product.Contract.ProductExist(&_Product.CallOpts, arg0)
}

// ProductExist is a free data retrieval call binding the contract method 0xe77a67ea.
//
// Solidity: function productExist(uint256 ) view returns(bool)
func (_Product *ProductCallerSession) ProductExist(arg0 *big.Int) (bool, error) {
	return _Product.Contract.ProductExist(&_Product.CallOpts, arg0)
}

// ShopStats is a free data retrieval call binding the contract method 0x38e3a2eb.
//
// Solidity: function shopStats() view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductCaller) ShopStats(opts *bind.CallOpts) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "shopStats")

	outstruct := new(struct {
		ProductCount *big.Int
		OrderCount   *big.Int
		SellerCount  *big.Int
		SalesCount   *big.Int
		Paid         *big.Int
		Balance      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProductCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OrderCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellerCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SalesCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Paid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ShopStats is a free data retrieval call binding the contract method 0x38e3a2eb.
//
// Solidity: function shopStats() view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductSession) ShopStats() (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Product.Contract.ShopStats(&_Product.CallOpts)
}

// ShopStats is a free data retrieval call binding the contract method 0x38e3a2eb.
//
// Solidity: function shopStats() view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductCallerSession) ShopStats() (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Product.Contract.ShopStats(&_Product.CallOpts)
}

// ShopStatsOf is a free data retrieval call binding the contract method 0x06f862cb.
//
// Solidity: function shopStatsOf(address ) view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductCaller) ShopStatsOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	var out []interface{}
	err := _Product.contract.Call(opts, &out, "shopStatsOf", arg0)

	outstruct := new(struct {
		ProductCount *big.Int
		OrderCount   *big.Int
		SellerCount  *big.Int
		SalesCount   *big.Int
		Paid         *big.Int
		Balance      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProductCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OrderCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SellerCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SalesCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Paid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ShopStatsOf is a free data retrieval call binding the contract method 0x06f862cb.
//
// Solidity: function shopStatsOf(address ) view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductSession) ShopStatsOf(arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Product.Contract.ShopStatsOf(&_Product.CallOpts, arg0)
}

// ShopStatsOf is a free data retrieval call binding the contract method 0x06f862cb.
//
// Solidity: function shopStatsOf(address ) view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Product *ProductCallerSession) ShopStatsOf(arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Product.Contract.ShopStatsOf(&_Product.CallOpts, arg0)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xc39e0b2b.
//
// Solidity: function createProduct(string _name, string _sku, string _description, string _imageURL, uint256 _price, uint256 _stock) payable returns()
func (_Product *ProductTransactor) CreateProduct(opts *bind.TransactOpts, _name string, _sku string, _description string, _imageURL string, _price *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Product.contract.Transact(opts, "createProduct", _name, _sku, _description, _imageURL, _price, _stock)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xc39e0b2b.
//
// Solidity: function createProduct(string _name, string _sku, string _description, string _imageURL, uint256 _price, uint256 _stock) payable returns()
func (_Product *ProductSession) CreateProduct(_name string, _sku string, _description string, _imageURL string, _price *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CreateProduct(&_Product.TransactOpts, _name, _sku, _description, _imageURL, _price, _stock)
}

// CreateProduct is a paid mutator transaction binding the contract method 0xc39e0b2b.
//
// Solidity: function createProduct(string _name, string _sku, string _description, string _imageURL, uint256 _price, uint256 _stock) payable returns()
func (_Product *ProductTransactorSession) CreateProduct(_name string, _sku string, _description string, _imageURL string, _price *big.Int, _stock *big.Int) (*types.Transaction, error) {
	return _Product.Contract.CreateProduct(&_Product.TransactOpts, _name, _sku, _description, _imageURL, _price, _stock)
}

// ProductNewProductIterator is returned from FilterNewProduct and is used to iterate over the raw logs and unpacked data for NewProduct events raised by the Product contract.
type ProductNewProductIterator struct {
	Event *ProductNewProduct // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProductNewProductIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProductNewProduct)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProductNewProduct)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProductNewProductIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProductNewProductIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProductNewProduct represents a NewProduct event raised by the Product contract.
type ProductNewProduct struct {
	Name        string
	Description string
	ImageURL    string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewProduct is a free log retrieval operation binding the contract event 0xa7bfe4e266155b451ea2f70cfe3c91628080e7dc35c6aec52839101d5002a105.
//
// Solidity: event NewProduct(string name, string description, string imageURL)
func (_Product *ProductFilterer) FilterNewProduct(opts *bind.FilterOpts) (*ProductNewProductIterator, error) {

	logs, sub, err := _Product.contract.FilterLogs(opts, "NewProduct")
	if err != nil {
		return nil, err
	}
	return &ProductNewProductIterator{contract: _Product.contract, event: "NewProduct", logs: logs, sub: sub}, nil
}

// WatchNewProduct is a free log subscription operation binding the contract event 0xa7bfe4e266155b451ea2f70cfe3c91628080e7dc35c6aec52839101d5002a105.
//
// Solidity: event NewProduct(string name, string description, string imageURL)
func (_Product *ProductFilterer) WatchNewProduct(opts *bind.WatchOpts, sink chan<- *ProductNewProduct) (event.Subscription, error) {

	logs, sub, err := _Product.contract.WatchLogs(opts, "NewProduct")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProductNewProduct)
				if err := _Product.contract.UnpackLog(event, "NewProduct", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewProduct is a log parse operation binding the contract event 0xa7bfe4e266155b451ea2f70cfe3c91628080e7dc35c6aec52839101d5002a105.
//
// Solidity: event NewProduct(string name, string description, string imageURL)
func (_Product *ProductFilterer) ParseNewProduct(log types.Log) (*ProductNewProduct, error) {
	event := new(ProductNewProduct)
	if err := _Product.contract.UnpackLog(event, "NewProduct", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
