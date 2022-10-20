package services

import (
	"database/sql"
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

func GetCasesById(assigneeId string) ([]Case,error){
	data,err:=db.Query("SELECT Id,AssigneeId,ReporterId,Status FROM cases WHERE AssigneeId=? OR ReporterId=?",assigneeId,assigneeId)
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

func AssignCaseToVolunteer(caseId string,assigneeId sql.NullString) error{
	_,err:=db.Exec("UPDATE cases SET AssigneeId=?,Status=? WHERE Id=?",assigneeId,"in-progress",caseId)
	if err!=nil{
		return err
	}
	return nil
}


func GetCases()[]Case{
	data,err:=db.Query("SELECT Id,AssigneeId,ReporterId,Status FROM cases")
	if err!=nil{
		return []Case{}
	}
	var _case Case
	cases:=[]Case{}
	for data.Next(){
		data.Scan(&_case.Id,&_case.AssigneeId,&_case.ReporterId,&_case.Status)
		cases=append(cases,_case)
	}
	return cases
}