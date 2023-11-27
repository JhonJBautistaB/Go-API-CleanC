package routes

import (
	"go-api-cleanc/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Get("/", handlers.HelloWorld)
	router.Get("/prueba", handlers.PruebaWeb)
}
