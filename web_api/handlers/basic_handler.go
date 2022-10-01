package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"oauth/web_api/services/auth"
	"strings"
)

type basic struct {
	ctx  *gin.Context
	body []byte
}

type Error struct {
	message string
	code    string
}

func (b *basic) Read() error {
	body, err := io.ReadAll(b.ctx.Request.Body)
	if err != nil {
		log.Printf("read body error: %s", err.Error())
		return err
	}

	b.body = body
	return nil
}

func (b *basic) Credentials() error {

	return nil
}

func (b *basic) JWT() error {
	return nil
}

func (b *basic) BasicAuth() error {
	basicAuth := strings.TrimPrefix(b.ctx.GetHeader("Authorization"), "Basic ")
	fmt.Printf(fmt.Sprintf("Basic auth: %s\n", basicAuth))

	splitBasicAuth := strings.Split(basicAuth, ":")
	if len(splitBasicAuth) != 2 {
		return fmt.Errorf("invalid basic auth")
	}

	return (&auth.Basic{
		Username: splitBasicAuth[0],
		Password: splitBasicAuth[1],
	}).Exec()
}
