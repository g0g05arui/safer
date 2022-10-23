package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
)
import . "safer.com/m/models"
func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	r.Route("/users", func(rx chi.Router) {
		rx.Post("/", CreateUser(Client))
		rx.Post("/admin",CreateAdmin)
		rx.Post("/volunteer",AuthMiddleWare([]string{"admin"},CreateUser(Volunteer)))
		rx.Put("/",AuthMiddleWare([]string{"client","admin","volunteer"},ChangeUserInformation))
	})

	r.Route("/cases", func(rx chi.Router) {
		rx.Post("/",AuthMiddleWare([]string{"client"},CreateCase))
		rx.Get("/",AuthMiddleWare([]string{"admin"},GetAllCases))
		rx.Post("/{id}/assign",AuthMiddleWare([]string{"admin"},AssignCase))
	})
	r.Route("/messages",func(rx chi.Router){
		rx.Post("/{caseId}",AuthMiddleWare([]string{"client","volunteer"},SendMessage))
	})
	r.Get("/my-cases",AuthMiddleWare([]string{"volunteer","client"},GetMyCases))
	r.Post("/login",AuthUser)
	r.Post("/test-role-authorization",AuthMiddleWare([]string{"client","admin"},TestAuth))
	return r
}
