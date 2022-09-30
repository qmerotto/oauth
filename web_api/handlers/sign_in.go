package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"oauth/common/database/models"
	"oauth/web_api/services/auth"
	"time"
)

type signIn struct{}

type signInCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	basic := &basic{ctx: ctx}
	if err := basic.Read(); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	credentials := &signInCredentials{}
	if err := json.Unmarshal(basic.body, credentials); err != nil {
		log.Printf("signInCredentials unmarshalling error: %s", err.Error())
	}

	if credentials == nil {
		ctx.Status(http.StatusForbidden)
	}
	fmt.Printf(fmt.Sprintf("username: %s password: %s", credentials.Username, credentials.Password))

	user := &models.User{}

	token, err := auth.Generator().Generate(&auth.Claims{UserUUID: user.UUID, ExpiredAt: time.Now()})
	if err != nil {
		log.Printf("auth generation error: %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("Bearer %s", token),
	})
}
