package api

import (
	"github.com/nohgo/go_networking/api/security"
	"log"
	"net/http"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", login)
	mux.HandleFunc("POST /api/users/{token}", getAll)
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	token, err := jwt.CreateToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	w.Write([]byte("Bearer " + token))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	token := r.PathValue("token")
	name, err := jwt.ParseToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	w.Write([]byte(name))
}
