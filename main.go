package main

import (
	"Employee-REST-API/routes"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	logrus.Infof("Server Listening at:%v", "5001")
	router := chi.NewRouter()
	routes.EmployeeRoutes(router)
	log.Fatal(http.ListenAndServe(":5001", router))
}
