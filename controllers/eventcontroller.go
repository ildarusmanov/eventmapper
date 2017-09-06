package controllers

import (
	"eventmapper/models"
	"eventmapper/mq"

	"encoding/json"
	"net/http"
)

type EventController struct {
	mqChannel *mq.Channel
}

/**
 * Event controller constructor
 * 
 * @param *mq.Session
 * @param *configs.Config
 * @return *EventController
 */
func CreateNewEventController(mqChannel *mq.Channel) *EventController {
	return &EventController{mqChannel}
}

/**
 * Publish event method
 * 
 * @param http.ResponseWriter
 * @param *http.Request
 */
func (c *EventController) CreateHandler(w http.ResponseWriter, r *http.Request) {
    event := models.CreateNewEvent()

    if err := json.NewDecoder(r.Body).Decode(event); err != nil {
   		panic(err)
    }

	if err := event.Validate(); err != nil {
		panic(err)
	}

 	if err := event.Publish(c.mqChannel); err != nil {
 		panic(err)
 	}

    if err := json.NewEncoder(w).Encode(event); err != nil {
        panic(err)
    }
}
