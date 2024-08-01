package repository

import (
	"database/sql"
	"homework-apirest/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) CreateUserNew(users *model.Users) (int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	users.Password = string(hashedPassword)
	query := `INSERT INTO users (first_name, last_name,  email, password, nacionality) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`

	var lastInsertId int
	err = repo.DB.QueryRow(query, users.FirstName, users.LastName, users.Email, users.Password, users.Nacionality).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil

	// result, err := repo.DB.Exec(`INSERT INTO users (first_name, last_name, email, password, nacionality) VALUES (?,?,?,?,?)`, users.FirstName, users.LastName, users.Email, users.Password, users.Nacionality)
	// if err != nil {
	// 	return 0, nil
	// }
	// lastInsertId, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, nil
	// }
	// return int(lastInsertId), nil

}

func (repo *UserRepository) GetAllUsers() ([]model.Users, error) {
	rows, err := repo.DB.Query("SELECT first_name, last_name ,email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users

	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
