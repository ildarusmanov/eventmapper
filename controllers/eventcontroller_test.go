package controllers

import (
	"eventmapper/db"
	"eventmapper/models"
	"eventmapper/tests"

	"io/ioutil"
	"bytes"
	"testing"
	"net/http/httptest"
	"encoding/json"
)

func TestCreateHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	config := tests.CreateConfig()
	mqSession := mq.CreateSession(config.MqUrl)
	defer mqSession.Close()

	controller := CreateNewEventController(dbSession, config)

	bodyJson := "{\"EventName\": \"authorized\", \"EventTarget\": \"user\", \"UserId\": \"some-user-id\", \"CreatedAt\": 1712311}"
	inBody := bytes.NewBufferString(bodyJson)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "http://127.0.0.1:8000/create/apply.events.new_user", inBody)

	controller.CreateHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	event := models.CreateNewEvent()

    if err := json.Unmarshal(body, event); err != nil {
   		t.Error("Invalid json response")
    }

    if event.ActionName != "authorized" {
    	t.Error("Incorrect data")
    }
}
