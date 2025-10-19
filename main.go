package main

import (
	"auth/config"
	"auth/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// loading env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// extracting port from env
	port := os.Getenv("PORT")

	// connecting db 
	err = config.ConnectDB()
	if err != nil {
		fmt.Println("DB connection error:", err)
		return
	}
	log.Println("DB connected now starting the server...")

	// initalize router
	r := mux.NewRouter()

	// passing route to register routes func
	routes.RegisterAuthRoutes(r)

	//initialize and listen server port
	fmt.Printf("Server is running on PORT %s\n", port)
	http.ListenAndServe(":"+port, r)
}
