package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
)


/* 

	Health Check

*/

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200 })
}