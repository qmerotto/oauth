package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"oauth/web_api/services/auth"

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

	token, err := auth.Generator().Generate(&auth.Claims{UserUUID: uuid.New()})
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	ctx.Status(http.StatusAccepted)
}
