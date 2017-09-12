package main

import (
	"eventmapper/configs"
	"eventmapper/models"
)

func BindEventsHandlers(config *configs.Config, closeCh chan bool) {
	for _, cfg := range config.MqHandlers {
		go models.StartHandler(cfg, closeCh)
	}
}
