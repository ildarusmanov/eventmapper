package models

import (
	"eventmapper/mq"
	"eventmapper/tests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateNewHandlerWithValidOptions(t *testing.T) {
	o := map[string]string{
		"mq_url":       "mq_url",
		"r_key":        "r_key",
		"handler_type": "http_json",
		"url":          "url",
	}

	h, err := CreateNewHandler(o)

	if err != nil {
		t.Error(err)
	}

	if h == nil {
		t.Error("Result should not be nil")
	}
}

func TestCreateNewHandlerWithInvalidOptions(t *testing.T) {
	h, err := CreateNewHandler(map[string]string{})

	if h != nil {
		t.Error("Result should be nil")
	}

	if err == nil {
		t.Error("Should be error")
	}
}

func TestPublishNListening(t *testing.T) {
	// load config
	config := tests.CreateConfig()

	// init amqp connection and channel
	conn, err := mq.CreateNewConnection(config.MqUrl)
	if err != nil {
		t.Error("Can not conenct to RabbitMQ")
	}
	ch, err := mq.CreateNewChannel(conn)
	if err != nil {
		t.Error("Can not create channel")
	}
	// create test http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		r
	}))
	defer ts.Close()

	// init handler params
	o := config.MqHandlers[0]
	o["url"] = ts.URL
	closeCh := make(chan bool)
	// start handler
	go StartHandler(o, closeCh)

	// create and send event
	e := createValidEvent()
	err = e.Publish(ch, "apply")

	// wait and stop listening
	time.Sleep(1 * time.Second)
	close(closeCh)

	if err != nil {
		t.Error("Can not pubslish event", err)
	}
}
