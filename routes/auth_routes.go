package routes

import (
	"auth/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
}
