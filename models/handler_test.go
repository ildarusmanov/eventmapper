package models

import (
	"eventmapper/mq"
	"eventmapper/tests"
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

	closeCh := make(chan bool)
	errCh := make(chan error)

	go h.StartListening(closeCh, errCh)

	e := createValidEvent()
	err = e.Publish(ch, "apply")

	time.Sleep(5 * time.Second)
	closeCh <- true

	if err != nil {
		t.Error("Can not pubslish event", err)
	}

	if len(errCh) > 0 {
		t.Error("Too much errors")
	}
}
