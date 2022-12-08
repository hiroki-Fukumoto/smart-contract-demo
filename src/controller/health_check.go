package controller

import "github.com/gin-gonic/gin"

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type healthCheckController struct {
}

func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

type healthCheckResponse struct {
	Message string `json:"message"`
}

// @Summary Health Check
// @Description Check the server communication.
// @Tags healthCheck
// @Accept json
// @Produce json
// @Success 200 {object} healthCheckResponse
// @Router /api/health-check [get]
func (h healthCheckController) HealthCheck(c *gin.Context) {
	res := healthCheckResponse{
		Message: "OK",
	}
	c.JSON(200, res)
}
