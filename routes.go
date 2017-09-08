package main

import (
	"eventmapper/configs"
	"eventmapper/controllers"
	"github.com/gorilla/mux"
	"log"
)

func CreateNewRouter(config *configs.Config) *mux.Router {
	router := mux.NewRouter()

	log.Printf("[x] Create controller")
	controller := controllers.CreateNewEventController(config.MqUrl)

	log.Printf("[x] Define routes")
	router.HandleFunc("/create/{r_key}", controller.CreateHandler).Methods("POST")

	return router
}
