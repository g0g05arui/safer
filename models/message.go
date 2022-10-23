package models

import (
	_"database/sql"
	"time"
)

type Message struct{
	Id        string         `json:"id"`
	SenderId  string         `json:"senderId"`
	CaseId    string         `json:"caseId" validate:"min=1,max=255"`
	Message   string         `json:"message" validate:"min=1,max=255"`
	Data      string         `json:"data"`
	CreatedAt time.Time		 `json:"createdAt"`
}