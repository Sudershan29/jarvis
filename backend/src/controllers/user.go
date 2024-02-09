package controllers

import (
	"log"
	"net/http"
	"backend/src/lib"
	"backend/src/models"
	"backend/src/middleware"
  	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) *models.JwtUser {
	token 	:= middleware.ExtractToken(c)
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
	Get your calendar information
*/

func UserCalendar(c *gin.Context) {
	user 	  := CurrentUser(c)
	// code, err := lib.GetSavedCalendarToken(user.UserId.String())
	userCalendarClient, err := lib.NewCalendarClient(user.UserId.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}

	calendarEvents, err := userCalendarClient.FetchEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "events": calendarEvents })
}

/*
	Connect with your Google Calendar
*/

func UserCalendarConnect(c *gin.Context) {
	user := CurrentUser(c)
	redirectUrl, err := lib.GenerateGCalendarAuthorizationLink(user.UserId.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}
	log.Println("Redirecting to ", redirectUrl)
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}