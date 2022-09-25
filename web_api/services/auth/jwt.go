package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"oauth/common"
	"time"
)

type Claims struct {
	jwt.Claims
	UserUUID  uuid.UUID
	ExpiredAt time.Time
}

type parser struct{}

func Parser() *parser {
	return &parser{}
}

func (p *parser) Parse(bearer string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(bearer, claims, func(t *jwt.Token) (interface{}, error) {
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

	return token.SignedString([]byte("my_private_key"))
}
