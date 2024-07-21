package auth

import (
	"log"
	"net/http"
)

// ProtectedMiddle returns an equivalent of the passed function with an authentication step beforehand.
// It adds the decrypted token to the header "Authorization"
// Proper calling is replacing YourFunction with ProtectedMiddle(YourFunction)
func ProtectedMiddle(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Authorization"]
		if !ok || len(token) < 1 {
			log.Println("No token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parsedName, err := ParseToken(token[0][7:])
		if err != nil {
			log.Println("Token is invalid")
			w.WriteHeader(401)
			return
		}
		r.Header.Set("Authorization", parsedName)
		next(w, r)
	}
}
