package main

import (
	"eventmapper/configs"
	"eventmapper/middlewares"
	"github.com/WajoxSoftware/middleware"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Printf("[x] Starting application...")

	log.Printf("[x] Load config")
	configFilePath, _ := filepath.Abs(os.Args[1])
	config := configs.LoadConfigFile(configFilePath)

	log.Printf("[x] Create router")
	router := CreateNewRouter(config)

	log.Printf("[x] Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.AuthToken))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(router)

	if config.DisableHandlers {
		log.Printf("[*] Handlers are disabled")
	} else {
		log.Printf("[x] Start events listener")
		closeCh := make(chan bool)
		BindEventsHandlers(config, closeCh)
	}

	if config.DisableGrpc {
		log.Printf("[*] GRPC is disabled")
	} else {
		log.Printf("[x] Start grpc server")
		StartGrpc(config)
	}

	log.Printf("[x] Start web-server")
	StartHttpServer(mware, config)
}
