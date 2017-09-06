package main

import (
	"eventmapper/controllers"
	"eventmapper/mq"
	"fmt"
	"github.com/gorilla/mux"
)

func CreateNewRouter(mqChannel *mq.Channel) *mux.Router {
	router := mux.NewRouter()

	fmt.Println("Create controller")
	controller := controllers.CreateNewEventController(mqChannel)

	fmt.Println("Define routes")
	router.HandleFunc("/create", controller.CreateHandler).Methods("POST")

	return router
}
