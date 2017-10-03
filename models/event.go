package models

import (
	"encoding/json"
	"eventmapper/mq"
	"gopkg.in/validator.v2"
)

type Event struct {
	EventName   string            `validate:"nonzero,min=1,max=255"`
	EventTarget string            `validate:"min=1,max=255"`
	UserId      string            `validate:"min=1,max=100"`
	CreatedAt   int32             `validate:"nonzero,min=1"`
	Params      map[string]string `validate:"max=200"`
}

func CreateNewEvent() *Event {
	return &Event{}
}

func BuildNewEvent(EventName, EventTarget, UserId string, CreatedAt int32, Params map[string]string) *Event {
	return &Event{EventName, EventTarget, UserId, CreatedAt, Params}
}

func (e *Event) Publish(ch mq.EventChannel, rKey string) error {
	return ch.PublishEvent(e, rKey)
}

func (e *Event) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) GetEventName() string {
	return e.EventName
}

func (e *Event) Validate() error {
	return validator.Validate(e)
}
