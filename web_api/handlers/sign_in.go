package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"oauth/web_api/services/auth"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	basic := basic{
		ctx: ctx,
	}
	basic.BasicAuth()
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("read body error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	credentials := &Credentials{}
	if err = json.Unmarshal(body, credentials); err != nil {
		log.Printf("credential unmarshalling error: %s", err.Error())
	}

	if credentials == nil {
		ctx.Status(http.StatusForbidden)
	}
	fmt.Printf(fmt.Sprintf("username: %s password: %s", credentials.Username, credentials.Password))

	token, err := auth.Generator().Generate(&auth.Claims{UserUUID: uuid.New()})
	if err != nil {
		log.Printf("auth generation error: %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("Bearer %s", token),
	})
}
