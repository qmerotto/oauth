package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"oauth/common/database/models"
	refresh_token "oauth/common/persistors/refreshToken"
	"oauth/common/persistors/user"
	"oauth/web_api/services/auth"
	"time"
)

type signIn struct {
	base       *basic
	persistors persistor
}

type persistor struct {
	user         user.Persistor
	refreshToken refresh_token.Persistor
}

type signInCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func SignIn(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			fmt.Printf("recovering error: %s", err)
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	resultChan := make(chan interface{}, 1)
	defer close(resultChan)

	err := (&signIn{
		base: &basic{ctx: ctx},
		persistors: persistor{
			user:         user.GetPersistor(),
			refreshToken: refresh_token.GetPersistor(),
		}}).Exec(resultChan)
	if err != nil {
		log.Printf("sign up error: %v", err)
		return
	}

	result := <-resultChan
	ctx.JSON(http.StatusOK, result)
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

	user, err := s.persistors.user.GetUserByMail(credentials.Email)
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

	accessToken, err := auth.Generator().Generate(
		&auth.Claims{
			StandardClaims: jwt.StandardClaims{
				Issuer:    "oauth",
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(15 * time.Hour).Unix(),
			},
			UserUUID: user.UUID,
		},
	)
	if err != nil {
		log.Printf("auth generation error: %s", err.Error())
		s.base.ctx.Status(http.StatusInternalServerError)
		return err
	}

	refreshTokenUUID := uuid.New()
	err = s.persistors.refreshToken.Create(&models.RefreshToken{
		UUID:     refreshTokenUUID,
		UserUUID: user.UUID,
	})

	refreshToken, err := auth.Generator().Generate(
		&auth.Claims{
			StandardClaims: jwt.StandardClaims{
				Id:        refreshTokenUUID.String(),
				Issuer:    "oauth",
				IssuedAt:  time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},
		},
	)

	if err != nil {
		fmt.Printf("error when saving refresh token: %s", err.Error())
		return err
	}

	ch <- signInResult{AccessToken: accessToken, RefreshToken: refreshToken}

	return nil
}
