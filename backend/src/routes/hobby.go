
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func HobbyRoute(g *gin.RouterGroup) {
	g.GET(""   , controllers.HobbyAll)
	g.POST(""   , controllers.HobbyCreate)
}