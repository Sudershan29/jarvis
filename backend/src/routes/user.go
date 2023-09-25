
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func UserRoute(g *gin.RouterGroup) {
	g.GET("/profile"   , controllers.UserProfile)
}