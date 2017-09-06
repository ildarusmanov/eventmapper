package main

import (
	"eventmapper/mq"
	"eventmapper/configs"
	"eventmapper/middlewares"

	"fmt"
	"os"
	"path/filepath"

	"github.com/WajoxSoftware/middleware"
)

func main() {
	fmt.Println("Starting application...")

	fmt.Println("Load config")
	configFilePath, _ := filepath.Abs(os.Args[1])
	config := configs.LoadConfigFile(configFilePath)

	fmt.Println("Open MQ session")
	mqChannel, err := mq.CreateChannelFromString(config.MqUrl)
	defer mqChannel.Close()

	if err != nil {
		return
	}

	fmt.Println("Create router")
	router := CreateNewRouter(mqChannel)

	fmt.Println("Define middleware")
	mware := middleware.CreateNewMiddleware()
	mware.AddHandler(middlewares.CreateNewAuth(config.AuthToken))
	mware.AddHandler(middlewares.CreateNewJsonOkResponse())
	mware.AddHandler(router)

	fmt.Println("Start web-server")
	StartServer(mware, config)

	fmt.Println("Start events listener")
	StartEventsListener(mqChannel)
}
