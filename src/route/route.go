package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/hiroki-Fukumoto/smart-contract-demo/docs"
	"github.com/hiroki-Fukumoto/smart-contract-demo/domain/contractaddress"
	"github.com/hiroki-Fukumoto/smart-contract-demo/domain/healthcheck"
	"github.com/hiroki-Fukumoto/smart-contract-demo/domain/hello"

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
		c := contractaddress.NewContractAddressController()
		gContractAddress.GET("", c.Get)
	}

	gHello := appApiV1.Group("hello")
	{
		c := hello.NewHelloController()
		gHello.GET("", c.HelloWorld)
	}

	return route
}
