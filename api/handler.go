package api

import (
	"log"
	"net/http"
)

type ApiHandler struct{}

func (apiHandler ApiHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/names", apiGetAll)
}

func apiGetAll(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	w.Write([]byte("hello"))
}
