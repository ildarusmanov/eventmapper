package models

import (
	"eventmapper/mq"
	"eventmapper/tests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
	}))
	defer ts.Close()

	// init handler params
	o := config.MqHandlers[0]
	o["url"] = ts.URL
	closeCh := make(chan bool)
	errCh := make(chan error)
	// start handler
	go StartHandler(o, closeCh, errCh)

	// create and send event
	e := createValidEvent()
	err = e.Publish(ch, "apply")

	// wait and stop listening
	time.Sleep(1 * time.Second)
	close(closeCh)

	if err != nil {
		t.Error("Can not pubslish event", err)
	}

	if len(errCh) > 0 {
		t.Error("Too much errors")
	}
}
