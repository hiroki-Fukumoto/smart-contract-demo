package product

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/contracts/product"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/ethereumclient"
)

type ProductService interface {
	Create(req CreateProductRequest) error
}

type productService struct{}

const (
	constantAddress = "0x9FE3fFb0ce5547568b7d9B9b6992ADe32B4c7118"
)

func NewProductService() ProductService {
	return &productService{}
}

func (s productService) Create(req CreateProductRequest) error {
	client, err := ethereumclient.GetEthereumClient()
	if err != nil {
		return err
	}

	conn, err := product.NewProduct(common.HexToAddress(constantAddress), client)
	if err != nil {
		return err
	}

	res, err := conn.CreateProduct(&bind.TransactOpts{
		From: common.HexToAddress(constantAddress),
	}, req.Name, req.Sku, req.Description, req.ImageURL, req.Price, req.Stock)
	if err != nil {
		return err
	}
	log.Println(res)

	return nil
}
