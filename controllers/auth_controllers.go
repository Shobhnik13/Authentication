package controllers

import (
	"auth/helpers"
	"auth/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validation
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Name, Email and Password are required", http.StatusBadRequest)
		return
	}

	// HASHING PASSWORD
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashed)

	createdUser, err := helpers.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdUser)
}
