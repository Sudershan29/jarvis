package controllers

import (
	"backend/src/algorithm/planner"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

	curl -XGET localhost:8080/api/v1/users/profile -H "Authorization: Bearer <token>"

*/

// TODO: Making the ScheduleMyWeek function async, and setting up some communication method
// NOTE: Does half week for now
func ScheduleMyWeek(c *gin.Context) {
	user := CurrentUser(c)

	// NOTE: Make this async
	err := planner.PrepareTimeTable(user.UserId.String(), 3)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "status": true})
}
