package main

import (
	"Employee-REST-API/routes"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	fmt.Println("Server Listening at:5001")
	router := chi.NewRouter()
	routes.EmployeeRoutes(router)
	http.ListenAndServe("localhost:5001", router)
}
