package main

import (
	"github.com/ildarusmanov/eventmapper/configs"
	"github.com/ildarusmanov/eventmapper/models"
)

func BindEventsHandlers(config *configs.Config, closeCh chan bool) {
	for _, cfg := range config.MqHandlers {
		go models.StartHandler(cfg, closeCh)
	}
}
