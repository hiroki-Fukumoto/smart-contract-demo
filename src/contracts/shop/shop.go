// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shop

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

// ShopMetaData contains all meta data concerning the Shop contract.
var ShopMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"shopStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"productCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shopStatsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"productCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610129806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806306f862cb14603757806338e3a2eb1460a8575b600080fd5b6077604236600460c5565b600660205260009081526040902080546001820154600283015460038401546004850154600590950154939492939192909186565b604080519687526020870195909552938501929092526060840152608083015260a082015260c00160405180910390f35b600054600154600254600354600454600554607795949392919086565b60006020828403121560d657600080fd5b81356001600160a01b038116811460ec57600080fd5b939250505056fea26469706673582212207aec591153535109d8b4c13b84bc072c0926f49f03d855fe567baa86d10c460064736f6c63430008110033",
}

// ShopABI is the input ABI used to generate the binding from.
// Deprecated: Use ShopMetaData.ABI instead.
var ShopABI = ShopMetaData.ABI

// ShopBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ShopMetaData.Bin instead.
var ShopBin = ShopMetaData.Bin

// DeployShop deploys a new Ethereum contract, binding an instance of Shop to it.
func DeployShop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Shop, error) {
	parsed, err := ShopMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShopBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Shop{ShopCaller: ShopCaller{contract: contract}, ShopTransactor: ShopTransactor{contract: contract}, ShopFilterer: ShopFilterer{contract: contract}}, nil
}

// Shop is an auto generated Go binding around an Ethereum contract.
type Shop struct {
	ShopCaller     // Read-only binding to the contract
	ShopTransactor // Write-only binding to the contract
	ShopFilterer   // Log filterer for contract events
}

// ShopCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShopCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShopTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShopTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShopFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShopFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShopSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShopSession struct {
	Contract     *Shop             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShopCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShopCallerSession struct {
	Contract *ShopCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ShopTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShopTransactorSession struct {
	Contract     *ShopTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShopRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShopRaw struct {
	Contract *Shop // Generic contract binding to access the raw methods on
}

// ShopCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShopCallerRaw struct {
	Contract *ShopCaller // Generic read-only contract binding to access the raw methods on
}

// ShopTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShopTransactorRaw struct {
	Contract *ShopTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShop creates a new instance of Shop, bound to a specific deployed contract.
func NewShop(address common.Address, backend bind.ContractBackend) (*Shop, error) {
	contract, err := bindShop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Shop{ShopCaller: ShopCaller{contract: contract}, ShopTransactor: ShopTransactor{contract: contract}, ShopFilterer: ShopFilterer{contract: contract}}, nil
}

// NewShopCaller creates a new read-only instance of Shop, bound to a specific deployed contract.
func NewShopCaller(address common.Address, caller bind.ContractCaller) (*ShopCaller, error) {
	contract, err := bindShop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShopCaller{contract: contract}, nil
}

// NewShopTransactor creates a new write-only instance of Shop, bound to a specific deployed contract.
func NewShopTransactor(address common.Address, transactor bind.ContractTransactor) (*ShopTransactor, error) {
	contract, err := bindShop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShopTransactor{contract: contract}, nil
}

// NewShopFilterer creates a new log filterer instance of Shop, bound to a specific deployed contract.
func NewShopFilterer(address common.Address, filterer bind.ContractFilterer) (*ShopFilterer, error) {
	contract, err := bindShop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShopFilterer{contract: contract}, nil
}

// bindShop binds a generic wrapper to an already deployed contract.
func bindShop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShopABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shop *ShopRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shop.Contract.ShopCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shop *ShopRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shop.Contract.ShopTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shop *ShopRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shop.Contract.ShopTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Shop *ShopCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Shop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Shop *ShopTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Shop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Shop *ShopTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Shop.Contract.contract.Transact(opts, method, params...)
}

// ShopStats is a free data retrieval call binding the contract method 0x38e3a2eb.
//
// Solidity: function shopStats() view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Shop *ShopCaller) ShopStats(opts *bind.CallOpts) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	var out []interface{}
	err := _Shop.contract.Call(opts, &out, "shopStats")

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
func (_Shop *ShopSession) ShopStats() (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Shop.Contract.ShopStats(&_Shop.CallOpts)
}

// ShopStats is a free data retrieval call binding the contract method 0x38e3a2eb.
//
// Solidity: function shopStats() view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Shop *ShopCallerSession) ShopStats() (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Shop.Contract.ShopStats(&_Shop.CallOpts)
}

// ShopStatsOf is a free data retrieval call binding the contract method 0x06f862cb.
//
// Solidity: function shopStatsOf(address ) view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Shop *ShopCaller) ShopStatsOf(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	var out []interface{}
	err := _Shop.contract.Call(opts, &out, "shopStatsOf", arg0)

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
func (_Shop *ShopSession) ShopStatsOf(arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Shop.Contract.ShopStatsOf(&_Shop.CallOpts, arg0)
}

// ShopStatsOf is a free data retrieval call binding the contract method 0x06f862cb.
//
// Solidity: function shopStatsOf(address ) view returns(uint256 productCount, uint256 orderCount, uint256 sellerCount, uint256 salesCount, uint256 paid, uint256 balance)
func (_Shop *ShopCallerSession) ShopStatsOf(arg0 common.Address) (struct {
	ProductCount *big.Int
	OrderCount   *big.Int
	SellerCount  *big.Int
	SalesCount   *big.Int
	Paid         *big.Int
	Balance      *big.Int
}, error) {
	return _Shop.Contract.ShopStatsOf(&_Shop.CallOpts, arg0)
}
