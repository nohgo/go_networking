// The auth package contains all the methods required for authentication.
package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key string = os.Getenv("GO_NETWORKING_KEY")

// Wrapper for the values contained in a JWT
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Returns a JWT token that has the provided username inside
func CreateToken(name string) (tokenString string, err error) {
	claims := JWTClaims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(key))
	return
}

// Parses the name and checks if the token is expired
// Should only be called through [auth.ProtectedMiddle] unless testing.
func ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return "", errors.New("Invalid claims")
	}

	if time.Now().After(claims.RegisteredClaims.ExpiresAt.Time) {
		return "", errors.New("Token expired")
	}

	return claims.Username, nil
}
