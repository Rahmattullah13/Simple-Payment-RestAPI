package model

type UserCredential struct {
	UserId   string
	Password string `json:"password"`
	Email    string `json:"email"`
}
