package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"oauth/common/persistors/user"
	"oauth/web_api/services/auth"
	"time"
)

type signIn struct {
	base      *basic
	persistor user.Persistor
}

type signInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	resultChan := make(chan interface{}, 1)
	defer close(resultChan)

	err := (&signIn{base: &basic{ctx: ctx}, persistor: user.GetPersistor()}).Exec(resultChan)
	if err != nil {
		log.Printf("sign up error: %v", err)
		return
	}

	token := <-resultChan
	ctx.JSON(http.StatusOK, gin.H{
		"token": fmt.Sprintf("Bearer %s", token.(string)),
	})
}

func (s *signIn) Exec(ch chan interface{}) error {
	if err := s.base.Read(); err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "read_error",
		})
		return err
	}

	credentials := &signInCredentials{}
	if err := json.Unmarshal(s.base.body, credentials); err != nil {
		log.Printf("signInCredentials unmarshalling error: %s", err.Error())
		s.base.ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unmarshall_error",
		})
		return err
	}

	if credentials == nil {
		s.base.ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "credentials_error",
		})
		return fmt.Errorf("invalid credentials")
	}

	user, err := s.persistor.GetUserByMail(credentials.Email)
	if err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "read_error",
		})
		return err
	}

	if user == nil {
		s.base.ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user_doesnt_exist",
		})
		return err
	}

	token, err := auth.Generator().Generate(
		&auth.Claims{
			StandardClaims: jwt.StandardClaims{
				Issuer:    "oauth",
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},
			UserUUID: user.UUID,
		},
	)
	if err != nil {
		log.Printf("auth generation error: %s", err.Error())
		s.base.ctx.Status(http.StatusInternalServerError)
		return err
	}

	ch <- token
	return nil
}
