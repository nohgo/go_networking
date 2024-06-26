package frontend

import (
	"log"
	"net/http"
)

type FrontendHandler struct{}

func (d FrontendHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	http.ServeFile(w, r, "frontend/html/index.html")
}
