package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"gopkg.in/validator.v2"
	. "safer.com/m/models"
	"safer.com/m/services"
	"safer.com/m/utils"
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

func AuthUser(w http.ResponseWriter, r *http.Request){
	var userBody User
	json.NewDecoder(r.Body).Decode(&userBody)

	if user,err:=services.GetUserBasicAuth(userBody.Email,userBody.Password);err!=nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		return
	}else{
		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%s",user.Id)
		user.Email=userBody.Email
		if token,err:=utils.GenerateJWT(user);err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}else{
			json.NewEncoder(w).Encode(map[string]interface{}{"token":token})
		}
	}
}



func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/users", CreateUser)
	r.Post("/login",AuthUser)
	return r
}
