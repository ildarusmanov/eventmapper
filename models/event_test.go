package models

import (
	"encoding/json"
	"errors"
	"eventmapper/mq"
	"github.com/streadway/amqp"
	"gopkg.in/validator.v2"
	"testing"
	"time"
)

type eventChannelMock struct {
	result error
}

func (ch *eventChannelMock) PublishEvent(event mq.Event, rKey string) error {
	return ch.result
}

func (ch *eventChannelMock) ConsumeEvents(rKeys string) (<-chan amqp.Delivery, error) {
	deliveries := make(chan amqp.Delivery)

	return (<-chan amqp.Delivery)(deliveries), nil
}

func (ch *eventChannelMock) Close() {
	return
}

func createValidEvent() mq.Event {
	validEvent := CreateNewEvent()
	validEvent.EventName = "authorized"
	validEvent.EventTarget = "user"
	validEvent.UserId = "some-user-id"
	validEvent.CreatedAt = time.Now().Unix()
	validEvent.Params = map[string]string{"key1": "value1"}

	return validEvent
}

func TestInvalidEventValidators(t *testing.T) {
	invalidEvent := CreateNewEvent()

	if err := validator.Validate(invalidEvent); err == nil {
		t.Error("Empty Event model validation: Error expected, but", err, "given")
	}
}

func TestValidEventValidators(t *testing.T) {
	validEvent := createValidEvent()

	if err := validator.Validate(validEvent); err != nil {
		t.Error("Valid Event model validation: Nil expected, but", err, "given")
	}
}

func TestGetEventName(t *testing.T) {
	e := CreateNewEvent()
	e.EventName = "authorized"

	if e.GetEventName() != e.EventName {
		t.Error(e.EventName, "expected, but", e.GetEventName(), "given")
	}
}

func TestGetEventBody(t *testing.T) {
	eSource := CreateNewEvent()
	eResult := CreateNewEvent()

	eSource.EventName = "authorized"
	jsonBody, err := eSource.GetBody()

	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(jsonBody, eResult)

	if err != nil {
		t.Error(err)
	}

	if eResult.EventName != eSource.EventName {
		t.Error("Encoded and decoded items are not equal")
	}
}

func TestPublishEvent(t *testing.T) {
	e := createValidEvent()
	result := errors.New("mock")
	ch := &eventChannelMock{result}
	err := e.Publish(ch, "test")

	if err != result {
		t.Error("Unexpected result")
	}
}
