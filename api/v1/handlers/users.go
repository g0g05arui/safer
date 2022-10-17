package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"gopkg.in/validator.v2"
	. "safer.com/m/models"
	"safer.com/m/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.Role = "client"
	if err := validator.Validate(user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})

	} else {
		if user,err:=services.AddUserToDataBase(user); err!=nil{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}else{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	}

}

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/users", CreateUser)
	return r
}
