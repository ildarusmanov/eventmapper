package controllers

import (
	"eventmapper/models"
	"eventmapper/tests"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestValidCreateHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	config := tests.CreateConfig()

	controller := CreateNewEventController(config.MqUrl)

	bodyJson := `{
		"source": {
			"source_id": "example.com",
			"source_type": "http",
			"origin": "/api/v1/endpoint/1.json",
			"params": {}
		},
		"target": {
			"target_type": "User",
			"target_id": "1",
			"params": {}
		},
		"event_name": "authorized",
		"user_id": "1",
		"created_at": 1712311,
		"params": {}
	}`

	inBody := bytes.NewBufferString(bodyJson)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "http://127.0.0.1:8000/apply.events.new_user/events", inBody)

	controller.CreateHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonResp := &models.JsonResponse{}

	if err := json.Unmarshal(body, jsonResp); err != nil {
		t.Error("Invalid json response")
	}

	if !jsonResp.IsOk {
		t.Error("Incorrect data")
	}
}

func TestInvalidCreateHandler(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Error", r)
		}
	}()

	config := tests.CreateConfig()

	controller := CreateNewEventController(config.MqUrl)

	bodyJson := `{"EventName": 1, 1712311}`
	inBody := bytes.NewBufferString(bodyJson)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "http://127.0.0.1:8000/apply.events.new_user/events", inBody)

	controller.CreateHandler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonResp := &models.JsonResponse{}

	if err := json.Unmarshal(body, jsonResp); err != nil {
		t.Error("Invalid json response")
	}

	if jsonResp.IsOk {
		t.Error("Incorrect data")
	}
}
