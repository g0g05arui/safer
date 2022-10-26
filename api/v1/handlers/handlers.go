package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pusher/pusher-http-go"
	"safer.com/m/internal/env"
	"time"
)
import . "safer.com/m/models"
func Routes() *chi.Mux {

	pusherClient := pusher.Client{
		AppID: env.Cfg["PUSHER_APP_ID"],
		Key: env.Cfg["PUSHER_KEY"],
		Secret: env.Cfg["PUSHER_SECRET"],
		Cluster: env.Cfg["PUSHER_CLUSTER"],
		Secure: true,
	}

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
		rx.Post("/{caseId}",AuthMiddleWare([]string{"client","volunteer"},SendMessage(pusherClient)))
		rx.Get("/token",AuthMiddleWare([]string{"client","volunteer","admin"},GetMessageToken))
		rx.Get("/{caseId}",AuthMiddleWare([]string{"client","volunteer"},GetMessages))
	})
	r.Get("/my-cases",AuthMiddleWare([]string{"volunteer","client"},GetMyCases))
	r.Post("/login",AuthUser)
	r.Post("/test-role-authorization",AuthMiddleWare([]string{"client","admin"},TestAuth))
	return r
}
