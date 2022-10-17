package services

import . "safer.com/m/models"
import  "github.com/google/uuid"

func AddUserToDataBase(user User) (User,error) {
	user.Id=uuid.New().String()
	_,err:=db.Exec("INSERT INTO users (Role, Email, Password, Id,FirstName,LastName,Phone) VALUES (?,?,?,?,?,?,?)",
		user.Role, user.Email, user.Password, user.Id, user.FirstName, user.LastName, user.Phone)
	if err != nil {
		return User{},err
	}else{
		return user,nil
	}
}
