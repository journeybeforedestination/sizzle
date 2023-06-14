package main

import (
	"fmt"
	"net/http"

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
	msg := fmt.Sprintf("PATH: %s", r.URL.Path)
	w.Write([]byte(msg))
}
