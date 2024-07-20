package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/help"
	"github.com/nohgo/go_networking/api/models"
	svc "github.com/nohgo/go_networking/api/services"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/sign-up", Register)
	mux.HandleFunc("POST /api/auth/login", login)
	mux.HandleFunc("GET /api/cars", auth.ProtectedMiddle(getAll))
	mux.HandleFunc("POST /api/cars", auth.ProtectedMiddle(postCar))
	mux.HandleFunc("DELETE /api/cars", auth.ProtectedMiddle(deleteCar))
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	us := svc.NewUserService(repo.NewPostgresUserRepository())

	var user models.User
	if err, code := help.DecodeStruct(r, &user); err != nil {
		http.Error(w, err.Error(), code)
	}

	if err := us.Register(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	us := svc.NewUserService(repo.NewPostgresUserRepository())

	var user models.User
	if err, code := help.DecodeStruct(r, &user); err != nil {
		http.Error(w, err.Error(), code)
	}

	token, err := us.Login(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	help.SendJson(w, token)
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
	help.SendJson(w, cars)
}

func postCar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	cs := svc.NewCarService(repo.NewCarRepository())

	var car models.Car
	if err, code := help.DecodeStruct(r, &car); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	username := r.Header.Get("authorization")
	if err := cs.Add(car, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	queries := r.URL.Query()

	idArr, ok := queries["id"]
	if !ok || len(idArr) < 1 {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idArr[0])
	if err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	cs := svc.NewCarService(repo.NewCarRepository())
	username := w.Header().Get("Authorization")

	if err := cs.Delete(id, username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
