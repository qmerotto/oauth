package middleware

import (
	"oauth/web_api/services/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtContent, err := auth.Parser().Parse(c.Request.Header.Get("AUTHORIZATION"))
		if err != nil {
			c.Set("JWTContent", jwtContent)
		}
	}
}
