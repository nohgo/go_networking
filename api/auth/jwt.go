package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var key string = os.Getenv("GO_NETWORKING_KEY")

func CreateToken(name string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
	})
	tokenString, err = token.SignedString([]byte(key))
	return
}

// middleware should be used to parse token
func parseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
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
