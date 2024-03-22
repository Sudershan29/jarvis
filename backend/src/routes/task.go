package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoute(g *gin.RouterGroup) {
	g.GET("", controllers.TaskAll)
	g.POST("", controllers.TaskCreate)
	g.DELETE("/:id", controllers.TaskDelete)
	g.GET(":id/proposals", controllers.TaskListProposals)
	g.POST("/:id/cancel/:proposal_id", controllers.TaskCancelProposal)
}
