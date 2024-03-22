package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func ScheduleRoute(g *gin.RouterGroup) {
	g.POST("/week", controllers.ScheduleMyWeek)
}
