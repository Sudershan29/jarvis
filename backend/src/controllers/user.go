package controllers

import (
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