package routes

import (
	"backend/src/controllers"

	"github.com/gin-gonic/gin"
)

func SkillRoute(g *gin.RouterGroup) {
	g.GET("", controllers.SkillAll)
	g.POST("", controllers.SkillCreate)
	g.DELETE("/:id", controllers.SkillDelete)
	g.GET(":id/proposals", controllers.SkillListProposals)
	g.POST("/:id/cancel/:proposal_id", controllers.SkillCancelProposal)
}
