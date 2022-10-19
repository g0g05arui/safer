package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"safer.com/m/internal/env"
	. "safer.com/m/models"
	"safer.com/m/services"
	"strings"
)
func CreateCase(w http.ResponseWriter,r *http.Request){
	var _case Case
	authToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	claims := jwt.MapClaims{}
	_,err:=jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Cfg["SECRET_KEY"]), nil
	})
	if err!=nil{
		return
	}
	_case.ReporterId=claims["id"].(string)
	_case.AssigneeId="";
	if _case, err := services.AddCaseToDataBase(_case); err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}else{
		json.NewEncoder(w).Encode(_case)
	}

}