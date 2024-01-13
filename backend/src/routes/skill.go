
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func SkillRoute(g *gin.RouterGroup) {
	g.GET(""   , controllers.SkillAll)
	g.POST(""   , controllers.SkillCreate)
}