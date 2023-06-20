package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/journeybeforedestination/sizzle/pan"
)

var templates = template.Must(template.ParseFiles("tmpl/home.html", "tmpl/sizzles.html"))

type server struct {
	pan pan.Pan
}

// NewHTTPServer creates an http server
func NewHTTPServer(addr string) *http.Server {
	s := &server{
		pan: pan.Pan{},
	}
	s.pan.TurnItOn()
	r := mux.NewRouter()

	// setup handlers
	r.HandleFunc("/", s.handleRoot).Methods("GET")
	r.HandleFunc("/sizzles", s.handleSizzles).Methods("GET")

	// register middleware
	var handler http.Handler = r
	handler = logRequestHandler(handler)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}

// logRequestHandler is a middleware that writes an http log for each request
func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)

		url := r.URL.String()
		method := r.Method

		log.Printf("%s :: %s", url, method)
	}
	return http.HandlerFunc(fn)
}

// handleRoot handles the root path "/"
func (s *server) handleRoot(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}
func (s *server) handleSizzles(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "sizzles", s.pan.Sizzles)
}
