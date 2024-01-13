package controllers

import (
	"net/http"
	"backend/src/models"
  	"github.com/gin-gonic/gin"
)

/*

	Create Hobby

	curl -XPOST localhost:8080/api/v1/hobbys -H "Authorization: Bearer <token>" -d '{"name": "Shuffle Cards", "level":"intermediate", "duration": 5, "categories": \["hmm"\]}'

*/

type createHobbyInput struct {
	Name  	    string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Categories  []string `json:"categories"`
}

func HobbyCreate(c *gin.Context) {
	var input createHobbyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hobby, err := models.HobbyCreate(input.Name, input.Description, input.Categories, CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "hobby": hobby.Marshal() })
}


func HobbyAll(c *gin.Context) {
	hobbies, err := models.HobbyShowAll(CurrentUser(c))
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	var result []models.HobbyJSON
	for _, hobby := range hobbies {
		result = append(result, hobby.Marshal())
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "hobbies": result })
}