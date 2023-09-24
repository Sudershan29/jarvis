
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func AuthenticateRoute(g *gin.RouterGroup) {
	g.GET("/status"   ,	controllers.HealthCheck)
	g.POST("/login"   , controllers.AuthenticateLogin)
	g.POST("/register", controllers.AuthenticateRegister)
	g.DELETE("/logout", controllers.AuthenticateLogout)
}