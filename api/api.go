package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
    port string
}

func CreateServer(port string) *Server {
    return &Server {
        port: port,
    }
}

func (s *Server) Run() {
    r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from GTC!"))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", s.port), r)
}
