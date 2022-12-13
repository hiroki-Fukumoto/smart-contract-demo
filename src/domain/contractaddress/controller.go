package contractaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/errorhandler"
)

type ContractAddressController interface {
	Get(c *gin.Context)
}

type contractAddressController struct {
	contractAddressService ContractAddressService
}

func NewContractAddressController(s ContractAddressService) ContractAddressController {
	return &contractAddressController{
		contractAddressService: s,
	}
}

// @Summary Get Contract Address
// @Description Get Contract Address
// @Tags contractAddress
// @Accept json
// @Produce json
// @Success 200 {object} ContractAddressResponse
// @Failure 400 {object} errorhandler.ErrorResponse
// @Router /api/v1/contract-address [get]
func (cn contractAddressController) Get(c *gin.Context) {
	address, err := cn.contractAddressService.GetContractAddress()
	if err != nil {
		apiError := errorhandler.ApiErrorHandle(errorhandler.ErrBadRequest, err.Error())
		c.JSON(apiError.Status, apiError)
		return
	}

	res := ContractAddressResponse{
		Address: *address,
	}

	c.JSON(http.StatusOK, res)
}
