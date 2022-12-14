package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/hiroki-Fukumoto/geth-dapp-demo/docs"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/domain/contractaddress"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/domain/healthcheck"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/domain/hello"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/domain/product"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			// "http://localhost:3000",
			"http://localhost:8080",
		},
		AllowMethods: []string{
			"*",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	appApiGroup := route.Group("/api")
	appApiV1 := appApiGroup.Group("/v1")

	gHealthCheck := appApiGroup.Group("health-check")
	{
		c := healthcheck.NewHealthCheckController()
		gHealthCheck.GET("", c.HealthCheck)
	}

	gContractAddress := appApiV1.Group("contract-address")
	{
		s := contractaddress.NewContractAddressService()
		c := contractaddress.NewContractAddressController(s)
		gContractAddress.GET("", c.Get)
	}

	gHello := appApiV1.Group("hello")
	{
		c := hello.NewHelloController()
		gHello.GET("", c.HelloWorld)
	}

	gProduct := appApiV1.Group("products")
	{
		s := product.NewProductService()
		c := product.NewProductController(s)
		gProduct.POST("", c.Create)
	}

	return route
}
