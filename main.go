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
        // setup log file
        logFilePath, _ := filepath.Abs(os.Args[1]);
	f, err := os.OpenFile(logFilePath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
        log.SetOutput(f)

	log.Printf("[x] Starting application...")

	log.Printf("[x] Load config")
	configFilePath, _ := filepath.Abs(os.Args[2])
	config := configs.LoadConfigFile(configFilePath)

	log.Printf("[x] Create router")
	routerHandler := CreateNewRouterHandler(config)

	log.Printf("[x] Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.AuthToken))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(routerHandler)

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
