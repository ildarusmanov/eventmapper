package mq

import (
	"github.com/ildarusmanov/eventmapper/tests"
	"testing"
)

func TestChannelConnection(t *testing.T) {
	config := tests.CreateConfig()

	conn, err := CreateNewConnection(config.MqUrl)

	if err != nil {
		t.Error("Can not conenct to RabbitMQ")
	}

	defer conn.Close()

	ch, err := CreateNewChannel(conn)

	if err != nil {
		t.Error("Can not create channel")
	}

	defer ch.Close()
}
