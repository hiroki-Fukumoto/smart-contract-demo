package controller

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/smart-contract-demo/contracts"
	"github.com/hiroki-Fukumoto/smart-contract-demo/env"
)

type HelloController interface {
	HelloWorld(c *gin.Context)
}

type helloController struct {
}

func NewHelloController() HelloController {
	return &helloController{}
}

const (
	constantAddress = "Your Contract Address"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

// @Summary Hello World
// @Description Smart Contract demo.
// @Tags helloWorld
// @Accept json
// @Produce json
// @Success 200 {object} helloWorldResponse
// @Router /api/v1/hello [get]
func (h helloController) HelloWorld(c *gin.Context) {
	url := fmt.Sprintf("%s:%s", env.GanacheHost(), env.GanachePort())
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	conn, err := contracts.NewContracts(common.HexToAddress(constantAddress), client)
	if err != nil {
		panic(err)
	}

	response, err := conn.HelloWorld(&bind.CallOpts{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	res := helloWorldResponse{
		Message: response,
	}

	c.JSON(http.StatusOK, res)
}
