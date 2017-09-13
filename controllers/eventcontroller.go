package controllers

import (
	"encoding/json"
	"eventmapper/models"
	"eventmapper/services"
	"github.com/gorilla/mux"
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
		return h.sendJsonResponse(w, false, err)
	}

	if err := services.PublishEvent(event, c.MqUrl, rKey); err != nil {
		return h.sendJsonResponse(w, false, err)
	}

	h.sendJsonResponse(w, true, "ok")
}

/**
 * send json response
 * @param  w http.ResponseWriter
 * @param isOk bool
 */
func (c *EventController) sendJsonResponse(w http.ResponseWriter, isOk bool, status string) {
	response := CreateNewJsonResponse(true, "ok")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
