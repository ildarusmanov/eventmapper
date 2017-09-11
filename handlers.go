package main

import (
	"eventmapper/configs"
	"eventmapper/models"
)

func BindEventsHandlers(config *configs.Config, closeCh chan bool, errCh chan error) {
	for _, cfg := range config.MqHandlers {
		mqH := models.BuildHandlerFromConfig(cfg)
		go mqH.StartListening(closeCh, errCh)
	}
}
