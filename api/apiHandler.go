package api

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	svc "github.com/nohgo/go_networking/api/services"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", register)
	mux.HandleFunc("GET /api/users", getAll)
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewUserRepository())
	if err := us.Register("arya", "goodbye"); err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewUserRepository())
	users, err := us.GetAll()

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", users)))
}
