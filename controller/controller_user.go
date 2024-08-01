package controller

import (
	"encoding/json"
	"homework-apirest/model"
	"homework-apirest/service"
	"net/http"
)

type UserController struct {
	Service *service.UserService
}

func (c *UserController) CreateUserNew(w http.ResponseWriter, r *http.Request) {
	var user model.Users

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}
	newUser, err := c.Service.CreateUserNew(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)

}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
