package main

import (
	"Employee-REST-API/routes"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	fmt.Sprintf("Server Listening at:%s", "5001")
	router := chi.NewRouter()
	routes.EmployeeRoutes(router)
	http.ListenAndServe(":5001", router)
}
