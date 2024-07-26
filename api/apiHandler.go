// The api package contains the routes starting with /api.
// Proper usage is creating an [api.ApiHandler] then calling its InitRoutes function
package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/nohgo/go_networking/api/auth"
	repo "github.com/nohgo/go_networking/api/database/repositories"
	"github.com/nohgo/go_networking/api/help"
	"github.com/nohgo/go_networking/api/models"
	svc "github.com/nohgo/go_networking/api/services"
)

// The ApiHandler struct adheres to the [main.Router] interface and initializes the api routes.
// Create the ApiHandler with a struct literal and then call InitRoutes
type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/sign-up", Register)
	mux.HandleFunc("POST /api/auth/login", Login)
	mux.HandleFunc("DELETE /api/auth", auth.ProtectedMiddle(DeleteUser))
	mux.HandleFunc("GET /api/cars", auth.ProtectedMiddle(GetAll))
	mux.HandleFunc("POST /api/cars", auth.ProtectedMiddle(PostCar))
	mux.HandleFunc("DELETE /api/cars", auth.ProtectedMiddle(DeleteCar))
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	us := svc.NewUserService(repo.NewPostgresUserRepository())

	var user models.User
	if err, code := help.DecodeStruct(r, &user); err != nil {
		http.Error(w, err.Error(), code)
	}

	err := us.Register(user)
	if errors.Is(err, &models.UserError{}) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
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

func GetAll(w http.ResponseWriter, r *http.Request) {
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

func PostCar(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	cs := svc.NewCarService(repo.NewCarRepository())

	var car models.Car
	if err, code := help.DecodeStruct(r, &car); err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	username := r.Header.Get("Authorization")
	if err := cs.Add(car, username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)

	username := r.Header.Get("Authorization")
	if len(username) == 0 {
		http.Error(w, "Invalid JWT token", http.StatusBadRequest)
	}
	us := svc.NewUserService(repo.NewPostgresUserRepository())

	err := us.Delete(models.User{Username: username})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
