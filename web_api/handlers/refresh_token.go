package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	refresh_token "oauth/common/persistors/refreshToken"
	"oauth/common/persistors/user"
	"oauth/web_api/services/auth"
)

type refreshToken struct {
	base       *basic
	persistors persistor
}

type refreshTokenInput struct {
	refreshToken string
}

type refreshTokenResult struct {
	newAccessToken string
}

func RefreshToken(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			log.Printf("panic %s", err)
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	resultChan := make(chan refreshTokenResult, 1)
	defer close(resultChan)

	err := (&refreshToken{
		base: &basic{ctx: ctx},
		persistors: persistor{
			user:         user.GetPersistor(),
			refreshToken: refresh_token.GetPersistor(),
		}}).Exec(resultChan)
	if err != nil {
		log.Printf("sign up error: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (s *refreshToken) Exec(ch chan refreshTokenResult) error {
	if err := s.base.Read(); err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "read_error",
		})
		return err
	}

	refreshTokenInput := &refreshTokenInput{}
	if err := json.Unmarshal(s.base.body, refreshTokenInput); err != nil {
		log.Printf("signInCredentials unmarshalling error: %s", err.Error())
		s.base.ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unmarshall_error",
		})
		return err
	}
	claims, err := auth.Parser().Parse(refreshTokenInput.refreshToken)
	if err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "token parse error",
		})
		return err
	}
	refreshToken, err := s.persistors.refreshToken.GetByUUID(uuid.MustParse(claims.Id))
	if err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "token parse error",
		})
		return err
	}

	if refreshToken == nil {
		s.base.ctx.JSON(http.StatusForbidden, gin.H{
			"message": "no refresh token",
		})
		return nil
	}
	return nil
}
