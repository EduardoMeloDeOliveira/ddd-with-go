package http

import "github.com/gin-gonic/gin"

type HealthcheckController struct {
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (c *HealthcheckController) RegisterRoute(r *gin.Engine) {
	healthRoutes := r.Group("/healthcheck")
	{
		healthRoutes.GET("", c.Healthcheck)
	}
}

func (c *HealthcheckController) Healthcheck(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "All ok"})
}
