package main

import (
	"eventmapper/configs"
	"eventmapper/controllers"
	"eventmapper/middlewares"
	"github.com/WajoxSoftware/middleware"
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

func StartHttpsServer(config *configs.Config) *http.Server {
	log.Printf("[x] Create router")
	routerHandler := CreateNewRouterHandler(config)

	log.Printf("[x] Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.HttpAuthType, config.HttpAuthParams))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(routerHandler)

	srv := &http.Server{
		Handler: mware,
		Addr:    config.ServerHost,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Fatal(srv.ListenAndServeTLS(config.ServerTLSCrt, config.ServerTLSKey))
	}()

	return srv
}
