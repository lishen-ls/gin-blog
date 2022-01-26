package models

type User struct {
	Id string `json:"id"`
	TimeStamp
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
