package main

import (
	"eventmapper/configs"
	"eventmapper/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type RouterHandler struct {
	router *mux.Router
}

func CreateNewRouterHandler(config *configs.Config) *RouterHandler {
	return &RouterHandler{createNewRouter(config)}
}

func (h *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
	h.router.ServeHTTP(w, r)

	return true
}

func createNewRouter(config *configs.Config) *mux.Router {
	router := mux.NewRouter()

	log.Printf("[x] Create controller")
	controller := controllers.CreateNewEventController(config.MqUrl)

	log.Printf("[x] Define routes")
	router.HandleFunc("/{r_key}/events", controller.CreateHandler).Methods("POST")

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
