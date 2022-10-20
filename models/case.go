package models

import "database/sql"

type CaseStatus string

const (
	Open       CaseStatus = "open"
	Closed     CaseStatus = "closed"
	InProgress CaseStatus = "in-progress"
)


type Case struct{
	Id 			string	   `json:"id"`
	AssigneeId  sql.NullString     `json:"assigneeId"`
	ReporterId  string	   `json:"reporterId"`
	Status 		CaseStatus `json:"status"`
}