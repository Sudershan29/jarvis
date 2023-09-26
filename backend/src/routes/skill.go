
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func SkillRoute(g *gin.RouterGroup) {
	g.POST("/"   , controllers.SkillCreate)
	g.GET("/"   , controllers.SkillAll)
}