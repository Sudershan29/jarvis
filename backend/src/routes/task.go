package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoute(g *gin.RouterGroup) {
	g.GET("", controllers.TaskAll)
	g.POST("", controllers.TaskCreate)
	g.DELETE("/:id", controllers.TaskDelete)
}
