package mq

import (
	"eventmapper/tests"
	"testing"
)

func TestConnection(t *testing.T) {
	config := tests.CreateConfig()

	_, err := CreateNewConnection(config.MqUrl)

	if err != nil {
		t.Error("Can not conenct to RabbitMQ")
	}
}