package models

import (
	"eventmapper/mq"
	"gopkg.in/validator.v2"
)

type Event struct {
	EventName    string				`validate:"nonzero,min=1,max=255"`
	EventTarget  string				`validate:"min=1,max=255"`
	UserId       string				`validate:"min=1,max=100"`
	CreatedAt    int64				`validate:"nonzero,min=1"`
	Params		 [string]string     `validate:"max=100"`
}

func CreateNewEvent() *Event {
	return &Event{}
}

func (e *Event) Publish(mqChannel *mq.Channel) error {
	return mqChannel.PublishEvent(event)
}

func (e *Event) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) Validate() error {
	return validator.Validate(e)
}

