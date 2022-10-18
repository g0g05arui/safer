package models

type UserType string

const (
	Admin     UserType = "admin"
	Client    UserType = "client"
	Volunteer UserType = "volunteer"
)

type User struct {
	Id        string   `json:"id"`
	Email     string   `json:"email" validate:"min=1,max=255"`
	Password  string   `json:"password" validate:"min=1,max=255"`
	FirstName string   `json:"firstName" validate:"min=1,max=16,regexp=^[a-zA-Z]*$"`
	LastName  string   `json:"lastName" validate:"min=1,max=16,regexp=^[a-zA-Z]*$"`
	Phone     string   `json:"phone" validate:"min=1,max=16,regexp=^[0-9]*$"`
	Role      UserType `json:"role"`
}
