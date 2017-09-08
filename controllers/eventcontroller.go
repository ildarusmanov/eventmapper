package controllers

import (
	"eventmapper/models"
	"eventmapper/mq"

	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
)

type EventController struct {
	mqUrl string
}

/**
 * Event controller constructor
 * 
 * @param *mq.Session
 * @param *configs.Config
 * @return *EventController
 */
func CreateNewEventController(mqUrl string) *EventController {
	return &EventController{mqUrl}
}

/**
 * Publish event method
 * 
 * @param http.ResponseWriter
 * @param *http.Request
 */
func (c *EventController) CreateHandler(w http.ResponseWriter, r *http.Request) {
    event := models.CreateNewEvent()
    vars := mux.Vars(r)
    rKey := vars["r_key"]

    if err := json.NewDecoder(r.Body).Decode(event); err != nil {
   		panic(err)
    }

	if err := event.Validate(); err != nil {
		panic(err)
	}

	mqConn, err := mq.CreateNewConnection(c.mqUrl)
	defer mqConn.Close()

	if err != nil {
		panic(err)
	}

	mqChannel, err := mq.CreateNewChannel(mqConn)	
	defer mqChannel.Close()

	if err != nil {
		panic(err)
	}

 	if err := event.Publish(mqChannel, rKey); err != nil {
 		panic(err)
 	}

    if err := json.NewEncoder(w).Encode(event); err != nil {
        panic(err)
    }
}
