package models

import (
	"eventmapper/mq"
	"encoding/json"
	"gopkg.in/validator.v2"
)

type Event struct {
	EventName    string				`validate:"nonzero,min=1,max=255"`
	EventTarget  string				`validate:"min=1,max=255"`
	UserId       string				`validate:"min=1,max=100"`
	CreatedAt    int64				`validate:"nonzero,min=1"`
	Params		 map[string]string  `validate:"max=100"`
}

func CreateNewEvent() *Event {
	return &Event{}
}

func (e *Event) Publish(mqChannel *mq.Channel, rKey string) error {
	return mqChannel.PublishEvent(e, rKey)
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

