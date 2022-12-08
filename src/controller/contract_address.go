package controller

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/smart-contract-demo/contracts"
	"github.com/hiroki-Fukumoto/smart-contract-demo/env"
)

type ContractAddressController interface {
	Get(c *gin.Context)
}

type contractAddressController struct {
}

func NewContractAddressController() ContractAddressController {
	return &contractAddressController{}
}

type contractAddressResponse struct {
	Address string `json:"address"`
}

// @Summary Get Contract Address
// @Description Get Contract Address
// @Tags contractAddress
// @Accept json
// @Produce json
// @Success 200 {object} contractAddressResponse
// @Router /api/v1/contract-address [get]
func (cn contractAddressController) Get(c *gin.Context) {
	client, err := ethclient.Dial(fmt.Sprintf("%s:%s", env.GanacheHost(), env.GanachePort()))
	if err != nil {
		log.Fatal(err)
	}

	priKey := "Your Private Key"
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address, _, _, err := contracts.DeployContracts(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	res := contractAddressResponse{
		Address: address.String(),
	}

	c.JSON(http.StatusOK, res)
}
