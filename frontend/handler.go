package frontend

import (
	"log"
	"net/http"
	"strings"
)

type FrontendHandler struct{}

func (d FrontendHandler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	p := r.URL.Path
	if p == "/" {
		http.ServeFile(w, r, "frontend/html")
		return
	}
	if strings.Contains(p, "../") {
		w.WriteHeader(401)
		return
	}
	http.ServeFile(w, r, "frontend/html"+p+".html")
}
