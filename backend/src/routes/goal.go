
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func GoalRoute(g *gin.RouterGroup) {
	g.GET(""   , controllers.GoalAll)
	g.POST(""   , controllers.GoalCreate)
}