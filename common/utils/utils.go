package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/pbkdf2"
	"io"
	"strconv"
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

func Hash(password string, pepper []byte) (string, error) {
	salt := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		fmt.Print("Err generating random salt")
		return "", fmt.Errorf("Err generating random salt")
	}

	hbts := pbkdf2.Key([]byte(password), salt, 10, 50, sha1.New)
	hbts = pbkdf2.Key(hbts, pepper, 10, 50, sha1.New)

	return fmt.Sprintf("%v:%v:%v",
		10,
		base64.StdEncoding.EncodeToString(salt),
		base64.StdEncoding.EncodeToString(hbts)), nil
}

func Verify(raw, hash string, pepper []byte) (bool, error) {
	hashParts := strings.Split(hash, ":")

	itr, err := strconv.Atoi(hashParts[0])
	if err != nil {
		fmt.Printf("wrong hash %v", hash)
		return false, errors.New("wrong hash, iteration is invalid")
	}
	salt, err := base64.StdEncoding.DecodeString(hashParts[1])
	if err != nil {
		fmt.Print("wrong hash, salt error:", err)
		return false, errors.New("wrong hash, salt error:" + err.Error())
	}

	hsh, err := base64.StdEncoding.DecodeString(hashParts[2])
	if err != nil {
		fmt.Print("wrong hash, hash error:", err)
		return false, errors.New("wrong hash, hash error:" + err.Error())
	}

	rhash := pbkdf2.Key([]byte(raw), salt, itr, len(hsh), sha1.New)
	rhash = pbkdf2.Key(rhash, pepper, 10, 50, sha1.New)

	return bytes.Equal(rhash, hsh), nil
}
