package routes

import (
	"backend/src/middleware"
	"github.com/gin-gonic/gin"
)

func CreateRouter(c *gin.Engine) {
	middleware.SetCors(c)

	// Auth
	apiGroup := c.Group("/api/v1")
	AuthenticateRoute(apiGroup)

	// User
	userGroup := c.Group("/api/v1/users")
	userGroup.Use(middleware.JwtAuthMiddleware())
	UserRoute(userGroup)

	// Skill
	skillGroup := c.Group("/api/v1/skills")
	skillGroup.Use(middleware.JwtAuthMiddleware())
	SkillRoute(skillGroup)
}