package main

import (
	"github.com/nohgo/go_networking/api"
	"log"
	"net/http"
)

func main() {
	var apiH = api.ApiHandler{}
	http.Handle("/", apiH)
	log.Printf("server started at %v", http.ListenAndServe(":8080", nil))
}
