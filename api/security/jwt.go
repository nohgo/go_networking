package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken() (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "noh",
	})
	tokenString, err = token.SignedString([]byte("hello"))
	return
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("hello"), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	switch claims["name"].(type) {
	case string:
		return claims["name"].(string), err
	default:
		return "", errors.New("invalid claims")
	}
}
