package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"net/http"
	. "safer.com/m/models"
	"safer.com/m/services"
)

func SendMessage(w http.ResponseWriter,r *http.Request){
	claims:=r.Context().Value("user").(jwt.MapClaims)

	if services.CheckMessagePermissions(claims["id"].(string),UserType(claims["role"].(string)),chi.URLParam(r,"caseId")){
		var message Message
		json.NewDecoder(r.Body).Decode(&message)
		message.SenderId=claims["id"].(string)
		message.CaseId=chi.URLParam(r,"caseId")
		if message,err:=services.CreateMessage(message);err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
		}else{
			json.NewEncoder(w).Encode(message)
		}

	}else{
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(HttpError{Message: "Unauthorized"})
	}
}