package middleware

import (
	"log"
	"net/http"
	"oauth/web_api/services/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtContent, err := auth.Parser().Parse(c.Request.Header.Get("AUTHORIZATION"))

		if err == nil && jwtContent != nil {
			c.Set("JWTContent", jwtContent)
			return
		}
		log.Printf(err.Error())
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
	}
}
