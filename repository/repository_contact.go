package repository

import (
	"database/sql"
	"homework-apirest/model"
)

type ContactRepository struct {
	DB *sql.DB
}

func (repo *ContactRepository) CreateContactNew(contacts *model.Contacts) (int, error) {
	result, err := repo.DB.Exec(`INSERT INTO contact (name, email,message) VALUES (?,?,?)`, contacts.Name, contacts.Email, contacts.Message)
	if err != nil {
		return 0, nil
	}
	lastInsetId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(lastInsetId), nil

}
