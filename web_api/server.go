package web_api

import (
	"github.com/gin-gonic/gin"
	"oauth/web_api/handlers"
	"oauth/web_api/middleware"
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

	v3.GET("/test", func(c *gin.Context) { handlers.TRR(c) })

	return r
}
