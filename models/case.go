package models

type CaseStatus string

const (
	Open       CaseStatus = "open"
	Closed     CaseStatus = "closed"
	InProgress CaseStatus = "in-progress"
)


type Case struct{
	Id 			string	   `json:"id"`
	AssigneeId  string     `json:"assigneeId"`
	ReporterId  string	   `json:"reporterId"`
	Status 		CaseStatus `json:"status"`
}