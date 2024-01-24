
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func UserRoute(g *gin.RouterGroup) {
	g.GET("/profile" , controllers.UserProfile)
	g.GET("/calendar", controllers.UserCalendar)
	g.GET("/calendar/connect", controllers.UserCalendarConnect)
}