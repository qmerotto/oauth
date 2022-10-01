package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"oauth/common/database/models"
	"oauth/common/persistors/user"
)

type signUp struct {
	base      *basic
	persistor user.Persistor
}

type signUpCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SignUp(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			log.Printf("panic %s", err)
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	err := (&signUp{base: &basic{ctx: ctx}, persistor: user.GetPersistor()}).Exec()
	if err != nil {
		log.Printf("sign up error: %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	log.Printf("success")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (s *signUp) Exec() error {
	if err := s.base.Read(); err != nil {
		return err
	}

	credentials := &signUpCredentials{}
	if err := json.Unmarshal(s.base.body, credentials); err != nil {
		log.Printf("signUpCredentials unmarshalling error: %s", err.Error())
	}

	if credentials == nil {
		s.base.ctx.Status(http.StatusForbidden)
	}
	fmt.Printf(fmt.Sprintf("username: %s password: %s", credentials.Username, credentials.Password))

	err := s.persistor.Create(&models.User{
		UUID:     uuid.New(),
		Name:     credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	})
	if err != nil {
		log.Printf("user creation error: %s", err.Error())
		return err
	}

	return nil
}
