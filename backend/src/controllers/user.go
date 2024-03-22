package controllers

import (
	"backend/src/lib"
	"backend/src/middleware"
	"backend/src/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) *models.JwtUser {
	token := middleware.ExtractToken(c)
	user_id := lib.GetUser(token)
	return models.NewJwtUser(user_id)
}

/*

	My Profile

	curl -XGET localhost:8080/api/v1/users/profile -H "Authorization: Bearer <token>"

*/

func UserProfile(c *gin.Context) {
	user := CurrentUser(c)
	user.Load()
	c.JSON(http.StatusOK, gin.H{"code": 200, "user": user.Model.Marshal()})
}

/*
	Connect with your Google Calendar
*/

func UserCalendarConnect(c *gin.Context) {
	token := c.Query("token")
	user_id := lib.GetUser(token)
	user := models.NewJwtUser(user_id)

	redirectUrl, err := lib.GenerateGCalendarAuthorizationLink(user.UserId.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Redirecting to ", redirectUrl)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}
