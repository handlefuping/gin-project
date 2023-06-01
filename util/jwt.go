package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKEY = []byte("AllYourBase")

type MyCustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateTokenStr (issuer string, expires time.Duration) (string, error) {
	claims := MyCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKEY)



}

func ParseTokenStr (tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.RegisteredClaims.Issuer, nil
	} else {
		return "", err
	}
}
