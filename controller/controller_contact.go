package controller

import (
	"encoding/json"
	"homework-apirest/model"
	"homework-apirest/service"
	"net/http"
)

type ContactController struct {
	Service *service.ContactService
}

func (c *ContactController) CreateContactNew(w http.ResponseWriter, r *http.Request) {
	var contacts model.Contacts

	err := json.NewDecoder(r.Body).Decode(&contacts)
	if err != nil {
		http.Error(w, "Datos invalidos", http.StatusBadRequest)
		return
	}
	newContact, err := c.Service.CreateContactNew(&contacts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newContact)
}
