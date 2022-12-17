package routes

import (
	"Employee-REST-API/config"
	"Employee-REST-API/controllers"
	"github.com/go-chi/chi"
)

var EmployeeRoutes = func(router *chi.Mux) {

	router.Post(config.ENDPOINT, controllers.CreateEmployee)
	router.Get(config.ENDPOINT, controllers.GetEmployees)
	router.Get("/empByID", controllers.GetEmployeeByID)
	router.Put(config.ENDPOINT, controllers.UpdateEmployee)
	router.Delete(config.ENDPOINT, controllers.DeleteEmployee)
}
