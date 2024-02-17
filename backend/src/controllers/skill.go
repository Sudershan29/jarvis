package controllers

import (
	"backend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*

	Create Skill

	curl -XPOST localhost:8080/api/v1/skills -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createSkillInput struct {
	Name       string   `json:"name" binding:"required"`
	Level      string   `json:"level"`
	Duration   int      `json:"duration"`
	Categories []string `json:"categories"`
}

func SkillCreate(c *gin.Context) {
	var input createSkillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	skill, err := models.SkillCreate(input.Name, input.Level, input.Duration, input.Categories, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "skill": skill.Marshal()})
}

func SkillAll(c *gin.Context) {
	skills, err := models.SkillShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	var result []models.SkillJSON
	for _, skill := range skills {
		result = append(result, skill.Marshal())
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "skills": result})
}

func SkillDelete(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}
	err = models.SkillDelete(skillID, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Skill deleted successfully"})
}
