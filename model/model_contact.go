package model

type Contacts struct {
	ID      int     `json:"contact_id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Message *string `json:"message"`
}
