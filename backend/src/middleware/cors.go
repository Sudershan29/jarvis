package middleware

import (
	// "time"
    // "fmt"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)


// func SetCors(router *gin.Engine) {
// 	router.Use(cors.New(cors.Config{
// 		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
// 		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		AllowAllOrigins:  true,
// 		MaxAge:           86400,
// 	}))
// }

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.IndentedJSON(204, "")
            return
        }

        c.Next()
    }
}