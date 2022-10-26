package services

import (
	. "safer.com/m/models"
	"github.com/google/uuid"
	"safer.com/m/utils"
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

func GetRecipientsChannelId(caseId string) ([]string,error){
	var assigneeId,reporterId string
	err:=db.QueryRow("SELECT ReporterId,AssigneeId FROM cases WHERE Id=?",caseId).Scan(&reporterId,&assigneeId)
	if err!=nil{
		return []string{}, err
	}

	data,err:=db.Query("SELECT Id,Password FROM users WHERE Id=? OR Id=?",reporterId,assigneeId)
	channels:=[]string{}
	for data.Next(){
		var id,pswd string
		data.Scan(&id,&pswd)
		channels=append(channels,utils.GetMD5Hash(id+pswd))
	}

	return channels,nil

}

func GetChannelToken(userId string) string{
	var id,pwd string
	db.QueryRow("SELECT Id,Password FROM users WHERE Id=?",userId).Scan(&id,&pwd)
	return utils.GetMD5Hash(id+pwd)
}

func GetMessagesByCaseId(caseId string) []Message{
	var message Message
	messages:=[]Message{}
	data,err:=db.Query("SELECT Id,SenderId,CaseId,CreatedAt,IFNULL(Data,''),IFNULL(Message,'') FROM messages WHERE CaseId=?",caseId)
	if err!=nil{
		return messages
	}
	for data.Next(){
		data.Scan(&message.Id,&message.SenderId,&message.CaseId,&message.CreatedAt,&message.Data,&message.Message)
		messages=append(messages,message)
	}
	return messages
}