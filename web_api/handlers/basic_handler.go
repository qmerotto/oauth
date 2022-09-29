package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oauth/web_api/services/auth"
	"strings"
)

type basic struct {
	ctx *gin.Context
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
