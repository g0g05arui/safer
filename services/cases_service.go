package services

import (
	. "safer.com/m/models"
	"github.com/google/uuid"
)

func AddCaseToDataBase(_case Case) (Case,error) {
	_case.Id=uuid.New().String()
	_,err:=db.Exec("INSERT INTO users () VALUES (?,?,?,?,)",
		)
	if err != nil {
		return Case{},err
	}else{
		return _case,nil
	}
}