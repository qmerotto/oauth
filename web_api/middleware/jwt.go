package middleware

import (
	"oauth/web_api/services/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		jwtContent := jwt.Parser().Parse(c.Request.Header.Get("AUTHORIZATION"))
		c.Set("JWTContent", jwtContent)
	})
}
