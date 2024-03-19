package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(g *gin.RouterGroup) {
	g.GET("/profile", controllers.UserProfile)
}
