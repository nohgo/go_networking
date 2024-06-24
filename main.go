package main

import (
	"log"
	"net/http"
)

type defaultHandler struct{}

func main() {
	var defh = defaultHandler{}
	http.Handle("/", defh)
	log.Printf("server started at %v", http.ListenAndServe(":8080", nil))
}

func (d defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	http.ServeFile(w, r, "html/index.html")
}
