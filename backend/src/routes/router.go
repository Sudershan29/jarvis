package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter(c *gin.Engine) {
	apiGroup := c.Group("/api/v1")

	AuthenticateRoute(apiGroup)
}