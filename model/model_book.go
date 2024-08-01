package model

type Books struct {
	ID          int    `json:"book_id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Status      *bool  `json:"status"`
	Pages       int    `json:"pages"`
	Description string `json:"description"`
	Published   string `json:"published"`
	Image       string `json:"image"`
}
