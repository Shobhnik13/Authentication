package controllers

import (
	"auth/helpers"
	"auth/models"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

	// checking dup email
	_, err = helpers.GetUserByEmail(user.Email)
	if err == nil {
		http.Error(w, "Email already in use", http.StatusBadRequest)
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

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var creds models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if creds.Email == "" || creds.Password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	user, err := helpers.GetUserByEmail(creds.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"uid":   user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// load env
	errEnv := godotenv.Load()
	if errEnv != nil {
		http.Error(w, "Error loading environment variables", http.StatusInternalServerError)
		return
	}

	secretKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// getting user info from context
	// like req.user in node.js
	userClaims := r.Context().Value("user")
	

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Your profile is this",
		"user":    userClaims,
	})
}
