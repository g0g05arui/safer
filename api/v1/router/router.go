package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"safer.com/m/api/v1/handlers"
)

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handlers.Routes())
	})
	return router
}
