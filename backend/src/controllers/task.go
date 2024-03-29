package controllers

import (
	"backend/src/helpers"
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
	var deadline time.Time
	currJWTUser := CurrentUser(c)
	if input.Deadline != "" {
		deadline = helpers.ParseTimeWithZone(input.Deadline, currJWTUser.Timezone())
	}
	task, err := models.TaskCreate(input.Name, input.Description, input.Duration, deadline,
		input.Categories, input.TimePreferences, currJWTUser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "task": task.Marshal()})
}

/*
Always show in descending order of creation time

TODO: Figure out Pagination
*/
func TaskAll(c *gin.Context) {
	tasks, err := models.TaskShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	result := make([]models.TaskJSON, 0)
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

func TaskListProposals(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := models.TaskFind(taskID, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	proposals, err := task.ProposalsWithTimeFilter(helpers.StartOfWeek(time.Now()), helpers.EndOfWeek(time.Now()))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "proposals not loading "})
		return
	}

	result := make([]models.ProposalJSON, 0)
	for _, proposal := range proposals {
		result = append(result, proposal.Marshal())
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "proposals": result})
}

func TaskCancelProposal(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, _ := models.TaskFind(taskID, CurrentUser(c))
	proposalID, err := strconv.Atoi(c.Param("proposal_id"))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	err = task.CancelProposal(proposalID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Proposal cancelled successfully"})
}
