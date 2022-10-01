package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"oauth/common"
	"strings"
)

type Claims struct {
	jwt.StandardClaims
	UserUUID uuid.UUID
}

type parser struct{}

func Parser() *parser {
	return &parser{}
}

func (p *parser) Parse(bearer string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(strings.TrimPrefix(bearer, "Bearer "), claims, func(t *jwt.Token) (interface{}, error) {
		return common.Settings().RsaPublicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("cannot parse bearer: %s", err.Error())
	}

	return claims, nil
}

type generator struct{}

func Generator() *generator {
	return &generator{}
}

func (g *generator) Generate(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(common.Settings().RsaPrivateKey)
}
