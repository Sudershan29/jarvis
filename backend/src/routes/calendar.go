package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func CalendarRoute(g *gin.RouterGroup) {
	g.GET("", controllers.CalendarShowAll)
	g.GET("events", controllers.CalendarEvents)
	g.GET("/:id", controllers.CalendarShow)
}
