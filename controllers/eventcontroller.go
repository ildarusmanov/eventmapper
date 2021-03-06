package controllers

import (
	"encoding/json"
	"github.com/ildarusmanov/eventmapper/models"
	"github.com/ildarusmanov/eventmapper/services"
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
		c.sendJsonResponse(w, false, err.Error())
		return
	}

	if _, err := services.PublishEvent(event, c.mqUrl, rKey); err != nil {
		c.sendJsonResponse(w, false, err.Error())
		return
	}

	c.sendJsonResponse(w, true, "ok")
}

/**
 * send json response
 * @param  w http.ResponseWriter
 * @param isOk bool
 */
func (c *EventController) sendJsonResponse(w http.ResponseWriter, isOk bool, status string) {
	response := models.CreateNewJsonResponse(isOk, status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
