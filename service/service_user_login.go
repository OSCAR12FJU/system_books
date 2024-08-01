package service

import (
	"errors"
	"homework-apirest/model"
	"homework-apirest/repository"
)

type UserLoginService struct {
	Repo *repository.UserLoginRepository
}

func (service *UserLoginService) Autheticate(email, password string) (*model.UserLogin, error) {

	user, err := service.Repo.GetUserByEmail(email)
	if err != nil {
		return &model.UserLogin{}, err
	}

	err = service.Repo.ComparePasswords(user.Password, password)
	if err != nil {
		return &model.UserLogin{}, errors.New("Credenciales invalidas")
	}
	return user, nil
}
