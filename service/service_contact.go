package service

import (
	"homework-apirest/model"
	"homework-apirest/repository"
)

type ContactService struct {
	Repo *repository.ContactRepository
}

func (service *ContactService) CreateContactNew(contacts *model.Contacts) (*model.Contacts, error) {
	id, err := service.Repo.CreateContactNew(contacts)
	if err != nil {
		return nil, err
	}
	contacts.ID = id
	return contacts, nil

}
