package models

import (
	"eventmapper/mq"
	"eventmapper/tests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateNewHandler(t *testing.T) {
	h := CreateNewHandler(
		"mqUrl",
		"rKey",
		"hType",
		map[string]string{"key1": "value1"},
	)

	if h.MqUrl != "mqUrl" || h.RKey != "rKey" || h.HandlerType != "hType" || h.Options["key1"] != "value1" {
		t.Error("incorrect initialization")
	}
}

func TestPublishNListening(t *testing.T) {
	config := tests.CreateConfig()

	conn, err := mq.CreateNewConnection(config.MqUrl)

	if err != nil {
		t.Error("Can not conenct to RabbitMQ")
	}

	ch, err := mq.CreateNewChannel(conn)

	if err != nil {
		t.Error("Can not create channel")
	}

	h := BuildHandlerFromConfig(config.MqHandlers[0])

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	}))

	defer ts.Close()

	h.Options["url"] = ts.URL

	closeCh := make(chan bool)
	errCh := make(chan error)

	go h.StartListening(closeCh, errCh)

	e := createValidEvent()
	err = e.Publish(ch, "apply")

	time.Sleep(1 * time.Second)

	close(closeCh)

	if err != nil {
		t.Error("Can not pubslish event", err)
	}

	if len(errCh) > 0 {
		t.Error("Too much errors")
	}
}
