package handlers

import (
	"encoding/json"
	"net/http"
	. "safer.com/m/models"
	"safer.com/m/services"
)
func CreateCase(w http.ResponseWriter,r *http.Request){
	var _case Case
	json.NewDecoder(r.Body).Decode(&_case)


	if _case, err := services.AddCaseToDataBase(_case); err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HttpError{Message: err.Error()})
	}else{
		json.NewEncoder(w).Encode(_case)
	}

}