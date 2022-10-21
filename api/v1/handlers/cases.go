package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"net/http"
	. "safer.com/m/models"
	"safer.com/m/services"

)
func CreateCase(w http.ResponseWriter,r *http.Request){
	var _case Case
	claims:=r.Context().Value("user").(jwt.MapClaims)
	_case.ReporterId=claims["id"].(string)
	if _case, err := services.AddCaseToDataBase(_case); err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}else{
		json.NewEncoder(w).Encode(_case)
	}

}

func AssignCase(w http.ResponseWriter,r* http.Request){
	var _case Case
	id:=chi.URLParam(r,"id")
	json.NewDecoder(r.Body).Decode(&_case)

	if err:=services.AssignCaseToVolunteer(id,_case.AssigneeId);err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}else{
		json.NewEncoder(w).Encode(map[string]interface{}{"message":"case updated succesfully"})
	}
}

func GetAllCases(w http.ResponseWriter,r *http.Request){
	cases:=services.GetCases()
	json.NewEncoder(w).Encode(cases)
}

func GetMyCases(w http.ResponseWriter,r *http.Request){
	claims:=r.Context().Value("user").(jwt.MapClaims)

	if cases, err :=services.GetCasesById(claims["id"].(string));err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}else{
		json.NewEncoder(w).Encode(cases)
	}
}