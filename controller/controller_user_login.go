package controller

import (
	"encoding/json"
	"homework-apirest/model"
	"homework-apirest/service"
	"net/http"
)

type UserLoginController struct {
	Service *service.UserLoginService
}

func (c *UserLoginController) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Solicitud no valida", http.StatusBadRequest)
		return
	}

	userFromDB, err := c.Service.Autheticate(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return

	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userFromDB)
}

func (c *UserLoginController) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cierre de sesion exitoso"})
}
