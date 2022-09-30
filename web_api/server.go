package web_api

import (
	"oauth/web_api/handlers"
	"oauth/web_api/middleware"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	if err := getRouter().Run(":8081"); err != nil {
		panic(err)
	}
}

func getRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	v1 := r.Group("/web_api")
	v2 := v1.Group("/auth")
	v2.POST("/sign_in", func(c *gin.Context) { handlers.SignIn(c) })
	v2.POST("/sign_up", func(c *gin.Context) { handlers.SignUp(c) })

	v3 := v1.Group("/authorization")
	v3.Use(middleware.AuthMiddleware())

	return r
}
