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
	r.Post("/cases/{id}/assign",AuthMiddleWare([]string{"admin"},AssignCase))
	r.Get("/my-cases",AuthMiddleWare([]string{"volunteer"},GetMyCases))
	r.Get("/cases",AuthMiddleWare([]string{"admin"},GetAllCases))
	return r
}
