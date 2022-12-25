package routes

import (
	"Employee-REST-API/config"
	"Employee-REST-API/handlers"
	"github.com/go-chi/chi"
)

var EmployeeRoutes = func(router *chi.Mux) {

	router.Post(config.ENDPOINT, handlers.CreateEmployee)
	router.Get(config.ENDPOINT, handlers.GetEmployees)
	router.Get("/empByID", handlers.GetEmployeeByID)
	router.Put(config.ENDPOINT, handlers.UpdateEmployee)
	router.Delete(config.ENDPOINT, handlers.DeleteEmployee)
}
