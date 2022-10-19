package services

import (
	. "safer.com/m/models"
	"github.com/google/uuid"
)

func AddCaseToDataBase(_case Case) (Case,error) {
	_case.Id=uuid.New().String()
	_,err:=db.Exec("INSERT INTO cases (Id,ReporterId) VALUES (?,?)",
					_case.Id,_case.ReporterId)
	if err != nil {
		return Case{},err
	}else{
		return _case,nil
	}
}

func GetCasesByAssigneeId(assigneeId string) ([]Case,error){
	data,err:=db.Query("SELECT Id,AssigneeId,ReporterId,Status FROM cases WHERE AssigneeId=?",assigneeId)
	if err!=nil{
		return []Case{}, err
	}
	cases:=[]Case{}
	for data.Next(){
		var _case Case
		data.Scan(&_case.Id,&_case.AssigneeId,&_case.ReporterId,&_case.Status)
		cases=append(cases,_case)
	}
	return cases,nil
}

func AssignCaseToVolunteer(caseId,assigneeId string) error{
	_,err:=db.Exec("UPDATE users SET AssigneeId=? WHERE Id=?",assigneeId,caseId)

	if err!=nil{
		return err
	}
	return nil
}