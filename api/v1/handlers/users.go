package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gopkg.in/validator.v2"
	"net/http"
	"safer.com/m/internal/env"
	. "safer.com/m/models"
	"safer.com/m/services"
	"safer.com/m/utils"
	"strings"
)

func CreateUser(userType UserType) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.Role = userType
		if err := validator.Validate(user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})

		} else {
			if user,err:=services.AddUserToDataBase(user); err!=nil{
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
			}else{
				json.NewEncoder(w).Encode(user)
			}
		}

	}
}
func CreateAdmin(w http.ResponseWriter,r *http.Request){
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	if authToken!=env.Cfg["SECRET_KEY"]{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.Role = Admin
	if err := validator.Validate(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})

	} else {
		if user,err:=services.AddUserToDataBase(user); err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}else{
			json.NewEncoder(w).Encode(user)
		}
	}
}

func AuthUser(w http.ResponseWriter, r *http.Request){
	var userBody User
	json.NewDecoder(r.Body).Decode(&userBody)

	if user,err:=services.GetUserBasicAuth(userBody.Email,userBody.Password);err!=nil{
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		return
	}else{
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

func TestAuth(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(map[string]interface{}{"ok":"ok"})
}


func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/users", CreateUser(Client))
	r.Post("/login",AuthUser)
	r.Post("/test-role-authorization",AuthMiddleWare([]string{"client","admin"},TestAuth))
	r.Post("/users/admin",CreateAdmin)
	r.Post("/users/volunteer",AuthMiddleWare([]string{"admin"},CreateUser(Volunteer)))
	return r
}
