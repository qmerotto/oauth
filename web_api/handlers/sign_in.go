package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	body := ctx.Request.Body

	_, err := io.ReadAll(body)
	if err != nil {
		log.Printf("read body error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusAccepted)
}
