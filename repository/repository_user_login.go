package repository

import (
	"database/sql"
	"errors"
	"homework-apirest/model"

	"golang.org/x/crypto/bcrypt"
)

type UserLoginRepository struct {
	DB *sql.DB
}

func (repo *UserLoginRepository) GetUserByEmail(email string) (*model.UserLogin, error) {
	var user model.UserLogin
	err := repo.DB.QueryRow("SELECT user_id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &model.UserLogin{}, errors.New("Usuario no existe")
		}
		return &model.UserLogin{}, err
	}
	return &user, nil
}

func (repo *UserLoginRepository) ComparePasswords(savedPassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(inputPassword))
}
