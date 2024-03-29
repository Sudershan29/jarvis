package controllers

import (
	"backend/src/helpers"
	"backend/src/lib"
	"backend/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
)

/*
Login

curl -XPOST localhost:8080/api/v1/login -d '{"email": "abc@google.com", "password":"test"}'
*/
type authenticateLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthenticateLogin(c *gin.Context) {
	var input authenticateLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.UserFind(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": err.Error()})
		return
	}
	token, err := user.Login(input.Password, false)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"code": http.StatusNotAcceptable, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "token": token})
}

/*

	Register

	curl -XPOST localhost:8080/api/v1/register -d '{"email": "abc@google.com", "username":"abc","password":"test"}'

*/

type authenticateRegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func AuthenticateRegister(c *gin.Context) {
	var input authenticateRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.UserCreate(input.Username, input.Password, input.Email)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "user": user.Marshal()})
}

/*

	Logout


*/

type AuthenticateLogoutInput struct {
	// Username string `json:"username" binding:"required"`
	// Password string `json:"password" binding:"required"`
}

func AuthenticateLogout(c *gin.Context) {
	var input AuthenticateLogoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200})
}

/*
Google Sign in
*/
func AuthenticateGoogleLogin(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	fmt.Println(c.Writer.Header())
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

/*

	Google Callback

*/

// TODO: Change the callback for multiple services?
func AuthenticateGoogleCallback(c *gin.Context) {
	q := c.Request.URL.Query()

	// Attempting to check if it was redirect request from calendar
	if calendarState, err := lib.ConvertCalendarState(c.Query("state")); err != nil {
		/*
			OAuth2.0 Callback
		*/
		q.Add("provider", "google")
		c.Request.URL.RawQuery = q.Encode()
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// NOTE: maybe explore options into redirect to different URL in backend to isolate logic
		googleUserJSON := models.ConvertToJSON(user)
		userModel, err := models.UserFind(googleUserJSON.Email)
		if err != nil {
			userModel, err = models.UserCreate(googleUserJSON.Name, uuid.New().String(), googleUserJSON.Email) // NOTE: Autogenerating passwords for privacy reasons
			if err != nil {
				c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
				return
			}
		}

		token, err := userModel.Login("", true)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}

		// c.JSON(http.StatusOK, gin.H{"code": 200, "user": googleUserJSON, "token": token})
		c.Redirect(http.StatusFound, helpers.GetEnv("FRONTEND_URL")+"/googleLoginSuccess/?token="+token)
	} else {
		/*
			Google Calendar Callback

			NOTES: Add calendars to database, and enable different views for different accounts | Research how to identify which email id? Maybe add it to Calendar State
		*/

		lib.SaveCalendarToken(calendarState, c.Query("code"))
		models.CalendarFindOrCreate("Google Calendar", "google", c.Query("code"), models.NewJwtUser(calendarState.UserSlug))
		// c.JSON(http.StatusOK, gin.H{"code": 200, "user": calendarState.UserSlug, "saved": status})
		c.Redirect(http.StatusFound, helpers.GetEnv("FRONTEND_URL")+"/Profile")
	}
}

/*

func AuthenticateGoogleStatus(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userJSON := models.ConvertToJSON(user)
	c.JSON(http.StatusOK, gin.H{"code": 200, "user": userJSON})
}

*/
