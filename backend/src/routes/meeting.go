
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/src/controllers"
)

func MeetingRoute(g *gin.RouterGroup) {
	g.GET(""   , controllers.MeetingAll)
	g.POST(""   , controllers.MeetingCreate)
}