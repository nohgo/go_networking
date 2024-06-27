package auth

import (
	"log"
	"net/http"
)

func ProtectedMiddle(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Authorization"]
		if !ok || len(token) < 1 {
			log.Println("No token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parsedName, err := parseToken(token[0][7:])
		if err != nil {
			log.Println("Token is invalid")
			w.WriteHeader(401)
			return
		}
		r.Header.Set("Authorization", parsedName)
		next(w, r)
	}
}
