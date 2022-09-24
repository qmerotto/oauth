package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Content struct {
	jwt.Claims
	UserUUID uuid.UUID
}

type parser struct{}

func Parser() *parser {
	return &parser{}
}

func (p *parser) Parse(token string) Content {
	jwt.ParseWithClaims(token, &Content{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("my_key"), nil
	})
	return Content{
		UserUUID: uuid.New(),
	}
}

type generator struct{}

func Generator() *generator {
	return &generator{}
}

func (g *generator) Generate(content Content) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	token.Claims = content

	return token.SignedString([]byte("my_key"))
}
