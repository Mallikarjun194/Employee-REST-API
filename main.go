package main

import (
	"Employee-REST-API/config"
	"Employee-REST-API/routes"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	logrus.Infof("Server Listening at:%v", config.PORT)
	router := chi.NewRouter()
	routes.EmployeeRoutes(router)
	log.Fatal(http.ListenAndServe(config.KEY_SEPARATOR+config.PORT, router))
}
