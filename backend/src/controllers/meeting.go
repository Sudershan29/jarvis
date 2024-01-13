package controllers

import (
	"time"
	"net/http"
	"backend/src/models"
  	"github.com/gin-gonic/gin"
)

/*

	Create Meeting

	curl -XPOST localhost:8080/api/v1/meetings -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createMeetingInput struct {
	Name  	    string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Where  	    string   `json:"where" binding:"required"`
	Whom  	    string   `json:"whom" binding:"required"`
	Duration  	int      `json:"duration" binding:"required"`
	When		string   `json:"when" binding:"required"`
}

func MeetingCreate(c *gin.Context) {
	var input createMeetingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	meeting, err := models.MeetingCreate(input.Name, input.Description, input.Where, input.Whom, 
		input.Duration, time.Now(), CurrentUser(c))  // TODO: Change deadline
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "meeting": meeting.Marshal() })
}


func MeetingAll(c *gin.Context) {
	meetings, err := models.MeetingShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	var result []models.MeetingJSON
	for _, meeting := range meetings {
		result = append(result, meeting.Marshal())
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "meetings": result })
}