package models

import (
	"gopkg.in/validator.v2"
	"testing"
	"time"
)

func TestInvalidEventValidators(t *testing.T) {
	invalidEvent := CreateNewEvent()

	if err := validator.Validate(invalidEvent); err == nil {
		t.Error("Empty Event model validation: Error expected, but", err, "given")
	}
}

func TestValidEventValidators(t *testing.T) {
	validEvent := CreateNewEvent()
	validEvent.EventName = "authorized"
	validEvent.EventTarget = "user"
	validEvent.UserId = "some-user-id"
	validEvent.CreatedAt = time.Now().Unix()
	validEvent.Params = map[string]string{"key1": "value1"}

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