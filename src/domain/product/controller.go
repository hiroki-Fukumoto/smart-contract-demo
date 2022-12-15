package product

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/errorhandler"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/validator"
)

type ProductController interface {
	Create(c *gin.Context)
}

type productController struct {
	productService ProductService
}

func NewProductController(s ProductService) ProductController {
	return &productController{
		productService: s,
	}
}

// @Summary Create product.
// @Description Create NFT product.
// @Tags product
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "New product info"
// @Success 204 {object} nil
// @Failure 400 {object} errorhandler.ErrorResponse
// @Failure 422 {object} errorhandler.ErrorResponse
// @Router /api/v1/products [post]
func (cn productController) Create(c *gin.Context) {
	var req CreateProductRequest
	c.BindJSON(&req)
	if err := validator.Validate(&req); err != nil {
		errors := validator.GetErrorMessages(err)
		apiError := errorhandler.ApiErrorHandle(errorhandler.ErrUnprocessableEntity, strings.Join(errors, ","))
		c.JSON(apiError.Status, apiError)
		return
	}

	err := cn.productService.Create(req)
	if err != nil {
		apiError := errorhandler.ApiErrorHandle(errorhandler.ErrBadRequest, err.Error())
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
