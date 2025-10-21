package routes

import (
	"auth/controllers"
	"auth/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST")

	// protected route
	router.Handle("/api/v1/auth/profile",
	 middleware.AuthProtect(http.HandlerFunc(controllers.Profile)),
	 ).Methods("GET")

}
