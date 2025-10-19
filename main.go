package main

import (
	"auth/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// initalize router
	r := mux.NewRouter()

	// passing route to register routes func
	routes.RegisterAuthRoutes(r)

	//initialize and listen server port
	fmt.Println("Server started at PORT 6969")
	http.ListenAndServe(":6969", r)
}
