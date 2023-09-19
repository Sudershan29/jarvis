package middleware

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/src/lib"
)

// Extracts the token headers
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}


// Verifies if the token is valid
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !lib.IsTokenValid(ExtractToken(c)) {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
