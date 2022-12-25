package routes

import (
	"Employee-REST-API/handlers"
	"github.com/go-chi/chi"
)

var EmployeeRoutes = func(router *chi.Mux) {

	router.Post("/emp", handlers.CreateEmployee)
	router.Get("/emp", handlers.GetEmployees)
	router.Get("/emp/{empID}", handlers.GetEmployeeByID)
	router.Put("/emp/{empID}", handlers.UpdateEmployee)
	router.Delete("/emp/{empID}", handlers.DeleteEmployee)
}
