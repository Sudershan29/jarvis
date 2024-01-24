package controllers

import (
	"net/http"
	"backend/src/lib"
	"backend/src/models"
  	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

/*

	Login

	curl -XPOST localhost:8080/api/v1/login -d '{"email": "abc@google.com", "password":"test"}'

*/
type authenticateLoginInput struct {
	Email string `json:"email" binding:"required"`
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
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": err.Error() })
		return
	}
	token, err := user.Login(input.Password)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"code": http.StatusNotAcceptable, "message": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "token": token })
}

/*

	Register

	curl -XPOST localhost:8080/api/v1/register -d '{"email": "abc@google.com", "username":"abc","password":"test"}'

*/

type authenticateRegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email 	 string `json:"email" binding:"required"`
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
	c.JSON(http.StatusOK, gin.H{"code": 200 })
}

/*

	Google Sign in

*/
func AuthenticateGoogleLogin(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", "google")
	c.Request.URL.RawQuery = q.Encode()
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
		userJSON := models.ConvertToJSON(user)

		// TODO: Generate JWT, maybe explore options into redirect to different URL in backend to isolate logic
		c.JSON(http.StatusOK, gin.H{"code": 200, "user": userJSON})
	} else {
		/*
			Google Calendar Callback

			NOTES: Add calendars to database, and enable different views for different accounts | Research how to identify which email id? Maybe add it to Calendar State
		*/
		status := lib.SaveCalendarToken(calendarState, c.Query("code"))
		c.JSON(http.StatusOK, gin.H{"code": 200, "user": calendarState.UserSlug, "saved": status })
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