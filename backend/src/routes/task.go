
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func TaskRoute(g *gin.RouterGroup) {
	g.GET(""   , controllers.TaskAll)
	g.POST(""   , controllers.TaskCreate)
}