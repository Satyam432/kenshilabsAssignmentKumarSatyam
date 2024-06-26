package models

type User struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserHash string `json:"userHash"`
}
