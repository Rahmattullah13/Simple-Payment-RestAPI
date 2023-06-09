package model

type User struct {
	UserId   string `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserRequest struct {
	Email    string `json:"email" db:"email" example:"johndoe@mail.com"`
	Password string `json:"password" db:"password" example:"qwerty"`
}
