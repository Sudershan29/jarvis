package controllers

import (
	"net/http"
	"backend/src/models"
  	"github.com/gin-gonic/gin"
)

/*

	Create Goal

	curl -XPOST localhost:8080/api/v1/goals -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createGoalInput struct {
	Name  	    string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Categories  []string `json:"categories"`
}

func GoalCreate(c *gin.Context) {
	var input createGoalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	goal, err := models.GoalCreate(input.Name, input.Description, input.Categories, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "goal": goal.Marshal() })
}


func GoalAll(c *gin.Context) {
	goals, err := models.GoalShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	var result []models.GoalJSON
	for _, goal := range goals {
		result = append(result, goal.Marshal())
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "goals": result })
}