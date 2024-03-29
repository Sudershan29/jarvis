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

	Create Skill

	curl -XPOST localhost:8080/api/v1/skills -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createSkillInput struct {
	Name           string   `json:"name" binding:"required"`
	Level          string   `json:"level"`
	Duration       int      `json:"duration"`
	Categories     []string `json:"categories"`
	TimePreference []string `json:"timepreference"`
}

func SkillCreate(c *gin.Context) {
	var input createSkillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	skill, err := models.SkillCreate(input.Name, input.Level, input.Duration, input.Categories, input.TimePreference, CurrentUser(c))
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
	result := make([]models.SkillJSON, 0)
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

func SkillListProposals(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	skill, err := models.SkillFind(skillID, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	proposals, err := skill.ProposalsWithTimeFilter(helpers.StartOfWeek(time.Now()), helpers.EndOfWeek(time.Now()))
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

func SkillCancelProposal(c *gin.Context) {
	skillID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	skill, err := models.SkillFind(skillID, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find skill with given ID"})
		return
	}

	proposalID, err := strconv.Atoi(c.Param("proposal_id"))

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	err = skill.CancelProposal(proposalID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Proposal cancelled successfully"})
}

// func SkillUpdate(c *gin.Context) {
// 	skillID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
// 		return
// 	}

// 	var input createSkillInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if any field is provided for update, if not return an error
// 	if input.Name == "" && input.Level == "" && input.Duration == 0 && len(input.Categories) == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No update parameters provided"})
// 		return
// 	}

// 	skill, err := models.SkillUpdate(skillID, input.Name, input.Level, input.Duration, input.Categories, CurrentUser(c))
// 	if err != nil {
// 		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"code": 200, "skill": skill.Marshal()})
// }
