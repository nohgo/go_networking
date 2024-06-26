package frontend

import (
	"log"
	"net/http"
)

type FrontendHandler struct{}

func (d FrontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiGet(w, r)
	}
}

func apiGet(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	http.ServeFile(w, r, "frontend/html/index.html")
}
