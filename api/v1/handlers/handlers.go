package handlers

import "github.com/go-chi/chi/v5"
import . "safer.com/m/models"
func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/users", CreateUser(Client))
	r.Post("/login",AuthUser)
	r.Post("/test-role-authorization",AuthMiddleWare([]string{"client","admin"},TestAuth))
	r.Post("/users/admin",CreateAdmin)
	r.Post("/users/volunteer",AuthMiddleWare([]string{"admin"},CreateUser(Volunteer)))
	r.Put("/users",AuthMiddleWare([]string{"client","admin","volunteer"},ChangeUserInformation))

	r.Post("/cases",AuthMiddleWare([]string{"client"},CreateCase))

	return r
}
