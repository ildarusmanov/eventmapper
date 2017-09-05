package main

import (
	"eventmapper/configs"

	"log"
	"net/http"
	"time"
)

func StartServer(handler http.Handler, config *configs.Config) {
	srv := &http.Server{
		Handler: handler,
		Addr:    config.ServerHost,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
