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
	mux.HandleFunc("GET /api/cars", auth.ProtectedMiddle(getAll))
	mux.HandleFunc("POST /api/cars", auth.ProtectedMiddle(postCar))
}

func register(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	us := svc.NewUserService(repo.NewUserRepository())

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := us.Register(user); err != nil {
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

	cs := svc.NewCarService(repo.NewCarRepository())

	username := r.Header.Get("authorization")
	cars, err := cs.GetAll(username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonCars, err := json.Marshal(cars)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonCars)
}

func postCar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	cs := svc.NewCarService(repo.NewCarRepository())

	var car models.Car

	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.Header.Get("authorization")
	if err := cs.Add(car, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
