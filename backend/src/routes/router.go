package routes

import (
	"backend/src/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRouter(c *gin.Engine) {
	// middleware.SetCors(c)

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

	// Goal
	goalGroup := c.Group("/api/v1/goals")
	goalGroup.Use(middleware.JwtAuthMiddleware())
	GoalRoute(goalGroup)

	// Task
	taskGroup := c.Group("/api/v1/tasks")
	taskGroup.Use(middleware.JwtAuthMiddleware())
	TaskRoute(taskGroup)

	// Meeting
	meetingGroup := c.Group("/api/v1/meetings")
	meetingGroup.Use(middleware.JwtAuthMiddleware())
	MeetingRoute(meetingGroup)

	// Hobby
	hobbyGroup := c.Group("/api/v1/hobbies")
	hobbyGroup.Use(middleware.JwtAuthMiddleware())
	HobbyRoute(hobbyGroup)

	// Calendar
	calendarGroup := c.Group("/api/v1/calendars")
	calendarGroup.Use(middleware.JwtAuthMiddleware())
	CalendarRoute(calendarGroup)
}
