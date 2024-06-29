package main

import (
	"github.com/nohgo/go_networking/api"
	"github.com/nohgo/go_networking/api/database"
	"github.com/nohgo/go_networking/frontend"
	"log"
	"net/http"
)

type Router interface {
	InitRoutes(*http.ServeMux)
}

func main() {
	mux := http.NewServeMux()
	routers := []Router{api.ApiHandler{}, frontend.FrontendHandler{}}
	for _, v := range routers {
		v.InitRoutes(mux)
	}

	db.CreatePool()
	defer db.ClosePool()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
