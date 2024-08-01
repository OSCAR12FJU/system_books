package model

type Users struct {
	ID          int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Nacionality string `json:"nacionality"`
}
