package routes

import (
	"auth/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
}
