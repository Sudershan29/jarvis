package controllers

import (
	"net/http"
	"backend/src/lib"
	"backend/src/models"
	"backend/src/middleware"
  	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) *models.UserModel {
	token 	:= middleware.ExtractToken(c)
	user_id := lib.GetUser(token)
	user, _ := models.UserFind(user_id)
	return user
}

/*

	My Profile

	curl -XGET localhost:8080/api/v1/users/profile -H "Authorization: Bearer <token>"

*/

func UserProfile(c *gin.Context) {
	user := CurrentUser(c)
	c.JSON(http.StatusOK, gin.H{"code": 200, "user": user.Marshal()})
}