package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/golang-jwt/jwt"
)

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

func ExtractClaim(tokenstring string) (interface{}, bool) {
	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) ([]byte, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, false
	}

	return token.Claims, true
}
