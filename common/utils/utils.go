package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
)

func Base64Decode(src []byte) (b []byte, err error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	decodedPublicKey, err := Base64Decode([]byte(strings.ReplaceAll(publicKey, "\\n", "\n")))
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error while decoding public key: %s", err.Error()))
	}

	decodedPublicKey = []byte(strings.ReplaceAll(string(decodedPublicKey), "\\n", "\n"))
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error while parsing RSA key: %s", err.Error()))
	}

	return rsaPublicKey, nil
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	decodedPrivateKey, err := Base64Decode([]byte(strings.ReplaceAll(privateKey, "\\n", "\n")))
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error while decoding public key: %s", err.Error()))
	}
	decodedPrivateKey = []byte(strings.ReplaceAll(string(decodedPrivateKey), "\\n", "\n"))
	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Error while parsing RSA key: %s", err.Error()))
	}

	return rsaPrivateKey, nil
}
