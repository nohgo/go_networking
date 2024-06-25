package api

import (
	"log"
	"net/http"
)

type ApiHandler struct{}

func (d ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiGet(w, r)
	}
}

func apiGet(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	w.Write([]byte("hello"))
}
