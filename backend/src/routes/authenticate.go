package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticateRoute(g *gin.RouterGroup) {
	g.GET("/status", controllers.HealthCheck)
	g.POST("/login", controllers.AuthenticateLogin)
	g.POST("/register", controllers.AuthenticateRegister)
	g.DELETE("/logout", controllers.AuthenticateLogout)

	/*  Google signin  */
	g.GET("/google/signin", controllers.AuthenticateGoogleLogin)
	g.GET("/google/callback", controllers.AuthenticateGoogleCallback)
	g.GET("/calendar/connect", controllers.UserCalendarConnect)
}
