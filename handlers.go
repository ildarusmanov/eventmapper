package main

import (
	"eventmapper/configs"
	"eventmapper/models"
)

func BindEventsHandlers(config *configs.Config) {
	for _, cfg := range config.MqHandlers {
		mqH := models.BuildHandlerFromConfig(cfg)
		mqH.StartListening()
	}
}
