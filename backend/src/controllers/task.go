package controllers

import (
	"backend/src/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*

	Create Task

	curl -XPOST localhost:8080/api/v1/tasks -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createTaskInput struct {
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	Duration        int      `json:"duration"`
	Deadline        string   `json:"deadline"`
	Categories      []string `json:"categories"`
	TimePreferences []string `json:"timepreference"`
}

func TaskCreate(c *gin.Context) {
	var input createTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, err := models.TaskCreate(input.Name, input.Description, input.Duration, time.Now(), // TODO: Change deadline
		input.Categories, input.TimePreferences, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "task": task.Marshal()})
}

func TaskAll(c *gin.Context) {
	tasks, err := models.TaskShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	var result []models.TaskJSON
	for _, task := range tasks {
		result = append(result, task.Marshal())
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "tasks": result})
}

func TaskDelete(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = models.TaskDelete(taskID, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Task deleted successfully"})
}
