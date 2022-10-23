package services

import (
	. "safer.com/m/models"
	"github.com/google/uuid"
	"time"
)

func CreateMessage(message Message) (Message,error){
	message.Id=uuid.New().String()
	message.CreatedAt=time.Now()
	_,err:=db.Exec("INSERT INTO messages (Id,CaseId,CreatedAt,Data,Message,SenderId) VALUES (?,?,?,?,?,?)",
		message.Id,message.CaseId,message.CreatedAt,message.Data,message.Message,message.SenderId)
	if err!=nil{
		return Message{}, err
	}
	return message,nil
}