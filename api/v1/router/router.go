package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"safer.com/m/api/v1/handlers"
)

func setJsonResponseType(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w,r)
	})
}

func Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		r.Use(setJsonResponseType)
		r.Mount("/", handlers.Routes())
	})
	return router
}
