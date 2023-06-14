package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type server struct{}

func NewHTTPServer(addr string) *http.Server {
	s := &server{}
	r := mux.NewRouter()

	r.HandleFunc("/", s.handleRoot).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/home.html")
	t.Execute(w, nil)
}
