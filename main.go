package main

import (
	"github.com/nohgo/go_networking/api"
	"github.com/nohgo/go_networking/frontend"
	"log"
	"net/http"
)

func main() {
	http.Handle("/api", api.ApiHandler{})
	http.Handle("/", frontend.FrontendHandler{})
	log.Printf("server started at %v", http.ListenAndServe(":8080", nil))
}
