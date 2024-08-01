package model

type UserLogin struct {
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
