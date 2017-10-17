package models

type EventTarget struct {
	TargetType string `validate:"nonzero,min=1,max=255" json:"target_type,omitempty"`
	TargetId string `validate:"nonzero,min=1,max=500" json:"target_id,omitempty"`
	Params map[string]string `validate:"max=200" json:"params,omitempty"`
}

func CreateNewEventTarget(targetType, targetId string, params map[string]string) *EventTarget {
	return &EventTarget{targetType, targetId, params}
}