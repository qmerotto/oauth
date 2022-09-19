package web_api

import (
	"oauth/web_api/handlers"

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
	{
		v1.POST("/ticket", func(c *gin.Context) { handlers.Ticket(c) })
	}

	return r
}
