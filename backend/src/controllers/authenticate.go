package controllers

import (
	"net/http"
	"backend/src/models"
  	"github.com/gin-gonic/gin"
)

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