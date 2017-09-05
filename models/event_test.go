package models

import (
	"gopkg.in/validator.v2"
	"testing"
	"time"
)

func TestValidators(t *testing.T) {
	invalidLog := CreateNewEvent()

	if err := validator.Validate(invalidLog); err == nil {
		t.Error("Empty ActionLog model validation: Error expected, but", err, "given")
	}

	validLog := CreateNewEvent()
	validLog.EventName = "authorized"
	validLog.EventTarget = "user"
	validLog.UserId = "some-user-id"
	validLog.CreatedAt = time.Now().Unix()
	validLog.Params = {"key1": "value1"}

	if err := validator.Validate(validLog); err != nil {
		t.Error("Valid ActionLog model validation: Nil expected, but", err, "given")
	}
}