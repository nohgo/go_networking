package api

import (
	"github.com/nohgo/go_networking/api/auth"
	"log"
	"net/http"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", login)
	mux.HandleFunc("GET /api/users", auth.ProtectedMiddle(getAll))
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	token, err := auth.CreateToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write([]byte("Bearer " + token))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	w.Write([]byte(r.Header["Authorization"][0]))
}
