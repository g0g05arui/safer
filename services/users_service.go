package services

import (
	"errors"
	. "safer.com/m/models"
	"safer.com/m/utils"
)
import  "github.com/google/uuid"

func AddUserToDataBase(user User) (User,error) {
	user.Id=uuid.New().String()
	_,err:=db.Exec("INSERT INTO users (Role, Email, Password, Id,FirstName,LastName,Phone) VALUES (?,?,?,?,?,?,?)",
		user.Role, user.Email, utils.GetMD5Hash(user.Password), user.Id, user.FirstName, user.LastName, user.Phone)
	if err != nil {
		return User{},err
	}else{
		return user,nil
	}
}

func GetUserBasicAuth(email,password string) (User,error){
	rows, err := db.Query(
		"SELECT Id, FirstName,LastName, Role FROM users WHERE email=? AND password=?",
		email, utils.GetMD5Hash(password))
	if err != nil {
		return User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var p User
		if err := rows.Scan(&p.Id, &p.FirstName, &p.LastName,&p.Role); err != nil {
			return User{}, err
		}
		return p,nil
	}
	return User{}, errors.New("User not found")
}

func UpdateUserInfo(user User) (User,string,error){

	_,err:=db.Exec("UPDATE users SET Email=?,FirstName=?,LastName=?,Phone=? WHERE Id=?",user.Email,user.FirstName,user.LastName,user.Phone,user.Id)
	if err!=nil{
		return User{}, "", err
	}else{
		if token,err:=utils.GenerateJWT(user);err!=nil{
			return User{}, "", err
		}else{
			return user,token,nil
		}
	}
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