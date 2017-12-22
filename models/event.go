package models

import (
	"encoding/json"
	"github.com/ildarusmanov/eventmapper/mq"
	"gopkg.in/validator.v2"
)

type Event struct {
	Source    *EventSource            `validate:"nonzero" json:"source,omitempty"`
	Target    *EventTarget           `validate:"nonzero" json:"target,omitempty"`
	EventName string            `validate:"nonzero,min=1,max=255" json:"event_name,omitempty"`
	UserId      string            `validate:"min=1,max=100" json:"user_id,omitempty"`
	CreatedAt   int32             `validate:"nonzero,min=1" json:"created_at,omitempty"`
	Params      map[string]string `validate:"max=200" json:"params,omitempty"`
}

func CreateNewEvent() *Event {
	return &Event{}
}

func BuildNewEvent(source *EventSource, target *EventTarget, eventName, userId string, createdAt int32, params map[string]string) *Event {
	return &Event{source, target, eventName, userId, createdAt, params}
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
