package main

import (
	"eventmapper/configs"
	"eventmapper/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func CreateNewRouter(config *configs.Config) *mux.Router {
	router := mux.NewRouter()

	log.Printf("[x] Create controller")
	controller := controllers.CreateNewEventController(config.MqUrl)

	log.Printf("[x] Define routes")
	router.HandleFunc("/create/{r_key}", controller.CreateHandler).Methods("POST")

	return router
}

func StartHttpServer(handler http.Handler, config *configs.Config) {
	srv := &http.Server{
		Handler: handler,
		Addr:    config.ServerHost,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}