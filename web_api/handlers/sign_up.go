package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"oauth/common"
	"oauth/common/database/models"
	"oauth/common/persistors/user"
	"oauth/common/utils"
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

type signUpResult struct {
	Status string `json:"status"`
}

func SignUp(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if err := recover(); err != nil {
			log.Printf("panic %s", err)
			ctx.Status(http.StatusInternalServerError)
		}
	}(ctx)

	result := &signUpResult{}
	err := (&signUp{base: &basic{ctx: ctx}, persistor: user.GetPersistor()}).Exec(result)
	if err != nil {
		log.Printf("sign up error: %v", err)
		return
	}

	log.Printf("success")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (s *signUp) Exec(result *signUpResult) error {
	if err := s.base.Read(); err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "read_error",
		})
		return err
	}

	credentials := &signUpCredentials{}
	if err := json.Unmarshal(s.base.body, credentials); err != nil {
		log.Printf("signUpCredentials unmarshalling error: %s", err.Error())
		s.base.ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unmarshall_error",
		})
		return err
	}

	if credentials == nil || !credentials.isValid() {
		s.base.ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "credentials_error",
		})
		return fmt.Errorf("invalid credentials")
	}

	if err := credentials.HashPassword(); err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "credentials_error",
		})
		return fmt.Errorf("password hashing error")
	}
	fmt.Printf(fmt.Sprintf("username: %s password: %s", credentials.Username, credentials.Password))

	err := s.persistor.Create(&models.User{
		UUID:     uuid.New(),
		Name:     credentials.Username,
		Password: credentials.Password,
		Email:    credentials.Email,
	})
	if err != nil {
		s.base.ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "user_creation_error",
		})
		log.Printf("user creation error: %s", err.Error())
		return err
	}

	*result = signUpResult{Status: "OK"}
	return nil
}

func (s *signUpCredentials) isValid() bool {
	if len(s.Password) < 8 {
		return false
	}

	if len(s.Email) == 0 {
		return false
	}

	return true
}

func (s *signUpCredentials) HashPassword() error {
	password, err := utils.Hash(s.Password, common.Settings().Pepper)
	if err != nil {
		return fmt.Errorf("password too short")
	}

	s.Password = password
	return nil
}
