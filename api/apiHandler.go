package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/models"
	svc "github.com/nohgo/go_networking/api/services"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/sign-up", register)
	mux.HandleFunc("POST /api/auth/login", login)
	mux.HandleFunc("GET /api/users", auth.ProtectedMiddle(getAll))
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewUserRepository())
	if err := us.Register(models.User{Username: "arya", Password: "hello"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewUserRepository())

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := us.Login(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonToken, err := json.Marshal(token)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonToken)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewUserRepository())
	users, err := us.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUsers, err := json.Marshal(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonUsers)
}
