package middleware

import (
	// "time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func SetCors(router *gin.Engine) {
	router.Use(cors.Default())
}