package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	NewRouter *chi.Mux
	Port      string
}

func NewServer(port string) *Server {
	return &Server{
		NewRouter: chi.NewRouter(),
		Port:      port,
	}
}

func (s *Server) InitRouter() *chi.Mux {
	return s.NewRouter
}

func (s *Server) RunServer(router http.Handler) error {
	log.Printf("server inicializado en http://localhost:%s", s.Port)
	if err := http.ListenAndServe(":"+s.Port, router); err != nil {
		log.Fatalf("No fue posible iniciar servidor en puerto %s %v", s.Port, router)
		return err
	}
	return nil
}
