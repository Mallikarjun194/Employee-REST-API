package main

import (
	"Employee-REST-API/routes"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	fmt.Println("Server Listening at:8080")
	router := chi.NewRouter()
	routes.EmployeeRoutes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":" + port, router)
}
